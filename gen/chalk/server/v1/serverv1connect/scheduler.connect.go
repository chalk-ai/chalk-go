// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/scheduler.proto

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
	// SchedulerServiceName is the fully-qualified name of the SchedulerService service.
	SchedulerServiceName = "chalk.server.v1.SchedulerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SchedulerServiceManualTriggerScheduledQueryProcedure is the fully-qualified name of the
	// SchedulerService's ManualTriggerScheduledQuery RPC.
	SchedulerServiceManualTriggerScheduledQueryProcedure = "/chalk.server.v1.SchedulerService/ManualTriggerScheduledQuery"
)

// SchedulerServiceClient is a client for the chalk.server.v1.SchedulerService service.
type SchedulerServiceClient interface {
	ManualTriggerScheduledQuery(context.Context, *connect.Request[v1.ManualTriggerScheduledQueryRequest]) (*connect.Response[v1.ManualTriggerScheduledQueryResponse], error)
}

// NewSchedulerServiceClient constructs a client for the chalk.server.v1.SchedulerService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchedulerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SchedulerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	schedulerServiceMethods := v1.File_chalk_server_v1_scheduler_proto.Services().ByName("SchedulerService").Methods()
	return &schedulerServiceClient{
		manualTriggerScheduledQuery: connect.NewClient[v1.ManualTriggerScheduledQueryRequest, v1.ManualTriggerScheduledQueryResponse](
			httpClient,
			baseURL+SchedulerServiceManualTriggerScheduledQueryProcedure,
			connect.WithSchema(schedulerServiceMethods.ByName("ManualTriggerScheduledQuery")),
			connect.WithClientOptions(opts...),
		),
	}
}

// schedulerServiceClient implements SchedulerServiceClient.
type schedulerServiceClient struct {
	manualTriggerScheduledQuery *connect.Client[v1.ManualTriggerScheduledQueryRequest, v1.ManualTriggerScheduledQueryResponse]
}

// ManualTriggerScheduledQuery calls chalk.server.v1.SchedulerService.ManualTriggerScheduledQuery.
func (c *schedulerServiceClient) ManualTriggerScheduledQuery(ctx context.Context, req *connect.Request[v1.ManualTriggerScheduledQueryRequest]) (*connect.Response[v1.ManualTriggerScheduledQueryResponse], error) {
	return c.manualTriggerScheduledQuery.CallUnary(ctx, req)
}

// SchedulerServiceHandler is an implementation of the chalk.server.v1.SchedulerService service.
type SchedulerServiceHandler interface {
	ManualTriggerScheduledQuery(context.Context, *connect.Request[v1.ManualTriggerScheduledQueryRequest]) (*connect.Response[v1.ManualTriggerScheduledQueryResponse], error)
}

// NewSchedulerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchedulerServiceHandler(svc SchedulerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	schedulerServiceMethods := v1.File_chalk_server_v1_scheduler_proto.Services().ByName("SchedulerService").Methods()
	schedulerServiceManualTriggerScheduledQueryHandler := connect.NewUnaryHandler(
		SchedulerServiceManualTriggerScheduledQueryProcedure,
		svc.ManualTriggerScheduledQuery,
		connect.WithSchema(schedulerServiceMethods.ByName("ManualTriggerScheduledQuery")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.SchedulerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SchedulerServiceManualTriggerScheduledQueryProcedure:
			schedulerServiceManualTriggerScheduledQueryHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSchedulerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSchedulerServiceHandler struct{}

func (UnimplementedSchedulerServiceHandler) ManualTriggerScheduledQuery(context.Context, *connect.Request[v1.ManualTriggerScheduledQueryRequest]) (*connect.Response[v1.ManualTriggerScheduledQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.SchedulerService.ManualTriggerScheduledQuery is not implemented"))
}
