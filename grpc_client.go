package chalk

import (
	"context"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/cockroachdb/errors"
)

// GRPCClient is the gRPC-native interface for interacting with Chalk.
// Our existing Client interface also works with gRPC, but this interface
// is more idiomatic for talking to our gRPC endpoints.
type GRPCClient interface {
	UpdateAggregates(ctx context.Context, params UpdateAggregatesParams) (*commonv1.UploadFeaturesBulkResponse, error)

	GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error)

	PlanAggregateBackfill(
		ctx context.Context,
		req *aggregatev1.PlanAggregateBackfillRequest,
	) (*aggregatev1.PlanAggregateBackfillResponse, error)
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
func NewGRPCClient(configs ...*ClientConfig) (GRPCClient, error) {
	var cfg *ClientConfig
	if len(configs) == 0 {
		cfg = &ClientConfig{}
	} else if len(configs) > 1 {
		return nil, errors.Newf("expected at most one ClientConfig, got %d", len(configs))
	} else {
		cfg = configs[len(configs)-1]
	}

	return newGrpcClient(*cfg)
}
