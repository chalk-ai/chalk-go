package chalk

import (
	"context"

	"github.com/chalk-ai/chalk-go/internal"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// instrumentationName identifies chalk-go as the source of the spans it
// emits. It shows up as `otel.library.name` on every span below.
const instrumentationName = "github.com/chalk-ai/chalk-go"

// Names of the spans that chalk-go emits. `chalk.online_query_bulk` is the
// parent of the request-side spans; the response-side spans are parented to
// whatever span was active when OnlineQueryBulk was called, since they are
// produced lazily, after the query itself has finished.
const (
	spanOnlineQueryBulk    = "chalk.online_query_bulk"
	spanBuildRequest       = "chalk.online_query_bulk.build_request"
	spanResolveParams      = "chalk.online_query_bulk.resolve_params"
	spanSerializeInputs    = "chalk.online_query_bulk.serialize_inputs"
	spanRPC                = "chalk.online_query_bulk.rpc"
	spanUnmarshalInto      = "chalk.online_query_bulk.unmarshal_into"
	spanDeserializeScalars = "chalk.online_query_bulk.deserialize_scalars"
	spanUnmarshalTable     = "chalk.online_query_bulk.unmarshal_table"
	spanGetTable           = "chalk.online_query_bulk.get_table"
	spanGetRow             = "chalk.online_query_bulk.get_row"
	spanExtractFeatures    = "chalk.online_query_bulk.extract_features"
)

// Attribute keys set on the spans above.
const (
	attrQueryName        = attribute.Key("chalk.query_name")
	attrQueryNameVersion = attribute.Key("chalk.query_name_version")
	attrCorrelationID    = attribute.Key("chalk.correlation_id")
	attrBranch           = attribute.Key("chalk.branch")
	attrEnvironment      = attribute.Key("chalk.environment")
	attrResourceGroup    = attribute.Key("chalk.resource_group")
	attrDeploymentTag    = attribute.Key("chalk.deployment_tag")
	attrNumInputs        = attribute.Key("chalk.num_inputs")
	attrNumOutputs       = attribute.Key("chalk.num_outputs")
	attrNumRows          = attribute.Key("chalk.num_rows")
	attrInputsBytes      = attribute.Key("chalk.inputs_feather_bytes")
	attrScalarsBytes     = attribute.Key("chalk.scalars_data_bytes")
	attrGroupsBytes      = attribute.Key("chalk.groups_data_bytes")
	attrNumErrors        = attribute.Key("chalk.num_errors")
	attrOperationID      = attribute.Key("chalk.operation_id")
	attrRowIndex         = attribute.Key("chalk.row_index")
)

// disabledSpan stands in for a span when tracing is off. It is a shared
// value rather than one per call: it holds no state, records nothing, and
// is never put into a context.
var disabledSpan trace.Span = noop.Span{}

// newTracer returns the tracer that a client should emit spans on, or nil
// when tracing is off.
//
// Tracing is off unless a TracerProvider is supplied, so a client that was
// not asked to trace emits nothing, even in a process that has an
// OpenTelemetry provider registered globally for its own use. Callers opt
// in by setting GRPCClientConfig.TracerProvider, typically to
// otel.GetTracerProvider().
func newTracer(provider trace.TracerProvider) trace.Tracer {
	if provider == nil {
		return nil
	}
	return provider.Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(internal.Version()),
	)
}

// startSpan begins a span on tracer.
//
// A nil tracer means tracing is off, and is also what results built
// outside of a client carry (see NewGRPCOnlineQueryBulkResult). In that
// case the context is handed back untouched: nothing is allocated, and the
// caller's own span stays the active one, so instrumentation elsewhere in
// the process is unaffected by chalk-go having declined to trace.
func startSpan(
	ctx context.Context,
	tracer trace.Tracer,
	name string,
	attrs ...attribute.KeyValue,
) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if tracer == nil {
		return ctx, disabledSpan
	}
	return tracer.Start(ctx, name, trace.WithAttributes(attrs...), trace.WithSpanKind(trace.SpanKindInternal))
}

// endSpan finishes span, recording err on it if there is one. Server errors
// returned alongside partial results are recorded like any other error: the
// span is what surfaces them, so a partially failed query is still visible
// as a failure in the trace.
func endSpan(span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()
}
