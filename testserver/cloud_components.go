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

func (h *cloudComponentsServiceHandler) CreateCloudComponentVpc(
	ctx context.Context,
	req *connect.Request[serverv1.CreateCloudComponentVpcRequest],
) (*connect.Response[serverv1.CreateCloudComponentVpcResponse], error) {
	h.registry.CaptureRequest("CreateCloudComponentVpc", req.Msg)
	if behavior := h.registry.GetBehavior("CreateCloudComponentVpc"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateCloudComponentVpcResponse)), nil
	}
	if err := h.registry.GetError("CreateCloudComponentVpc"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateCloudComponentVpc")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateCloudComponentVpc"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateCloudComponentVpcResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetCloudComponentVpc(
	ctx context.Context,
	req *connect.Request[serverv1.GetCloudComponentVpcRequest],
) (*connect.Response[serverv1.GetCloudComponentVpcResponse], error) {
	h.registry.CaptureRequest("GetCloudComponentVpc", req.Msg)
	if behavior := h.registry.GetBehavior("GetCloudComponentVpc"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetCloudComponentVpcResponse)), nil
	}
	if err := h.registry.GetError("GetCloudComponentVpc"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetCloudComponentVpc")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetCloudComponentVpc"))
	}
	return connect.NewResponse(resp.(*serverv1.GetCloudComponentVpcResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteCloudComponentVpc(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteCloudComponentVpcRequest],
) (*connect.Response[serverv1.DeleteCloudComponentVpcResponse], error) {
	h.registry.CaptureRequest("DeleteCloudComponentVpc", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteCloudComponentVpc"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentVpcResponse)), nil
	}
	if err := h.registry.GetError("DeleteCloudComponentVpc"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteCloudComponentVpc")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteCloudComponentVpcResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentVpcResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateCloudComponentCluster(
	ctx context.Context,
	req *connect.Request[serverv1.CreateCloudComponentClusterRequest],
) (*connect.Response[serverv1.CreateCloudComponentClusterResponse], error) {
	h.registry.CaptureRequest("CreateCloudComponentCluster", req.Msg)
	if behavior := h.registry.GetBehavior("CreateCloudComponentCluster"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateCloudComponentClusterResponse)), nil
	}
	if err := h.registry.GetError("CreateCloudComponentCluster"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateCloudComponentCluster")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateCloudComponentCluster"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateCloudComponentClusterResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetCloudComponentCluster(
	ctx context.Context,
	req *connect.Request[serverv1.GetCloudComponentClusterRequest],
) (*connect.Response[serverv1.GetCloudComponentClusterResponse], error) {
	h.registry.CaptureRequest("GetCloudComponentCluster", req.Msg)
	if behavior := h.registry.GetBehavior("GetCloudComponentCluster"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetCloudComponentClusterResponse)), nil
	}
	if err := h.registry.GetError("GetCloudComponentCluster"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetCloudComponentCluster")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetCloudComponentCluster"))
	}
	return connect.NewResponse(resp.(*serverv1.GetCloudComponentClusterResponse)), nil
}

func (h *cloudComponentsServiceHandler) UpdateCloudComponentCluster(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateCloudComponentClusterRequest],
) (*connect.Response[serverv1.UpdateCloudComponentClusterResponse], error) {
	h.registry.CaptureRequest("UpdateCloudComponentCluster", req.Msg)
	if behavior := h.registry.GetBehavior("UpdateCloudComponentCluster"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.UpdateCloudComponentClusterResponse)), nil
	}
	if err := h.registry.GetError("UpdateCloudComponentCluster"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("UpdateCloudComponentCluster")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for UpdateCloudComponentCluster"))
	}
	return connect.NewResponse(resp.(*serverv1.UpdateCloudComponentClusterResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteCloudComponentCluster(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteCloudComponentClusterRequest],
) (*connect.Response[serverv1.DeleteCloudComponentClusterResponse], error) {
	h.registry.CaptureRequest("DeleteCloudComponentCluster", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteCloudComponentCluster"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentClusterResponse)), nil
	}
	if err := h.registry.GetError("DeleteCloudComponentCluster"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteCloudComponentCluster")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteCloudComponentClusterResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentClusterResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateCloudComponentStorage(
	ctx context.Context,
	req *connect.Request[serverv1.CreateCloudComponentStorageRequest],
) (*connect.Response[serverv1.CreateCloudComponentStorageResponse], error) {
	h.registry.CaptureRequest("CreateCloudComponentStorage", req.Msg)
	if behavior := h.registry.GetBehavior("CreateCloudComponentStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateCloudComponentStorageResponse)), nil
	}
	if err := h.registry.GetError("CreateCloudComponentStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateCloudComponentStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateCloudComponentStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateCloudComponentStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetCloudComponentStorage(
	ctx context.Context,
	req *connect.Request[serverv1.GetCloudComponentStorageRequest],
) (*connect.Response[serverv1.GetCloudComponentStorageResponse], error) {
	h.registry.CaptureRequest("GetCloudComponentStorage", req.Msg)
	if behavior := h.registry.GetBehavior("GetCloudComponentStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetCloudComponentStorageResponse)), nil
	}
	if err := h.registry.GetError("GetCloudComponentStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetCloudComponentStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetCloudComponentStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.GetCloudComponentStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteCloudComponentStorage(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteCloudComponentStorageRequest],
) (*connect.Response[serverv1.DeleteCloudComponentStorageResponse], error) {
	h.registry.CaptureRequest("DeleteCloudComponentStorage", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteCloudComponentStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentStorageResponse)), nil
	}
	if err := h.registry.GetError("DeleteCloudComponentStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteCloudComponentStorage")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteCloudComponentStorageResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) ListCloudComponentStorage(
	ctx context.Context,
	req *connect.Request[serverv1.ListCloudComponentStorageRequest],
) (*connect.Response[serverv1.ListCloudComponentStorageResponse], error) {
	h.registry.CaptureRequest("ListCloudComponentStorage", req.Msg)
	if behavior := h.registry.GetBehavior("ListCloudComponentStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListCloudComponentStorageResponse)), nil
	}
	if err := h.registry.GetError("ListCloudComponentStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("ListCloudComponentStorage")
	if resp == nil {
		// List has a repeated response; default to empty when unconfigured.
		return connect.NewResponse(&serverv1.ListCloudComponentStorageResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.ListCloudComponentStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingEnvironmentCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingEnvironmentCloudStorageRequest],
) (*connect.Response[serverv1.CreateBindingEnvironmentCloudStorageResponse], error) {
	h.registry.CaptureRequest("CreateBindingEnvironmentCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingEnvironmentCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingEnvironmentCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingEnvironmentCloudStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingEnvironmentCloudStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingEnvironmentCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingEnvironmentCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingEnvironmentCloudStorageRequest],
) (*connect.Response[serverv1.GetBindingEnvironmentCloudStorageResponse], error) {
	h.registry.CaptureRequest("GetBindingEnvironmentCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingEnvironmentCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("GetBindingEnvironmentCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingEnvironmentCloudStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingEnvironmentCloudStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingEnvironmentCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingEnvironmentCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingEnvironmentCloudStorageRequest],
) (*connect.Response[serverv1.DeleteBindingEnvironmentCloudStorageResponse], error) {
	h.registry.CaptureRequest("DeleteBindingEnvironmentCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingEnvironmentCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingEnvironmentCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingEnvironmentCloudStorage")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteBindingEnvironmentCloudStorageResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingEnvironmentCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingClusterCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingClusterCloudStorageRequest],
) (*connect.Response[serverv1.CreateBindingClusterCloudStorageResponse], error) {
	h.registry.CaptureRequest("CreateBindingClusterCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingClusterCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingClusterCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingClusterCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingClusterCloudStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingClusterCloudStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingClusterCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingClusterCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingClusterCloudStorageRequest],
) (*connect.Response[serverv1.GetBindingClusterCloudStorageResponse], error) {
	h.registry.CaptureRequest("GetBindingClusterCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingClusterCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingClusterCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("GetBindingClusterCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingClusterCloudStorage")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingClusterCloudStorage"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingClusterCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingClusterCloudStorage(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingClusterCloudStorageRequest],
) (*connect.Response[serverv1.DeleteBindingClusterCloudStorageResponse], error) {
	h.registry.CaptureRequest("DeleteBindingClusterCloudStorage", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingClusterCloudStorage"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterCloudStorageResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingClusterCloudStorage"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingClusterCloudStorage")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteBindingClusterCloudStorageResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterCloudStorageResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateCloudComponentContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.CreateCloudComponentContainerRegistryRequest],
) (*connect.Response[serverv1.CreateCloudComponentContainerRegistryResponse], error) {
	h.registry.CaptureRequest("CreateCloudComponentContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("CreateCloudComponentContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateCloudComponentContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("CreateCloudComponentContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateCloudComponentContainerRegistry")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateCloudComponentContainerRegistry"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateCloudComponentContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetCloudComponentContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.GetCloudComponentContainerRegistryRequest],
) (*connect.Response[serverv1.GetCloudComponentContainerRegistryResponse], error) {
	h.registry.CaptureRequest("GetCloudComponentContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("GetCloudComponentContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetCloudComponentContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("GetCloudComponentContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetCloudComponentContainerRegistry")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetCloudComponentContainerRegistry"))
	}
	return connect.NewResponse(resp.(*serverv1.GetCloudComponentContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) ListCloudComponentContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.ListCloudComponentContainerRegistryRequest],
) (*connect.Response[serverv1.ListCloudComponentContainerRegistryResponse], error) {
	h.registry.CaptureRequest("ListCloudComponentContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("ListCloudComponentContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListCloudComponentContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("ListCloudComponentContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("ListCloudComponentContainerRegistry")
	if resp == nil {
		// List has a repeated response; default to empty when unconfigured.
		return connect.NewResponse(&serverv1.ListCloudComponentContainerRegistryResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.ListCloudComponentContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteCloudComponentContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteCloudComponentContainerRegistryRequest],
) (*connect.Response[serverv1.DeleteCloudComponentContainerRegistryResponse], error) {
	h.registry.CaptureRequest("DeleteCloudComponentContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteCloudComponentContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("DeleteCloudComponentContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteCloudComponentContainerRegistry")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteCloudComponentContainerRegistryResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteCloudComponentContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) CreateBindingClusterContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.CreateBindingClusterContainerRegistryRequest],
) (*connect.Response[serverv1.CreateBindingClusterContainerRegistryResponse], error) {
	h.registry.CaptureRequest("CreateBindingClusterContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("CreateBindingClusterContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.CreateBindingClusterContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("CreateBindingClusterContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("CreateBindingClusterContainerRegistry")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for CreateBindingClusterContainerRegistry"))
	}
	return connect.NewResponse(resp.(*serverv1.CreateBindingClusterContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) GetBindingClusterContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.GetBindingClusterContainerRegistryRequest],
) (*connect.Response[serverv1.GetBindingClusterContainerRegistryResponse], error) {
	h.registry.CaptureRequest("GetBindingClusterContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("GetBindingClusterContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.GetBindingClusterContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("GetBindingClusterContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("GetBindingClusterContainerRegistry")
	if resp == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("no mock response configured for GetBindingClusterContainerRegistry"))
	}
	return connect.NewResponse(resp.(*serverv1.GetBindingClusterContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) ListBindingClusterContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.ListBindingClusterContainerRegistryRequest],
) (*connect.Response[serverv1.ListBindingClusterContainerRegistryResponse], error) {
	h.registry.CaptureRequest("ListBindingClusterContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("ListBindingClusterContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.ListBindingClusterContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("ListBindingClusterContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("ListBindingClusterContainerRegistry")
	if resp == nil {
		// List has a repeated response; default to empty when unconfigured.
		return connect.NewResponse(&serverv1.ListBindingClusterContainerRegistryResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.ListBindingClusterContainerRegistryResponse)), nil
}

func (h *cloudComponentsServiceHandler) DeleteBindingClusterContainerRegistry(
	ctx context.Context,
	req *connect.Request[serverv1.DeleteBindingClusterContainerRegistryRequest],
) (*connect.Response[serverv1.DeleteBindingClusterContainerRegistryResponse], error) {
	h.registry.CaptureRequest("DeleteBindingClusterContainerRegistry", req.Msg)
	if behavior := h.registry.GetBehavior("DeleteBindingClusterContainerRegistry"); behavior != nil {
		resp, err := behavior(req.Msg)
		if err != nil {
			return nil, err
		}
		return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterContainerRegistryResponse)), nil
	}
	if err := h.registry.GetError("DeleteBindingClusterContainerRegistry"); err != nil {
		return nil, err
	}
	resp := h.registry.GetResponse("DeleteBindingClusterContainerRegistry")
	if resp == nil {
		// Delete has an empty response; default to success when unconfigured.
		return connect.NewResponse(&serverv1.DeleteBindingClusterContainerRegistryResponse{}), nil
	}
	return connect.NewResponse(resp.(*serverv1.DeleteBindingClusterContainerRegistryResponse)), nil
}

// OnCreateCloudComponentVpc configures the CreateCloudComponentVpc RPC method.
func (s *MockServer) OnCreateCloudComponentVpc() *MethodConfigBuilder[*serverv1.CreateCloudComponentVpcResponse] {
	return &MethodConfigBuilder[*serverv1.CreateCloudComponentVpcResponse]{
		methodName: "CreateCloudComponentVpc",
		registry:   s.registry,
	}
}

// OnGetCloudComponentVpc configures the GetCloudComponentVpc RPC method.
func (s *MockServer) OnGetCloudComponentVpc() *MethodConfigBuilder[*serverv1.GetCloudComponentVpcResponse] {
	return &MethodConfigBuilder[*serverv1.GetCloudComponentVpcResponse]{
		methodName: "GetCloudComponentVpc",
		registry:   s.registry,
	}
}

// OnDeleteCloudComponentVpc configures the DeleteCloudComponentVpc RPC method.
func (s *MockServer) OnDeleteCloudComponentVpc() *MethodConfigBuilder[*serverv1.DeleteCloudComponentVpcResponse] {
	return &MethodConfigBuilder[*serverv1.DeleteCloudComponentVpcResponse]{
		methodName: "DeleteCloudComponentVpc",
		registry:   s.registry,
	}
}

// OnCreateCloudComponentCluster configures the CreateCloudComponentCluster RPC method.
func (s *MockServer) OnCreateCloudComponentCluster() *MethodConfigBuilder[*serverv1.CreateCloudComponentClusterResponse] {
	return &MethodConfigBuilder[*serverv1.CreateCloudComponentClusterResponse]{
		methodName: "CreateCloudComponentCluster",
		registry:   s.registry,
	}
}

// OnGetCloudComponentCluster configures the GetCloudComponentCluster RPC method.
func (s *MockServer) OnGetCloudComponentCluster() *MethodConfigBuilder[*serverv1.GetCloudComponentClusterResponse] {
	return &MethodConfigBuilder[*serverv1.GetCloudComponentClusterResponse]{
		methodName: "GetCloudComponentCluster",
		registry:   s.registry,
	}
}

// OnUpdateCloudComponentCluster configures the UpdateCloudComponentCluster RPC method.
func (s *MockServer) OnUpdateCloudComponentCluster() *MethodConfigBuilder[*serverv1.UpdateCloudComponentClusterResponse] {
	return &MethodConfigBuilder[*serverv1.UpdateCloudComponentClusterResponse]{
		methodName: "UpdateCloudComponentCluster",
		registry:   s.registry,
	}
}

// OnDeleteCloudComponentCluster configures the DeleteCloudComponentCluster RPC method.
func (s *MockServer) OnDeleteCloudComponentCluster() *MethodConfigBuilder[*serverv1.DeleteCloudComponentClusterResponse] {
	return &MethodConfigBuilder[*serverv1.DeleteCloudComponentClusterResponse]{
		methodName: "DeleteCloudComponentCluster",
		registry:   s.registry,
	}
}
