// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/monitoring.proto

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
	// MonitoringServiceName is the fully-qualified name of the MonitoringService service.
	MonitoringServiceName = "chalk.server.v1.MonitoringService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MonitoringServiceTestPagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's TestPagerDutyIntegration RPC.
	MonitoringServiceTestPagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/TestPagerDutyIntegration"
	// MonitoringServiceAddPagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's AddPagerDutyIntegration RPC.
	MonitoringServiceAddPagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/AddPagerDutyIntegration"
	// MonitoringServiceDeletePagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's DeletePagerDutyIntegration RPC.
	MonitoringServiceDeletePagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/DeletePagerDutyIntegration"
	// MonitoringServiceSetDefaultPagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's SetDefaultPagerDutyIntegration RPC.
	MonitoringServiceSetDefaultPagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/SetDefaultPagerDutyIntegration"
	// MonitoringServiceUpdatePagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's UpdatePagerDutyIntegration RPC.
	MonitoringServiceUpdatePagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/UpdatePagerDutyIntegration"
	// MonitoringServiceGetAllPagerDutyIntegrationsProcedure is the fully-qualified name of the
	// MonitoringService's GetAllPagerDutyIntegrations RPC.
	MonitoringServiceGetAllPagerDutyIntegrationsProcedure = "/chalk.server.v1.MonitoringService/GetAllPagerDutyIntegrations"
	// MonitoringServiceGetPagerDutyIntegrationProcedure is the fully-qualified name of the
	// MonitoringService's GetPagerDutyIntegration RPC.
	MonitoringServiceGetPagerDutyIntegrationProcedure = "/chalk.server.v1.MonitoringService/GetPagerDutyIntegration"
)

// MonitoringServiceClient is a client for the chalk.server.v1.MonitoringService service.
type MonitoringServiceClient interface {
	TestPagerDutyIntegration(context.Context, *connect.Request[v1.TestPagerDutyIntegrationRequest]) (*connect.Response[v1.TestPagerDutyIntegrationResponse], error)
	AddPagerDutyIntegration(context.Context, *connect.Request[v1.AddPagerDutyIntegrationRequest]) (*connect.Response[v1.AddPagerDutyIntegrationResponse], error)
	DeletePagerDutyIntegration(context.Context, *connect.Request[v1.DeletePagerDutyIntegrationRequest]) (*connect.Response[v1.DeletePagerDutyIntegrationResponse], error)
	SetDefaultPagerDutyIntegration(context.Context, *connect.Request[v1.SetDefaultPagerDutyIntegrationRequest]) (*connect.Response[v1.SetDefaultPagerDutyIntegrationResponse], error)
	UpdatePagerDutyIntegration(context.Context, *connect.Request[v1.UpdatePagerDutyIntegrationRequest]) (*connect.Response[v1.UpdatePagerDutyIntegrationResponse], error)
	GetAllPagerDutyIntegrations(context.Context, *connect.Request[v1.GetAllPagerDutyIntegrationsRequest]) (*connect.Response[v1.GetAllPagerDutyIntegrationsResponse], error)
	GetPagerDutyIntegration(context.Context, *connect.Request[v1.GetPagerDutyIntegrationRequest]) (*connect.Response[v1.GetPagerDutyIntegrationResponse], error)
}

// NewMonitoringServiceClient constructs a client for the chalk.server.v1.MonitoringService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMonitoringServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MonitoringServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	monitoringServiceMethods := v1.File_chalk_server_v1_monitoring_proto.Services().ByName("MonitoringService").Methods()
	return &monitoringServiceClient{
		testPagerDutyIntegration: connect.NewClient[v1.TestPagerDutyIntegrationRequest, v1.TestPagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceTestPagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("TestPagerDutyIntegration")),
			connect.WithClientOptions(opts...),
		),
		addPagerDutyIntegration: connect.NewClient[v1.AddPagerDutyIntegrationRequest, v1.AddPagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceAddPagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("AddPagerDutyIntegration")),
			connect.WithClientOptions(opts...),
		),
		deletePagerDutyIntegration: connect.NewClient[v1.DeletePagerDutyIntegrationRequest, v1.DeletePagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceDeletePagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("DeletePagerDutyIntegration")),
			connect.WithClientOptions(opts...),
		),
		setDefaultPagerDutyIntegration: connect.NewClient[v1.SetDefaultPagerDutyIntegrationRequest, v1.SetDefaultPagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceSetDefaultPagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("SetDefaultPagerDutyIntegration")),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		updatePagerDutyIntegration: connect.NewClient[v1.UpdatePagerDutyIntegrationRequest, v1.UpdatePagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceUpdatePagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("UpdatePagerDutyIntegration")),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		getAllPagerDutyIntegrations: connect.NewClient[v1.GetAllPagerDutyIntegrationsRequest, v1.GetAllPagerDutyIntegrationsResponse](
			httpClient,
			baseURL+MonitoringServiceGetAllPagerDutyIntegrationsProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("GetAllPagerDutyIntegrations")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getPagerDutyIntegration: connect.NewClient[v1.GetPagerDutyIntegrationRequest, v1.GetPagerDutyIntegrationResponse](
			httpClient,
			baseURL+MonitoringServiceGetPagerDutyIntegrationProcedure,
			connect.WithSchema(monitoringServiceMethods.ByName("GetPagerDutyIntegration")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// monitoringServiceClient implements MonitoringServiceClient.
type monitoringServiceClient struct {
	testPagerDutyIntegration       *connect.Client[v1.TestPagerDutyIntegrationRequest, v1.TestPagerDutyIntegrationResponse]
	addPagerDutyIntegration        *connect.Client[v1.AddPagerDutyIntegrationRequest, v1.AddPagerDutyIntegrationResponse]
	deletePagerDutyIntegration     *connect.Client[v1.DeletePagerDutyIntegrationRequest, v1.DeletePagerDutyIntegrationResponse]
	setDefaultPagerDutyIntegration *connect.Client[v1.SetDefaultPagerDutyIntegrationRequest, v1.SetDefaultPagerDutyIntegrationResponse]
	updatePagerDutyIntegration     *connect.Client[v1.UpdatePagerDutyIntegrationRequest, v1.UpdatePagerDutyIntegrationResponse]
	getAllPagerDutyIntegrations    *connect.Client[v1.GetAllPagerDutyIntegrationsRequest, v1.GetAllPagerDutyIntegrationsResponse]
	getPagerDutyIntegration        *connect.Client[v1.GetPagerDutyIntegrationRequest, v1.GetPagerDutyIntegrationResponse]
}

// TestPagerDutyIntegration calls chalk.server.v1.MonitoringService.TestPagerDutyIntegration.
func (c *monitoringServiceClient) TestPagerDutyIntegration(ctx context.Context, req *connect.Request[v1.TestPagerDutyIntegrationRequest]) (*connect.Response[v1.TestPagerDutyIntegrationResponse], error) {
	return c.testPagerDutyIntegration.CallUnary(ctx, req)
}

// AddPagerDutyIntegration calls chalk.server.v1.MonitoringService.AddPagerDutyIntegration.
func (c *monitoringServiceClient) AddPagerDutyIntegration(ctx context.Context, req *connect.Request[v1.AddPagerDutyIntegrationRequest]) (*connect.Response[v1.AddPagerDutyIntegrationResponse], error) {
	return c.addPagerDutyIntegration.CallUnary(ctx, req)
}

// DeletePagerDutyIntegration calls chalk.server.v1.MonitoringService.DeletePagerDutyIntegration.
func (c *monitoringServiceClient) DeletePagerDutyIntegration(ctx context.Context, req *connect.Request[v1.DeletePagerDutyIntegrationRequest]) (*connect.Response[v1.DeletePagerDutyIntegrationResponse], error) {
	return c.deletePagerDutyIntegration.CallUnary(ctx, req)
}

// SetDefaultPagerDutyIntegration calls
// chalk.server.v1.MonitoringService.SetDefaultPagerDutyIntegration.
func (c *monitoringServiceClient) SetDefaultPagerDutyIntegration(ctx context.Context, req *connect.Request[v1.SetDefaultPagerDutyIntegrationRequest]) (*connect.Response[v1.SetDefaultPagerDutyIntegrationResponse], error) {
	return c.setDefaultPagerDutyIntegration.CallUnary(ctx, req)
}

// UpdatePagerDutyIntegration calls chalk.server.v1.MonitoringService.UpdatePagerDutyIntegration.
func (c *monitoringServiceClient) UpdatePagerDutyIntegration(ctx context.Context, req *connect.Request[v1.UpdatePagerDutyIntegrationRequest]) (*connect.Response[v1.UpdatePagerDutyIntegrationResponse], error) {
	return c.updatePagerDutyIntegration.CallUnary(ctx, req)
}

// GetAllPagerDutyIntegrations calls chalk.server.v1.MonitoringService.GetAllPagerDutyIntegrations.
func (c *monitoringServiceClient) GetAllPagerDutyIntegrations(ctx context.Context, req *connect.Request[v1.GetAllPagerDutyIntegrationsRequest]) (*connect.Response[v1.GetAllPagerDutyIntegrationsResponse], error) {
	return c.getAllPagerDutyIntegrations.CallUnary(ctx, req)
}

// GetPagerDutyIntegration calls chalk.server.v1.MonitoringService.GetPagerDutyIntegration.
func (c *monitoringServiceClient) GetPagerDutyIntegration(ctx context.Context, req *connect.Request[v1.GetPagerDutyIntegrationRequest]) (*connect.Response[v1.GetPagerDutyIntegrationResponse], error) {
	return c.getPagerDutyIntegration.CallUnary(ctx, req)
}

// MonitoringServiceHandler is an implementation of the chalk.server.v1.MonitoringService service.
type MonitoringServiceHandler interface {
	TestPagerDutyIntegration(context.Context, *connect.Request[v1.TestPagerDutyIntegrationRequest]) (*connect.Response[v1.TestPagerDutyIntegrationResponse], error)
	AddPagerDutyIntegration(context.Context, *connect.Request[v1.AddPagerDutyIntegrationRequest]) (*connect.Response[v1.AddPagerDutyIntegrationResponse], error)
	DeletePagerDutyIntegration(context.Context, *connect.Request[v1.DeletePagerDutyIntegrationRequest]) (*connect.Response[v1.DeletePagerDutyIntegrationResponse], error)
	SetDefaultPagerDutyIntegration(context.Context, *connect.Request[v1.SetDefaultPagerDutyIntegrationRequest]) (*connect.Response[v1.SetDefaultPagerDutyIntegrationResponse], error)
	UpdatePagerDutyIntegration(context.Context, *connect.Request[v1.UpdatePagerDutyIntegrationRequest]) (*connect.Response[v1.UpdatePagerDutyIntegrationResponse], error)
	GetAllPagerDutyIntegrations(context.Context, *connect.Request[v1.GetAllPagerDutyIntegrationsRequest]) (*connect.Response[v1.GetAllPagerDutyIntegrationsResponse], error)
	GetPagerDutyIntegration(context.Context, *connect.Request[v1.GetPagerDutyIntegrationRequest]) (*connect.Response[v1.GetPagerDutyIntegrationResponse], error)
}

// NewMonitoringServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMonitoringServiceHandler(svc MonitoringServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	monitoringServiceMethods := v1.File_chalk_server_v1_monitoring_proto.Services().ByName("MonitoringService").Methods()
	monitoringServiceTestPagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceTestPagerDutyIntegrationProcedure,
		svc.TestPagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("TestPagerDutyIntegration")),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceAddPagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceAddPagerDutyIntegrationProcedure,
		svc.AddPagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("AddPagerDutyIntegration")),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceDeletePagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceDeletePagerDutyIntegrationProcedure,
		svc.DeletePagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("DeletePagerDutyIntegration")),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceSetDefaultPagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceSetDefaultPagerDutyIntegrationProcedure,
		svc.SetDefaultPagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("SetDefaultPagerDutyIntegration")),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceUpdatePagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceUpdatePagerDutyIntegrationProcedure,
		svc.UpdatePagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("UpdatePagerDutyIntegration")),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceGetAllPagerDutyIntegrationsHandler := connect.NewUnaryHandler(
		MonitoringServiceGetAllPagerDutyIntegrationsProcedure,
		svc.GetAllPagerDutyIntegrations,
		connect.WithSchema(monitoringServiceMethods.ByName("GetAllPagerDutyIntegrations")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	monitoringServiceGetPagerDutyIntegrationHandler := connect.NewUnaryHandler(
		MonitoringServiceGetPagerDutyIntegrationProcedure,
		svc.GetPagerDutyIntegration,
		connect.WithSchema(monitoringServiceMethods.ByName("GetPagerDutyIntegration")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.MonitoringService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MonitoringServiceTestPagerDutyIntegrationProcedure:
			monitoringServiceTestPagerDutyIntegrationHandler.ServeHTTP(w, r)
		case MonitoringServiceAddPagerDutyIntegrationProcedure:
			monitoringServiceAddPagerDutyIntegrationHandler.ServeHTTP(w, r)
		case MonitoringServiceDeletePagerDutyIntegrationProcedure:
			monitoringServiceDeletePagerDutyIntegrationHandler.ServeHTTP(w, r)
		case MonitoringServiceSetDefaultPagerDutyIntegrationProcedure:
			monitoringServiceSetDefaultPagerDutyIntegrationHandler.ServeHTTP(w, r)
		case MonitoringServiceUpdatePagerDutyIntegrationProcedure:
			monitoringServiceUpdatePagerDutyIntegrationHandler.ServeHTTP(w, r)
		case MonitoringServiceGetAllPagerDutyIntegrationsProcedure:
			monitoringServiceGetAllPagerDutyIntegrationsHandler.ServeHTTP(w, r)
		case MonitoringServiceGetPagerDutyIntegrationProcedure:
			monitoringServiceGetPagerDutyIntegrationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMonitoringServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMonitoringServiceHandler struct{}

func (UnimplementedMonitoringServiceHandler) TestPagerDutyIntegration(context.Context, *connect.Request[v1.TestPagerDutyIntegrationRequest]) (*connect.Response[v1.TestPagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.TestPagerDutyIntegration is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) AddPagerDutyIntegration(context.Context, *connect.Request[v1.AddPagerDutyIntegrationRequest]) (*connect.Response[v1.AddPagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.AddPagerDutyIntegration is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) DeletePagerDutyIntegration(context.Context, *connect.Request[v1.DeletePagerDutyIntegrationRequest]) (*connect.Response[v1.DeletePagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.DeletePagerDutyIntegration is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) SetDefaultPagerDutyIntegration(context.Context, *connect.Request[v1.SetDefaultPagerDutyIntegrationRequest]) (*connect.Response[v1.SetDefaultPagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.SetDefaultPagerDutyIntegration is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) UpdatePagerDutyIntegration(context.Context, *connect.Request[v1.UpdatePagerDutyIntegrationRequest]) (*connect.Response[v1.UpdatePagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.UpdatePagerDutyIntegration is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) GetAllPagerDutyIntegrations(context.Context, *connect.Request[v1.GetAllPagerDutyIntegrationsRequest]) (*connect.Response[v1.GetAllPagerDutyIntegrationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.GetAllPagerDutyIntegrations is not implemented"))
}

func (UnimplementedMonitoringServiceHandler) GetPagerDutyIntegration(context.Context, *connect.Request[v1.GetPagerDutyIntegrationRequest]) (*connect.Response[v1.GetPagerDutyIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.MonitoringService.GetPagerDutyIntegration is not implemented"))
}
