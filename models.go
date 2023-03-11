package chalk

import (
	"github.com/chalk-ai/chalk-go/pkg/enum"
	"time"
)

// OnlineQueryParams defines the parameters
// that help you execute an online query.
type OnlineQueryParams struct {
	// The features for which there are known values, mapped to those values.
	Inputs map[string]any

	// The features that you'd like to compute from the inputs.
	Outputs []string

	// Maximum staleness overrides for any output features or intermediate features.
	// See https://docs.chalk.ai/docs/query-caching for more information.
	Staleness map[string]time.Duration

	// If true, returns metadata about the query execution in the response.
	IncludeMeta bool

	// If true, returns performance metrics about the query execution in the response.
	IncludeMetrics bool

	// The environment under which to run the resolvers. API tokens can be scoped to an
	// environment. If no environment is specified in the query, but the token supports
	// only a single environment, then that environment will be taken as the scope for
	// executing the request.
	EnvironmentId string

	// The tags used to scope the resolvers.
	Tags []string

	// If specified, Chalk will route your request to the relevant preview deployment.
	PreviewDeploymentId string

	// The name for class of query you're making, for example, "loan_application_model".
	QueryName string

	// A globally unique ID for the query, used alongside logs and available in web
	// interfaces. If None, a correlation ID will be generated for you and returned on
	// the response.
	CorrelationId string

	// Arbitrary key:value pairs to associate with a query.
	Meta map[string]string
}

// OnlineQueryResult holds the result of an online query.
type OnlineQueryResult struct {
	// The output features and any query metadata.
	Data []FeatureResult

	// Execution metadata for the query. See QueryMeta for details.
	Meta *QueryMeta

	features map[string]FeatureResult
}

func (result *OnlineQueryResult) GetFeature(feature string) *FeatureResult {
	featureResult, found := result.features[feature]
	if !found {
		return nil
	}
	return &featureResult
}

func (result *OnlineQueryResult) GetFeatureValue(feature string) any {
	featureResult := result.GetFeature(feature)
	if featureResult == nil {
		return nil
	}
	return featureResult.Value
}

type FeatureResult struct {
	// The name of the feature requested, e.g. 'user.identity.has_voip_phone'.
	Field string

	// The value of the requested feature.
	// If an error was encountered in resolving this feature,
	// this field will be empty.
	Value any

	// The primary key of the resolved feature.
	Pkey any

	// The time at which this feature was computed.
	// This value could be significantly in the past if you're using caching.
	Timestamp time.Time

	// Detailed information about how this feature was computed.
	Meta *FeatureResolutionMeta

	// The error encountered in resolving this feature.
	// If no error occurred, this field is empty.
	Error *ChalkServerError
}

type FeatureResolutionMeta struct {
	// The name of the resolver that computed the feature value.
	ChosenResolverFqn string `json:"chosen_resolver_fqn"`

	// Whether the feature request was satisfied by a cached value.
	CacheHit bool `json:"cache_hit"`

	// Primitive type name for the feature, e.g. `str` for `some_feature: str`.
	// Returned only if query-level 'include_meta' is True.
	PrimitiveType string `json:"primitive_type"`

	// The version that was selected for this feature. Defaults to `default_version`, if query
	// does not specify a constraint. If no versioning information is provided on the feature definition,
	// the default version is `1`.
	Version int `json:"version"`
}

// QueryMeta represents metadata about a Chalk query.
type QueryMeta struct {
	// The id of the deployment that served this query.
	DeploymentId string `json:"deployment_id"`

	// The id of the environment that served this query. Not intended to be human readable,
	// but helpful for support.
	EnvironmentId string `json:"environment_id"`

	// The short name of the environment that served this query. For example: "dev" or "prod".
	EnvironmentName string `json:"environment_name"`

	// A unique ID generated and persisted by Chalk for this query. All computed features,
	// metrics, and logs are associated with this ID. Your system can store this ID for
	// audit and debugging workflows.
	QueryId string `json:"query_id"`

	// At the start of query execution, Chalk computes 'datetime.now()'. This value is used
	// to timestamp computed features.
	QueryTimestamp *time.Time `json:"query_timestamp"`

	// Deterministic hash of the 'structure' of the query. Queries that have the same
	// input/output features will typically have the same hash; changes may be observed
	// over time as we adjust implementation details.
	QueryHash string `json:"query_hash"`
}

type ChalkException struct {
	// The name of the class of the exception.
	Kind string `json:"kind"`

	// The message taken from the exception.
	Message string `json:"message"`

	// The stacktrace produced by the code.
	Stacktrace string `json:"stacktrace"`
}

type ChalkErrorResponse struct {
	// Errors that occurred in Chalk's server.
	ServerErrors []ChalkServerError

	// Errors that occurred in Client or its dependencies.
	ClientError *ChalkClientError

	// Errors that are standard HTTP errors such as missing authorization.
	HttpError *ChalkHttpError
}

// ChalkServerError is an error that occurred in Chalk's server,
// for example, when a resolver unexpectedly fails to run.
type ChalkServerError struct {
	// The type of the error.
	Code enum.ErrorCode

	// The category of the error, given in the type field for the error codes.
	// This will be one of "REQUEST", "NETWORK", and "FIELD".
	Category enum.ErrorCodeCategory

	// A readable description of the error message.
	Message string

	// The exception that caused the failure, if applicable.
	Exception *ChalkException

	// The fully qualified name of the failing feature, e.g. `user.identity.has_voip_phone`
	Feature string

	// The fully qualified name of the failing resolver, e.g. `my.project.get_fraud_score`.
	Resolver string
}

// ChalkHttpError is a wrapper around a standard HTTP error such as missing authorization.
type ChalkHttpError struct {
	// The URL of the HTTP request made.
	Path string

	// The message describing the error.
	Message string

	// HTTP status code of the error.
	StatusCode int

	// The size of the message body, in bytes.
	ContentLength int64

	// A Chalk Trace ID, useful for when contacting Chalk Support.
	Trace *string
}

// ChalkClientError is an error that occurred in Client or its dependencies.
type ChalkClientError struct {
	Message string
}
