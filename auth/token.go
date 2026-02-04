package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
)

type Manager struct {
	mu         *sync.Mutex
	config     *config.Manager
	authClient serverv1connect.AuthServiceClient
	token      *serverv1.GetTokenResponse
}

type Inputs struct {
	// Token overrides the default token, if provided. It's unlikely you'll want to provide this.
	Token *serverv1.GetTokenResponse

	// HttpClient is used as the underlying http.client. Connect provides an interface for abstracting over the
	// standard library version of the auth client
	HttpClient connect.HTTPClient

	// Manager holds the credentials for this client to use. Non-optional.
	Config *config.Manager

	Timeout *time.Duration

	// SkipEnvironmentNameMapping controls whether to skip validating and mapping
	// environment names to IDs. If true, the EnvironmentId will be used verbatim.
	SkipEnvironmentNameMapping bool

	// SkipEngineMapping controls whether to skip setting the query server based
	// on the token's engine maps. If true, the query server will not be
	// automatically resolved from the token.
	SkipEngineMapping bool
}

func cleanEnvironmentId(
	provided config.SourcedConfig[string],
	token *serverv1.GetTokenResponse,
) (config.SourcedConfig[string], error) {
	if provided.Value == "" && token.PrimaryEnvironment != nil {
		return config.NewFromToken(
			*token.PrimaryEnvironment,
			fmt.Sprintf("Default environment %q", token.EnvironmentIdToName[*token.PrimaryEnvironment]),
		), nil
	} else if provided.Value == "" {
		var availableEnvironments []string
		for id, name := range token.EnvironmentIdToName {
			availableEnvironments = append(availableEnvironments, fmt.Sprintf("%s (%s)", name, id))
		}
		return config.SourcedConfig[string]{
				Value:  "",
				Source: "empty",
				Kind:   config.EmptySourceKind,
			}, errors.Newf(
				"environment was not specified, and the token did not include a primary environment; all available environments are %s; primary environment was %q",
				strings.Join(availableEnvironments, ", "),
				token.PrimaryEnvironment,
			)
	} else if _, ok := token.EnvironmentIdToName[provided.Value]; ok {
		return provided, nil
	} else {
		// The provided environment isn't valid, but it may be a name
		for envId := range token.EnvironmentIdToName {
			if strings.EqualFold(envId, provided.Value) {
				return provided.
					WithValue(envId).
					WithSourceF("%s (transformed from name %q)", provided.Source, envId), nil
			}
		}
		for envId, name := range token.EnvironmentIdToName {
			if strings.EqualFold(name, provided.Value) {
				return provided.
					WithValue(envId).
					WithSourceF("%s (transformed from name %q)", provided.Source, envId), nil
			}
		}
		var available []string
		for id, name := range token.EnvironmentIdToName {
			available = append(available, fmt.Sprintf("%s (%s)", name, id))
		}
		return config.SourcedConfig[string]{
				Value:  "",
				Source: "empty",
				Kind:   config.EmptySourceKind,
			}, errors.Newf(
				"could not find environment %q from source %q. available environments: %s",
				provided.Value, provided.Source, strings.Join(available, ", "),
			)
	}
}

func NewManager(ctx context.Context, opts *Inputs) (*Manager, error) {
	httpClient := opts.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if opts.Config == nil {
		return nil, errors.New("missing config manager")
	}
	r := &Manager{
		config: opts.Config,
		authClient: serverv1connect.NewAuthServiceClient(
			httpClient,
			opts.Config.GetAPIServer().Value,
			connect.WithInterceptors(
				connect.UnaryInterceptorFunc(
					func(next connect.UnaryFunc) connect.UnaryFunc {
						return func(c context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
							if opts.Timeout != nil {
								if _, deadlineSet := c.Deadline(); !deadlineSet {
									timeoutCtx, cancel := context.WithTimeout(c, *opts.Timeout)
									c = timeoutCtx
									defer cancel()
								}
							}
							req.Header().Set("x-chalk-server", "go-api")
							req.Header().Set("User-Agent", internal.UserAgent())
							return next(c, req)
						}
					},
				),
			),
		),
		mu:    &sync.Mutex{},
		token: opts.Token,
	}

	var err error
	if r.token == nil {
		r.token, err = r.GetJWT(ctx, time.Now())
		if err != nil {
			return nil, errors.Wrap(err, "initializing token refresher")
		}
	}

	if opts.SkipEnvironmentNameMapping {
		// Use environment ID verbatim without validation or mapping
		if r.config.EnvironmentId.Value == "" {
			return nil, errors.New("environment ID is required when SkipEnvironmentNameMapping is enabled")
		}
	} else {
		r.config.EnvironmentId, err = cleanEnvironmentId(r.config.EnvironmentId, r.token)
		if err != nil {
			return nil, errors.Wrap(err, "initializing environment id")
		}
	}

	if !opts.SkipEngineMapping {
		envName := r.token.EnvironmentIdToName[r.config.EnvironmentId.Value]
		if e := r.token.Engines[r.config.EnvironmentId.Value]; r.config.GetJSONQueryServer().Kind == config.DefaultSourceKind && e != "" {
			r.config.SetJSONQueryServer(config.NewFromToken(e, fmt.Sprintf("token for environment %q", envName)))
		}

		if e := r.token.GrpcEngines[r.config.EnvironmentId.Value]; r.config.GetGRPCQueryServer().Kind == config.DefaultSourceKind && e != "" {
			r.config.SetGRPCQueryServer(config.NewFromToken(e, fmt.Sprintf("token for environment %q", envName)))
		}
	}

	return r, nil
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
		ClientId:     string(r.config.ClientId.Value),
		ClientSecret: string(r.config.ClientSecret.Value),
		GrantType:    "client_credentials",
		Scope:        nil,
	}
	if r.config.Scope.Value != "" {
		req.Scope = &r.config.Scope.Value
	}

	t, err := r.authClient.GetToken(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "refreshing token")
	}
	r.token = t.Msg
	return r.token, nil
}

func (r *Manager) GetConfig() *config.Manager {
	return r.config
}
