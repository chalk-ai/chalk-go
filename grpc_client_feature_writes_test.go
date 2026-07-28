package chalk

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow/go/v16/arrow/memory"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeFeatureWriteClient struct {
	uploadParams UploadFeaturesParams
	deleteParams DeleteFeaturesParams
	uploadResult UploadFeaturesResult
	deleteResult DeleteFeaturesResult
	uploadErr    error
	deleteErr    error
}

type recordingConnectHTTPClient struct {
	requests []*http.Request
}

func (c *recordingConnectHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.requests = append(c.requests, req)
	body := `{"errors":[]}`
	if req.URL.Path == "/v1/upload_features/multi" {
		body = `{"operation_id":"operation-id"}`
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
	}, nil
}

func (c *fakeFeatureWriteClient) UploadFeatures(
	_ context.Context,
	params UploadFeaturesParams,
) (UploadFeaturesResult, error) {
	c.uploadParams = params
	return c.uploadResult, c.uploadErr
}

func (c *fakeFeatureWriteClient) DeleteFeatures(
	_ context.Context,
	params DeleteFeaturesParams,
) (DeleteFeaturesResult, error) {
	c.deleteParams = params
	return c.deleteResult, c.deleteErr
}

func TestGRPCClientUploadFeaturesDelegatesToHTTPClient(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("upload failed")
	delegate := &fakeFeatureWriteClient{
		uploadResult: UploadFeaturesResult{OperationId: "operation-id"},
		uploadErr:    expectedErr,
	}
	client := &grpcClientImpl{featureWriteClient: delegate}
	params := UploadFeaturesParams{
		Inputs: map[any]any{"user.email": []string{"user@example.com"}},
	}

	result, err := client.UploadFeatures(context.Background(), params)

	assert.Equal(t, params, delegate.uploadParams)
	assert.Equal(t, delegate.uploadResult, result)
	require.ErrorIs(t, err, expectedErr)
}

func TestGRPCClientDeleteFeaturesDelegatesToHTTPClient(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("delete failed")
	delegate := &fakeFeatureWriteClient{
		deleteResult: DeleteFeaturesResult{
			Errors: ServerErrors{{Message: "partial deletion"}},
		},
		deleteErr: expectedErr,
	}
	client := &grpcClientImpl{featureWriteClient: delegate}
	params := DeleteFeaturesParams{
		Namespace:   "user",
		Features:    []string{"email"},
		PrimaryKeys: []string{"user-1"},
	}

	result, err := client.DeleteFeatures(context.Background(), params)

	assert.Equal(t, params, delegate.deleteParams)
	assert.Equal(t, delegate.deleteResult, result)
	require.ErrorIs(t, err, expectedErr)
}

func TestNewGRPCClientSharesStateWithHTTPFeatureWriteClient(t *testing.T) {
	t.Parallel()

	allocator := memory.NewCheckedAllocator(memory.DefaultAllocator)
	client, err := newGrpcClient(context.Background(), &GRPCClientConfig{
		ClientId:                   "client-id",
		ClientSecret:               "client-secret",
		ApiServer:                  "https://api.example.com",
		QueryServer:                "https://query.example.com",
		EnvironmentId:              "environment-id",
		Branch:                     "branch-id",
		DeploymentTag:              "deployment-tag",
		ResourceGroup:              "resource-group",
		HTTPClient:                 http.DefaultClient,
		Allocator:                  allocator,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken: "token",
			ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
		},
	})
	require.NoError(t, err)

	delegate, ok := client.featureWriteClient.(*clientImpl)
	require.True(t, ok)
	assert.Same(t, client.config, delegate.config)
	assert.Same(t, client.tokenManager, delegate.tokenManager)
	assert.Same(t, allocator, delegate.allocator)
	assert.Same(t, http.DefaultClient, delegate.httpClient)
	assert.Equal(t, client.branch, delegate.Branch)
	assert.Equal(t, client.deploymentTag, delegate.DeploymentTag)
	assert.Equal(t, client.resourceGroup, delegate.resourceGroup)
}

func TestGRPCClientFeatureWritesUseHTTPAPI(t *testing.T) {
	httpClient := &recordingConnectHTTPClient{}
	client, err := NewGRPCClient(context.Background(), &GRPCClientConfig{
		ClientId:                   "client-id",
		ClientSecret:               "client-secret",
		ApiServer:                  "https://api.example.com",
		QueryServer:                "https://grpc-query.example.com",
		EnvironmentId:              "environment-id",
		HTTPClient:                 httpClient,
		SkipEnvironmentNameMapping: true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken: "token",
			ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
			Engines: map[string]string{
				"environment-id": "https://http-query.example.com",
			},
		},
	})
	require.NoError(t, err)

	uploadResult, err := client.UploadFeatures(context.Background(), UploadFeaturesParams{
		Inputs: map[any]any{"user.email": []string{"user@example.com"}},
	})
	require.NoError(t, err)
	assert.Equal(t, "operation-id", uploadResult.OperationId)

	deleteResult, err := client.DeleteFeatures(context.Background(), DeleteFeaturesParams{
		Namespace:   "user",
		Features:    []string{"email"},
		PrimaryKeys: []string{"user-1"},
	})
	require.NoError(t, err)
	assert.Empty(t, deleteResult.Errors)

	require.Len(t, httpClient.requests, 2)
	assert.Equal(t, http.MethodPost, httpClient.requests[0].Method)
	assert.Equal(t, "https://http-query.example.com/v1/upload_features/multi", httpClient.requests[0].URL.String())
	assert.Equal(t, http.MethodDelete, httpClient.requests[1].Method)
	assert.Equal(t, "https://api.example.com/v1/features/rows", httpClient.requests[1].URL.String())
	for _, req := range httpClient.requests {
		assert.Equal(t, "Bearer token", req.Header.Get("Authorization"))
		assert.Equal(t, "environment-id", req.Header.Get("X-Chalk-Env-Id"))
	}
}
