package chalk

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
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

func makeTokenInterceptor(configManager *configManager) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if err := configManager.refresh(false); err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			req.Header().Set(headerKeyEnvironmentId, configManager.environmentId.Value)
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", configManager.jwt.Token))
			return next(ctx, req)
		}
	}
}

func newAuthClient(httpClient HTTPClient, apiServer string) (serverv1connect.AuthServiceClient, error) {
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

func newQueryClient(httpClient HTTPClient, manager *configManager) (enginev1connect.QueryServiceClient, error) {
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
