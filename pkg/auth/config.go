package auth

import (
	"os"
)

func (cfg AuthConfig) GetProjectAuthConfigForWD() (*ProjectAuthConfig, string, error) {
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
		path = getwd
		tok = token
	}

	if token, ok := tokens["default"]; ok && tok == nil {
		path = "default"
		tok = token
	}

	return tok, path, nil
}
