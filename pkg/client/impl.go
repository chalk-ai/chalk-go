package client

import (
	"bytes"
	"net/http"
)

func (c *Client) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, error) {
	emptyResult := OnlineQueryResult{}

	if request.Context == nil {
		request.Context = &OnlineQueryContext{Environment: c.EnvironmentId}
	}

	jsonRequestBody, err := request.serialize()
	if err != nil {
		return emptyResult, err
	}

	httpRequest, err := http.NewRequest("POST", "v1/query/online", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return emptyResult, err
	}

	var httpResponse onlineQueryHttpResponse

	err = c.sendRequest(httpRequest, &httpResponse)
	if err != nil {
		return emptyResult, err
	}
	if len(httpResponse.Errors) > 0 {
		return emptyResult, &ChalkErrorResponse{ChalkErrors: httpResponse.Errors}
	}

	return httpResponse.deserialize(), nil
}
