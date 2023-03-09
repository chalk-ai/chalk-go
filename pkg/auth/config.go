package auth

import (
	"github.com/sirupsen/logrus"
	"os"
)

func (cfg authConfig) GetProjectAuthConfigForWD() (*projectAuthConfig, string, error) {
	getwd, err := os.Getwd()

	path := ""
	if err != nil {
		return nil, path, err
	}

	var tok *projectAuthConfig = nil
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
