// The StreamingInvokerService is used to communicate between the main process and remote (i.e. subprocess) invokers

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/streaming/v1/streaming_invoker_service.proto

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
	// StreamingInvokerServiceName is the fully-qualified name of the StreamingInvokerService service.
	StreamingInvokerServiceName = "chalk.streaming.v1.StreamingInvokerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StreamingInvokerServiceStreamingUnaryInvokeProcedure is the fully-qualified name of the
	// StreamingInvokerService's StreamingUnaryInvoke RPC.
	StreamingInvokerServiceStreamingUnaryInvokeProcedure = "/chalk.streaming.v1.StreamingInvokerService/StreamingUnaryInvoke"
)

// StreamingInvokerServiceClient is a client for the chalk.streaming.v1.StreamingInvokerService
// service.
type StreamingInvokerServiceClient interface {
	StreamingUnaryInvoke(context.Context, *connect.Request[v1.StreamingUnaryInvokeRequest]) (*connect.Response[v1.StreamingUnaryInvokeResponse], error)
}

// NewStreamingInvokerServiceClient constructs a client for the
// chalk.streaming.v1.StreamingInvokerService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStreamingInvokerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StreamingInvokerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	streamingInvokerServiceMethods := v1.File_chalk_streaming_v1_streaming_invoker_service_proto.Services().ByName("StreamingInvokerService").Methods()
	return &streamingInvokerServiceClient{
		streamingUnaryInvoke: connect.NewClient[v1.StreamingUnaryInvokeRequest, v1.StreamingUnaryInvokeResponse](
			httpClient,
			baseURL+StreamingInvokerServiceStreamingUnaryInvokeProcedure,
			connect.WithSchema(streamingInvokerServiceMethods.ByName("StreamingUnaryInvoke")),
			connect.WithClientOptions(opts...),
		),
	}
}

// streamingInvokerServiceClient implements StreamingInvokerServiceClient.
type streamingInvokerServiceClient struct {
	streamingUnaryInvoke *connect.Client[v1.StreamingUnaryInvokeRequest, v1.StreamingUnaryInvokeResponse]
}

// StreamingUnaryInvoke calls chalk.streaming.v1.StreamingInvokerService.StreamingUnaryInvoke.
func (c *streamingInvokerServiceClient) StreamingUnaryInvoke(ctx context.Context, req *connect.Request[v1.StreamingUnaryInvokeRequest]) (*connect.Response[v1.StreamingUnaryInvokeResponse], error) {
	return c.streamingUnaryInvoke.CallUnary(ctx, req)
}

// StreamingInvokerServiceHandler is an implementation of the
// chalk.streaming.v1.StreamingInvokerService service.
type StreamingInvokerServiceHandler interface {
	StreamingUnaryInvoke(context.Context, *connect.Request[v1.StreamingUnaryInvokeRequest]) (*connect.Response[v1.StreamingUnaryInvokeResponse], error)
}

// NewStreamingInvokerServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStreamingInvokerServiceHandler(svc StreamingInvokerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	streamingInvokerServiceMethods := v1.File_chalk_streaming_v1_streaming_invoker_service_proto.Services().ByName("StreamingInvokerService").Methods()
	streamingInvokerServiceStreamingUnaryInvokeHandler := connect.NewUnaryHandler(
		StreamingInvokerServiceStreamingUnaryInvokeProcedure,
		svc.StreamingUnaryInvoke,
		connect.WithSchema(streamingInvokerServiceMethods.ByName("StreamingUnaryInvoke")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.streaming.v1.StreamingInvokerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StreamingInvokerServiceStreamingUnaryInvokeProcedure:
			streamingInvokerServiceStreamingUnaryInvokeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStreamingInvokerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStreamingInvokerServiceHandler struct{}

func (UnimplementedStreamingInvokerServiceHandler) StreamingUnaryInvoke(context.Context, *connect.Request[v1.StreamingUnaryInvokeRequest]) (*connect.Response[v1.StreamingUnaryInvokeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.streaming.v1.StreamingInvokerService.StreamingUnaryInvoke is not implemented"))
}
