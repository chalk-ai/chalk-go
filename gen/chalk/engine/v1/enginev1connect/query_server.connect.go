// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/engine/v1/query_server.proto

package enginev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/engine/v1"
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
	// QueryServiceName is the fully-qualified name of the QueryService service.
	QueryServiceName = "chalk.engine.v1.QueryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueryServicePingProcedure is the fully-qualified name of the QueryService's Ping RPC.
	QueryServicePingProcedure = "/chalk.engine.v1.QueryService/Ping"
	// QueryServiceOnlineQueryProcedure is the fully-qualified name of the QueryService's OnlineQuery
	// RPC.
	QueryServiceOnlineQueryProcedure = "/chalk.engine.v1.QueryService/OnlineQuery"
	// QueryServiceOnlineQueryBulkProcedure is the fully-qualified name of the QueryService's
	// OnlineQueryBulk RPC.
	QueryServiceOnlineQueryBulkProcedure = "/chalk.engine.v1.QueryService/OnlineQueryBulk"
	// QueryServiceOnlineQueryMultiProcedure is the fully-qualified name of the QueryService's
	// OnlineQueryMulti RPC.
	QueryServiceOnlineQueryMultiProcedure = "/chalk.engine.v1.QueryService/OnlineQueryMulti"
	// QueryServiceQueryFromPlanProcedure is the fully-qualified name of the QueryService's
	// QueryFromPlan RPC.
	QueryServiceQueryFromPlanProcedure = "/chalk.engine.v1.QueryService/QueryFromPlan"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	queryServiceServiceDescriptor                = v1.File_chalk_engine_v1_query_server_proto.Services().ByName("QueryService")
	queryServicePingMethodDescriptor             = queryServiceServiceDescriptor.Methods().ByName("Ping")
	queryServiceOnlineQueryMethodDescriptor      = queryServiceServiceDescriptor.Methods().ByName("OnlineQuery")
	queryServiceOnlineQueryBulkMethodDescriptor  = queryServiceServiceDescriptor.Methods().ByName("OnlineQueryBulk")
	queryServiceOnlineQueryMultiMethodDescriptor = queryServiceServiceDescriptor.Methods().ByName("OnlineQueryMulti")
	queryServiceQueryFromPlanMethodDescriptor    = queryServiceServiceDescriptor.Methods().ByName("QueryFromPlan")
)

// QueryServiceClient is a client for the chalk.engine.v1.QueryService service.
type QueryServiceClient interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	OnlineQuery(context.Context, *connect.Request[v11.OnlineQueryRequest]) (*connect.Response[v11.OnlineQueryResponse], error)
	OnlineQueryBulk(context.Context, *connect.Request[v11.OnlineQueryBulkRequest]) (*connect.Response[v11.OnlineQueryBulkResponse], error)
	OnlineQueryMulti(context.Context, *connect.Request[v11.OnlineQueryMultiRequest]) (*connect.Response[v11.OnlineQueryMultiResponse], error)
	QueryFromPlan(context.Context, *connect.Request[v1.QueryFromPlanRequest]) (*connect.Response[v1.QueryFromPlanResponse], error)
}

// NewQueryServiceClient constructs a client for the chalk.engine.v1.QueryService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &queryServiceClient{
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+QueryServicePingProcedure,
			connect.WithSchema(queryServicePingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onlineQuery: connect.NewClient[v11.OnlineQueryRequest, v11.OnlineQueryResponse](
			httpClient,
			baseURL+QueryServiceOnlineQueryProcedure,
			connect.WithSchema(queryServiceOnlineQueryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onlineQueryBulk: connect.NewClient[v11.OnlineQueryBulkRequest, v11.OnlineQueryBulkResponse](
			httpClient,
			baseURL+QueryServiceOnlineQueryBulkProcedure,
			connect.WithSchema(queryServiceOnlineQueryBulkMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onlineQueryMulti: connect.NewClient[v11.OnlineQueryMultiRequest, v11.OnlineQueryMultiResponse](
			httpClient,
			baseURL+QueryServiceOnlineQueryMultiProcedure,
			connect.WithSchema(queryServiceOnlineQueryMultiMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryFromPlan: connect.NewClient[v1.QueryFromPlanRequest, v1.QueryFromPlanResponse](
			httpClient,
			baseURL+QueryServiceQueryFromPlanProcedure,
			connect.WithSchema(queryServiceQueryFromPlanMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	ping             *connect.Client[v1.PingRequest, v1.PingResponse]
	onlineQuery      *connect.Client[v11.OnlineQueryRequest, v11.OnlineQueryResponse]
	onlineQueryBulk  *connect.Client[v11.OnlineQueryBulkRequest, v11.OnlineQueryBulkResponse]
	onlineQueryMulti *connect.Client[v11.OnlineQueryMultiRequest, v11.OnlineQueryMultiResponse]
	queryFromPlan    *connect.Client[v1.QueryFromPlanRequest, v1.QueryFromPlanResponse]
}

// Ping calls chalk.engine.v1.QueryService.Ping.
func (c *queryServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// OnlineQuery calls chalk.engine.v1.QueryService.OnlineQuery.
func (c *queryServiceClient) OnlineQuery(ctx context.Context, req *connect.Request[v11.OnlineQueryRequest]) (*connect.Response[v11.OnlineQueryResponse], error) {
	return c.onlineQuery.CallUnary(ctx, req)
}

// OnlineQueryBulk calls chalk.engine.v1.QueryService.OnlineQueryBulk.
func (c *queryServiceClient) OnlineQueryBulk(ctx context.Context, req *connect.Request[v11.OnlineQueryBulkRequest]) (*connect.Response[v11.OnlineQueryBulkResponse], error) {
	return c.onlineQueryBulk.CallUnary(ctx, req)
}

// OnlineQueryMulti calls chalk.engine.v1.QueryService.OnlineQueryMulti.
func (c *queryServiceClient) OnlineQueryMulti(ctx context.Context, req *connect.Request[v11.OnlineQueryMultiRequest]) (*connect.Response[v11.OnlineQueryMultiResponse], error) {
	return c.onlineQueryMulti.CallUnary(ctx, req)
}

// QueryFromPlan calls chalk.engine.v1.QueryService.QueryFromPlan.
func (c *queryServiceClient) QueryFromPlan(ctx context.Context, req *connect.Request[v1.QueryFromPlanRequest]) (*connect.Response[v1.QueryFromPlanResponse], error) {
	return c.queryFromPlan.CallUnary(ctx, req)
}

// QueryServiceHandler is an implementation of the chalk.engine.v1.QueryService service.
type QueryServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	OnlineQuery(context.Context, *connect.Request[v11.OnlineQueryRequest]) (*connect.Response[v11.OnlineQueryResponse], error)
	OnlineQueryBulk(context.Context, *connect.Request[v11.OnlineQueryBulkRequest]) (*connect.Response[v11.OnlineQueryBulkResponse], error)
	OnlineQueryMulti(context.Context, *connect.Request[v11.OnlineQueryMultiRequest]) (*connect.Response[v11.OnlineQueryMultiResponse], error)
	QueryFromPlan(context.Context, *connect.Request[v1.QueryFromPlanRequest]) (*connect.Response[v1.QueryFromPlanResponse], error)
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServicePingHandler := connect.NewUnaryHandler(
		QueryServicePingProcedure,
		svc.Ping,
		connect.WithSchema(queryServicePingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceOnlineQueryHandler := connect.NewUnaryHandler(
		QueryServiceOnlineQueryProcedure,
		svc.OnlineQuery,
		connect.WithSchema(queryServiceOnlineQueryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceOnlineQueryBulkHandler := connect.NewUnaryHandler(
		QueryServiceOnlineQueryBulkProcedure,
		svc.OnlineQueryBulk,
		connect.WithSchema(queryServiceOnlineQueryBulkMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceOnlineQueryMultiHandler := connect.NewUnaryHandler(
		QueryServiceOnlineQueryMultiProcedure,
		svc.OnlineQueryMulti,
		connect.WithSchema(queryServiceOnlineQueryMultiMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceQueryFromPlanHandler := connect.NewUnaryHandler(
		QueryServiceQueryFromPlanProcedure,
		svc.QueryFromPlan,
		connect.WithSchema(queryServiceQueryFromPlanMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.engine.v1.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServicePingProcedure:
			queryServicePingHandler.ServeHTTP(w, r)
		case QueryServiceOnlineQueryProcedure:
			queryServiceOnlineQueryHandler.ServeHTTP(w, r)
		case QueryServiceOnlineQueryBulkProcedure:
			queryServiceOnlineQueryBulkHandler.ServeHTTP(w, r)
		case QueryServiceOnlineQueryMultiProcedure:
			queryServiceOnlineQueryMultiHandler.ServeHTTP(w, r)
		case QueryServiceQueryFromPlanProcedure:
			queryServiceQueryFromPlanHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueryServiceHandler struct{}

func (UnimplementedQueryServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.QueryService.Ping is not implemented"))
}

func (UnimplementedQueryServiceHandler) OnlineQuery(context.Context, *connect.Request[v11.OnlineQueryRequest]) (*connect.Response[v11.OnlineQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.QueryService.OnlineQuery is not implemented"))
}

func (UnimplementedQueryServiceHandler) OnlineQueryBulk(context.Context, *connect.Request[v11.OnlineQueryBulkRequest]) (*connect.Response[v11.OnlineQueryBulkResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.QueryService.OnlineQueryBulk is not implemented"))
}

func (UnimplementedQueryServiceHandler) OnlineQueryMulti(context.Context, *connect.Request[v11.OnlineQueryMultiRequest]) (*connect.Response[v11.OnlineQueryMultiResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.QueryService.OnlineQueryMulti is not implemented"))
}

func (UnimplementedQueryServiceHandler) QueryFromPlan(context.Context, *connect.Request[v1.QueryFromPlanRequest]) (*connect.Response[v1.QueryFromPlanResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.QueryService.QueryFromPlan is not implemented"))
}
