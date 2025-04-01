package integration

import (
	"context"
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

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
			t.Parallel()
			_, err := restClient.OnlineQuery(
				context.Background(),
				chalk.OnlineQueryParams{PlannerOptions: optionFixture.plannerOptions}.
					WithInput(testFeatures.AllTypes.Id, 1).
					WithOutputs(testFeatures.AllTypes.StrFeat),
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
				t.Parallel()
				params := chalk.OnlineQueryParams{
					PlannerOptions: optionFixture.plannerOptions,
				}.
					WithInput(testFeatures.AllTypes.Id, []int{1}).
					WithOutputs(testFeatures.AllTypes.StrFeat)

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
