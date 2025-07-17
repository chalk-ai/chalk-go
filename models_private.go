package chalk

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"time"
)

type onlineQueryResponseSerialized struct {
	Data   []featureResultSerialized `json:"data"`
	Errors []chalkErrorSerialized    `json:"errors"`
	Meta   *QueryMeta                `json:"meta"`
}

type onlineQueryResultFeather struct {
	HasData    bool
	ScalarData arrow.Table
	GroupsData map[Fqn]arrow.Table
	Errors     serverErrorsT
	Meta       *QueryMeta
}

type QueryName = string
type Fqn = string

type OnlineQueryBulkResponse struct {
	QueryResults map[QueryName]onlineQueryResultFeather

	allocator memory.Allocator
}

type featureResultSerialized struct {
	Field     string                 `json:"field"`
	Value     any                    `json:"value"`
	Pkey      any                    `json:"pkey"`
	Timestamp string                 `json:"ts"`
	Meta      *FeatureResolutionMeta `json:"meta"`
	Error     *chalkErrorSerialized  `json:"error"`
}

type DatasetSampleFilter struct {
	LowerBound *time.Time `json:"lower_bound"`
	UpperBound *time.Time `json:"upper_bound"`
	MaxSamples *int       `json:"max_samples"`
}

type DatasetFilter struct {
	SampleFilters DatasetSampleFilter `json:"sample_filters"`
	MaxCacheAge   *float64            `json:"max_cache_age_secs"`
}

type chalkHttpException struct {
	Detail any     `json:"detail"`
	Trace  *string `json:"trace"`
}

type sendRequestParams struct {
	Body   any
	Method string
	URL    string

	Response    any
	DontRefresh bool

	EnvironmentOverride   string
	PreviewDeploymentId   string
	Versioned             bool
	Branch                *string
	ResourceGroupOverride *string

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
