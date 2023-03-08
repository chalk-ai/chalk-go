package client

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go/pkg/utils"
)

func (request *OnlineQueryRequest) serialize() ([]byte, error) {
	httpRequestBody := OnlineQueryHttpRequest{
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
