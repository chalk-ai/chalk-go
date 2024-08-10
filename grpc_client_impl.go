package chalk

import (
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	"net/http"
)

type grpcClientImpl struct {
	GRPCClient
	config *configManager

	branch     string
	logger     LeveledLogger
	httpClient *http.Client

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newGrpcClient(cfg ClientConfig) (*grpcClientImpl, error) {
	config, err := getConfigManager(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}
	var logger LeveledLogger
	if cfg.Logger == nil {
		logger = DefaultLeveledLogger
	}
	client := &grpcClientImpl{
		branch:     cfg.Branch,
		httpClient: http.DefaultClient,
		logger:     logger,
		config:     config,
	}
	if err := client.init(); err != nil {
		return nil, errors.Wrap(err, "error initializing gRPC service clients")
	}

	return client, nil
}

func (c *grpcClientImpl) init() error {
	//authClient, err := c.NewAuthClient()
	//if err != nil {
	//	return errors.Wrap(err, "error creating auth client")
	//}
	//c.authClient = authClient
	//
	//c.config.getToken = c.getToken
	//// Necessary to get GRPC engines URL
	//if err := c.config.refresh(false); err != nil {
	//	return errors.Wrap(err, "error fetching initial config")
	//}
	//
	//queryClient, err := c.NewQueryClient()
	//if err != nil {
	//	return errors.Wrap(err, "error creating query client")
	//}
	//c.queryClient = queryClient
	//
	//return nil
	return nil
}
