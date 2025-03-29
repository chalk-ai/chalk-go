package integration

import (
	"context"
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
				bulkParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Crashing.Id, []int{1}).
					WithOutputs(testFeatures.Crashing.Name)

				resp, err := grpcClient.OnlineQueryBulk(context.Background(), bulkParams)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "3.14")

				row, err := resp.GetRow(0)
				assert.NoError(t, err)

				crashingFeature, err := row.GetFeature(testFeatures.Crashing.Name)
				assert.NoError(t, err)
				assert.Nil(t, crashingFeature.Value)
			} else {
				singularParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Crashing.Id, 1).
					WithOutputs(testFeatures.Crashing.Name)
				_, err := restClient.OnlineQuery(context.Background(), singularParams, nil)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "3.14")
			}
		})
	}
}

// TestOnlineQueryGrpcErringHasMany tests requests with an erring has-many feature as the sole output
func TestOnlineQueryGrpcErringHasMany(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.CrashingHasManyRoot.Id, []string{"id_a"}).
		WithOutputs(testFeatures.CrashingHasManyRoot.CrashingHasMany)
	resp, err := grpcClient.OnlineQueryBulk(context.Background(), params)
	assert.Error(t, err)

	row, err := resp.GetRow(0)
	assert.NoError(t, err)

	hmFeat, err := row.GetFeature(testFeatures.CrashingHasManyRoot.CrashingHasMany)
	assert.NoError(t, err)
	assert.NotNil(t, hmFeat)
	castVal, ok := hmFeat.Value.([]any)
	assert.True(t, ok)
	assert.Equal(t, 0, len(castVal))
}
