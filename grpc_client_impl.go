package chalk

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/chalk-ai/chalk-go/auth"

	"connectrpc.com/connect"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/config"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
)

type grpcClientImpl struct {
	GRPCClient
	config    *config.Manager
	allocator memory.Allocator

	branch        string
	queryServer   *string
	resourceGroup *string
	logger        LeveledLogger
	httpClient    connect.HTTPClient
	timeout       *time.Duration

	queryClient      enginev1connect.QueryServiceClient
	graphClient      serverv1connect.GraphServiceClient
	tokenInterceptor connect.UnaryInterceptorFunc
	deploymentTag    string
	tokenManager     *auth.Manager
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
	if cfg.Logger == nil {
		cfg.Logger = DefaultLeveledLogger
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	if cfg.Allocator == nil {
		cfg.Allocator = memory.DefaultAllocator
	}

	c, err := config.NewManager(
		ctx,
		config.NewFromArg[string](cfg.ApiServer),
		config.NewFromArg[config.ClientId](config.ClientId(cfg.ClientId)),
		config.NewFromArg[config.ClientSecret](config.ClientSecret(cfg.ClientSecret)),
		config.NewFromArg[string](cfg.EnvironmentId),
		cfg.ConfigDir,
	)
	if err != nil {
		return nil, errors.Wrap(err, "getting resolved config")
	}
	tokenManager, err := auth.NewManager(
		ctx,
		&auth.Inputs{
			Token:      cfg.JWT,
			HttpClient: cfg.HTTPClient,
			Manager:    c,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "initializing token manager")
	}

	var timeout *time.Duration
	if cfg.Timeout != 0 { // If unspecified (zero value)
		timeout = &cfg.Timeout
	}
	var resourceGroup *string
	if cfg.ResourceGroup != "" {
		resourceGroup = &cfg.ResourceGroup
	}

	resolvedQueryServer := tokenManager.GetQueryServerURL(cfg.QueryServer)
	if strings.HasPrefix(resolvedQueryServer, "http://") {
		// Unsecured client
		// From https://connectrpc.com/docs/go/deployment#h2c
		cfg.HTTPClient = &http.Client{
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
	tokenInterceptor := makeTokenInterceptor(tokenManager)
	engineInterceptors := []connect.Interceptor{
		tokenInterceptor,
		headerInterceptor(headers),
	}
	if timeout != nil {
		engineInterceptors = append(engineInterceptors, timeoutInterceptor(timeout))
	}

	queryClient := enginev1connect.NewQueryServiceClient(
		cfg.HTTPClient,
		resolvedQueryServer,
		connect.WithInterceptors(append(cfg.Interceptors, engineInterceptors...)...),
		connect.WithGRPC(),
	)

	// Create GraphServiceClient with API server endpoint
	apiServerURL := c.ApiServer.Value
	apiInterceptors := []connect.Interceptor{
		tokenInterceptor,
		headerInterceptor(map[string]string{
			HeaderKeyServerType: serverTypeApi,
		}),
	}
	if timeout != nil {
		apiInterceptors = append(apiInterceptors, timeoutInterceptor(timeout))
	}

	graphClient := serverv1connect.NewGraphServiceClient(
		cfg.HTTPClient,
		apiServerURL,
		connect.WithInterceptors(append(cfg.Interceptors, apiInterceptors...)...),
		connect.WithGRPC(),
	)

	return &grpcClientImpl{
		deploymentTag:    cfg.DeploymentTag,
		branch:           cfg.Branch,
		httpClient:       cfg.HTTPClient,
		logger:           cfg.Logger,
		config:           c,
		tokenManager:     tokenManager,
		queryClient:      queryClient,
		graphClient:      graphClient,
		queryServer:      ptr.OrNil(cfg.QueryServer),
		resourceGroup:    resourceGroup,
		timeout:          timeout,
		allocator:        cfg.Allocator,
		tokenInterceptor: tokenInterceptor,
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
	Allocator memory.Allocator
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
		if opt.Allocator != nil {
			allocator = opt.Allocator
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
	return c.tokenManager.GetQueryServerURL("")
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
		ClientId:      string(c.config.ClientId.Value),
		ClientSecret:  string(c.config.ClientSecret.Value),
		ApiServer:     c.config.ApiServer.Value,
		EnvironmentId: c.tokenManager.GetEnvironmentId(),
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
	res, err := c.tokenManager.GetJWT(ctx, time.Now().Add(time.Minute))
	if err != nil {
		return nil, errors.Wrap(err, "getting JWT token")
	}

	return &TokenResult{
		AccessToken:        res.AccessToken,
		ValidUntil:         res.ExpiresAt.AsTime(),
		PrimaryEnvironment: c.tokenManager.GetEnvironmentId(),
		Engines:            res.Engines,
	}, nil
}

type GRPCGetGraphResult struct {
	RawResponse *serverv1.GetGraphResponse
}

func (c *grpcClientImpl) GetGraph(ctx context.Context, deploymentId string) (*GRPCGetGraphResult, error) {
	req := connect.NewRequest(&serverv1.GetGraphRequest{
		DeploymentId: deploymentId,
	})
	
	res, err := c.graphClient.GetGraph(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "getting graph")
	}

	return &GRPCGetGraphResult{RawResponse: res.Msg}, nil
}

type GRPCUpdateGraphResult struct {
	RawResponse *serverv1.UpdateGraphResponse
}

func (c *grpcClientImpl) UpdateGraph(ctx context.Context, req *serverv1.UpdateGraphRequest) (*GRPCUpdateGraphResult, error) {
	connectReq := connect.NewRequest(req)
	
	res, err := c.graphClient.UpdateGraph(ctx, connectReq)
	if err != nil {
		return nil, errors.Wrap(err, "updating graph")
	}

	return &GRPCUpdateGraphResult{RawResponse: res.Msg}, nil
}
