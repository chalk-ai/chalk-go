package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// graphServiceHandler implements the GraphService RPC handler
// using a ResponseRegistry to provide configurable mock responses.
type graphServiceHandler struct {
	serverv1connect.UnimplementedGraphServiceHandler
	registry *ResponseRegistry
}

// newGraphServiceHandler creates a new GraphService handler.
func newGraphServiceHandler(registry *ResponseRegistry) *graphServiceHandler {
	return &graphServiceHandler{registry: registry}
}

func (h *graphServiceHandler) GetOfflineStoreTable(
	ctx context.Context,
	req *connect.Request[serverv1.GetOfflineStoreTableRequest],
) (*connect.Response[serverv1.GetOfflineStoreTableResponse], error) {
	h.registry.CaptureRequest("GetOfflineStoreTable", req.Msg)
	if behavior := h.registry.GetBehavior("GetOfflineStoreTable"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetOfflineStoreTableResponse)), nil
	}
	if err := h.registry.GetError("GetOfflineStoreTable"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetOfflineStoreTable")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetOfflineStoreTable"))
	}
	return connect.NewResponse(resp.(*serverv1.GetOfflineStoreTableResponse)), nil
}
