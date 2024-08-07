package integration

import (
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestOnlineQueryAllTypesUnmarshalling mainly tests
// unmarshalling real data from the staging server
// does not crash. Correctness is partially tested here,
// but is mainly tested in TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryAllTypesUnmarshalling(t *testing.T) {
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
			req := chalk.OnlineQueryParams{}.
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

			var implicitUser user
			res, queryErr := client.OnlineQuery(req, &implicitUser)
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

			testUserValues := func(testUser user) {
				assert.Equal(t, *testUser.Id, int64(1))
				assert.Equal(t, *testUser.Gender, "f")
				assert.NotNil(t, testUser.Today)
				assert.Equal(t, *testUser.NiceNewFeature, int64(9))
				assert.Equal(t, *testUser.SocureScore, 123.0)
				assert.Equal(t, *testUser.FavoriteNumbers, []int64{1, 2, 3})
				assert.Equal(t, *testUser.FavoriteColors, []string{"red", "green", "blue"})
				assert.NotNil(t, testUser.FranchiseSet)
			}
			testUserValues(implicitUser)
			testUserValues(explicitUser)
		})
	}
}

// Test that we can execute an OnlineQuery
// with has-manys as both inputs and outputs.
func TestOnlineQueryHasManyInputsAndOutputs(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	assert.NoError(t, chalk.InitFeatures(&testFeatures))
	client, err := chalk.NewClient()
	assert.NoError(t, err)
	investorsInput := []newGradAngelInvestor{
		{Id: lo.ToPtr("amylase"), SeriesId: lo.ToPtr("seed"), HowBroke: lo.ToPtr(int64(1))},
		{Id: lo.ToPtr("lipase"), SeriesId: lo.ToPtr("seed"), HowBroke: lo.ToPtr(int64(2))},
	}
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Investors, investorsInput).
		WithOutputs(testFeatures.Series.Name, testFeatures.Series.Investors)

	var resultSeries series
	res, err := client.OnlineQuery(params, &resultSeries)
	assert.NoError(t, err)
	assert.Equal(t, len(investorsInput), len(*resultSeries.Investors))

	investorsFeature, err := chalk.UnwrapFeature(testFeatures.Series.Investors)
	assert.NoError(t, err)

	// has many result should be a map that has "columns" and "values" as keys
	resultInvestors, err := res.GetFeatureValue(investorsFeature)
	assert.NotNil(t, resultInvestors)
	// TODO: Check resultInvestors["columns"][0] length
}
