// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/deploy.proto

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
	// DeployServiceListDeploymentsProcedure is the fully-qualified name of the DeployService's
	// ListDeployments RPC.
	DeployServiceListDeploymentsProcedure = "/chalk.server.v1.DeployService/ListDeployments"
	// DeployServiceGetActiveDeploymentsProcedure is the fully-qualified name of the DeployService's
	// GetActiveDeployments RPC.
	DeployServiceGetActiveDeploymentsProcedure = "/chalk.server.v1.DeployService/GetActiveDeployments"
	// DeployServiceSuspendDeploymentProcedure is the fully-qualified name of the DeployService's
	// SuspendDeployment RPC.
	DeployServiceSuspendDeploymentProcedure = "/chalk.server.v1.DeployService/SuspendDeployment"
	// DeployServiceScaleDeploymentProcedure is the fully-qualified name of the DeployService's
	// ScaleDeployment RPC.
	DeployServiceScaleDeploymentProcedure = "/chalk.server.v1.DeployService/ScaleDeployment"
	// DeployServiceTagDeploymentProcedure is the fully-qualified name of the DeployService's
	// TagDeployment RPC.
	DeployServiceTagDeploymentProcedure = "/chalk.server.v1.DeployService/TagDeployment"
)

// DeployServiceClient is a client for the chalk.server.v1.DeployService service.
type DeployServiceClient interface {
	DeployBranch(context.Context, *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error)
	GetDeployment(context.Context, *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error)
	ListDeployments(context.Context, *connect.Request[v1.ListDeploymentsRequest]) (*connect.Response[v1.ListDeploymentsResponse], error)
	GetActiveDeployments(context.Context, *connect.Request[v1.GetActiveDeploymentsRequest]) (*connect.Response[v1.GetActiveDeploymentsResponse], error)
	SuspendDeployment(context.Context, *connect.Request[v1.SuspendDeploymentRequest]) (*connect.Response[v1.SuspendDeploymentResponse], error)
	ScaleDeployment(context.Context, *connect.Request[v1.ScaleDeploymentRequest]) (*connect.Response[v1.ScaleDeploymentResponse], error)
	TagDeployment(context.Context, *connect.Request[v1.TagDeploymentRequest]) (*connect.Response[v1.TagDeploymentResponse], error)
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
	deployServiceMethods := v1.File_chalk_server_v1_deploy_proto.Services().ByName("DeployService").Methods()
	return &deployServiceClient{
		deployBranch: connect.NewClient[v1.DeployBranchRequest, v1.DeployBranchResponse](
			httpClient,
			baseURL+DeployServiceDeployBranchProcedure,
			connect.WithSchema(deployServiceMethods.ByName("DeployBranch")),
			connect.WithClientOptions(opts...),
		),
		getDeployment: connect.NewClient[v1.GetDeploymentRequest, v1.GetDeploymentResponse](
			httpClient,
			baseURL+DeployServiceGetDeploymentProcedure,
			connect.WithSchema(deployServiceMethods.ByName("GetDeployment")),
			connect.WithClientOptions(opts...),
		),
		listDeployments: connect.NewClient[v1.ListDeploymentsRequest, v1.ListDeploymentsResponse](
			httpClient,
			baseURL+DeployServiceListDeploymentsProcedure,
			connect.WithSchema(deployServiceMethods.ByName("ListDeployments")),
			connect.WithClientOptions(opts...),
		),
		getActiveDeployments: connect.NewClient[v1.GetActiveDeploymentsRequest, v1.GetActiveDeploymentsResponse](
			httpClient,
			baseURL+DeployServiceGetActiveDeploymentsProcedure,
			connect.WithSchema(deployServiceMethods.ByName("GetActiveDeployments")),
			connect.WithClientOptions(opts...),
		),
		suspendDeployment: connect.NewClient[v1.SuspendDeploymentRequest, v1.SuspendDeploymentResponse](
			httpClient,
			baseURL+DeployServiceSuspendDeploymentProcedure,
			connect.WithSchema(deployServiceMethods.ByName("SuspendDeployment")),
			connect.WithClientOptions(opts...),
		),
		scaleDeployment: connect.NewClient[v1.ScaleDeploymentRequest, v1.ScaleDeploymentResponse](
			httpClient,
			baseURL+DeployServiceScaleDeploymentProcedure,
			connect.WithSchema(deployServiceMethods.ByName("ScaleDeployment")),
			connect.WithClientOptions(opts...),
		),
		tagDeployment: connect.NewClient[v1.TagDeploymentRequest, v1.TagDeploymentResponse](
			httpClient,
			baseURL+DeployServiceTagDeploymentProcedure,
			connect.WithSchema(deployServiceMethods.ByName("TagDeployment")),
			connect.WithClientOptions(opts...),
		),
	}
}

// deployServiceClient implements DeployServiceClient.
type deployServiceClient struct {
	deployBranch         *connect.Client[v1.DeployBranchRequest, v1.DeployBranchResponse]
	getDeployment        *connect.Client[v1.GetDeploymentRequest, v1.GetDeploymentResponse]
	listDeployments      *connect.Client[v1.ListDeploymentsRequest, v1.ListDeploymentsResponse]
	getActiveDeployments *connect.Client[v1.GetActiveDeploymentsRequest, v1.GetActiveDeploymentsResponse]
	suspendDeployment    *connect.Client[v1.SuspendDeploymentRequest, v1.SuspendDeploymentResponse]
	scaleDeployment      *connect.Client[v1.ScaleDeploymentRequest, v1.ScaleDeploymentResponse]
	tagDeployment        *connect.Client[v1.TagDeploymentRequest, v1.TagDeploymentResponse]
}

// DeployBranch calls chalk.server.v1.DeployService.DeployBranch.
func (c *deployServiceClient) DeployBranch(ctx context.Context, req *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error) {
	return c.deployBranch.CallUnary(ctx, req)
}

// GetDeployment calls chalk.server.v1.DeployService.GetDeployment.
func (c *deployServiceClient) GetDeployment(ctx context.Context, req *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error) {
	return c.getDeployment.CallUnary(ctx, req)
}

// ListDeployments calls chalk.server.v1.DeployService.ListDeployments.
func (c *deployServiceClient) ListDeployments(ctx context.Context, req *connect.Request[v1.ListDeploymentsRequest]) (*connect.Response[v1.ListDeploymentsResponse], error) {
	return c.listDeployments.CallUnary(ctx, req)
}

// GetActiveDeployments calls chalk.server.v1.DeployService.GetActiveDeployments.
func (c *deployServiceClient) GetActiveDeployments(ctx context.Context, req *connect.Request[v1.GetActiveDeploymentsRequest]) (*connect.Response[v1.GetActiveDeploymentsResponse], error) {
	return c.getActiveDeployments.CallUnary(ctx, req)
}

// SuspendDeployment calls chalk.server.v1.DeployService.SuspendDeployment.
func (c *deployServiceClient) SuspendDeployment(ctx context.Context, req *connect.Request[v1.SuspendDeploymentRequest]) (*connect.Response[v1.SuspendDeploymentResponse], error) {
	return c.suspendDeployment.CallUnary(ctx, req)
}

// ScaleDeployment calls chalk.server.v1.DeployService.ScaleDeployment.
func (c *deployServiceClient) ScaleDeployment(ctx context.Context, req *connect.Request[v1.ScaleDeploymentRequest]) (*connect.Response[v1.ScaleDeploymentResponse], error) {
	return c.scaleDeployment.CallUnary(ctx, req)
}

// TagDeployment calls chalk.server.v1.DeployService.TagDeployment.
func (c *deployServiceClient) TagDeployment(ctx context.Context, req *connect.Request[v1.TagDeploymentRequest]) (*connect.Response[v1.TagDeploymentResponse], error) {
	return c.tagDeployment.CallUnary(ctx, req)
}

// DeployServiceHandler is an implementation of the chalk.server.v1.DeployService service.
type DeployServiceHandler interface {
	DeployBranch(context.Context, *connect.Request[v1.DeployBranchRequest]) (*connect.Response[v1.DeployBranchResponse], error)
	GetDeployment(context.Context, *connect.Request[v1.GetDeploymentRequest]) (*connect.Response[v1.GetDeploymentResponse], error)
	ListDeployments(context.Context, *connect.Request[v1.ListDeploymentsRequest]) (*connect.Response[v1.ListDeploymentsResponse], error)
	GetActiveDeployments(context.Context, *connect.Request[v1.GetActiveDeploymentsRequest]) (*connect.Response[v1.GetActiveDeploymentsResponse], error)
	SuspendDeployment(context.Context, *connect.Request[v1.SuspendDeploymentRequest]) (*connect.Response[v1.SuspendDeploymentResponse], error)
	ScaleDeployment(context.Context, *connect.Request[v1.ScaleDeploymentRequest]) (*connect.Response[v1.ScaleDeploymentResponse], error)
	TagDeployment(context.Context, *connect.Request[v1.TagDeploymentRequest]) (*connect.Response[v1.TagDeploymentResponse], error)
}

// NewDeployServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDeployServiceHandler(svc DeployServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	deployServiceMethods := v1.File_chalk_server_v1_deploy_proto.Services().ByName("DeployService").Methods()
	deployServiceDeployBranchHandler := connect.NewUnaryHandler(
		DeployServiceDeployBranchProcedure,
		svc.DeployBranch,
		connect.WithSchema(deployServiceMethods.ByName("DeployBranch")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceGetDeploymentHandler := connect.NewUnaryHandler(
		DeployServiceGetDeploymentProcedure,
		svc.GetDeployment,
		connect.WithSchema(deployServiceMethods.ByName("GetDeployment")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceListDeploymentsHandler := connect.NewUnaryHandler(
		DeployServiceListDeploymentsProcedure,
		svc.ListDeployments,
		connect.WithSchema(deployServiceMethods.ByName("ListDeployments")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceGetActiveDeploymentsHandler := connect.NewUnaryHandler(
		DeployServiceGetActiveDeploymentsProcedure,
		svc.GetActiveDeployments,
		connect.WithSchema(deployServiceMethods.ByName("GetActiveDeployments")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceSuspendDeploymentHandler := connect.NewUnaryHandler(
		DeployServiceSuspendDeploymentProcedure,
		svc.SuspendDeployment,
		connect.WithSchema(deployServiceMethods.ByName("SuspendDeployment")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceScaleDeploymentHandler := connect.NewUnaryHandler(
		DeployServiceScaleDeploymentProcedure,
		svc.ScaleDeployment,
		connect.WithSchema(deployServiceMethods.ByName("ScaleDeployment")),
		connect.WithHandlerOptions(opts...),
	)
	deployServiceTagDeploymentHandler := connect.NewUnaryHandler(
		DeployServiceTagDeploymentProcedure,
		svc.TagDeployment,
		connect.WithSchema(deployServiceMethods.ByName("TagDeployment")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.DeployService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DeployServiceDeployBranchProcedure:
			deployServiceDeployBranchHandler.ServeHTTP(w, r)
		case DeployServiceGetDeploymentProcedure:
			deployServiceGetDeploymentHandler.ServeHTTP(w, r)
		case DeployServiceListDeploymentsProcedure:
			deployServiceListDeploymentsHandler.ServeHTTP(w, r)
		case DeployServiceGetActiveDeploymentsProcedure:
			deployServiceGetActiveDeploymentsHandler.ServeHTTP(w, r)
		case DeployServiceSuspendDeploymentProcedure:
			deployServiceSuspendDeploymentHandler.ServeHTTP(w, r)
		case DeployServiceScaleDeploymentProcedure:
			deployServiceScaleDeploymentHandler.ServeHTTP(w, r)
		case DeployServiceTagDeploymentProcedure:
			deployServiceTagDeploymentHandler.ServeHTTP(w, r)
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

func (UnimplementedDeployServiceHandler) ListDeployments(context.Context, *connect.Request[v1.ListDeploymentsRequest]) (*connect.Response[v1.ListDeploymentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.ListDeployments is not implemented"))
}

func (UnimplementedDeployServiceHandler) GetActiveDeployments(context.Context, *connect.Request[v1.GetActiveDeploymentsRequest]) (*connect.Response[v1.GetActiveDeploymentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.GetActiveDeployments is not implemented"))
}

func (UnimplementedDeployServiceHandler) SuspendDeployment(context.Context, *connect.Request[v1.SuspendDeploymentRequest]) (*connect.Response[v1.SuspendDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.SuspendDeployment is not implemented"))
}

func (UnimplementedDeployServiceHandler) ScaleDeployment(context.Context, *connect.Request[v1.ScaleDeploymentRequest]) (*connect.Response[v1.ScaleDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.ScaleDeployment is not implemented"))
}

func (UnimplementedDeployServiceHandler) TagDeployment(context.Context, *connect.Request[v1.TagDeploymentRequest]) (*connect.Response[v1.TagDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.DeployService.TagDeployment is not implemented"))
}
