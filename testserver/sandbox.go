package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	sandboxv1 "github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1/sandboxv1connect"
)

// sandboxServiceHandler implements the SandboxService RPC handler using a
// ResponseRegistry to provide configurable mock responses. Exec is left to the
// embedded Unimplemented handler: it is a bidi stream with no meaningful
// mock-response shape, and no caller in these tests drives it.
type sandboxServiceHandler struct {
	sandboxv1connect.UnimplementedSandboxServiceHandler
	registry *ResponseRegistry
}

func newSandboxServiceHandler(registry *ResponseRegistry) *sandboxServiceHandler {
	return &sandboxServiceHandler{registry: registry}
}

func (h *sandboxServiceHandler) CreateSandbox(
	ctx context.Context,
	req *connect.Request[sandboxv1.CreateSandboxRequest],
) (*connect.Response[sandboxv1.CreateSandboxResponse], error) {
	h.registry.CaptureRequest("CreateSandbox", req.Msg)

	if behavior := h.registry.GetBehavior("CreateSandbox"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*sandboxv1.CreateSandboxResponse)), nil
	}

	if err := h.registry.GetError("CreateSandbox"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("CreateSandbox")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for CreateSandbox"))
	}
	return connect.NewResponse(resp.(*sandboxv1.CreateSandboxResponse)), nil
}

func (h *sandboxServiceHandler) GetSandbox(
	ctx context.Context,
	req *connect.Request[sandboxv1.GetSandboxRequest],
) (*connect.Response[sandboxv1.GetSandboxResponse], error) {
	h.registry.CaptureRequest("GetSandbox", req.Msg)

	if behavior := h.registry.GetBehavior("GetSandbox"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*sandboxv1.GetSandboxResponse)), nil
	}

	if err := h.registry.GetError("GetSandbox"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("GetSandbox")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for GetSandbox"))
	}
	return connect.NewResponse(resp.(*sandboxv1.GetSandboxResponse)), nil
}

func (h *sandboxServiceHandler) ListSandboxes(
	ctx context.Context,
	req *connect.Request[sandboxv1.ListSandboxesRequest],
) (*connect.Response[sandboxv1.ListSandboxesResponse], error) {
	h.registry.CaptureRequest("ListSandboxes", req.Msg)

	if behavior := h.registry.GetBehavior("ListSandboxes"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*sandboxv1.ListSandboxesResponse)), nil
	}

	if err := h.registry.GetError("ListSandboxes"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("ListSandboxes")
	if resp == nil {
		return connect.NewResponse(&sandboxv1.ListSandboxesResponse{}), nil
	}
	return connect.NewResponse(resp.(*sandboxv1.ListSandboxesResponse)), nil
}

func (h *sandboxServiceHandler) TerminateSandbox(
	ctx context.Context,
	req *connect.Request[sandboxv1.TerminateSandboxRequest],
) (*connect.Response[sandboxv1.TerminateSandboxResponse], error) {
	h.registry.CaptureRequest("TerminateSandbox", req.Msg)

	if behavior := h.registry.GetBehavior("TerminateSandbox"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*sandboxv1.TerminateSandboxResponse)), nil
	}

	if err := h.registry.GetError("TerminateSandbox"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("TerminateSandbox")
	if resp == nil {
		return connect.NewResponse(&sandboxv1.TerminateSandboxResponse{}), nil
	}
	return connect.NewResponse(resp.(*sandboxv1.TerminateSandboxResponse)), nil
}

// OnCreateSandbox configures the CreateSandbox RPC method.
func (s *MockServer) OnCreateSandbox() *MethodConfigBuilder[*sandboxv1.CreateSandboxResponse] {
	return &MethodConfigBuilder[*sandboxv1.CreateSandboxResponse]{
		methodName: "CreateSandbox",
		registry:   s.registry,
	}
}

// OnGetSandbox configures the GetSandbox RPC method.
func (s *MockServer) OnGetSandbox() *MethodConfigBuilder[*sandboxv1.GetSandboxResponse] {
	return &MethodConfigBuilder[*sandboxv1.GetSandboxResponse]{
		methodName: "GetSandbox",
		registry:   s.registry,
	}
}

// OnListSandboxes configures the ListSandboxes RPC method.
func (s *MockServer) OnListSandboxes() *MethodConfigBuilder[*sandboxv1.ListSandboxesResponse] {
	return &MethodConfigBuilder[*sandboxv1.ListSandboxesResponse]{
		methodName: "ListSandboxes",
		registry:   s.registry,
	}
}

// OnTerminateSandbox configures the TerminateSandbox RPC method.
func (s *MockServer) OnTerminateSandbox() *MethodConfigBuilder[*sandboxv1.TerminateSandboxResponse] {
	return &MethodConfigBuilder[*sandboxv1.TerminateSandboxResponse]{
		methodName: "TerminateSandbox",
		registry:   s.registry,
	}
}
