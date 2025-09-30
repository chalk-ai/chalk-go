package graph

import (
	"fmt"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
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

func Years(n int64) time.Duration {
	return time.Duration(n) * 365 * 24 * time.Hour
}

type WindowedFeatureBuilder struct {
	proto           *graphv1.WindowedFeatureType
	ofType          FeatureBuilder
	Default         expr.Expr
	Expression      expr.Expr
	MaxStaleness    time.Duration
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

func (w *WindowedFeatureBuilder) WithMaxStaleness(d time.Duration) *WindowedFeatureBuilder {
	w.MaxStaleness = d
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

type ParsedAggregation struct {
	aggregateOn      *graphv1.FeatureReference
	filters          []*expressionv1.LogicalExprNode
	foreignNamespace string
}

func extractFromExpr(namespace string, features []*graphv1.FeatureType, aggPtr *expr.AggregateExprImpl) (*ParsedAggregation, error) {
	df := aggPtr.DataFrame.(*expr.DataFrameExprImpl)
	dfName := strcase.ToSnake(df.Name)

	foreignNamespace := ""
	for _, f := range features {
		switch t := f.Type.(type) {
		case *graphv1.FeatureType_HasMany:
			if t.HasMany.Name == dfName {
				foreignNamespace = t.HasMany.ForeignNamespace
			}
		}
	}
	if foreignNamespace == "" {
		return nil, fmt.Errorf("could not find definition of dataframe %s prior to windowed aggregate", dfName)
	}

	filters := make([]*expressionv1.LogicalExprNode, len(aggPtr.Conditions))
	for i, cond := range aggPtr.Conditions {
		f, err := toFilterParsedProto(cond, foreignNamespace)
		if err != nil {
			return nil, err
		}
		filters[i] = f
	}

	parsedAgg := &ParsedAggregation{
		filters:          filters,
		foreignNamespace: foreignNamespace,
	}

	s := aggPtr.Selection
	name := ""
	switch e := s.(type) {
	case nil:
		if aggPtr.Function != "count" {
			return nil, fmt.Errorf("did not select an expression in non-count dataframe aggregation: %s", aggPtr.Function)
		} else { // return null feature reference
			return parsedAgg, nil
		}
	case *expr.ColumnExpr:
		name = e.Name
	case *expr.IdentifierExpr:
		name = e.Name
	case *expr.GetAttributeExpr:
		name = e.Attribute
	default:
		return nil, fmt.Errorf("invalid expression selected in dataframe aggregation: %T", s)
	}
	if name == "" {
		return nil, fmt.Errorf("incorrectly extracted name from %s", s.String())
	}

	parsedAgg.aggregateOn = &graphv1.FeatureReference{
		Name:      name,
		Namespace: foreignNamespace,
	}
	return parsedAgg, nil
}

func durationProto(duration time.Duration) *durationpb.Duration {
	if duration != 0 {
		return durationpb.New(duration)
	}
	return nil
}

func timeProto(time time.Time) *timestamppb.Timestamp {
	if !time.IsZero() {
		return timestamppb.New(time)
	}
	return nil
}

func stringProto(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}

func int64Proto(i int64) *int64 {
	if i != 0 {
		return &i
	}
	return nil
}

func (w *WindowedFeatureBuilder) AppendFeatures(features []*graphv1.FeatureType, fieldName string, namespace string) ([]*graphv1.FeatureType, error) {
	if w.err != nil {
		return []*graphv1.FeatureType{}, w.err
	}

	numPeriods := len(w.proto.WindowDurations)
	m := w.Materialization

	scalarPtr, ok := w.ofType.(*ScalarFeatureBuilder)
	if !ok {
		return nil, fmt.Errorf("windowed features must be scalar")
	}
	aggPtr, ok := w.Expression.(*expr.AggregateExprImpl)
	if !ok {
		return nil, fmt.Errorf("windowed feature expression must be dataframe aggregation")
	}
	parsedAgg, err := extractFromExpr(namespace, features, aggPtr)
	if err != nil {
		return nil, err
	}

	exprProto := expr.ToProto(w.Expression)
	var defaultProto *arrowv1.ScalarValue
	if w.Default != nil {
		lit, ok := w.Default.(*expr.LiteralExpr)
		if !ok {
			return []*graphv1.FeatureType{}, w.err
		}
		defaultProto = lit.ScalarValue
	}

	featureForWindow := func(suffixedFieldName string, d time.Duration, dp *durationpb.Duration) *graphv1.FeatureType {
		f := scalarPtr.ToProto(suffixedFieldName, namespace)
		scalar := f.Type.(*graphv1.FeatureType_Scalar).Scalar
		scalar.MaxStalenessDuration = durationpb.New(w.MaxStaleness)

		scalar.WindowInfo = &graphv1.WindowInfo{
			Duration: dp,
			Aggregation: &graphv1.WindowAggregation{
				Namespace:   parsedAgg.foreignNamespace,
				Aggregation: aggPtr.Function,
				AggregateOn: parsedAgg.aggregateOn,
				Filters:     parsedAgg.filters,
				ArrowType:   scalarPtr.proto.ArrowType,

				// set from MaterializationOptions
				BucketDuration: durationpb.New(GetDefault(
					m.BucketDurations,
					d,
					m.DefaultBucketDuration,
				)),
				ContinuousBufferDuration: durationProto(m.ContinuousBufferDuration),
				BackfillSchedule:         stringProto(m.BackfillSchedule),
				BucketStart:              timeProto(m.BucketStart),
				ApproxTopKArgK:           int64Proto(m.ApproxTopKArgK),
				BackfillResolver:         stringProto(m.BackfillResolver),
				BackfillLookbackDuration: durationProto(m.BackfillLookbackDuration),
				BackfillStartTime:        timeProto(m.BackfillStartTime),
				ContinuousResolver:       stringProto(m.ContinuousResolver),
			},
		}
		scalar.Expression = exprProto
		scalar.DefaultValue = defaultProto
		return f
	}

	windowed := proto.Clone(w.proto).(*graphv1.WindowedFeatureType)
	windowed.Name = fieldName
	windowed.Namespace = namespace
	windowed.AttributeName = fieldName
	windowed.UnversionedAttributeName = fieldName

	oldLength := len(features)
	if numPeriods == 0 {
		newFeatures := make([]*graphv1.FeatureType, oldLength+2)
		copy(newFeatures, features)
		suffixedFieldName := fmt.Sprintf("%s__all__", fieldName)
		duration := Years(10)
		durationProto := durationpb.New(duration)
		newFeatures[oldLength] = featureForWindow(suffixedFieldName, duration, durationProto)

		windowed.WindowDurations = []*durationpb.Duration{durationProto}
		newFeatures[oldLength+1] = &graphv1.FeatureType{
			Type: &graphv1.FeatureType_Windowed{
				Windowed: windowed,
			},
		}
		return newFeatures, nil
	} else {
		newFeatures := make([]*graphv1.FeatureType, oldLength+numPeriods+1)
		copy(newFeatures, features)
		for i := range numPeriods {
			durationProto := w.proto.WindowDurations[i]
			duration := durationProto.AsDuration()

			// Add duration suffix to the scalar feature name
			durationSeconds := int64(duration.Seconds())
			suffixedFieldName := fmt.Sprintf("%s__%d__", fieldName, durationSeconds)
			newFeatures[oldLength+i] = featureForWindow(suffixedFieldName, duration, durationProto)
		}

		newFeatures[oldLength+numPeriods] = &graphv1.FeatureType{
			Type: &graphv1.FeatureType_Windowed{
				Windowed: windowed,
			},
		}
		return newFeatures, nil
	}
}
