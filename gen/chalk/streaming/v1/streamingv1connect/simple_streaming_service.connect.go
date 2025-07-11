// We expect to remove this file.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/streaming/v1/simple_streaming_service.proto

package streamingv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/streaming/v1"
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
	// SimpleStreamingServiceName is the fully-qualified name of the SimpleStreamingService service.
	SimpleStreamingServiceName = "chalk.streaming.v1.SimpleStreamingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SimpleStreamingServiceSimpleStreamingUnaryInvokeProcedure is the fully-qualified name of the
	// SimpleStreamingService's SimpleStreamingUnaryInvoke RPC.
	SimpleStreamingServiceSimpleStreamingUnaryInvokeProcedure = "/chalk.streaming.v1.SimpleStreamingService/SimpleStreamingUnaryInvoke"
)

// SimpleStreamingServiceClient is a client for the chalk.streaming.v1.SimpleStreamingService
// service.
type SimpleStreamingServiceClient interface {
	// Runs a simple streaming plan with the given request.
	// This is a simplified version of the streaming invoker service.
	SimpleStreamingUnaryInvoke(context.Context, *connect.Request[v1.SimpleStreamingUnaryInvokeRequest]) (*connect.Response[v1.SimpleStreamingUnaryInvokeResponse], error)
}

// NewSimpleStreamingServiceClient constructs a client for the
// chalk.streaming.v1.SimpleStreamingService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSimpleStreamingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SimpleStreamingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	simpleStreamingServiceMethods := v1.File_chalk_streaming_v1_simple_streaming_service_proto.Services().ByName("SimpleStreamingService").Methods()
	return &simpleStreamingServiceClient{
		simpleStreamingUnaryInvoke: connect.NewClient[v1.SimpleStreamingUnaryInvokeRequest, v1.SimpleStreamingUnaryInvokeResponse](
			httpClient,
			baseURL+SimpleStreamingServiceSimpleStreamingUnaryInvokeProcedure,
			connect.WithSchema(simpleStreamingServiceMethods.ByName("SimpleStreamingUnaryInvoke")),
			connect.WithClientOptions(opts...),
		),
	}
}

// simpleStreamingServiceClient implements SimpleStreamingServiceClient.
type simpleStreamingServiceClient struct {
	simpleStreamingUnaryInvoke *connect.Client[v1.SimpleStreamingUnaryInvokeRequest, v1.SimpleStreamingUnaryInvokeResponse]
}

// SimpleStreamingUnaryInvoke calls
// chalk.streaming.v1.SimpleStreamingService.SimpleStreamingUnaryInvoke.
func (c *simpleStreamingServiceClient) SimpleStreamingUnaryInvoke(ctx context.Context, req *connect.Request[v1.SimpleStreamingUnaryInvokeRequest]) (*connect.Response[v1.SimpleStreamingUnaryInvokeResponse], error) {
	return c.simpleStreamingUnaryInvoke.CallUnary(ctx, req)
}

// SimpleStreamingServiceHandler is an implementation of the
// chalk.streaming.v1.SimpleStreamingService service.
type SimpleStreamingServiceHandler interface {
	// Runs a simple streaming plan with the given request.
	// This is a simplified version of the streaming invoker service.
	SimpleStreamingUnaryInvoke(context.Context, *connect.Request[v1.SimpleStreamingUnaryInvokeRequest]) (*connect.Response[v1.SimpleStreamingUnaryInvokeResponse], error)
}

// NewSimpleStreamingServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSimpleStreamingServiceHandler(svc SimpleStreamingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	simpleStreamingServiceMethods := v1.File_chalk_streaming_v1_simple_streaming_service_proto.Services().ByName("SimpleStreamingService").Methods()
	simpleStreamingServiceSimpleStreamingUnaryInvokeHandler := connect.NewUnaryHandler(
		SimpleStreamingServiceSimpleStreamingUnaryInvokeProcedure,
		svc.SimpleStreamingUnaryInvoke,
		connect.WithSchema(simpleStreamingServiceMethods.ByName("SimpleStreamingUnaryInvoke")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.streaming.v1.SimpleStreamingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SimpleStreamingServiceSimpleStreamingUnaryInvokeProcedure:
			simpleStreamingServiceSimpleStreamingUnaryInvokeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSimpleStreamingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSimpleStreamingServiceHandler struct{}

func (UnimplementedSimpleStreamingServiceHandler) SimpleStreamingUnaryInvoke(context.Context, *connect.Request[v1.SimpleStreamingUnaryInvokeRequest]) (*connect.Response[v1.SimpleStreamingUnaryInvokeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.streaming.v1.SimpleStreamingService.SimpleStreamingUnaryInvoke is not implemented"))
}
