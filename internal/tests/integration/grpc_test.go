package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestOnlineQueryGrpcErringScalar tests requests with an erring scalar feature as the sole output
func TestOnlineQueryGrpcErringScalar(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Crashing.Id, []int{1}).
		WithOutputs(testFeatures.Crashing.Name)
	resp, err := grpcClient.OnlineQueryBulk(context.Background(), params)
	assert.Error(t, err)

	row, err := resp.GetRow(0)
	assert.NoError(t, err)

	crashingFeature, err := row.GetFeature(testFeatures.Crashing.Name)
	assert.NoError(t, err)

	assert.NotNil(t, crashingFeature)
	assert.Nil(t, crashingFeature.Value)
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
