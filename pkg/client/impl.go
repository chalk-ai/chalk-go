package client

import (
	"bytes"
	"net/http"
)

func (c *Client) OnlineQuery(request OnlineQueryRequest) (OnlineQueryResponse, error) {
	if request.Context == nil {
		request.Context = &OnlineQueryContext{Environment: c.EnvironmentId}
	}

	jsonRequestBody, err := request.serialize()
	if err != nil {
		return OnlineQueryResponse{}, err
	}

	httpRequest, err := http.NewRequest("POST", "v1/query/online", bytes.NewBuffer(jsonRequestBody))
	var httpResponse OnlineQueryHttpResponse

	err = c.sendRequest(httpRequest, &httpResponse)
	if err != nil {
		return OnlineQueryResponse{}, err
	}
	if len(httpResponse.Errors) > 0 {
		return OnlineQueryResponse{}, &ChalkErrorResponse{ChalkErrors: httpResponse.Errors}
	}

	return OnlineQueryResponse{
		Data: httpResponse.Data,
		Meta: httpResponse.Meta,
	}, err
}
