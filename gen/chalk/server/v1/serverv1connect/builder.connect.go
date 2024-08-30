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
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	builderServiceServiceDescriptor                  = v1.File_chalk_server_v1_builder_proto.Services().ByName("BuilderService")
	builderServiceActivateDeploymentMethodDescriptor = builderServiceServiceDescriptor.Methods().ByName("ActivateDeployment")
)

// BuilderServiceClient is a client for the chalk.server.v1.BuilderService service.
type BuilderServiceClient interface {
	// Takes an existing (past) deployment and promotes the k8s resources / other things associated with it.
	// Useful for debugging in local development where the auto activation doesn't work b/c no pubsub.
	ActivateDeployment(context.Context, *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error)
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
	}
}

// builderServiceClient implements BuilderServiceClient.
type builderServiceClient struct {
	activateDeployment *connect.Client[v1.ActivateDeploymentRequest, v1.ActivateDeploymentResponse]
}

// ActivateDeployment calls chalk.server.v1.BuilderService.ActivateDeployment.
func (c *builderServiceClient) ActivateDeployment(ctx context.Context, req *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error) {
	return c.activateDeployment.CallUnary(ctx, req)
}

// BuilderServiceHandler is an implementation of the chalk.server.v1.BuilderService service.
type BuilderServiceHandler interface {
	// Takes an existing (past) deployment and promotes the k8s resources / other things associated with it.
	// Useful for debugging in local development where the auto activation doesn't work b/c no pubsub.
	ActivateDeployment(context.Context, *connect.Request[v1.ActivateDeploymentRequest]) (*connect.Response[v1.ActivateDeploymentResponse], error)
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
	return "/chalk.server.v1.BuilderService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BuilderServiceActivateDeploymentProcedure:
			builderServiceActivateDeploymentHandler.ServeHTTP(w, r)
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
