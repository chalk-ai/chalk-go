package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// hostPoolServiceHandler implements the HostPoolService RPC handler using a
// ResponseRegistry to provide configurable mock responses.
type hostPoolServiceHandler struct {
	serverv1connect.UnimplementedHostPoolServiceHandler
	registry *ResponseRegistry
}

func newHostPoolServiceHandler(registry *ResponseRegistry) *hostPoolServiceHandler {
	return &hostPoolServiceHandler{registry: registry}
}

func (h *hostPoolServiceHandler) CreateEnvironmentHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.CreateEnvironmentHostPoolRequest],
) (*connect.Response[serverv1.CreateEnvironmentHostPoolResponse], error) {
	h.registry.CaptureRequest("CreateEnvironmentHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("CreateEnvironmentHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateEnvironmentHostPoolResponse)), nil
	}
	if err := h.registry.GetError("CreateEnvironmentHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateEnvironmentHostPool")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateEnvironmentHostPool"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateEnvironmentHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) UpdateEnvironmentHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateEnvironmentHostPoolRequest],
) (*connect.Response[serverv1.UpdateEnvironmentHostPoolResponse], error) {
	h.registry.CaptureRequest("UpdateEnvironmentHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("UpdateEnvironmentHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateEnvironmentHostPoolResponse)), nil
	}
	if err := h.registry.GetError("UpdateEnvironmentHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("UpdateEnvironmentHostPool")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for UpdateEnvironmentHostPool"))
	}
	return connect.NewResponse(resp.(*serverv1.UpdateEnvironmentHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) DeleteEnvironmentHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteEnvironmentHostPoolRequest],
) (*connect.Response[serverv1.DeleteEnvironmentHostPoolResponse], error) {
	h.registry.CaptureRequest("DeleteEnvironmentHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteEnvironmentHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteEnvironmentHostPoolResponse)), nil
	}
	if err := h.registry.GetError("DeleteEnvironmentHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteEnvironmentHostPool")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteEnvironmentHostPoolResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteEnvironmentHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) CreateClusterHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.CreateClusterHostPoolRequest],
) (*connect.Response[serverv1.CreateClusterHostPoolResponse], error) {
	h.registry.CaptureRequest("CreateClusterHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("CreateClusterHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateClusterHostPoolResponse)), nil
	}
	if err := h.registry.GetError("CreateClusterHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateClusterHostPool")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateClusterHostPool"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateClusterHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) UpdateClusterHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateClusterHostPoolRequest],
) (*connect.Response[serverv1.UpdateClusterHostPoolResponse], error) {
	h.registry.CaptureRequest("UpdateClusterHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("UpdateClusterHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateClusterHostPoolResponse)), nil
	}
	if err := h.registry.GetError("UpdateClusterHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("UpdateClusterHostPool")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for UpdateClusterHostPool"))
	}
	return connect.NewResponse(resp.(*serverv1.UpdateClusterHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) DeleteClusterHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteClusterHostPoolRequest],
) (*connect.Response[serverv1.DeleteClusterHostPoolResponse], error) {
	h.registry.CaptureRequest("DeleteClusterHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteClusterHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteClusterHostPoolResponse)), nil
	}
	if err := h.registry.GetError("DeleteClusterHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteClusterHostPool")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteClusterHostPoolResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteClusterHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) GetHostPool(
	ctx context.Context,
	req *connect.Request[serverv1.GetHostPoolRequest],
) (*connect.Response[serverv1.GetHostPoolResponse], error) {
	h.registry.CaptureRequest("GetHostPool", req.Msg)
	if behavior := h.registry.GetBehavior("GetHostPool"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetHostPoolResponse)), nil
	}
	if err := h.registry.GetError("GetHostPool"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetHostPool")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetHostPool"))
	}
	return connect.NewResponse(resp.(*serverv1.GetHostPoolResponse)), nil
}

func (h *hostPoolServiceHandler) ListHostPools(
	ctx context.Context,
	req *connect.Request[serverv1.ListHostPoolsRequest],
) (*connect.Response[serverv1.ListHostPoolsResponse], error) {
	h.registry.CaptureRequest("ListHostPools", req.Msg)
	if behavior := h.registry.GetBehavior("ListHostPools"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListHostPoolsResponse)), nil
	}
	if err := h.registry.GetError("ListHostPools"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("ListHostPools")
	if resp == nil {
		return connect.NewResponse(&serverv1.ListHostPoolsResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.ListHostPoolsResponse)), nil
}
