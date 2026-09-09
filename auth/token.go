package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
)

type Manager struct {
	mu           *sync.Mutex
	config       *config.Manager
	auth         atomic.Pointer[AuthSnapshot]
	authClient   serverv1connect.AuthServiceClient
	authProvider AuthProvider
}

// AuthSnapshot is the bearer token and effective environment that must be sent
// together on an authenticated request.
type AuthSnapshot struct {
	Token         *serverv1.GetTokenResponse
	EnvironmentID string
}

// AuthProvider returns a currently-valid authentication snapshot. It is called
// whenever the cached token goes stale, in place of the client-credentials
// exchange. Token and EnvironmentID must be resolved atomically from the same
// credential state.
// Implementations must not call back into the client that owns this Manager;
// they are invoked while its lock is held.
type AuthProvider func(ctx context.Context) (*AuthSnapshot, error)

type Inputs struct {
	// Token is a pre-issued JWT to authenticate with, instead of exchanging client
	// credentials
	Token *serverv1.GetTokenResponse

	// AuthProvider supplies a token and effective environment whenever the cached
	// token goes stale, instead of running the client-credentials exchange.
	// Sufficient on its own -- the first snapshot is obtained from it if Token is
	// unset. See AuthProvider.
	AuthProvider AuthProvider

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

// validatePreIssuedToken rejects a token that cannot authenticate a request at
// all. Both a caller-supplied Token and an AuthProvider result go through here.
func validatePreIssuedToken(token *serverv1.GetTokenResponse) error {
	if token == nil {
		return errors.New("token is nil")
	}
	if token.GetAccessToken() == "" {
		return errors.New("token has an empty access token")
	}
	return nil
}

func validateAuthSnapshot(snapshot *AuthSnapshot) error {
	if snapshot == nil {
		return errors.New("auth snapshot is nil")
	}
	if err := validatePreIssuedToken(snapshot.Token); err != nil {
		return err
	}
	if snapshot.EnvironmentID == "" {
		return errors.New("auth snapshot has an empty environment ID")
	}
	return nil
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
		mu:           &sync.Mutex{},
		authProvider: opts.AuthProvider,
	}
	// Validate before storing so nothing invalid is ever observable in r.auth,
	// matching how GetAuth handles an AuthProvider result.
	if opts.Token != nil {
		if err := validatePreIssuedToken(opts.Token); err != nil {
			return nil, errors.Wrap(err, "invalid pre-issued JWT")
		}
		r.auth.Store(&AuthSnapshot{
			Token:         opts.Token,
			EnvironmentID: r.config.EnvironmentId.Value,
		})
	} else {
		// If none of a pre-issued JWT, an auth provider, or a complete pair of
		// client credentials was given, short-circuit with a clearer error than
		// letting GetToken fail with "Client ID and secret are invalid" against
		// the api-server.
		if r.authProvider == nil && (r.config.ClientId.Value == "" || r.config.ClientSecret.Value == "") {
			credentialsErr := errors.New(
				"no JWT and no ClientId/ClientSecret provided; pass a pre-issued JWT (e.g. --access-token) or set client credentials",
			)
			if configErr := r.config.ProjectConfigErr(); configErr != nil {
				return nil, errors.Wrap(configErr, credentialsErr.Error())
			}
			return nil, credentialsErr
		}
		if _, err := r.GetAuth(ctx, time.Now()); err != nil {
			return nil, errors.Wrap(err, "initializing token refresher")
		}
	}

	activeAuth := r.auth.Load()
	if opts.SkipEnvironmentNameMapping {
		// Use environment ID verbatim without validation or mapping
		if activeAuth.EnvironmentID == "" {
			return nil, errors.New("environment ID is required when SkipEnvironmentNameMapping is enabled")
		}
		if r.config.EnvironmentId.Value == "" {
			r.config.EnvironmentId = config.NewFromArg(activeAuth.EnvironmentID)
		}
	} else {
		cleaned, err := cleanEnvironmentId(r.config.EnvironmentId, activeAuth.Token)
		if err != nil {
			return nil, errors.Wrap(err, "initializing environment id")
		}
		r.config.EnvironmentId = cleaned
		activeAuth = &AuthSnapshot{
			Token:         activeAuth.Token,
			EnvironmentID: r.config.EnvironmentId.Value,
		}
		r.auth.Store(activeAuth)
	}

	if !opts.SkipEngineMapping {
		activeToken := activeAuth.Token
		envName := activeToken.EnvironmentIdToName[r.config.EnvironmentId.Value]
		if e := activeToken.Engines[r.config.EnvironmentId.Value]; r.config.GetJSONQueryServer().Kind == config.DefaultSourceKind && e != "" {
			r.config.SetJSONQueryServer(config.NewFromToken(e, fmt.Sprintf("token for environment %q", envName)))
		}

		if e := activeToken.GrpcEngines[r.config.EnvironmentId.Value]; r.config.GetGRPCQueryServer().Kind == config.DefaultSourceKind && e != "" {
			r.config.SetGRPCQueryServer(config.NewFromToken(e, fmt.Sprintf("token for environment %q", envName)))
		}
	}

	return r, nil
}

// GetAuth returns a token and effective environment from the same credential
// snapshot, refreshing both atomically when the token goes stale.
func (r *Manager) GetAuth(
	ctx context.Context,
	newerThan time.Time,
) (*AuthSnapshot, error) {
	if snapshot := r.auth.Load(); snapshot != nil && snapshot.Token.GetExpiresAt().AsTime().After(newerThan) {
		return snapshot, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if snapshot := r.auth.Load(); snapshot != nil && snapshot.Token.GetExpiresAt().AsTime().After(newerThan) {
		return snapshot, nil
	}

	// A rotating credential is re-read at its source rather than exchanged. There
	// are no client credentials to exchange with on that path, so without this
	// the request below would be sent with two empty strings and rejected --
	// stranding the client the moment its token expires, even though a fresh one
	// is already available.
	if r.authProvider != nil {
		snapshot, err := r.authProvider(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "refreshing pre-issued token")
		}
		if err := validateAuthSnapshot(snapshot); err != nil {
			return nil, errors.Wrap(err, "refreshing pre-issued token")
		}

		if !snapshot.Token.GetExpiresAt().AsTime().After(newerThan) {
			return nil, errors.Newf(
				"pre-issued token provider returned a token expiring at %s, which is not valid until %s; the token source may no longer be refreshed",
				snapshot.Token.GetExpiresAt().AsTime(),
				newerThan,
			)
		}
		r.auth.Store(snapshot)
		return snapshot, nil
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
	snapshot := &AuthSnapshot{
		Token:         t.Msg,
		EnvironmentID: r.config.EnvironmentId.Value,
	}
	r.auth.Store(snapshot)
	return snapshot, nil
}

// GetJWT is a token-only convenience wrapper. Authenticated request paths
// should use GetAuth so the token cannot be paired with a stale environment.
func (r *Manager) GetJWT(
	ctx context.Context,
	newerThan time.Time,
) (*serverv1.GetTokenResponse, error) {
	snapshot, err := r.GetAuth(ctx, newerThan)
	if err != nil {
		return nil, err
	}
	return snapshot.Token, nil
}

func (r *Manager) GetConfig() *config.Manager {
	return r.config
}
