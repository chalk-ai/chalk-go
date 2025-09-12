package graph

import (
	"fmt"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
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
	// NOTE: OPPOSITE OF chalkpy
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
	Filters        []expr.Expr
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
	w.Materialization.DefaultBucketDuration = duration
	return w
}

func (w *WindowedFeatureBuilder) WithDurationForWindows(duration time.Duration, windows ...time.Duration) *WindowedFeatureBuilder {
	for _, d := range windows {
		w.Materialization.BucketDurations[d] = duration
	}
	return w
}

func (w *WindowedFeatureBuilder) WithMaterialization(mo MaterializationOptions) *WindowedFeatureBuilder {
	w.Materialization = mo
	return w
}

func GetDefault[M ~map[K]V, K comparable, V any](m M, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
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
		f := scalarPtr.ToProto(fieldName, namespace)
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
				//do we get this from the expression?
				//Aggregation    string
				//AggregateOn    *graphv1.FeatureReference
				//ArrowType      *v11.ArrowType
			},
		}
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
