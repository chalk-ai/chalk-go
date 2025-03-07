package chalk

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
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

type GRPCOnlineQueryBulkResult struct {
	RawResponse *commonv1.OnlineQueryBulkResponse
}

type GRPCFeatureResult struct {
	Field string
	Value any
	Meta  *FeatureMeta
}

type FeatureMeta struct {
	Pkey        any
	ResolverFqn string
	SourceType  string
	SourceId    string
}

func (r *GRPCOnlineQueryBulkResult) GetRow(rowIndex int) ([]GRPCFeatureResult, error) {
	if len(r.RawResponse.GetScalarsData()) == 0 {
		return nil, errors.New("results table empty, either the query has errors or the data is malformed")
	}

	scalarsTable, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData())
	if err != nil {
		return nil, errors.Wrap(err, "converting scalars data to table")
	}

	results := make([]GRPCFeatureResult, 0, scalarsTable.NumCols())
	rows, meta, err := internal.ExtractFeaturesFromTable(scalarsTable, false)
	if err != nil {
		return nil, errors.Wrap(err, "extracting features from scalars table")
	}

	var rowMeta map[string]internal.FeatureMeta
	if len(meta) != len(rows) {
		return nil, errors.New("metadata length does not match rows length")
	}
	rowMeta = meta[rowIndex]
	for fqn, value := range rows[rowIndex] {
		featureRes := GRPCFeatureResult{
			Field: fqn,
			Value: value,
		}
		if rowMeta != nil {
			internalMeta, ok := rowMeta[fqn]
			if !ok {
				// Features such as has-many features do not have a metadata column.
				continue
			}
			publicMeta := &FeatureMeta{
				Pkey: internalMeta.Pkey,
			}
			if internalMeta.ResolverFqn != nil {
				publicMeta.ResolverFqn = *internalMeta.ResolverFqn
			}
			if internalMeta.SourceType != nil {
				publicMeta.SourceType = *internalMeta.SourceType
			}
			if internalMeta.SourceId != nil {
				publicMeta.SourceId = *internalMeta.SourceId
			}
		}
		results = append(results, featureRes)
	}
	return results, nil
}

func (r *GRPCOnlineQueryBulkResult) UnmarshalInto(resultHolders any) error {
	scalars, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData())
	if err != nil {
		return errors.Wrap(err, "deserializing scalars table")
	}
	return internal.UnmarshalTableInto(scalars, resultHolders)
}

func (c *grpcClientImpl) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (*GRPCOnlineQueryBulkResult, error) {
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

	return &GRPCOnlineQueryBulkResult{
		RawResponse: res.Msg,
	}, nil
}

type GRPCUpdateAggregatesResult struct {
	RawResponse *commonv1.UploadFeaturesBulkResponse
}

func (c *grpcClientImpl) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (*GRPCUpdateAggregatesResult, error) {
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

	return &GRPCUpdateAggregatesResult{
		RawResponse: res.Msg,
	}, nil
}

type GRPCGetAggregatesResult struct {
	RawResponse *aggregatev1.GetAggregatesResponse
}

func (c *grpcClientImpl) GetAggregates(ctx context.Context, features []string) (*GRPCGetAggregatesResult, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making get aggregates request")
	}

	return &GRPCGetAggregatesResult{
		RawResponse: res.Msg,
	}, nil
}

type GRPCPlanAggregateBackfillResult struct {
	RawResponse *aggregatev1.PlanAggregateBackfillResponse
}

func (c *grpcClientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*GRPCPlanAggregateBackfillResult, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "making plan aggregate backfill request")
	}
	return &GRPCPlanAggregateBackfillResult{
		RawResponse: res.Msg,
	}, nil
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
