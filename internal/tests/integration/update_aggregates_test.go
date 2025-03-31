package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

type MatAggTxn struct {
	Id     *string
	Amount *float64
	AggId  *int64
	Ts     *time.Time
}

type MatAggs struct {
	Id       *int64
	Name     *string
	TxnCount map[string]*float64 `windows:"3d,10d,30d"`
	TxnSum   map[string]*float64 `windows:"3d,10d,30d"`
}

var matAggFeatures struct {
	MatAggTxn *MatAggTxn
	MatAggs   *MatAggs
}

func init() {
	if err := chalk.InitFeatures(&matAggFeatures); err != nil {
		panic(err)
	}
}

func TestWindowedAggregates(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	now := time.Now().UTC()

	randomId1 := now.UnixMicro()
	randomId2 := randomId1 + 1

	txnIds := []string{"txn_1", "txn_2", "txn_3", "txn_4"}
	txnAmounts := []float64{100, 200, 300, 400}
	txnAggIds := []int64{randomId1, randomId1, randomId1, randomId2}
	txnTimes := []time.Time{
		now.Add(-1 * time.Minute),     // 1 minute ago for Alice
		now.Add(-7 * 24 * time.Hour),  // 7 days ago for Alice
		now.Add(-15 * 24 * time.Hour), // 15 days ago for Alice
		now.Add(-1 * 24 * time.Hour),  // 1 day ago for Bob
	}

	txnParams := chalk.UpdateAggregatesParams{Inputs: map[any]any{
		matAggFeatures.MatAggTxn.Id:     txnIds,
		matAggFeatures.MatAggTxn.Amount: txnAmounts,
		matAggFeatures.MatAggTxn.AggId:  txnAggIds,
		matAggFeatures.MatAggTxn.Ts:     txnTimes,
	}}
	_, err := grpcClient.UpdateAggregates(context.Background(), txnParams)
	assert.NoError(t, err)

	queryParams := chalk.OnlineQueryParams{}.WithInput(
		matAggFeatures.MatAggs.Id,
		[]int64{randomId1, randomId2},
	).WithOutputs(
		matAggFeatures.MatAggs.TxnCount["3d"],
		matAggFeatures.MatAggs.TxnCount["10d"],
		matAggFeatures.MatAggs.TxnCount["30d"],
		matAggFeatures.MatAggs.TxnSum["3d"],
		matAggFeatures.MatAggs.TxnSum["10d"],
		matAggFeatures.MatAggs.TxnSum["30d"],
	)

	bulkRes, err := grpcClient.OnlineQueryBulk(context.Background(), queryParams)
	assert.NoError(t, err)

	var aggResults []MatAggs
	assert.NoError(t, bulkRes.UnmarshalInto(&aggResults))
	assert.Equal(t, 2, len(aggResults))

	// Check Alice's counts (id=1)
	assert.Equal(t, float64(1), *aggResults[0].TxnCount["3d"])  // 1 txn in last 3 days
	assert.Equal(t, float64(2), *aggResults[0].TxnCount["10d"]) // 2 txns in last 10 days
	assert.Equal(t, float64(3), *aggResults[0].TxnCount["30d"]) // 3 txns in last 30 days

	// Check Alice's sums
	assert.Equal(t, float64(100), *aggResults[0].TxnSum["3d"])  // $100 in last 3 days
	assert.Equal(t, float64(300), *aggResults[0].TxnSum["10d"]) // $300 in last 10 days
	assert.Equal(t, float64(600), *aggResults[0].TxnSum["30d"]) // $600 in last 30 days

	// Check Bob's counts (id=2)
	assert.Equal(t, float64(1), *aggResults[1].TxnCount["3d"])  // 1 txn in last 3 days
	assert.Equal(t, float64(1), *aggResults[1].TxnCount["10d"]) // 1 txn in last 10 days
	assert.Equal(t, float64(1), *aggResults[1].TxnCount["30d"]) // 1 txn in last 30 days

	// Check Bob's sums
	assert.Equal(t, float64(400), *aggResults[1].TxnSum["3d"])  // $400 in last 3 days
	assert.Equal(t, float64(400), *aggResults[1].TxnSum["10d"]) // $400 in last 10 days
	assert.Equal(t, float64(400), *aggResults[1].TxnSum["30d"]) // $400 in last 30 days
}
