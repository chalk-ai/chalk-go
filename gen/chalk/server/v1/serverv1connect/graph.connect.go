// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/graph.proto

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
	// GraphServiceName is the fully-qualified name of the GraphService service.
	GraphServiceName = "chalk.server.v1.GraphService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GraphServiceGetFeatureSQLProcedure is the fully-qualified name of the GraphService's
	// GetFeatureSQL RPC.
	GraphServiceGetFeatureSQLProcedure = "/chalk.server.v1.GraphService/GetFeatureSQL"
	// GraphServiceGetFeaturesMetadataProcedure is the fully-qualified name of the GraphService's
	// GetFeaturesMetadata RPC.
	GraphServiceGetFeaturesMetadataProcedure = "/chalk.server.v1.GraphService/GetFeaturesMetadata"
	// GraphServiceGetGraphProcedure is the fully-qualified name of the GraphService's GetGraph RPC.
	GraphServiceGetGraphProcedure = "/chalk.server.v1.GraphService/GetGraph"
	// GraphServiceUpdateGraphProcedure is the fully-qualified name of the GraphService's UpdateGraph
	// RPC.
	GraphServiceUpdateGraphProcedure = "/chalk.server.v1.GraphService/UpdateGraph"
	// GraphServiceGetPythonFeaturesFromGraphProcedure is the fully-qualified name of the GraphService's
	// GetPythonFeaturesFromGraph RPC.
	GraphServiceGetPythonFeaturesFromGraphProcedure = "/chalk.server.v1.GraphService/GetPythonFeaturesFromGraph"
)

// GraphServiceClient is a client for the chalk.server.v1.GraphService service.
type GraphServiceClient interface {
	// GetFeatureSQL returns the feature SQLs for a given deployment.
	GetFeatureSQL(context.Context, *connect.Request[v1.GetFeatureSQLRequest]) (*connect.Response[v1.GetFeatureSQLResponse], error)
	GetFeaturesMetadata(context.Context, *connect.Request[v1.GetFeaturesMetadataRequest]) (*connect.Response[v1.GetFeaturesMetadataResponse], error)
	GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error)
	// UpdateGraph uploads the protobuf graph for a given deployment.
	UpdateGraph(context.Context, *connect.Request[v1.UpdateGraphRequest]) (*connect.Response[v1.UpdateGraphResponse], error)
	// GetPythonFeaturesFromGraph returns chalk python features generated from the protograph
	GetPythonFeaturesFromGraph(context.Context, *connect.Request[v1.GetPythonFeaturesFromGraphRequest]) (*connect.Response[v1.GetPythonFeaturesFromGraphResponse], error)
}

// NewGraphServiceClient constructs a client for the chalk.server.v1.GraphService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGraphServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GraphServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	graphServiceMethods := v1.File_chalk_server_v1_graph_proto.Services().ByName("GraphService").Methods()
	return &graphServiceClient{
		getFeatureSQL: connect.NewClient[v1.GetFeatureSQLRequest, v1.GetFeatureSQLResponse](
			httpClient,
			baseURL+GraphServiceGetFeatureSQLProcedure,
			connect.WithSchema(graphServiceMethods.ByName("GetFeatureSQL")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getFeaturesMetadata: connect.NewClient[v1.GetFeaturesMetadataRequest, v1.GetFeaturesMetadataResponse](
			httpClient,
			baseURL+GraphServiceGetFeaturesMetadataProcedure,
			connect.WithSchema(graphServiceMethods.ByName("GetFeaturesMetadata")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getGraph: connect.NewClient[v1.GetGraphRequest, v1.GetGraphResponse](
			httpClient,
			baseURL+GraphServiceGetGraphProcedure,
			connect.WithSchema(graphServiceMethods.ByName("GetGraph")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		updateGraph: connect.NewClient[v1.UpdateGraphRequest, v1.UpdateGraphResponse](
			httpClient,
			baseURL+GraphServiceUpdateGraphProcedure,
			connect.WithSchema(graphServiceMethods.ByName("UpdateGraph")),
			connect.WithClientOptions(opts...),
		),
		getPythonFeaturesFromGraph: connect.NewClient[v1.GetPythonFeaturesFromGraphRequest, v1.GetPythonFeaturesFromGraphResponse](
			httpClient,
			baseURL+GraphServiceGetPythonFeaturesFromGraphProcedure,
			connect.WithSchema(graphServiceMethods.ByName("GetPythonFeaturesFromGraph")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// graphServiceClient implements GraphServiceClient.
type graphServiceClient struct {
	getFeatureSQL              *connect.Client[v1.GetFeatureSQLRequest, v1.GetFeatureSQLResponse]
	getFeaturesMetadata        *connect.Client[v1.GetFeaturesMetadataRequest, v1.GetFeaturesMetadataResponse]
	getGraph                   *connect.Client[v1.GetGraphRequest, v1.GetGraphResponse]
	updateGraph                *connect.Client[v1.UpdateGraphRequest, v1.UpdateGraphResponse]
	getPythonFeaturesFromGraph *connect.Client[v1.GetPythonFeaturesFromGraphRequest, v1.GetPythonFeaturesFromGraphResponse]
}

// GetFeatureSQL calls chalk.server.v1.GraphService.GetFeatureSQL.
func (c *graphServiceClient) GetFeatureSQL(ctx context.Context, req *connect.Request[v1.GetFeatureSQLRequest]) (*connect.Response[v1.GetFeatureSQLResponse], error) {
	return c.getFeatureSQL.CallUnary(ctx, req)
}

// GetFeaturesMetadata calls chalk.server.v1.GraphService.GetFeaturesMetadata.
func (c *graphServiceClient) GetFeaturesMetadata(ctx context.Context, req *connect.Request[v1.GetFeaturesMetadataRequest]) (*connect.Response[v1.GetFeaturesMetadataResponse], error) {
	return c.getFeaturesMetadata.CallUnary(ctx, req)
}

// GetGraph calls chalk.server.v1.GraphService.GetGraph.
func (c *graphServiceClient) GetGraph(ctx context.Context, req *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error) {
	return c.getGraph.CallUnary(ctx, req)
}

// UpdateGraph calls chalk.server.v1.GraphService.UpdateGraph.
func (c *graphServiceClient) UpdateGraph(ctx context.Context, req *connect.Request[v1.UpdateGraphRequest]) (*connect.Response[v1.UpdateGraphResponse], error) {
	return c.updateGraph.CallUnary(ctx, req)
}

// GetPythonFeaturesFromGraph calls chalk.server.v1.GraphService.GetPythonFeaturesFromGraph.
func (c *graphServiceClient) GetPythonFeaturesFromGraph(ctx context.Context, req *connect.Request[v1.GetPythonFeaturesFromGraphRequest]) (*connect.Response[v1.GetPythonFeaturesFromGraphResponse], error) {
	return c.getPythonFeaturesFromGraph.CallUnary(ctx, req)
}

// GraphServiceHandler is an implementation of the chalk.server.v1.GraphService service.
type GraphServiceHandler interface {
	// GetFeatureSQL returns the feature SQLs for a given deployment.
	GetFeatureSQL(context.Context, *connect.Request[v1.GetFeatureSQLRequest]) (*connect.Response[v1.GetFeatureSQLResponse], error)
	GetFeaturesMetadata(context.Context, *connect.Request[v1.GetFeaturesMetadataRequest]) (*connect.Response[v1.GetFeaturesMetadataResponse], error)
	GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error)
	// UpdateGraph uploads the protobuf graph for a given deployment.
	UpdateGraph(context.Context, *connect.Request[v1.UpdateGraphRequest]) (*connect.Response[v1.UpdateGraphResponse], error)
	// GetPythonFeaturesFromGraph returns chalk python features generated from the protograph
	GetPythonFeaturesFromGraph(context.Context, *connect.Request[v1.GetPythonFeaturesFromGraphRequest]) (*connect.Response[v1.GetPythonFeaturesFromGraphResponse], error)
}

// NewGraphServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGraphServiceHandler(svc GraphServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	graphServiceMethods := v1.File_chalk_server_v1_graph_proto.Services().ByName("GraphService").Methods()
	graphServiceGetFeatureSQLHandler := connect.NewUnaryHandler(
		GraphServiceGetFeatureSQLProcedure,
		svc.GetFeatureSQL,
		connect.WithSchema(graphServiceMethods.ByName("GetFeatureSQL")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceGetFeaturesMetadataHandler := connect.NewUnaryHandler(
		GraphServiceGetFeaturesMetadataProcedure,
		svc.GetFeaturesMetadata,
		connect.WithSchema(graphServiceMethods.ByName("GetFeaturesMetadata")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceGetGraphHandler := connect.NewUnaryHandler(
		GraphServiceGetGraphProcedure,
		svc.GetGraph,
		connect.WithSchema(graphServiceMethods.ByName("GetGraph")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceUpdateGraphHandler := connect.NewUnaryHandler(
		GraphServiceUpdateGraphProcedure,
		svc.UpdateGraph,
		connect.WithSchema(graphServiceMethods.ByName("UpdateGraph")),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceGetPythonFeaturesFromGraphHandler := connect.NewUnaryHandler(
		GraphServiceGetPythonFeaturesFromGraphProcedure,
		svc.GetPythonFeaturesFromGraph,
		connect.WithSchema(graphServiceMethods.ByName("GetPythonFeaturesFromGraph")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.GraphService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GraphServiceGetFeatureSQLProcedure:
			graphServiceGetFeatureSQLHandler.ServeHTTP(w, r)
		case GraphServiceGetFeaturesMetadataProcedure:
			graphServiceGetFeaturesMetadataHandler.ServeHTTP(w, r)
		case GraphServiceGetGraphProcedure:
			graphServiceGetGraphHandler.ServeHTTP(w, r)
		case GraphServiceUpdateGraphProcedure:
			graphServiceUpdateGraphHandler.ServeHTTP(w, r)
		case GraphServiceGetPythonFeaturesFromGraphProcedure:
			graphServiceGetPythonFeaturesFromGraphHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGraphServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGraphServiceHandler struct{}

func (UnimplementedGraphServiceHandler) GetFeatureSQL(context.Context, *connect.Request[v1.GetFeatureSQLRequest]) (*connect.Response[v1.GetFeatureSQLResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.GraphService.GetFeatureSQL is not implemented"))
}

func (UnimplementedGraphServiceHandler) GetFeaturesMetadata(context.Context, *connect.Request[v1.GetFeaturesMetadataRequest]) (*connect.Response[v1.GetFeaturesMetadataResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.GraphService.GetFeaturesMetadata is not implemented"))
}

func (UnimplementedGraphServiceHandler) GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.GraphService.GetGraph is not implemented"))
}

func (UnimplementedGraphServiceHandler) UpdateGraph(context.Context, *connect.Request[v1.UpdateGraphRequest]) (*connect.Response[v1.UpdateGraphResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.GraphService.UpdateGraph is not implemented"))
}

func (UnimplementedGraphServiceHandler) GetPythonFeaturesFromGraph(context.Context, *connect.Request[v1.GetPythonFeaturesFromGraphRequest]) (*connect.Response[v1.GetPythonFeaturesFromGraphResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.GraphService.GetPythonFeaturesFromGraph is not implemented"))
}
