// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/billing.proto

package serverv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BillingServiceName is the fully-qualified name of the BillingService service.
	BillingServiceName = "chalk.server.v1.BillingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BillingServiceGetNodesAndPodsUIProcedure is the fully-qualified name of the BillingService's
	// GetNodesAndPodsUI RPC.
	BillingServiceGetNodesAndPodsUIProcedure = "/chalk.server.v1.BillingService/GetNodesAndPodsUI"
	// BillingServiceGetNodesAndPodsProcedure is the fully-qualified name of the BillingService's
	// GetNodesAndPods RPC.
	BillingServiceGetNodesAndPodsProcedure = "/chalk.server.v1.BillingService/GetNodesAndPods"
	// BillingServiceGetUsageChartProcedure is the fully-qualified name of the BillingService's
	// GetUsageChart RPC.
	BillingServiceGetUsageChartProcedure = "/chalk.server.v1.BillingService/GetUsageChart"
	// BillingServiceGetUtilizationRatesProcedure is the fully-qualified name of the BillingService's
	// GetUtilizationRates RPC.
	BillingServiceGetUtilizationRatesProcedure = "/chalk.server.v1.BillingService/GetUtilizationRates"
	// BillingServiceGetPodRequestChartsProcedure is the fully-qualified name of the BillingService's
	// GetPodRequestCharts RPC.
	BillingServiceGetPodRequestChartsProcedure = "/chalk.server.v1.BillingService/GetPodRequestCharts"
)

// BillingServiceClient is a client for the chalk.server.v1.BillingService service.
type BillingServiceClient interface {
	// GetNodesAndPodsUI returns the nodes and pods for the team by default,
	// not just a single environment. To limit the scope, add filters to
	// the request object.
	// Use this endpoint going forward; GetNodesAndPods should be deprecated because
	// it reuses PubSub types that are dangerous to update and are not intended for UI use.
	GetNodesAndPodsUI(context.Context, *connect.Request[v1.GetNodesAndPodsUIRequest]) (*connect.Response[v1.GetNodesAndPodsUIResponse], error)
	// GetNodesAndPods returns the nodes and pods for the team by default,
	// not just a single environment. To limit the scope, add filters to
	// the request object.
	GetNodesAndPods(context.Context, *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error)
	// GetUsageChart shows the Chalk credit usage between a provided start and
	// end period. The usage can be grouped by UsageChartPeriod for daily or
	// monthly usage, and by UsageChartGrouping for instance type or cluster usage.
	GetUsageChart(context.Context, *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error)
	// GetUtilizationRates returns the current utilization rates for all
	// instance types.
	GetUtilizationRates(context.Context, *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error)
	GetPodRequestCharts(context.Context, *connect.Request[v1.GetPodRequestChartsRequest]) (*connect.Response[v1.GetPodRequestChartsResponse], error)
}

// NewBillingServiceClient constructs a client for the chalk.server.v1.BillingService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBillingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BillingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	billingServiceMethods := v1.File_chalk_server_v1_billing_proto.Services().ByName("BillingService").Methods()
	return &billingServiceClient{
		getNodesAndPodsUI: connect.NewClient[v1.GetNodesAndPodsUIRequest, v1.GetNodesAndPodsUIResponse](
			httpClient,
			baseURL+BillingServiceGetNodesAndPodsUIProcedure,
			connect.WithSchema(billingServiceMethods.ByName("GetNodesAndPodsUI")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getNodesAndPods: connect.NewClient[v1.GetNodesAndPodsRequest, v1.GetNodesAndPodsResponse](
			httpClient,
			baseURL+BillingServiceGetNodesAndPodsProcedure,
			connect.WithSchema(billingServiceMethods.ByName("GetNodesAndPods")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getUsageChart: connect.NewClient[v1.GetUsageChartRequest, v1.GetUsageChartResponse](
			httpClient,
			baseURL+BillingServiceGetUsageChartProcedure,
			connect.WithSchema(billingServiceMethods.ByName("GetUsageChart")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getUtilizationRates: connect.NewClient[v1.GetUtilizationRatesRequest, v1.GetUtilizationRatesResponse](
			httpClient,
			baseURL+BillingServiceGetUtilizationRatesProcedure,
			connect.WithSchema(billingServiceMethods.ByName("GetUtilizationRates")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getPodRequestCharts: connect.NewClient[v1.GetPodRequestChartsRequest, v1.GetPodRequestChartsResponse](
			httpClient,
			baseURL+BillingServiceGetPodRequestChartsProcedure,
			connect.WithSchema(billingServiceMethods.ByName("GetPodRequestCharts")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// billingServiceClient implements BillingServiceClient.
type billingServiceClient struct {
	getNodesAndPodsUI   *connect.Client[v1.GetNodesAndPodsUIRequest, v1.GetNodesAndPodsUIResponse]
	getNodesAndPods     *connect.Client[v1.GetNodesAndPodsRequest, v1.GetNodesAndPodsResponse]
	getUsageChart       *connect.Client[v1.GetUsageChartRequest, v1.GetUsageChartResponse]
	getUtilizationRates *connect.Client[v1.GetUtilizationRatesRequest, v1.GetUtilizationRatesResponse]
	getPodRequestCharts *connect.Client[v1.GetPodRequestChartsRequest, v1.GetPodRequestChartsResponse]
}

// GetNodesAndPodsUI calls chalk.server.v1.BillingService.GetNodesAndPodsUI.
func (c *billingServiceClient) GetNodesAndPodsUI(ctx context.Context, req *connect.Request[v1.GetNodesAndPodsUIRequest]) (*connect.Response[v1.GetNodesAndPodsUIResponse], error) {
	return c.getNodesAndPodsUI.CallUnary(ctx, req)
}

// GetNodesAndPods calls chalk.server.v1.BillingService.GetNodesAndPods.
func (c *billingServiceClient) GetNodesAndPods(ctx context.Context, req *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error) {
	return c.getNodesAndPods.CallUnary(ctx, req)
}

// GetUsageChart calls chalk.server.v1.BillingService.GetUsageChart.
func (c *billingServiceClient) GetUsageChart(ctx context.Context, req *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error) {
	return c.getUsageChart.CallUnary(ctx, req)
}

// GetUtilizationRates calls chalk.server.v1.BillingService.GetUtilizationRates.
func (c *billingServiceClient) GetUtilizationRates(ctx context.Context, req *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error) {
	return c.getUtilizationRates.CallUnary(ctx, req)
}

// GetPodRequestCharts calls chalk.server.v1.BillingService.GetPodRequestCharts.
func (c *billingServiceClient) GetPodRequestCharts(ctx context.Context, req *connect.Request[v1.GetPodRequestChartsRequest]) (*connect.Response[v1.GetPodRequestChartsResponse], error) {
	return c.getPodRequestCharts.CallUnary(ctx, req)
}

// BillingServiceHandler is an implementation of the chalk.server.v1.BillingService service.
type BillingServiceHandler interface {
	// GetNodesAndPodsUI returns the nodes and pods for the team by default,
	// not just a single environment. To limit the scope, add filters to
	// the request object.
	// Use this endpoint going forward; GetNodesAndPods should be deprecated because
	// it reuses PubSub types that are dangerous to update and are not intended for UI use.
	GetNodesAndPodsUI(context.Context, *connect.Request[v1.GetNodesAndPodsUIRequest]) (*connect.Response[v1.GetNodesAndPodsUIResponse], error)
	// GetNodesAndPods returns the nodes and pods for the team by default,
	// not just a single environment. To limit the scope, add filters to
	// the request object.
	GetNodesAndPods(context.Context, *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error)
	// GetUsageChart shows the Chalk credit usage between a provided start and
	// end period. The usage can be grouped by UsageChartPeriod for daily or
	// monthly usage, and by UsageChartGrouping for instance type or cluster usage.
	GetUsageChart(context.Context, *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error)
	// GetUtilizationRates returns the current utilization rates for all
	// instance types.
	GetUtilizationRates(context.Context, *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error)
	GetPodRequestCharts(context.Context, *connect.Request[v1.GetPodRequestChartsRequest]) (*connect.Response[v1.GetPodRequestChartsResponse], error)
}

// NewBillingServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBillingServiceHandler(svc BillingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	billingServiceMethods := v1.File_chalk_server_v1_billing_proto.Services().ByName("BillingService").Methods()
	billingServiceGetNodesAndPodsUIHandler := connect.NewUnaryHandler(
		BillingServiceGetNodesAndPodsUIProcedure,
		svc.GetNodesAndPodsUI,
		connect.WithSchema(billingServiceMethods.ByName("GetNodesAndPodsUI")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetNodesAndPodsHandler := connect.NewUnaryHandler(
		BillingServiceGetNodesAndPodsProcedure,
		svc.GetNodesAndPods,
		connect.WithSchema(billingServiceMethods.ByName("GetNodesAndPods")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetUsageChartHandler := connect.NewUnaryHandler(
		BillingServiceGetUsageChartProcedure,
		svc.GetUsageChart,
		connect.WithSchema(billingServiceMethods.ByName("GetUsageChart")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetUtilizationRatesHandler := connect.NewUnaryHandler(
		BillingServiceGetUtilizationRatesProcedure,
		svc.GetUtilizationRates,
		connect.WithSchema(billingServiceMethods.ByName("GetUtilizationRates")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetPodRequestChartsHandler := connect.NewUnaryHandler(
		BillingServiceGetPodRequestChartsProcedure,
		svc.GetPodRequestCharts,
		connect.WithSchema(billingServiceMethods.ByName("GetPodRequestCharts")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.BillingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BillingServiceGetNodesAndPodsUIProcedure:
			billingServiceGetNodesAndPodsUIHandler.ServeHTTP(w, r)
		case BillingServiceGetNodesAndPodsProcedure:
			billingServiceGetNodesAndPodsHandler.ServeHTTP(w, r)
		case BillingServiceGetUsageChartProcedure:
			billingServiceGetUsageChartHandler.ServeHTTP(w, r)
		case BillingServiceGetUtilizationRatesProcedure:
			billingServiceGetUtilizationRatesHandler.ServeHTTP(w, r)
		case BillingServiceGetPodRequestChartsProcedure:
			billingServiceGetPodRequestChartsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBillingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBillingServiceHandler struct{}

func (UnimplementedBillingServiceHandler) GetNodesAndPodsUI(context.Context, *connect.Request[v1.GetNodesAndPodsUIRequest]) (*connect.Response[v1.GetNodesAndPodsUIResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetNodesAndPodsUI is not implemented"))
}

func (UnimplementedBillingServiceHandler) GetNodesAndPods(context.Context, *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetNodesAndPods is not implemented"))
}

func (UnimplementedBillingServiceHandler) GetUsageChart(context.Context, *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetUsageChart is not implemented"))
}

func (UnimplementedBillingServiceHandler) GetUtilizationRates(context.Context, *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetUtilizationRates is not implemented"))
}

func (UnimplementedBillingServiceHandler) GetPodRequestCharts(context.Context, *connect.Request[v1.GetPodRequestChartsRequest]) (*connect.Response[v1.GetPodRequestChartsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetPodRequestCharts is not implemented"))
}
