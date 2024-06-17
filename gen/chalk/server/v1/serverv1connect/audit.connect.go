// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/audit.proto

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
	// AuditServiceName is the fully-qualified name of the AuditService service.
	AuditServiceName = "chalk.server.v1.AuditService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuditServiceGetAuditLogsProcedure is the fully-qualified name of the AuditService's GetAuditLogs
	// RPC.
	AuditServiceGetAuditLogsProcedure = "/chalk.server.v1.AuditService/GetAuditLogs"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	auditServiceServiceDescriptor            = v1.File_chalk_server_v1_audit_proto.Services().ByName("AuditService")
	auditServiceGetAuditLogsMethodDescriptor = auditServiceServiceDescriptor.Methods().ByName("GetAuditLogs")
)

// AuditServiceClient is a client for the chalk.server.v1.AuditService service.
type AuditServiceClient interface {
	GetAuditLogs(context.Context, *connect.Request[v1.GetAuditLogsRequest]) (*connect.Response[v1.GetAuditLogsResponse], error)
}

// NewAuditServiceClient constructs a client for the chalk.server.v1.AuditService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuditServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuditServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &auditServiceClient{
		getAuditLogs: connect.NewClient[v1.GetAuditLogsRequest, v1.GetAuditLogsResponse](
			httpClient,
			baseURL+AuditServiceGetAuditLogsProcedure,
			connect.WithSchema(auditServiceGetAuditLogsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// auditServiceClient implements AuditServiceClient.
type auditServiceClient struct {
	getAuditLogs *connect.Client[v1.GetAuditLogsRequest, v1.GetAuditLogsResponse]
}

// GetAuditLogs calls chalk.server.v1.AuditService.GetAuditLogs.
func (c *auditServiceClient) GetAuditLogs(ctx context.Context, req *connect.Request[v1.GetAuditLogsRequest]) (*connect.Response[v1.GetAuditLogsResponse], error) {
	return c.getAuditLogs.CallUnary(ctx, req)
}

// AuditServiceHandler is an implementation of the chalk.server.v1.AuditService service.
type AuditServiceHandler interface {
	GetAuditLogs(context.Context, *connect.Request[v1.GetAuditLogsRequest]) (*connect.Response[v1.GetAuditLogsResponse], error)
}

// NewAuditServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuditServiceHandler(svc AuditServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	auditServiceGetAuditLogsHandler := connect.NewUnaryHandler(
		AuditServiceGetAuditLogsProcedure,
		svc.GetAuditLogs,
		connect.WithSchema(auditServiceGetAuditLogsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.AuditService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuditServiceGetAuditLogsProcedure:
			auditServiceGetAuditLogsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuditServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuditServiceHandler struct{}

func (UnimplementedAuditServiceHandler) GetAuditLogs(context.Context, *connect.Request[v1.GetAuditLogsRequest]) (*connect.Response[v1.GetAuditLogsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.AuditService.GetAuditLogs is not implemented"))
}
