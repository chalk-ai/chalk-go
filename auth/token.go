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
	// mu serializes token refreshes so concurrent requests do not each trigger
	// one. Reads go through token, which is atomic, so the common case of an
	// already-valid token takes no lock.
	mu     *sync.Mutex
	config *config.Manager
	// token is read on every request from any number of goroutines and written on
	// refresh, so it must be accessed atomically rather than directly.
	token         atomic.Pointer[serverv1.GetTokenResponse]
	authClient    serverv1connect.AuthServiceClient
	tokenProvider TokenProvider
}

// TokenProvider returns a currently-valid Chalk JWT. It is called whenever the
// cached token goes stale, in place of the client-credentials exchange.
//
// Callers holding a rotating credential -- one an external process replaces
// before it expires -- supply this so the token can be re-read at its source.
// Without it a pre-issued token is a one-shot snapshot: there are no credentials
// to exchange with, so the client is stranded the moment the token expires.
//
// Implementations must not call back into the client that owns this Manager;
// they are invoked while its lock is held.
type TokenProvider func(ctx context.Context) (*serverv1.GetTokenResponse, error)

type Inputs struct {
	// Token is a pre-issued JWT to authenticate with, instead of exchanging client
	// credentials. Supply it when the caller has already authenticated and has no
	// client secret -- and set TokenProvider too if the token rotates, since a
	// Token on its own cannot be renewed once it expires.
	Token *serverv1.GetTokenResponse

	// TokenProvider supplies a token whenever the cached one goes stale, instead
	// of running the client-credentials exchange. Sufficient on its own -- the
	// first token is obtained from it if Token is unset. See TokenProvider.
	TokenProvider TokenProvider

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
// all. Both a caller-supplied Token and a TokenProvider result go through here.
//
// Each check otherwise fails far from its cause: a nil token gets dereferenced
// inside an interceptor, and an empty access token produces a bare "Bearer "
// header that the server rejects without explaining why.
//
// Deliberately does not require ExpiresAt. A token without one reads as the Unix
// epoch and so is always stale, which is wasteful but survivable when client
// credentials are also configured -- a combination that predates this change and
// that callers, including this repo's own tests, do use. The provider path has no
// such fallback and does require it; see GetJWT.
func validatePreIssuedToken(token *serverv1.GetTokenResponse) error {
	if token == nil {
		return errors.New("token is nil")
	}
	if token.GetAccessToken() == "" {
		return errors.New("token has an empty access token")
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
		mu:            &sync.Mutex{},
		tokenProvider: opts.TokenProvider,
	}
	r.token.Store(opts.Token)

	// Reject an unusable pre-issued token here rather than at first request. The
	// common trap is omitting ExpiresAt, which reads as the Unix epoch and makes
	// the token permanently stale -- so every request would try to exchange
	// credentials the caller deliberately does not have.
	if token := r.token.Load(); token != nil {
		if err := validatePreIssuedToken(token); err != nil {
			return nil, errors.Wrap(err, "invalid pre-issued JWT")
		}
	}

	var err error
	if r.token.Load() == nil {
		// If none of a pre-issued JWT, a token provider, or a complete pair of
		// client credentials was given, short-circuit with a clearer error than
		// letting GetToken fail with "Client ID and secret are invalid" against
		// the api-server.
		if r.tokenProvider == nil && (r.config.ClientId.Value == "" || r.config.ClientSecret.Value == "") {
			credentialsErr := errors.New(
				"no JWT and no ClientId/ClientSecret provided; pass a pre-issued JWT (e.g. --access-token) or set client credentials",
			)
			if configErr := r.config.ProjectConfigErr(); configErr != nil {
				return nil, errors.Wrap(configErr, credentialsErr.Error())
			}
			return nil, credentialsErr
		}
		token, err := r.GetJWT(ctx, time.Now())
		if err == nil {
			r.token.Store(token)
		}
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
		r.config.EnvironmentId, err = cleanEnvironmentId(r.config.EnvironmentId, r.token.Load())
		if err != nil {
			return nil, errors.Wrap(err, "initializing environment id")
		}
	}

	if !opts.SkipEngineMapping {
		activeToken := r.token.Load()
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

func (r *Manager) GetJWT(
	ctx context.Context,
	newerThan time.Time,
) (*serverv1.GetTokenResponse, error) {
	if to := r.token.Load(); to != nil && to.GetExpiresAt().AsTime().After(newerThan) {
		return to, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if to := r.token.Load(); to != nil && to.GetExpiresAt().AsTime().After(newerThan) {
		return to, nil
	}

	// A rotating credential is re-read at its source rather than exchanged. There
	// are no client credentials to exchange with on that path, so without this
	// the request below would be sent with two empty strings and rejected --
	// stranding the client the moment its token expires, even though a fresh one
	// is already available.
	if r.tokenProvider != nil {
		token, err := r.tokenProvider(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "refreshing pre-issued token")
		}
		// Validate rather than trust: this is a caller-supplied callback, and a
		// nil or empty token here would otherwise be cached and only surface much
		// later as a nil dereference in an interceptor or a bare "Bearer " header.
		if err := validatePreIssuedToken(token); err != nil {
			return nil, errors.Wrap(err, "refreshing pre-issued token")
		}
		// The provider is expected to return something usable. If it hands back a
		// token that is already stale, say so rather than stamping it onto a request
		// and letting the server reject it as an opaque 401. Callers ask for a
		// minute of headroom, so this also reports a producer that has stopped
		// rotating instead of silently re-invoking the provider on every request.
		//
		// A missing ExpiresAt lands here too, since it reads as the Unix epoch.
		// Unlike the caller-supplied Token path there is no credential fallback
		// here, so an unbounded expiry is never usable.
		if !token.GetExpiresAt().AsTime().After(newerThan) {
			return nil, errors.Newf(
				"pre-issued token provider returned a token expiring at %s, which is not valid until %s; the token source may no longer be refreshed",
				token.GetExpiresAt().AsTime(),
				newerThan,
			)
		}
		r.token.Store(token)
		return token, nil
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
	r.token.Store(t.Msg)
	return t.Msg, nil
}

func (r *Manager) GetConfig() *config.Manager {
	return r.config
}
