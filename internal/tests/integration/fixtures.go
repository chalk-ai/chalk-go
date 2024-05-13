package integration

import (
	"bytes"
	"io"
	"net/http"
)

type user struct {
	Id               *int64
	FullNameOptional *string
	SocureScore      *float64
}

var testFeatures struct {
	User *user
}

type Intercepted struct {
	Header http.Header
	Body   []byte
}

type InterceptorHTTPClient struct {
	Intercepted Intercepted
}

func NewInterceptorHTTPClient() *InterceptorHTTPClient {
	return &InterceptorHTTPClient{}
}

func (c *InterceptorHTTPClient) Do(req *http.Request) (*http.Response, error) {
	bodyBytes, bodyBytesErr := io.ReadAll(req.Body)
	if bodyBytesErr != nil {
		return nil, bodyBytesErr
	}
	req.Body.Close()
	c.Intercepted = Intercepted{
		Header: req.Header,
		Body:   bodyBytes,
	}
	body := io.NopCloser(bytes.NewBufferString(`{"data": {"something": "exciting"}}`))
	return &http.Response{StatusCode: 200, Body: body}, nil
}

func (c *InterceptorHTTPClient) Get(url string) (*http.Response, error) {
	actualClient := &http.Client{}
	return actualClient.Get(url)
}
