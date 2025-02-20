package chalk

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	"github.com/apache/arrow/go/v16/arrow"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
	"google.golang.org/protobuf/types/known/structpb"
	"net"
	"net/http"
	"strings"
	"time"
)

type grpcClientImpl struct {
	GRPCClient
	config *configManager

	branch        string
	queryServer   *string
	resourceGroup *string
	logger        LeveledLogger
	httpClient    HTTPClient
	timeout       *time.Duration

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newGrpcClient(ctx context.Context, configs ...*GRPCClientConfig) (*grpcClientImpl, error) {
	var cfg *GRPCClientConfig
	if len(configs) == 0 {
		cfg = &GRPCClientConfig{}
	} else if len(configs) == 1 {
		cfg = configs[len(configs)-1]
	} else {
		return nil, errors.Newf("expected at most one GRPCClientConfig, got %d", len(configs))
	}

	config, err := newConfigManager(cfg.ApiServer, cfg.ClientId, cfg.ClientSecret, cfg.EnvironmentId, cfg.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "getting resolved config")
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var timeout *time.Duration
	if cfg.Timeout != 0 { // If unspecified (zero value)
		timeout = &cfg.Timeout
	}

	authInterceptors := []connect.Interceptor{
		headerInterceptor(map[string]string{
			HeaderKeyServerType: serverTypeApi,
		}),
	}
	if timeout != nil {
		authInterceptors = append(authInterceptors, timeoutInterceptor(timeout))
	}
	authClient := serverv1connect.NewAuthServiceClient(
		httpClient,
		config.apiServer.Value,
		connect.WithInterceptors(authInterceptors...),
	)

	config.getToken = func(ctx context.Context, clientId string, clientSecret string) (*getTokenResult, error) {
		return getToken(ctx, clientId, clientSecret, config.logger, authClient)
	}

	// Necessary to get GRPC engines URL
	if err := config.refresh(ctx, false); err != nil {
		return nil, errors.Wrap(err, "fetching initial config")
	}

	var queryServer *string
	if cfg.QueryServer != "" {
		queryServer = &cfg.QueryServer
	}

	var resourceGroup *string
	if cfg.ResourceGroup != "" {
		resourceGroup = &cfg.ResourceGroup
	}

	resolvedQueryServer := config.getQueryServer(queryServer)
	if strings.HasPrefix(resolvedQueryServer, "http://") {
		// Unsecured client
		// From https://connectrpc.com/docs/go/deployment#h2c
		httpClient = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}
	}

	headers := map[string]string{
		HeaderKeyDeploymentType: "engine-grpc",
		HeaderKeyServerType:     serverTypeEngine,
	}
	if cfg.DeploymentTag != "" {
		headers[HeaderKeyDeploymentTag] = cfg.DeploymentTag
	}
	engineInterceptors := []connect.Interceptor{
		makeTokenInterceptor(config),
		headerInterceptor(headers),
	}
	if timeout != nil {
		engineInterceptors = append(engineInterceptors, timeoutInterceptor(timeout))
	}

	queryClient := enginev1connect.NewQueryServiceClient(
		httpClient,
		ensureHTTPSPrefix(resolvedQueryServer),
		connect.WithInterceptors(engineInterceptors...),
		connect.WithGRPC(),
	)

	return &grpcClientImpl{
		branch:        cfg.Branch,
		httpClient:    httpClient,
		logger:        config.logger,
		config:        config,
		authClient:    authClient,
		queryClient:   queryClient,
		queryServer:   queryServer,
		resourceGroup: resourceGroup,
		timeout:       timeout,
	}, nil
}

func getToken(ctx context.Context, clientId string, clientSecret string, logger LeveledLogger, client serverv1connect.AuthServiceClient) (*getTokenResult, error) {
	logger.Debugf("Getting new token via gRPC")
	authRequest := connect.NewRequest(
		&serverv1.GetTokenRequest{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
		},
	)

	token, err := client.GetToken(ctx, authRequest)
	if err != nil {
		logger.Debugf("Failed to get a new token: %s", err.Error())
		return nil, err
	}

	expiresAt := token.Msg.GetExpiresAt()
	if token.Msg.GetExpiresAt() == nil {
		return nil, errors.New("token has no expiration date")
	}

	return &getTokenResult{
		ValidUntil:         expiresAt.AsTime(),
		AccessToken:        token.Msg.GetAccessToken(),
		PrimaryEnvironment: token.Msg.GetPrimaryEnvironment(),
		Engines:            token.Msg.GetGrpcEngines(),
	}, nil
}

func (c *grpcClientImpl) OnlineQuery(ctx context.Context, args OnlineQueryParamsComplete) (*commonv1.OnlineQueryResponse, error) {
	newInputs, err := internal.SingleInputsToBulkInputs(args.underlying.inputs)
	if err != nil {
		return nil, errors.Wrap(err, "converting inputs to bulk inputs")
	}
	args.underlying.inputs = newInputs

	bulkRes, err := c.OnlineQueryBulk(ctx, args)
	if err != nil {
		// intentionally don't wrap, original error is good enough
		return nil, err
	}

	features := make(map[string]*commonv1.FeatureResult)
	if len(bulkRes.GetScalarsData()) > 0 {
		scalarsTable, err := internal.ConvertBytesToTable(bulkRes.GetScalarsData())
		if err != nil {
			return nil, errors.Wrap(err, "converting scalars data to table")
		}

		// Need to obtain time.Time values as string because structpb.NewValue does not support time.Time.
		rows, meta, err := internal.ExtractFeaturesFromTable(scalarsTable, true)
		if err != nil {
			return nil, errors.Wrap(err, "extracting features from scalars table")
		}

		if len(rows) == 1 {
			var rowMeta map[string]internal.FeatureMeta
			if len(meta) > 0 {
				if len(meta) != 1 {
					return nil, errors.Newf("expected exactly one metadata row, found %v", meta)
				}
				rowMeta = meta[0]
			}
			for fqn, value := range rows[0] {
				// Needed to obtain time.Time values as string because structpb.NewValue does not support time.Time.
				newValue, err := structpb.NewValue(value)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"converting value for feature '%s' from `any` to `structpb.Value`",
						fqn,
					)
				}
				featureRes := commonv1.FeatureResult{
					Field: fqn,
					Value: newValue,
				}
				features[fqn] = &featureRes
				if rowMeta != nil {
					featureMeta, ok := rowMeta[fqn]
					if !ok {
						// Features such as has-many features do not have a metadata column.
						continue
					}
					if featureMeta.Pkey != nil {
						val, err := structpb.NewValue(featureMeta.Pkey)
						if err != nil {
							return nil, errors.Wrapf(
								err,
								"converting primary key for feature '%s' to `structpb.Value`",
								fqn,
							)
						}
						featureRes.Pkey = val
					}

					featureRes.Meta = &commonv1.FeatureMeta{}
					if featureMeta.ResolverFqn != nil {
						featureRes.Meta.ChosenResolverFqn = *featureMeta.ResolverFqn
					}
					if featureMeta.SourceType != nil && *featureMeta.SourceType == string(internal.SourceTypeOnlineStore) {
						featureRes.Meta.CacheHit = true
					}
				}
			}
		}
	}

	for fqn, tableBytes := range bulkRes.GetGroupsData() {
		table, err := internal.ConvertBytesToTable(tableBytes)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"converting bytes for feature '%s' to table",
				fqn,
			)
		}

		rowsHm, _, err := internal.ExtractFeaturesFromTable(table, false)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"extracting features from has-many table for feature '%s'",
				fqn,
			)
		}

		colNames := colls.Map(
			table.Schema().Fields(),
			func(f arrow.Field) string {
				return f.Name
			},
		)
		colValues := make([][]any, 0, len(rowsHm))
		for _, col := range colNames {
			colValues = append(
				colValues,
				colls.Map(rowsHm, func(row map[string]any) any {
					return row[col]
				}),
			)
		}
		hmResult := map[string]any{
			"columns": colNames,
			"values":  colValues,
		}
		hmProto, err := structpb.NewValue(hmResult)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"converting has-many result for feature '%s' to `structpb.Value`",
				fqn,
			)
		}

		features[fqn] = &commonv1.FeatureResult{
			Field: fqn,
			Value: hmProto,
		}
	}

	return &commonv1.OnlineQueryResponse{
		Data: &commonv1.OnlineQueryResult{
			Results: colls.Values(features),
		},
		Errors:       bulkRes.Errors,
		ResponseMeta: bulkRes.ResponseMeta,
	}, nil
}

func (c *grpcClientImpl) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (*commonv1.OnlineQueryBulkResponse, error) {
	paramsProto, err := convertOnlineQueryParamsToProto(&args.underlying)
	if err != nil {
		return nil, errors.Wrap(err, "converting online query params to proto")
	}
	req := connect.NewRequest(paramsProto)
	if args.underlying.ResourceGroup != "" {
		req.Header().Set(HeaderKeyResourceGroup, args.underlying.ResourceGroup)
	} else if c.resourceGroup != nil {
		req.Header().Set(HeaderKeyResourceGroup, *c.resourceGroup)
	}
	res, err := c.queryClient.OnlineQueryBulk(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "executing online query")
	}
	return res.Msg, nil
}

func (c *grpcClientImpl) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (*commonv1.UploadFeaturesBulkResponse, error) {
	inputsConverted, err := getConvertedInputsMap(args.Inputs)
	if err != nil {
		return nil, errors.Wrap(err, "converting inputs map")
	}
	inputsFeather, err := internal.InputsToArrowBytes(inputsConverted)
	if err != nil {
		return nil, errors.Wrap(err, "serializing inputs as feather")
	}

	req := connect.NewRequest(&commonv1.UploadFeaturesBulkRequest{
		InputsFeather: inputsFeather,
		BodyType:      commonv1.FeatherBodyType_FEATHER_BODY_TYPE_TABLE,
	})

	res, err := c.queryClient.UploadFeaturesBulk(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making update aggregates request")
	}
	return res.Msg, nil
}

func (c *grpcClientImpl) GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making get aggregates request")
	}

	return res.Msg, err
}

func (c *grpcClientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "making plan aggregate backfill request")
	}
	return res.Msg, err
}

func (c *grpcClientImpl) GetToken(ctx context.Context) (*TokenResult, error) {
	res, err := c.config.getToken(ctx, c.config.clientId.Value, c.config.clientSecret.Value)
	if err != nil {
		return nil, err
	}
	return &TokenResult{
		AccessToken:        res.AccessToken,
		ValidUntil:         res.ValidUntil,
		PrimaryEnvironment: res.PrimaryEnvironment,
		Engines:            res.Engines,
	}, nil
}
