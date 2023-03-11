package chalk

import (
	"net/http"
)

func (c *ChalkClientImpl) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse) {
	emptyResult := OnlineQueryResult{}

	if request.EnvironmentId == "" {
		request.EnvironmentId = c.EnvironmentId.Value
	}

	var serializedResponse onlineQueryResponseSerialized

	err := c.sendRequest(sendRequestParams{Method: "POST", URL: "v1/query/online", Body: request.serialize(), Response: &serializedResponse})
	if err != nil {
		httpError, ok := err.(*ChalkHttpError)
		if ok {
			return OnlineQueryResult{}, &ChalkErrorResponse{HttpError: httpError}
		}
		return OnlineQueryResult{}, &ChalkErrorResponse{ClientError: &ChalkClientError{Message: err.Error()}}
	}
	if len(serializedResponse.Errors) > 0 {
		serverErrors, deserializationErr := deserializeChalkErrors(serializedResponse.Errors)
		if deserializationErr != nil {
			return OnlineQueryResult{}, &ChalkErrorResponse{
				ClientError: &ChalkClientError{deserializationErr.Error()},
			}
		}

		return emptyResult, &ChalkErrorResponse{ServerErrors: serverErrors}
	}

	response, err := serializedResponse.deserialize()
	if err != nil {
		return OnlineQueryResult{}, &ChalkErrorResponse{
			ClientError: &ChalkClientError{err.Error()},
		}
	}

	return response, nil
}

func (c *ChalkClientImpl) SetLogger(logger *LeveledLogger) {
	c.logger = logger
}

func (c *ChalkClientImpl) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}
