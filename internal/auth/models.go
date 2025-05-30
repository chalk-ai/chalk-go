package auth

import "time"

var validityBuffer = 5 * time.Minute

type JWT struct {
	Token      string    `yaml:"value"`
	ValidUntil time.Time `yaml:"validUntil"`
}

type ProjectToken struct {
	Name              string `yaml:"name"`
	ClientId          string `yaml:"clientId"`
	ClientSecret      string `yaml:"clientSecret"`
	ValidUntil        string `yaml:"validUntil"`
	ApiServer         string `yaml:"apiServer"`
	ActiveEnvironment string `yaml:"activeEnvironment"`
	JWT               JWT    `yaml:"jwt,omitempty"`
}

func (t *JWT) IsValid() bool {
	expiry := t.ValidUntil.Add(-validityBuffer)
	return time.Now().UTC().Before(expiry)
}

type ProjectTokens struct {
	Tokens map[string]*ProjectToken `yaml:"tokens,omitempty"`
}
