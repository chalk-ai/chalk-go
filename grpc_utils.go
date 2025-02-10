package chalk

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"strings"
	"time"
)

func withChalkInterceptors(serverType string, timeout *time.Duration, interceptors ...connect.Interceptor) connect.Option {
	interceptors = append(interceptors, headerInterceptor(map[string]string{
		HeaderKeyServerType: serverType,
	}))
	if timeout != nil {
		interceptors = append(interceptors, timeoutInterceptor(timeout))
	}
	return connect.WithInterceptors(interceptors...)
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

func timeoutInterceptor(timeout *time.Duration) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			ctx, cancel := internal.GetContextWithTimeout(ctx, timeout)
			defer cancel()
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
