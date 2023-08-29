package chalk

import (
	"net/http"
)

// Client is the primary interface for interacting with Chalk. You can use
// it to query data, trigger resolver runs, gather offline data, and more.
type Client interface {
	// OnlineQuery computes features values using online resolvers.
	// See [query basics] for more information.
	//
	// resultHolder is a pointer to the struct that the result should be
	// unmarshalled into. The struct passed in should be the struct that
	// represents the feature set corresponding to the root output namespace.
	// For instance, in the example below, 'user' is the root output
	// namespace, so a pointer to a 'User' struct is passed in.
	// You can also choose to pass 'nil' as the resultHolder, in which case
	// you should use OnlineQueryResult.UnmarshalInto to populate a struct.
	//
	// The Chalk CLI can codegen structs for all available features with
	// the [chalk codegen] command.
	//
	// Example:
	//
	//		user := User{}
	//		res, err := client.OnlineQuery(
	//			OnlineQueryParams{
	//				IncludeMeta: true,
	//				EnvironmentId: "pipkjlfc3gtmn",
	//			}.
	//			WithInput(Features.User.Card.Id, 4).
	//			WithOutputs(Features.User.Email, Features.User.Card.Id),
	//			&user,
	//		)
	//		fmt.Println("User email: ", user.Email)
	//		fmt.Println("User card ID: ", user.Card.Id)
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	// [query basics]: https://docs.chalk.ai/docs/query-basics
	OnlineQuery(args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, *ErrorResponse)

	// OnlineQueryBulk computes features values using online resolvers,
	// and has the ability to query multiple primary keys at once.
	// OnlineQueryBulk also has the ability to return a has-many feature
	// in the form of an arrow.Record. The usual features are returned
	// in bulk in an arrow.Record too, with each column name corresponding
	// to the feature name.
	//
	// The Chalk CLI can codegen structs for all available features with
	// the [chalk codegen] command.
	//
	// Example:
	//
	//		res, err := client.OnlineQueryBulk(
	//			OnlineQueryParams{
	//				IncludeMeta: true,
	//				EnvironmentId: "pipkjlfc3gtmn",
	//			}.
	//			WithInput(Features.User.Card.Id, []int{1, 2, 3, 4}).
	//			WithOutputs(Features.User.Email, Features.User.Card.Id),
	//		)
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	// [query basics]: https://docs.chalk.ai/docs/query-basics
	OnlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, *ErrorResponse)

	// OfflineQuery queries feature values from the offline store.
	// See Dataset for more information.
	//
	// Example:
	//
	//		client.OfflineQuery(
	//			OfflineQueryParams{
	//				EnvironmentId: "pipkjlfc3gtmn",
	//			}.
	//	 		WithRequiredOutputs(Features.User.Email, Features.User.Card.Id),
	//		)
	//
	OfflineQuery(args OfflineQueryParamsComplete) (Dataset, *ErrorResponse)

	// TriggerResolverRun triggers an offline resolver to run.
	// See https://docs.chalk.ai/docs/runs for more information.
	TriggerResolverRun(args TriggerResolverRunParams) (TriggerResolverRunResult, *ErrorResponse)

	// GetRunStatus retrieves the status of an offline resolver run.
	// See https://docs.chalk.ai/docs/runs for more information.
	GetRunStatus(args GetRunStatusParams) (GetRunStatusResult, *ErrorResponse)
}

type ClientConfig struct {
	ClientId      string
	ClientSecret  string
	ApiServer     string
	EnvironmentId string

	// If specified, Chalk will route all requests from this client
	// instance to the relevant branch.
	Branch string

	// Logger is the logger that the backend will use to log errors,
	// warnings, and informational messages.
	//
	// LeveledLogger is implemented by StdOutLeveledLogger, and one can be
	// initialized at the desired level of logging.  LeveledLogger
	// also provides out-of-the-box compatibility with a Logrus Logger, but may
	// require a thin shim for use with other logging libraries that use less
	// standard conventions like Zap.
	//
	// Defaults to DefaultLeveledLogger.
	//
	// To set a logger that logs nothing, set this to a chalk.LeveledLogger
	// with a Level of LevelNull (simply setting this field to nil will not
	// work).
	Logger LeveledLogger

	// HTTPClient is an HTTP client instance to use when making API requests.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient *http.Client
}

// NewClient creates a Client with authentication settings configured.
// These settings can be overriden by passing in a ClientConfig
// object. Otherwise, for each configuration variable, NewClient uses its
// corresponding environment variable if it exists. The environment variables
// that NewClient looks for are:
//
//	CHALK_ACTIVE_ENVIRONMENT
//	CHALK_API_SERVER
//	CHALK_CLIENT_ID
//	CHALK_CLIENT_SECRET
//
// For each config variable, if it is still not found, NewClient will look for a
// `~/.chalk.yml` file, which is updated when you run [chalk login].
// If a configuration for the specific project directory if found,
// that configuration will be used. Otherwise, the configuration under
// the key `default` will be used.
//
// [chalk login]: https://docs.chalk.ai/cli#login
func NewClient(config ...*ClientConfig) (Client, error) {
	return newClientImpl(config...)
}
