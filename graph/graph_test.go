package graph

import (
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
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
func TestWindowedSum(t *testing.T) {
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
				WithExpr(expr.DataFrame("transactions").
					Filter(__("at").Gt(__("chalk_window"))).
					Filter(__("at").Lt(__("chalk_now"))).
					Select(__("amount")).
					Agg("sum"),
				)),
	).ToGraph()

	assert.NoError(t, err)
	println(graph)
	// 4 regular features + feature time
	assert.Equal(t, 5, len(graph.FeatureSets[0].Features))
	// 3 regular features + 4 windowed columns + feature time
	assert.Equal(t, 8, len(graph.FeatureSets[1].Features))
}

func TestWindowedGroupedCount(t *testing.T) {
	// Create the windowed aggregation expression
	expression := expr.DataFrame("events").
		Filter(expr.Col("ts").Gt(expr.Col("chalk_window"))).
		Filter(expr.Col("ts").Lt(expr.Col("chalk_now"))).
		Agg("count")

	// Define feature sets using the graph builder API
	definitions := Definitions{}.
		WithFeatureSets(
			// SimpleEvent feature set
			FeatureSet{Name: "SimpleEvent"}.
				WithPrimary("id", Int()).
				WithForeignKey("user_id", "User").
				With("amount", Float()).
				With("ts", Datetime()),

			// User feature set with windowed aggregation
			FeatureSet{Name: "User"}.
				WithPrimary("id", Int()).
				With("address_country", String()).
				With("address_line_1", String()).
				With("address_line_2", String()).
				With("address_locality", String()).
				With("address_major_admin_division", String()).
				With("address_minor_admin_division", String()).
				With("address_postal_code", String()).
				With("address_type", String()).
				With("dob", String()).
				With("first_name", String()).
				With("last_name", String()).
				With("primary_email", String()).
				With("primary_phone", String()).
				With("events", DataFrame("SimpleEvent")).
				With("count_events",
					Windowed(Int(), Minutes(1), Minutes(5)).WithMaterialization(
						MaterializationOptions{
							DefaultBucketDuration: time.Minute * 1,
							GroupBy:               []*graphv1.FeatureReference{{Name: "id", Namespace: "user"}},
						},
					).WithExpr(expression)),
		)

	// Convert to protobuf Graph
	_, err := definitions.ToGraph()
	assert.NoError(t, err)
}
