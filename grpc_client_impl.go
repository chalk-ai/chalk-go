package chalk

import (
	"connectrpc.com/connect"
	"context"
	"crypto/tls"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
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
	config    *configManager
	allocator memory.Allocator

	branch        string
	queryServer   *string
	resourceGroup *string
	logger        LeveledLogger
	httpClient    HTTPClient
	timeout       *time.Duration

	authClient       serverv1connect.AuthServiceClient
	queryClient      enginev1connect.QueryServiceClient
	tokenInterceptor connect.UnaryInterceptorFunc
	deploymentTag    string
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

	allocator := memory.DefaultAllocator
	if cfg.Allocator != nil {
		allocator = cfg.Allocator
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
	tokenInterceptor := makeTokenInterceptor(config)
	engineInterceptors := []connect.Interceptor{
		tokenInterceptor,
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
		deploymentTag:    cfg.DeploymentTag,
		branch:           cfg.Branch,
		httpClient:       httpClient,
		logger:           config.logger,
		config:           config,
		authClient:       authClient,
		queryClient:      queryClient,
		queryServer:      queryServer,
		resourceGroup:    resourceGroup,
		timeout:          timeout,
		allocator:        allocator,
		tokenInterceptor: tokenInterceptor,
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

type FeatureMeta struct {
	ResolverFqn string
	SourceType  string
	SourceId    string
}

type FeatureOutput struct {
	Fqn   string
	Value any
	Meta  *FeatureMeta
}

type RowResult struct {
	Features map[string]FeatureOutput
}

func newRowResult() *RowResult {
	return &RowResult{
		Features: make(map[string]FeatureOutput),
	}
}

// GetFeature takes in a feature string or a codegen'd
// feature reference and returns the `FeatureOutput` object.
// Given this codegen'd snippet:
//
//	type User struct {
//	 	Id                       *int64
//	 	FullName                 *string
//	}
//
//	var Features struct {
//	 	User *User
//	}
//
//	func init() {
//	 	InitFeaturesErr = chalk.InitFeatures(&Features)
//	}
//
// You would get the feature object for "user.full_name" as follows:
//
//	feature, err := row.GetFeature(Features.User.FullName)
func (r *RowResult) GetFeature(feature any) (*FeatureOutput, error) {
	fqn, ok := feature.(string)
	if !ok {
		unwrapped, err := UnwrapFeature(feature)
		if err != nil {
			return nil, errors.Wrap(err, "please provide a feature string or a codegen'd feature reference")
		}
		fqn = unwrapped.Fqn
	}
	res, ok := r.Features[fqn]
	if !ok {
		return nil, errors.Newf("feature '%s' not found", fqn)
	}
	return &res, nil
}

func (r *RowResult) GetFeatureValue(feature any) (any, error) {
	res, err := r.GetFeature(feature)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

type GRPCOnlineQueryBulkResult struct {
	RawResponse *commonv1.OnlineQueryBulkResponse
	allocator   memory.Allocator
}

type NewGRPCOnlineQueryBulkResultOptions struct {
	allocator memory.Allocator
}

// NewGRPCOnlineQueryBulkResult creates a GRPCOnlineQueryBulkResult
// for testing. This function sets up a result object with Arrow
// artifacts such as a `memory.Allocator` which is required during
// unmarshalling operations.
func NewGRPCOnlineQueryBulkResult(
	response *commonv1.OnlineQueryBulkResponse,
	options ...NewGRPCOnlineQueryBulkResultOptions,
) (*GRPCOnlineQueryBulkResult, error) {
	allocator := memory.DefaultAllocator
	if len(options) == 1 {
		opt := options[0]
		if opt.allocator != nil {
			allocator = opt.allocator
		}
	} else if len(options) > 1 {
		return nil, errors.Newf("expected only one set of options, found %d", len(options))
	}
	return &GRPCOnlineQueryBulkResult{
		RawResponse: response,
		allocator:   allocator,
	}, nil
}

func (r *GRPCOnlineQueryBulkResult) GetTable() (arrow.Table, error) {
	return internal.ConvertBytesToTable(r.RawResponse.GetScalarsData(), r.allocator)
}

func (r *GRPCOnlineQueryBulkResult) GetRow(rowIndex int) (*RowResult, error) {
	row := newRowResult()
	if len(r.RawResponse.GetScalarsData()) == 0 {
		return nil, errors.New("results table empty, either the query has errors or the data is malformed")
	}

	scalarsTable, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData(), r.allocator)
	if err != nil {
		return nil, errors.Wrap(err, "converting scalars data to table")
	}

	rows, meta, err := internal.ExtractFeaturesFromTable(scalarsTable, false)
	if err != nil {
		return nil, errors.Wrap(err, "extracting features from scalars table")
	}

	if rowIndex < 0 || rowIndex >= len(rows) {
		return nil, errors.Newf(
			"out of bounds: accessing index %d of table with %d rows",
			rowIndex, len(rows),
		)
	}

	var rowMeta map[string]internal.FeatureMeta
	if len(meta) > 0 {
		if len(meta) != len(rows) {
			return nil, errors.Newf(
				"metadata length %v does not match rows length %v",
				len(meta), len(rows),
			)
		}
		rowMeta = meta[rowIndex]
	}

	for fqn, value := range rows[rowIndex] {
		featureRes := FeatureOutput{
			Fqn:   fqn,
			Value: value,
		}
		if rowMeta != nil {
			internalMeta, ok := rowMeta[fqn]
			if !ok {
				// Features such as has-many features do not have a metadata column.
				continue
			}
			featureRes.Meta = &FeatureMeta{
				ResolverFqn: internalMeta.ResolverFqn,
				SourceType:  internalMeta.SourceType,
				SourceId:    internalMeta.SourceId,
			}
		}
		row.Features[fqn] = featureRes
	}

	return row, nil
}

func (r *GRPCOnlineQueryBulkResult) GetQueryMeta() *QueryMeta {
	return queryMetaFromProto(r.RawResponse.GetResponseMeta())
}

func (r *GRPCOnlineQueryBulkResult) GetErrors() ([]ServerError, error) {
	return serverErrorsFromProto(r.RawResponse.GetErrors())
}

func (r *GRPCOnlineQueryBulkResult) UnmarshalInto(resultHolders any) error {
	allocator := r.allocator
	if allocator == nil {
		allocator = memory.DefaultAllocator
	}
	scalars, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData(), allocator)
	if err != nil {
		return errors.Wrap(err, "deserializing scalars table")
	}
	return internal.UnmarshalTableInto(scalars, resultHolders)
}

func (c *grpcClientImpl) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (*GRPCOnlineQueryBulkResult, error) {
	req, err := c.GetOnlineQueryBulkRequest(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, "generating online query request")
	}
	res, err := c.queryClient.OnlineQueryBulk(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "executing online query")
	}

	result := &GRPCOnlineQueryBulkResult{RawResponse: res.Msg, allocator: c.allocator}
	if len(res.Msg.GetErrors()) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Msg.GetErrors())
		if err != nil {
			return nil, errors.Wrap(err, "converting proto errors")
		}
		// Must return result even upon error, since there could be partial results
		return result, convertedErrs
	}
	return result, nil
}

func (c *grpcClientImpl) GetOnlineQueryBulkRequest(ctx context.Context, args OnlineQueryParamsComplete) (*connect.Request[commonv1.OnlineQueryBulkRequest], error) {
	paramsProto, err := convertOnlineQueryParamsToProto(&args.underlying, c.allocator)
	if err != nil {
		return nil, errors.Wrap(err, "converting online query params to proto")
	}
	req := connect.NewRequest(paramsProto)
	if args.underlying.ResourceGroup != "" {
		req.Header().Set(HeaderKeyResourceGroup, args.underlying.ResourceGroup)
	} else if c.resourceGroup != nil {
		req.Header().Set(HeaderKeyResourceGroup, *c.resourceGroup)
	}
	return req, nil
}

func (c *grpcClientImpl) GetQueryEndpoint() string {
	return c.config.apiServer.Value
}

func (c *grpcClientImpl) GetMetadataServerInterceptor() []connect.ClientOption {
	//httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption
	return []connect.ClientOption{
		connect.WithInterceptors(
			headerInterceptor(map[string]string{
				HeaderKeyServerType: serverTypeApi,
			}),
			c.tokenInterceptor,
		),
	}
}

func (c *grpcClientImpl) GetConfig() *GRPCClientConfig {
	return &GRPCClientConfig{
		ClientId:      c.config.clientId.Value,
		ClientSecret:  c.config.clientSecret.Value,
		ApiServer:     c.config.apiServer.Value,
		EnvironmentId: c.config.environmentId.Value,
		Branch:        c.branch,
		QueryServer:   ptr.OrZero(c.queryServer),
		Logger:        c.logger,
		HTTPClient:    c.httpClient,
		DeploymentTag: c.deploymentTag,
		ResourceGroup: ptr.OrZero(c.resourceGroup),
		Timeout:       ptr.OrZero(c.timeout),
		Allocator:     c.allocator,
	}
}

type GRPCUpdateAggregatesResult struct {
	RawResponse *commonv1.UploadFeaturesBulkResponse
}

func (r *GRPCUpdateAggregatesResult) GetErrors() ([]ServerError, error) {
	return serverErrorsFromProto(r.RawResponse.GetErrors())
}

func (c *grpcClientImpl) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (*GRPCUpdateAggregatesResult, error) {
	inputsConverted, err := getConvertedInputsMap(args.Inputs)
	if err != nil {
		return nil, errors.Wrap(err, "converting inputs map")
	}
	inputsFeather, err := internal.InputsToArrowBytes(inputsConverted, c.allocator)
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

	result := &GRPCUpdateAggregatesResult{RawResponse: res.Msg}
	if len(res.Msg.GetErrors()) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Msg.GetErrors())
		if err != nil {
			return nil, errors.Wrap(err, "converting proto errors")
		}
		return result, convertedErrs
	}
	return result, nil
}

type GRPCGetAggregatesResult struct {
	RawResponse *aggregatev1.GetAggregatesResponse
}

func (c *grpcClientImpl) GetAggregates(ctx context.Context, features []string) (*GRPCGetAggregatesResult, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "making get aggregates request")
	}

	result := &GRPCGetAggregatesResult{RawResponse: res.Msg}
	if len(res.Msg.GetErrors()) > 0 {
		var allErrors []error
		for _, errStr := range res.Msg.GetErrors() {
			allErrors = append(allErrors, errors.New(errStr))
		}
		return result, errors.Join(allErrors...)
	}
	return result, nil
}

type GRPCPlanAggregateBackfillResult struct {
	RawResponse *aggregatev1.PlanAggregateBackfillResponse
}

func (c *grpcClientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*GRPCPlanAggregateBackfillResult, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, errors.Wrap(err, "making plan aggregate backfill request")
	}

	result := &GRPCPlanAggregateBackfillResult{RawResponse: res.Msg}
	if len(res.Msg.GetErrors()) > 0 {
		var allErrors []error
		for _, errStr := range res.Msg.GetErrors() {
			allErrors = append(allErrors, errors.New(errStr))
		}
		return result, errors.Join(allErrors...)
	}
	return result, nil
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
