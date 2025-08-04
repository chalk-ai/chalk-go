package auth

import (
	"connectrpc.com/connect"
	"context"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	"sync"
	"time"
)

type TokenRefresher struct {
	mu         *sync.Mutex
	manager    *config.Manager
	AuthClient serverv1connect.AuthServiceClient

	// Initially null
	token *serverv1.GetTokenResponse
}

func NewTokenRefresher(
	ctx context.Context,
	httpClient connect.HTTPClient,
	manager *config.Manager,
) (*TokenRefresher, error) {
	r := &TokenRefresher{
		manager: manager,
		AuthClient: serverv1connect.NewAuthServiceClient(
			httpClient,
			manager.ApiServer.Value,
			connect.WithInterceptors(
				connect.UnaryInterceptorFunc(
					func(next connect.UnaryFunc) connect.UnaryFunc {
						return func(c context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
							req.Header().Set("x-chalk-server", "go-api")
							return next(c, req)
						}
					},
				),
			),
		),
		mu: &sync.Mutex{},
	}
	if _, err := r.GetJWT(ctx, time.Now()); err != nil {
		return nil, errors.Wrap(err, "initializing token refresher")
	}
	return r, nil
}

func (r *TokenRefresher) GetEnvironmentId(override string) string {
	if override != "" {
		return override
	}
	token := r.token
	if token.PrimaryEnvironment != nil && *token.PrimaryEnvironment != "" {
		return *token.PrimaryEnvironment
	}
	return r.manager.EnvironmentId.Value
}

func (r *TokenRefresher) GetJWT(
	ctx context.Context,
	newerThan time.Time,
) (*serverv1.GetTokenResponse, error) {
	to := r.token
	if to != nil && to.GetExpiresAt().AsTime().After(newerThan) {
		return to, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.token != nil && r.token.GetExpiresAt().AsTime().After(newerThan) {
		return r.token, nil
	}
	req := &serverv1.GetTokenRequest{
		ClientId:     string(r.manager.ClientId.Value),
		ClientSecret: string(r.manager.ClientSecret.Value),
		GrantType:    "client_credentials",
		Scope:        nil,
	}
	if r.manager.Scope.Value != "" {
		req.Scope = &r.manager.Scope.Value
	}

	t, err := r.AuthClient.GetToken(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "refreshing token")
	}
	r.token = t.Msg
	return r.token, nil
}

func (r *TokenRefresher) GetQueryServerURL(envOverride string) string {
	token := r.token
	if token == nil {
		return r.manager.ApiServer.Value
	}

	env := envOverride
	if env == "" {
		env = r.manager.EnvironmentId.Value
	}

	if engine, foundEngine := token.Engines[env]; foundEngine && env != "" {
		return engine
	}

	return r.manager.ApiServer.Value
}
