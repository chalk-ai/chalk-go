package integration

import (
	"context"
	"fmt"
	"github.com/chalk-ai/chalk-go/v2"
	"github.com/chalk-ai/chalk-go/v2/internal/ptr"
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

	for _, optionFixture := range plannerOptionsFixtures {
		t.Run(fmt.Sprintf("plannerOptionValid=%v", optionFixture.isValid), func(t *testing.T) {
			_, err := restClient.OnlineQuery(
				context.Background(),
				chalk.OnlineQueryParams{PlannerOptions: optionFixture.plannerOptions}.
					WithInput("user.id", 1).
					WithOutputs("user.socure_score"),
				nil,
			)
			if optionFixture.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestOnlineQueryBulkPlannerOptions(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, optionFixture := range plannerOptionsFixtures {
			t.Run(fmt.Sprintf("grpc=%v, plannerOptionValid=%v", useGrpc, optionFixture.isValid), func(t *testing.T) {
				params := chalk.OnlineQueryParams{
					PlannerOptions: optionFixture.plannerOptions,
				}.
					WithInput("user.id", []int{1}).
					WithOutputs("user.socure_score")

				var err error
				if useGrpc {
					_, err = grpcClient.OnlineQueryBulk(context.Background(), params)
				} else {
					_, err = restClient.OnlineQueryBulk(context.Background(), params)
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

var timeouts = []struct {
	name       string
	timeout    time.Duration
	shouldFail bool
}{
	{name: "1 nanosecond", timeout: 1 * time.Nanosecond, shouldFail: true},
	{name: "5 seconds", timeout: 5 * time.Second},
	{name: "unspecified (zero value)", timeout: 0},
}

func TestTimeoutClientLevel(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				var err error
				if useGrpc {
					_, err = chalk.NewGRPCClient(context.Background(), &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
				} else {
					_, err = chalk.NewClient(context.Background(), &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
				}
				if timeoutFixture.shouldFail {
					assert.Error(t, err)
					return
				} else {
					assert.NoError(t, err)
				}

			})
		}
	}
}

func TestTimeoutClientOverrides(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*1)
				defer cancelFunc()
				var err error
				if useGrpc {
					_, err = chalk.NewGRPCClient(ctx, &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
				} else {
					_, err = chalk.NewClient(ctx, &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
				}
				// Since we've passed in a context with lenient timeout override,
				// all client instantiation should succeed.
				assert.NoError(t, err)

			})
		}
	}
}

func TestTimeoutRequestOverrides(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	lenientTimeout := time.Minute * 1
	params := chalk.OnlineQueryParams{}.
		WithInput("user.id", []int{1}).
		WithOutputs("user.socure_score")
	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				ctx, cancelFunc := context.WithTimeout(context.Background(), lenientTimeout)
				defer cancelFunc()
				if useGrpc {
					timeoutClient, err := chalk.NewGRPCClient(ctx, &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
					assert.NoError(t, err)

					// lenient override
					requestCtx, requestCancelFunc := context.WithTimeout(ctx, lenientTimeout)
					defer requestCancelFunc()
					res, err := timeoutClient.OnlineQueryBulk(requestCtx, params)
					assert.NoError(t, err)
					assert.Equal(t, 0, len(res.RawResponse.GetErrors()))

					// no override
					_, err = timeoutClient.OnlineQueryBulk(context.Background(), params)
					if timeoutFixture.shouldFail {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				} else {
					timeoutClient, err := chalk.NewClient(ctx, &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
					assert.NoError(t, err)

					// lenient override
					requestCtx, requestCancelFunc := context.WithTimeout(ctx, lenientTimeout)
					defer requestCancelFunc()
					_, err = timeoutClient.OnlineQueryBulk(requestCtx, params)
					assert.NoError(t, err)

					// no override
					_, err = timeoutClient.OnlineQueryBulk(context.Background(), params)
					if timeoutFixture.shouldFail {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				}
			})
		}
	}
}
