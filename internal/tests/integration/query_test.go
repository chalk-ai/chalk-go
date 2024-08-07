package integration

import (
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

// TestFeatherHeaderParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInFeatherHeader. Correctness
// of the results is *not* tested here.
func TestFeatherHeaderParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	err := chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}

	expectedTags := []string{"named-integration"}
	requiredResolverTags := []string{"named-integration"}
	now := []time.Time{time.Now(), time.Now()}
	staleness := map[any]time.Duration{
		testFeatures.User.SocureScore: time.Minute * 10,
	}
	storePlanStages := true
	correlationId := "chalk-go-int-test-correlation-id"
	queryName := "chalk-go-int-test-query"
	queryNameVersion := "1"
	meta := map[string]string{
		"test_meta_1": "test_meta_value_1",
		"test_meta_2": "test_meta_value_2",
	}

	client, err := chalk.NewClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1, 2}

	req := chalk.OnlineQueryParams{
		Tags:                 expectedTags,
		RequiredResolverTags: requiredResolverTags,
		Now:                  now,
		StorePlanStages:      storePlanStages,
		CorrelationId:        correlationId,
		QueryName:            queryName,
		QueryNameVersion:     queryNameVersion,
		Meta:                 meta,
	}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.FullName)
	for k, v := range staleness {
		req = req.WithStaleness(k, v)
	}

	_, err = client.OnlineQueryBulk(req)
	assert.NoError(t, err)
}
