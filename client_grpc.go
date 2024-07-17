package chalk

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"net/http"
	"time"
)

var (
	headerDeploymentType = "X-Chalk-Deployment-Type"
	headerEnvironmentId  = "X-Chalk-Env-Id"
	headerServerType     = "X-Chalk-Server"
)

type clientGrpc struct {
	apiServer     auth2.SourcedConfig
	clientId      auth2.SourcedConfig
	environmentId auth2.SourcedConfig
	branch        string

	httpClient  *http.Client
	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient

	clientSecret auth2.SourcedConfig
	jwt          *auth2.JWT
	logger       LeveledLogger
}

func newClientGrpc(cfg ClientConfig) (*clientGrpc, error) {
	resolved, err := getResolvedConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}
	client := &clientGrpc{
		apiServer:     resolved.ApiServer,
		clientId:      resolved.ClientId,
		clientSecret:  resolved.ClientSecret,
		environmentId: resolved.EnvironmentId,
		branch:        cfg.Branch,
		httpClient:    http.DefaultClient,
	}
	return client, nil
}

func (c *clientGrpc) initClients() error {
	c.authClient = c.NewAuthClient()
	c.queryClient = c.NewQueryClient()
	return nil
}

func (c *clientGrpc) NewAuthClient() serverv1connect.AuthServiceClient {
	return serverv1connect.NewAuthServiceClient(
		c.httpClient,
		c.apiServer.Value,
		connect.WithInterceptors(
			headerInterceptor(map[string]string{
				headerServerType: "go-api",
			}),
		),
	)
}

func (c *clientGrpc) NewQueryClient() enginev1connect.QueryServiceClient {
	headers := map[string]string{
		headerDeploymentType: "engine-grpc",
	}
	return enginev1connect.NewQueryServiceClient(
		c.httpClient,
		c.apiServer.Value,
		connect.WithInterceptors(
			c.tokenInterceptor(c.authClient),
			headerInterceptor(headers),
		),
	)
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

func (c *clientGrpc) tokenInterceptor(
	authClient serverv1connect.AuthServiceClient,
) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if c.jwt == nil || !c.jwt.IsValid() {
				c.logger.Debugf("Getting new token")
				authRequest := connect.NewRequest(
					&serverv1.GetTokenRequest{
						ClientId:     c.clientId.Value,
						ClientSecret: c.clientSecret.Value,
						GrantType:    "client_credentials",
					},
				)
				req.Header().Set(headerEnvironmentId, c.environmentId.Value)
				newToken, err := authClient.GetToken(ctx, authRequest)
				if err != nil {
					c.logger.Debugf("Failed to get a new token: %s", err.Error())
					return nil, err
				}
				c.jwt = &auth2.JWT{
					Token:      newToken.Msg.AccessToken,
					ValidUntil: newToken.Msg.ExpiresAt.AsTime(),
				}
			}

			req.Header().Set(headerEnvironmentId, c.environmentId.Value)
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.Token))
			return next(ctx, req)
		}
	}
}

func (c *clientGrpc) OnlineQuery(args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	return OnlineQueryResult{}, errors.New("not implemented")
}

func (c *clientGrpc) OnlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	inputsFeather, err := args.ToBytes()
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
				Explain:     nil,
				IncludeMeta: args.underlying.IncludeMeta,
				Metadata:    args.underlying.Meta,
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
