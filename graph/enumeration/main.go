package main

import (
	"log"
	"os"

	"github.com/chalk-ai/chalk-go/expr"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	"github.com/chalk-ai/chalk-go/graph"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Define a standalone feature set with an int column and expressions using every operator
	standaloneFS := graph.FeatureSet{Name: "StandaloneFeatures"}.
		WithPrimary("id", graph.Int).
		WithAll(graph.Features{
			"value":     graph.Int,
			"amount":    graph.Float,
			"name":      graph.String,
			"is_active": graph.Boolean,
			"timestamp": graph.Datetime,

			// Arithmetic operators
			"add_result": graph.Int.Expr(expr.Col("value").Add(expr.Int64(10))),
			"sub_result": graph.Int.Expr(expr.Col("value").Sub(expr.Int64(5))),
			"mul_result": graph.Int.Expr(expr.Col("value").Mul(expr.Int64(2))),
			"div_result": graph.Float.Expr(expr.Col("amount").Div(expr.Float(2.0))),

			// Comparison operators
			"eq_check": graph.Boolean.Expr(expr.Col("value").Eq(expr.Int64(100))),
			"ne_check": graph.Boolean.Expr(expr.Col("value").Ne(expr.Int64(0))),
			"lt_check": graph.Boolean.Expr(expr.Col("value").Lt(expr.Int64(1000))),
			"le_check": graph.Boolean.Expr(expr.Col("value").Le(expr.Int64(1000))),
			"gt_check": graph.Boolean.Expr(expr.Col("value").Gt(expr.Int64(0))),
			"ge_check": graph.Boolean.Expr(expr.Col("value").Ge(expr.Int64(0))),

			// Logical operators
			"and_check": graph.Boolean.Expr(expr.Col("value").Gt(expr.Int64(0)).And(expr.Col("value").Lt(expr.Int64(100)))),
			"or_check":  graph.Boolean.Expr(expr.Col("value").Eq(expr.Int64(0)).Or(expr.Col("value").Eq(expr.Int64(100)))),
			"not_check": graph.Boolean.Expr(expr.Col("is_active").Not()),

			// Null checking operators
			"is_null_check":     graph.Boolean.Expr(expr.Col("name").IsNull()),
			"is_not_null_check": graph.Boolean.Expr(expr.Col("name").IsNotNull()),
		})

	// Define an Event feature set that will be used for DataFrame aggregations
	eventFS := graph.FeatureSet{Name: "Event"}.
		WithPrimary("id", graph.Int).
		WithForeignKey("user_id", "User").
		WithAll(graph.Features{
			"amount":    graph.Float,
			"item_name": graph.String,
			"timestamp": graph.Datetime,
			"category":  graph.String,
			"quantity":  graph.Int,
		})

	// Define User feature set with DataFrame and all possible aggregations
	countExpr := expr.DataFrame("events").
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	sumExpr := expr.DataFrame("events").
		Select(expr.Col("amount")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("sum")

	meanExpr := expr.DataFrame("events").
		Select(expr.Col("amount")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("mean")

	minExpr := expr.DataFrame("events").
		Select(expr.Col("amount")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("min")

	maxExpr := expr.DataFrame("events").
		Select(expr.Col("amount")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("max")

	stddevExpr := expr.DataFrame("events").
		Select(expr.Col("amount")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("stddev_sample")

	approxCountDistinctExpr := expr.DataFrame("events").
		Select(expr.Col("item_name")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("approx_count_distinct")

	maxByNExpr := expr.DataFrame("events").
		Select(expr.Col("id")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("max_by_n", expr.Col("timestamp"), expr.Int64(5))

	minByNExpr := expr.DataFrame("events").
		Select(expr.Col("id")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("min_by_n", expr.Col("timestamp"), expr.Int64(5))

	approxTopKExpr := expr.DataFrame("events").
		Select(expr.Col("item_name")).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("approx_top_k", expr.Int64(10))

	// Test different filter operations (using all comparison operators in filters)
	countWithEqFilter := expr.DataFrame("events").
		Filter(expr.Col("category").Eq(expr.String("electronics"))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithNeFilter := expr.DataFrame("events").
		Filter(expr.Col("category").Ne(expr.String("food"))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithLtFilter := expr.DataFrame("events").
		Filter(expr.Col("amount").Lt(expr.Float(100.0))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithLeFilter := expr.DataFrame("events").
		Filter(expr.Col("amount").Le(expr.Float(100.0))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithGtFilter := expr.DataFrame("events").
		Filter(expr.Col("amount").Gt(expr.Float(50.0))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithGeFilter := expr.DataFrame("events").
		Filter(expr.Col("amount").Ge(expr.Float(50.0))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithAndFilter := expr.DataFrame("events").
		Filter(expr.Col("amount").Gt(expr.Float(10.0)).And(expr.Col("amount").Lt(expr.Float(100.0)))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithOrFilter := expr.DataFrame("events").
		Filter(expr.Col("category").Eq(expr.String("electronics")).Or(expr.Col("category").Eq(expr.String("books")))).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	countWithNotFilter := expr.DataFrame("events").
		Filter(expr.Col("category").Eq(expr.String("food")).Not()).
		Filter(expr.Col("timestamp").Gt(expr.ChalkWindow())).
		Filter(expr.Col("timestamp").Lt(expr.ChalkNow())).
		Agg("count")

	userFS := graph.FeatureSet{Name: "User"}.
		WithPrimary("id", graph.Int).
		WithAll(graph.Features{
			"events": graph.DataFrame("Event"),

			// Test all aggregation functions
			"count_events": graph.Windowed(graph.Int, graph.Minutes(1), graph.Minutes(5), graph.Hours(1), graph.Days(1)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countExpr),

			"sum_amount": graph.Windowed(graph.Float, graph.Minutes(5), graph.Hours(1)).
				WithDefault(expr.Float(0.0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(sumExpr),

			"mean_amount": graph.Windowed(graph.Float, graph.Minutes(5), graph.Hours(1)).
				WithDefault(expr.Float(0.0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(meanExpr),

			"min_amount": graph.Windowed(graph.Float, graph.Minutes(5), graph.Hours(1)).
				WithDefault(expr.Float(0.0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(minExpr),

			"max_amount": graph.Windowed(graph.Float, graph.Minutes(5), graph.Hours(1)).
				WithDefault(expr.Float(0.0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(maxExpr),

			"stddev_amount": graph.Windowed(graph.Float, graph.Hours(1)).
				WithDefault(expr.Float(0.0)).
				WithBucketDuration(graph.Minutes(5)).
				WithExpr(stddevExpr),

			"distinct_items_approx": graph.Windowed(graph.Int, graph.Hours(1), graph.Days(1)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(10)).
				WithExpr(approxCountDistinctExpr),

			"top_event_ids": graph.Windowed(graph.List(graph.Int), graph.Hours(1)).
				WithBucketDuration(graph.Minutes(10)).
				WithExpr(maxByNExpr),

			"bottom_event_ids": graph.Windowed(graph.List(graph.Int), graph.Hours(1)).
				WithBucketDuration(graph.Minutes(10)).
				WithExpr(minByNExpr),

			"top_items": graph.Windowed(graph.List(graph.String), graph.Hours(1)).
				WithBucketDuration(graph.Minutes(10)).
				WithExpr(approxTopKExpr),

			// Test all filter operations
			"count_eq_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithEqFilter),

			"count_ne_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithNeFilter),

			"count_lt_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithLtFilter),

			"count_le_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithLeFilter),

			"count_gt_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithGtFilter),

			"count_ge_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithGeFilter),

			"count_and_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithAndFilter),

			"count_or_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithOrFilter),

			"count_not_filter": graph.Windowed(graph.Int, graph.Minutes(5)).
				WithDefault(expr.Int(0)).
				WithBucketDuration(graph.Minutes(1)).
				WithExpr(countWithNotFilter),
		})

	// Create definitions with all feature sets
	definitions := graph.Definitions{}.WithFeatureSets(standaloneFS, eventFS, userFS)

	// Convert to graph
	g, err := definitions.ToGraph()
	if err != nil {
		log.Fatalf("Failed to create graph: %v", err)
	}

	log.Printf("Successfully created graph with %d feature sets", len(g.FeatureSets))

	// Write the graph to graph.pb file in the current directory
	err = writeGraphToFile(g, "graph.pb")
	if err != nil {
		log.Fatalf("Failed to write graph to file: %v", err)
	}

	log.Printf("Graph written to graph.pb")

	// Print summary
	log.Printf("\nGraph Summary:")
	for _, fs := range g.FeatureSets {
		log.Printf("  Feature Set: %s (%d features)", fs.Name, len(fs.Features))
	}
}

// writeGraphToFile writes a protobuf Graph to a file
func writeGraphToFile(g *graphv1.Graph, filename string) error {
	data, err := proto.Marshal(g)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
