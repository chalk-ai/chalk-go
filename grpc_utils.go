package chalk

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go/auth"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"time"
)

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

func timeoutInterceptor(clientLevelTimeout *time.Duration) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			ctx, cancel := internal.GetContextWithTimeout(ctx, clientLevelTimeout)
			defer cancel()
			return next(ctx, req)
		}
	}
}

func makeTokenInterceptor(tm *auth.TokenRefresher) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			token, err := tm.GetJWT(ctx, time.Now().Add(time.Minute))
			if err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			req.Header().Set("x-chalk-env-id", tm.GetEnvironmentId(""))
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
			return next(ctx, req)
		}
	}
}
