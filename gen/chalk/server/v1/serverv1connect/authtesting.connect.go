// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/authtesting.proto

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
	// AuthTestingServiceName is the fully-qualified name of the AuthTestingService service.
	AuthTestingServiceName = "chalk.server.v1.AuthTestingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthTestingServiceGetUnauthedTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetUnauthedTestEndpoint RPC.
	AuthTestingServiceGetUnauthedTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetUnauthedTestEndpoint"
	// AuthTestingServiceGetAuthedTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetAuthedTestEndpoint RPC.
	AuthTestingServiceGetAuthedTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetAuthedTestEndpoint"
	// AuthTestingServiceGetViewerTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetViewerTestEndpoint RPC.
	AuthTestingServiceGetViewerTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetViewerTestEndpoint"
	// AuthTestingServiceGetDataScientistTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetDataScientistTestEndpoint RPC.
	AuthTestingServiceGetDataScientistTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetDataScientistTestEndpoint"
	// AuthTestingServiceGetDeveloperTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetDeveloperTestEndpoint RPC.
	AuthTestingServiceGetDeveloperTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetDeveloperTestEndpoint"
	// AuthTestingServiceGetAdminTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetAdminTestEndpoint RPC.
	AuthTestingServiceGetAdminTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetAdminTestEndpoint"
	// AuthTestingServiceGetOwnerTestEndpointProcedure is the fully-qualified name of the
	// AuthTestingService's GetOwnerTestEndpoint RPC.
	AuthTestingServiceGetOwnerTestEndpointProcedure = "/chalk.server.v1.AuthTestingService/GetOwnerTestEndpoint"
)

// AuthTestingServiceClient is a client for the chalk.server.v1.AuthTestingService service.
type AuthTestingServiceClient interface {
	GetUnauthedTestEndpoint(context.Context, *connect.Request[v1.GetUnauthedTestEndpointRequest]) (*connect.Response[v1.GetUnauthedTestEndpointResponse], error)
	GetAuthedTestEndpoint(context.Context, *connect.Request[v1.GetAuthedTestEndpointRequest]) (*connect.Response[v1.GetAuthedTestEndpointResponse], error)
	GetViewerTestEndpoint(context.Context, *connect.Request[v1.GetViewerTestEndpointRequest]) (*connect.Response[v1.GetViewerTestEndpointResponse], error)
	GetDataScientistTestEndpoint(context.Context, *connect.Request[v1.GetDataScientistTestEndpointRequest]) (*connect.Response[v1.GetDataScientistTestEndpointResponse], error)
	GetDeveloperTestEndpoint(context.Context, *connect.Request[v1.GetDeveloperTestEndpointRequest]) (*connect.Response[v1.GetDeveloperTestEndpointResponse], error)
	GetAdminTestEndpoint(context.Context, *connect.Request[v1.GetAdminTestEndpointRequest]) (*connect.Response[v1.GetAdminTestEndpointResponse], error)
	GetOwnerTestEndpoint(context.Context, *connect.Request[v1.GetOwnerTestEndpointRequest]) (*connect.Response[v1.GetOwnerTestEndpointResponse], error)
}

// NewAuthTestingServiceClient constructs a client for the chalk.server.v1.AuthTestingService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthTestingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthTestingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	authTestingServiceMethods := v1.File_chalk_server_v1_authtesting_proto.Services().ByName("AuthTestingService").Methods()
	return &authTestingServiceClient{
		getUnauthedTestEndpoint: connect.NewClient[v1.GetUnauthedTestEndpointRequest, v1.GetUnauthedTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetUnauthedTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetUnauthedTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getAuthedTestEndpoint: connect.NewClient[v1.GetAuthedTestEndpointRequest, v1.GetAuthedTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetAuthedTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetAuthedTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getViewerTestEndpoint: connect.NewClient[v1.GetViewerTestEndpointRequest, v1.GetViewerTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetViewerTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetViewerTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getDataScientistTestEndpoint: connect.NewClient[v1.GetDataScientistTestEndpointRequest, v1.GetDataScientistTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetDataScientistTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetDataScientistTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getDeveloperTestEndpoint: connect.NewClient[v1.GetDeveloperTestEndpointRequest, v1.GetDeveloperTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetDeveloperTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetDeveloperTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getAdminTestEndpoint: connect.NewClient[v1.GetAdminTestEndpointRequest, v1.GetAdminTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetAdminTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetAdminTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getOwnerTestEndpoint: connect.NewClient[v1.GetOwnerTestEndpointRequest, v1.GetOwnerTestEndpointResponse](
			httpClient,
			baseURL+AuthTestingServiceGetOwnerTestEndpointProcedure,
			connect.WithSchema(authTestingServiceMethods.ByName("GetOwnerTestEndpoint")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// authTestingServiceClient implements AuthTestingServiceClient.
type authTestingServiceClient struct {
	getUnauthedTestEndpoint      *connect.Client[v1.GetUnauthedTestEndpointRequest, v1.GetUnauthedTestEndpointResponse]
	getAuthedTestEndpoint        *connect.Client[v1.GetAuthedTestEndpointRequest, v1.GetAuthedTestEndpointResponse]
	getViewerTestEndpoint        *connect.Client[v1.GetViewerTestEndpointRequest, v1.GetViewerTestEndpointResponse]
	getDataScientistTestEndpoint *connect.Client[v1.GetDataScientistTestEndpointRequest, v1.GetDataScientistTestEndpointResponse]
	getDeveloperTestEndpoint     *connect.Client[v1.GetDeveloperTestEndpointRequest, v1.GetDeveloperTestEndpointResponse]
	getAdminTestEndpoint         *connect.Client[v1.GetAdminTestEndpointRequest, v1.GetAdminTestEndpointResponse]
	getOwnerTestEndpoint         *connect.Client[v1.GetOwnerTestEndpointRequest, v1.GetOwnerTestEndpointResponse]
}

// GetUnauthedTestEndpoint calls chalk.server.v1.AuthTestingService.GetUnauthedTestEndpoint.
func (c *authTestingServiceClient) GetUnauthedTestEndpoint(ctx context.Context, req *connect.Request[v1.GetUnauthedTestEndpointRequest]) (*connect.Response[v1.GetUnauthedTestEndpointResponse], error) {
	return c.getUnauthedTestEndpoint.CallUnary(ctx, req)
}

// GetAuthedTestEndpoint calls chalk.server.v1.AuthTestingService.GetAuthedTestEndpoint.
func (c *authTestingServiceClient) GetAuthedTestEndpoint(ctx context.Context, req *connect.Request[v1.GetAuthedTestEndpointRequest]) (*connect.Response[v1.GetAuthedTestEndpointResponse], error) {
	return c.getAuthedTestEndpoint.CallUnary(ctx, req)
}

// GetViewerTestEndpoint calls chalk.server.v1.AuthTestingService.GetViewerTestEndpoint.
func (c *authTestingServiceClient) GetViewerTestEndpoint(ctx context.Context, req *connect.Request[v1.GetViewerTestEndpointRequest]) (*connect.Response[v1.GetViewerTestEndpointResponse], error) {
	return c.getViewerTestEndpoint.CallUnary(ctx, req)
}

// GetDataScientistTestEndpoint calls
// chalk.server.v1.AuthTestingService.GetDataScientistTestEndpoint.
func (c *authTestingServiceClient) GetDataScientistTestEndpoint(ctx context.Context, req *connect.Request[v1.GetDataScientistTestEndpointRequest]) (*connect.Response[v1.GetDataScientistTestEndpointResponse], error) {
	return c.getDataScientistTestEndpoint.CallUnary(ctx, req)
}

// GetDeveloperTestEndpoint calls chalk.server.v1.AuthTestingService.GetDeveloperTestEndpoint.
func (c *authTestingServiceClient) GetDeveloperTestEndpoint(ctx context.Context, req *connect.Request[v1.GetDeveloperTestEndpointRequest]) (*connect.Response[v1.GetDeveloperTestEndpointResponse], error) {
	return c.getDeveloperTestEndpoint.CallUnary(ctx, req)
}

// GetAdminTestEndpoint calls chalk.server.v1.AuthTestingService.GetAdminTestEndpoint.
func (c *authTestingServiceClient) GetAdminTestEndpoint(ctx context.Context, req *connect.Request[v1.GetAdminTestEndpointRequest]) (*connect.Response[v1.GetAdminTestEndpointResponse], error) {
	return c.getAdminTestEndpoint.CallUnary(ctx, req)
}

// GetOwnerTestEndpoint calls chalk.server.v1.AuthTestingService.GetOwnerTestEndpoint.
func (c *authTestingServiceClient) GetOwnerTestEndpoint(ctx context.Context, req *connect.Request[v1.GetOwnerTestEndpointRequest]) (*connect.Response[v1.GetOwnerTestEndpointResponse], error) {
	return c.getOwnerTestEndpoint.CallUnary(ctx, req)
}

// AuthTestingServiceHandler is an implementation of the chalk.server.v1.AuthTestingService service.
type AuthTestingServiceHandler interface {
	GetUnauthedTestEndpoint(context.Context, *connect.Request[v1.GetUnauthedTestEndpointRequest]) (*connect.Response[v1.GetUnauthedTestEndpointResponse], error)
	GetAuthedTestEndpoint(context.Context, *connect.Request[v1.GetAuthedTestEndpointRequest]) (*connect.Response[v1.GetAuthedTestEndpointResponse], error)
	GetViewerTestEndpoint(context.Context, *connect.Request[v1.GetViewerTestEndpointRequest]) (*connect.Response[v1.GetViewerTestEndpointResponse], error)
	GetDataScientistTestEndpoint(context.Context, *connect.Request[v1.GetDataScientistTestEndpointRequest]) (*connect.Response[v1.GetDataScientistTestEndpointResponse], error)
	GetDeveloperTestEndpoint(context.Context, *connect.Request[v1.GetDeveloperTestEndpointRequest]) (*connect.Response[v1.GetDeveloperTestEndpointResponse], error)
	GetAdminTestEndpoint(context.Context, *connect.Request[v1.GetAdminTestEndpointRequest]) (*connect.Response[v1.GetAdminTestEndpointResponse], error)
	GetOwnerTestEndpoint(context.Context, *connect.Request[v1.GetOwnerTestEndpointRequest]) (*connect.Response[v1.GetOwnerTestEndpointResponse], error)
}

// NewAuthTestingServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthTestingServiceHandler(svc AuthTestingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authTestingServiceMethods := v1.File_chalk_server_v1_authtesting_proto.Services().ByName("AuthTestingService").Methods()
	authTestingServiceGetUnauthedTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetUnauthedTestEndpointProcedure,
		svc.GetUnauthedTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetUnauthedTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetAuthedTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetAuthedTestEndpointProcedure,
		svc.GetAuthedTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetAuthedTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetViewerTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetViewerTestEndpointProcedure,
		svc.GetViewerTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetViewerTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetDataScientistTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetDataScientistTestEndpointProcedure,
		svc.GetDataScientistTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetDataScientistTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetDeveloperTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetDeveloperTestEndpointProcedure,
		svc.GetDeveloperTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetDeveloperTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetAdminTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetAdminTestEndpointProcedure,
		svc.GetAdminTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetAdminTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	authTestingServiceGetOwnerTestEndpointHandler := connect.NewUnaryHandler(
		AuthTestingServiceGetOwnerTestEndpointProcedure,
		svc.GetOwnerTestEndpoint,
		connect.WithSchema(authTestingServiceMethods.ByName("GetOwnerTestEndpoint")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.AuthTestingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthTestingServiceGetUnauthedTestEndpointProcedure:
			authTestingServiceGetUnauthedTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetAuthedTestEndpointProcedure:
			authTestingServiceGetAuthedTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetViewerTestEndpointProcedure:
			authTestingServiceGetViewerTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetDataScientistTestEndpointProcedure:
			authTestingServiceGetDataScientistTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetDeveloperTestEndpointProcedure:
			authTestingServiceGetDeveloperTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetAdminTestEndpointProcedure:
			authTestingServiceGetAdminTestEndpointHandler.ServeHTTP(w, r)
		case AuthTestingServiceGetOwnerTestEndpointProcedure:
			authTestingServiceGetOwnerTestEndpointHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthTestingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthTestingServiceHandler struct{}

func (UnimplementedAuthTestingServiceHandler) GetUnauthedTestEndpoint(context.Context, *connect.Request[v1.GetUnauthedTestEndpointRequest]) (*connect.Response[v1.GetUnauthedTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetUnauthedTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetAuthedTestEndpoint(context.Context, *connect.Request[v1.GetAuthedTestEndpointRequest]) (*connect.Response[v1.GetAuthedTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetAuthedTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetViewerTestEndpoint(context.Context, *connect.Request[v1.GetViewerTestEndpointRequest]) (*connect.Response[v1.GetViewerTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetViewerTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetDataScientistTestEndpoint(context.Context, *connect.Request[v1.GetDataScientistTestEndpointRequest]) (*connect.Response[v1.GetDataScientistTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetDataScientistTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetDeveloperTestEndpoint(context.Context, *connect.Request[v1.GetDeveloperTestEndpointRequest]) (*connect.Response[v1.GetDeveloperTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetDeveloperTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetAdminTestEndpoint(context.Context, *connect.Request[v1.GetAdminTestEndpointRequest]) (*connect.Response[v1.GetAdminTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetAdminTestEndpoint is not implemented"))
}

func (UnimplementedAuthTestingServiceHandler) GetOwnerTestEndpoint(context.Context, *connect.Request[v1.GetOwnerTestEndpointRequest]) (*connect.Response[v1.GetOwnerTestEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthTestingService.GetOwnerTestEndpoint is not implemented"))
}
