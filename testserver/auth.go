package testserver

import (
	"context"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// authServiceHandler implements the AuthService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type authServiceHandler struct {
	serverv1connect.UnimplementedAuthServiceHandler
	registry *ResponseRegistry
}

// newAuthServiceHandler creates a new AuthService handler.
func newAuthServiceHandler(registry *ResponseRegistry) *authServiceHandler {
	return &authServiceHandler{registry: registry}
}

// GetToken implements the GetToken RPC method.
// By default, returns a mock token without validation.
func (h *authServiceHandler) GetToken(
	ctx context.Context,
	req *connect.Request[serverv1.GetTokenRequest],
) (*connect.Response[serverv1.GetTokenResponse], error) {
	// Capture request for test assertions
	h.registry.CaptureRequest("GetToken", req.Msg)

	// Check for custom behavior
	if behavior := h.registry.GetBehavior("GetToken"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetTokenResponse)), nil
	}

	// Check for configured error
	if err := h.registry.GetError("GetToken"); err != nil {
		return nil, err
	}

	// Return configured response or default mock token
	resp := h.registry.GetResponse("GetToken")
	if resp != nil {
		return connect.NewResponse(resp.(*serverv1.GetTokenResponse)), nil
	}

	// Default: return a mock token that passes client-side validation
	return connect.NewResponse(&serverv1.GetTokenResponse{
		AccessToken: "mock-test-token",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}), nil
}
