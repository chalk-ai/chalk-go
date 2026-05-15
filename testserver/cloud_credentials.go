package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// cloudAccountCredentialsServiceHandler implements the CloudAccountCredentialsService RPC
// handler using a ResponseRegistry to provide configurable mock responses.
type cloudAccountCredentialsServiceHandler struct {
	serverv1connect.UnimplementedCloudAccountCredentialsServiceHandler
	registry *ResponseRegistry
}

func newCloudAccountCredentialsServiceHandler(registry *ResponseRegistry) *cloudAccountCredentialsServiceHandler {
	return &cloudAccountCredentialsServiceHandler{registry: registry}
}

func (h *cloudAccountCredentialsServiceHandler) GetCloudCredentials(
	ctx context.Context,
	req *connect.Request[serverv1.GetCloudCredentialsRequest],
) (*connect.Response[serverv1.GetCloudCredentialsResponse], error) {
	h.registry.CaptureRequest("GetCloudCredentials", req.Msg)
	if behavior := h.registry.GetBehavior("GetCloudCredentials"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetCloudCredentialsResponse)), nil
	}
	if err := h.registry.GetError("GetCloudCredentials"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetCloudCredentials")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetCloudCredentials"))
	}
	return connect.NewResponse(resp.(*serverv1.GetCloudCredentialsResponse)), nil
}
