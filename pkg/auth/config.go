package auth

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func (cfg AuthConfig) GetProjectAuthConfigByDirectory() (*ProjectAuthConfig, string, error) {
	getwd, err := os.Getwd()

	path := ""
	if err != nil {
		return nil, path, err
	}

	var tok *ProjectAuthConfig = nil
	if cfg.Tokens == nil {
		return tok, path, nil
	}
	tokens := *cfg.Tokens
	if token, ok := tokens[getwd]; ok {
		logrus.Debug("Found directory-scoped config")
		path = getwd
		tok = token
	}

	if token, ok := tokens["default"]; ok && tok == nil {
		logrus.Debug("Found default config")
		path = "default"
		tok = token
	}

	return tok, path, nil
}

func (cfg AuthConfig) UpsertJWT(token string, expiresIn int) error {
	tokenConfig, _, err := cfg.GetTokenForWD()
	if err != nil || tokenConfig == nil {
		return err
	}

	expiry := time.Now().UTC().Add(time.Duration(expiresIn) * time.Second)
	tokenConfig.JWT = &JWT{
		Token:      &token,
		ValidUntil: &expiry,
	}
	return nil
}

func (cfg AuthConfig) ResetJWTs() {
	tokens := *cfg.Tokens
	for _, token := range tokens {
		token.JWT = nil
	}
}

func (cfg AuthConfig) Save() error {
	path, err := getConfigPath()
	if err != nil || path == nil {
		return err
	}

	yamlContents, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = utils.AtomicWriteFile(*path, yamlContents, 0666)
	return err
}
