package chalk

import (
	"connectrpc.com/connect"
	"context"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"net/http"
)

type grpcClientImpl struct {
	GRPCClient
	config *configManager

	branch     string
	logger     LeveledLogger
	httpClient *http.Client

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
	logger := cfg.Logger
	if logger == nil {
		logger = DefaultLeveledLogger
	}
	config.getToken = func(clientId string, clientSecret string) (*getTokenResult, error) {
		return getToken(clientId, clientSecret, logger, authClient)
	}

	// Necessary to get GRPC engines URL
	if err := config.refresh(false); err != nil {
		return nil, errors.Wrap(err, "error fetching initial config")
	}

	queryClient, err := newQueryClient(httpClient, config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating query client")
	}

	return &grpcClientImpl{
		branch:      cfg.Branch,
		httpClient:  httpClient,
		logger:      logger,
		config:      config,
		authClient:  authClient,
		queryClient: queryClient,
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
