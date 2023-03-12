package auth

import "time"

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

type ProjectTokens struct {
	Tokens *map[string]*ProjectToken `yaml:"tokens,omitempty"`
}
