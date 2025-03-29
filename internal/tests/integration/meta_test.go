package integration

import (
	"context"
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

// TestCacheHitMeta mainly tests that the response
// includes the correct cache hit metadata.
func TestCacheHitMeta(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	pkey := "chalk-go-cache-hit-meta-test"
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

// TestResolverFqnMeta mainly tests that the response
// includes the correct resolver metadata.
func TestResolverMeta(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	resolverFqn := "registry.all_feature_types.get_all_types"
	for _, useGrpc := range []bool{true, false} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			if useGrpc {
				resolverRes, err := grpcClient.OnlineQueryBulk(context.Background(), chalk.OnlineQueryParams{IncludeMeta: true}.
					WithInput(testFeatures.AllTypes.Id, []int64{1}).
					WithOutputs(testFeatures.AllTypes.StrFeat),
				)

				resolvedRow, err := resolverRes.GetRow(0)
				assert.NoError(t, err)
				resolvedFeat, err := resolvedRow.GetFeature(testFeatures.AllTypes.StrFeat)
				assert.NoError(t, err)
				assert.Equal(t, resolverFqn, resolvedFeat.Meta.ResolverFqn)
			} else {
				res, err := restClient.OnlineQuery(context.Background(), chalk.OnlineQueryParams{IncludeMeta: true}.
					WithInput(testFeatures.AllTypes.Id, 1).
					WithOutputs(testFeatures.AllTypes.StrFeat),
					nil,
				)
				assert.NoError(t, err)
				resolvedFeat, err := res.GetFeature(testFeatures.AllTypes.StrFeat)
				assert.NoError(t, err)
				assert.Equal(t, resolverFqn, resolvedFeat.Meta.ChosenResolverFqn)
			}
		})
	}

}
