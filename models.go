package chalk

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

type Feature struct {
	Fqn string
}

// OnlineQueryParams defines the parameters
// that help you execute an online query.
// OnlineQueryParams is the starting point
// of the method chain that can help you
// obtain an object of type [OnlineQueryParamsComplete]
// that you can pass into Client.OnlineQuery.
type OnlineQueryParams struct {
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

func (result *OnlineQueryResult) GetFeature(feature any) *FeatureResult {
	castedFeature := unwrapFeatureInterface(feature)
	featureResult, found := result.features[castedFeature.Fqn]
	if !found {
		return nil
	}
	return &featureResult
}

func (result *OnlineQueryResult) GetFeatureValue(feature any) any {
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
//  2. UnmarshalInto validates that all expected output features (as specified in OnlineQueryParams)
//     are not nil pointers, and returns a ClientError otherwise.
//
//  3. UnmarshalInto also returns a ClientError if its argument is not a pointer to a struct.
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
func (result *OnlineQueryResult) UnmarshalInto(resultHolder any) *ClientError {
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

type OfflineQueryParams struct {
	Inputs         map[string][]any
	Output         []string
	RequiredOutput []string
	EnvironmentId  string
	DatasetName    string
	Branch         string
	MaxSamples     *int
}

type QueryStatus int

const (
	PENDING_SUBMISSION QueryStatus = 1
	SUBMITTED          QueryStatus = 2
	RUNNING            QueryStatus = 3
	ERROR              QueryStatus = 4
	EXPIRED            QueryStatus = 5
	CANCELLED          QueryStatus = 6
	SUCCESSFUL         QueryStatus = 7
)

type Dataset struct {
	IsFinished  bool              `json:"is_finished"`
	Version     int               `json:"version"`
	DatasetId   *string           `json:"dataset_id"`
	DatasetName *string           `json:"dataset_name"`
	Revisions   []DatasetRevision `json:"revisions"`
	Errors      []ServerError     `json:"errors"`
}

type DatasetRevision struct {
	RevisionId    string        `json:"revision_id"`
	CreatorId     string        `json:"creator_id"`
	Outputs       []string      `json:"outputs"`
	GivensUri     *string       `json:"givens_uri"`
	Status        QueryStatus   `json:"status"`
	Filters       DatasetFilter `json:"filters"`
	NumPartitions int           `json:"num_partitions"`
	OutputUris    string        `json:"output_uris"`
	OutputVersion int           `json:"output_version"`
	NumBytes      *int          `json:"num_bytes"`
	CreatedAt     *time.Time    `json:"created_at"`
	StartedAt     *string       `json:"started_at"`
	TerminatedAt  *time.Time    `json:"terminated_at"`
	DatasetName   *string       `json:"dataset_name"`
	DatasetId     *string       `json:"dataset_id"`

	client *clientImpl
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

type IDatasetRevision interface {
	DownloadData()
}

func deferFunctionWithError(function func() error, originalError error) error {
	err := originalError
	if bodyCloseErr := function(); bodyCloseErr != nil && err == nil {
		err = bodyCloseErr
	}
	return err
}

func (c *clientImpl) saveUrlToDirectory(URL string, directory string) error {
	resp, err := c.httpClient.Get(URL)
	if err != nil {
		return err
	}
	defer func() {
		err = deferFunctionWithError(resp.Body.Close, err)
	}()

	parsedUrl, urlParseErr := url.Parse(URL)
	if urlParseErr != nil {
		return urlParseErr
	}
	destinationFilepath := filepath.Join(parsedUrl.Path[4:])
	destinationDirectory := filepath.Join(directory, filepath.Dir(destinationFilepath))

	if err = os.MkdirAll(destinationDirectory, os.ModePerm); err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	destinationPath := filepath.Join(directory, destinationFilepath)
	if err = os.WriteFile(destinationPath, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (d *DatasetRevision) DownloadData(client Client, path string) *ErrorResponse {
	urls, getUrlsErr := client.GetDatasetUrls(d.RevisionId, "")
	if getUrlsErr != nil {
		return getUrlsErr
	}

	for _, url := range urls {
		saveErr := d.client.saveUrlToDirectory(url, path)
		if saveErr != nil {
			return &ErrorResponse{ClientError: &ClientError{Message: saveErr.Error()}}
		}
	}
	return nil
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
	Feature string `json:"feature"`

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
