package chalk

import (
	"connectrpc.com/connect"
	"context"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	"strings"
)

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

func ensureHTTPSPrefix(inputURL string) string {
	if strings.HasPrefix(inputURL, "https://") || strings.HasPrefix(inputURL, "http://") {
		return inputURL
	}
	return "https://" + inputURL
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

func NewAuthClient(httpClient HTTPClient, apiServer string) (serverv1connect.AuthServiceClient, error) {
	return serverv1connect.NewAuthServiceClient(
		httpClient,
		apiServer,
		withChalkInterceptors(
			serverTypeApi,
			headerInterceptor(map[string]string{
				headerKeyServerType: serverTypeApi,
			}),
		),
	), nil
}

func NewQueryClient(httpClient HTTPClient, manager *configManager) (enginev1connect.QueryServiceClient, error) {
	return enginev1connect.NewQueryServiceClient(
		httpClient,
		ensureHTTPSPrefix(manager.getQueryServer()),
		withChalkInterceptors(
			serverTypeEngine,
			makeTokenInterceptor(manager),
			headerInterceptor(map[string]string{
				headerKeyDeploymentType: "engine-grpc",
			}),
		),
		connect.WithGRPC(),
	), nil
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

func MakeGetTokenFunc(logger LeveledLogger, client serverv1connect.AuthServiceClient) func(clientId string, clientSecret string) (*getTokenResult, error) {
	return func(clientId string, clientSecret string) (*getTokenResult, error) {
		return getToken(clientId, clientSecret, logger, client)
	}
}
