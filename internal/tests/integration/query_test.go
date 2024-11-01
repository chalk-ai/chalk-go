package integration

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getParams() chalk.OnlineQueryParamsComplete {
	return chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(
			testFeatures.User.Id,
			testFeatures.User.Gender,
			testFeatures.User.Today,
			testFeatures.User.NiceNewFeature,
			testFeatures.User.SocureScore,
			testFeatures.User.FavoriteNumbers,
			testFeatures.User.FavoriteColors,
			testFeatures.User.FranchiseSet,
		)
}

func testUserValues(t *testing.T, testUser *user) {
	t.Helper()
	assert.Equal(t, *testUser.Id, int64(1))
	assert.Equal(t, *testUser.Gender, "f")
	assert.NotNil(t, testUser.Today)
	assert.Equal(t, *testUser.NiceNewFeature, int64(9))
	assert.Equal(t, *testUser.SocureScore, 123.0)
	assert.Equal(t, *testUser.FavoriteNumbers, []int64{1, 2, 3})
	assert.Equal(t, *testUser.FavoriteColors, []string{"red", "green", "blue"})
	assert.NotNil(t, testUser.FranchiseSet)
}

// TestOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			res, queryErr := client.OnlineQuery(getParams(), &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}

			var explicitUser user
			err = res.UnmarshalInto(&explicitUser)
			// TODO: We need to fix this nil check
			//       to just be `!= nil`
			if err != (*chalk.ClientError)(nil) {
				t.Fatal("Failed unmarshaling result", err)
			}

			testUserValues(t, &implicitUser)
			testUserValues(t, &explicitUser)
		})
	}
}

// TestNamedQueriesE2E tests that querying with a query name works.
func TestNamedQueriesE2E(t *testing.T) {
	t.Skip("CHA-5086")
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			params := chalk.OnlineQueryParams{}.
				WithInput("user.id", 1).
				WithQueryName("user_socure_score")

			_, queryErr := client.OnlineQuery(params, &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}
			assert.Equal(t, 123.0, *implicitUser.SocureScore)
		})
	}
}

// TestGRPCOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
//
// This test is also notably different from the E2E test
// where a gRPC client is also tested but is built on top
// of the existing REST `Client` interface.
func TestGRPCOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewGRPCClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}

	res, queryErr := client.OnlineQuery(context.Background(), getParams())
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var testUser user
	assert.NoError(t, chalk.UnmarshalOnlineQueryResponse(res, &testUser))
	testUserValues(t, &testUser)
}

// TestOnlineQueryBulkParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInFeatherHeader. Correctness
// of the results is *not* tested here.
func TestOnlineQueryBulkParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			err := chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			userIds := []int{1, 2}

			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now(), time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.User.Id, userIds).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			_, err = client.OnlineQueryBulk(req)
			assert.NoError(t, err)
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

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			if fixture.useGrpc {
				t.Skip("CHA-4780")
			}
			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.User.Id, 1).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			_, err = client.OnlineQuery(req, nil)
			assert.NoError(t, err)
		})
	}
}

// Test that we can execute an OnlineQuery
// with has-manys as both inputs and outputs.
// Correctness of unmarshalling all data types
// within a has-many feature is tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryHasManyInputsAndOutputs(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	assert.NoError(t, chalk.InitFeatures(&testFeatures))
	client, err := chalk.NewClient()
	assert.NoError(t, err)
	investorsInput := []newGradAngelInvestor{
		{Id: ptr.Ptr("amylase"), SeriesId: ptr.Ptr("seed"), HowBroke: ptr.Ptr(int64(1))},
		{Id: ptr.Ptr("lipase"), SeriesId: ptr.Ptr("seed"), HowBroke: ptr.Ptr(int64(2))},
	}
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Id, "seed").
		WithInput(testFeatures.Series.Investors, investorsInput).
		WithOutputs(testFeatures.Series.Name, testFeatures.Series.Investors)

	var resultSeries series
	res, err := client.OnlineQuery(params, &resultSeries)
	assert.NoError(t, err)
	assert.Equal(t, len(investorsInput), len(*resultSeries.Investors))
	assert.Equal(t, "amylase", *(*resultSeries.Investors)[0].Id)
	assert.Equal(t, "lipase", *(*resultSeries.Investors)[1].Id)
	assert.Equal(t, int64(1), *(*resultSeries.Investors)[0].HowBroke)
	assert.Equal(t, int64(2), *(*resultSeries.Investors)[1].HowBroke)
	assert.Equal(t, "seed", *(*resultSeries.Investors)[0].SeriesId)
	assert.Equal(t, "seed", *(*resultSeries.Investors)[1].SeriesId)

	investorsFeature, err := chalk.UnwrapFeature(testFeatures.Series.Investors)
	assert.NoError(t, err)

	// has many result should be a map that has
	// "columns" and "values" as keys. Correctness
	// of this GetFeatureValue result is guaranteed
	// if the result of UnmarshalInto is correct,
	// and that is being checked above.
	resultInvestors, err := res.GetFeatureValue(investorsFeature)
	assert.NoError(t, err)
	assert.NotNil(t, resultInvestors)
}
