package graph

import (
	"testing"

	"github.com/chalk-ai/chalk-go/expr"
	"github.com/stretchr/testify/assert"
)

/*
@features
class Example:

	id: int
	val: _.id + 1
*/
func TestSimpleAdd(t *testing.T) {
	_, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Example"}.
			WithPrimary("id", Int()).
			With("val", Int().Expr(__("id").Add(expr.Int64(1)))),
	).ToGraph()

	assert.NoError(t, err)
}

/*
@features
class Transaction:

	id: int
	user_id: "User.id"
	amount: float
	at: datetime
	name: str

@features
class User:

	id: int
	transactions: DataFrame[Transaction]
	total_spend: Windowed[float] = windowed(
	    "30d", "60d", "90d",
	    default=0,
	    expression=_.transactions[
	        _.amount,
	        _.at > _.chalk_window,
	        _.at < _.chalk_now,
	    ].sum(),
	    materialization={"bucket_duration": "1d"},
	)
*/
func TestWindowed(t *testing.T) {
	graph, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Transaction"}.
			WithPrimary("id", Int()).
			WithForeignKey("user_id", "User").
			With("amount", Float()).
			With("at", Datetime()),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int()).
			With("name", String()).
			With("transactions", DataFrame("Transaction")).
			With("total_spend", Windowed(Float(), Days(30), Days(60), Days(90)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(Days(1)).
				WithExpr(__("transactions").Get(
					__("amount"),
					__("at").Gt(__("chalk_window")),
					__("at").Lt(__("chalk_now")),
				).Attr("sum").Apply())),
	).ToGraph()

	assert.NoError(t, err)
	println(graph)
	// 4 regular features + feature time
	assert.Equal(t, 5, len(graph.FeatureSets[0].Features))
	// 3 regular features + 4 windowed columns + feature time
	assert.Equal(t, 8, len(graph.FeatureSets[1].Features))
}
