package client

import (
	"bytes"
	"net/http"
)

func (c *Client) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, error) {
	emptyResult := OnlineQueryResult{}

	if request.EnvironmentId == "" {
		request.EnvironmentId = c.EnvironmentId.Value
	}

	jsonRequestBody, err := request.serialize()
	if err != nil {
		return emptyResult, err
	}

	httpRequest, err := http.NewRequest("POST", "v1/query/online", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return emptyResult, err
	}

	var httpResponse onlineQueryResponseSerialized

	err = c.sendRequest(httpRequest, &httpResponse)
	if err != nil {
		return emptyResult, err
	}
	if len(httpResponse.Errors) > 0 {
		serverErrors, deserializationErr := deserializeChalkErrors(httpResponse.Errors)
		if deserializationErr != nil {
			return OnlineQueryResult{}, &ChalkErrorResponse{
				ClientError: deserializationErr,
			}
		}

		return emptyResult, &ChalkErrorResponse{ServerErrors: serverErrors}
	}

	return httpResponse.deserialize(), nil
}
