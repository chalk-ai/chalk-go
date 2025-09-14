package chalk

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/apache/arrow/go/v16/arrow/memory"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
)

// GRPCClient is the gRPC interface for interacting with Chalk.
type GRPCClient interface {
	// OnlineQueryBulk computes features values using online resolvers,
	// and has the ability to query multiple primary keys at once.
	//
	// The Chalk CLI can codegen structs for all available features with
	// the [chalk codegen] command.
	//
	//	Example usages:
	//
	//      // Single-namespace online query
	//  	var users []User
	//  	res, err := chalk.OnlineQueryBulk(
	//  		context.Background(),
	//  		chalk.OnlineQueryParams{}.
	//  			WithInput(Features.User.Id, []string{"u273489056"}).
	//  			WithInput(Features.User.Transactions, [][]Transaction{
	//  				{
	//  					{Id: utils.ToPtr("txn8f76"), Amount: utils.ToPtr(13.23)},
	//  					{Id: utils.ToPtr("txn546d"), Amount: utils.ToPtr(48.95)},
	//  				},
	//  			}).
	//  			WithOutputs(Features.User.Id, Features.User.WeightedScore),
	//  	)
	//  	if err != nil {
	//  		return errors.Wrap(err, "querying weighted score")
	//  	}
	//  	if err = res.UnmarshalInto(&users); err != nil {
	//  		return errors.Wrap(err, "unmarshalling into users")
	//  	}
	//  	fmt.Println("user %s has weighted score %v", users[0].Id, users[0].WeightedScore)
	//
	//
	//  	// Multi-namespace online query
	//  	type underwriting struct {
	//  		User
	//  		Loan
	//  	}
	//
	//  	res, err := chalk.OnlineQueryBulk(
	//  		context.Background(),
	//  		chalk.OnlineQueryParams{}.
	//  			WithInput(Features.User.Id, []string{"u273489056"}).
	//  			WithInput(Features.Loan.Id, []string{"l273489056"}).
	//  			WithOutputs(
	//  				Features.User.Id,
	//  				Features.User.WeightedScore,
	//  				Features.Loan.Id,
	//  				Features.Loan.ApprovalStatus,
	//  			),
	//  	)
	//  	if err != nil {
	//  		return errors.Wrap(err, "querying weighted score and loan approval status")
	//  	}
	//
	//  	var root []underwriting
	//  	if err = res.UnmarshalInto(&root); err != nil {
	//  		return errors.Wrap(err, "unmarshalling into underwriting")
	//  	}
	//  	fmt.Println("user %s has weighted score %v", root[0].User.Id, root[0].User.WeightedScore)
	//  	fmt.Println("loan %s has approval status %v", root[0].Loan.Id, root[0].Loan.ApprovalStatus)
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	// [query basics]: https://docs.chalk.ai/docs/query-basics
	OnlineQueryBulk(ctx context.Context, params OnlineQueryParamsComplete) (*GRPCOnlineQueryBulkResult, error)
	GetOnlineQueryBulkRequest(ctx context.Context, params OnlineQueryParamsComplete) (*connect.Request[commonv1.OnlineQueryBulkRequest], error)
	GetQueryEndpoint() string

	// UpdateAggregates synchronously persists feature values that back windowed aggregations,
	// while updating the corresponding aggregate values themselves.
	// The `Inputs` parameter should be a map of features to values. The features should either
	// be a string or codegen-ed feature reference, and the values a slice of the appropriate type.
	// All slices should be the same length.
	//
	// The update is successful if the response contains no errors.
	//
	// Example:
	//
	// 		res, err := client.UpdateAggregates(
	//			context.Background(),
	// 			UpdateAggregatesParams{
	// 				Inputs: map[any]any{
	// 					Features.Txns.Id: []string{5555-5555", "4444-4444"},
	// 				    "txns.merchant_id": []string{"amezon", "pacman studios"},
	//					"txns.amount": []float64{126.58, 100.03},
	// 				},
	// 			}
	// 		)
	//      if err != nil {
	//          return errors.Wrap(err, "updating aggregates for merchant")
	//      }
	//
	// [chalk codegen]: https://docs.chalk.ai/cli#codegen
	UpdateAggregates(ctx context.Context, params UpdateAggregatesParams) (*GRPCUpdateAggregatesResult, error)

	GetAggregates(ctx context.Context, features []string) (*GRPCGetAggregatesResult, error)

	PlanAggregateBackfill(
		ctx context.Context,
		req *aggregatev1.PlanAggregateBackfillRequest,
	) (*GRPCPlanAggregateBackfillResult, error)

	// GetToken retrieves a token that can be used to authenticate requests to the Chalk API
	// along with other using the client's credentials.
	GetToken(ctx context.Context) (*TokenResult, error)

	// GetConfig retrieves the current configuration of the GRPCClient.
	// This will not necessarily be the same config as was used to create the
	// GRPCClient, as default values may have been set.
	GetConfig() *GRPCClientConfig

	GetMetadataServerInterceptor() []connect.ClientOption

	// GetGraph retrieves the graph for a deployment.
	GetGraph(ctx context.Context, deploymentId string) (*GRPCGetGraphResult, error)

	// UpdateGraph updates the graph for a deployment.
	UpdateGraph(ctx context.Context, req *serverv1.UpdateGraphRequest) (*GRPCUpdateGraphResult, error)
}

type GRPCClientConfig struct {
	ClientId      string
	ClientSecret  string
	ApiServer     string
	EnvironmentId string

	// ConfigDir specifies the directory to look for configuration files.
	// If nil, will use the default XDG_CONFIG_HOME or user home directory.
	ConfigDir *string

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

	// HTTPClient is an HTTP client instance to use when instantiating a
	// Connect gRPC-compatible client.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient connect.HTTPClient

	// Chalk can route queries to specific deployments using deployment
	// tags.
	DeploymentTag string

	// ResourceGroup specifies the resource group to route all requests to. If set
	// on the request or query level, this will be overridden.
	ResourceGroup string

	// Timeout specifies the timeout for all requests. Defaults to no timeout.
	// Timeout of 0 means no timeout. Deadline or timeout set on the request
	// context will override this timeout.
	Timeout time.Duration

	// Allocator specifies the allocator to use for creating Arrow objects.
	// Defaults to `memory.DefaultAllocator`.
	Allocator memory.Allocator

	// Interceptors are middleware functions that can intercept and modify
	// gRPC requests and responses for logging, auth, metrics, etc.
	Interceptors []connect.Interceptor

	// JWT is a valid Chalk JWT that can be used to authenticate requests
	// uncommon to be used, prefer ClientId and ClientSecret instead.
	JWT *serverv1.GetTokenResponse
}

// NewGRPCClient creates a GRPCClient with authentication settings configured.
// These settings can be overriden by passing in a GRPCClientConfig
// object. Otherwise, for each configuration variable, NewGRPCClient uses its
// corresponding environment variable if it exists. The environment variables
// that NewGRPCClient looks for are:
//
//	CHALK_ACTIVE_ENVIRONMENT
//	CHALK_API_SERVER
//	CHALK_CLIENT_ID
//	CHALK_CLIENT_SECRET
//
// For each config variable, if it is still not found, NewGRPCClient will look for a
// `~/.chalk.yml` file, which is updated when you run [chalk login].
// If a configuration for the specific project directory if found,
// that configuration will be used. Otherwise, the configuration under
// the key `default` will be used.
//
// Example:
//
//		     chalkClient, err := chalk.NewGRPCClient(
//	             context.Background(),
//		         &chalk.GRPCClientConfig{
//			         ClientId:      "id-89140a6614886982a6782106759e30",
//			         ClientSecret:  "sec-b1ba98e658d7ada4ff4c7464fb0fcee65fe2cbd86b3dd34141e16f6314267b7b",
//			         ApiServer:     "https://api.chalk.ai",
//			         EnvironmentId: "qa",
//			         Branch:        "jorges-december",
//		         },
//		     )
//
// [chalk login]: https://docs.chalk.ai/cli#login
func NewGRPCClient(ctx context.Context, configs ...*GRPCClientConfig) (GRPCClient, error) {
	return newGrpcClient(ctx, configs...)
}
