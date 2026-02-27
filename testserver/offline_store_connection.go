package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// offlineStoreConnectionServiceHandler implements the OfflineStoreConnectionService RPC handler
// using a ResponseRegistry to provide configurable mock responses.
type offlineStoreConnectionServiceHandler struct {
	serverv1connect.UnimplementedOfflineStoreConnectionServiceHandler
	registry *ResponseRegistry
}

// newOfflineStoreConnectionServiceHandler creates a new OfflineStoreConnectionService handler.
func newOfflineStoreConnectionServiceHandler(registry *ResponseRegistry) *offlineStoreConnectionServiceHandler {
	return &offlineStoreConnectionServiceHandler{registry: registry}
}

func (h *offlineStoreConnectionServiceHandler) CreateOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.CreateOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.CreateOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("CreateOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("CreateOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("CreateOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) GetOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.GetOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.GetOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("GetOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("GetOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("GetOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.GetOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) ListOfflineStoreConnections(
	ctx context.Context,
	req *connect.Request[serverv1.ListOfflineStoreConnectionsRequest],
) (*connect.Response[serverv1.ListOfflineStoreConnectionsResponse], error) {
	h.registry.CaptureRequest("ListOfflineStoreConnections", req.Msg)
	if behavior := h.registry.GetBehavior("ListOfflineStoreConnections"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListOfflineStoreConnectionsResponse)), nil
	}
	if err := h.registry.GetError("ListOfflineStoreConnections"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("ListOfflineStoreConnections")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for ListOfflineStoreConnections"))
	}
	return connect.NewResponse(resp.(*serverv1.ListOfflineStoreConnectionsResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) UpdateOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.UpdateOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("UpdateOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("UpdateOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("UpdateOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("UpdateOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for UpdateOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.UpdateOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) DeleteOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.DeleteOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("DeleteOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("DeleteOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) TestOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.TestOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.TestOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("TestOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("TestOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.TestOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("TestOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("TestOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for TestOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.TestOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) CreateBindingEnvironmentOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingEnvironmentOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.CreateBindingEnvironmentOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("CreateBindingEnvironmentOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingEnvironmentOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingEnvironmentOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingEnvironmentOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingEnvironmentOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) GetBindingEnvironmentOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingEnvironmentOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.GetBindingEnvironmentOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("GetBindingEnvironmentOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingEnvironmentOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("GetBindingEnvironmentOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingEnvironmentOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingEnvironmentOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) DeleteBindingEnvironmentOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingEnvironmentOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.DeleteBindingEnvironmentOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("DeleteBindingEnvironmentOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingEnvironmentOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingEnvironmentOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingEnvironmentOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingEnvironmentOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentOfflineStoreConnectionResponse)), nil
}

func (h *offlineStoreConnectionServiceHandler) MigrateOfflineStoreConnection(
	ctx context.Context,
	req *connect.Request[serverv1.MigrateOfflineStoreConnectionRequest],
) (*connect.Response[serverv1.MigrateOfflineStoreConnectionResponse], error) {
	h.registry.CaptureRequest("MigrateOfflineStoreConnection", req.Msg)
	if behavior := h.registry.GetBehavior("MigrateOfflineStoreConnection"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.MigrateOfflineStoreConnectionResponse)), nil
	}
	if err := h.registry.GetError("MigrateOfflineStoreConnection"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("MigrateOfflineStoreConnection")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for MigrateOfflineStoreConnection"))
	}
	return connect.NewResponse(resp.(*serverv1.MigrateOfflineStoreConnectionResponse)), nil
}