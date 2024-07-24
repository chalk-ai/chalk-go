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
	// BillingServiceGetNodesProcedure is the fully-qualified name of the BillingService's GetNodes RPC.
	BillingServiceGetNodesProcedure = "/chalk.server.v1.BillingService/GetNodes"
	// BillingServiceGetNodesAndPodsProcedure is the fully-qualified name of the BillingService's
	// GetNodesAndPods RPC.
	BillingServiceGetNodesAndPodsProcedure = "/chalk.server.v1.BillingService/GetNodesAndPods"
	// BillingServiceGetUsageChartProcedure is the fully-qualified name of the BillingService's
	// GetUsageChart RPC.
	BillingServiceGetUsageChartProcedure = "/chalk.server.v1.BillingService/GetUsageChart"
	// BillingServiceGetUtilizationRatesProcedure is the fully-qualified name of the BillingService's
	// GetUtilizationRates RPC.
	BillingServiceGetUtilizationRatesProcedure = "/chalk.server.v1.BillingService/GetUtilizationRates"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	billingServiceServiceDescriptor                   = v1.File_chalk_server_v1_billing_proto.Services().ByName("BillingService")
	billingServiceGetNodesMethodDescriptor            = billingServiceServiceDescriptor.Methods().ByName("GetNodes")
	billingServiceGetNodesAndPodsMethodDescriptor     = billingServiceServiceDescriptor.Methods().ByName("GetNodesAndPods")
	billingServiceGetUsageChartMethodDescriptor       = billingServiceServiceDescriptor.Methods().ByName("GetUsageChart")
	billingServiceGetUtilizationRatesMethodDescriptor = billingServiceServiceDescriptor.Methods().ByName("GetUtilizationRates")
)

// BillingServiceClient is a client for the chalk.server.v1.BillingService service.
type BillingServiceClient interface {
	GetNodes(context.Context, *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error)
	// Gets the nodes and pods for the team.
	GetNodesAndPods(context.Context, *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error)
	GetUsageChart(context.Context, *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error)
	GetUtilizationRates(context.Context, *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error)
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
	return &billingServiceClient{
		getNodes: connect.NewClient[v1.GetNodesRequest, v1.GetNodesResponse](
			httpClient,
			baseURL+BillingServiceGetNodesProcedure,
			connect.WithSchema(billingServiceGetNodesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getNodesAndPods: connect.NewClient[v1.GetNodesAndPodsRequest, v1.GetNodesAndPodsResponse](
			httpClient,
			baseURL+BillingServiceGetNodesAndPodsProcedure,
			connect.WithSchema(billingServiceGetNodesAndPodsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getUsageChart: connect.NewClient[v1.GetUsageChartRequest, v1.GetUsageChartResponse](
			httpClient,
			baseURL+BillingServiceGetUsageChartProcedure,
			connect.WithSchema(billingServiceGetUsageChartMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getUtilizationRates: connect.NewClient[v1.GetUtilizationRatesRequest, v1.GetUtilizationRatesResponse](
			httpClient,
			baseURL+BillingServiceGetUtilizationRatesProcedure,
			connect.WithSchema(billingServiceGetUtilizationRatesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// billingServiceClient implements BillingServiceClient.
type billingServiceClient struct {
	getNodes            *connect.Client[v1.GetNodesRequest, v1.GetNodesResponse]
	getNodesAndPods     *connect.Client[v1.GetNodesAndPodsRequest, v1.GetNodesAndPodsResponse]
	getUsageChart       *connect.Client[v1.GetUsageChartRequest, v1.GetUsageChartResponse]
	getUtilizationRates *connect.Client[v1.GetUtilizationRatesRequest, v1.GetUtilizationRatesResponse]
}

// GetNodes calls chalk.server.v1.BillingService.GetNodes.
func (c *billingServiceClient) GetNodes(ctx context.Context, req *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error) {
	return c.getNodes.CallUnary(ctx, req)
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

// BillingServiceHandler is an implementation of the chalk.server.v1.BillingService service.
type BillingServiceHandler interface {
	GetNodes(context.Context, *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error)
	// Gets the nodes and pods for the team.
	GetNodesAndPods(context.Context, *connect.Request[v1.GetNodesAndPodsRequest]) (*connect.Response[v1.GetNodesAndPodsResponse], error)
	GetUsageChart(context.Context, *connect.Request[v1.GetUsageChartRequest]) (*connect.Response[v1.GetUsageChartResponse], error)
	GetUtilizationRates(context.Context, *connect.Request[v1.GetUtilizationRatesRequest]) (*connect.Response[v1.GetUtilizationRatesResponse], error)
}

// NewBillingServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBillingServiceHandler(svc BillingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	billingServiceGetNodesHandler := connect.NewUnaryHandler(
		BillingServiceGetNodesProcedure,
		svc.GetNodes,
		connect.WithSchema(billingServiceGetNodesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetNodesAndPodsHandler := connect.NewUnaryHandler(
		BillingServiceGetNodesAndPodsProcedure,
		svc.GetNodesAndPods,
		connect.WithSchema(billingServiceGetNodesAndPodsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetUsageChartHandler := connect.NewUnaryHandler(
		BillingServiceGetUsageChartProcedure,
		svc.GetUsageChart,
		connect.WithSchema(billingServiceGetUsageChartMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	billingServiceGetUtilizationRatesHandler := connect.NewUnaryHandler(
		BillingServiceGetUtilizationRatesProcedure,
		svc.GetUtilizationRates,
		connect.WithSchema(billingServiceGetUtilizationRatesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.BillingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BillingServiceGetNodesProcedure:
			billingServiceGetNodesHandler.ServeHTTP(w, r)
		case BillingServiceGetNodesAndPodsProcedure:
			billingServiceGetNodesAndPodsHandler.ServeHTTP(w, r)
		case BillingServiceGetUsageChartProcedure:
			billingServiceGetUsageChartHandler.ServeHTTP(w, r)
		case BillingServiceGetUtilizationRatesProcedure:
			billingServiceGetUtilizationRatesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBillingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBillingServiceHandler struct{}

func (UnimplementedBillingServiceHandler) GetNodes(context.Context, *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BillingService.GetNodes is not implemented"))
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
