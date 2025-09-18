package integration

import (
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestErringScalar tests requests with an erring scalar feature as the sole output
func TestErringScalar(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{true, false} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			if useGrpc {
				grpcClient := newGRPCClient(t)
				bulkParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Crashing.Id, []int{1}).
					WithOutputs(testFeatures.Crashing.Name)

				resp, err := grpcClient.OnlineQueryBulk(t.Context(), bulkParams)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "3.14")

				row, err := resp.GetRow(0)
				assert.NoError(t, err)

				crashingFeature, err := row.GetFeature(testFeatures.Crashing.Name)
				assert.NoError(t, err)
				assert.Nil(t, crashingFeature.Value)
			} else {
				restClient := newRestClient(t)
				singularParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Crashing.Id, 1).
					WithOutputs(testFeatures.Crashing.Name)
				_, err := restClient.OnlineQuery(t.Context(), singularParams, nil)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "3.14")
			}
		})
	}
}

// TestErringHasMany tests requests with an erring has-many feature as the sole output
func TestErringHasMany(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{true, false} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			if useGrpc {
				grpcClient := newGRPCClient(t)
				bulkParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.CrashingHasManyRoot.Id, []string{"id_a"}).
					WithOutputs(testFeatures.CrashingHasManyRoot.CrashingHasMany)
				resp, err := grpcClient.OnlineQueryBulk(t.Context(), bulkParams)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "42")

				row, err := resp.GetRow(0)
				assert.NoError(t, err)

				hmFeat, err := row.GetFeature(testFeatures.CrashingHasManyRoot.CrashingHasMany)
				assert.NoError(t, err)
				assert.NotNil(t, hmFeat)
				castVal, ok := hmFeat.Value.([]any)
				assert.True(t, ok)
				assert.Equal(t, 0, len(castVal))
			} else {
				restClient := newRestClient(t)
				singularParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.CrashingHasManyRoot.Id, "id_a").
					WithOutputs(testFeatures.CrashingHasManyRoot.CrashingHasMany)
				_, err := restClient.OnlineQuery(t.Context(), singularParams, nil)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "42")
			}
		})
	}
}
