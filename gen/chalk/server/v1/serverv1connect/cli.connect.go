// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/cli.proto

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
	// CommandLineInterfaceServiceName is the fully-qualified name of the CommandLineInterfaceService
	// service.
	CommandLineInterfaceServiceName = "chalk.server.v1.CommandLineInterfaceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CommandLineInterfaceServiceGetVersionsProcedure is the fully-qualified name of the
	// CommandLineInterfaceService's GetVersions RPC.
	CommandLineInterfaceServiceGetVersionsProcedure = "/chalk.server.v1.CommandLineInterfaceService/GetVersions"
)

// CommandLineInterfaceServiceClient is a client for the chalk.server.v1.CommandLineInterfaceService
// service.
type CommandLineInterfaceServiceClient interface {
	GetVersions(context.Context, *connect.Request[v1.GetVersionsRequest]) (*connect.Response[v1.GetVersionsResponse], error)
}

// NewCommandLineInterfaceServiceClient constructs a client for the
// chalk.server.v1.CommandLineInterfaceService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCommandLineInterfaceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CommandLineInterfaceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	commandLineInterfaceServiceMethods := v1.File_chalk_server_v1_cli_proto.Services().ByName("CommandLineInterfaceService").Methods()
	return &commandLineInterfaceServiceClient{
		getVersions: connect.NewClient[v1.GetVersionsRequest, v1.GetVersionsResponse](
			httpClient,
			baseURL+CommandLineInterfaceServiceGetVersionsProcedure,
			connect.WithSchema(commandLineInterfaceServiceMethods.ByName("GetVersions")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// commandLineInterfaceServiceClient implements CommandLineInterfaceServiceClient.
type commandLineInterfaceServiceClient struct {
	getVersions *connect.Client[v1.GetVersionsRequest, v1.GetVersionsResponse]
}

// GetVersions calls chalk.server.v1.CommandLineInterfaceService.GetVersions.
func (c *commandLineInterfaceServiceClient) GetVersions(ctx context.Context, req *connect.Request[v1.GetVersionsRequest]) (*connect.Response[v1.GetVersionsResponse], error) {
	return c.getVersions.CallUnary(ctx, req)
}

// CommandLineInterfaceServiceHandler is an implementation of the
// chalk.server.v1.CommandLineInterfaceService service.
type CommandLineInterfaceServiceHandler interface {
	GetVersions(context.Context, *connect.Request[v1.GetVersionsRequest]) (*connect.Response[v1.GetVersionsResponse], error)
}

// NewCommandLineInterfaceServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCommandLineInterfaceServiceHandler(svc CommandLineInterfaceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	commandLineInterfaceServiceMethods := v1.File_chalk_server_v1_cli_proto.Services().ByName("CommandLineInterfaceService").Methods()
	commandLineInterfaceServiceGetVersionsHandler := connect.NewUnaryHandler(
		CommandLineInterfaceServiceGetVersionsProcedure,
		svc.GetVersions,
		connect.WithSchema(commandLineInterfaceServiceMethods.ByName("GetVersions")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.CommandLineInterfaceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CommandLineInterfaceServiceGetVersionsProcedure:
			commandLineInterfaceServiceGetVersionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCommandLineInterfaceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCommandLineInterfaceServiceHandler struct{}

func (UnimplementedCommandLineInterfaceServiceHandler) GetVersions(context.Context, *connect.Request[v1.GetVersionsRequest]) (*connect.Response[v1.GetVersionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.CommandLineInterfaceService.GetVersions is not implemented"))
}
