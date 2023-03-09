package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"github.com/chalk-ai/chalk-go/pkg/enum"
	"net/http"
)

type Client struct {
	ApiServer     config
	ClientId      config
	EnvironmentId config

	clientSecret config
	jwt          *auth.JWT
	httpClient   *http.Client
}

type OnlineQueryParams struct {
	Inputs  map[string]any
	Outputs []string
	// TODO: Use Duration. Drop JSON where
	Staleness      map[string]string
	IncludeMeta    bool
	IncludeMetrics bool
	DeploymentId   string
	EnvironmentId  string
	QueryName      string
	CorrelationId  string
	Meta           map[string]string
	Tags           []string
}

type OnlineQueryResult struct {
	Data []FeatureResult
	Meta *queryMeta

	features map[string]FeatureResult
}

// TODO create FeatureResultSerialized
type FeatureResult struct {
	Field     string                `json:"field"`
	Value     any                   `json:"Value"`
	Timestamp string                `json:"ts"`
	Meta      map[string]any        `json:"meta"`
	Error     *chalkErrorSerialized `json:"error"`
}

type ChalkErrorResponse struct {
	ServerErrors []ChalkServerError
	ClientError  *ChalkClientError
	HttpError    *ChalkHttpError
}

type ChalkServerError struct {
	Code      enum.ErrorCode
	Category  enum.ErrorCodeCategory
	Message   string
	Exception *chalkException
	Feature   string
	Resolver  string
}

type ChalkHttpError struct {
	Path          string
	Message       string
	StatusCode    int
	ContentLength int64
	Trace         *string
}

type ChalkClientError struct {
	Message string
}
