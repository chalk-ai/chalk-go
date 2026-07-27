package chalk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/auth"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDeleteFeaturesRequestWireFormat(t *testing.T) {
	t.Parallel()
	// Must match chalkpy's FeatureObservationDeletionRequest exactly.
	body, err := json.Marshal(deleteFeaturesRequest(DeleteFeaturesParams{
		Namespace:     "user",
		Features:      []string{"email", "name"},
		Tags:          nil,
		PrimaryKeys:   []string{"1", "2"},
		RetainOffline: true,
		RetainOnline:  false,
	}))
	require.NoError(t, err)

	var got map[string]any
	require.NoError(t, json.Unmarshal(body, &got))
	assert.Equal(t, "user", got["namespace"])
	assert.Equal(t, []any{"email", "name"}, got["features"])
	assert.Nil(t, got["tags"])
	assert.Equal(t, []any{"1", "2"}, got["primary_keys"])
	assert.Equal(t, true, got["retain_offline"])
	assert.Equal(t, false, got["retain_online"])

	// The booleans must be present even when false -- omitempty would drop
	// retain_online and change the server-side default.
	_, hasRetainOnline := got["retain_online"]
	assert.True(t, hasRetainOnline, "retain_online must be sent explicitly, not omitted")
}

func TestDeleteFeaturesRequestAndResponse(t *testing.T) {
	t.Parallel()

	var gotMethod string
	var gotPath string
	var gotEnvironment string
	var gotBody deleteFeaturesRequest
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotEnvironment = r.Header.Get("X-Chalk-Env-Id")
		require.NoError(t, json.NewDecoder(r.Body).Decode(&gotBody))

		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
			"errors": []map[string]any{{
				"code":     "PARSE_FAILED",
				"category": "FIELD",
				"message":  "could not delete one feature",
			}},
		}))
	}))
	t.Cleanup(apiServer.Close)

	ctx := context.Background()
	cfg, err := config.NewManager(ctx, &config.ManagerInputs{
		APIServer:     apiServer.URL,
		ClientId:      "client-id",
		ClientSecret:  "client-secret",
		EnvironmentId: "test-env",
	})
	require.NoError(t, err)
	tokenManager, err := auth.NewManager(ctx, &auth.Inputs{
		Token: &serverv1.GetTokenResponse{
			AccessToken:         "token",
			ExpiresAt:           timestamppb.New(time.Now().Add(time.Hour)),
			PrimaryEnvironment:  ptrTo("test-env"),
			EnvironmentIdToName: map[string]string{"test-env": "Test"},
			Engines:             map[string]string{},
			GrpcEngines:         map[string]string{},
		},
		HttpClient:                 http.DefaultClient,
		Config:                     cfg,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	client := &clientImpl{
		config:       cfg,
		httpClient:   http.DefaultClient,
		logger:       DefaultLeveledLogger,
		tokenManager: tokenManager,
	}
	params := DeleteFeaturesParams{
		Namespace:     "user",
		Tags:          []string{"pii"},
		PrimaryKeys:   []string{"1", "2"},
		RetainOffline: true,
	}
	result, err := client.DeleteFeatures(ctx, params)
	require.NoError(t, err)

	assert.Equal(t, http.MethodDelete, gotMethod)
	assert.Equal(t, "/v1/features/rows", gotPath)
	assert.Equal(t, "test-env", gotEnvironment)
	assert.Equal(t, deleteFeaturesRequest(params), gotBody)
	require.Len(t, result.Errors, 1)
	assert.Equal(t, "could not delete one feature", result.Errors[0].Message)
}

func TestDeleteFeaturesRejectsBranchDeployment(t *testing.T) {
	t.Parallel()

	client := &clientImpl{Branch: "my-branch"}
	_, err := client.DeleteFeatures(context.Background(), DeleteFeaturesParams{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not supported for branch deployments")
}
