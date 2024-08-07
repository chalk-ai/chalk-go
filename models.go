package chalk

import (
	"context"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"golang.org/x/sync/errgroup"
	"reflect"
	"time"
)

// OnlineQueryParams defines the parameters
// that help you execute an online query.
// OnlineQueryParams is the starting point
// of the method chain that can help you
// obtain an object of type [OnlineQueryParamsComplete]
// that you can pass into Client.OnlineQuery.
type OnlineQueryParams struct {
	/*************
	 PUBLIC FIELDS
	**************/

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

	// RequiredResolverTags are all the [tags] that must be present on a resolver for
	// it to be considered eligible to execute.
	// [tags]: https://docs.chalk.ai/docs/resolver-tags
	RequiredResolverTags []string

	// If specified, Chalk will route your request to the relevant preview deployment.
	PreviewDeploymentId string

	// The name for class of query you're making, for example, "loan_application_model".
	QueryName string

	// The version of the query you're making, for example, "v1".
	QueryNameVersion string

	// A globally unique ID for the query, used alongside logs and available in web
	// interfaces. If None, a correlation ID will be generated for you and returned on
	// the response.
	CorrelationId string

	// Arbitrary key:value pairs to associate with a query.
	Meta map[string]string

	// The branch id
	BranchId *string

	// Now is the time value that will be passed into [Now] dependent resolvers.
	// [Now]: https://docs.chalk.ai/docs/time#now-explicitly-time-dependent-resolvers
	Now []time.Time

	// StorePlanStages triggers storing the output of each of the query plan stages in
	// S3/GCS. This will dramatically impact the performance of the query, so it should
	// only be used for debugging. These files will be visible in the web dashboard's
	// query detail view, and can be downloaded in full by clicking on a plan node in
	// the query plan visualizer.
	StorePlanStages bool

	/**************
	 PRIVATE FIELDS
	***************/

	// The features for which there are known values, mapped to those values.
	// Set by OnlineQueryParams.WithInput.
	inputs map[string]any

	// The features that you'd like to compute from the inputs.
	// Set by OnlineQueryParams.WithOutputs.
	outputs []string

	// Maximum staleness overrides for any output features or intermediate features.
	// Set by OnlineQueryParams.WithStaleness.
	staleness map[string]time.Duration

	builderErrors BuilderErrors

	// Whether features have been versioned. Features have been versioned if
	// codegen-ed structs were used to specify inputs or outputs.
	versioned bool
}

// WithInput returns a copy of Online Query parameters with the specified inputs added.
// For use via method chaining. See [OnlineQueryParamsComplete] for usage examples.
func (p OnlineQueryParams) WithInput(feature any, value any) onlineQueryParamsWithInputs {
	return onlineQueryParamsWithInputs{underlying: p.withInput(feature, value)}
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParams) WithOutputs(features ...any) onlineQueryParamsWithOutputs {
	return onlineQueryParamsWithOutputs{underlying: p.withOutputs(features...)}
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p OnlineQueryParams) WithStaleness(feature any, duration time.Duration) OnlineQueryParams {
	return p.withStaleness(feature, duration)
}

func (p OnlineQueryParams) WithBranchId(branchId string) OnlineQueryParams {
	p.BranchId = &branchId
	return p
}

func (p OnlineQueryParams) WithQueryName(queryName string) OnlineQueryParams {
	p.QueryName = queryName
	return p
}

func (p OnlineQueryParams) WithQueryNameVersion(version string) OnlineQueryParams {
	p.QueryNameVersion = version
	return p
}

// OnlineQueryResult holds the result of an online query.
type OnlineQueryResult struct {
	// The output features and any query metadata.
	Data []FeatureResult

	// Execution metadata for the query. See QueryMeta for details.
	Meta *QueryMeta

	// Used to efficiently get a FeatureResult by FQN.
	features map[string]FeatureResult

	// Used to validate result holder expected outputs are not nil.
	expectedOutputs []string
}

// GetFeature returns a wrapper for the raw, uncasted value of the specified feature.
// To get the value of a feature as its appropriate Go type, use the UnmarshalInto method.
func (result *OnlineQueryResult) GetFeature(feature any) (*FeatureResult, error) {
	var key string
	if fqn, ok := feature.(string); ok {
		key = fqn
	} else {
		castedFeature, err := UnwrapFeature(feature)
		if err != nil {
			return nil, fmt.Errorf("exception occurred while getting feature: %w", err)
		}
		key = castedFeature.Fqn
	}
	featureResult, found := result.features[key]
	if !found {
		return nil, fmt.Errorf("feature '%s' not found in OnlineQueryResult.Data", key)
	}
	return &featureResult, nil
}

// GetFeatureValue returns the raw, uncasted value of the specified feature.
// To get the value of a feature as its appropriate Go type, use the UnmarshalInto method.
func (result *OnlineQueryResult) GetFeatureValue(feature any) (any, error) {
	featureResult, err := result.GetFeature(feature)
	if err != nil {
		return nil, err
	}
	if featureResult == nil {
		return nil, fmt.Errorf("feature object unexpectedly nil")
	}
	return featureResult.Value, nil
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
	Error *ServerError
}

// UnmarshalInto unmarshals OnlineQueryResult.Data into the specified struct (passed by pointer).
// The input argument should be a pointer to the struct that represents the output namespace.
//
//  1. UnmarshalInto populates fields corresponding to outputs specified in OnlineQueryParams,
//     while leaving all other fields as nil. If the struct has fields that point to other
//     structs (has-one relations), those nested structs will also be populated with their
//     respective feature values.
//
//  2. UnmarshalInto also returns a ClientError if its argument is not a pointer to a struct.
//
// Implicit usage example (pass result struct into OnlineQuery):
//
//	func printUserDetails(chalkClient chalk.Client) {
//		user := User{}
//		chalkClient.OnlineQuery(chalk.OnlineQueryParams{}.WithOutputs(
//			 Features.User.Family.Size,
//			 Features.User.SocureScore
//		).WithInput(Features.User.Id, 1), &user)
//
//		fmt.Println("User family size: ", *user.Family.Size)
//		fmt.Println("User Socure score: ", *user.SocureScore)
//	}
//
// Equivalent explicit usage example:
//
//	func printUserDetails(chalkClient chalk.Client) {
//		result, _ := chalkClient.OnlineQuery(chalk.OnlineQueryParams{}.WithOutputs(
//			Features.User.Family.Size,
//			Features.User.SocureScore
//		).WithInput(Features.User.Id, 1), nil)
//
//		user := User{}
//		result.UnmarshalInto(&user)
//
//		fmt.Println("User family size: ", *user.Family.Size)
//		fmt.Println("User Socure score: ", *user.SocureScore)
//	}
func (result *OnlineQueryResult) UnmarshalInto(resultHolder any) (returnErr *ClientError) {
	defer func() {
		if panicContents := recover(); panicContents != nil {
			detail := "details irretrievable"
			switch typedContents := panicContents.(type) {
			case *reflect.ValueError:
				detail = typedContents.Error()
			case string:
				detail = typedContents
			}
			returnErr = &ClientError{Message: fmt.Sprintf("exception occurred while unmarshalling result: %s", detail)}
		}
	}()

	value := reflect.ValueOf(resultHolder)
	kind := value.Type().Kind()
	if kind != reflect.Pointer {
		return &ClientError{Message: fmt.Sprintf("argument should be a pointer, got '%s' instead", kind.String())}
	}

	kindPointedTo := value.Elem().Kind()
	if kindPointedTo != reflect.Struct {
		return &ClientError{Message: fmt.Sprintf("argument should be pointer to a struct, got a pointer to a '%s' instead", kindPointedTo.String())}
	}

	return result.unmarshal(resultHolder)
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
	// Execution duration in seconds
	ExecutionDurationS float64 `json:"execution_duration_s"`

	// The id of the deployment that served this query.
	DeploymentId string `json:"deployment_id"`

	// The id of the environment that served this query. Not intended to be human-readable,
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

// OnlineQueryBulkResult holds the result of a bulk online query.
type OnlineQueryBulkResult struct {
	// ScalarsTable is an Arrow Table containing
	// scalar features of the target feature class.
	ScalarsTable arrow.Table

	// GroupsTables is a map from a has-many feature to its
	// corresponding Arrow Table.
	GroupsTables map[string]arrow.Table

	// Execution metadata for the query. See QueryMeta for details.
	Meta *QueryMeta
}

// UnmarshalInto unmarshals Arrow tables in OnlineQueryBulkResult into the specified slice of
// structs (passed by pointer). The input argument should be a pointer to the empty slice of
// structs that represents the output namespace.
//
//  1. UnmarshalInto populates fields corresponding to outputs specified in OnlineQueryParams,
//     while leaving all other fields as nil. If the struct has fields that point to other
//     structs (has-one relations), those nested structs will also be populated with their
//     respective feature values.
//
//  2. UnmarshalInto also returns a ClientError if its argument is not a pointer to a list of
//     structs.
//
//  3. UnmarshalInto does not currently handle unmarshalling:
//     a. has-many features
//     In the meantime, you can manually unmarshal a has-many feature across all root
//     structs by using the UnmarshalTableInto function.
//     b. struct-like features, such as Python `dataclass` or Pydantic models
//
// Usage example:
//
//	func printUserDetails(chalkClient chalk.Client) {
//		result, _ := chalkClient.OnlineQueryBulk(chalk.OnlineQueryParams{}.WithOutputs(
//			Features.User.Family.Size,
//			Features.User.SocureScore
//		).WithInput(Features.User.Id, []int{1, 2}), nil)
//
//		var users []User
//		result.UnmarshalInto(&users)
//
//		fmt.Println("User 1 family size: ", *user[0].Family.Size)
//		fmt.Println("User 2 Socure score: ", *user[1].SocureScore)
//	}
func (r *OnlineQueryBulkResult) UnmarshalInto(resultHolders any) *ClientError {
	if err := unmarshalTableInto(r.ScalarsTable, resultHolders); err != nil {
		return &ClientError{Message: err.Error()}
	}
	return nil
}

// UploadFeaturesParams defines the parameters
// that help you execute an upload features request.
type UploadFeaturesParams struct {
	// Inputs is a map of features to values. The features should
	// either be a string or codegen-ed Feature object. The values
	// should be a slice of the appropriate type. All slices should
	// be the same length as the number of entities you want to upload
	// features for.
	Inputs map[any]any

	// EnvironmentOverride is the environment to which you want to upload
	// features. If not specified, defaults to the environment specified
	// in the client configuration.
	EnvironmentOverride string

	// PreviewDeploymentId is the preview deployment to which you want to upload
	// features. If not specified, defaults to the main deployment.
	PreviewDeploymentId string
}

// UploadFeaturesResult holds the result of an upload features request.
type UploadFeaturesResult struct {
	OperationId string `json:"operation_id"`
}

// UpdateAggregatesParams lets you specify
// parameters for updating aggregated values
// of your windowed aggregation features.
type UpdateAggregatesParams struct {
	// Inputs is a map of features to values. The features should
	// either be a string or codegen-ed Feature object. The values
	// should be a slice of the appropriate type. All slices should
	// be the same length as the number of entities you want to upload
	// features for.
	Inputs map[any]any

	Context context.Context
}

// UpdateAggregatesResult holds the result of an upload aggregates request.
type UpdateAggregatesResult struct {
	TraceId string `json:"trace_id"`
}

// OfflineQueryParams defines the parameters
// that help you execute an online query.
// OfflineQueryParams is the starting point
// of the method chain that can help you
// obtain an object of type [OfflineQueryParamsComplete]
// that you can pass into Client.OfflineQuery.
type OfflineQueryParams struct {
	/**************
	 PUBLIC FIELDS
	**************/

	// The environment under which to run the resolvers.
	// API tokens can be scoped to an environment.
	// If no environment is specified in the query,
	// but the token supports only a single environment,
	// then that environment will be taken as the scope
	// for executing the request.
	EnvironmentId string

	// A unique name that if provided will be used to generate and
	// save a Dataset constructed from the list of features computed
	// from the inputs.
	DatasetName string

	// The branch under which to run the resolvers.
	Branch string

	// The maximum number of samples to include in the `DataFrame`.
	MaxSamples *int

	// DefaultTime indicates the default time at which you would like to observe the features.
	// If not specified, the current time will be used as the default observation time.
	// The default observation time will be used when:
	// 1. A feature value is passed into [OfflineQueryParams.WithInput] as a [TsFeatureValue] with a nil time.
	// 2. A feature value is passed into [OfflineQueryParams.WithInput] as a raw value (not a [TsFeatureValue]).
	// For more information about observation time, see https://docs.chalk.ai/docs/temporal-consistency
	DefaultTime *time.Time

	// The tags used to scope the resolvers.
	Tags []string

	/***************
	 PRIVATE FIELDS
	***************/

	inputs          map[string][]TsFeatureValue
	outputs         []string
	requiredOutputs []string
	builderErrors   BuilderErrors
	// Whether features have been versioned. Features have been versioned if
	// codegen-ed structs were used to specify inputs or outputs.
	versioned bool
}

// WithInput returns a copy of Offline Query parameters with the specified inputs added.
// For use via method chaining. See [OfflineQueryParamsComplete] for usage examples.
// The "values" argument can contain a raw value (int or string), or it can also contain
// a [TsFeatureValue] if you want to query with a specific observation time. The observation
// time for raw values will be the default observation time specified as [OfflineQueryParams.DefaultTime].
// If no default observation time is specified, the current time will be used.
// For more information about observation time, see [Temporal Consistency].
//
// [Temporal Consistency]: https://docs.chalk.ai/docs/temporal-consistency
func (p OfflineQueryParams) WithInput(feature any, values []any) offlineQueryParamsWithInputs {
	return offlineQueryParamsWithInputs{underlying: p.withInput(feature, values)}
}

// WithOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParams) WithOutputs(features ...any) OfflineQueryParamsComplete {
	return OfflineQueryParamsComplete{underlying: p.withOutputs(features...)}
}

// WithRequiredOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParams) WithRequiredOutputs(features ...any) OfflineQueryParamsComplete {
	return OfflineQueryParamsComplete{underlying: p.withRequiredOutputs(features...)}
}

// TsFeatureValue is a struct that can be passed to OfflineQueryParams.WithInput
// to specify the value of a feature along with a timestamp. This timestamp indicates
// the observation time at which you would like the output feature values to be queried.
// For more information about observation time, see [Temporal Consistency].
//
// [Temporal Consistency]: https://docs.chalk.ai/docs/temporal-consistency
type TsFeatureValue struct {
	// The value of the feature. In the context of offline query,
	// this is always a value of a primary feature.
	Value any

	// The observation time at which you would like the output
	// feature values to be queried. If nil, [OfflineQueryParams.DefaultTime]
	// will be used as the observation time. If [OfflineQueryParams.DefaultTime]
	// is also nil, the current time will be used as the observation time.
	ObservationTime *time.Time
}

// QueryStatus represents the status of an offline query.
type QueryStatus int

const (
	// QueryStatusPendingSubmission to the database.
	QueryStatusPendingSubmission QueryStatus = 1

	// QueryStatusSubmitted to the database, but not yet running.
	QueryStatusSubmitted QueryStatus = 2

	// QueryStatusRunning in the database.
	QueryStatusRunning QueryStatus = 3

	// QueryStatusError with either submitting or running the job.
	QueryStatusError QueryStatus = 4

	// QueryStatusExpired indicates the job did not complete before an expiration
	// deadline, so there are no results.
	QueryStatusExpired QueryStatus = 5

	// QueryStatusCancelled indicates the job was manually cancelled before it
	// errored or finished successfully.
	QueryStatusCancelled QueryStatus = 6

	// QueryStatusSuccessful indicates the job successfully ran.
	QueryStatusSuccessful QueryStatus = 7
)

type Dataset struct {
	// Whether the export job is finished (it runs asynchronously)
	IsFinished bool `json:"is_finished"`

	// Version number representing the format of the data. The client
	// uses this version number to properly decode and load the query
	// results into DataFrames.
	Version     int               `json:"version"`
	DatasetId   *string           `json:"dataset_id"`
	DatasetName *string           `json:"dataset_name"`
	Revisions   []DatasetRevision `json:"revisions"`
	Errors      []ServerError     `json:"errors"`
}

type DatasetRevision struct {
	// UUID for the revision job.
	RevisionId string `json:"revision_id"`

	// UUID for the creator of the job.
	CreatorId string `json:"creator_id"`

	// Output features for the dataset revision.
	Outputs []string `json:"outputs"`

	// Location of the givens stored for the dataset.
	GivensUri *string `json:"givens_uri"`

	// Status of the revision job.
	Status QueryStatus `json:"status"`

	// Filters performed on the dataset.
	Filters DatasetFilter `json:"filters"`

	// Number of partitions for revision job.
	NumPartitions int `json:"num_partitions"`

	// Location of the outputs stored fo the dataset.
	OutputUris string `json:"output_uris"`

	// Storage version of the outputs.
	OutputVersion int `json:"output_version"`

	// Number of bytes of the output, updated upon success.
	NumBytes *int `json:"num_bytes"`

	// Timestamp for creation of revision job.
	CreatedAt *time.Time `json:"created_at"`

	// Timestamp for start of revision job.
	StartedAt *string `json:"started_at"`

	// Timestamp for end of revision job.
	TerminatedAt *time.Time `json:"terminated_at"`

	// Name of revision, if given.
	DatasetName *string `json:"dataset_name"`

	// ID of revision, if name is given.
	DatasetId *string `json:"dataset_id"`

	client *clientImpl
}

// DownloadData downloads output files pertaining to the revision to given path.
// Datasets are stored in Chalk as sharded Parquet files. With this
// method, you can download those raw files into a directory for processing
// with other tools.
func (d *DatasetRevision) DownloadData(directory string) error {
	urls, getUrlsErr := d.client.getDatasetUrls(d.RevisionId, "")
	if getUrlsErr != nil {
		return getUrlsErr
	}
	g, _ := errgroup.WithContext(context.Background())
	for _, url := range urls {
		// Capture the loop variables in the closure.
		u := url
		g.Go(func() error {
			return d.client.saveUrlToDirectory(u, directory)
		})
	}
	saveErr := g.Wait()
	if saveErr != nil {
		return &ErrorResponse{ClientError: &ClientError{Message: saveErr.Error()}}
	}
	return nil
}

type ColumnMetadata struct {
	FeatureFqn string `json:"feature_fqn"`
	ColumnName string `json:"column_name"`
	Dtype      string `json:"dtype"`
}

type GetOfflineQueryJobResponse struct {
	IsFinished bool             `json:"is_finished"`
	Version    int              `json:"version"`
	Urls       []string         `json:"urls"`
	Errors     []ServerError    `json:"errors"`
	Columns    []ColumnMetadata `json:"columns"`
}

type TriggerResolverRunParams struct {
	// ResolverFqn is the fully qualified name of the offline resolver to trigger.
	ResolverFqn string `json:"resolver_fqn"`

	// EnvironmentId is the environment under which you'd like to query your data.
	EnvironmentId string `json:"environment_id"`

	// PreviewDeploymentId, if specified, will be used by Chalk to route
	// your request to the relevant preview deployment.
	PreviewDeploymentId string `json:"preview_deployment_id"`
}

type TriggerResolverRunResult struct {
	// Id is the ID of the offline resolver run.
	Id string `json:"id"`

	// Status is the current status of the resolver run.
	Status string `json:"status"`
}

type GetRunStatusParams struct {
	// RunId is the ID of the resolver run to check.
	RunId string `json:"resolver_fqn"`

	// PreviewDeploymentId, if specified, will be used by Chalk to route
	// your request to the relevant preview deployment.
	PreviewDeploymentId string `json:"preview_deployment_id"`
}

type GetRunStatusResult struct {
	// Id is the ID of the resolver run.
	Id string `json:"id"`

	// Status is the current status of the resolver run.
	Status string `json:"status"`
}

type ResolverException struct {
	// The name of the class of the exception.
	Kind string `json:"kind"`

	// The message taken from the exception.
	Message string `json:"message"`

	// The stacktrace produced by the code.
	Stacktrace string `json:"stacktrace"`
}

type ErrorResponse struct {
	// Errors that occurred in Chalk's server.
	ServerErrors []ServerError

	// Errors that occurred in Client or its dependencies.
	ClientError *ClientError

	// Errors that are standard HTTP errors such as missing authorization.
	HttpError *HTTPError
}

// ServerError is an error that occurred in Chalk's server,
// for example, when a resolver unexpectedly fails to run.
type ServerError struct {
	// The type of the error.
	Code ErrorCode `json:"code"`

	// The category of the error, given in the type field for the error codes.
	// This will be one of "REQUEST", "NETWORK", and "FIELD".
	Category ErrorCodeCategory `json:"category"`

	// A readable description of the error message.
	Message string `json:"message"`

	// The exception that caused the failure, if applicable.
	Exception *ResolverException `json:"exception"`

	// The fully qualified name of the failing feature, e.g. `user.identity.has_voip_phone`
	Feature string `json:"12feature"`

	// The fully qualified name of the failing resolver, e.g. `my.project.get_fraud_score`.
	Resolver string `json:"resolver"`
}

// HTTPError is a wrapper around a standard HTTP error such as missing authorization.
type HTTPError struct {
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

// ClientError is an error that occurred in Client or its dependencies.
type ClientError struct {
	Message string
}

// TokenResult holds the result of a GetToken request.
type TokenResult struct {
	// The access token that can be used to authenticate requests to the Chalk API.
	AccessToken string `json:"access_token"`

	// The primary environment that the token is scoped to.
	PrimaryEnvironment string `json:"primary_environment"`

	// The time at which the token expires.
	ValidUntil time.Time `json:"valid_until"`

	// The GRPC endpoint for the engine.
	Engines map[string]string `json:"engines"`
}
