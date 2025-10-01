package graph

import (
	"fmt"
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
			WithPrimary("id", Int).
			With("val", Int.Expr(expr.Col("id").Add(expr.Int64(1)))),
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
			WithPrimary("id", Int).
			WithForeignKey("user_id", "User").
			With("amount", Float).
			With("at", Datetime),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int).
			With("name", String).
			With("transactions", DataFrame("Transaction")).
			With("total_spend", Windowed(Float, Days(30), Days(60), Days(90)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(Days(1)).
				WithExpr(expr.DataFrame("transactions").
					Filter(expr.Col("at").Gt(expr.ChalkWindow())).
					Filter(expr.Col("at").Lt(expr.ChalkNow())).
					Select(expr.Col("amount")).
					Agg("sum"),
				)),
	).ToGraph()

	assert.NoError(t, err)
	// 4 regular features + feature time + singleton relation
	assert.Equal(t, 6, len(graph.FeatureSets[0].Features))
	// 3 regular features + 4 windowed columns + feature time + singleton relation
	assert.Equal(t, 9, len(graph.FeatureSets[1].Features))
}

// getEvalRecordFeatureSet creates the EvalRecord FeatureSet for the given namespace
func getEvalRecordFeatureSet(accountNamespace string) FeatureSet {
	evalRecordFeatureSetName := fmt.Sprintf("eval_record_%s", accountNamespace)

	return FeatureSet{Name: evalRecordFeatureSetName, EtlOfflineToOnline: false}.
		WithPrimary("id", String).
		WithAll(Features{
			"case_status":                          String,
			"decision":                             String,
			"sub_status":                           String,
			"decision_and_sub_status_both_approve": Boolean,
			"eval_at":                              Datetime,
			"reviewer_updated_at":                  Datetime,
			"reason_codes":                         String,
			"is_approve":                           Boolean,
			"is_first_time_seen":                   Boolean,
			"is_first_transaction":                 Boolean,
			"data":                                 String,
		})
}

func TestWindowedGroupedCount(t *testing.T) {
	// Create the windowed aggregation expression
	expression := expr.DataFrame("events").
		Filter(expr.Col("ts").Gt(expr.ChalkWindow())).
		Filter(expr.Col("ts").Lt(expr.ChalkNow())).
		Agg("count")

	evalRecordFS := getEvalRecordFeatureSet("acctns")

	// Define feature sets using the graph builder API
	definitions := Definitions{}.
		WithFeatureSets(
			// SimpleEvent feature set
			FeatureSet{Name: "SimpleEvent"}.
				WithPrimary("id", Int).
				WithForeignKey("user_id", "User").
				With("amount", Float).
				With("ts", Datetime),

			// User feature set with windowed aggregation
			FeatureSet{Name: "User"}.
				WithPrimary("id", Int).
				WithAll(Features{
					"address_country":              String,
					"address_line_1":               String,
					"address_line_2":               String,
					"address_locality":             String,
					"address_major_admin_division": String,
					"address_minor_admin_division": String,
					"address_postal_code":          String,
					"address_type":                 String,
					"dob":                          String,
					"first_name":                   String,
					"last_name":                    String,
					"primary_email":                String,
					"primary_phone":                String,
					"events":                       DataFrame("SimpleEvent"),
					"count_events": Windowed(Int, Minutes(1), Minutes(5)).
						WithBucketDuration(Minutes(1)).
						WithExpr(expression),
				}),

			evalRecordFS,
		)

	// Convert to protobuf Graph
	_, err := definitions.ToGraph()
	assert.NoError(t, err)
}

func TestWindowedAllTime(t *testing.T) {
	definitions := Definitions{}.
		WithFeatureSets(
			// SimpleEvent feature set
			FeatureSet{Name: "other_account_id_with_same_stripe_fingerprint_found"}.
				WithPrimary("account_id", Int).
				With("other_account_id", String).
				With("ts", Datetime),

			// User feature set with windowed aggregation
			FeatureSet{Name: "other_account_id_with_same_stripe_fingerprint_found_count"}.
				WithPrimary("account_id", Int).
				With("other_accounts_df", DataFrame("other_account_id_with_same_stripe_fingerprint_found").WithMaxStaleness(Days(1))).
				With("other_account_id_count",
					Windowed(Int).
						WithExpr(expr.DataFrame("other_accounts_df").Select(expr.Col("other_account_id")).Agg("approx_count_distinct"))),
		)
	graph, err := definitions.ToGraph()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(graph.FeatureSets[0].Features))
	assert.Equal(t, 6, len(graph.FeatureSets[1].Features))
}
