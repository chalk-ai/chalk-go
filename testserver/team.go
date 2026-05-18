package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// teamServiceHandler implements the TeamService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type teamServiceHandler struct {
	serverv1connect.UnimplementedTeamServiceHandler
	registry *ResponseRegistry
}

// newTeamServiceHandler creates a new TeamService handler.
func newTeamServiceHandler(registry *ResponseRegistry) *teamServiceHandler {
	return &teamServiceHandler{registry: registry}
}

// GetEnv implements the GetEnv RPC method.
func (h *teamServiceHandler) GetEnv(
	ctx context.Context,
	req *connect.Request[serverv1.GetEnvRequest],
) (*connect.Response[serverv1.GetEnvResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("GetEnv", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("GetEnv"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetEnvResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("GetEnv"); err != nil {
		return nil, err
	}

	// Return configured response
	resp := h.registry.GetResponse("GetEnv")
	if resp == nil {
		// Return default response with a test cluster ID
		clusterID := "test-cluster-id"
		return connect.NewResponse(&serverv1.GetEnvResponse{
			Environment: &serverv1.Environment{
				KubeClusterId: &clusterID,
			},
		}), nil
	}

	return connect.NewResponse(resp.(*serverv1.GetEnvResponse)), nil
}

// GetTeam implements the GetTeam RPC method.
func (h *teamServiceHandler) GetTeam(
	ctx context.Context,
	req *connect.Request[serverv1.GetTeamRequest],
) (*connect.Response[serverv1.GetTeamResponse], error) {
	h.registry.CaptureRequest("GetTeam", req.Msg)
	if behavior := h.registry.GetBehavior("GetTeam"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetTeamResponse)), nil
	}
	if err := h.registry.GetError("GetTeam"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetTeam")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetTeam"))
	}
	return connect.NewResponse(resp.(*serverv1.GetTeamResponse)), nil
}
