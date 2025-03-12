// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/named_query.proto

package serverv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/server/v1"
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
	// NamedQueryServiceName is the fully-qualified name of the NamedQueryService service.
	NamedQueryServiceName = "chalk.server.v1.NamedQueryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NamedQueryServiceGetAllNamedQueriesProcedure is the fully-qualified name of the
	// NamedQueryService's GetAllNamedQueries RPC.
	NamedQueryServiceGetAllNamedQueriesProcedure = "/chalk.server.v1.NamedQueryService/GetAllNamedQueries"
	// NamedQueryServiceGetAllNamedQueriesActiveDeploymentProcedure is the fully-qualified name of the
	// NamedQueryService's GetAllNamedQueriesActiveDeployment RPC.
	NamedQueryServiceGetAllNamedQueriesActiveDeploymentProcedure = "/chalk.server.v1.NamedQueryService/GetAllNamedQueriesActiveDeployment"
	// NamedQueryServiceGetNamedQueryByNameProcedure is the fully-qualified name of the
	// NamedQueryService's GetNamedQueryByName RPC.
	NamedQueryServiceGetNamedQueryByNameProcedure = "/chalk.server.v1.NamedQueryService/GetNamedQueryByName"
)

// NamedQueryServiceClient is a client for the chalk.server.v1.NamedQueryService service.
type NamedQueryServiceClient interface {
	GetAllNamedQueries(context.Context, *connect.Request[v1.GetAllNamedQueriesRequest]) (*connect.Response[v1.GetAllNamedQueriesResponse], error)
	GetAllNamedQueriesActiveDeployment(context.Context, *connect.Request[v1.GetAllNamedQueriesActiveDeploymentRequest]) (*connect.Response[v1.GetAllNamedQueriesActiveDeploymentResponse], error)
	GetNamedQueryByName(context.Context, *connect.Request[v1.GetNamedQueryByNameRequest]) (*connect.Response[v1.GetNamedQueryByNameResponse], error)
}

// NewNamedQueryServiceClient constructs a client for the chalk.server.v1.NamedQueryService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNamedQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NamedQueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	namedQueryServiceMethods := v1.File_chalk_server_v1_named_query_proto.Services().ByName("NamedQueryService").Methods()
	return &namedQueryServiceClient{
		getAllNamedQueries: connect.NewClient[v1.GetAllNamedQueriesRequest, v1.GetAllNamedQueriesResponse](
			httpClient,
			baseURL+NamedQueryServiceGetAllNamedQueriesProcedure,
			connect.WithSchema(namedQueryServiceMethods.ByName("GetAllNamedQueries")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getAllNamedQueriesActiveDeployment: connect.NewClient[v1.GetAllNamedQueriesActiveDeploymentRequest, v1.GetAllNamedQueriesActiveDeploymentResponse](
			httpClient,
			baseURL+NamedQueryServiceGetAllNamedQueriesActiveDeploymentProcedure,
			connect.WithSchema(namedQueryServiceMethods.ByName("GetAllNamedQueriesActiveDeployment")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getNamedQueryByName: connect.NewClient[v1.GetNamedQueryByNameRequest, v1.GetNamedQueryByNameResponse](
			httpClient,
			baseURL+NamedQueryServiceGetNamedQueryByNameProcedure,
			connect.WithSchema(namedQueryServiceMethods.ByName("GetNamedQueryByName")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// namedQueryServiceClient implements NamedQueryServiceClient.
type namedQueryServiceClient struct {
	getAllNamedQueries                 *connect.Client[v1.GetAllNamedQueriesRequest, v1.GetAllNamedQueriesResponse]
	getAllNamedQueriesActiveDeployment *connect.Client[v1.GetAllNamedQueriesActiveDeploymentRequest, v1.GetAllNamedQueriesActiveDeploymentResponse]
	getNamedQueryByName                *connect.Client[v1.GetNamedQueryByNameRequest, v1.GetNamedQueryByNameResponse]
}

// GetAllNamedQueries calls chalk.server.v1.NamedQueryService.GetAllNamedQueries.
func (c *namedQueryServiceClient) GetAllNamedQueries(ctx context.Context, req *connect.Request[v1.GetAllNamedQueriesRequest]) (*connect.Response[v1.GetAllNamedQueriesResponse], error) {
	return c.getAllNamedQueries.CallUnary(ctx, req)
}

// GetAllNamedQueriesActiveDeployment calls
// chalk.server.v1.NamedQueryService.GetAllNamedQueriesActiveDeployment.
func (c *namedQueryServiceClient) GetAllNamedQueriesActiveDeployment(ctx context.Context, req *connect.Request[v1.GetAllNamedQueriesActiveDeploymentRequest]) (*connect.Response[v1.GetAllNamedQueriesActiveDeploymentResponse], error) {
	return c.getAllNamedQueriesActiveDeployment.CallUnary(ctx, req)
}

// GetNamedQueryByName calls chalk.server.v1.NamedQueryService.GetNamedQueryByName.
func (c *namedQueryServiceClient) GetNamedQueryByName(ctx context.Context, req *connect.Request[v1.GetNamedQueryByNameRequest]) (*connect.Response[v1.GetNamedQueryByNameResponse], error) {
	return c.getNamedQueryByName.CallUnary(ctx, req)
}

// NamedQueryServiceHandler is an implementation of the chalk.server.v1.NamedQueryService service.
type NamedQueryServiceHandler interface {
	GetAllNamedQueries(context.Context, *connect.Request[v1.GetAllNamedQueriesRequest]) (*connect.Response[v1.GetAllNamedQueriesResponse], error)
	GetAllNamedQueriesActiveDeployment(context.Context, *connect.Request[v1.GetAllNamedQueriesActiveDeploymentRequest]) (*connect.Response[v1.GetAllNamedQueriesActiveDeploymentResponse], error)
	GetNamedQueryByName(context.Context, *connect.Request[v1.GetNamedQueryByNameRequest]) (*connect.Response[v1.GetNamedQueryByNameResponse], error)
}

// NewNamedQueryServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNamedQueryServiceHandler(svc NamedQueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	namedQueryServiceMethods := v1.File_chalk_server_v1_named_query_proto.Services().ByName("NamedQueryService").Methods()
	namedQueryServiceGetAllNamedQueriesHandler := connect.NewUnaryHandler(
		NamedQueryServiceGetAllNamedQueriesProcedure,
		svc.GetAllNamedQueries,
		connect.WithSchema(namedQueryServiceMethods.ByName("GetAllNamedQueries")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	namedQueryServiceGetAllNamedQueriesActiveDeploymentHandler := connect.NewUnaryHandler(
		NamedQueryServiceGetAllNamedQueriesActiveDeploymentProcedure,
		svc.GetAllNamedQueriesActiveDeployment,
		connect.WithSchema(namedQueryServiceMethods.ByName("GetAllNamedQueriesActiveDeployment")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	namedQueryServiceGetNamedQueryByNameHandler := connect.NewUnaryHandler(
		NamedQueryServiceGetNamedQueryByNameProcedure,
		svc.GetNamedQueryByName,
		connect.WithSchema(namedQueryServiceMethods.ByName("GetNamedQueryByName")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.NamedQueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NamedQueryServiceGetAllNamedQueriesProcedure:
			namedQueryServiceGetAllNamedQueriesHandler.ServeHTTP(w, r)
		case NamedQueryServiceGetAllNamedQueriesActiveDeploymentProcedure:
			namedQueryServiceGetAllNamedQueriesActiveDeploymentHandler.ServeHTTP(w, r)
		case NamedQueryServiceGetNamedQueryByNameProcedure:
			namedQueryServiceGetNamedQueryByNameHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNamedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNamedQueryServiceHandler struct{}

func (UnimplementedNamedQueryServiceHandler) GetAllNamedQueries(context.Context, *connect.Request[v1.GetAllNamedQueriesRequest]) (*connect.Response[v1.GetAllNamedQueriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.NamedQueryService.GetAllNamedQueries is not implemented"))
}

func (UnimplementedNamedQueryServiceHandler) GetAllNamedQueriesActiveDeployment(context.Context, *connect.Request[v1.GetAllNamedQueriesActiveDeploymentRequest]) (*connect.Response[v1.GetAllNamedQueriesActiveDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.NamedQueryService.GetAllNamedQueriesActiveDeployment is not implemented"))
}

func (UnimplementedNamedQueryServiceHandler) GetNamedQueryByName(context.Context, *connect.Request[v1.GetNamedQueryByNameRequest]) (*connect.Response[v1.GetNamedQueryByNameResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.NamedQueryService.GetNamedQueryByName is not implemented"))
}
