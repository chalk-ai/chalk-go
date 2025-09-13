package auth

import (
	"context"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
)

type Manager struct {
	mu         *sync.Mutex
	manager    *config.Manager
	authClient serverv1connect.AuthServiceClient
	token      *serverv1.GetTokenResponse
}

type Inputs struct {
	Token      *serverv1.GetTokenResponse
	HttpClient connect.HTTPClient
	Manager    *config.Manager
}

func NewManager(ctx context.Context, opts *Inputs) (*Manager, error) {
	r := &Manager{
		manager: opts.Manager,
		authClient: serverv1connect.NewAuthServiceClient(
			opts.HttpClient,
			opts.Manager.ApiServer.Value,
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
		mu:    &sync.Mutex{},
		token: opts.Token,
	}

	if _, err := r.GetJWT(ctx, time.Now()); err != nil {
		return nil, errors.Wrap(err, "initializing token refresher")
	}
	return r, nil
}

func (r *Manager) GetEnvironmentId(override string) string {
	if override != "" {
		return override
	}
	token := r.token
	if token.PrimaryEnvironment != nil && *token.PrimaryEnvironment != "" {
		return *token.PrimaryEnvironment
	}
	return r.manager.EnvironmentId.Value
}

func (r *Manager) GetJWT(
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

	t, err := r.authClient.GetToken(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "refreshing token")
	}
	r.token = t.Msg
	return r.token, nil
}

func (r *Manager) GetQueryServerURL(envOverride string) string {
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
