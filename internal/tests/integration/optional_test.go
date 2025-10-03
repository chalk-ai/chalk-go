package integration

import (
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestQueryOptionalFeatures(t *testing.T) {
	t.Parallel()
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient(t.Context())
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	idWithNone := 1
	res := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Optionals.Id, idWithNone).
		WithOutputs(testFeatures.Optionals.Name)
	resWithNone := optionals{}
	_, err = client.OnlineQuery(t.Context(), res, &resWithNone)
	assert.NoError(t, err)
	assert.Nil(t, resWithNone.Name)

	idNotNone := 0
	res = chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Optionals.Id, idNotNone).
		WithOutputs(testFeatures.Optionals.Name)
	resNotNone := optionals{}
	_, err = client.OnlineQuery(t.Context(), res, &resNotNone)
	assert.NoError(t, err)
	assert.NotNil(t, resNotNone.Name)
	assert.Equal(t, "name_0", *resNotNone.Name)
}

// TestBulkQueryOptionalFeatures tests bulk queries with optional features
// parameterized between gRPC and REST clients
func TestBulkQueryOptionalFeatures(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	// Test IDs where one returns nil and one returns a value
	ids := []int64{0, 1}

	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()

			// Create params for bulk query
			params := chalk.OnlineQueryParams{}.
				WithInput(testFeatures.Optionals.Id, ids).
				WithOutputs(testFeatures.Optionals.Name)

			var results []optionals

			if useGrpc {
				res, err := grpcClient.OnlineQueryBulk(t.Context(), params)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&results))
			} else {
				res, err := restClient.OnlineQueryBulk(t.Context(), params)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&results))
			}

			assert.Equal(t, len(ids), len(results))
			assert.NotNil(t, results[0].Name)
			assert.Equal(t, "name_0", *results[0].Name)
			assert.Nil(t, results[1].Name)
		})
	}
}
