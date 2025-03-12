// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chalk/server/v1/prompt.proto

package serverv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/server/v1"
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
	// PromptServiceName is the fully-qualified name of the PromptService service.
	PromptServiceName = "chalk.server.v1.PromptService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PromptServiceListNamedPromptsProcedure is the fully-qualified name of the PromptService's
	// ListNamedPrompts RPC.
	PromptServiceListNamedPromptsProcedure = "/chalk.server.v1.PromptService/ListNamedPrompts"
	// PromptServiceGetNamedPromptProcedure is the fully-qualified name of the PromptService's
	// GetNamedPrompt RPC.
	PromptServiceGetNamedPromptProcedure = "/chalk.server.v1.PromptService/GetNamedPrompt"
	// PromptServiceCreateNamedPromptProcedure is the fully-qualified name of the PromptService's
	// CreateNamedPrompt RPC.
	PromptServiceCreateNamedPromptProcedure = "/chalk.server.v1.PromptService/CreateNamedPrompt"
	// PromptServiceUpdateNamedPromptProcedure is the fully-qualified name of the PromptService's
	// UpdateNamedPrompt RPC.
	PromptServiceUpdateNamedPromptProcedure = "/chalk.server.v1.PromptService/UpdateNamedPrompt"
	// PromptServiceListPromptVariantsProcedure is the fully-qualified name of the PromptService's
	// ListPromptVariants RPC.
	PromptServiceListPromptVariantsProcedure = "/chalk.server.v1.PromptService/ListPromptVariants"
	// PromptServiceListPromptEvaluationRunsProcedure is the fully-qualified name of the PromptService's
	// ListPromptEvaluationRuns RPC.
	PromptServiceListPromptEvaluationRunsProcedure = "/chalk.server.v1.PromptService/ListPromptEvaluationRuns"
	// PromptServiceGetPromptEvaluationRunProcedure is the fully-qualified name of the PromptService's
	// GetPromptEvaluationRun RPC.
	PromptServiceGetPromptEvaluationRunProcedure = "/chalk.server.v1.PromptService/GetPromptEvaluationRun"
	// PromptServiceCreatePromptEvaluationRunProcedure is the fully-qualified name of the
	// PromptService's CreatePromptEvaluationRun RPC.
	PromptServiceCreatePromptEvaluationRunProcedure = "/chalk.server.v1.PromptService/CreatePromptEvaluationRun"
)

// PromptServiceClient is a client for the chalk.server.v1.PromptService service.
type PromptServiceClient interface {
	// Return a list of named prompts
	ListNamedPrompts(context.Context, *connect.Request[v1.ListNamedPromptsRequest]) (*connect.Response[v1.ListNamedPromptsResponse], error)
	// Return a single named prompt with the latest prompt variant
	GetNamedPrompt(context.Context, *connect.Request[v1.GetNamedPromptRequest]) (*connect.Response[v1.GetNamedPromptResponse], error)
	// Create a new named prompt and its corresponding prompt variant
	CreateNamedPrompt(context.Context, *connect.Request[v1.CreateNamedPromptRequest]) (*connect.Response[v1.CreateNamedPromptResponse], error)
	// Modify a named prompt; may create a new prompt variant if variant settings changed
	UpdateNamedPrompt(context.Context, *connect.Request[v1.UpdateNamedPromptRequest]) (*connect.Response[v1.UpdateNamedPromptResponse], error)
	// Return a list of prompt variants
	ListPromptVariants(context.Context, *connect.Request[v1.ListPromptVariantsRequest]) (*connect.Response[v1.ListPromptVariantsResponse], error)
	// Return a list of prompt evaluation runs
	ListPromptEvaluationRuns(context.Context, *connect.Request[v1.ListPromptEvaluationRunsRequest]) (*connect.Response[v1.ListPromptEvaluationRunsResponse], error)
	// Return a prompt evaluation run and its prompt variants
	GetPromptEvaluationRun(context.Context, *connect.Request[v1.GetPromptEvaluationRunRequest]) (*connect.Response[v1.GetPromptEvaluationRunResponse], error)
	// Start a new prompt evaluation run
	CreatePromptEvaluationRun(context.Context, *connect.Request[v1.CreatePromptEvaluationRunRequest]) (*connect.Response[v1.CreatePromptEvaluationRunResponse], error)
}

// NewPromptServiceClient constructs a client for the chalk.server.v1.PromptService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPromptServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PromptServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	promptServiceMethods := v1.File_chalk_server_v1_prompt_proto.Services().ByName("PromptService").Methods()
	return &promptServiceClient{
		listNamedPrompts: connect.NewClient[v1.ListNamedPromptsRequest, v1.ListNamedPromptsResponse](
			httpClient,
			baseURL+PromptServiceListNamedPromptsProcedure,
			connect.WithSchema(promptServiceMethods.ByName("ListNamedPrompts")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getNamedPrompt: connect.NewClient[v1.GetNamedPromptRequest, v1.GetNamedPromptResponse](
			httpClient,
			baseURL+PromptServiceGetNamedPromptProcedure,
			connect.WithSchema(promptServiceMethods.ByName("GetNamedPrompt")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createNamedPrompt: connect.NewClient[v1.CreateNamedPromptRequest, v1.CreateNamedPromptResponse](
			httpClient,
			baseURL+PromptServiceCreateNamedPromptProcedure,
			connect.WithSchema(promptServiceMethods.ByName("CreateNamedPrompt")),
			connect.WithClientOptions(opts...),
		),
		updateNamedPrompt: connect.NewClient[v1.UpdateNamedPromptRequest, v1.UpdateNamedPromptResponse](
			httpClient,
			baseURL+PromptServiceUpdateNamedPromptProcedure,
			connect.WithSchema(promptServiceMethods.ByName("UpdateNamedPrompt")),
			connect.WithClientOptions(opts...),
		),
		listPromptVariants: connect.NewClient[v1.ListPromptVariantsRequest, v1.ListPromptVariantsResponse](
			httpClient,
			baseURL+PromptServiceListPromptVariantsProcedure,
			connect.WithSchema(promptServiceMethods.ByName("ListPromptVariants")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listPromptEvaluationRuns: connect.NewClient[v1.ListPromptEvaluationRunsRequest, v1.ListPromptEvaluationRunsResponse](
			httpClient,
			baseURL+PromptServiceListPromptEvaluationRunsProcedure,
			connect.WithSchema(promptServiceMethods.ByName("ListPromptEvaluationRuns")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getPromptEvaluationRun: connect.NewClient[v1.GetPromptEvaluationRunRequest, v1.GetPromptEvaluationRunResponse](
			httpClient,
			baseURL+PromptServiceGetPromptEvaluationRunProcedure,
			connect.WithSchema(promptServiceMethods.ByName("GetPromptEvaluationRun")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createPromptEvaluationRun: connect.NewClient[v1.CreatePromptEvaluationRunRequest, v1.CreatePromptEvaluationRunResponse](
			httpClient,
			baseURL+PromptServiceCreatePromptEvaluationRunProcedure,
			connect.WithSchema(promptServiceMethods.ByName("CreatePromptEvaluationRun")),
			connect.WithClientOptions(opts...),
		),
	}
}

// promptServiceClient implements PromptServiceClient.
type promptServiceClient struct {
	listNamedPrompts          *connect.Client[v1.ListNamedPromptsRequest, v1.ListNamedPromptsResponse]
	getNamedPrompt            *connect.Client[v1.GetNamedPromptRequest, v1.GetNamedPromptResponse]
	createNamedPrompt         *connect.Client[v1.CreateNamedPromptRequest, v1.CreateNamedPromptResponse]
	updateNamedPrompt         *connect.Client[v1.UpdateNamedPromptRequest, v1.UpdateNamedPromptResponse]
	listPromptVariants        *connect.Client[v1.ListPromptVariantsRequest, v1.ListPromptVariantsResponse]
	listPromptEvaluationRuns  *connect.Client[v1.ListPromptEvaluationRunsRequest, v1.ListPromptEvaluationRunsResponse]
	getPromptEvaluationRun    *connect.Client[v1.GetPromptEvaluationRunRequest, v1.GetPromptEvaluationRunResponse]
	createPromptEvaluationRun *connect.Client[v1.CreatePromptEvaluationRunRequest, v1.CreatePromptEvaluationRunResponse]
}

// ListNamedPrompts calls chalk.server.v1.PromptService.ListNamedPrompts.
func (c *promptServiceClient) ListNamedPrompts(ctx context.Context, req *connect.Request[v1.ListNamedPromptsRequest]) (*connect.Response[v1.ListNamedPromptsResponse], error) {
	return c.listNamedPrompts.CallUnary(ctx, req)
}

// GetNamedPrompt calls chalk.server.v1.PromptService.GetNamedPrompt.
func (c *promptServiceClient) GetNamedPrompt(ctx context.Context, req *connect.Request[v1.GetNamedPromptRequest]) (*connect.Response[v1.GetNamedPromptResponse], error) {
	return c.getNamedPrompt.CallUnary(ctx, req)
}

// CreateNamedPrompt calls chalk.server.v1.PromptService.CreateNamedPrompt.
func (c *promptServiceClient) CreateNamedPrompt(ctx context.Context, req *connect.Request[v1.CreateNamedPromptRequest]) (*connect.Response[v1.CreateNamedPromptResponse], error) {
	return c.createNamedPrompt.CallUnary(ctx, req)
}

// UpdateNamedPrompt calls chalk.server.v1.PromptService.UpdateNamedPrompt.
func (c *promptServiceClient) UpdateNamedPrompt(ctx context.Context, req *connect.Request[v1.UpdateNamedPromptRequest]) (*connect.Response[v1.UpdateNamedPromptResponse], error) {
	return c.updateNamedPrompt.CallUnary(ctx, req)
}

// ListPromptVariants calls chalk.server.v1.PromptService.ListPromptVariants.
func (c *promptServiceClient) ListPromptVariants(ctx context.Context, req *connect.Request[v1.ListPromptVariantsRequest]) (*connect.Response[v1.ListPromptVariantsResponse], error) {
	return c.listPromptVariants.CallUnary(ctx, req)
}

// ListPromptEvaluationRuns calls chalk.server.v1.PromptService.ListPromptEvaluationRuns.
func (c *promptServiceClient) ListPromptEvaluationRuns(ctx context.Context, req *connect.Request[v1.ListPromptEvaluationRunsRequest]) (*connect.Response[v1.ListPromptEvaluationRunsResponse], error) {
	return c.listPromptEvaluationRuns.CallUnary(ctx, req)
}

// GetPromptEvaluationRun calls chalk.server.v1.PromptService.GetPromptEvaluationRun.
func (c *promptServiceClient) GetPromptEvaluationRun(ctx context.Context, req *connect.Request[v1.GetPromptEvaluationRunRequest]) (*connect.Response[v1.GetPromptEvaluationRunResponse], error) {
	return c.getPromptEvaluationRun.CallUnary(ctx, req)
}

// CreatePromptEvaluationRun calls chalk.server.v1.PromptService.CreatePromptEvaluationRun.
func (c *promptServiceClient) CreatePromptEvaluationRun(ctx context.Context, req *connect.Request[v1.CreatePromptEvaluationRunRequest]) (*connect.Response[v1.CreatePromptEvaluationRunResponse], error) {
	return c.createPromptEvaluationRun.CallUnary(ctx, req)
}

// PromptServiceHandler is an implementation of the chalk.server.v1.PromptService service.
type PromptServiceHandler interface {
	// Return a list of named prompts
	ListNamedPrompts(context.Context, *connect.Request[v1.ListNamedPromptsRequest]) (*connect.Response[v1.ListNamedPromptsResponse], error)
	// Return a single named prompt with the latest prompt variant
	GetNamedPrompt(context.Context, *connect.Request[v1.GetNamedPromptRequest]) (*connect.Response[v1.GetNamedPromptResponse], error)
	// Create a new named prompt and its corresponding prompt variant
	CreateNamedPrompt(context.Context, *connect.Request[v1.CreateNamedPromptRequest]) (*connect.Response[v1.CreateNamedPromptResponse], error)
	// Modify a named prompt; may create a new prompt variant if variant settings changed
	UpdateNamedPrompt(context.Context, *connect.Request[v1.UpdateNamedPromptRequest]) (*connect.Response[v1.UpdateNamedPromptResponse], error)
	// Return a list of prompt variants
	ListPromptVariants(context.Context, *connect.Request[v1.ListPromptVariantsRequest]) (*connect.Response[v1.ListPromptVariantsResponse], error)
	// Return a list of prompt evaluation runs
	ListPromptEvaluationRuns(context.Context, *connect.Request[v1.ListPromptEvaluationRunsRequest]) (*connect.Response[v1.ListPromptEvaluationRunsResponse], error)
	// Return a prompt evaluation run and its prompt variants
	GetPromptEvaluationRun(context.Context, *connect.Request[v1.GetPromptEvaluationRunRequest]) (*connect.Response[v1.GetPromptEvaluationRunResponse], error)
	// Start a new prompt evaluation run
	CreatePromptEvaluationRun(context.Context, *connect.Request[v1.CreatePromptEvaluationRunRequest]) (*connect.Response[v1.CreatePromptEvaluationRunResponse], error)
}

// NewPromptServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPromptServiceHandler(svc PromptServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	promptServiceMethods := v1.File_chalk_server_v1_prompt_proto.Services().ByName("PromptService").Methods()
	promptServiceListNamedPromptsHandler := connect.NewUnaryHandler(
		PromptServiceListNamedPromptsProcedure,
		svc.ListNamedPrompts,
		connect.WithSchema(promptServiceMethods.ByName("ListNamedPrompts")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceGetNamedPromptHandler := connect.NewUnaryHandler(
		PromptServiceGetNamedPromptProcedure,
		svc.GetNamedPrompt,
		connect.WithSchema(promptServiceMethods.ByName("GetNamedPrompt")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceCreateNamedPromptHandler := connect.NewUnaryHandler(
		PromptServiceCreateNamedPromptProcedure,
		svc.CreateNamedPrompt,
		connect.WithSchema(promptServiceMethods.ByName("CreateNamedPrompt")),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceUpdateNamedPromptHandler := connect.NewUnaryHandler(
		PromptServiceUpdateNamedPromptProcedure,
		svc.UpdateNamedPrompt,
		connect.WithSchema(promptServiceMethods.ByName("UpdateNamedPrompt")),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceListPromptVariantsHandler := connect.NewUnaryHandler(
		PromptServiceListPromptVariantsProcedure,
		svc.ListPromptVariants,
		connect.WithSchema(promptServiceMethods.ByName("ListPromptVariants")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceListPromptEvaluationRunsHandler := connect.NewUnaryHandler(
		PromptServiceListPromptEvaluationRunsProcedure,
		svc.ListPromptEvaluationRuns,
		connect.WithSchema(promptServiceMethods.ByName("ListPromptEvaluationRuns")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceGetPromptEvaluationRunHandler := connect.NewUnaryHandler(
		PromptServiceGetPromptEvaluationRunProcedure,
		svc.GetPromptEvaluationRun,
		connect.WithSchema(promptServiceMethods.ByName("GetPromptEvaluationRun")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	promptServiceCreatePromptEvaluationRunHandler := connect.NewUnaryHandler(
		PromptServiceCreatePromptEvaluationRunProcedure,
		svc.CreatePromptEvaluationRun,
		connect.WithSchema(promptServiceMethods.ByName("CreatePromptEvaluationRun")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chalk.server.v1.PromptService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PromptServiceListNamedPromptsProcedure:
			promptServiceListNamedPromptsHandler.ServeHTTP(w, r)
		case PromptServiceGetNamedPromptProcedure:
			promptServiceGetNamedPromptHandler.ServeHTTP(w, r)
		case PromptServiceCreateNamedPromptProcedure:
			promptServiceCreateNamedPromptHandler.ServeHTTP(w, r)
		case PromptServiceUpdateNamedPromptProcedure:
			promptServiceUpdateNamedPromptHandler.ServeHTTP(w, r)
		case PromptServiceListPromptVariantsProcedure:
			promptServiceListPromptVariantsHandler.ServeHTTP(w, r)
		case PromptServiceListPromptEvaluationRunsProcedure:
			promptServiceListPromptEvaluationRunsHandler.ServeHTTP(w, r)
		case PromptServiceGetPromptEvaluationRunProcedure:
			promptServiceGetPromptEvaluationRunHandler.ServeHTTP(w, r)
		case PromptServiceCreatePromptEvaluationRunProcedure:
			promptServiceCreatePromptEvaluationRunHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPromptServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPromptServiceHandler struct{}

func (UnimplementedPromptServiceHandler) ListNamedPrompts(context.Context, *connect.Request[v1.ListNamedPromptsRequest]) (*connect.Response[v1.ListNamedPromptsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.ListNamedPrompts is not implemented"))
}

func (UnimplementedPromptServiceHandler) GetNamedPrompt(context.Context, *connect.Request[v1.GetNamedPromptRequest]) (*connect.Response[v1.GetNamedPromptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.GetNamedPrompt is not implemented"))
}

func (UnimplementedPromptServiceHandler) CreateNamedPrompt(context.Context, *connect.Request[v1.CreateNamedPromptRequest]) (*connect.Response[v1.CreateNamedPromptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.CreateNamedPrompt is not implemented"))
}

func (UnimplementedPromptServiceHandler) UpdateNamedPrompt(context.Context, *connect.Request[v1.UpdateNamedPromptRequest]) (*connect.Response[v1.UpdateNamedPromptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.UpdateNamedPrompt is not implemented"))
}

func (UnimplementedPromptServiceHandler) ListPromptVariants(context.Context, *connect.Request[v1.ListPromptVariantsRequest]) (*connect.Response[v1.ListPromptVariantsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.ListPromptVariants is not implemented"))
}

func (UnimplementedPromptServiceHandler) ListPromptEvaluationRuns(context.Context, *connect.Request[v1.ListPromptEvaluationRunsRequest]) (*connect.Response[v1.ListPromptEvaluationRunsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.ListPromptEvaluationRuns is not implemented"))
}

func (UnimplementedPromptServiceHandler) GetPromptEvaluationRun(context.Context, *connect.Request[v1.GetPromptEvaluationRunRequest]) (*connect.Response[v1.GetPromptEvaluationRunResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.GetPromptEvaluationRun is not implemented"))
}

func (UnimplementedPromptServiceHandler) CreatePromptEvaluationRun(context.Context, *connect.Request[v1.CreatePromptEvaluationRunRequest]) (*connect.Response[v1.CreatePromptEvaluationRunResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chalk.server.v1.PromptService.CreatePromptEvaluationRun is not implemented"))
}
