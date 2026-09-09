package chalk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/auth"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authCaptureTransport struct {
	request *http.Request
}

func (t *authCaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		Request:    req,
	}, nil
}

func staleAuthToken() *serverv1.GetTokenResponse {
	return &serverv1.GetTokenResponse{
		AccessToken: "token-initial",
		ExpiresAt:   timestamppb.New(time.Now().Add(30 * time.Second)),
	}
}

func rotatedAuthProvider(context.Context) (*auth.AuthSnapshot, error) {
	return &auth.AuthSnapshot{
		Token: &serverv1.GetTokenResponse{
			AccessToken: "token-rotated",
			ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
		},
		EnvironmentID: "env-rotated",
	}, nil
}

func TestHTTPClientSendsRotatedAuthSnapshot(t *testing.T) {
	capture := &authCaptureTransport{}
	client, err := NewClient(t.Context(), &ClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		QueryServer:                "https://engine.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		HTTPClient:                 &http.Client{Transport: capture},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	impl := client.(*clientImpl)
	var response map[string]any
	require.NoError(t, impl.sendRequest(t.Context(), &sendRequestParams{
		Method:   http.MethodGet,
		URL:      "https://api.chalk.ai/test",
		Response: &response,
	}))
	require.Equal(t, "Bearer token-rotated", capture.request.Header.Get("Authorization"))
	require.Equal(t, "env-rotated", capture.request.Header.Get("X-Chalk-Env-Id"))
}

func TestGRPCClientSendsRotatedAuthSnapshot(t *testing.T) {
	client, err := NewGRPCClient(t.Context(), &GRPCClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		QueryServer:                "https://engine.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	var captured http.Header
	next := func(_ context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		captured = req.Header().Clone()
		return nil, nil
	}
	req := connect.NewRequest(&serverv1.GetGraphRequest{})
	_, err = client.(*grpcClientImpl).engineInterceptor(next)(t.Context(), req)
	require.NoError(t, err)
	require.Equal(t, "Bearer token-rotated", captured.Get("Authorization"))
	require.Equal(t, "env-rotated", captured.Get("x-chalk-env-id"))
}

func TestVolumeClientSendsRotatedAuthSnapshot(t *testing.T) {
	client, err := NewVolumeClient(t.Context(), &VolumeClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	header := http.Header{}
	require.NoError(t, client.(*volumeClientImpl).addAuthHeaders(t.Context(), header))
	require.Equal(t, "Bearer token-rotated", header.Get("Authorization"))
	require.Equal(t, "env-rotated", header.Get("x-chalk-env-id"))
}
