package graph

import (
	"fmt"
	"testing"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
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
	assertValid(t, FeatureSet{Name: "Example"}.
		WithPrimary("id", Int).
		With("val", Int.Expr(expr.Col("id").Add(expr.Int64(1)))))
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
	checkCorrectNumFeatures(t, graph, map[string]int{
		// 4 regular features + feature time + singleton relation
		"transaction": 6,
		// 3 regular features + 4 windowed columns + feature time + singleton relation
		"user": 9,
	})
	assert.NoError(t, CheckHasFeature(graph, "transaction", "id"))
	assert.NoError(t, CheckHasFeature(graph, "user", "id"))
}

// getEvalRecordFeatureSet creates the EvalRecord FeatureSet for the given namespace
func getEvalRecordFeatureSet(accountNamespace string) FeatureSet {
	evalRecordFeatureSetName := fmt.Sprintf("eval_record_%s", accountNamespace)

	return FeatureSet{Name: evalRecordFeatureSetName}.
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

func TestWindowedCount(t *testing.T) {
	// Create the windowed aggregation expression
	expression := expr.DataFrame("events").
		Filter(expr.Col("ts").Gt(expr.ChalkWindow())).
		Filter(expr.Col("ts").Lt(expr.ChalkNow())).
		Agg("count")

	evalRecordFS := getEvalRecordFeatureSet("acctxns")

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
	graph, err := definitions.ToGraph()
	assert.NoError(t, err)
	assert.NoError(t, CheckHasFeature(graph, "simple_event", "id"))
	assert.NoError(t, CheckHasFeature(graph, "user", "id"))
	assert.NoError(t, CheckHasFeature(graph, "user", "count_events__60__"))
	assert.NoError(t, CheckHasFeature(graph, "user", "count_events__300__"))
	checkCorrectNumFeatures(t, graph, map[string]int{
		"simple_events":        5,
		"user":                 20,
		"eval_record_accttxns": 24,
	})
}

func checkCorrectNumFeatures(t *testing.T, graph *graphv1.Graph, nameToNumFeatures map[string]int) {
	for _, fs := range graph.FeatureSets {
		num, ok := nameToNumFeatures[fs.Name]
		if ok {
			println(fs.Name)
			for _, f := range fs.Features {
				println(FeatureName(f))
			}
			println()
			assert.Equal(t, num, len(fs.Features), "incorrect number of features in %s", fs.Name)
		}
	}
}

func TestWindowedAllTime(t *testing.T) {
	definitions := Definitions{}.
		WithFeatureSets(
			// SimpleEvent feature set
			FeatureSet{Name: "other_account_id_with_same_stripe_fingerprint_found"}.
				WithPrimary("account_id", Int).
				With("ts", Datetime).
				WithForeignKey("other_account_id", "other_account_id_with_same_stripe_fingerprint_found_count"),

			// User feature set with windowed aggregation
			FeatureSet{Name: "other_account_id_with_same_stripe_fingerprint_found_count"}.
				WithPrimary("account_id", Int).
				With("other_accounts_df", DataFrame("other_account_id_with_same_stripe_fingerprint_found").WithMaxStaleness(Days(1))).
				With("other_account_id_count",
					Windowed(Int, All).
						WithExpr(expr.DataFrame("other_accounts_df").Select(expr.Col("other_account_id")).Agg("approx_count_distinct"))),
		)
	graph, err := definitions.ToGraph()
	assert.NoError(t, err)
	checkCorrectNumFeatures(t, graph, map[string]int{
		"other_account_id_with_same_stripe_fingerprint_found":       4,
		"other_account_id_with_same_stripe_fingerprint_found_count": 6,
	})

	assert.NoError(t, CheckHasFeature(
		graph,
		"other_account_id_with_same_stripe_fingerprint_found_count",
		"other_account_id_count__all__",
	))
}

func assertValid(t *testing.T, fs FeatureSet) {
	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.NoError(t, err)
}

func assertInvalid(t *testing.T, fs FeatureSet) {
	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestDuplicateFeatureTimeNotAllowed(t *testing.T) {
	assertInvalid(t, FeatureSet{Name: "event"}.
		WithPrimary("id", Int).
		With("ts", Datetime).
		With("at", FeatureTime))
}

func TestWithForeignKey(t *testing.T) {
	assertInvalid(t, FeatureSet{Name: "event"}.
		WithPrimary("id", Int).
		WithForeignKey("other_id", "DoesNotExist"))
}

func TestMaxByN(t *testing.T) {
	_, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Transaction"}.
			WithPrimary("id", Int).
			WithForeignKey("user_id", "User").
			With("amount", Float).
			With("at", Datetime),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int).
			With("transactions", DataFrame("Transaction")).
			With("top_5_txns", Windowed(List(Int), Days(30), Days(60), Days(90)).
				WithDefault(expr.Int(0)). //
				WithBucketDuration(Days(1)).
				WithExpr(expr.DataFrame("transactions").
					Filter(expr.Col("at").Gt(expr.ChalkWindow())).
					Filter(expr.Col("at").Lt(expr.ChalkNow())).
					Select(expr.Col("id")).
					Agg("max_by_n", expr.Col("amount"), expr.Int64(5)),
				)),
	).ToGraph()
	assert.NoError(t, err)
}

func TestForeignKeyWithExpr(t *testing.T) {
	graph, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Transaction"}.
			WithPrimary("id", Int).
			WithForeignKey("user_id", "User").
			With("user_id", Int.Expr(expr.Int64(69))),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int),
	).ToGraph()
	assert.NoError(t, err)
	checkCorrectNumFeatures(t, graph, map[string]int{
		"user":        3,
		"transaction": 4,
	})
}

func TestHasOneRelationship(t *testing.T) {
	_, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Transaction"}.
			WithPrimary("id", Int).
			With("user_id", Int).
			With("user", HasOne("User", expr.ColIn("User", "id").Eq(expr.ColIn("Transaction", "user_id")))),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int),
	).ToGraph()
	assert.NoError(t, err)
}

func TestDuplicatePrimaryKeyError(t *testing.T) {
	fs := FeatureSet{Name: "User"}.
		WithPrimary("id", Int).
		WithPrimary("other_id", Int)

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestDuplicateFeatureNameError(t *testing.T) {
	fs := FeatureSet{Name: "User"}.
		WithPrimary("id", Int).
		With("name", String).
		With("name", String)

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestTsFieldMustBeDatetimeError(t *testing.T) {
	fs := FeatureSet{Name: "Event"}.
		WithPrimary("id", Int).
		With("ts", Int)

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestPrimaryKeyMustBeScalarError(t *testing.T) {
	fs := FeatureSet{Name: "User"}.
		WithPrimary("id", DataFrame("Other"))

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestWindowedMissingDataFrameError(t *testing.T) {
	fs := FeatureSet{Name: "User"}.
		WithPrimary("id", Int).
		With("event_count", Windowed(Int, Minutes(5)).
			WithExpr(expr.DataFrame("events").Agg("count")))

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestWindowedMustBeScalarError(t *testing.T) {
	fs := FeatureSet{Name: "User"}.
		WithPrimary("id", Int).
		With("events", DataFrame("Event")).
		With("bad", Windowed(DataFrame("Event"), Minutes(5)).
			WithExpr(expr.DataFrame("events").Agg("count")))

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestForeignKeyNonExistentError(t *testing.T) {
	fs := FeatureSet{Name: "Event"}.
		WithPrimary("id", Int).
		WithForeignKey("other_id", "DoesNotExist")

	_, err := Definitions{}.WithFeatureSets(fs).ToGraph()
	assert.Error(t, err)
}

func TestStreamResolverInvalidSourceTypeError(t *testing.T) {
	evalRecordFS := FeatureSet{Name: "eval_record"}.
		WithPrimary("id", String).
		With("decision", String)

	defs := Definitions{}.
		WithFeatureSets(evalRecordFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "eval_record_stream",
				StreamSourceName: "invalid_source",
				StreamSourceType: "invalid",
				OutputFeatureSet: "eval_record",
				OutputFeatures: map[string]expr.Expr{
					"decision": expr.Col("decision"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"decision": "str",
				}},
			},
		)

	_, err := defs.ToGraph()
	assert.Error(t, err)
}

func TestStreamResolverMissingFeatureSetError(t *testing.T) {
	defs := Definitions{}.
		WithStreamResolvers(
			StreamResolver{
				Name:             "eval_record_stream",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "nonexistent",
				OutputFeatures: map[string]expr.Expr{
					"decision": expr.Col("decision"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"decision": "str",
				}},
			},
		)

	_, err := defs.ToGraph()
	assert.Error(t, err)
}

func TestStreamResolverMissingOutputFeatureError(t *testing.T) {
	evalRecordFS := FeatureSet{Name: "eval_record"}.
		WithPrimary("id", String).
		With("decision", String)

	defs := Definitions{}.
		WithFeatureSets(evalRecordFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "eval_record_stream",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "eval_record",
				OutputFeatures: map[string]expr.Expr{
					"decision":    expr.Col("decision"),
					"nonexistent": expr.Col("nonexistent"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"decision":    "str",
					"nonexistent": "str",
				}},
			},
		)

	_, err := defs.ToGraph()
	assert.Error(t, err)
}

func TestStreamResolverEmptyNameError(t *testing.T) {
	evalRecordFS := FeatureSet{Name: "eval_record"}.
		WithPrimary("id", String).
		With("decision", String)

	defs := Definitions{}.
		WithFeatureSets(evalRecordFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "eval_record",
				OutputFeatures: map[string]expr.Expr{
					"decision": expr.Col("decision"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"decision": "str",
				}},
			},
		)

	_, err := defs.ToGraph()
	assert.Error(t, err)
}

func TestStreamResolverDotsInNameError(t *testing.T) {
	evalRecordFS := FeatureSet{Name: "eval_record"}.
		WithPrimary("id", String).
		With("decision", String)

	defs := Definitions{}.
		WithFeatureSets(evalRecordFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "eval.record.stream",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "eval_record",
				OutputFeatures: map[string]expr.Expr{
					"decision": expr.Col("decision"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"decision": "str",
				}},
			},
		)

	_, err := defs.ToGraph()
	assert.Error(t, err)
}

func TestSingletonFeatureSet(t *testing.T) {
	graph, err := Definitions{}.WithFeatureSets(
		FeatureSet{
			Name:        "Config",
			IsSingleton: true,
		}.With("api_key", String),
	).ToGraph()
	assert.NoError(t, err)

	configFS := getFeatureSetNamed(graph, "config")
	assert.NotNil(t, configFS)
	assert.True(t, configFS.IsSingleton)

	err = checkHasFeatures(configFS, "__chalk_fk_singleton__")
	assert.Error(t, err, "singleton should not have __chalk_fk_singleton__")

	assert.NoError(t, checkHasFeatures(configFS, "api_key"))
}

func TestWindowedNonCountWithoutSelectionError(t *testing.T) {
	_, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Event"}.
			WithPrimary("id", Int).
			WithForeignKey("user_id", "User").
			With("amount", Float),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int).
			With("events", DataFrame("Event")).
			With("total", Windowed(Float, Days(1)).
				WithExpr(expr.DataFrame("events").Agg("sum"))),
	).ToGraph()
	assert.Error(t, err)
}

func TestApproxTopKMustBeListError(t *testing.T) {
	_, err := Definitions{}.WithFeatureSets(
		FeatureSet{Name: "Event"}.
			WithPrimary("id", Int).
			WithForeignKey("user_id", "User").
			With("item", String),

		FeatureSet{Name: "User"}.
			WithPrimary("id", Int).
			With("events", DataFrame("Event")).
			With("top_items", Windowed(String, Days(1)).
				WithExpr(expr.DataFrame("events").
					Select(expr.Col("item")).
					Agg("approx_top_k", expr.Int64(5)))),
	).ToGraph()
	assert.Error(t, err)
}

/*
from pydantic import BaseModel

from chalk import functions as F
from chalk import features
from chalk.features import _
from chalk.features.resolver import make_stream_resolver
from chalk.streams import KafkaSource

@features
class Transaction:

	id: int
	user_id: str
	amount: float

class TxnMessage(BaseModel):

	id: int
	user_id: str
	amount: float

message_str = F.bytes_to_string(_, encoding="utf-8")

parse_expression = F.if_then_else(

	F.cast(F.json_value(message_str, "$.amount"), pa.float64) > 0,
	F.json_value(message_str, "$"),
	None,

)

kafka_source = KafkaSource(...)

streaming_resolver = make_stream_resolver(

	name="get_txn_stream",
	source=kafka_source,
	message_type=TxnMessage,
	parse=parse_expression,
	output_features={
	    Transaction.id: _.id,
	    Transaction.user_id: _.user_id,
	    Transaction.amount: _.amount,
	},

)
*/
func TestStreamResolverValid(t *testing.T) {
	transactionFS := FeatureSet{Name: "Transaction"}.
		WithPrimary("id", Int).
		WithAll(Features{
			"user_id": String,
			"amount":  Float,
		})

	messageStr := expr.BytesToUtf8(expr.Col("_"))

	defs := Definitions{}.
		WithFeatureSets(transactionFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "get_txn_stream",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "transaction",
				OutputFeatures: map[string]expr.Expr{
					"id":      expr.Col("id"),
					"user_id": expr.Col("user_id"),
					"amount":  expr.Col("amount"),
				},
				MessageType: StreamStructMessage{Fields: map[string]string{
					"id":      "int",
					"user_id": "str",
					"amount":  "float",
				}},
				Parse: expr.IfElse(
					expr.Cast(
						expr.GetJsonValue(messageStr, "$.amount"),
						&arrowv1.ArrowType{ArrowTypeEnum: &arrowv1.ArrowType_Float64{}},
					).Gt(expr.Float(0)),
					messageStr,
					expr.Null(),
				),
			},
		)

	graph, err := defs.ToGraph()
	assert.NoError(t, err)
	assert.Len(t, graph.StreamResolvers, 1)
}

//# def _build_parse_fn():
//#     message_str = F.bytes_to_string(_, encoding="utf-8")
//#     primary_email_field = F.cast(F.json_value(message_str, "$.primary_email"), pa.large_utf8())
//#     secondary_email_field = F.cast(F.json_value(message_str, "$.secondary_email"), pa.large_utf8())
//#     primary_email_struct = F.struct_pack(
//#         {
//#             "id": F.cast(F.rand(), pa.large_utf8()),
//#             "email": primary_email_field,
//#         }
//#     )
//#     secondary_email_struct = F.struct_pack(
//#         {
//#             "id": F.cast(F.rand(), pa.large_utf8()),
//#             "email": secondary_email_field,
//#         }
//#     )
//#     return F.array(primary_email_struct, secondary_email_struct)
//#
//#
//# socure_pii_resolver = make_stream_resolver(
//#     source=source,
//#     name="socure_pii_resolver",
//#     message_type=list[EvalRecordMessage],
//#     parse=_build_parse_fn(),
//#     output_features={EvalRecord.id: _.id, EvalRecord.email: _.email},
//# )

func TestStreamResolverParseFnReturnsList(t *testing.T) {
	evalRecordFS := FeatureSet{Name: "EvalRecord"}.
		WithPrimary("id", Int).
		WithAll(Features{
			"email": String,
		})

	messageStr := expr.BytesToUtf8(expr.Col("_"))
	emailPrimaryField := expr.GetJsonValue(messageStr, "$.primary_email")
	emailSecondaryField := expr.GetJsonValue(messageStr, "$.secondary_email")
	emailPrimaryRecord := expr.StructPack([]expr.StructField{
		{"id", expr.Cast(expr.Rand(), &arrowv1.ArrowType{ArrowTypeEnum: &arrowv1.ArrowType_LargeUtf8{}})},
		{"email", emailPrimaryField},
	})
	emailSecondaryRecord := expr.StructPack([]expr.StructField{
		{"id", expr.Cast(expr.Rand(), &arrowv1.ArrowType{ArrowTypeEnum: &arrowv1.ArrowType_LargeUtf8{}})},
		{"email", emailSecondaryField},
	})
	parseExpr := expr.ArrayConstructor(emailPrimaryRecord, emailSecondaryRecord)

	defs := Definitions{}.
		WithFeatureSets(evalRecordFS).
		WithStreamResolvers(
			StreamResolver{
				Name:             "parse_email_records",
				StreamSourceName: "kafka_source",
				StreamSourceType: "kafka",
				OutputFeatureSet: "eval_record",
				OutputFeatures: map[string]expr.Expr{
					"id":    expr.Col("id"),
					"email": expr.Col("email"),
				},
				MessageType: StreamListMessage{ValueType: StreamStructMessage{Fields: map[string]string{
					"id":    "int",
					"email": "str",
				}}},
				Parse: parseExpr,
			},
		)

	graph, err := defs.ToGraph()
	assert.NoError(t, err)
	assert.Len(t, graph.StreamResolvers, 1)
}
