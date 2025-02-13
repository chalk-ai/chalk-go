package chalk

import (
	"context"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"time"
)

// GRPCClient is the gRPC-native interface for interacting with Chalk.
// Our existing Client interface also works with gRPC, but this interface
// is more idiomatic for talking to our gRPC endpoints.
type GRPCClient interface {
	OnlineQuery(ctx context.Context, params OnlineQueryParamsComplete) (*commonv1.OnlineQueryResponse, error)

	OnlineQueryBulk(ctx context.Context, params OnlineQueryParamsComplete) (*commonv1.OnlineQueryBulkResponse, error)

	UpdateAggregates(ctx context.Context, params UpdateAggregatesParams) (*commonv1.UploadFeaturesBulkResponse, error)

	GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error)

	PlanAggregateBackfill(
		ctx context.Context,
		req *aggregatev1.PlanAggregateBackfillRequest,
	) (*aggregatev1.PlanAggregateBackfillResponse, error)

	// GetToken retrieves a token that can be used to authenticate requests to the Chalk API
	// along with other using the client's credentials.
	GetToken() (*TokenResult, error)
}

type GRPCClientConfig struct {
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

	// HTTPClient is an HTTP client instance to use when instantiating a
	// Connect gRPC-compatible client.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient HTTPClient

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
}

// NewGRPCClient creates a GRPCClient with authentication settings configured.
// These settings can be overriden by passing in a ClientConfig
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
//	     chalkClient, err := chalk.NewGRPCClient(&chalk.ClientConfig{
//		        ClientId:      "id-89140a6614886982a6782106759e30",
//		        ClientSecret:  "sec-b1ba98e658d7ada4ff4c7464fb0fcee65fe2cbd86b3dd34141e16f6314267b7b",
//		        ApiServer:     "https://api.chalk.ai",
//		        EnvironmentId: "qa",
//		        Branch:        "jorges-december",
//	})
//
// [chalk login]: https://docs.chalk.ai/cli#login
func NewGRPCClient(configs ...*GRPCClientConfig) (GRPCClient, error) {
	return newGrpcClient(configs...)
}
