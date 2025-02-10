package chalk

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"strings"
	"time"
)

func getContextWithTimeout(ctx context.Context, timeout *time.Duration) context.Context {
	if _, deadlineSet := ctx.Deadline(); !deadlineSet && timeout != nil {
		ctx, _ = context.WithTimeout(ctx, *timeout)
	}
	return ctx
}

func withChalkInterceptors(serverType string, interceptors ...connect.Interceptor) connect.Option {
	return connect.WithInterceptors(
		append(
			interceptors,
			headerInterceptor(map[string]string{
				HeaderKeyServerType: serverType,
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
			req.Header().Set(HeaderKeyEnvironmentId, configManager.environmentId.Value)
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
				HeaderKeyServerType: serverTypeApi,
			}),
		),
	), nil
}

func newInsecureClient() *http.Client {
	// From https://connectrpc.com/docs/go/deployment#h2c
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
}

func newQueryClient(httpClient HTTPClient, manager *configManager, deploymentTag string, queryServerOverride *string) (enginev1connect.QueryServiceClient, error) {
	endpoint := manager.getQueryServer(queryServerOverride)
	if strings.HasPrefix(endpoint, "http://") {
		httpClient = newInsecureClient()
	}
	headers := map[string]string{
		HeaderKeyDeploymentType: "engine-grpc",
	}
	if deploymentTag != "" {
		headers[HeaderKeyDeploymentTag] = deploymentTag
	}
	return enginev1connect.NewQueryServiceClient(
		httpClient,
		ensureHTTPSPrefix(endpoint),
		withChalkInterceptors(
			serverTypeEngine,
			makeTokenInterceptor(manager),
			headerInterceptor(headers),
		),
		connect.WithGRPC(),
	), nil
}
