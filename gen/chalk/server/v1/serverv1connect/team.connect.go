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
	// TeamServiceGetTeamProcedure is the fully-qualified name of the TeamService's GetTeam RPC.
	TeamServiceGetTeamProcedure = "/chalk.server.v1.TeamService/GetTeam"
	// TeamServiceCreateTeamProcedure is the fully-qualified name of the TeamService's CreateTeam RPC.
	TeamServiceCreateTeamProcedure = "/chalk.server.v1.TeamService/CreateTeam"
	// TeamServiceCreateProjectProcedure is the fully-qualified name of the TeamService's CreateProject
	// RPC.
	TeamServiceCreateProjectProcedure = "/chalk.server.v1.TeamService/CreateProject"
	// TeamServiceCreateEnvironmentProcedure is the fully-qualified name of the TeamService's
	// CreateEnvironment RPC.
	TeamServiceCreateEnvironmentProcedure = "/chalk.server.v1.TeamService/CreateEnvironment"
	// TeamServiceCreateServiceTokenProcedure is the fully-qualified name of the TeamService's
	// CreateServiceToken RPC.
	TeamServiceCreateServiceTokenProcedure = "/chalk.server.v1.TeamService/CreateServiceToken"
	// TeamServiceGetAvailablePermissionsProcedure is the fully-qualified name of the TeamService's
	// GetAvailablePermissions RPC.
	TeamServiceGetAvailablePermissionsProcedure = "/chalk.server.v1.TeamService/GetAvailablePermissions"
	// TeamServiceDeleteServiceTokenProcedure is the fully-qualified name of the TeamService's
	// DeleteServiceToken RPC.
	TeamServiceDeleteServiceTokenProcedure = "/chalk.server.v1.TeamService/DeleteServiceToken"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	teamServiceServiceDescriptor                       = v1.File_chalk_server_v1_team_proto.Services().ByName("TeamService")
	teamServiceGetEnvMethodDescriptor                  = teamServiceServiceDescriptor.Methods().ByName("GetEnv")
	teamServiceGetEnvironmentsMethodDescriptor         = teamServiceServiceDescriptor.Methods().ByName("GetEnvironments")
	teamServiceGetAgentMethodDescriptor                = teamServiceServiceDescriptor.Methods().ByName("GetAgent")
	teamServiceGetDisplayAgentMethodDescriptor         = teamServiceServiceDescriptor.Methods().ByName("GetDisplayAgent")
	teamServiceGetTeamMethodDescriptor                 = teamServiceServiceDescriptor.Methods().ByName("GetTeam")
	teamServiceCreateTeamMethodDescriptor              = teamServiceServiceDescriptor.Methods().ByName("CreateTeam")
	teamServiceCreateProjectMethodDescriptor           = teamServiceServiceDescriptor.Methods().ByName("CreateProject")
	teamServiceCreateEnvironmentMethodDescriptor       = teamServiceServiceDescriptor.Methods().ByName("CreateEnvironment")
	teamServiceCreateServiceTokenMethodDescriptor      = teamServiceServiceDescriptor.Methods().ByName("CreateServiceToken")
	teamServiceGetAvailablePermissionsMethodDescriptor = teamServiceServiceDescriptor.Methods().ByName("GetAvailablePermissions")
	teamServiceDeleteServiceTokenMethodDescriptor      = teamServiceServiceDescriptor.Methods().ByName("DeleteServiceToken")
)

// TeamServiceClient is a client for the chalk.server.v1.TeamService service.
type TeamServiceClient interface {
	GetEnv(context.Context, *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error)
	GetEnvironments(context.Context, *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error)
	GetAgent(context.Context, *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error)
	GetDisplayAgent(context.Context, *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error)
	GetTeam(context.Context, *connect.Request[v1.GetTeamRequest]) (*connect.Response[v1.GetTeamResponse], error)
	CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error)
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error)
	CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error)
	CreateServiceToken(context.Context, *connect.Request[v1.CreateServiceTokenRequest]) (*connect.Response[v1.CreateServiceTokenResponse], error)
	GetAvailablePermissions(context.Context, *connect.Request[v1.GetAvailablePermissionsRequest]) (*connect.Response[v1.GetAvailablePermissionsResponse], error)
	DeleteServiceToken(context.Context, *connect.Request[v1.DeleteServiceTokenRequest]) (*connect.Response[v1.DeleteServiceTokenResponse], error)
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
		getTeam: connect.NewClient[v1.GetTeamRequest, v1.GetTeamResponse](
			httpClient,
			baseURL+TeamServiceGetTeamProcedure,
			connect.WithSchema(teamServiceGetTeamMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createTeam: connect.NewClient[v1.CreateTeamRequest, v1.CreateTeamResponse](
			httpClient,
			baseURL+TeamServiceCreateTeamProcedure,
			connect.WithSchema(teamServiceCreateTeamMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProject: connect.NewClient[v1.CreateProjectRequest, v1.CreateProjectResponse](
			httpClient,
			baseURL+TeamServiceCreateProjectProcedure,
			connect.WithSchema(teamServiceCreateProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createEnvironment: connect.NewClient[v1.CreateEnvironmentRequest, v1.CreateEnvironmentResponse](
			httpClient,
			baseURL+TeamServiceCreateEnvironmentProcedure,
			connect.WithSchema(teamServiceCreateEnvironmentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createServiceToken: connect.NewClient[v1.CreateServiceTokenRequest, v1.CreateServiceTokenResponse](
			httpClient,
			baseURL+TeamServiceCreateServiceTokenProcedure,
			connect.WithSchema(teamServiceCreateServiceTokenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAvailablePermissions: connect.NewClient[v1.GetAvailablePermissionsRequest, v1.GetAvailablePermissionsResponse](
			httpClient,
			baseURL+TeamServiceGetAvailablePermissionsProcedure,
			connect.WithSchema(teamServiceGetAvailablePermissionsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		deleteServiceToken: connect.NewClient[v1.DeleteServiceTokenRequest, v1.DeleteServiceTokenResponse](
			httpClient,
			baseURL+TeamServiceDeleteServiceTokenProcedure,
			connect.WithSchema(teamServiceDeleteServiceTokenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// teamServiceClient implements TeamServiceClient.
type teamServiceClient struct {
	getEnv                  *connect.Client[v1.GetEnvRequest, v1.GetEnvResponse]
	getEnvironments         *connect.Client[v1.GetEnvironmentsRequest, v1.GetEnvironmentsResponse]
	getAgent                *connect.Client[v1.GetAgentRequest, v1.GetAgentResponse]
	getDisplayAgent         *connect.Client[v1.GetDisplayAgentRequest, v1.GetDisplayAgentResponse]
	getTeam                 *connect.Client[v1.GetTeamRequest, v1.GetTeamResponse]
	createTeam              *connect.Client[v1.CreateTeamRequest, v1.CreateTeamResponse]
	createProject           *connect.Client[v1.CreateProjectRequest, v1.CreateProjectResponse]
	createEnvironment       *connect.Client[v1.CreateEnvironmentRequest, v1.CreateEnvironmentResponse]
	createServiceToken      *connect.Client[v1.CreateServiceTokenRequest, v1.CreateServiceTokenResponse]
	getAvailablePermissions *connect.Client[v1.GetAvailablePermissionsRequest, v1.GetAvailablePermissionsResponse]
	deleteServiceToken      *connect.Client[v1.DeleteServiceTokenRequest, v1.DeleteServiceTokenResponse]
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

// GetTeam calls chalk.server.v1.TeamService.GetTeam.
func (c *teamServiceClient) GetTeam(ctx context.Context, req *connect.Request[v1.GetTeamRequest]) (*connect.Response[v1.GetTeamResponse], error) {
	return c.getTeam.CallUnary(ctx, req)
}

// CreateTeam calls chalk.server.v1.TeamService.CreateTeam.
func (c *teamServiceClient) CreateTeam(ctx context.Context, req *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error) {
	return c.createTeam.CallUnary(ctx, req)
}

// CreateProject calls chalk.server.v1.TeamService.CreateProject.
func (c *teamServiceClient) CreateProject(ctx context.Context, req *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error) {
	return c.createProject.CallUnary(ctx, req)
}

// CreateEnvironment calls chalk.server.v1.TeamService.CreateEnvironment.
func (c *teamServiceClient) CreateEnvironment(ctx context.Context, req *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
	return c.createEnvironment.CallUnary(ctx, req)
}

// CreateServiceToken calls chalk.server.v1.TeamService.CreateServiceToken.
func (c *teamServiceClient) CreateServiceToken(ctx context.Context, req *connect.Request[v1.CreateServiceTokenRequest]) (*connect.Response[v1.CreateServiceTokenResponse], error) {
	return c.createServiceToken.CallUnary(ctx, req)
}

// GetAvailablePermissions calls chalk.server.v1.TeamService.GetAvailablePermissions.
func (c *teamServiceClient) GetAvailablePermissions(ctx context.Context, req *connect.Request[v1.GetAvailablePermissionsRequest]) (*connect.Response[v1.GetAvailablePermissionsResponse], error) {
	return c.getAvailablePermissions.CallUnary(ctx, req)
}

// DeleteServiceToken calls chalk.server.v1.TeamService.DeleteServiceToken.
func (c *teamServiceClient) DeleteServiceToken(ctx context.Context, req *connect.Request[v1.DeleteServiceTokenRequest]) (*connect.Response[v1.DeleteServiceTokenResponse], error) {
	return c.deleteServiceToken.CallUnary(ctx, req)
}

// TeamServiceHandler is an implementation of the chalk.server.v1.TeamService service.
type TeamServiceHandler interface {
	GetEnv(context.Context, *connect.Request[v1.GetEnvRequest]) (*connect.Response[v1.GetEnvResponse], error)
	GetEnvironments(context.Context, *connect.Request[v1.GetEnvironmentsRequest]) (*connect.Response[v1.GetEnvironmentsResponse], error)
	GetAgent(context.Context, *connect.Request[v1.GetAgentRequest]) (*connect.Response[v1.GetAgentResponse], error)
	GetDisplayAgent(context.Context, *connect.Request[v1.GetDisplayAgentRequest]) (*connect.Response[v1.GetDisplayAgentResponse], error)
	GetTeam(context.Context, *connect.Request[v1.GetTeamRequest]) (*connect.Response[v1.GetTeamResponse], error)
	CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error)
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error)
	CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error)
	CreateServiceToken(context.Context, *connect.Request[v1.CreateServiceTokenRequest]) (*connect.Response[v1.CreateServiceTokenResponse], error)
	GetAvailablePermissions(context.Context, *connect.Request[v1.GetAvailablePermissionsRequest]) (*connect.Response[v1.GetAvailablePermissionsResponse], error)
	DeleteServiceToken(context.Context, *connect.Request[v1.DeleteServiceTokenRequest]) (*connect.Response[v1.DeleteServiceTokenResponse], error)
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
	teamServiceGetTeamHandler := connect.NewUnaryHandler(
		TeamServiceGetTeamProcedure,
		svc.GetTeam,
		connect.WithSchema(teamServiceGetTeamMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceCreateTeamHandler := connect.NewUnaryHandler(
		TeamServiceCreateTeamProcedure,
		svc.CreateTeam,
		connect.WithSchema(teamServiceCreateTeamMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceCreateProjectHandler := connect.NewUnaryHandler(
		TeamServiceCreateProjectProcedure,
		svc.CreateProject,
		connect.WithSchema(teamServiceCreateProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceCreateEnvironmentHandler := connect.NewUnaryHandler(
		TeamServiceCreateEnvironmentProcedure,
		svc.CreateEnvironment,
		connect.WithSchema(teamServiceCreateEnvironmentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceCreateServiceTokenHandler := connect.NewUnaryHandler(
		TeamServiceCreateServiceTokenProcedure,
		svc.CreateServiceToken,
		connect.WithSchema(teamServiceCreateServiceTokenMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceGetAvailablePermissionsHandler := connect.NewUnaryHandler(
		TeamServiceGetAvailablePermissionsProcedure,
		svc.GetAvailablePermissions,
		connect.WithSchema(teamServiceGetAvailablePermissionsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	teamServiceDeleteServiceTokenHandler := connect.NewUnaryHandler(
		TeamServiceDeleteServiceTokenProcedure,
		svc.DeleteServiceToken,
		connect.WithSchema(teamServiceDeleteServiceTokenMethodDescriptor),
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
		case TeamServiceGetTeamProcedure:
			teamServiceGetTeamHandler.ServeHTTP(w, r)
		case TeamServiceCreateTeamProcedure:
			teamServiceCreateTeamHandler.ServeHTTP(w, r)
		case TeamServiceCreateProjectProcedure:
			teamServiceCreateProjectHandler.ServeHTTP(w, r)
		case TeamServiceCreateEnvironmentProcedure:
			teamServiceCreateEnvironmentHandler.ServeHTTP(w, r)
		case TeamServiceCreateServiceTokenProcedure:
			teamServiceCreateServiceTokenHandler.ServeHTTP(w, r)
		case TeamServiceGetAvailablePermissionsProcedure:
			teamServiceGetAvailablePermissionsHandler.ServeHTTP(w, r)
		case TeamServiceDeleteServiceTokenProcedure:
			teamServiceDeleteServiceTokenHandler.ServeHTTP(w, r)
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

func (UnimplementedTeamServiceHandler) GetTeam(context.Context, *connect.Request[v1.GetTeamRequest]) (*connect.Response[v1.GetTeamResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetTeam is not implemented"))
}

func (UnimplementedTeamServiceHandler) CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.CreateTeam is not implemented"))
}

func (UnimplementedTeamServiceHandler) CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.CreateProject is not implemented"))
}

func (UnimplementedTeamServiceHandler) CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.CreateEnvironment is not implemented"))
}

func (UnimplementedTeamServiceHandler) CreateServiceToken(context.Context, *connect.Request[v1.CreateServiceTokenRequest]) (*connect.Response[v1.CreateServiceTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.CreateServiceToken is not implemented"))
}

func (UnimplementedTeamServiceHandler) GetAvailablePermissions(context.Context, *connect.Request[v1.GetAvailablePermissionsRequest]) (*connect.Response[v1.GetAvailablePermissionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.GetAvailablePermissions is not implemented"))
}

func (UnimplementedTeamServiceHandler) DeleteServiceToken(context.Context, *connect.Request[v1.DeleteServiceTokenRequest]) (*connect.Response[v1.DeleteServiceTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.TeamService.DeleteServiceToken is not implemented"))
}
