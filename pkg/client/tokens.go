package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
	"time"
)

func (c *Client) getJwt() (*auth.JWT, *ChalkClientError) {
	if c.jwt != nil && !time.Time.IsZero(c.jwt.ValidUntil) &&
		c.jwt.ValidUntil.After(time.Now().UTC().Add(-10*time.Second)) {
		return c.jwt, nil
	}

	jsonBody, err := json.Marshal(getTokenRequest{
		ClientId:     c.ClientId.Value,
		ClientSecret: c.ClientSecret.Value,
		GrantType:    "client_credentials",
	})
	if err != nil {
		return nil, &ChalkClientError{Message: err.Error()}
	}

	req, err := http.NewRequest("POST", "v1/oauth/token", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, &ChalkClientError{Message: err.Error()}
	}

	response := getTokenResponse{}
	err = c.sendRequest(req, &response)
	if err != nil {
		return nil, &ChalkClientError{Message: fmt.Sprintf(
			"Error obtaining access token with these credentials: "+
				"api_server=%q (%q), "+
				"client_id=%q (%q), "+
				"client_secret=%q (%q), "+
				"environment_id=%q (%q)",
			c.ApiServer.Value,
			c.ApiServer.source,
			c.ClientId.Value,
			c.ClientId.source,
			c.ClientSecret.Value,
			c.ClientSecret.source,
			c.EnvironmentId.Value,
			c.EnvironmentId.source,
		)}
	}

	expiry := time.Now().UTC().Add(time.Duration(response.ExpiresIn) * time.Second)
	jwt := &auth.JWT{
		Token:      response.AccessToken,
		ValidUntil: expiry,
	}
	return jwt, nil
}
