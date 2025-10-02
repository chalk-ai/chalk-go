package integration

import (
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

// TestOnlineQueryGrpcIncludeMeta mainly tests that the response
// includes the correct metadata when requested.
func TestCachedFeatures(t *testing.T) {
	t.Parallel()
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	SkipIfNotIntegrationTester(t)
	for _, useGrpc := range []bool{true, false} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			if useGrpc {
				pkey := "chalk-go-upload-features-test-gRPC"
				expectedNum := rand.Float64()
				_, err := restClient.UploadFeatures(t.Context(), chalk.UploadFeaturesParams{
					Inputs: map[any]any{
						testFeatures.Cached.Id:                   []string{pkey},
						testFeatures.Cached.RandomUploadedNumber: []float64{expectedNum},
					},
				})
				assert.NoError(t, err)
				bulkReq := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Cached.Id, []string{pkey}).
					WithOutputs(testFeatures.Cached.Id, testFeatures.Cached.RandomUploadedNumber)
				cachedRes, err := grpcClient.OnlineQueryBulk(t.Context(), bulkReq)
				assert.NoError(t, err)

				cachedRow, err := cachedRes.GetRow(0)
				assert.NoError(t, err)

				actualNum, err := cachedRow.GetFeature(testFeatures.Cached.RandomUploadedNumber)
				assert.Nil(t, err)
				assert.NotNil(t, actualNum)
				assert.Equal(t, expectedNum, actualNum.Value)
			} else {
				pkey := "chalk-go-upload-features-test-REST"
				expectedNum := rand.Float64()
				_, err := restClient.UploadFeatures(t.Context(), chalk.UploadFeaturesParams{
					Inputs: map[any]any{
						testFeatures.Cached.Id:                   []string{pkey},
						testFeatures.Cached.RandomUploadedNumber: []float64{expectedNum},
					},
				})
				assert.NoError(t, err)
				singularReq := chalk.OnlineQueryParams{IncludeMeta: true}.
					WithInput(testFeatures.Cached.Id, pkey).
					WithOutputs(testFeatures.Cached.Id, testFeatures.Cached.RandomUploadedNumber)
				cachedRes, err := restClient.OnlineQuery(t.Context(), singularReq, nil)
				assert.NoError(t, err)
				actualNum, err := cachedRes.GetFeatureValue(testFeatures.Cached.RandomUploadedNumber)
				assert.NoError(t, err)
				assert.Equal(t, expectedNum, actualNum)
			}
		})
	}
}
