package chalk

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"strings"
	"time"
)

type grpcClientImpl struct {
	GRPCClient
	config *configManager

	branch        string
	queryServer   *string
	resourceGroup *string
	logger        LeveledLogger
	httpClient    HTTPClient
	timeout       *time.Duration

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newGrpcClient(ctx context.Context, configs ...*GRPCClientConfig) (*grpcClientImpl, error) {
	var cfg *GRPCClientConfig
	if len(configs) == 0 {
		cfg = &GRPCClientConfig{}
	} else if len(configs) == 1 {
		cfg = configs[len(configs)-1]
	} else {
		return nil, errors.Newf("expected at most one GRPCClientConfig, got %d", len(configs))
	}

	config, err := newConfigManager(cfg.ApiServer, cfg.ClientId, cfg.ClientSecret, cfg.EnvironmentId, cfg.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "getting resolved config")
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var timeout *time.Duration
	if cfg.Timeout != 0 { // If unspecified (zero value)
		timeout = &cfg.Timeout
	}

	authInterceptors := []connect.Interceptor{
		headerInterceptor(map[string]string{
			HeaderKeyServerType: serverTypeApi,
		}),
	}
	if timeout != nil {
		authInterceptors = append(authInterceptors, timeoutInterceptor(timeout))
	}
	authClient := serverv1connect.NewAuthServiceClient(
		httpClient,
		config.apiServer.Value,
		connect.WithInterceptors(authInterceptors...),
	)

	config.getToken = func(ctx context.Context, clientId string, clientSecret string) (*getTokenResult, error) {
		return getToken(ctx, clientId, clientSecret, config.logger, authClient)
	}

	// Necessary to get GRPC engines URL
	if err := config.refresh(ctx, false); err != nil {
		return nil, errors.Wrap(err, "fetching initial config")
	}

	var queryServer *string
	if cfg.QueryServer != "" {
		queryServer = &cfg.QueryServer
	}

	var resourceGroup *string
	if cfg.ResourceGroup != "" {
		resourceGroup = &cfg.ResourceGroup
	}

	resolvedQueryServer := config.getQueryServer(queryServer)
	if strings.HasPrefix(resolvedQueryServer, "http://") {
		// Unsecured client
		// From https://connectrpc.com/docs/go/deployment#h2c
		httpClient = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}
	}

	headers := map[string]string{
		HeaderKeyDeploymentType: "engine-grpc",
		HeaderKeyServerType:     serverTypeEngine,
	}
	if cfg.DeploymentTag != "" {
		headers[HeaderKeyDeploymentTag] = cfg.DeploymentTag
	}
	engineInterceptors := []connect.Interceptor{
		makeTokenInterceptor(config),
		headerInterceptor(headers),
	}
	if timeout != nil {
		engineInterceptors = append(engineInterceptors, timeoutInterceptor(timeout))
	}

	queryClient := enginev1connect.NewQueryServiceClient(
		httpClient,
		ensureHTTPSPrefix(resolvedQueryServer),
		connect.WithInterceptors(engineInterceptors...),
		connect.WithGRPC(),
	)

	return &grpcClientImpl{
		branch:        cfg.Branch,
		httpClient:    httpClient,
		logger:        config.logger,
		config:        config,
		authClient:    authClient,
		queryClient:   queryClient,
		queryServer:   queryServer,
		resourceGroup: resourceGroup,
		timeout:       timeout,
	}, nil
}

func getToken(ctx context.Context, clientId string, clientSecret string, logger LeveledLogger, client serverv1connect.AuthServiceClient) (*getTokenResult, error) {
	logger.Debugf("Getting new token via gRPC")
	authRequest := connect.NewRequest(
		&serverv1.GetTokenRequest{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
		},
	)

	token, err := client.GetToken(ctx, authRequest)
	if err != nil {
		logger.Debugf("Failed to get a new token: %s", err.Error())
		return nil, err
	}

	expiresAt := token.Msg.GetExpiresAt()
	if token.Msg.GetExpiresAt() == nil {
		return nil, errors.New("token has no expiration date")
	}

	return &getTokenResult{
		ValidUntil:         expiresAt.AsTime(),
		AccessToken:        token.Msg.GetAccessToken(),
		PrimaryEnvironment: token.Msg.GetPrimaryEnvironment(),
		Engines:            token.Msg.GetGrpcEngines(),
	}, nil
}

type QueryExplainInfo struct {
	Plan string
}

type OnlineQueryMetadata struct {
	ExecutionDuration *time.Duration
	DeploymentId      string
	EnvironmentId     string
	EnvironmentName   string
	QueryId           string
	QueryTimestamp    *time.Time
	QueryHash         string
	ExplainOutput     *QueryExplainInfo
}
type GRPCOnlineQueryBulkResult struct {
	RawTable []byte
	Meta     OnlineQueryMetadata
	Errors   []ServerError

	RawResponse *commonv1.OnlineQueryBulkResponse
}

func (c *grpcClientImpl) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (*GRPCOnlineQueryBulkResult, error) {
	paramsProto, err := convertOnlineQueryParamsToProto(&args.underlying)
	if err != nil {
		return nil, errors.Wrap(err, "converting online query params to proto")
	}
	req := connect.NewRequest(paramsProto)
	if args.underlying.ResourceGroup != "" {
		req.Header().Set(HeaderKeyResourceGroup, args.underlying.ResourceGroup)
	} else if c.resourceGroup != nil {
		req.Header().Set(HeaderKeyResourceGroup, *c.resourceGroup)
	}
	res, err := c.queryClient.OnlineQueryBulk(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "executing online query")
	}

	grpcMeta := res.Msg.GetResponseMeta()

	var explainOutput *QueryExplainInfo
	if grpcMeta.GetExplainOutput() != nil {
		grpcExplain := grpcMeta.GetExplainOutput()
		explainOutput = &QueryExplainInfo{
			Plan: grpcExplain.GetPlanString(),
		}
	}

	meta := OnlineQueryMetadata{
		ExecutionDuration: ptr.Ptr(grpcMeta.GetExecutionDuration().AsDuration()),
		DeploymentId:      grpcMeta.GetDeploymentId(),
		EnvironmentId:     grpcMeta.GetEnvironmentId(),
		EnvironmentName:   grpcMeta.GetEnvironmentName(),
		QueryId:           grpcMeta.GetQueryId(),
		QueryTimestamp:    ptr.Ptr(grpcMeta.GetQueryTimestamp().AsTime()),
		QueryHash:         grpcMeta.GetQueryHash(),
		ExplainOutput:     explainOutput,
	}

	serverErrors, err := serverErrorsFromProto(res.Msg.GetErrors())
	if err != nil {
		return nil, errors.Wrapf(err, "converting raw proto errors: %v", serverErrors)
	}

	return &GRPCOnlineQueryBulkResult{
		RawTable:    res.Msg.GetScalarsData(),
		Meta:        meta,
		Errors:      serverErrors,
		RawResponse: res.Msg,
	}, nil
}

type GRPCUpdateAggregatesResult struct {
	Errors      []ServerError
	RawResponse *commonv1.UploadFeaturesBulkResponse
}

func (c *grpcClientImpl) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (*GRPCUpdateAggregatesResult, error) {
	inputsConverted, err := getConvertedInputsMap(args.Inputs)
	if err != nil {
		return nil, errors.Wrap(err, "converting inputs map")
	}
	inputsFeather, err := internal.InputsToArrowBytes(inputsConverted)
	if err != nil {
		return nil, errors.Wrap(err, "serializing inputs as feather")
	}

	req := connect.NewRequest(&commonv1.UploadFeaturesBulkRequest{
		InputsFeather: inputsFeather,
		BodyType:      commonv1.FeatherBodyType_FEATHER_BODY_TYPE_TABLE,
	})

	res, err := c.queryClient.UploadFeaturesBulk(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making update aggregates request")
	}

	serverErrors, err := serverErrorsFromProto(res.Msg.GetErrors())
	if err != nil {
		return nil, errors.Wrapf(err, "converting raw proto errors: %v", serverErrors)
	}
	return &GRPCUpdateAggregatesResult{
		Errors:      serverErrors,
		RawResponse: res.Msg,
	}, nil
}

type AggregateTimeSeriesRule struct {
	Aggregation       string
	BucketDuration    *time.Duration
	DependentFeatures []string
	Rentention        *time.Duration
	DatetimeFeature   string
}

type AggregateTimeSeries struct {
	Namespace          string
	AggregateOn        string
	GroupBy            []string
	Rules              []AggregateTimeSeriesRule
	FiltersDescription string
	BucketFeature      string
}

type GRPCGetAggregatesResult struct {
	Series      []AggregateTimeSeries
	RawResponse *aggregatev1.GetAggregatesResponse
}

func aggregateTimeSeriesFromProto(raw []*aggregatev1.AggregateTimeSeries) []AggregateTimeSeries {
	var series []AggregateTimeSeries
	for _, oneSeries := range raw {
		var rules []AggregateTimeSeriesRule
		for _, rawRule := range oneSeries.GetRules() {
			var bucketDuration *time.Duration
			if rawRule.GetBucketDuration() != nil {
				bucketDuration = ptr.Ptr(rawRule.GetBucketDuration().AsDuration())
			}

			var retention *time.Duration
			if rawRule.GetRetention() != nil {
				retention = ptr.Ptr(rawRule.GetRetention().AsDuration())
			}

			rules = append(rules, AggregateTimeSeriesRule{
				Aggregation:       rawRule.GetAggregation(),
				BucketDuration:    bucketDuration,
				DependentFeatures: rawRule.GetDependentFeatures(),
				Rentention:        retention,
				DatetimeFeature:   rawRule.GetDatetimeFeature(),
			})
		}
		series = append(series, AggregateTimeSeries{
			Namespace:          oneSeries.GetNamespace(),
			AggregateOn:        oneSeries.GetAggregateOn(),
			GroupBy:            oneSeries.GetGroupBy(),
			Rules:              rules,
			FiltersDescription: oneSeries.GetFiltersDescription(),
			BucketFeature:      oneSeries.GetBucketFeature(),
		})
	}
	return series
}

func (c *grpcClientImpl) GetAggregates(ctx context.Context, features []string) (*GRPCGetAggregatesResult, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making get aggregates request")
	}

	return &GRPCGetAggregatesResult{
		Series:      aggregateTimeSeriesFromProto(res.Msg.GetSeries()),
		RawResponse: res.Msg,
	}, nil
}

type AggregateBackfillCostEstimate struct {
	MaxBuckets          int64
	ExpectedBuckets     int64
	ExpectedBytes       int64
	ExpectedStorageCost float64
	ExpectedRuntime     *time.Duration
}

type AggregateBackfill struct {
	Series             []AggregateTimeSeries
	Resolver           string
	DatetimeFeature    string
	BucketDuration     *time.Duration
	FiltersDescription string
	GroupBy            []string
	MaxRetention       *time.Duration
	LowerBound         *time.Time
	UpperBound         *time.Time
}

type AggregateBackfillWithCostEstimate struct {
	Backfill *AggregateBackfill
	Estimate *AggregateBackfillCostEstimate
}

type GRPCPlanAggregateBackfillResult struct {
	Estimate            *AggregateBackfillCostEstimate
	Errors              []ServerError
	Backfills           []AggregateBackfillWithCostEstimate
	AggregateBackfillId string

	RawResponse *aggregatev1.PlanAggregateBackfillResponse
}

func aggregateBackfillCostEstimateFromProto(raw *aggregatev1.AggregateBackfillCostEstimate) *AggregateBackfillCostEstimate {
	var estimate *AggregateBackfillCostEstimate
	if raw != nil {
		var runtime *time.Duration
		if raw.GetExpectedRuntime() != nil {
			runtime = ptr.Ptr(raw.GetExpectedRuntime().AsDuration())
		}
		estimate = &AggregateBackfillCostEstimate{
			MaxBuckets:          raw.GetMaxBuckets(),
			ExpectedBuckets:     raw.GetExpectedBuckets(),
			ExpectedBytes:       raw.GetExpectedBytes(),
			ExpectedStorageCost: raw.GetExpectedStorageCost(),
			ExpectedRuntime:     runtime,
		}
	}
	return estimate
}

func (c *grpcClientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*GRPCPlanAggregateBackfillResult, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "making plan aggregate backfill request")
	}

	var backfills []AggregateBackfillWithCostEstimate
	for _, backfill := range res.Msg.GetBackfills() {
		var innerBackfill *AggregateBackfill
		rawBackfill := backfill.GetBackfill()
		if rawBackfill != nil {
			var bucketDuration *time.Duration
			if rawBackfill.GetBucketDuration() != nil {
				bucketDuration = ptr.Ptr(rawBackfill.GetBucketDuration().AsDuration())
			}

			var maxRetention *time.Duration
			if rawBackfill.GetMaxRetention() != nil {
				maxRetention = ptr.Ptr(rawBackfill.GetMaxRetention().AsDuration())
			}

			var lowerBound *time.Time
			if rawBackfill.GetLowerBound() != nil {
				lowerBound = ptr.Ptr(rawBackfill.GetLowerBound().AsTime())
			}

			var upperBound *time.Time
			if rawBackfill.GetUpperBound() != nil {
				upperBound = ptr.Ptr(rawBackfill.GetUpperBound().AsTime())
			}

			innerBackfill = &AggregateBackfill{
				Series:             aggregateTimeSeriesFromProto(rawBackfill.GetSeries()),
				Resolver:           rawBackfill.GetResolver(),
				DatetimeFeature:    rawBackfill.GetDatetimeFeature(),
				BucketDuration:     bucketDuration,
				FiltersDescription: rawBackfill.GetFiltersDescription(),
				GroupBy:            rawBackfill.GetGroupBy(),
				MaxRetention:       maxRetention,
				LowerBound:         lowerBound,
				UpperBound:         upperBound,
			}
		}

		backfills = append(backfills, AggregateBackfillWithCostEstimate{
			Backfill: innerBackfill,
			Estimate: aggregateBackfillCostEstimateFromProto(backfill.GetEstimate()),
		})
	}

	return &GRPCPlanAggregateBackfillResult{
		Estimate:    aggregateBackfillCostEstimateFromProto(res.Msg.GetEstimate()),
		RawResponse: res.Msg,
	}, err
}

func (c *grpcClientImpl) GetToken(ctx context.Context) (*TokenResult, error) {
	res, err := c.config.getToken(ctx, c.config.clientId.Value, c.config.clientSecret.Value)
	if err != nil {
		return nil, err
	}
	return &TokenResult{
		AccessToken:        res.AccessToken,
		ValidUntil:         res.ValidUntil,
		PrimaryEnvironment: res.PrimaryEnvironment,
		Engines:            res.Engines,
	}, nil
}
