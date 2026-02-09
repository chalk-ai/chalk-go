package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// builderServiceHandler implements the BuilderService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type builderServiceHandler struct {
	serverv1connect.UnimplementedBuilderServiceHandler
	registry *ResponseRegistry
}

// newBuilderServiceHandler creates a new BuilderService handler.
func newBuilderServiceHandler(registry *ResponseRegistry) *builderServiceHandler {
	return &builderServiceHandler{registry: registry}
}

// GetClusterTimescaleDB implements the GetClusterTimescaleDB RPC method.
func (h *builderServiceHandler) GetClusterTimescaleDB(
	ctx context.Context,
	req *connect.Request[serverv1.GetClusterTimescaleDBRequest],
) (*connect.Response[serverv1.GetClusterTimescaleDBResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("GetClusterTimescaleDB", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("GetClusterTimescaleDB"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetClusterTimescaleDBResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("GetClusterTimescaleDB"); err != nil {
		return nil, err
	}

	// Return configured response
	resp := h.registry.GetResponse("GetClusterTimescaleDB")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for GetClusterTimescaleDB"))
	}

	return connect.NewResponse(resp.(*serverv1.GetClusterTimescaleDBResponse)), nil
}

// UpdateClusterTimescaleDB implements the UpdateClusterTimescaleDB RPC method.
func (h *builderServiceHandler) UpdateClusterTimescaleDB(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateClusterTimescaleDBRequest],
) (*connect.Response[serverv1.UpdateClusterTimescaleDBResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("UpdateClusterTimescaleDB", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("UpdateClusterTimescaleDB"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateClusterTimescaleDBResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("UpdateClusterTimescaleDB"); err != nil {
		return nil, err
	}

	// Return configured response
	resp := h.registry.GetResponse("UpdateClusterTimescaleDB")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for UpdateClusterTimescaleDB"))
	}

	return connect.NewResponse(resp.(*serverv1.UpdateClusterTimescaleDBResponse)), nil
}

// CreateClusterTimescaleDB implements the CreateClusterTimescaleDB RPC method.
func (h *builderServiceHandler) CreateClusterTimescaleDB(
	ctx context.Context,
	req *connect.Request[serverv1.CreateClusterTimescaleDBRequest],
) (*connect.Response[serverv1.CreateClusterTimescaleDBResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("CreateClusterTimescaleDB", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("CreateClusterTimescaleDB"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateClusterTimescaleDBResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("CreateClusterTimescaleDB"); err != nil {
		return nil, err
	}

	// Return configured response
	resp := h.registry.GetResponse("CreateClusterTimescaleDB")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for CreateClusterTimescaleDB"))
	}

	return connect.NewResponse(resp.(*serverv1.CreateClusterTimescaleDBResponse)), nil
}

// DeleteClusterTimescaleDB implements the DeleteClusterTimescaleDB RPC method.
func (h *builderServiceHandler) DeleteClusterTimescaleDB(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteClusterTimescaleDBRequest],
) (*connect.Response[serverv1.DeleteClusterTimescaleDBResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("DeleteClusterTimescaleDB", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("DeleteClusterTimescaleDB"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteClusterTimescaleDBResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("DeleteClusterTimescaleDB"); err != nil {
		return nil, err
	}

	// Return configured response
	resp := h.registry.GetResponse("DeleteClusterTimescaleDB")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for DeleteClusterTimescaleDB"))
	}

	return connect.NewResponse(resp.(*serverv1.DeleteClusterTimescaleDBResponse)), nil
}
