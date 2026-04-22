package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	scalinggroupv1 "github.com/chalk-ai/chalk-go/gen/chalk/scalinggroup/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/scalinggroup/v1/scalinggroupv1connect"
)

// scalingGroupServiceHandler implements the ScalingGroupManagerService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type scalingGroupServiceHandler struct {
	scalinggroupv1connect.UnimplementedScalingGroupManagerServiceHandler
	registry *ResponseRegistry
}

func newScalingGroupServiceHandler(registry *ResponseRegistry) *scalingGroupServiceHandler {
	return &scalingGroupServiceHandler{registry: registry}
}

func (h *scalingGroupServiceHandler) CreateScalingGroup(
	ctx context.Context,
	req *connect.Request[scalinggroupv1.CreateScalingGroupRequest],
) (*connect.Response[scalinggroupv1.CreateScalingGroupResponse], error) {
	h.registry.CaptureRequest("CreateScalingGroup", req.Msg)

	if behavior := h.registry.GetBehavior("CreateScalingGroup"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*scalinggroupv1.CreateScalingGroupResponse)), nil
	}

	if err := h.registry.GetError("CreateScalingGroup"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("CreateScalingGroup")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for CreateScalingGroup"))
	}
	return connect.NewResponse(resp.(*scalinggroupv1.CreateScalingGroupResponse)), nil
}

func (h *scalingGroupServiceHandler) GetScalingGroup(
	ctx context.Context,
	req *connect.Request[scalinggroupv1.GetScalingGroupRequest],
) (*connect.Response[scalinggroupv1.GetScalingGroupResponse], error) {
	h.registry.CaptureRequest("GetScalingGroup", req.Msg)

	if behavior := h.registry.GetBehavior("GetScalingGroup"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*scalinggroupv1.GetScalingGroupResponse)), nil
	}

	if err := h.registry.GetError("GetScalingGroup"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("GetScalingGroup")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for GetScalingGroup"))
	}
	return connect.NewResponse(resp.(*scalinggroupv1.GetScalingGroupResponse)), nil
}

func (h *scalingGroupServiceHandler) ListScalingGroups(
	ctx context.Context,
	req *connect.Request[scalinggroupv1.ListScalingGroupsRequest],
) (*connect.Response[scalinggroupv1.ListScalingGroupsResponse], error) {
	h.registry.CaptureRequest("ListScalingGroups", req.Msg)

	if behavior := h.registry.GetBehavior("ListScalingGroups"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*scalinggroupv1.ListScalingGroupsResponse)), nil
	}

	if err := h.registry.GetError("ListScalingGroups"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("ListScalingGroups")
	if resp == nil {
		return connect.NewResponse(&scalinggroupv1.ListScalingGroupsResponse{}), nil
	}
	return connect.NewResponse(resp.(*scalinggroupv1.ListScalingGroupsResponse)), nil
}

func (h *scalingGroupServiceHandler) DeleteScalingGroup(
	ctx context.Context,
	req *connect.Request[scalinggroupv1.DeleteScalingGroupRequest],
) (*connect.Response[scalinggroupv1.DeleteScalingGroupResponse], error) {
	h.registry.CaptureRequest("DeleteScalingGroup", req.Msg)

	if behavior := h.registry.GetBehavior("DeleteScalingGroup"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*scalinggroupv1.DeleteScalingGroupResponse)), nil
	}

	if err := h.registry.GetError("DeleteScalingGroup"); err != nil {
		return nil, err
	}

	resp := h.registry.GetResponse("DeleteScalingGroup")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound,
			errors.New("no mock response configured for DeleteScalingGroup"))
	}
	return connect.NewResponse(resp.(*scalinggroupv1.DeleteScalingGroupResponse)), nil
}

func (h *scalingGroupServiceHandler) BatchUpdateScalingGroupStatus(
	ctx context.Context,
	req *connect.Request[scalinggroupv1.BatchUpdateScalingGroupStatusRequest],
) (*connect.Response[scalinggroupv1.BatchUpdateScalingGroupStatusResponse], error) {
	h.registry.CaptureRequest("BatchUpdateScalingGroupStatus", req.Msg)
	return connect.NewResponse(&scalinggroupv1.BatchUpdateScalingGroupStatusResponse{}), nil
}

// OnCreateScalingGroup configures the CreateScalingGroup RPC method.
func (s *MockServer) OnCreateScalingGroup() *MethodConfigBuilder[*scalinggroupv1.CreateScalingGroupResponse] {
	return &MethodConfigBuilder[*scalinggroupv1.CreateScalingGroupResponse]{
		methodName: "CreateScalingGroup",
		registry:   s.registry,
	}
}

// OnGetScalingGroup configures the GetScalingGroup RPC method.
func (s *MockServer) OnGetScalingGroup() *MethodConfigBuilder[*scalinggroupv1.GetScalingGroupResponse] {
	return &MethodConfigBuilder[*scalinggroupv1.GetScalingGroupResponse]{
		methodName: "GetScalingGroup",
		registry:   s.registry,
	}
}

// OnDeleteScalingGroup configures the DeleteScalingGroup RPC method.
func (s *MockServer) OnDeleteScalingGroup() *MethodConfigBuilder[*scalinggroupv1.DeleteScalingGroupResponse] {
	return &MethodConfigBuilder[*scalinggroupv1.DeleteScalingGroupResponse]{
		methodName: "DeleteScalingGroup",
		registry:   s.registry,
	}
}
