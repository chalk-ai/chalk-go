// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/status.proto

package serverv1connect

import (
	v1 "chalk/chalk/server/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
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
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "chalk.server.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServiceCheckHealthProcedure is the fully-qualified name of the HealthService's CheckHealth
	// RPC.
	HealthServiceCheckHealthProcedure = "/chalk.server.v1.HealthService/CheckHealth"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	healthServiceServiceDescriptor           = v1.File_chalk_server_v1_status_proto.Services().ByName("HealthService")
	healthServiceCheckHealthMethodDescriptor = healthServiceServiceDescriptor.Methods().ByName("CheckHealth")
)

// HealthServiceClient is a client for the chalk.server.v1.HealthService service.
type HealthServiceClient interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
}

// NewHealthServiceClient constructs a client for the chalk.server.v1.HealthService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthServiceClient{
		checkHealth: connect.NewClient[v1.CheckHealthRequest, v1.CheckHealthResponse](
			httpClient,
			baseURL+HealthServiceCheckHealthProcedure,
			connect.WithSchema(healthServiceCheckHealthMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	checkHealth *connect.Client[v1.CheckHealthRequest, v1.CheckHealthResponse]
}

// CheckHealth calls chalk.server.v1.HealthService.CheckHealth.
func (c *healthServiceClient) CheckHealth(ctx context.Context, req *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return c.checkHealth.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the chalk.server.v1.HealthService service.
type HealthServiceHandler interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	healthServiceCheckHealthHandler := connect.NewUnaryHandler(
		HealthServiceCheckHealthProcedure,
		svc.CheckHealth,
		connect.WithSchema(healthServiceCheckHealthMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServiceCheckHealthProcedure:
			healthServiceCheckHealthHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.HealthService.CheckHealth is not implemented"))
}