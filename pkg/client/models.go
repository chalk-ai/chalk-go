package client

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

type Client struct {
	BaseUrl       string
	JWT           *auth.JWT
	HTTPClient    *http.Client
	EnvironmentId string
	ClientId      string
	ClientSecret  string
}

type ChalkHTTPException struct {
	Detail *string `json:"detail"`
	Trace  *string `json:"trace"`
}

type GetTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ApiServer   string `json:"api_server"`
}

type ClientError struct {
	Path          string
	Message       string
	StatusCode    int
	ContentLength int64
	Trace         *string
}

func (e *ClientError) Error() string {
	if e.Trace != nil {
		return fmt.Sprintf("HTTPClient Error: path=%q, message=%q, status=%d, content-length=%d, trace=%q",
			e.Path, e.Message, e.StatusCode, e.ContentLength, *e.Trace)
	}
	return fmt.Sprintf("HTTPClient Error: path=%q, message=%q, status=%d, content-length=%d",
		e.Path, e.Message, e.StatusCode, e.ContentLength)
}
