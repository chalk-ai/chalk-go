package chalk

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

type clientImpl struct {
	ApiServer     config
	ClientId      config
	EnvironmentId config

	clientSecret config
	jwt          *auth.JWT
	httpClient   *http.Client
	logger       *LeveledLogger
}

// OnlineQuery computes features values using online resolvers.
// See https://docs.chalk.ai/docs/query-basics for more information.
func (c *clientImpl) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse) {
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
