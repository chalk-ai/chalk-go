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
	httpClient := http.DefaultClient
	authClient, err := NewAuthClient(httpClient, config.apiServer.Value)
	if err != nil {
		return nil, errors.Wrap(err, "error creating auth client")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = DefaultLeveledLogger
	}
	config.getToken = MakeGetTokenFunc(logger, authClient)

	// Necessary to get GRPC engines URL
	if err := config.refresh(false); err != nil {
		return nil, errors.Wrap(err, "error fetching initial config")
	}

	queryClient, err := NewQueryClient(httpClient, config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating query client")
	}

	return &grpcClientImpl{
		branch:      cfg.Branch,
		httpClient:  httpClient,
		logger:      logger,
		config:      config,
		authClient:  authClient,
		queryClient: queryClient,
	}, nil
}
