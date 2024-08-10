package chalk

import (
	"context"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	"github.com/cockroachdb/errors"
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
	OnlineQuery(args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error)

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
	// import "github.com/apache/arrow/go/v16/arrow/array"
	//
	//
	//		res, err := client.OnlineQueryBulk(
	//			OnlineQueryParams{
	//				IncludeMeta: true,
	//				EnvironmentId: "pipkjlfc3gtmn",
	//			}.
	//			WithInput("user.id", []int{1, 2, 3, 4}).
	//			WithOutputs("user.email", "user.transactions"),
	//		)
	//		defer res.Release()
	//
	//		reader := array.NewTableReader(res.ScalarsTable, 10_000)
	//		defer reader.Release()
	//		for reader.Next() {
	//		    record := reader.Record()
	//		    // Do something with the record
	//		}
	//
	//		txnTable := res.GroupsTables["user.transactions"]
	//		txnReader := array.NewTableReader(txnTable, 10_000)
	//		// Do something with the txnReader
	//		defer txnReader.Release()
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	// [query basics]: https://docs.chalk.ai/docs/query-basics
	OnlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error)

	// UploadFeatures synchronously persists feature values to the online store and
	// offline store. The `Inputs` parameter should be a map of features to values.
	// The features should either be a string or codegen-ed Feature object. The values
	// should be a slice of the appropriate type. All slices should be the same length
	// as the number of entities you want to upload features for.
	//
	// The upload is successful if the response contains no errors.
	//
	// Example:
	//
	// 		res, err := client.UploadFeatures(
	// 			UploadFeaturesParams{
	// 				Inputs: map[any]any{
	// 					Features.User.Card.Id: []string{"5555-5555-5555-5555", "4444-4444-4444-4444"},
	// 				    "new_user_model.card_id": []string{"5555-5555-5555-5555", "4444-4444-4444-4444"},
	//					"new_user_model.email": []string{"borges@chalk.ai", "jorge@chalk.ai"},
	// 				},
	//              BranchOverride: "jorge-december",
	// 			}
	// 		)
	//      if err != nil {
	//          return err.Error()
	//      }
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	UploadFeatures(args UploadFeaturesParams) (UploadFeaturesResult, error)

	// UpdateAggregates synchronously persists feature values that back windowed aggregations,
	// while updating the corresponding aggregate values themselves.
	// The `Inputs` parameter should be a map of features to values. The features should either
	// be a string or codegen-ed Feature object, and the values a slice of the appropriate type.
	// All slices should be the same length.
	//
	// The update is successful if the response contains no errors.
	//
	// Example:
	//
	// 		res, err := client.UpdateAggregates(
	// 			UpdateAggregatesParams{
	// 				Inputs: map[any]any{
	// 					Features.Txns.Id: []string{5555-5555", "4444-4444"},
	// 				    "txns.merchant_id": []string{"amezon", "pacman studios"},
	//					"txns.amount": []float64{126.58, 100.03},
	// 				},
	// 			}
	// 		)
	//      if err != nil {
	//          return err.Error()
	//      }
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	UpdateAggregates(args UpdateAggregatesParams) (UpdateAggregatesResult, error)

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
	OfflineQuery(args OfflineQueryParamsComplete) (Dataset, error)

	// TriggerResolverRun triggers an offline resolver to run.
	// See https://docs.chalk.ai/docs/runs for more information.
	TriggerResolverRun(args TriggerResolverRunParams) (TriggerResolverRunResult, error)

	// GetRunStatus retrieves the status of an offline resolver run.
	// See https://docs.chalk.ai/docs/runs for more information.
	GetRunStatus(args GetRunStatusParams) (GetRunStatusResult, error)

	// GetToken retrieves a token that can be used to authenticate requests to the Chalk API
	// along with other using the client's credentials.
	GetToken() (*TokenResult, error)

	GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error)

	PlanAggregateBackfill(
		ctx context.Context,
		req *aggregatev1.PlanAggregateBackfillRequest,
	) (*aggregatev1.PlanAggregateBackfillResponse, error)
}

type ClientConfig struct {
	ClientId      string
	ClientSecret  string
	ApiServer     string
	EnvironmentId string

	// If specified, Chalk will route all requests from this client
	// instance to the relevant branch.
	Branch string

	// Chalk routes performance sensitive requests like online query
	// directly to the query server that runs the engine. Populate
	// this field if you would like to route these requests to a
	// different query server than the one automatically resolved
	// by Chalk.
	QueryServer string

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
	HTTPClient HTTPClient

	// UseGrpc, if set to true, will create a gRPC client instead of a REST client.
	UseGrpc bool
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
// Example:
//
//	     chalkClient, chalkClientErr := chalk.NewClient(&chalk.ClientConfig{
//		        ClientId:      "id-89140a6614886982a6782106759e30",
//		        ClientSecret:  "sec-b1ba98e658d7ada4ff4c7464fb0fcee65fe2cbd86b3dd34141e16f6314267b7b",
//		        ApiServer:     "https://api.chalk.ai",
//		        EnvironmentId: "qa",
//		        Branch:        "jorges-december",
//	})
//
// [chalk login]: https://docs.chalk.ai/cli#login
func NewClient(configs ...*ClientConfig) (Client, error) {
	var cfg *ClientConfig
	if len(configs) == 0 {
		cfg = &ClientConfig{}
	} else if len(configs) > 1 {
		return nil, errors.Newf("expected at most one ClientConfig, got %d", len(configs))
	} else {
		cfg = configs[len(configs)-1]
	}

	if cfg.UseGrpc {
		return newClientGrpc(*cfg)
	}

	return newClientImpl(*cfg)
}
