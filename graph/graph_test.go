package graph

import (
	"testing"

	"github.com/chalk-ai/chalk-go/expr"
)

func TestSimpleAdd(t *testing.T) {
	Definitions{}.WithFeatureSets(
		FeatureSet{
			Name: "Example",
		}.WithFeature("id",
			Int(),
		).WithFeature("val",
			Int().Expr(__("id").Add(expr.Int64(1))),
		),
	).ToGraph()
}
