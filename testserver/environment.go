package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// environmentServiceHandler implements the EnvironmentService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type environmentServiceHandler struct {
	serverv1connect.UnimplementedEnvironmentServiceHandler
	registry *ResponseRegistry
}

// newEnvironmentServiceHandler creates a new EnvironmentService handler.
func newEnvironmentServiceHandler(registry *ResponseRegistry) *environmentServiceHandler {
	return &environmentServiceHandler{registry: registry}
}

// CreateEnvironmentV2 implements the CreateEnvironmentV2 RPC method.
func (h *environmentServiceHandler) CreateEnvironmentV2(
	ctx context.Context,
	req *connect.Request[serverv1.CreateEnvironmentV2Request],
) (*connect.Response[serverv1.CreateEnvironmentV2Response], error) {
	h.registry.CaptureRequest("CreateEnvironmentV2", req.Msg)

	if behavior := h.registry.GetBehavior("CreateEnvironmentV2"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateEnvironmentV2Response)), nil
	}

	if err := h.registry.GetError("CreateEnvironmentV2"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("CreateEnvironmentV2")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for CreateEnvironmentV2"))
	}

	return connect.NewResponse(resp.(*serverv1.CreateEnvironmentV2Response)), nil
}

// UpdateEnvironmentV2 implements the UpdateEnvironmentV2 RPC method.
func (h *environmentServiceHandler) UpdateEnvironmentV2(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateEnvironmentV2Request],
) (*connect.Response[serverv1.UpdateEnvironmentV2Response], error) {
	h.registry.CaptureRequest("UpdateEnvironmentV2", req.Msg)

	if behavior := h.registry.GetBehavior("UpdateEnvironmentV2"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateEnvironmentV2Response)), nil
	}

	if err := h.registry.GetError("UpdateEnvironmentV2"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("UpdateEnvironmentV2")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for UpdateEnvironmentV2"))
	}

	return connect.NewResponse(resp.(*serverv1.UpdateEnvironmentV2Response)), nil
}

// DeleteEnvironment implements the DeleteEnvironment RPC method.
func (h *environmentServiceHandler) DeleteEnvironment(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteEnvironmentRequest],
) (*connect.Response[serverv1.DeleteEnvironmentResponse], error) {
	h.registry.CaptureRequest("DeleteEnvironment", req.Msg)

	if behavior := h.registry.GetBehavior("DeleteEnvironment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteEnvironmentResponse)), nil
	}

	if err := h.registry.GetError("DeleteEnvironment"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("DeleteEnvironment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for DeleteEnvironment"))
	}

	return connect.NewResponse(resp.(*serverv1.DeleteEnvironmentResponse)), nil
}
