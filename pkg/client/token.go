package client

import (
	"bytes"
	"encoding/json"
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
	"time"
)

func (c *Client) GetJWT() (*auth.JWT, error) {
	if c.JWT != nil && !time.Time.IsZero(c.JWT.ValidUntil) &&
		c.JWT.ValidUntil.After(time.Now().UTC().Add(-10*time.Second)) {
		return c.JWT, nil
	}

	jsonBody, err := json.Marshal(GetTokenRequest{
		ClientId:     c.ClientId,
		ClientSecret: c.ClientSecret,
		GrantType:    "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "v1/oauth/token", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	response := GetTokenResponse{}
	err = c.sendRequest(req, &response)

	expiry := time.Now().UTC().Add(time.Duration(response.ExpiresIn) * time.Second)
	jwt := &auth.JWT{
		Token:      response.AccessToken,
		ValidUntil: expiry,
	}
	return jwt, err
}
