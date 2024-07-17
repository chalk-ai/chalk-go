package chalk

import (
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"time"
)

type configManager struct {
	apiServer     auth2.SourcedConfig
	clientId      auth2.SourcedConfig
	clientSecret  auth2.SourcedConfig
	environmentId auth2.SourcedConfig

	jwt                *auth2.JWT
	initialEnvironment auth2.SourcedConfig
	engines            map[string]string
	getToken           func() (*getTokenResponse, *ClientError)
}

func (r *configManager) refreshConfig(forceRefresh bool) *ClientError {
	if !forceRefresh && r.jwt != nil && r.jwt.IsValid() {
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

	expiry := time.Now().UTC().Add(time.Duration(config.ExpiresIn) * time.Second)
	r.jwt = &auth2.JWT{
		Token:      config.AccessToken,
		ValidUntil: expiry,
	}

	r.engines = config.Engines

	return nil
}
