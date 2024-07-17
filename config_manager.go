package chalk

import (
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"time"
)

// getTokenResult is agnostic to whether the token
// was obtained via gRPC or REST.
type getTokenResult struct {
	AccessToken        string
	PrimaryEnvironment string
	ValidUntil         time.Time
	Engines            map[string]string
}

type configManager struct {
	apiServer     auth2.SourcedConfig
	clientId      auth2.SourcedConfig
	clientSecret  auth2.SourcedConfig
	environmentId auth2.SourcedConfig

	jwt                *auth2.JWT
	initialEnvironment auth2.SourcedConfig
	engines            map[string]string
	getToken           func() (*getTokenResult, error)
}

func (r *configManager) refresh(force bool) error {
	if !force && r.jwt != nil && r.jwt.IsValid() {
		return nil
	}

	config, getTokenErr := r.getToken()
	if getTokenErr != nil {
		return getTokenErr
	}

	if r.initialEnvironment.Value == "" {
		r.environmentId = auth2.SourcedConfig{
			Value:  config.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		r.environmentId = r.initialEnvironment
	}

	r.jwt = &auth2.JWT{
		Token:      config.AccessToken,
		ValidUntil: config.ValidUntil,
	}

	r.engines = config.Engines

	return nil
}
