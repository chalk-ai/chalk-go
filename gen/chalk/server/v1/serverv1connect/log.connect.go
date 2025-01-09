// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/log.proto

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
	// LogSearchServiceName is the fully-qualified name of the LogSearchService service.
	LogSearchServiceName = "chalk.server.v1.LogSearchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LogSearchServiceSearchLogEntriesProcedure is the fully-qualified name of the LogSearchService's
	// SearchLogEntries RPC.
	LogSearchServiceSearchLogEntriesProcedure = "/chalk.server.v1.LogSearchService/SearchLogEntries"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	logSearchServiceServiceDescriptor                = v1.File_chalk_server_v1_log_proto.Services().ByName("LogSearchService")
	logSearchServiceSearchLogEntriesMethodDescriptor = logSearchServiceServiceDescriptor.Methods().ByName("SearchLogEntries")
)

// LogSearchServiceClient is a client for the chalk.server.v1.LogSearchService service.
type LogSearchServiceClient interface {
	SearchLogEntries(context.Context, *connect.Request[v1.SearchLogEntriesRequest]) (*connect.Response[v1.SearchLogEntriesResponse], error)
}

// NewLogSearchServiceClient constructs a client for the chalk.server.v1.LogSearchService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLogSearchServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LogSearchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &logSearchServiceClient{
		searchLogEntries: connect.NewClient[v1.SearchLogEntriesRequest, v1.SearchLogEntriesResponse](
			httpClient,
			baseURL+LogSearchServiceSearchLogEntriesProcedure,
			connect.WithSchema(logSearchServiceSearchLogEntriesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// logSearchServiceClient implements LogSearchServiceClient.
type logSearchServiceClient struct {
	searchLogEntries *connect.Client[v1.SearchLogEntriesRequest, v1.SearchLogEntriesResponse]
}

// SearchLogEntries calls chalk.server.v1.LogSearchService.SearchLogEntries.
func (c *logSearchServiceClient) SearchLogEntries(ctx context.Context, req *connect.Request[v1.SearchLogEntriesRequest]) (*connect.Response[v1.SearchLogEntriesResponse], error) {
	return c.searchLogEntries.CallUnary(ctx, req)
}

// LogSearchServiceHandler is an implementation of the chalk.server.v1.LogSearchService service.
type LogSearchServiceHandler interface {
	SearchLogEntries(context.Context, *connect.Request[v1.SearchLogEntriesRequest]) (*connect.Response[v1.SearchLogEntriesResponse], error)
}

// NewLogSearchServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLogSearchServiceHandler(svc LogSearchServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	logSearchServiceSearchLogEntriesHandler := connect.NewUnaryHandler(
		LogSearchServiceSearchLogEntriesProcedure,
		svc.SearchLogEntries,
		connect.WithSchema(logSearchServiceSearchLogEntriesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.LogSearchService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LogSearchServiceSearchLogEntriesProcedure:
			logSearchServiceSearchLogEntriesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLogSearchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLogSearchServiceHandler struct{}

func (UnimplementedLogSearchServiceHandler) SearchLogEntries(context.Context, *connect.Request[v1.SearchLogEntriesRequest]) (*connect.Response[v1.SearchLogEntriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.LogSearchService.SearchLogEntries is not implemented"))
}
