package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

// TestOnlineQueryGrpcIncludeMeta mainly tests that the response
// includes the correct metadata when requested.
func TestOnlineQueryGrpcIncludeMeta(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	pkey := "chalk-go-include-meta-test"
	randomNumber := rand.Float64()
	_, err := restClient.UploadFeatures(context.Background(), chalk.UploadFeaturesParams{
		Inputs: map[any]any{
			testFeatures.Cached.Id:                   []string{pkey},
			testFeatures.Cached.RandomUploadedNumber: []float64{randomNumber},
		},
	})
	assert.NoError(t, err)
	req := chalk.OnlineQueryParams{IncludeMeta: true}.
		WithInput(testFeatures.Cached.Id, []string{pkey}).
		WithOutputs(testFeatures.Cached.Id, testFeatures.Cached.RandomUploadedNumber)
	cachedRes, err := grpcClient.OnlineQueryBulk(context.Background(), req)
	assert.NoError(t, err)

	cachedRow, err := cachedRes.GetRow(0)
	assert.NoError(t, err)

	actualNumber, err := cachedRow.GetFeature(testFeatures.Cached.RandomUploadedNumber)
	assert.Nil(t, err)
	assert.NotNil(t, actualNumber)
	assert.NotNil(t, actualNumber.Meta)
	assert.Equal(t, randomNumber, actualNumber.Value)
	assert.Equal(t, "online_store", actualNumber.Meta.SourceType)

	resolverRes, err := grpcClient.OnlineQueryBulk(context.Background(), chalk.OnlineQueryParams{IncludeMeta: true}.
		WithInput(testFeatures.AllTypes.Id, []int64{1}).
		WithOutputs(testFeatures.AllTypes.StrFeat),
	)

	resolvedRow, err := resolverRes.GetRow(0)
	resolvedFeat, err := resolvedRow.GetFeature(testFeatures.AllTypes.StrFeat)
	assert.NoError(t, err)
	assert.Equal(t, "registry.all_feature_types.get_all_types", resolvedFeat.Meta.ResolverFqn)
}

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
