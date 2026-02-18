package integration

import (
	"fmt"
	"testing"
	"time"

	chalk "github.com/chalk-ai/chalk-go"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// TestOnlineQueryE2E tests the singular OnlineQuery method
// which is only available on the REST client.
// The bulk query functionality is tested in TestOnlineQueryBulk for both clients.
func TestOnlineQuery(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	params := chalk.OnlineQueryParams{}.
		WithOutputs(
			testFeatures.AllTypes.Id,
			testFeatures.AllTypes.StrFeat,
			testFeatures.AllTypes.IntFeat,
		)

	// Test the REST client's singular OnlineQuery method
	var implicitResult allTypes
	res, err := restClient.OnlineQuery(
		t.Context(),
		params.WithInput(testFeatures.AllTypes.Id, 1),
		&implicitResult,
	)
	assert.NoError(t, err)

	var explicitResult allTypes
	assert.NoError(t, res.UnmarshalInto(&explicitResult))

	// Verify results from both implicit and explicit unmarshalling
	for _, result := range []allTypes{implicitResult, explicitResult} {
		assert.NotNil(t, result.Id)
		assert.Equal(t, int64(1), lo.FromPtr(result.Id))
		assert.NotNil(t, result.StrFeat)
		assert.Equal(t, "1", lo.FromPtr(result.StrFeat))
		assert.NotNil(t, result.IntFeat)
		assert.Equal(t, int64(1), lo.FromPtr(result.IntFeat))
	}
}

// TestOnlineQueryBulk mainly tests that a
// real query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulk(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			ids := []int64{1, 2}
			var results []allTypes
			req := chalk.OnlineQueryParams{}.
				WithInput(testFeatures.AllTypes.Id, ids).
				WithOutputs(
					testFeatures.AllTypes.Id,
					testFeatures.AllTypes.StrFeat,
					testFeatures.AllTypes.IntFeat,
				)
			if useGrpc {
				res, err := grpcClient.OnlineQueryBulk(t.Context(), req)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&results))
			} else {
				res, err := restClient.OnlineQueryBulk(t.Context(), req)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&results))
			}
			assert.Equal(t, 2, len(results))
			assert.Equal(t, ids[0], lo.FromPtr(results[0].Id))
			assert.Equal(t, "1", lo.FromPtr(results[0].StrFeat))
			assert.Equal(t, int64(1), lo.FromPtr(results[0].IntFeat))
			assert.Equal(t, ids[1], lo.FromPtr(results[1].Id))
			assert.Equal(t, "2", lo.FromPtr(results[1].StrFeat))
			assert.Equal(t, int64(2), lo.FromPtr(results[1].IntFeat))
		})
	}
}

// Test that we can execute an OnlineQuery
// with has-manys as both inputs and outputs.
// Correctness of unmarshalling all data types
// within a has-many feature is tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestHasManyInputsAndOutputs(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	hmInput := []hasManyFeature{
		{Id: new("id_a"), Name: new("name_a"), AllTypesId: new(int64(1))},
		{Id: new("id_b"), Name: new("name_b"), AllTypesId: new(int64(1))},
	}

	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			var row allTypes

			if useGrpc {
				var allResults []allTypes
				bulkParams := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.AllTypes.Id, []int64{1}).
					WithInput(testFeatures.AllTypes.HasMany, [][]hasManyFeature{hmInput}).
					WithOutputs(testFeatures.AllTypes.StrFeat, testFeatures.AllTypes.HasMany)
				res, err := grpcClient.OnlineQueryBulk(t.Context(), bulkParams)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&allResults))
				assert.Equal(t, 1, len(allResults))
				row = allResults[0]

				row, err := res.GetRow(0)
				assert.NoError(t, err)
				hmOutput, err := row.GetFeatureValue(testFeatures.AllTypes.HasMany)
				assert.NoError(t, err)
				assert.NotNil(t, hmOutput)
			} else {
				params := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.AllTypes.Id, 1).
					WithInput(testFeatures.AllTypes.HasMany, hmInput).
					WithOutputs(testFeatures.AllTypes.StrFeat, testFeatures.AllTypes.HasMany)
				res, err := restClient.OnlineQuery(t.Context(), params, &row)
				assert.NoError(t, err)

				resultInvestors, err := res.GetFeatureValue(testFeatures.AllTypes.HasMany)
				assert.NoError(t, err)
				assert.NotNil(t, resultInvestors)
			}

			if row.HasMany != nil {
				assert.Equal(t, len(hmInput), len(*row.HasMany))
				assert.Equal(t, "id_a", *(*row.HasMany)[0].Id)
				assert.Equal(t, "id_b", *(*row.HasMany)[1].Id)
				assert.Equal(t, "name_a", *(*row.HasMany)[0].Name)
				assert.Equal(t, "name_b", *(*row.HasMany)[1].Name)
				assert.Equal(t, int64(1), *(*row.HasMany)[0].AllTypesId)
				assert.Equal(t, int64(1), *(*row.HasMany)[1].AllTypesId)
			} else {
				t.Fatal("expected has-many output to be non-nil")
				return
			}
		})
	}
}

// TestOnlineQueryBulkParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInFeatherHeader. Correctness
// of the results is *not* tested here.
func TestOnlineQueryBulkParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()

			// Test IDs for bulk query
			ids := []int64{1, 2}

			req := chalk.OnlineQueryParams{
				// TODO: Add separate test for tags
				//Tags:                 []string{"named-integration"},
				//RequiredResolverTags: []string{"named-integration"},
				Now:              []time.Time{time.Now(), time.Now()},
				StorePlanStages:  true,
				CorrelationId:    "chalk-go-int-test-correlation-id",
				QueryName:        "chalk-go-int-test-query",
				QueryNameVersion: "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.AllTypes.Id, ids).
				WithOutputs(testFeatures.AllTypes.StrFeat).
				WithStaleness(testFeatures.AllTypes.IntFeat, time.Minute*10)

			if useGrpc {
				_, err := grpcClient.OnlineQueryBulk(t.Context(), req)
				assert.NoError(t, err)
			} else {
				_, err := restClient.OnlineQueryBulk(t.Context(), req)
				assert.NoError(t, err)
			}
		})
	}
}

// TestOnlineQueryParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInOnlineQuery. Correctness
// of the results is *not* tested here.
func TestOnlineQueryParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			t.Parallel()
			req := chalk.OnlineQueryParams{
				// TODO: Add separate tag tests
				//Tags:                 []string{"named-integration"},
				//RequiredResolverTags: []string{"named-integration"},
				Now:              []time.Time{time.Now()},
				StorePlanStages:  true,
				CorrelationId:    "chalk-go-int-test-correlation-id",
				QueryName:        "chalk-go-int-test-query",
				QueryNameVersion: "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithOutputs(testFeatures.AllTypes.StrFeat).
				WithStaleness(testFeatures.AllTypes.IntFeat, time.Minute*10)

			if useGrpc {
				_, err := grpcClient.OnlineQueryBulk(
					t.Context(),
					req.WithInput(testFeatures.AllTypes.Id, []int64{1}),
				)
				assert.NoError(t, err)
			} else {
				_, err := restClient.OnlineQuery(
					t.Context(),
					req.WithInput(testFeatures.AllTypes.Id, int64(1)),
					nil,
				)
				assert.NoError(t, err)
			}
		})
	}
}
