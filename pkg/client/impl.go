package client

import (
	"bytes"
	"net/http"
)

func (c *Client) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, error) {
	if request.Context == nil {
		request.Context = &OnlineQueryContext{Environment: c.EnvironmentId}
	}

	jsonRequestBody, err := request.serialize()
	if err != nil {
		return OnlineQueryResult{}, err
	}

	httpRequest, err := http.NewRequest("POST", "v1/query/online", bytes.NewBuffer(jsonRequestBody))
	// TODO: Handle err
	var httpResponse onlineQueryHttpResponse

	err = c.sendRequest(httpRequest, &httpResponse)
	if err != nil {
		return OnlineQueryResult{}, err
	}
	if len(httpResponse.Errors) > 0 {
		return OnlineQueryResult{}, &ChalkErrorResponse{ChalkErrors: httpResponse.Errors}
	}

	return httpResponse.deserialize()
}
