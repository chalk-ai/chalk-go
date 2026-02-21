package testserver

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

// cloudComponentsServiceHandler implements the CloudComponentsService RPC handler using
// a ResponseRegistry to provide configurable mock responses.
type cloudComponentsServiceHandler struct {
	serverv1connect.UnimplementedCloudComponentsServiceHandler
	registry *ResponseRegistry
}

// newCloudComponentsServiceHandler creates a new CloudComponentsService handler.
func newCloudComponentsServiceHandler(registry *ResponseRegistry) *cloudComponentsServiceHandler {
	return &cloudComponentsServiceHandler{registry: registry}
}

func (h *cloudComponentsServiceHandler) CreateBindingClusterGateway(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingClusterGatewayRequest],
) (*connect.Response[serverv1.CreateBindingClusterGatewayResponse], error) {
	h.registry.CaptureRequest("CreateBindingClusterGateway", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingClusterGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingClusterGatewayResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingClusterGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingClusterGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingClusterGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingClusterGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingClusterGateway(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingClusterGatewayRequest],
) (*connect.Response[serverv1.GetBindingClusterGatewayResponse], error) {
	h.registry.CaptureRequest("GetBindingClusterGateway", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingClusterGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingClusterGatewayResponse)), nil
	}
	if err := h.registry.GetError("GetBindingClusterGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingClusterGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingClusterGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingClusterGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingClusterGateway(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingClusterGatewayRequest],
) (*connect.Response[serverv1.DeleteBindingClusterGatewayResponse], error) {
	h.registry.CaptureRequest("DeleteBindingClusterGateway", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingClusterGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterGatewayResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingClusterGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingClusterGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingClusterGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingPrivateGateway(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingPrivateGatewayRequest],
) (*connect.Response[serverv1.CreateBindingPrivateGatewayResponse], error) {
	h.registry.CaptureRequest("CreateBindingPrivateGateway", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingPrivateGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingPrivateGatewayResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingPrivateGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingPrivateGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingPrivateGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingPrivateGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingPrivateGateway(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingPrivateGatewayRequest],
) (*connect.Response[serverv1.GetBindingPrivateGatewayResponse], error) {
	h.registry.CaptureRequest("GetBindingPrivateGateway", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingPrivateGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingPrivateGatewayResponse)), nil
	}
	if err := h.registry.GetError("GetBindingPrivateGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingPrivateGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingPrivateGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingPrivateGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingPrivateGateway(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingPrivateGatewayRequest],
) (*connect.Response[serverv1.DeleteBindingPrivateGatewayResponse], error) {
	h.registry.CaptureRequest("DeleteBindingPrivateGateway", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingPrivateGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingPrivateGatewayResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingPrivateGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingPrivateGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingPrivateGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingPrivateGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingClusterBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingClusterBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.CreateBindingClusterBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("CreateBindingClusterBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingClusterBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingClusterBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingClusterBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingClusterBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingClusterBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingClusterBackgroundPersistenceDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingClusterBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingClusterBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.GetBindingClusterBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("GetBindingClusterBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingClusterBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingClusterBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("GetBindingClusterBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingClusterBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingClusterBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingClusterBackgroundPersistenceDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingClusterBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingClusterBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.DeleteBindingClusterBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("DeleteBindingClusterBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingClusterBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingClusterBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingClusterBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingClusterBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterBackgroundPersistenceDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingClusterTelemetryDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingClusterTelemetryDeploymentRequest],
) (*connect.Response[serverv1.CreateBindingClusterTelemetryDeploymentResponse], error) {
	h.registry.CaptureRequest("CreateBindingClusterTelemetryDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingClusterTelemetryDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingClusterTelemetryDeploymentResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingClusterTelemetryDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingClusterTelemetryDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingClusterTelemetryDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingClusterTelemetryDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingClusterTelemetryDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingClusterTelemetryDeploymentRequest],
) (*connect.Response[serverv1.GetBindingClusterTelemetryDeploymentResponse], error) {
	h.registry.CaptureRequest("GetBindingClusterTelemetryDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingClusterTelemetryDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingClusterTelemetryDeploymentResponse)), nil
	}
	if err := h.registry.GetError("GetBindingClusterTelemetryDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingClusterTelemetryDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingClusterTelemetryDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingClusterTelemetryDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingClusterTelemetryDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingClusterTelemetryDeploymentRequest],
) (*connect.Response[serverv1.DeleteBindingClusterTelemetryDeploymentResponse], error) {
	h.registry.CaptureRequest("DeleteBindingClusterTelemetryDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingClusterTelemetryDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterTelemetryDeploymentResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingClusterTelemetryDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingClusterTelemetryDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingClusterTelemetryDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterTelemetryDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingEnvironmentGateway(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingEnvironmentGatewayRequest],
) (*connect.Response[serverv1.CreateBindingEnvironmentGatewayResponse], error) {
	h.registry.CaptureRequest("CreateBindingEnvironmentGateway", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingEnvironmentGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentGatewayResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingEnvironmentGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingEnvironmentGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingEnvironmentGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingEnvironmentGateway(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingEnvironmentGatewayRequest],
) (*connect.Response[serverv1.GetBindingEnvironmentGatewayResponse], error) {
	h.registry.CaptureRequest("GetBindingEnvironmentGateway", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingEnvironmentGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentGatewayResponse)), nil
	}
	if err := h.registry.GetError("GetBindingEnvironmentGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingEnvironmentGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingEnvironmentGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingEnvironmentGateway(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingEnvironmentGatewayRequest],
) (*connect.Response[serverv1.DeleteBindingEnvironmentGatewayResponse], error) {
	h.registry.CaptureRequest("DeleteBindingEnvironmentGateway", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingEnvironmentGateway"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentGatewayResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingEnvironmentGateway"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingEnvironmentGateway")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingEnvironmentGateway"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentGatewayResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingEnvironmentBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingEnvironmentBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.CreateBindingEnvironmentBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("CreateBindingEnvironmentBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingEnvironmentBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingEnvironmentBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingEnvironmentBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingEnvironmentBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingEnvironmentBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingEnvironmentBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.GetBindingEnvironmentBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("GetBindingEnvironmentBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingEnvironmentBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("GetBindingEnvironmentBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingEnvironmentBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingEnvironmentBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingEnvironmentBackgroundPersistenceDeployment(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingEnvironmentBackgroundPersistenceDeploymentRequest],
) (*connect.Response[serverv1.DeleteBindingEnvironmentBackgroundPersistenceDeploymentResponse], error) {
	h.registry.CaptureRequest("DeleteBindingEnvironmentBackgroundPersistenceDeployment", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingEnvironmentBackgroundPersistenceDeployment"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingEnvironmentBackgroundPersistenceDeployment"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingEnvironmentBackgroundPersistenceDeployment")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for DeleteBindingEnvironmentBackgroundPersistenceDeployment"))
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentBackgroundPersistenceDeploymentResponse)), nil
}
