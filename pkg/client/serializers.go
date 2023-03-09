package client

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go/pkg/client/clientenums"
	"github.com/chalk-ai/chalk-go/pkg/utils"
)

func (request *OnlineQueryParams) serialize() ([]byte, error) {
	context := OnlineQueryContext{
		Environment: utils.StrPtrOrNil(request.EnvironmentId),
		Tags:        request.Tags,
	}

	httpRequestBody := onlineQueryHttpRequest{
		Inputs:         request.Inputs,
		Outputs:        request.Outputs,
		Context:        context,
		Staleness:      request.Staleness,
		IncludeMeta:    request.IncludeMeta,
		IncludeMetrics: request.IncludeMetrics,
		DeploymentId:   utils.StrPtrOrNil(request.DeploymentId),
		QueryName:      utils.StrPtrOrNil(request.QueryName),
		CorrelationId:  utils.StrPtrOrNil(request.CorrelationId),
		Meta:           request.Meta,
	}
	jsonRequestBody, err := json.Marshal(httpRequestBody)
	if err != nil {
		return []byte{}, err
	}

	return jsonRequestBody, nil
}

func (response *onlineQueryHttpResponse) deserialize() OnlineQueryResult {
	values := make(map[string]any)

	for _, result := range response.Data {
		values[result.Field] = result.Value
	}

	return OnlineQueryResult{
		Data:   response.Data,
		Meta:   response.Meta,
		values: values,
	}
}

func (e *chalkErrorSerialized) deserialize() (ChalkServerError, error) {
	errorCode, getErrorCodeErr := clientenums.GetErrorCode(e.Code)
	if getErrorCodeErr != nil {
		return ChalkServerError{}, nil
	}
	return ChalkServerError{
		Code:      *errorCode,
		Category:  e.Category,
		Message:   e.Message,
		Exception: e.Exception,
		Feature:   e.Feature,
		Resolver:  e.Resolver,
	}, nil
}

func deserializeChalkErrors(errors []chalkErrorSerialized) ([]ChalkServerError, *ChalkClientError) {
	deserializedErrors := make([]ChalkServerError, 0)
	for _, serializedErr := range errors {
		deserializedError, deserializationFailure := serializedErr.deserialize()
		if deserializationFailure != nil {
			return []ChalkServerError{}, &ChalkClientError{
				Message: deserializationFailure.Error(),
			}
		}
		deserializedErrors = append(deserializedErrors, deserializedError)

	}
	return deserializedErrors, nil
}
