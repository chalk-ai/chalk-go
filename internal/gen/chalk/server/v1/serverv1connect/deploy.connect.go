// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/deploy.proto

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
	// DeployServiceName is the fully-qualified name of the DeployService service.
	DeployServiceName = "chalk.server.v1.DeployService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DeployServiceDeployBranchProcedure is the fully-qualified name of the DeployService's
	// DeployBranch RPC.
	DeployServiceDeployBranchProcedure = "/chalk.server.v1.DeployService/DeployBranch"
	// DeployServiceGetDeploymentProcedure is the fully-qualified name of the DeployService's
	// GetDeployment RPC.
	DeployServiceGetDeploymentProcedure = "/chalk.server.v1.DeployService/GetDeployment"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	deployServiceServiceDescriptor             = v1.File_chalk_server_v1_deploy_proto.Services().ByName("DeployService")
	deployServiceDeployBranchMethodDescriptor  = deployServiceServiceDescriptor.Methods().ByName("DeployBranch")
	deployServiceGetDeploymentMethodDescriptor = deployServiceServiceDescriptor.Methods().ByName("GetDeployment")
)

// DeployServiceClient is a client for the chalk.server.v1.DeployService service.
type DeployServiceClient interface {
	DeployBranch(context.Context, *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error)
	GetDeployment(context.Context, *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error)
}

// NewDeployServiceClient constructs a client for the chalk.server.v1.DeployService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDeployServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DeployServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &deployServiceClient{
		deployBranch: connect.NewClient[v1.DeployBranchRequest, v1.DeployBranchResponse](
			httpClient,
			baseURL+DeployServiceDeployBranchProcedure,
			connect.WithSchema(deployServiceDeployBranchMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDeployment: connect.NewClient[v1.GetDeploymentRequest, v1.GetDeploymentResponse](
			httpClient,
			baseURL+DeployServiceGetDeploymentProcedure,
			connect.WithSchema(deployServiceGetDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// deployServiceClient implements DeployServiceClient.
type deployServiceClient struct {
	deployBranch  *connect.Client[v1.DeployBranchRequest, v1.DeployBranchResponse]
	getDeployment *connect.Client[v1.GetDeploymentRequest, v1.GetDeploymentResponse]
}

// DeployBranch calls chalk.server.v1.DeployService.DeployBranch.
func (c *deployServiceClient) DeployBranch(ctx context.Context, req *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error) {
	return c.deployBranch.CallUnary(ctx, req)
}

// GetDeployment calls chalk.server.v1.DeployService.GetDeployment.
func (c *deployServiceClient) GetDeployment(ctx context.Context, req *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error) {
	return c.getDeployment.CallUnary(ctx, req)
}

// DeployServiceHandler is an implementation of the chalk.server.v1.DeployService service.
type DeployServiceHandler interface {
	DeployBranch(context.Context, *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error)
	GetDeployment(context.Context, *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error)
}

// NewDeployServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDeployServiceHandler(svc DeployServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	deployServiceDeployBranchHandler := connect.NewUnaryHandler(
		DeployServiceDeployBranchProcedure,
		svc.DeployBranch,
		connect.WithSchema(deployServiceDeployBranchMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceGetDeploymentHandler := connect.NewUnaryHandler(
		DeployServiceGetDeploymentProcedure,
		svc.GetDeployment,
		connect.WithSchema(deployServiceGetDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.DeployService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DeployServiceDeployBranchProcedure:
			deployServiceDeployBranchHandler.ServeHTTP(w, r)
		case DeployServiceGetDeploymentProcedure:
			deployServiceGetDeploymentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDeployServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDeployServiceHandler struct{}

func (UnimplementedDeployServiceHandler) DeployBranch(context.Context, *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.DeployBranch is not implemented"))
}

func (UnimplementedDeployServiceHandler) GetDeployment(context.Context, *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.GetDeployment is not implemented"))
}
