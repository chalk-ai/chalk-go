// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/team.proto

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
	// TeamServiceName is the fully-qualified name of the TeamService service.
	TeamServiceName = "chalk.server.v1.TeamService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TeamServiceGetEnvProcedure is the fully-qualified name of the TeamService's GetEnv RPC.
	TeamServiceGetEnvProcedure = "/chalk.server.v1.TeamService/GetEnv"
	// TeamServiceGetEnvironmentsProcedure is the fully-qualified name of the TeamService's
	// GetEnvironments RPC.
	TeamServiceGetEnvironmentsProcedure = "/chalk.server.v1.TeamService/GetEnvironments"
	// TeamServiceGetAgentProcedure is the fully-qualified name of the TeamService's GetAgent RPC.
	TeamServiceGetAgentProcedure = "/chalk.server.v1.TeamService/GetAgent"
	// TeamServiceGetDisplayAgentProcedure is the fully-qualified name of the TeamService's
	// GetDisplayAgent RPC.
	TeamServiceGetDisplayAgentProcedure = "/chalk.server.v1.TeamService/GetDisplayAgent"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	teamServiceServiceDescriptor               = v1.File_chalk_server_v1_team_proto.Services().ByName("TeamService")
	teamServiceGetEnvMethodDescriptor          = teamServiceServiceDescriptor.Methods().ByName("GetEnv")
	teamServiceGetEnvironmentsMethodDescriptor = teamServiceServiceDescriptor.Methods().ByName("GetEnvironments")
	teamServiceGetAgentMethodDescriptor        = teamServiceServiceDescriptor.Methods().ByName("GetAgent")
	teamServiceGetDisplayAgentMethodDescriptor = teamServiceServiceDescriptor.Methods().ByName("GetDisplayAgent")
)

// TeamServiceClient is a client for the chalk.server.v1.TeamService service.
type TeamServiceClient interface {
	GetEnv(context.Context, *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error)
	GetEnvironments(context.Context, *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error)
	GetAgent(context.Context, *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error)
	GetDisplayAgent(context.Context, *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error)
}

// NewTeamServiceClient constructs a client for the chalk.server.v1.TeamService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTeamServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TeamServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &teamServiceClient{
		getEnv: connect.NewClient[v1.GetEnvRequest, v1.GetEnvResponse](
			httpClient,
			baseURL+TeamServiceGetEnvProcedure,
			connect.WithSchema(teamServiceGetEnvMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getEnvironments: connect.NewClient[v1.GetEnvironmentsRequest, v1.GetEnvironmentsResponse](
			httpClient,
			baseURL+TeamServiceGetEnvironmentsProcedure,
			connect.WithSchema(teamServiceGetEnvironmentsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getAgent: connect.NewClient[v1.GetAgentRequest, v1.GetAgentResponse](
			httpClient,
			baseURL+TeamServiceGetAgentProcedure,
			connect.WithSchema(teamServiceGetAgentMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getDisplayAgent: connect.NewClient[v1.GetDisplayAgentRequest, v1.GetDisplayAgentResponse](
			httpClient,
			baseURL+TeamServiceGetDisplayAgentProcedure,
			connect.WithSchema(teamServiceGetDisplayAgentMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// teamServiceClient implements TeamServiceClient.
type teamServiceClient struct {
	getEnv          *connect.Client[v1.GetEnvRequest, v1.GetEnvResponse]
	getEnvironments *connect.Client[v1.GetEnvironmentsRequest, v1.GetEnvironmentsResponse]
	getAgent        *connect.Client[v1.GetAgentRequest, v1.GetAgentResponse]
	getDisplayAgent *connect.Client[v1.GetDisplayAgentRequest, v1.GetDisplayAgentResponse]
}

// GetEnv calls chalk.server.v1.TeamService.GetEnv.
func (c *teamServiceClient) GetEnv(ctx context.Context, req *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error) {
	return c.getEnv.CallUnary(ctx, req)
}

// GetEnvironments calls chalk.server.v1.TeamService.GetEnvironments.
func (c *teamServiceClient) GetEnvironments(ctx context.Context, req *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error) {
	return c.getEnvironments.CallUnary(ctx, req)
}

// GetAgent calls chalk.server.v1.TeamService.GetAgent.
func (c *teamServiceClient) GetAgent(ctx context.Context, req *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error) {
	return c.getAgent.CallUnary(ctx, req)
}

// GetDisplayAgent calls chalk.server.v1.TeamService.GetDisplayAgent.
func (c *teamServiceClient) GetDisplayAgent(ctx context.Context, req *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error) {
	return c.getDisplayAgent.CallUnary(ctx, req)
}

// TeamServiceHandler is an implementation of the chalk.server.v1.TeamService service.
type TeamServiceHandler interface {
	GetEnv(context.Context, *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error)
	GetEnvironments(context.Context, *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error)
	GetAgent(context.Context, *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error)
	GetDisplayAgent(context.Context, *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error)
}

// NewTeamServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTeamServiceHandler(svc TeamServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	teamServiceGetEnvHandler := connect.NewUnaryHandler(
		TeamServiceGetEnvProcedure,
		svc.GetEnv,
		connect.WithSchema(teamServiceGetEnvMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceGetEnvironmentsHandler := connect.NewUnaryHandler(
		TeamServiceGetEnvironmentsProcedure,
		svc.GetEnvironments,
		connect.WithSchema(teamServiceGetEnvironmentsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceGetAgentHandler := connect.NewUnaryHandler(
		TeamServiceGetAgentProcedure,
		svc.GetAgent,
		connect.WithSchema(teamServiceGetAgentMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceGetDisplayAgentHandler := connect.NewUnaryHandler(
		TeamServiceGetDisplayAgentProcedure,
		svc.GetDisplayAgent,
		connect.WithSchema(teamServiceGetDisplayAgentMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.TeamService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TeamServiceGetEnvProcedure:
			teamServiceGetEnvHandler.ServeHTTP(w, r)
		case TeamServiceGetEnvironmentsProcedure:
			teamServiceGetEnvironmentsHandler.ServeHTTP(w, r)
		case TeamServiceGetAgentProcedure:
			teamServiceGetAgentHandler.ServeHTTP(w, r)
		case TeamServiceGetDisplayAgentProcedure:
			teamServiceGetDisplayAgentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTeamServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTeamServiceHandler struct{}

func (UnimplementedTeamServiceHandler) GetEnv(context.Context, *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetEnv is not implemented"))
}

func (UnimplementedTeamServiceHandler) GetEnvironments(context.Context, *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetEnvironments is not implemented"))
}

func (UnimplementedTeamServiceHandler) GetAgent(context.Context, *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetAgent is not implemented"))
}

func (UnimplementedTeamServiceHandler) GetDisplayAgent(context.Context, *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetDisplayAgent is not implemented"))
}
