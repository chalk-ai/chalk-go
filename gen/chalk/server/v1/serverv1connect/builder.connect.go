// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/builder.proto

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
	// BuilderServiceName is the fully-qualified name of the BuilderService service.
	BuilderServiceName = "chalk.server.v1.BuilderService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BuilderServiceActivateDeploymentProcedure is the fully-qualified name of the BuilderService's
	// ActivateDeployment RPC.
	BuilderServiceActivateDeploymentProcedure = "/chalk.server.v1.BuilderService/ActivateDeployment"
	// BuilderServiceIndexDeploymentProcedure is the fully-qualified name of the BuilderService's
	// IndexDeployment RPC.
	BuilderServiceIndexDeploymentProcedure = "/chalk.server.v1.BuilderService/IndexDeployment"
	// BuilderServiceDeployKubeComponentsProcedure is the fully-qualified name of the BuilderService's
	// DeployKubeComponents RPC.
	BuilderServiceDeployKubeComponentsProcedure = "/chalk.server.v1.BuilderService/DeployKubeComponents"
	// BuilderServiceRebuildDeploymentProcedure is the fully-qualified name of the BuilderService's
	// RebuildDeployment RPC.
	BuilderServiceRebuildDeploymentProcedure = "/chalk.server.v1.BuilderService/RebuildDeployment"
	// BuilderServiceRedeployDeploymentProcedure is the fully-qualified name of the BuilderService's
	// RedeployDeployment RPC.
	BuilderServiceRedeployDeploymentProcedure = "/chalk.server.v1.BuilderService/RedeployDeployment"
	// BuilderServiceUploadSourceProcedure is the fully-qualified name of the BuilderService's
	// UploadSource RPC.
	BuilderServiceUploadSourceProcedure = "/chalk.server.v1.BuilderService/UploadSource"
	// BuilderServiceGetDeploymentStepsProcedure is the fully-qualified name of the BuilderService's
	// GetDeploymentSteps RPC.
	BuilderServiceGetDeploymentStepsProcedure = "/chalk.server.v1.BuilderService/GetDeploymentSteps"
	// BuilderServiceGetDeploymentLogsProcedure is the fully-qualified name of the BuilderService's
	// GetDeploymentLogs RPC.
	BuilderServiceGetDeploymentLogsProcedure = "/chalk.server.v1.BuilderService/GetDeploymentLogs"
	// BuilderServiceGetClusterTimescaleDBProcedure is the fully-qualified name of the BuilderService's
	// GetClusterTimescaleDB RPC.
	BuilderServiceGetClusterTimescaleDBProcedure = "/chalk.server.v1.BuilderService/GetClusterTimescaleDB"
	// BuilderServiceGetClusterGatewayProcedure is the fully-qualified name of the BuilderService's
	// GetClusterGateway RPC.
	BuilderServiceGetClusterGatewayProcedure = "/chalk.server.v1.BuilderService/GetClusterGateway"
	// BuilderServiceCreateClusterTimescaleDBProcedure is the fully-qualified name of the
	// BuilderService's CreateClusterTimescaleDB RPC.
	BuilderServiceCreateClusterTimescaleDBProcedure = "/chalk.server.v1.BuilderService/CreateClusterTimescaleDB"
	// BuilderServiceCreateClusterGatewayProcedure is the fully-qualified name of the BuilderService's
	// CreateClusterGateway RPC.
	BuilderServiceCreateClusterGatewayProcedure = "/chalk.server.v1.BuilderService/CreateClusterGateway"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	builderServiceServiceDescriptor                        = v1.File_chalk_server_v1_builder_proto.Services().ByName("BuilderService")
	builderServiceActivateDeploymentMethodDescriptor       = builderServiceServiceDescriptor.Methods().ByName("ActivateDeployment")
	builderServiceIndexDeploymentMethodDescriptor          = builderServiceServiceDescriptor.Methods().ByName("IndexDeployment")
	builderServiceDeployKubeComponentsMethodDescriptor     = builderServiceServiceDescriptor.Methods().ByName("DeployKubeComponents")
	builderServiceRebuildDeploymentMethodDescriptor        = builderServiceServiceDescriptor.Methods().ByName("RebuildDeployment")
	builderServiceRedeployDeploymentMethodDescriptor       = builderServiceServiceDescriptor.Methods().ByName("RedeployDeployment")
	builderServiceUploadSourceMethodDescriptor             = builderServiceServiceDescriptor.Methods().ByName("UploadSource")
	builderServiceGetDeploymentStepsMethodDescriptor       = builderServiceServiceDescriptor.Methods().ByName("GetDeploymentSteps")
	builderServiceGetDeploymentLogsMethodDescriptor        = builderServiceServiceDescriptor.Methods().ByName("GetDeploymentLogs")
	builderServiceGetClusterTimescaleDBMethodDescriptor    = builderServiceServiceDescriptor.Methods().ByName("GetClusterTimescaleDB")
	builderServiceGetClusterGatewayMethodDescriptor        = builderServiceServiceDescriptor.Methods().ByName("GetClusterGateway")
	builderServiceCreateClusterTimescaleDBMethodDescriptor = builderServiceServiceDescriptor.Methods().ByName("CreateClusterTimescaleDB")
	builderServiceCreateClusterGatewayMethodDescriptor     = builderServiceServiceDescriptor.Methods().ByName("CreateClusterGateway")
)

// BuilderServiceClient is a client for the chalk.server.v1.BuilderService service.
type BuilderServiceClient interface {
	// Takes an existing (past) deployment and promotes the k8s resources / other things associated with it.
	// Useful for debugging in local development where the auto activation doesn't work b/c no pubsub.
	ActivateDeployment(context.Context, *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error)
	IndexDeployment(context.Context, *connect.Request[v1.IndexDeploymentRequest]) (*connect.Response[v1.IndexDeploymentResponse], error)
	// Intermediate step in the deployment activation process. Allows for partial migration to the new
	// go-api-server builder service.
	DeployKubeComponents(context.Context, *connect.Request[v1.DeployKubeComponentsRequest]) (*connect.Response[v1.DeployKubeComponentsResponse], error)
	// Takes an existing (past) deployment and re-creates the image associated with it,
	// publishing the image as 'new_image_tag'.
	RebuildDeployment(context.Context, *connect.Request[v1.RebuildDeploymentRequest]) (*connect.Response[v1.RebuildDeploymentResponse], error)
	// Triggers a new build with the source code from this deployment and deploys the result
	RedeployDeployment(context.Context, *connect.Request[v1.RedeployDeploymentRequest]) (*connect.Response[v1.RedeployDeploymentResponse], error)
	// Triggers a new build with the provided source code archive and deploys the result
	UploadSource(context.Context, *connect.Request[v1.UploadSourceRequest]) (*connect.Response[v1.UploadSourceResponse], error)
	GetDeploymentSteps(context.Context, *connect.Request[v1.GetDeploymentStepsRequest]) (*connect.Response[v1.GetDeploymentStepsResponse], error)
	GetDeploymentLogs(context.Context, *connect.Request[v1.GetDeploymentLogsRequest]) (*connect.Response[v1.GetDeploymentLogsResponse], error)
	GetClusterTimescaleDB(context.Context, *connect.Request[v1.GetClusterTimescaleDBRequest]) (*connect.Response[v1.GetClusterTimescaleDBResponse], error)
	GetClusterGateway(context.Context, *connect.Request[v1.GetClusterGatewayRequest]) (*connect.Response[v1.GetClusterGatewayResponse], error)
	CreateClusterTimescaleDB(context.Context, *connect.Request[v1.CreateClusterTimescaleDBRequest]) (*connect.Response[v1.CreateClusterTimescaleDBResponse], error)
	CreateClusterGateway(context.Context, *connect.Request[v1.CreateClusterGatewayRequest]) (*connect.Response[v1.CreateClusterGatewayResponse], error)
}

// NewBuilderServiceClient constructs a client for the chalk.server.v1.BuilderService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBuilderServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BuilderServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &builderServiceClient{
		activateDeployment: connect.NewClient[v1.ActivateDeploymentRequest, v1.ActivateDeploymentResponse](
			httpClient,
			baseURL+BuilderServiceActivateDeploymentProcedure,
			connect.WithSchema(builderServiceActivateDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		indexDeployment: connect.NewClient[v1.IndexDeploymentRequest, v1.IndexDeploymentResponse](
			httpClient,
			baseURL+BuilderServiceIndexDeploymentProcedure,
			connect.WithSchema(builderServiceIndexDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deployKubeComponents: connect.NewClient[v1.DeployKubeComponentsRequest, v1.DeployKubeComponentsResponse](
			httpClient,
			baseURL+BuilderServiceDeployKubeComponentsProcedure,
			connect.WithSchema(builderServiceDeployKubeComponentsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		rebuildDeployment: connect.NewClient[v1.RebuildDeploymentRequest, v1.RebuildDeploymentResponse](
			httpClient,
			baseURL+BuilderServiceRebuildDeploymentProcedure,
			connect.WithSchema(builderServiceRebuildDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		redeployDeployment: connect.NewClient[v1.RedeployDeploymentRequest, v1.RedeployDeploymentResponse](
			httpClient,
			baseURL+BuilderServiceRedeployDeploymentProcedure,
			connect.WithSchema(builderServiceRedeployDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		uploadSource: connect.NewClient[v1.UploadSourceRequest, v1.UploadSourceResponse](
			httpClient,
			baseURL+BuilderServiceUploadSourceProcedure,
			connect.WithSchema(builderServiceUploadSourceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDeploymentSteps: connect.NewClient[v1.GetDeploymentStepsRequest, v1.GetDeploymentStepsResponse](
			httpClient,
			baseURL+BuilderServiceGetDeploymentStepsProcedure,
			connect.WithSchema(builderServiceGetDeploymentStepsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDeploymentLogs: connect.NewClient[v1.GetDeploymentLogsRequest, v1.GetDeploymentLogsResponse](
			httpClient,
			baseURL+BuilderServiceGetDeploymentLogsProcedure,
			connect.WithSchema(builderServiceGetDeploymentLogsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getClusterTimescaleDB: connect.NewClient[v1.GetClusterTimescaleDBRequest, v1.GetClusterTimescaleDBResponse](
			httpClient,
			baseURL+BuilderServiceGetClusterTimescaleDBProcedure,
			connect.WithSchema(builderServiceGetClusterTimescaleDBMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getClusterGateway: connect.NewClient[v1.GetClusterGatewayRequest, v1.GetClusterGatewayResponse](
			httpClient,
			baseURL+BuilderServiceGetClusterGatewayProcedure,
			connect.WithSchema(builderServiceGetClusterGatewayMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createClusterTimescaleDB: connect.NewClient[v1.CreateClusterTimescaleDBRequest, v1.CreateClusterTimescaleDBResponse](
			httpClient,
			baseURL+BuilderServiceCreateClusterTimescaleDBProcedure,
			connect.WithSchema(builderServiceCreateClusterTimescaleDBMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createClusterGateway: connect.NewClient[v1.CreateClusterGatewayRequest, v1.CreateClusterGatewayResponse](
			httpClient,
			baseURL+BuilderServiceCreateClusterGatewayProcedure,
			connect.WithSchema(builderServiceCreateClusterGatewayMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// builderServiceClient implements BuilderServiceClient.
type builderServiceClient struct {
	activateDeployment       *connect.Client[v1.ActivateDeploymentRequest, v1.ActivateDeploymentResponse]
	indexDeployment          *connect.Client[v1.IndexDeploymentRequest, v1.IndexDeploymentResponse]
	deployKubeComponents     *connect.Client[v1.DeployKubeComponentsRequest, v1.DeployKubeComponentsResponse]
	rebuildDeployment        *connect.Client[v1.RebuildDeploymentRequest, v1.RebuildDeploymentResponse]
	redeployDeployment       *connect.Client[v1.RedeployDeploymentRequest, v1.RedeployDeploymentResponse]
	uploadSource             *connect.Client[v1.UploadSourceRequest, v1.UploadSourceResponse]
	getDeploymentSteps       *connect.Client[v1.GetDeploymentStepsRequest, v1.GetDeploymentStepsResponse]
	getDeploymentLogs        *connect.Client[v1.GetDeploymentLogsRequest, v1.GetDeploymentLogsResponse]
	getClusterTimescaleDB    *connect.Client[v1.GetClusterTimescaleDBRequest, v1.GetClusterTimescaleDBResponse]
	getClusterGateway        *connect.Client[v1.GetClusterGatewayRequest, v1.GetClusterGatewayResponse]
	createClusterTimescaleDB *connect.Client[v1.CreateClusterTimescaleDBRequest, v1.CreateClusterTimescaleDBResponse]
	createClusterGateway     *connect.Client[v1.CreateClusterGatewayRequest, v1.CreateClusterGatewayResponse]
}

// ActivateDeployment calls chalk.server.v1.BuilderService.ActivateDeployment.
func (c *builderServiceClient) ActivateDeployment(ctx context.Context, req *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error) {
	return c.activateDeployment.CallUnary(ctx, req)
}

// IndexDeployment calls chalk.server.v1.BuilderService.IndexDeployment.
func (c *builderServiceClient) IndexDeployment(ctx context.Context, req *connect.Request[v1.IndexDeploymentRequest]) (*connect.Response[v1.IndexDeploymentResponse], error) {
	return c.indexDeployment.CallUnary(ctx, req)
}

// DeployKubeComponents calls chalk.server.v1.BuilderService.DeployKubeComponents.
func (c *builderServiceClient) DeployKubeComponents(ctx context.Context, req *connect.Request[v1.DeployKubeComponentsRequest]) (*connect.Response[v1.DeployKubeComponentsResponse], error) {
	return c.deployKubeComponents.CallUnary(ctx, req)
}

// RebuildDeployment calls chalk.server.v1.BuilderService.RebuildDeployment.
func (c *builderServiceClient) RebuildDeployment(ctx context.Context, req *connect.Request[v1.RebuildDeploymentRequest]) (*connect.Response[v1.RebuildDeploymentResponse], error) {
	return c.rebuildDeployment.CallUnary(ctx, req)
}

// RedeployDeployment calls chalk.server.v1.BuilderService.RedeployDeployment.
func (c *builderServiceClient) RedeployDeployment(ctx context.Context, req *connect.Request[v1.RedeployDeploymentRequest]) (*connect.Response[v1.RedeployDeploymentResponse], error) {
	return c.redeployDeployment.CallUnary(ctx, req)
}

// UploadSource calls chalk.server.v1.BuilderService.UploadSource.
func (c *builderServiceClient) UploadSource(ctx context.Context, req *connect.Request[v1.UploadSourceRequest]) (*connect.Response[v1.UploadSourceResponse], error) {
	return c.uploadSource.CallUnary(ctx, req)
}

// GetDeploymentSteps calls chalk.server.v1.BuilderService.GetDeploymentSteps.
func (c *builderServiceClient) GetDeploymentSteps(ctx context.Context, req *connect.Request[v1.GetDeploymentStepsRequest]) (*connect.Response[v1.GetDeploymentStepsResponse], error) {
	return c.getDeploymentSteps.CallUnary(ctx, req)
}

// GetDeploymentLogs calls chalk.server.v1.BuilderService.GetDeploymentLogs.
func (c *builderServiceClient) GetDeploymentLogs(ctx context.Context, req *connect.Request[v1.GetDeploymentLogsRequest]) (*connect.Response[v1.GetDeploymentLogsResponse], error) {
	return c.getDeploymentLogs.CallUnary(ctx, req)
}

// GetClusterTimescaleDB calls chalk.server.v1.BuilderService.GetClusterTimescaleDB.
func (c *builderServiceClient) GetClusterTimescaleDB(ctx context.Context, req *connect.Request[v1.GetClusterTimescaleDBRequest]) (*connect.Response[v1.GetClusterTimescaleDBResponse], error) {
	return c.getClusterTimescaleDB.CallUnary(ctx, req)
}

// GetClusterGateway calls chalk.server.v1.BuilderService.GetClusterGateway.
func (c *builderServiceClient) GetClusterGateway(ctx context.Context, req *connect.Request[v1.GetClusterGatewayRequest]) (*connect.Response[v1.GetClusterGatewayResponse], error) {
	return c.getClusterGateway.CallUnary(ctx, req)
}

// CreateClusterTimescaleDB calls chalk.server.v1.BuilderService.CreateClusterTimescaleDB.
func (c *builderServiceClient) CreateClusterTimescaleDB(ctx context.Context, req *connect.Request[v1.CreateClusterTimescaleDBRequest]) (*connect.Response[v1.CreateClusterTimescaleDBResponse], error) {
	return c.createClusterTimescaleDB.CallUnary(ctx, req)
}

// CreateClusterGateway calls chalk.server.v1.BuilderService.CreateClusterGateway.
func (c *builderServiceClient) CreateClusterGateway(ctx context.Context, req *connect.Request[v1.CreateClusterGatewayRequest]) (*connect.Response[v1.CreateClusterGatewayResponse], error) {
	return c.createClusterGateway.CallUnary(ctx, req)
}

// BuilderServiceHandler is an implementation of the chalk.server.v1.BuilderService service.
type BuilderServiceHandler interface {
	// Takes an existing (past) deployment and promotes the k8s resources / other things associated with it.
	// Useful for debugging in local development where the auto activation doesn't work b/c no pubsub.
	ActivateDeployment(context.Context, *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error)
	IndexDeployment(context.Context, *connect.Request[v1.IndexDeploymentRequest]) (*connect.Response[v1.IndexDeploymentResponse], error)
	// Intermediate step in the deployment activation process. Allows for partial migration to the new
	// go-api-server builder service.
	DeployKubeComponents(context.Context, *connect.Request[v1.DeployKubeComponentsRequest]) (*connect.Response[v1.DeployKubeComponentsResponse], error)
	// Takes an existing (past) deployment and re-creates the image associated with it,
	// publishing the image as 'new_image_tag'.
	RebuildDeployment(context.Context, *connect.Request[v1.RebuildDeploymentRequest]) (*connect.Response[v1.RebuildDeploymentResponse], error)
	// Triggers a new build with the source code from this deployment and deploys the result
	RedeployDeployment(context.Context, *connect.Request[v1.RedeployDeploymentRequest]) (*connect.Response[v1.RedeployDeploymentResponse], error)
	// Triggers a new build with the provided source code archive and deploys the result
	UploadSource(context.Context, *connect.Request[v1.UploadSourceRequest]) (*connect.Response[v1.UploadSourceResponse], error)
	GetDeploymentSteps(context.Context, *connect.Request[v1.GetDeploymentStepsRequest]) (*connect.Response[v1.GetDeploymentStepsResponse], error)
	GetDeploymentLogs(context.Context, *connect.Request[v1.GetDeploymentLogsRequest]) (*connect.Response[v1.GetDeploymentLogsResponse], error)
	GetClusterTimescaleDB(context.Context, *connect.Request[v1.GetClusterTimescaleDBRequest]) (*connect.Response[v1.GetClusterTimescaleDBResponse], error)
	GetClusterGateway(context.Context, *connect.Request[v1.GetClusterGatewayRequest]) (*connect.Response[v1.GetClusterGatewayResponse], error)
	CreateClusterTimescaleDB(context.Context, *connect.Request[v1.CreateClusterTimescaleDBRequest]) (*connect.Response[v1.CreateClusterTimescaleDBResponse], error)
	CreateClusterGateway(context.Context, *connect.Request[v1.CreateClusterGatewayRequest]) (*connect.Response[v1.CreateClusterGatewayResponse], error)
}

// NewBuilderServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBuilderServiceHandler(svc BuilderServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	builderServiceActivateDeploymentHandler := connect.NewUnaryHandler(
		BuilderServiceActivateDeploymentProcedure,
		svc.ActivateDeployment,
		connect.WithSchema(builderServiceActivateDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceIndexDeploymentHandler := connect.NewUnaryHandler(
		BuilderServiceIndexDeploymentProcedure,
		svc.IndexDeployment,
		connect.WithSchema(builderServiceIndexDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceDeployKubeComponentsHandler := connect.NewUnaryHandler(
		BuilderServiceDeployKubeComponentsProcedure,
		svc.DeployKubeComponents,
		connect.WithSchema(builderServiceDeployKubeComponentsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceRebuildDeploymentHandler := connect.NewUnaryHandler(
		BuilderServiceRebuildDeploymentProcedure,
		svc.RebuildDeployment,
		connect.WithSchema(builderServiceRebuildDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceRedeployDeploymentHandler := connect.NewUnaryHandler(
		BuilderServiceRedeployDeploymentProcedure,
		svc.RedeployDeployment,
		connect.WithSchema(builderServiceRedeployDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceUploadSourceHandler := connect.NewUnaryHandler(
		BuilderServiceUploadSourceProcedure,
		svc.UploadSource,
		connect.WithSchema(builderServiceUploadSourceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceGetDeploymentStepsHandler := connect.NewUnaryHandler(
		BuilderServiceGetDeploymentStepsProcedure,
		svc.GetDeploymentSteps,
		connect.WithSchema(builderServiceGetDeploymentStepsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceGetDeploymentLogsHandler := connect.NewUnaryHandler(
		BuilderServiceGetDeploymentLogsProcedure,
		svc.GetDeploymentLogs,
		connect.WithSchema(builderServiceGetDeploymentLogsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceGetClusterTimescaleDBHandler := connect.NewUnaryHandler(
		BuilderServiceGetClusterTimescaleDBProcedure,
		svc.GetClusterTimescaleDB,
		connect.WithSchema(builderServiceGetClusterTimescaleDBMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceGetClusterGatewayHandler := connect.NewUnaryHandler(
		BuilderServiceGetClusterGatewayProcedure,
		svc.GetClusterGateway,
		connect.WithSchema(builderServiceGetClusterGatewayMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceCreateClusterTimescaleDBHandler := connect.NewUnaryHandler(
		BuilderServiceCreateClusterTimescaleDBProcedure,
		svc.CreateClusterTimescaleDB,
		connect.WithSchema(builderServiceCreateClusterTimescaleDBMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	builderServiceCreateClusterGatewayHandler := connect.NewUnaryHandler(
		BuilderServiceCreateClusterGatewayProcedure,
		svc.CreateClusterGateway,
		connect.WithSchema(builderServiceCreateClusterGatewayMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.BuilderService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BuilderServiceActivateDeploymentProcedure:
			builderServiceActivateDeploymentHandler.ServeHTTP(w, r)
		case BuilderServiceIndexDeploymentProcedure:
			builderServiceIndexDeploymentHandler.ServeHTTP(w, r)
		case BuilderServiceDeployKubeComponentsProcedure:
			builderServiceDeployKubeComponentsHandler.ServeHTTP(w, r)
		case BuilderServiceRebuildDeploymentProcedure:
			builderServiceRebuildDeploymentHandler.ServeHTTP(w, r)
		case BuilderServiceRedeployDeploymentProcedure:
			builderServiceRedeployDeploymentHandler.ServeHTTP(w, r)
		case BuilderServiceUploadSourceProcedure:
			builderServiceUploadSourceHandler.ServeHTTP(w, r)
		case BuilderServiceGetDeploymentStepsProcedure:
			builderServiceGetDeploymentStepsHandler.ServeHTTP(w, r)
		case BuilderServiceGetDeploymentLogsProcedure:
			builderServiceGetDeploymentLogsHandler.ServeHTTP(w, r)
		case BuilderServiceGetClusterTimescaleDBProcedure:
			builderServiceGetClusterTimescaleDBHandler.ServeHTTP(w, r)
		case BuilderServiceGetClusterGatewayProcedure:
			builderServiceGetClusterGatewayHandler.ServeHTTP(w, r)
		case BuilderServiceCreateClusterTimescaleDBProcedure:
			builderServiceCreateClusterTimescaleDBHandler.ServeHTTP(w, r)
		case BuilderServiceCreateClusterGatewayProcedure:
			builderServiceCreateClusterGatewayHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBuilderServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBuilderServiceHandler struct{}

func (UnimplementedBuilderServiceHandler) ActivateDeployment(context.Context, *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.ActivateDeployment is not implemented"))
}

func (UnimplementedBuilderServiceHandler) IndexDeployment(context.Context, *connect.Request[v1.IndexDeploymentRequest]) (*connect.Response[v1.IndexDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.IndexDeployment is not implemented"))
}

func (UnimplementedBuilderServiceHandler) DeployKubeComponents(context.Context, *connect.Request[v1.DeployKubeComponentsRequest]) (*connect.Response[v1.DeployKubeComponentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.DeployKubeComponents is not implemented"))
}

func (UnimplementedBuilderServiceHandler) RebuildDeployment(context.Context, *connect.Request[v1.RebuildDeploymentRequest]) (*connect.Response[v1.RebuildDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.RebuildDeployment is not implemented"))
}

func (UnimplementedBuilderServiceHandler) RedeployDeployment(context.Context, *connect.Request[v1.RedeployDeploymentRequest]) (*connect.Response[v1.RedeployDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.RedeployDeployment is not implemented"))
}

func (UnimplementedBuilderServiceHandler) UploadSource(context.Context, *connect.Request[v1.UploadSourceRequest]) (*connect.Response[v1.UploadSourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.UploadSource is not implemented"))
}

func (UnimplementedBuilderServiceHandler) GetDeploymentSteps(context.Context, *connect.Request[v1.GetDeploymentStepsRequest]) (*connect.Response[v1.GetDeploymentStepsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.GetDeploymentSteps is not implemented"))
}

func (UnimplementedBuilderServiceHandler) GetDeploymentLogs(context.Context, *connect.Request[v1.GetDeploymentLogsRequest]) (*connect.Response[v1.GetDeploymentLogsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.GetDeploymentLogs is not implemented"))
}

func (UnimplementedBuilderServiceHandler) GetClusterTimescaleDB(context.Context, *connect.Request[v1.GetClusterTimescaleDBRequest]) (*connect.Response[v1.GetClusterTimescaleDBResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.GetClusterTimescaleDB is not implemented"))
}

func (UnimplementedBuilderServiceHandler) GetClusterGateway(context.Context, *connect.Request[v1.GetClusterGatewayRequest]) (*connect.Response[v1.GetClusterGatewayResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.GetClusterGateway is not implemented"))
}

func (UnimplementedBuilderServiceHandler) CreateClusterTimescaleDB(context.Context, *connect.Request[v1.CreateClusterTimescaleDBRequest]) (*connect.Response[v1.CreateClusterTimescaleDBResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.CreateClusterTimescaleDB is not implemented"))
}

func (UnimplementedBuilderServiceHandler) CreateClusterGateway(context.Context, *connect.Request[v1.CreateClusterGatewayRequest]) (*connect.Response[v1.CreateClusterGatewayResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.BuilderService.CreateClusterGateway is not implemented"))
}
