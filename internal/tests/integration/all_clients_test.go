package integration

import (
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	assert "github.com/stretchr/testify/require"
	"os"
	"testing"
)

var clients []ClientFixture

type ClientFixture struct {
	name   string
	client chalk.Client
}

func init() {
	if err := chalk.InitFeatures(&testFeatures); err != nil {
		panic(err)
	}

	if os.Getenv("INTEGRATION_TESTER") == "" {
		return
	}

	restClient, err := chalk.NewClient()
	if err != nil {
		panic(err)
	}
	clients = append(clients, ClientFixture{name: "rest", client: restClient})

	grpcClient, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		panic(err)
	}
	clients = append(clients, ClientFixture{name: "grpc", client: grpcClient})

}

// Test that we can execute an OnlineQuery
// with has-manys as both inputs and outputs.
// Correctness of unmarshalling all data types
// within a has-many feature is tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestHasManyInputsAndOutputs(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, clientFixture := range clients {
		t.Run(clientFixture.name, func(t *testing.T) {
			client := clientFixture.client
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
		})
	}
}
