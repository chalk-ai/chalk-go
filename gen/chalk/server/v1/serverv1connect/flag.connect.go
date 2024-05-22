// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/flag.proto

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
	// FeatureFlagServiceName is the fully-qualified name of the FeatureFlagService service.
	FeatureFlagServiceName = "chalk.server.v1.FeatureFlagService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FeatureFlagServiceGetFeatureFlagsProcedure is the fully-qualified name of the
	// FeatureFlagService's GetFeatureFlags RPC.
	FeatureFlagServiceGetFeatureFlagsProcedure = "/chalk.server.v1.FeatureFlagService/GetFeatureFlags"
	// FeatureFlagServiceSetFeatureFlagProcedure is the fully-qualified name of the FeatureFlagService's
	// SetFeatureFlag RPC.
	FeatureFlagServiceSetFeatureFlagProcedure = "/chalk.server.v1.FeatureFlagService/SetFeatureFlag"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	featureFlagServiceServiceDescriptor               = v1.File_chalk_server_v1_flag_proto.Services().ByName("FeatureFlagService")
	featureFlagServiceGetFeatureFlagsMethodDescriptor = featureFlagServiceServiceDescriptor.Methods().ByName("GetFeatureFlags")
	featureFlagServiceSetFeatureFlagMethodDescriptor  = featureFlagServiceServiceDescriptor.Methods().ByName("SetFeatureFlag")
)

// FeatureFlagServiceClient is a client for the chalk.server.v1.FeatureFlagService service.
type FeatureFlagServiceClient interface {
	GetFeatureFlags(context.Context, *connect.Request[v1.GetFeatureFlagsRequest]) (*connect.Response[v1.GetFeatureFlagsResponse], error)
	SetFeatureFlag(context.Context, *connect.Request[v1.SetFeatureFlagRequest]) (*connect.Response[v1.SetFeatureFlagResponse], error)
}

// NewFeatureFlagServiceClient constructs a client for the chalk.server.v1.FeatureFlagService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFeatureFlagServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FeatureFlagServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &featureFlagServiceClient{
		getFeatureFlags: connect.NewClient[v1.GetFeatureFlagsRequest, v1.GetFeatureFlagsResponse](
			httpClient,
			baseURL+FeatureFlagServiceGetFeatureFlagsProcedure,
			connect.WithSchema(featureFlagServiceGetFeatureFlagsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		setFeatureFlag: connect.NewClient[v1.SetFeatureFlagRequest, v1.SetFeatureFlagResponse](
			httpClient,
			baseURL+FeatureFlagServiceSetFeatureFlagProcedure,
			connect.WithSchema(featureFlagServiceSetFeatureFlagMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
	}
}

// featureFlagServiceClient implements FeatureFlagServiceClient.
type featureFlagServiceClient struct {
	getFeatureFlags *connect.Client[v1.GetFeatureFlagsRequest, v1.GetFeatureFlagsResponse]
	setFeatureFlag  *connect.Client[v1.SetFeatureFlagRequest, v1.SetFeatureFlagResponse]
}

// GetFeatureFlags calls chalk.server.v1.FeatureFlagService.GetFeatureFlags.
func (c *featureFlagServiceClient) GetFeatureFlags(ctx context.Context, req *connect.Request[v1.GetFeatureFlagsRequest]) (*connect.Response[v1.GetFeatureFlagsResponse], error) {
	return c.getFeatureFlags.CallUnary(ctx, req)
}

// SetFeatureFlag calls chalk.server.v1.FeatureFlagService.SetFeatureFlag.
func (c *featureFlagServiceClient) SetFeatureFlag(ctx context.Context, req *connect.Request[v1.SetFeatureFlagRequest]) (*connect.Response[v1.SetFeatureFlagResponse], error) {
	return c.setFeatureFlag.CallUnary(ctx, req)
}

// FeatureFlagServiceHandler is an implementation of the chalk.server.v1.FeatureFlagService service.
type FeatureFlagServiceHandler interface {
	GetFeatureFlags(context.Context, *connect.Request[v1.GetFeatureFlagsRequest]) (*connect.Response[v1.GetFeatureFlagsResponse], error)
	SetFeatureFlag(context.Context, *connect.Request[v1.SetFeatureFlagRequest]) (*connect.Response[v1.SetFeatureFlagResponse], error)
}

// NewFeatureFlagServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFeatureFlagServiceHandler(svc FeatureFlagServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	featureFlagServiceGetFeatureFlagsHandler := connect.NewUnaryHandler(
		FeatureFlagServiceGetFeatureFlagsProcedure,
		svc.GetFeatureFlags,
		connect.WithSchema(featureFlagServiceGetFeatureFlagsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	featureFlagServiceSetFeatureFlagHandler := connect.NewUnaryHandler(
		FeatureFlagServiceSetFeatureFlagProcedure,
		svc.SetFeatureFlag,
		connect.WithSchema(featureFlagServiceSetFeatureFlagMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.FeatureFlagService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FeatureFlagServiceGetFeatureFlagsProcedure:
			featureFlagServiceGetFeatureFlagsHandler.ServeHTTP(w, r)
		case FeatureFlagServiceSetFeatureFlagProcedure:
			featureFlagServiceSetFeatureFlagHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFeatureFlagServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFeatureFlagServiceHandler struct{}

func (UnimplementedFeatureFlagServiceHandler) GetFeatureFlags(context.Context, *connect.Request[v1.GetFeatureFlagsRequest]) (*connect.Response[v1.GetFeatureFlagsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.FeatureFlagService.GetFeatureFlags is not implemented"))
}

func (UnimplementedFeatureFlagServiceHandler) SetFeatureFlag(context.Context, *connect.Request[v1.SetFeatureFlagRequest]) (*connect.Response[v1.SetFeatureFlagResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.FeatureFlagService.SetFeatureFlag is not implemented"))
}