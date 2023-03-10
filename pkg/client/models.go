package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"github.com/chalk-ai/chalk-go/pkg/enum"
	"net/http"
	"time"
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
	Inputs         map[string]any
	Outputs        []string
	Staleness      map[string]time.Duration
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

type FeatureResult struct {
	Field     string
	Value     any
	Timestamp time.Time
	Meta      map[string]any
	Error     *ChalkServerError
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
