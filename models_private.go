package chalk

import (
	"github.com/apache/arrow/go/v16/arrow"
	"time"
)

type onlineQueryRequestSerialized struct {
	Inputs         map[string]any     `json:"inputs,string"`
	Outputs        []string           `json:"outputs"`
	Context        onlineQueryContext `json:"context"`
	Staleness      map[string]string  `json:"staleness"`
	IncludeMeta    bool               `json:"include_meta"`
	IncludeMetrics bool               `json:"include_metrics"`
	DeploymentId   *string            `json:"deployment_id"`
	QueryName      *string            `json:"query_name"`
	CorrelationId  *string            `json:"correlation_id"`
	Meta           map[string]string  `json:"meta"`
}

type onlineQueryResponseSerialized struct {
	Data   []featureResultSerialized `json:"data"`
	Errors []chalkErrorSerialized    `json:"errors"`
	Meta   *QueryMeta                `json:"meta"`
}

type onlineQueryResultFeather struct {
	HasData    bool
	ScalarData arrow.Table
	GroupsData map[Fqn]arrow.Table
	Errors     []ServerError
	Meta       *QueryMeta
}

type QueryName = string
type Fqn = string

type OnlineQueryBulkResponse struct {
	QueryResults map[QueryName]onlineQueryResultFeather
}

type onlineQueryContext struct {
	Environment *string  `json:"environment"`
	Tags        []string `json:"tags"`
}

type featureResultSerialized struct {
	Field     string                 `json:"field"`
	Value     any                    `json:"value"`
	Pkey      any                    `json:"pkey"`
	Timestamp string                 `json:"ts"`
	Meta      *FeatureResolutionMeta `json:"meta"`
	Error     *chalkErrorSerialized  `json:"error"`
}

type offlineQueryInputSerialized struct {
	Columns []string `json:"columns"`
	Values  [][]any  `json:"values"`
}

type offlineQueryRequestSerialized struct {
	Input                offlineQueryInputSerialized `json:"input"`
	Output               []string                    `json:"output"`
	RequiredOutput       []string                    `json:"required_output"`
	DatasetName          *string                     `json:"dataset_name"`
	Branch               *string                     `json:"branch"`
	MaxSamples           *int                        `json:"max_samples"`
	DestinationFormat    string                      `json:"destination_format"`
	JobId                *string                     `json:"job_id"`
	MaxCacheAge          *int                        `json:"max_cache_age_secs"`
	ObservedAtLowerBound *time.Time                  `json:"observed_at_lower_bound"`
	ObservedAtUpperBound *time.Time                  `json:"observed_at_upper_bound"`
}

type DatasetSampleFilter struct {
	LowerBound *time.Time
	UpperBound *time.Time
	MaxSamples *int
}

type DatasetFilter struct {
	SampleFilters DatasetSampleFilter `json:"sample_filters"`
	MaxCacheAge   *float64            `json:"max_cache_age_secs"`
}

type chalkHttpException struct {
	Detail *string `json:"detail"`
	Trace  *string `json:"trace"`
}

type sendRequestParams struct {
	Body   any
	Method string
	URL    string

	Response    any
	DontRefresh bool

	EnvironmentOverride string
	PreviewDeploymentId string
	Versioned           bool
	Branch              *string
	Tags                []string

	IsEngineRequest bool
}

type getTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type getTokenResponse struct {
	AccessToken        string            `json:"access_token"`
	TokenType          string            `json:"token_type"`
	ExpiresIn          int               `json:"expires_in"`
	ApiServer          string            `json:"api_server"`
	PrimaryEnvironment string            `json:"primary_environment"`
	Engines            map[string]string `json:"engines"`
}

type chalkErrorSerialized struct {
	Code      string             `json:"code"`
	Category  string             `json:"category"`
	Message   string             `json:"message"`
	Exception *ResolverException `json:"exception"`
	Feature   string             `json:"feature"`
	Resolver  string             `json:"resolver"`
}
