// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/engine/v1/offline_store_service.proto

package enginev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/engine/v1"
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
	// OfflineStoreServiceName is the fully-qualified name of the OfflineStoreService service.
	OfflineStoreServiceName = "chalk.engine.v1.OfflineStoreService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OfflineStoreServiceGetQueryLogEntriesProcedure is the fully-qualified name of the
	// OfflineStoreService's GetQueryLogEntries RPC.
	OfflineStoreServiceGetQueryLogEntriesProcedure = "/chalk.engine.v1.OfflineStoreService/GetQueryLogEntries"
	// OfflineStoreServiceGetQueryValuesProcedure is the fully-qualified name of the
	// OfflineStoreService's GetQueryValues RPC.
	OfflineStoreServiceGetQueryValuesProcedure = "/chalk.engine.v1.OfflineStoreService/GetQueryValues"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	offlineStoreServiceServiceDescriptor                  = v1.File_chalk_engine_v1_offline_store_service_proto.Services().ByName("OfflineStoreService")
	offlineStoreServiceGetQueryLogEntriesMethodDescriptor = offlineStoreServiceServiceDescriptor.Methods().ByName("GetQueryLogEntries")
	offlineStoreServiceGetQueryValuesMethodDescriptor     = offlineStoreServiceServiceDescriptor.Methods().ByName("GetQueryValues")
)

// OfflineStoreServiceClient is a client for the chalk.engine.v1.OfflineStoreService service.
type OfflineStoreServiceClient interface {
	GetQueryLogEntries(context.Context, *connect.Request[v11.GetQueryLogEntriesRequest]) (*connect.Response[v11.GetQueryLogEntriesResponse], error)
	GetQueryValues(context.Context, *connect.Request[v11.GetQueryValuesRequest]) (*connect.Response[v11.GetQueryValuesResponse], error)
}

// NewOfflineStoreServiceClient constructs a client for the chalk.engine.v1.OfflineStoreService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOfflineStoreServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OfflineStoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &offlineStoreServiceClient{
		getQueryLogEntries: connect.NewClient[v11.GetQueryLogEntriesRequest, v11.GetQueryLogEntriesResponse](
			httpClient,
			baseURL+OfflineStoreServiceGetQueryLogEntriesProcedure,
			connect.WithSchema(offlineStoreServiceGetQueryLogEntriesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getQueryValues: connect.NewClient[v11.GetQueryValuesRequest, v11.GetQueryValuesResponse](
			httpClient,
			baseURL+OfflineStoreServiceGetQueryValuesProcedure,
			connect.WithSchema(offlineStoreServiceGetQueryValuesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// offlineStoreServiceClient implements OfflineStoreServiceClient.
type offlineStoreServiceClient struct {
	getQueryLogEntries *connect.Client[v11.GetQueryLogEntriesRequest, v11.GetQueryLogEntriesResponse]
	getQueryValues     *connect.Client[v11.GetQueryValuesRequest, v11.GetQueryValuesResponse]
}

// GetQueryLogEntries calls chalk.engine.v1.OfflineStoreService.GetQueryLogEntries.
func (c *offlineStoreServiceClient) GetQueryLogEntries(ctx context.Context, req *connect.Request[v11.GetQueryLogEntriesRequest]) (*connect.Response[v11.GetQueryLogEntriesResponse], error) {
	return c.getQueryLogEntries.CallUnary(ctx, req)
}

// GetQueryValues calls chalk.engine.v1.OfflineStoreService.GetQueryValues.
func (c *offlineStoreServiceClient) GetQueryValues(ctx context.Context, req *connect.Request[v11.GetQueryValuesRequest]) (*connect.Response[v11.GetQueryValuesResponse], error) {
	return c.getQueryValues.CallUnary(ctx, req)
}

// OfflineStoreServiceHandler is an implementation of the chalk.engine.v1.OfflineStoreService
// service.
type OfflineStoreServiceHandler interface {
	GetQueryLogEntries(context.Context, *connect.Request[v11.GetQueryLogEntriesRequest]) (*connect.Response[v11.GetQueryLogEntriesResponse], error)
	GetQueryValues(context.Context, *connect.Request[v11.GetQueryValuesRequest]) (*connect.Response[v11.GetQueryValuesResponse], error)
}

// NewOfflineStoreServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOfflineStoreServiceHandler(svc OfflineStoreServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	offlineStoreServiceGetQueryLogEntriesHandler := connect.NewUnaryHandler(
		OfflineStoreServiceGetQueryLogEntriesProcedure,
		svc.GetQueryLogEntries,
		connect.WithSchema(offlineStoreServiceGetQueryLogEntriesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	offlineStoreServiceGetQueryValuesHandler := connect.NewUnaryHandler(
		OfflineStoreServiceGetQueryValuesProcedure,
		svc.GetQueryValues,
		connect.WithSchema(offlineStoreServiceGetQueryValuesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.engine.v1.OfflineStoreService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OfflineStoreServiceGetQueryLogEntriesProcedure:
			offlineStoreServiceGetQueryLogEntriesHandler.ServeHTTP(w, r)
		case OfflineStoreServiceGetQueryValuesProcedure:
			offlineStoreServiceGetQueryValuesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOfflineStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOfflineStoreServiceHandler struct{}

func (UnimplementedOfflineStoreServiceHandler) GetQueryLogEntries(context.Context, *connect.Request[v11.GetQueryLogEntriesRequest]) (*connect.Response[v11.GetQueryLogEntriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.OfflineStoreService.GetQueryLogEntries is not implemented"))
}

func (UnimplementedOfflineStoreServiceHandler) GetQueryValues(context.Context, *connect.Request[v11.GetQueryValuesRequest]) (*connect.Response[v11.GetQueryValuesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.engine.v1.OfflineStoreService.GetQueryValues is not implemented"))
}