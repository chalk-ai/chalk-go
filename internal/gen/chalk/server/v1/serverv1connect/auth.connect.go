// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/auth.proto

package serverv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/chalk-ai/chalk-go/internal/gen/chalk/server/v1"
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
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "chalk.server.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceGetTokenProcedure is the fully-qualified name of the AuthService's GetToken RPC.
	AuthServiceGetTokenProcedure = "/chalk.server.v1.AuthService/GetToken"
	// AuthServiceCreateLinkSessionProcedure is the fully-qualified name of the AuthService's
	// CreateLinkSession RPC.
	AuthServiceCreateLinkSessionProcedure = "/chalk.server.v1.AuthService/CreateLinkSession"
	// AuthServiceGetLinkSessionProcedure is the fully-qualified name of the AuthService's
	// GetLinkSession RPC.
	AuthServiceGetLinkSessionProcedure = "/chalk.server.v1.AuthService/GetLinkSession"
	// AuthServiceUpdateLinkSessionProcedure is the fully-qualified name of the AuthService's
	// UpdateLinkSession RPC.
	AuthServiceUpdateLinkSessionProcedure = "/chalk.server.v1.AuthService/UpdateLinkSession"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authServiceServiceDescriptor                 = v1.File_chalk_server_v1_auth_proto.Services().ByName("AuthService")
	authServiceGetTokenMethodDescriptor          = authServiceServiceDescriptor.Methods().ByName("GetToken")
	authServiceCreateLinkSessionMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("CreateLinkSession")
	authServiceGetLinkSessionMethodDescriptor    = authServiceServiceDescriptor.Methods().ByName("GetLinkSession")
	authServiceUpdateLinkSessionMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("UpdateLinkSession")
)

// AuthServiceClient is a client for the chalk.server.v1.AuthService service.
type AuthServiceClient interface {
	GetToken(context.Context, *connect.Request[v1.GetTokenRequest]) (*connect.Response[v1.GetTokenResponse], error)
	CreateLinkSession(context.Context, *connect.Request[v1.CreateLinkSessionRequest]) (*connect.Response[v1.CreateLinkSessionResponse], error)
	GetLinkSession(context.Context, *connect.Request[v1.GetLinkSessionRequest]) (*connect.Response[v1.GetLinkSessionResponse], error)
	UpdateLinkSession(context.Context, *connect.Request[v1.UpdateLinkSessionRequest]) (*connect.Response[v1.UpdateLinkSessionResponse], error)
}

// NewAuthServiceClient constructs a client for the chalk.server.v1.AuthService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		getToken: connect.NewClient[v1.GetTokenRequest, v1.GetTokenResponse](
			httpClient,
			baseURL+AuthServiceGetTokenProcedure,
			connect.WithSchema(authServiceGetTokenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createLinkSession: connect.NewClient[v1.CreateLinkSessionRequest, v1.CreateLinkSessionResponse](
			httpClient,
			baseURL+AuthServiceCreateLinkSessionProcedure,
			connect.WithSchema(authServiceCreateLinkSessionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getLinkSession: connect.NewClient[v1.GetLinkSessionRequest, v1.GetLinkSessionResponse](
			httpClient,
			baseURL+AuthServiceGetLinkSessionProcedure,
			connect.WithSchema(authServiceGetLinkSessionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateLinkSession: connect.NewClient[v1.UpdateLinkSessionRequest, v1.UpdateLinkSessionResponse](
			httpClient,
			baseURL+AuthServiceUpdateLinkSessionProcedure,
			connect.WithSchema(authServiceUpdateLinkSessionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	getToken          *connect.Client[v1.GetTokenRequest, v1.GetTokenResponse]
	createLinkSession *connect.Client[v1.CreateLinkSessionRequest, v1.CreateLinkSessionResponse]
	getLinkSession    *connect.Client[v1.GetLinkSessionRequest, v1.GetLinkSessionResponse]
	updateLinkSession *connect.Client[v1.UpdateLinkSessionRequest, v1.UpdateLinkSessionResponse]
}

// GetToken calls chalk.server.v1.AuthService.GetToken.
func (c *authServiceClient) GetToken(ctx context.Context, req *connect.Request[v1.GetTokenRequest]) (*connect.Response[v1.GetTokenResponse], error) {
	return c.getToken.CallUnary(ctx, req)
}

// CreateLinkSession calls chalk.server.v1.AuthService.CreateLinkSession.
func (c *authServiceClient) CreateLinkSession(ctx context.Context, req *connect.Request[v1.CreateLinkSessionRequest]) (*connect.Response[v1.CreateLinkSessionResponse], error) {
	return c.createLinkSession.CallUnary(ctx, req)
}

// GetLinkSession calls chalk.server.v1.AuthService.GetLinkSession.
func (c *authServiceClient) GetLinkSession(ctx context.Context, req *connect.Request[v1.GetLinkSessionRequest]) (*connect.Response[v1.GetLinkSessionResponse], error) {
	return c.getLinkSession.CallUnary(ctx, req)
}

// UpdateLinkSession calls chalk.server.v1.AuthService.UpdateLinkSession.
func (c *authServiceClient) UpdateLinkSession(ctx context.Context, req *connect.Request[v1.UpdateLinkSessionRequest]) (*connect.Response[v1.UpdateLinkSessionResponse], error) {
	return c.updateLinkSession.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the chalk.server.v1.AuthService service.
type AuthServiceHandler interface {
	GetToken(context.Context, *connect.Request[v1.GetTokenRequest]) (*connect.Response[v1.GetTokenResponse], error)
	CreateLinkSession(context.Context, *connect.Request[v1.CreateLinkSessionRequest]) (*connect.Response[v1.CreateLinkSessionResponse], error)
	GetLinkSession(context.Context, *connect.Request[v1.GetLinkSessionRequest]) (*connect.Response[v1.GetLinkSessionResponse], error)
	UpdateLinkSession(context.Context, *connect.Request[v1.UpdateLinkSessionRequest]) (*connect.Response[v1.UpdateLinkSessionResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceGetTokenHandler := connect.NewUnaryHandler(
		AuthServiceGetTokenProcedure,
		svc.GetToken,
		connect.WithSchema(authServiceGetTokenMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceCreateLinkSessionHandler := connect.NewUnaryHandler(
		AuthServiceCreateLinkSessionProcedure,
		svc.CreateLinkSession,
		connect.WithSchema(authServiceCreateLinkSessionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceGetLinkSessionHandler := connect.NewUnaryHandler(
		AuthServiceGetLinkSessionProcedure,
		svc.GetLinkSession,
		connect.WithSchema(authServiceGetLinkSessionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceUpdateLinkSessionHandler := connect.NewUnaryHandler(
		AuthServiceUpdateLinkSessionProcedure,
		svc.UpdateLinkSession,
		connect.WithSchema(authServiceUpdateLinkSessionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceGetTokenProcedure:
			authServiceGetTokenHandler.ServeHTTP(w, r)
		case AuthServiceCreateLinkSessionProcedure:
			authServiceCreateLinkSessionHandler.ServeHTTP(w, r)
		case AuthServiceGetLinkSessionProcedure:
			authServiceGetLinkSessionHandler.ServeHTTP(w, r)
		case AuthServiceUpdateLinkSessionProcedure:
			authServiceUpdateLinkSessionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) GetToken(context.Context, *connect.Request[v1.GetTokenRequest]) (*connect.Response[v1.GetTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthService.GetToken is not implemented"))
}

func (UnimplementedAuthServiceHandler) CreateLinkSession(context.Context, *connect.Request[v1.CreateLinkSessionRequest]) (*connect.Response[v1.CreateLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthService.CreateLinkSession is not implemented"))
}

func (UnimplementedAuthServiceHandler) GetLinkSession(context.Context, *connect.Request[v1.GetLinkSessionRequest]) (*connect.Response[v1.GetLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthService.GetLinkSession is not implemented"))
}

func (UnimplementedAuthServiceHandler) UpdateLinkSession(context.Context, *connect.Request[v1.UpdateLinkSessionRequest]) (*connect.Response[v1.UpdateLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuthService.UpdateLinkSession is not implemented"))
}
