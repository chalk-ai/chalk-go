package client

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go/pkg/utils"
)

func (request *OnlineQueryParams) serialize() ([]byte, error) {
	httpRequestBody := onlineQueryHttpRequest{
		Inputs:         request.Inputs,
		Outputs:        request.Outputs,
		Context:        request.Context,
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

func (response *onlineQueryHttpResponse) deserialize() (OnlineQueryResult, error) {
	values := make(map[string]any)

	for _, result := range response.Data {
		values[result.Field] = result.Value
	}

	return OnlineQueryResult{
		Data:   response.Data,
		Meta:   response.Meta,
		values: values,
	}, nil
}
