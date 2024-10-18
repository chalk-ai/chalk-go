package chalk

import (
	"connectrpc.com/connect"
	"context"
	"github.com/apache/arrow/go/v16/arrow"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
	"reflect"
	"time"
)

var (
	headerKeyDeploymentType = "x-chalk-deployment-type"
	headerKeyEnvironmentId  = "x-chalk-env-id"
	headerKeyServerType     = "x-chalk-server"
	//headerKeyTraceId        = "x-chalk-trace-id"
	headerKeyDeploymentTag = "x-chalk-deployment-tag"

	serverTypeApi    = "go-api"
	serverTypeEngine = "engine"
)

type grpcClientImpl struct {
	GRPCClient
	config *configManager

	branch      string
	queryServer *string
	logger      LeveledLogger
	httpClient  *http.Client

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newGrpcClient(cfg GRPCClientConfig) (*grpcClientImpl, error) {
	config, err := newConfigManager(cfg.ApiServer, cfg.ClientId, cfg.ClientSecret, cfg.EnvironmentId, cfg.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}
	httpClient := http.DefaultClient
	authClient, err := newAuthClient(httpClient, config.apiServer.Value)
	if err != nil {
		return nil, errors.Wrap(err, "error creating auth client")
	}
	config.getToken = func(clientId string, clientSecret string) (*getTokenResult, error) {
		return getToken(clientId, clientSecret, config.logger, authClient)
	}

	// Necessary to get GRPC engines URL
	if err := config.refresh(false); err != nil {
		return nil, errors.Wrap(err, "error fetching initial config")
	}

	var queryServer *string
	if cfg.QueryServer != "" {
		queryServer = ptr.Ptr(cfg.QueryServer)
	}

	queryClient, err := newQueryClient(httpClient, config, cfg.DeploymentTag, queryServer)
	if err != nil {
		return nil, errors.Wrap(err, "error creating query client")
	}

	return &grpcClientImpl{
		branch:      cfg.Branch,
		httpClient:  httpClient,
		logger:      config.logger,
		config:      config,
		authClient:  authClient,
		queryClient: queryClient,
		queryServer: queryServer,
	}, nil
}

func getToken(clientId string, clientSecret string, logger LeveledLogger, client serverv1connect.AuthServiceClient) (*getTokenResult, error) {
	logger.Debugf("Getting new token via gRPC")
	authRequest := connect.NewRequest(
		&serverv1.GetTokenRequest{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
		},
	)
	token, err := client.GetToken(context.Background(), authRequest)
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
	bulkInputs := make(map[string]any)
	for k, singleValue := range args.underlying.inputs {
		slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(singleValue)), 1, 1)
		slice.Index(0).Set(reflect.ValueOf(singleValue))
		bulkInputs[k] = slice.Interface()
	}
	args.underlying.inputs = bulkInputs

	bulkRes, err := c.OnlineQueryBulk(ctx, args)
	if err != nil {
		// intentionally don't wrap, original error is good enough
		return nil, err
	}

	scalarsTable, err := internal.ConvertBytesToTable(bulkRes.GetScalarsData())
	if err != nil {
		return nil, errors.Wrap(err, "error converting scalars data to table")
	}

	rows, err := internal.ExtractFeaturesFromTable(scalarsTable)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting features from scalars table")
	}

	features := make(map[string]*commonv1.FeatureResult)
	if len(rows) != 1 {
		return nil, errors.Newf(
			"expected 1 row from scalars table, got %d",
			len(rows),
		)
	}
	for fqn, value := range rows[0] {
		if reflect.TypeOf(value) == reflect.TypeOf(time.Time{}) {
			value = value.(time.Time).Format(time.RFC3339)
		}
		newValue, err := structpb.NewValue(value)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"error converting value for feature '%s' from `any` to `structpb.Value`",
				fqn,
			)
		}
		features[fqn] = &commonv1.FeatureResult{
			Field: fqn,
			Value: newValue,
		}
	}

	for fqn, tableBytes := range bulkRes.GetGroupsData() {
		table, err := internal.ConvertBytesToTable(tableBytes)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"error converting bytes for feature '%s' to table",
				fqn,
			)
		}

		rowsHm, err := internal.ExtractFeaturesFromTable(table)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"error extracting features from has-many table for feature '%s'",
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
				"error converting has-many result for feature '%s' to `structpb.Value`",
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
		return nil, errors.Wrap(err, "error converting online query params to proto")
	}
	req := connect.NewRequest(paramsProto)
	res, err := c.queryClient.OnlineQueryBulk(ctx, req)
	if err != nil {
		return nil, wrapClientError(err, "error executing online query")
	}
	return res.Msg, nil
}

func (c *grpcClientImpl) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (*commonv1.UploadFeaturesBulkResponse, error) {
	inputsConverted, err := getConvertedInputsMap(args.Inputs)
	if err != nil {
		return nil, wrapClientError(err, "error converting inputs map")
	}
	inputsFeather, err := internal.InputsToArrowBytes(inputsConverted)
	if err != nil {
		return nil, wrapClientError(err, "error serializing inputs as feather")
	}

	req := connect.NewRequest(&commonv1.UploadFeaturesBulkRequest{
		InputsFeather: inputsFeather,
		BodyType:      commonv1.FeatherBodyType_FEATHER_BODY_TYPE_TABLE,
	})

	res, err := c.queryClient.UploadFeaturesBulk(ctx, req)
	if err != nil {
		return nil, wrapClientError(err, "error making update aggregates request")
	}
	return res.Msg, nil
}

func (c *grpcClientImpl) GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, wrapClientError(err, "error making get aggregates request")
	}

	return res.Msg, err
}

func (c *grpcClientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, wrapClientError(err, "error making plan aggregate backfill request")
	}
	return res.Msg, err
}

func (c *grpcClientImpl) GetToken() (*TokenResult, error) {
	res, err := c.config.getToken(c.config.clientId.Value, c.config.clientSecret.Value)
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
