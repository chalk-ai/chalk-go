package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// integrationsServiceHandler implements the IntegrationsService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type integrationsServiceHandler struct {
	serverv1connect.UnimplementedIntegrationsServiceHandler
	registry *ResponseRegistry
}

// newIntegrationsServiceHandler creates a new IntegrationsService handler.
func newIntegrationsServiceHandler(registry *ResponseRegistry) *integrationsServiceHandler {
	return &integrationsServiceHandler{registry: registry}
}

// InsertIntegration implements the InsertIntegration RPC method.
func (h *integrationsServiceHandler) InsertIntegration(
	ctx context.Context,
	req *connect.Request[serverv1.InsertIntegrationRequest],
) (*connect.Response[serverv1.InsertIntegrationResponse], error) {
	h.registry.CaptureRequest("InsertIntegration", req.Msg)

	if behavior := h.registry.GetBehavior("InsertIntegration"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.InsertIntegrationResponse)), nil
	}

	if err := h.registry.GetError("InsertIntegration"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("InsertIntegration")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for InsertIntegration"))
	}

	return connect.NewResponse(resp.(*serverv1.InsertIntegrationResponse)), nil
}

// GetIntegration implements the GetIntegration RPC method.
func (h *integrationsServiceHandler) GetIntegration(
	ctx context.Context,
	req *connect.Request[serverv1.GetIntegrationRequest],
) (*connect.Response[serverv1.GetIntegrationResponse], error) {
	h.registry.CaptureRequest("GetIntegration", req.Msg)

	if behavior := h.registry.GetBehavior("GetIntegration"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetIntegrationResponse)), nil
	}

	if err := h.registry.GetError("GetIntegration"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("GetIntegration")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for GetIntegration"))
	}

	return connect.NewResponse(resp.(*serverv1.GetIntegrationResponse)), nil
}

// GetIntegrationValue implements the GetIntegrationValue RPC method.
func (h *integrationsServiceHandler) GetIntegrationValue(
	ctx context.Context,
	req *connect.Request[serverv1.GetIntegrationValueRequest],
) (*connect.Response[serverv1.GetIntegrationValueResponse], error) {
	h.registry.CaptureRequest("GetIntegrationValue", req.Msg)

	if behavior := h.registry.GetBehavior("GetIntegrationValue"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetIntegrationValueResponse)), nil
	}

	if err := h.registry.GetError("GetIntegrationValue"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("GetIntegrationValue")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for GetIntegrationValue"))
	}

	return connect.NewResponse(resp.(*serverv1.GetIntegrationValueResponse)), nil
}

// UpdateIntegration implements the UpdateIntegration RPC method.
func (h *integrationsServiceHandler) UpdateIntegration(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateIntegrationRequest],
) (*connect.Response[serverv1.UpdateIntegrationResponse], error) {
	h.registry.CaptureRequest("UpdateIntegration", req.Msg)

	if behavior := h.registry.GetBehavior("UpdateIntegration"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateIntegrationResponse)), nil
	}

	if err := h.registry.GetError("UpdateIntegration"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("UpdateIntegration")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for UpdateIntegration"))
	}

	return connect.NewResponse(resp.(*serverv1.UpdateIntegrationResponse)), nil
}

// DeleteIntegration implements the DeleteIntegration RPC method.
func (h *integrationsServiceHandler) DeleteIntegration(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteIntegrationRequest],
) (*connect.Response[serverv1.DeleteIntegrationResponse], error) {
	h.registry.CaptureRequest("DeleteIntegration", req.Msg)

	if behavior := h.registry.GetBehavior("DeleteIntegration"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteIntegrationResponse)), nil
	}

	if err := h.registry.GetError("DeleteIntegration"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("DeleteIntegration")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for DeleteIntegration"))
	}

	return connect.NewResponse(resp.(*serverv1.DeleteIntegrationResponse)), nil
}

// ListIntegrations implements the ListIntegrations RPC method.
func (h *integrationsServiceHandler) ListIntegrations(
	ctx context.Context,
	req *connect.Request[serverv1.ListIntegrationsRequest],
) (*connect.Response[serverv1.ListIntegrationsResponse], error) {
	h.registry.CaptureRequest("ListIntegrations", req.Msg)

	if behavior := h.registry.GetBehavior("ListIntegrations"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListIntegrationsResponse)), nil
	}

	if err := h.registry.GetError("ListIntegrations"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("ListIntegrations")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for ListIntegrations"))
	}

	return connect.NewResponse(resp.(*serverv1.ListIntegrationsResponse)), nil
}
