package chalk

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/chalk-ai/chalk-go/auth"

	"connectrpc.com/connect"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/chalk-ai/chalk-go/config"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/http2"
)

type grpcClientImpl struct {
	GRPCClient
	config    *config.Manager
	allocator memory.Allocator

	branch        string
	resourceGroup *string
	logger        LeveledLogger
	httpClient    connect.HTTPClient
	timeout       *time.Duration

	featureWriteClient    featureWriteClient
	queryClient           enginev1connect.QueryServiceClient
	branchQueryClient     enginev1connect.QueryServiceClient
	graphClient           serverv1connect.GraphServiceClient
	datasetMetadataClient serverv1connect.DatasetMetadataServiceClient
	deploymentTag         string
	tokenManager          *auth.Manager
	metadataInterceptor   connect.UnaryInterceptorFunc
	engineInterceptor     connect.UnaryInterceptorFunc
	tracerProvider        trace.TracerProvider
	tracer                trace.Tracer
}

type featureWriteClient interface {
	UploadFeatures(context.Context, UploadFeaturesParams) (UploadFeaturesResult, error)
	DeleteFeatures(context.Context, DeleteFeaturesParams) (DeleteFeaturesResult, error)
}

type grpcHTTPClientAdapter struct {
	connect.HTTPClient
}

func (c grpcHTTPClientAdapter) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func asHTTPClient(client connect.HTTPClient) HTTPClient {
	if client, ok := client.(HTTPClient); ok {
		return client
	}
	return grpcHTTPClientAdapter{HTTPClient: client}
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
		if cfg.InsecureSkipVerify {
			// Create HTTP client with TLS config that skips certificate verification
			cfg.HTTPClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
		} else {
			cfg.HTTPClient = http.DefaultClient
		}
	}
	if cfg.Allocator == nil {
		cfg.Allocator = memory.DefaultAllocator
	}

	configManager, err := config.NewManager(
		ctx,
		&config.ManagerInputs{
			APIServer:       cfg.ApiServer,
			GRPCQueryServer: cfg.QueryServer,
			ClientId:        config.ClientId(cfg.ClientId),
			ClientSecret:    config.ClientSecret(cfg.ClientSecret),
			EnvironmentId:   cfg.EnvironmentId,
			ConfigDir:       cfg.ConfigDir,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "getting resolved config")
	}
	var timeout *time.Duration
	if cfg.Timeout != 0 {
		timeout = &cfg.Timeout
	}
	tokenManager, err := auth.NewManager(
		ctx,
		&auth.Inputs{
			Token:                      cfg.JWT,
			HttpClient:                 cfg.HTTPClient,
			Config:                     configManager,
			Timeout:                    timeout,
			SkipEnvironmentNameMapping: cfg.SkipEnvironmentNameMapping,
			SkipEngineMapping:          cfg.SkipEngineMapping,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "initializing token manager")
	}

	var resourceGroup *string
	if cfg.ResourceGroup != "" {
		resourceGroup = &cfg.ResourceGroup
	}

	resolvedQueryServer := configManager.GetGRPCQueryServer().Value
	if strings.HasPrefix(resolvedQueryServer, "http://") {
		// Unsecured client
		// From https://connectrpc.com/docs/go/deployment#h2c
		cfg.HTTPClient = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(_ context.Context, network, addr string, tlsConfig *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}
	} else if cfg.InsecureSkipVerify {
		// For HTTPS with InsecureSkipVerify, we need HTTP/2 support
		cfg.HTTPClient = &http.Client{
			Transport: &http2.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	clientBranch := cfg.Branch
	// makeEngineInterceptor builds a connect interceptor for QueryService
	// requests. The serverHeader value is sent as x-chalk-server, which the
	// upstream Envoy uses to route the request: "engine" goes to the engine
	// cluster (used by queryClient, which talks directly to the resolved
	// query server), while "go-api" goes to the API server's branch
	// router (used by branchQueryClient, which talks to api.chalk.ai). This
	// mirrors the Python SDK, where the API_SERVER channel is built with
	// server="go-api" and the engine channel with server="engine".
	makeEngineInterceptor := func(serverHeader string) func(connect.UnaryFunc) connect.UnaryFunc {
		return func(next connect.UnaryFunc) connect.UnaryFunc {
			return func(
				ctx context.Context,
				req connect.AnyRequest,
			) (connect.AnyResponse, error) {
				if timeout != nil {
					if _, deadlineSet := ctx.Deadline(); !deadlineSet {
						timeoutCtx, cancel := context.WithTimeout(ctx, *timeout)
						ctx = timeoutCtx
						defer cancel()
					}
				}

				if req.Header().Get("x-chalk-branch-id") == "" && clientBranch != "" {
					req.Header().Set("x-chalk-branch-id", clientBranch)
				}
				if req.Header().Get("x-chalk-branch-id") != "" {
					req.Header().Set("x-chalk-deployment-type", "branch-grpc")
				} else {
					req.Header().Set("x-chalk-deployment-type", "engine-grpc")
				}
				req.Header().Set("x-chalk-server", serverHeader)
				req.Header().Set("User-Agent", internal.UserAgent())
				if cfg.DeploymentTag != "" {
					req.Header().Set("x-chalk-deployment-tag", cfg.DeploymentTag)
				}

				if envId := tokenManager.GetConfig().EnvironmentId.Value; envId != "" {
					req.Header().Set("x-chalk-env-id", envId)
				}
				token, err := tokenManager.GetJWT(ctx, time.Now().Add(time.Minute))
				if err != nil {
					return nil, errors.Wrap(err, "error refreshing config")
				}
				req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
				return next(ctx, req)
			}
		}
	}

	// engineInterceptor is retained on the client struct for callers of
	// GetEngineServerInterceptor (which expect the engine routing).
	engineInterceptor := makeEngineInterceptor("engine")
	branchQueryInterceptor := makeEngineInterceptor("go-api")

	queryClient := enginev1connect.NewQueryServiceClient(
		cfg.HTTPClient,
		resolvedQueryServer,
		connect.WithInterceptors(cfg.Interceptors...),
		connect.WithInterceptors(connect.UnaryInterceptorFunc(engineInterceptor)),
		connect.WithGRPC(),
	)

	apiServerURL := configManager.GetAPIServer().Value
	branchQueryClient := enginev1connect.NewQueryServiceClient(
		cfg.HTTPClient,
		apiServerURL,
		connect.WithInterceptors(cfg.Interceptors...),
		connect.WithInterceptors(connect.UnaryInterceptorFunc(branchQueryInterceptor)),
		connect.WithGRPC(),
	)

	// Create GraphServiceClient with API server endpoint
	authedServerInterceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if timeout != nil {
				if _, deadlineSet := ctx.Deadline(); !deadlineSet {
					timeoutCtx, cancel := context.WithTimeout(ctx, *timeout)
					ctx = timeoutCtx
					defer cancel()
				}
			}
			req.Header().Set("x-chalk-server", "go-api")
			req.Header().Set("User-Agent", internal.UserAgent())
			if envId := tokenManager.GetConfig().EnvironmentId.Value; envId != "" {
				req.Header().Set("x-chalk-env-id", envId)
			}
			token, err := tokenManager.GetJWT(ctx, time.Now().Add(time.Minute))
			if err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
			return next(ctx, req)
		}
	}

	graphClient := serverv1connect.NewGraphServiceClient(
		cfg.HTTPClient,
		apiServerURL,
		connect.WithInterceptors(connect.UnaryInterceptorFunc(authedServerInterceptor)),
		connect.WithInterceptors(cfg.Interceptors...),
	)

	datasetMetadataClient := serverv1connect.NewDatasetMetadataServiceClient(
		cfg.HTTPClient,
		apiServerURL,
		connect.WithInterceptors(connect.UnaryInterceptorFunc(authedServerInterceptor)),
		connect.WithInterceptors(cfg.Interceptors...),
	)

	featureWriteClient := &clientImpl{
		config:        configManager,
		allocator:     cfg.Allocator,
		Branch:        cfg.Branch,
		DeploymentTag: cfg.DeploymentTag,
		resourceGroup: resourceGroup,
		timeout:       timeout,
		httpClient:    asHTTPClient(cfg.HTTPClient),
		logger:        cfg.Logger,
		tokenManager:  tokenManager,
	}

	return &grpcClientImpl{
		deploymentTag:         cfg.DeploymentTag,
		branch:                cfg.Branch,
		httpClient:            cfg.HTTPClient,
		logger:                cfg.Logger,
		config:                configManager,
		tokenManager:          tokenManager,
		featureWriteClient:    featureWriteClient,
		queryClient:           queryClient,
		branchQueryClient:     branchQueryClient,
		graphClient:           graphClient,
		datasetMetadataClient: datasetMetadataClient,
		resourceGroup:         resourceGroup,
		timeout:               timeout,
		allocator:             cfg.Allocator,
		metadataInterceptor:   authedServerInterceptor,
		engineInterceptor:     engineInterceptor,
		tracerProvider:        cfg.TracerProvider,
		tracer:                newTracer(cfg.TracerProvider),
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
	RawResponse     *commonv1.OnlineQueryBulkResponse
	ResponseHeaders *http.Header
	allocator       memory.Allocator

	// ctx and tracer are carried over from the OnlineQueryBulk call that
	// produced this result so that the deserialization done by the methods
	// below, which happens after the query returns and takes no context of
	// its own, still lands in the caller's trace. Both are nil for results
	// built by NewGRPCOnlineQueryBulkResult.
	ctx    context.Context
	tracer trace.Tracer
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
	_, span := startSpan(r.ctx, r.tracer, spanGetTable)
	if span.IsRecording() {
		span.SetAttributes(attrScalarsBytes.Int(len(r.RawResponse.GetScalarsData())))
	}
	table, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData(), r.allocator)
	if err == nil && table != nil && span.IsRecording() {
		span.SetAttributes(attrNumRows.Int64(table.NumRows()))
	}
	endSpan(span, err)
	return table, err
}

func (r *GRPCOnlineQueryBulkResult) GetRow(rowIndex int) (*RowResult, error) {
	ctx, span := startSpan(r.ctx, r.tracer, spanGetRow)
	if span.IsRecording() {
		span.SetAttributes(
			attrRowIndex.Int(rowIndex),
			attrScalarsBytes.Int(len(r.RawResponse.GetScalarsData())),
		)
	}
	row, err := r.getRow(ctx, rowIndex)
	endSpan(span, err)
	return row, err
}

func (r *GRPCOnlineQueryBulkResult) getRow(ctx context.Context, rowIndex int) (*RowResult, error) {
	row := newRowResult()
	if len(r.RawResponse.GetScalarsData()) == 0 {
		return nil, errors.New("results table empty, either the query has errors or the data is malformed")
	}

	scalarsTable, err := r.deserializeScalars(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "converting scalars data to table")
	}

	_, extractSpan := startSpan(ctx, r.tracer, spanExtractFeatures)
	if extractSpan.IsRecording() {
		extractSpan.SetAttributes(attrNumRows.Int64(scalarsTable.NumRows()))
	}
	rows, meta, err := internal.ExtractFeaturesFromTable(scalarsTable, false)
	endSpan(extractSpan, err)
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
	ctx, span := startSpan(r.ctx, r.tracer, spanUnmarshalInto)
	if span.IsRecording() {
		span.SetAttributes(attrScalarsBytes.Int(len(r.RawResponse.GetScalarsData())))
	}
	err := r.unmarshalInto(ctx, span, resultHolders)
	endSpan(span, err)
	return err
}

func (r *GRPCOnlineQueryBulkResult) unmarshalInto(ctx context.Context, span trace.Span, resultHolders any) error {
	scalars, err := r.deserializeScalars(ctx)
	if err != nil {
		return errors.Wrap(err, "deserializing scalars table")
	}
	if span.IsRecording() {
		span.SetAttributes(attrNumRows.Int64(scalars.NumRows()))
	}

	_, unmarshalSpan := startSpan(ctx, r.tracer, spanUnmarshalTable)
	if unmarshalSpan.IsRecording() {
		unmarshalSpan.SetAttributes(attrNumRows.Int64(scalars.NumRows()))
	}
	err = internal.UnmarshalTableInto(scalars, resultHolders)
	endSpan(unmarshalSpan, err)
	return err
}

// deserializeScalars converts the Arrow IPC bytes on the response into a
// table, under its own span: on wide responses this dominates the time
// spent between the RPC returning and the caller getting typed values.
func (r *GRPCOnlineQueryBulkResult) deserializeScalars(ctx context.Context) (arrow.Table, error) {
	allocator := r.allocator
	if allocator == nil {
		allocator = memory.DefaultAllocator
	}
	_, span := startSpan(ctx, r.tracer, spanDeserializeScalars)
	if span.IsRecording() {
		span.SetAttributes(attrScalarsBytes.Int(len(r.RawResponse.GetScalarsData())))
	}
	table, err := internal.ConvertBytesToTable(r.RawResponse.GetScalarsData(), allocator)
	if err == nil && table != nil && span.IsRecording() {
		span.SetAttributes(attrNumRows.Int64(table.NumRows()))
	}
	endSpan(span, err)
	return table, err
}

func (c *grpcClientImpl) getQueryClient(hasBranch bool) enginev1connect.QueryServiceClient {
	if hasBranch {
		return c.branchQueryClient
	}
	return c.queryClient
}

func (c *grpcClientImpl) UploadFeatures(
	ctx context.Context,
	params UploadFeaturesParams,
) (UploadFeaturesResult, error) {
	return c.featureWriteClient.UploadFeatures(ctx, params)
}

func (c *grpcClientImpl) DeleteFeatures(
	ctx context.Context,
	params DeleteFeaturesParams,
) (DeleteFeaturesResult, error) {
	return c.featureWriteClient.DeleteFeatures(ctx, params)
}

func (c *grpcClientImpl) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (*GRPCOnlineQueryBulkResult, error) {
	// callerCtx is held on to before the query span is started: the
	// response is deserialized lazily, after this function (and its span)
	// has returned, so those spans belong next to the query rather than
	// inside it.
	callerCtx := ctx

	// Describing the query costs a slice and a dozen field reads, so it is
	// only done when something is going to read it. Passing a nil slice
	// through the variadic does not allocate.
	var queryAttrs []attribute.KeyValue
	if c.tracer != nil {
		queryAttrs = c.queryAttributes(&args.underlying)
	}
	queryCtx, querySpan := startSpan(ctx, c.tracer, spanOnlineQueryBulk, queryAttrs...)
	result, err := c.onlineQueryBulk(queryCtx, querySpan, args)
	if result != nil {
		result.ctx = callerCtx
		result.tracer = c.tracer
	}
	endSpan(querySpan, err)
	return result, err
}

func (c *grpcClientImpl) onlineQueryBulk(
	ctx context.Context,
	querySpan trace.Span,
	args OnlineQueryParamsComplete,
) (*GRPCOnlineQueryBulkResult, error) {
	req, err := c.GetOnlineQueryBulkRequest(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, "generating online query request")
	}
	perRequestBranch := args.underlying.BranchId != nil && *args.underlying.BranchId != ""
	if perRequestBranch {
		req.Header().Set("x-chalk-branch-id", *args.underlying.BranchId)
	}
	hasBranch := perRequestBranch || c.branch != ""

	// The RPC gets its own span so that time on the wire is separable from
	// the serialization work around it, even when the caller has not
	// installed a connect interceptor that traces the call.
	rpcCtx, rpcSpan := startSpan(ctx, c.tracer, spanRPC)
	if rpcSpan.IsRecording() {
		rpcSpan.SetAttributes(attrInputsBytes.Int(len(req.Msg.GetInputsFeather())))
	}
	res, err := c.getQueryClient(hasBranch).OnlineQueryBulk(rpcCtx, req)
	if err != nil {
		endSpan(rpcSpan, err)
		return nil, errors.Wrap(err, "executing online query")
	}
	if rpcSpan.IsRecording() {
		responseAttrs := responseAttributes(res.Msg)
		rpcSpan.SetAttributes(responseAttrs...)
		querySpan.SetAttributes(responseAttrs...)
	}
	endSpan(rpcSpan, nil)

	result := &GRPCOnlineQueryBulkResult{RawResponse: res.Msg, allocator: c.allocator, ResponseHeaders: new(res.Header())}
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

// queryAttributes describes a query in a way that is safe to export: names,
// ids and counts, never feature values.
func (c *grpcClientImpl) queryAttributes(params *OnlineQueryParams) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, 8)
	if params.QueryName != "" {
		attrs = append(attrs, attrQueryName.String(params.QueryName))
	}
	if params.QueryNameVersion != "" {
		attrs = append(attrs, attrQueryNameVersion.String(params.QueryNameVersion))
	}
	if params.CorrelationId != "" {
		attrs = append(attrs, attrCorrelationID.String(params.CorrelationId))
	}
	if branch := ptr.OrZero(params.BranchId); branch != "" {
		attrs = append(attrs, attrBranch.String(branch))
	} else if c.branch != "" {
		attrs = append(attrs, attrBranch.String(c.branch))
	}
	if params.ResourceGroup != "" {
		attrs = append(attrs, attrResourceGroup.String(params.ResourceGroup))
	} else if c.resourceGroup != nil {
		attrs = append(attrs, attrResourceGroup.String(*c.resourceGroup))
	}
	if c.deploymentTag != "" {
		attrs = append(attrs, attrDeploymentTag.String(c.deploymentTag))
	}
	if envId := c.config.EnvironmentId.Value; envId != "" {
		attrs = append(attrs, attrEnvironment.String(envId))
	}
	return attrs
}

// responseAttributes reports the size of the payload the client has to
// deserialize, which is the main driver of the response-side spans.
func responseAttributes(res *commonv1.OnlineQueryBulkResponse) []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		attrScalarsBytes.Int(len(res.GetScalarsData())),
		attrNumErrors.Int(len(res.GetErrors())),
	}
	if groups := res.GetGroupsData(); len(groups) > 0 {
		var groupsBytes int
		for _, group := range groups {
			groupsBytes += len(group)
		}
		attrs = append(attrs, attrGroupsBytes.Int(groupsBytes))
	}
	if queryId := res.GetResponseMeta().GetQueryId(); queryId != "" {
		attrs = append(attrs, attrOperationID.String(queryId))
	}
	return attrs
}

func (c *grpcClientImpl) GetOnlineQueryBulkRequest(ctx context.Context, args OnlineQueryParamsComplete) (*connect.Request[commonv1.OnlineQueryBulkRequest], error) {
	ctx, span := startSpan(ctx, c.tracer, spanBuildRequest)
	paramsProto, err := convertOnlineQueryParamsToProto(ctx, c.tracer, &args.underlying, c.allocator)
	if err != nil {
		endSpan(span, err)
		return nil, errors.Wrap(err, "converting online query params to proto")
	}
	if span.IsRecording() {
		span.SetAttributes(
			attrNumOutputs.Int(len(paramsProto.GetOutputs())),
			attrInputsBytes.Int(len(paramsProto.GetInputsFeather())),
		)
	}
	endSpan(span, nil)

	req := connect.NewRequest(paramsProto)
	if args.underlying.ResourceGroup != "" {
		req.Header().Set(HeaderKeyResourceGroup, args.underlying.ResourceGroup)
	} else if c.resourceGroup != nil {
		req.Header().Set(HeaderKeyResourceGroup, *c.resourceGroup)
	}
	return req, nil
}

func (c *grpcClientImpl) GetMetadataServerInterceptor() []connect.ClientOption {
	return []connect.ClientOption{connect.WithInterceptors(c.metadataInterceptor)}
}

func (c *grpcClientImpl) GetEngineServerInterceptor() []connect.ClientOption {
	return []connect.ClientOption{connect.WithInterceptors(c.engineInterceptor)}
}

func (c *grpcClientImpl) GetConfig() *GRPCClientConfig {
	return &GRPCClientConfig{
		ClientId:       string(c.config.ClientId.Value),
		ClientSecret:   string(c.config.ClientSecret.Value),
		ApiServer:      c.config.GetAPIServer().Value,
		EnvironmentId:  c.config.EnvironmentId.Value,
		Branch:         c.branch,
		QueryServer:    c.config.GetGRPCQueryServer().Value,
		Logger:         c.logger,
		HTTPClient:     c.httpClient,
		DeploymentTag:  c.deploymentTag,
		ResourceGroup:  ptr.OrZero(c.resourceGroup),
		Timeout:        ptr.OrZero(c.timeout),
		Allocator:      c.allocator,
		TracerProvider: c.tracerProvider,
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

	res, err := c.getQueryClient(c.branch != "").UploadFeaturesBulk(ctx, req)
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
	res, err := c.getQueryClient(c.branch != "").GetAggregates(ctx, req)
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
	res, err := c.getQueryClient(c.branch != "").PlanAggregateBackfill(ctx, connect.NewRequest(req))
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
		PrimaryEnvironment: c.config.EnvironmentId.Value,
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

type GRPCListDatasetsResult struct {
	RawResponse *serverv1.ListDatasetsResponse
}

func (c *grpcClientImpl) ListDatasets(ctx context.Context, params ListDatasetsParams) (*GRPCListDatasetsResult, error) {
	req := connect.NewRequest(&serverv1.ListDatasetsRequest{
		Cursor: ptr.OrNil(params.Cursor),
		Limit:  ptr.OrNil(params.Limit),
		Search: ptr.OrNil(params.Search),
	})

	res, err := c.datasetMetadataClient.ListDatasets(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "listing datasets")
	}

	return &GRPCListDatasetsResult{RawResponse: res.Msg}, nil
}
