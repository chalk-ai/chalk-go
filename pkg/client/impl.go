package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) OnlineQuery(request OnlineQueryRequest) (OnlineQueryResponse, error) {
	if request.Context == nil {
		request.Context = &OnlineQueryContext{Environment: c.EnvironmentId}
	}

	stringifiedInputs := getStringifiedInputs(request.Inputs)
	httpRequestBody := OnlineQueryHttpRequest{
		Inputs:         stringifiedInputs,
		Outputs:        request.Outputs,
		Context:        request.Context,
		Staleness:      request.Staleness,
		IncludeMeta:    request.IncludeMeta,
		IncludeMetrics: request.IncludeMetrics,
		DeploymentId:   request.DeploymentId,
		QueryName:      request.QueryName,
		CorrelationId:  request.CorrelationId,
		Meta:           request.Meta,
	}

	var httpResponse OnlineQueryHttpResponse
	jsonRequestBody, err := json.Marshal(httpRequestBody)
	if err != nil {
		return OnlineQueryResponse{}, err
	}
	httpRequest, err := http.NewRequest("POST", "v1/query/online", bytes.NewBuffer(jsonRequestBody))

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
