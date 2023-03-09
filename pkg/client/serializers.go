package client

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go/pkg/client/clientenums"
)

func (request *OnlineQueryParams) serialize() ([]byte, error) {
	context := OnlineQueryContext{
		Environment: stringPointerOrNil(request.EnvironmentId),
		Tags:        request.Tags,
	}

	httpRequestBody := onlineQueryHttpRequest{
		Inputs:         request.Inputs,
		Outputs:        request.Outputs,
		Context:        context,
		Staleness:      request.Staleness,
		IncludeMeta:    request.IncludeMeta,
		IncludeMetrics: request.IncludeMetrics,
		DeploymentId:   stringPointerOrNil(request.DeploymentId),
		QueryName:      stringPointerOrNil(request.QueryName),
		CorrelationId:  stringPointerOrNil(request.CorrelationId),
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
		return ChalkServerError{}, getErrorCodeErr
	}

	errorCodeCategory, getCategoryErr := clientenums.GetErrorCodeCategory(e.Category)
	if getCategoryErr != nil {
		return ChalkServerError{}, getCategoryErr
	}

	return ChalkServerError{
		Code:      *errorCode,
		Category:  *errorCodeCategory,
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
