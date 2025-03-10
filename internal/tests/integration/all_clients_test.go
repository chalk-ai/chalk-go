package integration

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	assert "github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var restClient chalk.Client
var grpcClient chalk.GRPCClient

func init() {
	if err := chalk.InitFeatures(&testFeatures); err != nil {
		panic(err)
	}

	if os.Getenv("INTEGRATION_TESTER") == "" {
		return
	}

	ctx := context.Background()

	client, err := chalk.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	restClient = client

	clientGrpc, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	grpcClient = clientGrpc

}

// Test that we can execute an OnlineQuery
// with has-manys as both inputs and outputs.
// Correctness of unmarshalling all data types
// within a has-many feature is tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestHasManyInputsAndOutputs(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	investorsInput := []newGradAngelInvestor{
		{Id: ptr.Ptr("amylase"), SeriesId: ptr.Ptr("seed"), HowBroke: ptr.Ptr(int64(1))},
		{Id: ptr.Ptr("lipase"), SeriesId: ptr.Ptr("seed"), HowBroke: ptr.Ptr(int64(2))},
	}

	for _, useGrpc := range []bool{false, true} {
		t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
			var resultSeries series
			if useGrpc {
				var resultSeriesBulk []series
				params := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Series.Id, []string{"seed"}).
					WithInput(testFeatures.Series.Investors, [][]newGradAngelInvestor{investorsInput}).
					WithOutputs(testFeatures.Series.Name, testFeatures.Series.Investors)
				res, err := grpcClient.OnlineQueryBulk(context.Background(), params)
				assert.NoError(t, err)
				assert.NoError(t, res.UnmarshalInto(&resultSeriesBulk))
				assert.Equal(t, 1, len(resultSeriesBulk))
				resultSeries = resultSeriesBulk[0]

				row, err := res.GetRow(0)
				assert.NoError(t, err)
				resultInvestors, err := row.GetFeatureValue(testFeatures.Series.Investors)
				assert.NoError(t, err)
				assert.NotNil(t, resultInvestors)
			} else {
				params := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.Series.Id, "seed").
					WithInput(testFeatures.Series.Investors, investorsInput).
					WithOutputs(testFeatures.Series.Name, testFeatures.Series.Investors)
				res, err := restClient.OnlineQuery(context.Background(), params, &resultSeries)
				assert.NoError(t, err)

				resultInvestors, err := res.GetFeatureValue(testFeatures.Series.Investors)
				assert.NoError(t, err)
				assert.NotNil(t, resultInvestors)
			}

			assert.Equal(t, len(investorsInput), len(*resultSeries.Investors))
			assert.Equal(t, "amylase", *(*resultSeries.Investors)[0].Id)
			assert.Equal(t, "lipase", *(*resultSeries.Investors)[1].Id)
			assert.Equal(t, int64(1), *(*resultSeries.Investors)[0].HowBroke)
			assert.Equal(t, int64(2), *(*resultSeries.Investors)[1].HowBroke)
			assert.Equal(t, "seed", *(*resultSeries.Investors)[0].SeriesId)
			assert.Equal(t, "seed", *(*resultSeries.Investors)[1].SeriesId)
		})
	}
}

type plannerOptionsFixture struct {
	isValid        bool
	plannerOptions map[string]any
}

var plannerOptionsFixtures = []plannerOptionsFixture{
	{isValid: true, plannerOptions: map[string]any{"planner_version": "2"}},
	{isValid: false, plannerOptions: map[string]any{"planner_version": "abcdefg"}},
}

func TestOnlineQueryPlannerOptions(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, optionFixture := range plannerOptionsFixtures {
			t.Run(fmt.Sprintf("grpc=%v, plannerOptionValid=%v", useGrpc, optionFixture.isValid), func(t *testing.T) {
				var err error
				baseParams := chalk.OnlineQueryParams{
					PlannerOptions: optionFixture.plannerOptions,
				}.
					WithOutputs("user.socure_score")
				userId := 1
				if useGrpc {
					_, err = grpcClient.OnlineQueryBulk(
						context.Background(),
						baseParams.WithInput("user.id", []int{userId}),
					)
				} else {
					_, err = restClient.OnlineQuery(
						context.Background(),
						baseParams.WithInput("user.id", userId),
						nil,
					)
				}
				if optionFixture.isValid {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
			})
		}
	}
}

//func TestOnlineQueryBulkPlannerOptions(t *testing.T) {
//	t.Parallel()
//	SkipIfNotIntegrationTester(t)
//
//	for _, clientFixture := range clients {
//		for _, optionFixture := range plannerOptionsFixtures {
//			t.Run(fmt.Sprintf("grpc=%v, plannerOptionValid=%v", clientFixture.name, optionFixture.isValid), func(t *testing.T) {
//				client := clientFixture.client
//				params := chalk.OnlineQueryParams{
//					PlannerOptions: optionFixture.plannerOptions,
//				}.
//					WithInput("user.id", []int{1}).
//					WithOutputs("user.socure_score")
//				_, err := client.OnlineQueryBulk(context.Background(), params)
//				if optionFixture.isValid {
//					assert.NoError(t, err)
//				} else {
//					assert.Error(t, err)
//				}
//			})
//		}
//	}
//}

func TestTimeout(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	timeouts := []struct {
		name       string
		timeout    time.Duration
		shouldFail bool
	}{
		{name: "1 nanosecond", timeout: 1 * time.Nanosecond, shouldFail: true},
		{name: "5 seconds", timeout: 5 * time.Second},
		{name: "unspecified (zero value)", timeout: 0},
	}

	// TODO: Reintroduce GRPC client to this test
	for _, timeoutFixture := range timeouts {
		client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
		t.Run(fmt.Sprintf("timeoutFixture=%v", timeoutFixture.name), func(t *testing.T) {
			t.Parallel()
			if timeoutFixture.shouldFail {
				assert.Error(t, err)
				return
			} else {
				params := chalk.OnlineQueryParams{}.
					WithInput("user.id", 1).
					WithOutputs("user.socure_score")
				assert.NoError(t, err)
				myUser := user{}
				_, err := client.OnlineQuery(context.Background(), params, &myUser)
				assert.NoError(t, err)
				assert.NotNil(t, myUser.SocureScore)
				assert.Equal(t, 123.0, *myUser.SocureScore)
			}
		})
	}
}
