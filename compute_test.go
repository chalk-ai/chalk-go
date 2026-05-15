package chalk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	containerv1 "github.com/chalk-ai/chalk-go/gen/chalk/container/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/container/v1/containerv1connect"
	sandboxv1 "github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1/sandboxv1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImageBuilderProducesImageSpec(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "entrypoint.sh")
	require.NoError(t, os.WriteFile(filePath, []byte("#!/bin/sh\npython -m app\n"), 0o644))

	base := DebianSlimImage().PipInstall("requests")
	image, err := base.
		RunCommands("apt-get update").
		AddLocalFile(filePath, "/app/entrypoint.sh", FileMode(0o755))
	require.NoError(t, err)
	image = image.
		EnvVar("PYTHONDONTWRITEBYTECODE", "1").
		Workdir("/app").
		Entrypoint("/app/entrypoint.sh").
		Cmd("serve")

	require.Empty(t, base.toProto().GetEnv())
	require.Len(t, base.toProto().GetSteps(), 1)

	proto := image.toProto()
	require.Equal(t, "python:3.14-slim-trixie", proto.GetBaseImage())
	require.Equal(t, "/app", proto.GetWorkdir())
	require.Equal(t, []string{"/app/entrypoint.sh"}, proto.GetEntrypoint())
	require.Equal(t, []string{"serve"}, proto.GetCmd())
	require.Equal(t, "1", proto.GetEnv()["PYTHONDONTWRITEBYTECODE"])
	require.Len(t, proto.GetSteps(), 3)
	require.Equal(t, []string{"requests"}, proto.GetSteps()[0].GetPipInstall().GetPackages())
	require.Equal(t, []string{"apt-get update"}, proto.GetSteps()[1].GetRunCommands().GetCommands())
	require.Equal(t, "/app/entrypoint.sh", proto.GetSteps()[2].GetAddFile().GetDestination())
	require.Equal(t, uint32(0o755), proto.GetSteps()[2].GetAddFile().GetMode())
}

func TestContainerSpecUsesBuiltImageAndVolumes(t *testing.T) {
	container := NewContainerFromImageRef("worker", "registry.example.com/worker:sha").
		WithCPU("1").
		WithMemory("2Gi").
		WithGPU("nvidia-tesla-t4:1").
		WithEnv("LOG_LEVEL", "debug").
		WithPort(8080).
		WithLifetime(30 * time.Minute).
		WithVolume(EmptyDirVolume("scratch", "/scratch", "2Gi")).
		WithProtocol(ContainerProtocolHTTP).
		WithRouting(ContainerRoutingPrivate).
		WithAuthentication(ContainerAuthenticationAuthenticated)

	proto, err := container.toProto()
	require.NoError(t, err)
	require.Equal(t, "worker", proto.GetName())
	require.Equal(t, "registry.example.com/worker:sha", proto.GetImage())
	require.Equal(t, "1", proto.GetResources().GetCpu())
	require.Equal(t, "2Gi", proto.GetResources().GetMemory())
	require.Equal(t, "nvidia-tesla-t4:1", proto.GetResources().GetGpu())
	require.Equal(t, "debug", proto.GetEnvVars()["LOG_LEVEL"])
	require.Equal(t, int32(8080), proto.GetPort())
	require.Equal(t, int64(1800), proto.GetLifetime().GetSeconds())
	require.Len(t, proto.GetVolumes(), 1)
	require.Equal(t, "scratch", proto.GetVolumes()[0].GetName())
	require.Equal(t, "/scratch", proto.GetVolumes()[0].GetMountPath())
	require.Equal(t, "empty_dir", proto.GetVolumes()[0].GetType())
	require.Equal(t, "2Gi", proto.GetVolumes()[0].GetSizeLimit())
	require.Equal(t, "http", proto.GetProtocol())
	require.Equal(t, "PRIVATE", proto.GetRouting())
	require.Equal(t, "AUTHENTICATED", proto.GetAuthentication())
}

type computeAuthHandler struct {
	serverv1connect.UnimplementedAuthServiceHandler
}

func (h *computeAuthHandler) GetToken(
	_ context.Context,
	_ *connect.Request[serverv1.GetTokenRequest],
) (*connect.Response[serverv1.GetTokenResponse], error) {
	primaryEnvironment := "test-env"
	return connect.NewResponse(&serverv1.GetTokenResponse{
		AccessToken:        "test-token",
		TokenType:          "Bearer",
		ExpiresIn:          3600,
		ExpiresAt:          timestamppb.New(time.Now().Add(time.Hour)),
		PrimaryEnvironment: &primaryEnvironment,
		EnvironmentIdToName: map[string]string{
			"test-env": "test",
		},
	}), nil
}

type imageRPCHandler struct {
	sandboxv1connect.UnimplementedCustomImageServiceHandler
	mu               sync.Mutex
	submitted        *sandboxv1.GetOrBuildCustomImageRequest
	submittedHeaders http.Header
	getBuildIDs      []string
}

func (h *imageRPCHandler) GetOrBuildCustomImage(
	_ context.Context,
	req *connect.Request[sandboxv1.GetOrBuildCustomImageRequest],
) (*connect.Response[sandboxv1.GetOrBuildCustomImageResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.submitted = req.Msg
	h.submittedHeaders = req.Header().Clone()
	return connect.NewResponse(&sandboxv1.GetOrBuildCustomImageResponse{
		BuildId: "build-1",
		Status:  "building",
	}), nil
}

func (h *imageRPCHandler) GetCustomImage(
	_ context.Context,
	req *connect.Request[sandboxv1.GetCustomImageRequest],
) (*connect.Response[sandboxv1.GetCustomImageResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.getBuildIDs = append(h.getBuildIDs, req.Msg.GetBuildId())
	return connect.NewResponse(&sandboxv1.GetCustomImageResponse{
		BuildId: req.Msg.GetBuildId(),
		Status:  "succeeded",
		Image:   "registry.example.com/chalk/image:built",
	}), nil
}

type containerRPCHandler struct {
	containerv1connect.UnimplementedContainerServiceHandler
	mu         sync.Mutex
	runRequest *containerv1.RunContainerRequest
	headers    http.Header
}

func (h *containerRPCHandler) RunContainer(
	_ context.Context,
	req *connect.Request[containerv1.RunContainerRequest],
) (*connect.Response[containerv1.RunContainerResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.runRequest = req.Msg
	h.headers = req.Header().Clone()
	return connect.NewResponse(&containerv1.RunContainerResponse{
		Container: &containerv1.ContainerResponse{
			Id:     "container-1",
			Name:   req.Msg.GetContainer().GetSpec().GetName(),
			Status: "Running",
			Spec:   req.Msg.GetContainer().GetSpec(),
		},
	}), nil
}

func TestClientRunContainerBuildsImageAndSubmitsContainer(t *testing.T) {
	imageHandler := &imageRPCHandler{}
	containerHandler := &containerRPCHandler{}

	mux := http.NewServeMux()
	authPath, authHandler := serverv1connect.NewAuthServiceHandler(&computeAuthHandler{})
	mux.Handle(authPath, authHandler)
	imagePath, imageServiceHandler := sandboxv1connect.NewCustomImageServiceHandler(imageHandler)
	mux.Handle(imagePath, imageServiceHandler)
	containerPath, containerServiceHandler := containerv1connect.NewContainerServiceHandler(containerHandler)
	mux.Handle(containerPath, containerServiceHandler)

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient(context.Background(), &ClientConfig{
		ClientId:      "client-id",
		ClientSecret:  "client-secret",
		ApiServer:     server.URL,
		EnvironmentId: "test-env",
	})
	require.NoError(t, err)

	container := NewContainer(
		"worker",
		DebianSlimImage().PipInstall("requests").Workdir("/app"),
	).WithCPU("1").
		WithMemory("2Gi").
		WithVolume(SharedMemoryVolume("dshm", "/dev/shm", "1Gi"))

	info, err := client.RunContainer(
		context.Background(),
		container,
		WithRunContainerBuildOptions(WithImageBuildPollInterval(time.Millisecond)),
	)
	require.NoError(t, err)
	require.Equal(t, "container-1", info.ID)
	require.Equal(t, "Running", info.Status)

	require.NotNil(t, imageHandler.submitted)
	require.Equal(t, "python:3.14-slim-trixie", imageHandler.submitted.GetImageSpec().GetBaseImage())
	require.Equal(t, "Bearer test-token", imageHandler.submittedHeaders.Get("Authorization"))
	require.Equal(t, "go-api", imageHandler.submittedHeaders.Get("X-Chalk-Server"))
	require.Equal(t, "test-env", imageHandler.submittedHeaders.Get("X-Chalk-Env-Id"))
	require.Equal(t, []string{"build-1"}, imageHandler.getBuildIDs)

	require.NotNil(t, containerHandler.runRequest)
	spec := containerHandler.runRequest.GetContainer().GetSpec()
	require.Equal(t, "worker", spec.GetName())
	require.Equal(t, "registry.example.com/chalk/image:built", spec.GetImage())
	require.Equal(t, "1", spec.GetResources().GetCpu())
	require.Equal(t, "2Gi", spec.GetResources().GetMemory())
	require.Len(t, spec.GetVolumes(), 1)
	require.Equal(t, "shared_memory", spec.GetVolumes()[0].GetType())
	require.Equal(t, "Bearer test-token", containerHandler.headers.Get("Authorization"))
}
