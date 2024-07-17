package chalk

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	enginev1 "github.com/chalk-ai/chalk-go/gen/chalk/engine/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"net/http"
	"time"
)

var (
	headerKeyDeploymentType = "x-chalk-deployment-type"
	headerKeyEnvironmentId  = "x-chalk-env-id"
	headerKeyServerType     = "x-chalk-server"

	serverTypeApi    = "go-api"
	serverTypeEngine = "engine"
)

type clientGrpc struct {
	config *configManager

	branch     string
	logger     LeveledLogger
	httpClient *http.Client

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newClientGrpc(cfg ClientConfig) (*clientGrpc, error) {
	resolved, err := getResolvedConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}
	var logger LeveledLogger
	if cfg.Logger == nil {
		logger = DefaultLeveledLogger
	}
	client := &clientGrpc{
		branch:     cfg.Branch,
		httpClient: http.DefaultClient,
		logger:     logger,

		config: &configManager{
			apiServer:          resolved.ApiServer,
			clientId:           resolved.ClientId,
			clientSecret:       resolved.ClientSecret,
			environmentId:      resolved.EnvironmentId,
			initialEnvironment: resolved.EnvironmentId,
		},
	}
	if err := client.init(); err != nil {
		return nil, errors.Wrap(err, "error initializing gRPC service clients")
	}

	return client, nil
}

func (c *clientGrpc) init() error {
	authClient, err := c.NewAuthClient()
	if err != nil {
		return errors.Wrap(err, "error creating auth client")
	}
	c.authClient = authClient

	c.config.getToken = c.getToken
	// Necessary to get GRPC engines URL
	if err := c.config.refresh(false); err != nil {
		return errors.Wrap(err, "error fetching initial config")
	}

	queryClient, err := c.NewQueryClient()
	if err != nil {
		return errors.Wrap(err, "error creating query client")
	}
	c.queryClient = queryClient

	return nil
}

func withChalkInterceptors(serverType string, interceptors ...connect.Interceptor) connect.Option {
	return connect.WithInterceptors(
		append(
			interceptors,
			headerInterceptor(map[string]string{
				headerKeyServerType: serverType,
			}),
		)...,
	)
}

func (c *clientGrpc) NewAuthClient() (serverv1connect.AuthServiceClient, error) {
	return serverv1connect.NewAuthServiceClient(
		c.httpClient,
		c.config.apiServer.Value,
		withChalkInterceptors(
			serverTypeApi,
			headerInterceptor(map[string]string{
				headerKeyServerType: "go-api",
			}),
		),
	), nil
}

func (c *clientGrpc) NewQueryClient() (enginev1connect.QueryServiceClient, error) {
	headers := map[string]string{
		headerKeyDeploymentType: "engine-grpc",
	}

	endpoint, ok := c.config.engines[c.config.environmentId.Value]
	if !ok {
		return nil, errors.Newf(
			"no engine found for environment '%s' - engine map keys: '%s'",
			c.config.environmentId.Value,
			lo.Keys(c.config.engines),
		)
	}

	return enginev1connect.NewQueryServiceClient(
		c.httpClient,
		"https://"+endpoint,
		withChalkInterceptors(
			serverTypeEngine,
			c.tokenInterceptor(),
			headerInterceptor(headers),
		),
		connect.WithGRPC(),
	), nil
}

func headerInterceptor(headers map[string]string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			for k, v := range headers {
				req.Header().Set(k, v)
			}
			return next(ctx, req)
		}
	}
}

func (c *clientGrpc) getToken() (*getTokenResult, error) {
	c.logger.Debugf("Getting new token via gRPC")
	authRequest := connect.NewRequest(
		&serverv1.GetTokenRequest{
			ClientId:     c.config.clientId.Value,
			ClientSecret: c.config.clientSecret.Value,
			GrantType:    "client_credentials",
		},
	)
	token, err := c.authClient.GetToken(context.Background(), authRequest)
	if err != nil {
		c.logger.Debugf("Failed to get a new token: %s", err.Error())
		return nil, err
	}
	return &getTokenResult{
		ValidUntil:         token.Msg.GetExpiresAt().AsTime(),
		AccessToken:        token.Msg.GetAccessToken(),
		PrimaryEnvironment: token.Msg.GetPrimaryEnvironment(),
		Engines:            token.Msg.GetGrpcEngines(),
	}, nil
}

func (c *clientGrpc) tokenInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if err := c.config.refresh(false); err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			req.Header().Set(headerKeyEnvironmentId, c.config.environmentId.Value)
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", c.config.jwt.Token))
			return next(ctx, req)
		}
	}
}

func (c *clientGrpc) OnlineQuery(args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	return OnlineQueryResult{}, errors.New("not implemented")
}

func (c *clientGrpc) OnlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	inputsFeather, err := internal.InputsToArrowBytes(args.underlying.inputs)
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "error serializing inputs as feather")
	}
	outputs := lo.Map(args.underlying.outputs, func(v string, _ int) *commonv1.OutputExpr {
		return &commonv1.OutputExpr{
			Expr: &commonv1.OutputExpr_FeatureFqn{
				FeatureFqn: v,
			},
		}
	})
	staleness := lo.MapValues(args.underlying.staleness, func(v time.Duration, k string) string {
		return internal.FormatBucketDuration(int(v.Seconds()))
	})

	pingRes, err := c.queryClient.Ping(context.Background(), connect.NewRequest(&enginev1.PingRequest{Num: 5525}))
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "error pinging server")
	}
	fmt.Println("Ping response:", pingRes.Msg.Num)

	req := connect.NewRequest(
		&commonv1.OnlineQueryBulkRequest{
			InputsFeather: inputsFeather,
			Outputs:       outputs,
			Staleness:     staleness,
			Now:           nil,
			Context: &commonv1.OnlineQueryContext{
				Environment:          args.underlying.EnvironmentId,
				Tags:                 args.underlying.Tags,
				DeploymentId:         lo.ToPtr(args.underlying.PreviewDeploymentId),
				BranchId:             args.underlying.BranchId,
				CorrelationId:        lo.ToPtr(args.underlying.CorrelationId),
				QueryName:            lo.ToPtr(args.underlying.QueryName),
				RequiredResolverTags: nil,
				QueryNameVersion:     nil,
				Options:              nil,
			},
			ResponseOptions: &commonv1.OnlineQueryResponseOptions{
				IncludeMeta:     args.underlying.IncludeMeta,
				Metadata:        args.underlying.Meta,
				EncodingOptions: nil,
				Explain:         nil,
			},
		},
	)

	res, err := c.queryClient.OnlineQueryBulk(context.Background(), req)
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "error executing online query")
	}

	if len(res.Msg.Errors) > 0 {
		var serverErrs []ServerError

		for _, e := range res.Msg.Errors {
			serverErr, err := serverErrorFromProto(e)
			if err != nil {
				return OnlineQueryBulkResult{}, errors.Wrap(err, "error converting server error")
			}
			serverErrs = append(serverErrs, *serverErr)
		}
		return OnlineQueryBulkResult{}, &ErrorResponse{
			ServerErrors: serverErrs,
		}
	}

	scalars, err := internal.ConvertBytesToTable(res.Msg.GetScalarsData())
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "error deserializing scalars table")
	}

	groups := make(map[string]arrow.Table)
	for k, v := range res.Msg.GetGroupsData() {
		g, err := internal.ConvertBytesToTable(v)
		if err != nil {
			return OnlineQueryBulkResult{}, errors.Wrapf(
				err,
				"error deserializing has-many table for feature '%s'",
				k,
			)
		}
		groups[k] = g
	}

	metaRaw := res.Msg.GetResponseMeta()

	var executionDuration float64
	if metaRaw.ExecutionDuration != nil {
		executionDuration = metaRaw.ExecutionDuration.AsDuration().Seconds()
	}

	var queryTimestamp *time.Time
	if metaRaw.QueryTimestamp != nil {
		queryTimestamp = lo.ToPtr(metaRaw.QueryTimestamp.AsTime())
	}

	return OnlineQueryBulkResult{
		ScalarsTable: scalars,
		GroupsTables: groups,
		Meta: &QueryMeta{
			ExecutionDurationS: executionDuration,
			DeploymentId:       metaRaw.DeploymentId,
			EnvironmentId:      metaRaw.EnvironmentId,
			EnvironmentName:    metaRaw.EnvironmentName,
			QueryId:            metaRaw.QueryId,
			QueryTimestamp:     queryTimestamp,
			QueryHash:          metaRaw.QueryHash,
		},
	}, nil

}

func (c *clientGrpc) UploadFeatures(args UploadFeaturesParams) (UploadFeaturesResult, error) {
	return UploadFeaturesResult{}, errors.New("not implemented")
}

func (c *clientGrpc) OfflineQuery(args OfflineQueryParamsComplete) (Dataset, error) {
	return Dataset{}, errors.New("not implemented")
}

func (c *clientGrpc) TriggerResolverRun(args TriggerResolverRunParams) (TriggerResolverRunResult, error) {
	return TriggerResolverRunResult{}, errors.New("not implemented")
}

func (c *clientGrpc) GetRunStatus(args GetRunStatusParams) (GetRunStatusResult, error) {
	return GetRunStatusResult{}, errors.New("not implemented")
}