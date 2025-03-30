// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/webhook.proto

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
	// WebhookServiceName is the fully-qualified name of the WebhookService service.
	WebhookServiceName = "chalk.server.v1.WebhookService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WebhookServiceCreateWebhookProcedure is the fully-qualified name of the WebhookService's
	// CreateWebhook RPC.
	WebhookServiceCreateWebhookProcedure = "/chalk.server.v1.WebhookService/CreateWebhook"
	// WebhookServiceUpdateWebhookProcedure is the fully-qualified name of the WebhookService's
	// UpdateWebhook RPC.
	WebhookServiceUpdateWebhookProcedure = "/chalk.server.v1.WebhookService/UpdateWebhook"
	// WebhookServiceDeleteWebhookProcedure is the fully-qualified name of the WebhookService's
	// DeleteWebhook RPC.
	WebhookServiceDeleteWebhookProcedure = "/chalk.server.v1.WebhookService/DeleteWebhook"
	// WebhookServiceGetWebhookProcedure is the fully-qualified name of the WebhookService's GetWebhook
	// RPC.
	WebhookServiceGetWebhookProcedure = "/chalk.server.v1.WebhookService/GetWebhook"
	// WebhookServiceListWebhooksProcedure is the fully-qualified name of the WebhookService's
	// ListWebhooks RPC.
	WebhookServiceListWebhooksProcedure = "/chalk.server.v1.WebhookService/ListWebhooks"
)

// WebhookServiceClient is a client for the chalk.server.v1.WebhookService service.
type WebhookServiceClient interface {
	CreateWebhook(context.Context, *connect.Request[v1.CreateWebhookRequest]) (*connect.Response[v1.CreateWebhookResponse], error)
	UpdateWebhook(context.Context, *connect.Request[v1.UpdateWebhookRequest]) (*connect.Response[v1.UpdateWebhookResponse], error)
	DeleteWebhook(context.Context, *connect.Request[v1.DeleteWebhookRequest]) (*connect.Response[v1.DeleteWebhookResponse], error)
	GetWebhook(context.Context, *connect.Request[v1.GetWebhookRequest]) (*connect.Response[v1.GetWebhookResponse], error)
	ListWebhooks(context.Context, *connect.Request[v1.ListWebhooksRequest]) (*connect.Response[v1.ListWebhooksResponse], error)
}

// NewWebhookServiceClient constructs a client for the chalk.server.v1.WebhookService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWebhookServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WebhookServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	webhookServiceMethods := v1.File_chalk_server_v1_webhook_proto.Services().ByName("WebhookService").Methods()
	return &webhookServiceClient{
		createWebhook: connect.NewClient[v1.CreateWebhookRequest, v1.CreateWebhookResponse](
			httpClient,
			baseURL+WebhookServiceCreateWebhookProcedure,
			connect.WithSchema(webhookServiceMethods.ByName("CreateWebhook")),
			connect.WithClientOptions(opts...),
		),
		updateWebhook: connect.NewClient[v1.UpdateWebhookRequest, v1.UpdateWebhookResponse](
			httpClient,
			baseURL+WebhookServiceUpdateWebhookProcedure,
			connect.WithSchema(webhookServiceMethods.ByName("UpdateWebhook")),
			connect.WithClientOptions(opts...),
		),
		deleteWebhook: connect.NewClient[v1.DeleteWebhookRequest, v1.DeleteWebhookResponse](
			httpClient,
			baseURL+WebhookServiceDeleteWebhookProcedure,
			connect.WithSchema(webhookServiceMethods.ByName("DeleteWebhook")),
			connect.WithClientOptions(opts...),
		),
		getWebhook: connect.NewClient[v1.GetWebhookRequest, v1.GetWebhookResponse](
			httpClient,
			baseURL+WebhookServiceGetWebhookProcedure,
			connect.WithSchema(webhookServiceMethods.ByName("GetWebhook")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listWebhooks: connect.NewClient[v1.ListWebhooksRequest, v1.ListWebhooksResponse](
			httpClient,
			baseURL+WebhookServiceListWebhooksProcedure,
			connect.WithSchema(webhookServiceMethods.ByName("ListWebhooks")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// webhookServiceClient implements WebhookServiceClient.
type webhookServiceClient struct {
	createWebhook *connect.Client[v1.CreateWebhookRequest, v1.CreateWebhookResponse]
	updateWebhook *connect.Client[v1.UpdateWebhookRequest, v1.UpdateWebhookResponse]
	deleteWebhook *connect.Client[v1.DeleteWebhookRequest, v1.DeleteWebhookResponse]
	getWebhook    *connect.Client[v1.GetWebhookRequest, v1.GetWebhookResponse]
	listWebhooks  *connect.Client[v1.ListWebhooksRequest, v1.ListWebhooksResponse]
}

// CreateWebhook calls chalk.server.v1.WebhookService.CreateWebhook.
func (c *webhookServiceClient) CreateWebhook(ctx context.Context, req *connect.Request[v1.CreateWebhookRequest]) (*connect.Response[v1.CreateWebhookResponse], error) {
	return c.createWebhook.CallUnary(ctx, req)
}

// UpdateWebhook calls chalk.server.v1.WebhookService.UpdateWebhook.
func (c *webhookServiceClient) UpdateWebhook(ctx context.Context, req *connect.Request[v1.UpdateWebhookRequest]) (*connect.Response[v1.UpdateWebhookResponse], error) {
	return c.updateWebhook.CallUnary(ctx, req)
}

// DeleteWebhook calls chalk.server.v1.WebhookService.DeleteWebhook.
func (c *webhookServiceClient) DeleteWebhook(ctx context.Context, req *connect.Request[v1.DeleteWebhookRequest]) (*connect.Response[v1.DeleteWebhookResponse], error) {
	return c.deleteWebhook.CallUnary(ctx, req)
}

// GetWebhook calls chalk.server.v1.WebhookService.GetWebhook.
func (c *webhookServiceClient) GetWebhook(ctx context.Context, req *connect.Request[v1.GetWebhookRequest]) (*connect.Response[v1.GetWebhookResponse], error) {
	return c.getWebhook.CallUnary(ctx, req)
}

// ListWebhooks calls chalk.server.v1.WebhookService.ListWebhooks.
func (c *webhookServiceClient) ListWebhooks(ctx context.Context, req *connect.Request[v1.ListWebhooksRequest]) (*connect.Response[v1.ListWebhooksResponse], error) {
	return c.listWebhooks.CallUnary(ctx, req)
}

// WebhookServiceHandler is an implementation of the chalk.server.v1.WebhookService service.
type WebhookServiceHandler interface {
	CreateWebhook(context.Context, *connect.Request[v1.CreateWebhookRequest]) (*connect.Response[v1.CreateWebhookResponse], error)
	UpdateWebhook(context.Context, *connect.Request[v1.UpdateWebhookRequest]) (*connect.Response[v1.UpdateWebhookResponse], error)
	DeleteWebhook(context.Context, *connect.Request[v1.DeleteWebhookRequest]) (*connect.Response[v1.DeleteWebhookResponse], error)
	GetWebhook(context.Context, *connect.Request[v1.GetWebhookRequest]) (*connect.Response[v1.GetWebhookResponse], error)
	ListWebhooks(context.Context, *connect.Request[v1.ListWebhooksRequest]) (*connect.Response[v1.ListWebhooksResponse], error)
}

// NewWebhookServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWebhookServiceHandler(svc WebhookServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	webhookServiceMethods := v1.File_chalk_server_v1_webhook_proto.Services().ByName("WebhookService").Methods()
	webhookServiceCreateWebhookHandler := connect.NewUnaryHandler(
		WebhookServiceCreateWebhookProcedure,
		svc.CreateWebhook,
		connect.WithSchema(webhookServiceMethods.ByName("CreateWebhook")),
		connect.WithHandlerOptions(opts...),
	)
	webhookServiceUpdateWebhookHandler := connect.NewUnaryHandler(
		WebhookServiceUpdateWebhookProcedure,
		svc.UpdateWebhook,
		connect.WithSchema(webhookServiceMethods.ByName("UpdateWebhook")),
		connect.WithHandlerOptions(opts...),
	)
	webhookServiceDeleteWebhookHandler := connect.NewUnaryHandler(
		WebhookServiceDeleteWebhookProcedure,
		svc.DeleteWebhook,
		connect.WithSchema(webhookServiceMethods.ByName("DeleteWebhook")),
		connect.WithHandlerOptions(opts...),
	)
	webhookServiceGetWebhookHandler := connect.NewUnaryHandler(
		WebhookServiceGetWebhookProcedure,
		svc.GetWebhook,
		connect.WithSchema(webhookServiceMethods.ByName("GetWebhook")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	webhookServiceListWebhooksHandler := connect.NewUnaryHandler(
		WebhookServiceListWebhooksProcedure,
		svc.ListWebhooks,
		connect.WithSchema(webhookServiceMethods.ByName("ListWebhooks")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.WebhookService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WebhookServiceCreateWebhookProcedure:
			webhookServiceCreateWebhookHandler.ServeHTTP(w, r)
		case WebhookServiceUpdateWebhookProcedure:
			webhookServiceUpdateWebhookHandler.ServeHTTP(w, r)
		case WebhookServiceDeleteWebhookProcedure:
			webhookServiceDeleteWebhookHandler.ServeHTTP(w, r)
		case WebhookServiceGetWebhookProcedure:
			webhookServiceGetWebhookHandler.ServeHTTP(w, r)
		case WebhookServiceListWebhooksProcedure:
			webhookServiceListWebhooksHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWebhookServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWebhookServiceHandler struct{}

func (UnimplementedWebhookServiceHandler) CreateWebhook(context.Context, *connect.Request[v1.CreateWebhookRequest]) (*connect.Response[v1.CreateWebhookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.WebhookService.CreateWebhook is not implemented"))
}

func (UnimplementedWebhookServiceHandler) UpdateWebhook(context.Context, *connect.Request[v1.UpdateWebhookRequest]) (*connect.Response[v1.UpdateWebhookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.WebhookService.UpdateWebhook is not implemented"))
}

func (UnimplementedWebhookServiceHandler) DeleteWebhook(context.Context, *connect.Request[v1.DeleteWebhookRequest]) (*connect.Response[v1.DeleteWebhookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.WebhookService.DeleteWebhook is not implemented"))
}

func (UnimplementedWebhookServiceHandler) GetWebhook(context.Context, *connect.Request[v1.GetWebhookRequest]) (*connect.Response[v1.GetWebhookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.WebhookService.GetWebhook is not implemented"))
}

func (UnimplementedWebhookServiceHandler) ListWebhooks(context.Context, *connect.Request[v1.ListWebhooksRequest]) (*connect.Response[v1.ListWebhooksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.WebhookService.ListWebhooks is not implemented"))
}
