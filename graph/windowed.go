package graph

import (
	"fmt"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	"github.com/iancoleman/strcase"
	proto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Duration helper functions
func Seconds(n int64) time.Duration {
	return time.Duration(n) * time.Second
}

func Minutes(n int64) time.Duration {
	return time.Duration(n) * time.Minute
}

func Hours(n int64) time.Duration {
	return time.Duration(n) * time.Hour
}

func Days(n int64) time.Duration {
	return time.Duration(n) * 24 * time.Hour
}

func Weeks(n int64) time.Duration {
	return time.Duration(n) * 7 * 24 * time.Hour
}

type WindowedFeatureBuilder struct {
	proto           *graphv1.WindowedFeatureType
	ofType          FeatureBuilder
	Default         expr.Expr
	Expression      expr.Expr
	isMaterialized  bool
	Materialization MaterializationOptions
	err             error
}

func Windowed(ofType FeatureBuilder, windows ...time.Duration) *WindowedFeatureBuilder {
	durations := make([]*durationpb.Duration, len(windows))
	for i, d := range windows {
		durations[i] = durationpb.New(d)
	}
	return &WindowedFeatureBuilder{
		proto: &graphv1.WindowedFeatureType{
			WindowDurations: durations,
		},
		ofType: ofType,
	}
}

func (w *WindowedFeatureBuilder) WithDefault(expression expr.Expr) *WindowedFeatureBuilder {
	w.Default = expression
	return w
}

func (w *WindowedFeatureBuilder) WithExpr(expression expr.Expr) *WindowedFeatureBuilder {
	w.Expression = expression
	return w
}

// https://docs.chalk.ai/api-docs#windowed.materialization
type MaterializationOptions struct {
	// map window -> bucket duration
	// NOTE: OPPOSITE ORDER AS chalkpy if construction is being done via struct
	BucketDurations       map[time.Duration]time.Duration
	DefaultBucketDuration time.Duration
	// The period for which to use the continuous resolver, instead
	// of relying upon the last backfill. If not provided, and a continuous
	// resolver is provided, this will be set to backfill_lookback_duration.
	ContinuousBufferDuration time.Duration
	// A crontab or duration string to specify the schedule for back filling the
	// materialized aggregate.
	BackfillSchedule string
	// The lower bound of the first bucket. All buckets are aligned to this time.
	BucketStart time.Time

	// technically not part of materialization kwarg, but still useful to expose (for now)
	GroupBy []*graphv1.FeatureReference
	// The 'k' arg of approx_top_k.
	ApproxTopKArgK int64
	// The resolver to use for back-filling the materialized aggregate.
	// If not provided, the data will be back filled using the resolver
	// that would run for an offline query.
	BackfillResolver string
	// The amount of time before the start of the previous backfill
	// to consider when running the backfill resolver. Set this parameter
	// to the be equal to the latest arriving data in the backfill window.
	BackfillLookbackDuration time.Duration
	// The time at which to start back filling the materialized aggregate.
	// If not provided, the backfill consider the earliest available data returned
	// by the `backfill_resolver`.
	BackfillStartTime time.Time
	// The resolver to use for continuous updates to the materialized aggregate.
	// If not provided, the data will be updated using the resolver that would run
	// for an online query.
	ContinuousResolver string
}

func (w *WindowedFeatureBuilder) WithBucketDuration(duration time.Duration) *WindowedFeatureBuilder {
	w.isMaterialized = true
	w.Materialization.DefaultBucketDuration = duration
	return w
}

func (w *WindowedFeatureBuilder) WithDurationForWindows(duration time.Duration, windows ...time.Duration) *WindowedFeatureBuilder {
	w.isMaterialized = true
	for _, d := range windows {
		w.Materialization.BucketDurations[d] = duration
	}
	return w
}

// NOTE: this will override all previous materialization settings (WithBucketDuration,WithDurationForWindows)
func (w *WindowedFeatureBuilder) WithMaterialization(mo MaterializationOptions) *WindowedFeatureBuilder {
	w.isMaterialized = true
	w.Materialization = mo
	return w
}

func GetDefault[M ~map[K]V, K comparable, V any](m M, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
}

func combineFilters(filters []*expressionv1.LogicalExprNode) *expressionv1.LogicalExprNode {
	return &expressionv1.LogicalExprNode{
		ExprType: &expressionv1.LogicalExprNode_BinaryExpr{
			BinaryExpr: &expressionv1.BinaryExprNode{
				Op:       "and",
				Operands: filters,
			},
		},
	}
}

func aggregateOnFromExpr(namespace string, aggPtr *expr.AggregateExprImpl) (*graphv1.FeatureReference, []*expressionv1.LogicalExprNode, error) {
	s := aggPtr.Selection
	name := ""
	switch e := s.(type) {
	case nil:
		if aggPtr.Function != "count" {
			return nil, nil, fmt.Errorf("did not select an expression in non-count dataframe aggregation: %s", aggPtr.Function)
		} else { // return null feature reference
			return nil, nil, nil
		}
	case *expr.ColumnExpr:
		name = e.Name
	case *expr.IdentifierExpr:
		name = e.Name
	case *expr.GetAttributeExpr:
		name = e.Attribute
	default:
		return nil, nil, fmt.Errorf("invalid expression selected in dataframe aggregation: %T", s)
	}
	if name == "" {
		return nil, nil, fmt.Errorf("incorrectly extracted name from %s", s.String())
	}
	df := aggPtr.DataFrame.(*expr.DataFrameExprImpl)

	foreignNamespace := strcase.ToSnake(df.Name)
	filters := make([]*expressionv1.LogicalExprNode, len(aggPtr.Conditions))
	for i, cond := range aggPtr.Conditions {
		f, err := toFilterParsedProto(cond, foreignNamespace)
		if err != nil {
			return nil, nil, err
		}
		filters[i] = f
	}

	return &graphv1.FeatureReference{
		Name:      name,
		Namespace: namespace,
		Df: &graphv1.DataFrameType{
			RootNamespace: foreignNamespace,
			Filter:        combineFilters(filters),
		},
	}, filters, nil
}

func (w *WindowedFeatureBuilder) ToProtos(fieldName string, namespace string) ([]*graphv1.FeatureType, error) {
	if w.err != nil {
		return []*graphv1.FeatureType{}, w.err
	}

	numPeriods := len(w.proto.WindowDurations)
	m := w.Materialization
	features := make([]*graphv1.FeatureType, numPeriods+1)

	for i := range numPeriods {
		d := w.proto.WindowDurations[i]
		scalarPtr, ok := w.ofType.(*ScalarFeatureBuilder)
		if !ok {
			return nil, fmt.Errorf("windowed features must be scalar")
		}
		aggPtr, ok := w.Expression.(*expr.AggregateExprImpl)
		if !ok {
			return nil, fmt.Errorf("windowed feature expression must be dataframe aggregation")
		}
		aggOn, filters, err := aggregateOnFromExpr(namespace, aggPtr)
		if err != nil {
			return nil, err
		}

		// Add duration suffix to the scalar feature name
		durationSeconds := int64(d.AsDuration().Seconds())
		suffixedFieldName := fmt.Sprintf("%s__%d__", fieldName, durationSeconds)
		f := scalarPtr.ToProto(suffixedFieldName, namespace)
		f.Type.(*graphv1.FeatureType_Scalar).Scalar.WindowInfo = &graphv1.WindowInfo{
			Duration: d,
			Aggregation: &graphv1.WindowAggregation{
				Namespace: namespace,
				BucketDuration: durationpb.New(GetDefault(
					m.BucketDurations,
					d.AsDuration(),
					m.DefaultBucketDuration,
				)),
				ContinuousBufferDuration: durationpb.New(m.ContinuousBufferDuration),
				BackfillSchedule:         MaybeStr(m.BackfillSchedule),
				BucketStart:              timestamppb.New(m.BucketStart),
				Aggregation:              aggPtr.Function,
				AggregateOn:              aggOn,
				Filters:                  filters,
				ArrowType:                scalarPtr.proto.ArrowType,

				// set from MaterializationOptions
				//GroupBy:                  m.GroupBy,
				ApproxTopKArgK:           &m.ApproxTopKArgK,
				BackfillResolver:         &m.BackfillResolver,
				BackfillLookbackDuration: durationpb.New(m.BackfillLookbackDuration),
				BackfillStartTime:        timestamppb.New(m.BackfillStartTime),
				ContinuousResolver:       &m.ContinuousResolver,
			},
		}
		f.Type.(*graphv1.FeatureType_Scalar).Scalar.Expression = expr.ToProto(w.Expression)
		if w.Default != nil {
			lit, ok := w.Default.(*expr.LiteralExpr)
			if !ok {
				return []*graphv1.FeatureType{}, w.err
			}
			f.Type.(*graphv1.FeatureType_Scalar).Scalar.DefaultValue = lit.ScalarValue
		}
		features[i] = f
	}

	windowed := proto.Clone(w.proto).(*graphv1.WindowedFeatureType)
	windowed.Name = fieldName
	windowed.Namespace = namespace
	windowed.AttributeName = fieldName
	windowed.UnversionedAttributeName = fieldName
	features[numPeriods] = &graphv1.FeatureType{
		Type: &graphv1.FeatureType_Windowed{
			Windowed: windowed,
		},
	}

	return features, nil
}
