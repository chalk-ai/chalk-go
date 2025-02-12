// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/queries.proto

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
	// QueriesServiceName is the fully-qualified name of the QueriesService service.
	QueriesServiceName = "chalk.server.v1.QueriesService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueriesServiceGetQueryPerformanceSummaryProcedure is the fully-qualified name of the
	// QueriesService's GetQueryPerformanceSummary RPC.
	QueriesServiceGetQueryPerformanceSummaryProcedure = "/chalk.server.v1.QueriesService/GetQueryPerformanceSummary"
	// QueriesServiceListQueryErrorsProcedure is the fully-qualified name of the QueriesService's
	// ListQueryErrors RPC.
	QueriesServiceListQueryErrorsProcedure = "/chalk.server.v1.QueriesService/ListQueryErrors"
)

// QueriesServiceClient is a client for the chalk.server.v1.QueriesService service.
type QueriesServiceClient interface {
	GetQueryPerformanceSummary(context.Context, *connect.Request[v1.GetQueryPerformanceSummaryRequest]) (*connect.Response[v1.GetQueryPerformanceSummaryResponse], error)
	ListQueryErrors(context.Context, *connect.Request[v1.ListQueryErrorsRequest]) (*connect.Response[v1.ListQueryErrorsResponse], error)
}

// NewQueriesServiceClient constructs a client for the chalk.server.v1.QueriesService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueriesServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueriesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	queriesServiceMethods := v1.File_chalk_server_v1_queries_proto.Services().ByName("QueriesService").Methods()
	return &queriesServiceClient{
		getQueryPerformanceSummary: connect.NewClient[v1.GetQueryPerformanceSummaryRequest, v1.GetQueryPerformanceSummaryResponse](
			httpClient,
			baseURL+QueriesServiceGetQueryPerformanceSummaryProcedure,
			connect.WithSchema(queriesServiceMethods.ByName("GetQueryPerformanceSummary")),
			connect.WithClientOptions(opts...),
		),
		listQueryErrors: connect.NewClient[v1.ListQueryErrorsRequest, v1.ListQueryErrorsResponse](
			httpClient,
			baseURL+QueriesServiceListQueryErrorsProcedure,
			connect.WithSchema(queriesServiceMethods.ByName("ListQueryErrors")),
			connect.WithClientOptions(opts...),
		),
	}
}

// queriesServiceClient implements QueriesServiceClient.
type queriesServiceClient struct {
	getQueryPerformanceSummary *connect.Client[v1.GetQueryPerformanceSummaryRequest, v1.GetQueryPerformanceSummaryResponse]
	listQueryErrors            *connect.Client[v1.ListQueryErrorsRequest, v1.ListQueryErrorsResponse]
}

// GetQueryPerformanceSummary calls chalk.server.v1.QueriesService.GetQueryPerformanceSummary.
func (c *queriesServiceClient) GetQueryPerformanceSummary(ctx context.Context, req *connect.Request[v1.GetQueryPerformanceSummaryRequest]) (*connect.Response[v1.GetQueryPerformanceSummaryResponse], error) {
	return c.getQueryPerformanceSummary.CallUnary(ctx, req)
}

// ListQueryErrors calls chalk.server.v1.QueriesService.ListQueryErrors.
func (c *queriesServiceClient) ListQueryErrors(ctx context.Context, req *connect.Request[v1.ListQueryErrorsRequest]) (*connect.Response[v1.ListQueryErrorsResponse], error) {
	return c.listQueryErrors.CallUnary(ctx, req)
}

// QueriesServiceHandler is an implementation of the chalk.server.v1.QueriesService service.
type QueriesServiceHandler interface {
	GetQueryPerformanceSummary(context.Context, *connect.Request[v1.GetQueryPerformanceSummaryRequest]) (*connect.Response[v1.GetQueryPerformanceSummaryResponse], error)
	ListQueryErrors(context.Context, *connect.Request[v1.ListQueryErrorsRequest]) (*connect.Response[v1.ListQueryErrorsResponse], error)
}

// NewQueriesServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueriesServiceHandler(svc QueriesServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queriesServiceMethods := v1.File_chalk_server_v1_queries_proto.Services().ByName("QueriesService").Methods()
	queriesServiceGetQueryPerformanceSummaryHandler := connect.NewUnaryHandler(
		QueriesServiceGetQueryPerformanceSummaryProcedure,
		svc.GetQueryPerformanceSummary,
		connect.WithSchema(queriesServiceMethods.ByName("GetQueryPerformanceSummary")),
		connect.WithHandlerOptions(opts...),
	)
	queriesServiceListQueryErrorsHandler := connect.NewUnaryHandler(
		QueriesServiceListQueryErrorsProcedure,
		svc.ListQueryErrors,
		connect.WithSchema(queriesServiceMethods.ByName("ListQueryErrors")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.QueriesService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueriesServiceGetQueryPerformanceSummaryProcedure:
			queriesServiceGetQueryPerformanceSummaryHandler.ServeHTTP(w, r)
		case QueriesServiceListQueryErrorsProcedure:
			queriesServiceListQueryErrorsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueriesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueriesServiceHandler struct{}

func (UnimplementedQueriesServiceHandler) GetQueryPerformanceSummary(context.Context, *connect.Request[v1.GetQueryPerformanceSummaryRequest]) (*connect.Response[v1.GetQueryPerformanceSummaryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.QueriesService.GetQueryPerformanceSummary is not implemented"))
}

func (UnimplementedQueriesServiceHandler) ListQueryErrors(context.Context, *connect.Request[v1.ListQueryErrorsRequest]) (*connect.Response[v1.ListQueryErrorsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.QueriesService.ListQueryErrors is not implemented"))
}
