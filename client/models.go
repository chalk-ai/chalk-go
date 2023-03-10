package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"github.com/chalk-ai/chalk-go/pkg/enum"
	"net/http"
	"time"
)

type ChalkClientImpl struct {
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
	Meta *QueryMeta

	features map[string]FeatureResult
}

type FeatureResult struct {
	Field     string
	Value     any
	Timestamp time.Time
	Meta      *FeatureResolutionMeta
	Error     *ChalkServerError
}

type FeatureResolutionMeta struct {
	ChosenResolverFqn string `json:"chosen_resolver_fqn"`
	CacheHit          bool   `json:"cache_hit"`
	PrimitiveType     string `json:"primitive_type"`
	Version           int    `json:"version"`
}

type QueryMeta struct {
	ExecutionDurationS float64 `json:"execution_duration_s"`
	DeploymentId       string  `json:"deployment_id"`
	QueryId            string  `json:"query_id"`
}

type ChalkException struct {
	Kind       string `json:"kind"`
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
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
	Exception *ChalkException
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
