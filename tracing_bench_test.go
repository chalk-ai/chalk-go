package chalk

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// discardExporter keeps nothing, so these benchmarks measure the cost of
// producing spans rather than the cost of retaining them.
type discardExporter struct{}

func (discardExporter) ExportSpans(context.Context, []sdktrace.ReadOnlySpan) error { return nil }
func (discardExporter) Shutdown(context.Context) error                             { return nil }

// benchQueryPath is the exact span shape of one OnlineQueryBulk plus
// UnmarshalInto: 9 spans with the attributes chalk-go sets on them.
func benchQueryPath(tr trace.Tracer) {
	ctx := context.Background()

	var queryAttrs []attribute.KeyValue
	if tr != nil {
		queryAttrs = []attribute.KeyValue{
			attrQueryName.String("loan_application_model"),
			attrQueryNameVersion.String("v1"),
			attrCorrelationID.String("abc-123"),
			attrEnvironment.String("prod"),
		}
	}
	qctx, query := startSpan(ctx, tr, spanOnlineQueryBulk, queryAttrs...)

	bctx, build := startSpan(qctx, tr, spanBuildRequest)
	if build.IsRecording() {
		build.SetAttributes(attrNumOutputs.Int(12), attrInputsBytes.Int(4096))
	}
	_, resolve := startSpan(bctx, tr, spanResolveParams)
	if resolve.IsRecording() {
		resolve.SetAttributes(attrNumInputs.Int(3), attrNumOutputs.Int(12))
	}
	endSpan(resolve, nil)
	_, serialize := startSpan(bctx, tr, spanSerializeInputs)
	if serialize.IsRecording() {
		serialize.SetAttributes(attrNumInputs.Int(3), attrNumRows.Int(100), attrInputsBytes.Int(4096))
	}
	endSpan(serialize, nil)
	endSpan(build, nil)

	_, rpc := startSpan(qctx, tr, spanRPC)
	if rpc.IsRecording() {
		rpc.SetAttributes(attrInputsBytes.Int(4096), attrScalarsBytes.Int(65536), attrNumErrors.Int(0))
	}
	endSpan(rpc, nil)
	endSpan(query, nil)

	uctx, unmarshal := startSpan(ctx, tr, spanUnmarshalInto)
	if unmarshal.IsRecording() {
		unmarshal.SetAttributes(attrScalarsBytes.Int(65536), attrNumRows.Int64(100))
	}
	_, deser := startSpan(uctx, tr, spanDeserializeScalars)
	if deser.IsRecording() {
		deser.SetAttributes(attrScalarsBytes.Int(65536), attrNumRows.Int64(100))
	}
	endSpan(deser, nil)
	_, table := startSpan(uctx, tr, spanUnmarshalTable)
	if table.IsRecording() {
		table.SetAttributes(attrNumRows.Int64(100))
	}
	endSpan(table, nil)
	endSpan(unmarshal, nil)
}

// BenchmarkTracingOff is the cost paid by callers who never set a
// TracerProvider. It must stay at zero allocations; see
// TestDisabledTracingDoesNotAllocate for the enforcing test.
func BenchmarkTracingOff(b *testing.B) {
	tr := newTracer(nil)
	b.ReportAllocs()
	for b.Loop() {
		benchQueryPath(tr)
	}
}

// Sampled out: the 99% case for anyone running head sampling below 100%.
func BenchmarkTracingDropped(b *testing.B) {
	tr := newTracer(sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(discardExporter{}),
		sdktrace.WithSampler(sdktrace.NeverSample()),
	))
	b.ReportAllocs()
	for b.Loop() {
		benchQueryPath(tr)
	}
}

// Fully recorded and exported.
func BenchmarkTracingSampled(b *testing.B) {
	tr := newTracer(sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(discardExporter{}),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	))
	b.ReportAllocs()
	for b.Loop() {
		benchQueryPath(tr)
	}
}
