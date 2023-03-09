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
		ClientSecret: c.clientSecret.Value,
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
	err = c.sendRequest(requestParams{Request: req, Response: &response, DontRefresh: true})
	if err != nil {
		return nil, &ChalkClientError{Message: fmt.Sprintf(
			"Error obtaining access token: %s.\n"+
				"  Auth config:\n"+
				"    api_server=%q (source: %s),\n"+
				"    client_id=%q (source: %s),\n"+
				"    client_secret=*** (source: %s),\n"+
				"    environment_id=%q (source: %s)\n",
			err.Error(),
			c.ApiServer.Value,
			c.ApiServer.Source,
			c.ClientId.Value,
			c.ClientId.Source,
			c.clientSecret.Source,
			c.EnvironmentId.Value,
			c.EnvironmentId.Source,
		)}
	}

	expiry := time.Now().UTC().Add(time.Duration(response.ExpiresIn) * time.Second)
	jwt := &auth.JWT{
		Token:      response.AccessToken,
		ValidUntil: expiry,
	}
	return jwt, nil
}

func (c *Client) upsertJwt() *ChalkClientError {
	jwt, getJwtErr := c.getJwt()
	if getJwtErr != nil {
		return getJwtErr
	}
	c.jwt = jwt
	return nil
}
