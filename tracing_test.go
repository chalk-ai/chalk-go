package chalk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	assert "github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
)

// scalarsQueryHandler serves a canned OnlineQueryBulk response so that the
// response-side spans have real Arrow bytes to deserialize.
type scalarsQueryHandler struct {
	enginev1connect.UnimplementedQueryServiceHandler
	scalarsData []byte
	queryId     string
	failWith    error
}

func (h *scalarsQueryHandler) OnlineQueryBulk(
	_ context.Context,
	_ *connect.Request[commonv1.OnlineQueryBulkRequest],
) (*connect.Response[commonv1.OnlineQueryBulkResponse], error) {
	if h.failWith != nil {
		return nil, h.failWith
	}
	return connect.NewResponse(&commonv1.OnlineQueryBulkResponse{
		ScalarsData: h.scalarsData,
		ResponseMeta: &commonv1.OnlineQueryMetadata{
			QueryId: h.queryId,
		},
	}), nil
}

func startScalarsQueryServer(t *testing.T, handler *scalarsQueryHandler) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	queryPath, queryHandler := enginev1connect.NewQueryServiceHandler(handler)
	mux.Handle(queryPath, queryHandler)
	authPath, authHandler := serverv1connect.NewAuthServiceHandler(&minimalAuthHandler{})
	mux.Handle(authPath, authHandler)
	return startUnencryptedHTTP2Server(t, mux)
}

func newTracedGRPCClient(t *testing.T, serverURL string, provider trace.TracerProvider) GRPCClient {
	t.Helper()
	client, err := NewGRPCClient(context.Background(), &GRPCClientConfig{
		ClientId:                   "test-client-id",
		ClientSecret:               "test-client-secret",
		ApiServer:                  serverURL,
		QueryServer:                serverURL,
		EnvironmentId:              "test-env",
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
		TracerProvider:             provider,
		JWT: &serverv1.GetTokenResponse{
			AccessToken: "mock-test-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		},
	})
	assert.NoError(t, err)
	return client
}

// userScalarsData builds the Arrow payload for a single-row `user.id`
// response, matching the `user` struct used elsewhere in these tests.
func userScalarsData(t *testing.T) []byte {
	t.Helper()
	table, err := MakeFeatureTable(map[any]any{
		"user.id": []int64{1},
	})
	assert.NoError(t, err)
	tableBytes, err := table.ToBytes()
	assert.NoError(t, err)
	return tableBytes
}

type recordedSpans map[string]sdktrace.ReadOnlySpan

func recorded(t *testing.T, recorder *tracetest.SpanRecorder) recordedSpans {
	t.Helper()
	byName := recordedSpans{}
	for _, span := range recorder.Ended() {
		byName[span.Name()] = span
	}
	return byName
}

func (s recordedSpans) get(t *testing.T, name string) sdktrace.ReadOnlySpan {
	t.Helper()
	span, ok := s[name]
	assert.Truef(t, ok, "expected a span named %q, got %v", name, s.names())
	return span
}

func (s recordedSpans) names() []string {
	names := make([]string, 0, len(s))
	for name := range s {
		names = append(names, name)
	}
	return names
}

func assertParent(t *testing.T, child sdktrace.ReadOnlySpan, parent sdktrace.ReadOnlySpan) {
	t.Helper()
	assert.Equalf(
		t,
		parent.SpanContext().SpanID().String(),
		child.Parent().SpanID().String(),
		"expected %q to be a child of %q", child.Name(), parent.Name(),
	)
}

func spanAttribute(t *testing.T, span sdktrace.ReadOnlySpan, key attribute.Key) attribute.Value {
	t.Helper()
	for _, attr := range span.Attributes() {
		if attr.Key == key {
			return attr.Value
		}
	}
	assert.FailNowf(t, "missing attribute", "span %q has no attribute %q", span.Name(), key)
	return attribute.Value{}
}

func TestOnlineQueryBulkEmitsSpans(t *testing.T) {
	t.Parallel()
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })

	server := startScalarsQueryServer(t, &scalarsQueryHandler{
		scalarsData: userScalarsData(t),
		queryId:     "query-id-1",
	})
	client := newTracedGRPCClient(t, server.URL, provider)

	ctx, callerSpan := provider.Tracer("test").Start(context.Background(), "caller")
	result, err := client.OnlineQueryBulk(ctx, OnlineQueryParams{
		QueryName:        "my_query",
		QueryNameVersion: "v1",
		CorrelationId:    "correlation-1",
	}.
		WithInput("user.id", []int64{1}).
		WithOutputs("user.id"))
	assert.NoError(t, err)

	var users []user
	assert.NoError(t, result.UnmarshalInto(&users))
	callerSpan.End()

	spans := recorded(t, recorder)
	query := spans.get(t, spanOnlineQueryBulk)
	buildRequest := spans.get(t, spanBuildRequest)
	unmarshalInto := spans.get(t, spanUnmarshalInto)

	// The request-side spans nest under the query span.
	assertParent(t, query, spans.get(t, "caller"))
	assertParent(t, buildRequest, query)
	assertParent(t, spans.get(t, spanResolveParams), buildRequest)
	assertParent(t, spans.get(t, spanSerializeInputs), buildRequest)
	assertParent(t, spans.get(t, spanRPC), query)

	// Deserialization happens after OnlineQueryBulk returns, so it is a
	// sibling of the query rather than a child of it.
	assertParent(t, unmarshalInto, spans.get(t, "caller"))
	assertParent(t, spans.get(t, spanDeserializeScalars), unmarshalInto)
	assertParent(t, spans.get(t, spanUnmarshalTable), unmarshalInto)

	assert.Equal(t, "my_query", spanAttribute(t, query, attrQueryName).AsString())
	assert.Equal(t, "v1", spanAttribute(t, query, attrQueryNameVersion).AsString())
	assert.Equal(t, "correlation-1", spanAttribute(t, query, attrCorrelationID).AsString())
	assert.Equal(t, "test-env", spanAttribute(t, query, attrEnvironment).AsString())
	assert.Equal(t, "query-id-1", spanAttribute(t, query, attrOperationID).AsString())
	assert.Equal(t, int64(0), spanAttribute(t, query, attrNumErrors).AsInt64())

	assert.Equal(t, int64(1), spanAttribute(t, spans.get(t, spanResolveParams), attrNumInputs).AsInt64())
	assert.Equal(t, int64(1), spanAttribute(t, spans.get(t, spanResolveParams), attrNumOutputs).AsInt64())
	assert.Equal(t, int64(1), spanAttribute(t, spans.get(t, spanSerializeInputs), attrNumRows).AsInt64())
	assert.Positive(t, spanAttribute(t, spans.get(t, spanSerializeInputs), attrInputsBytes).AsInt64())
	assert.Positive(t, spanAttribute(t, spans.get(t, spanRPC), attrScalarsBytes).AsInt64())
	assert.Equal(t, int64(1), spanAttribute(t, unmarshalInto, attrNumRows).AsInt64())

	for _, span := range recorder.Ended() {
		assert.Equalf(t, codes.Unset, span.Status().Code, "span %q should not be an error", span.Name())
	}
}

func TestOnlineQueryBulkGetRowEmitsSpans(t *testing.T) {
	t.Parallel()
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })

	server := startScalarsQueryServer(t, &scalarsQueryHandler{scalarsData: userScalarsData(t)})
	client := newTracedGRPCClient(t, server.URL, provider)

	result, err := client.OnlineQueryBulk(
		context.Background(),
		OnlineQueryParams{}.WithInput("user.id", []int64{1}).WithOutputs("user.id"),
	)
	assert.NoError(t, err)

	row, err := result.GetRow(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), row.Features["user.id"].Value)

	_, err = result.GetTable()
	assert.NoError(t, err)

	spans := recorded(t, recorder)
	getRow := spans.get(t, spanGetRow)
	assertParent(t, spans.get(t, spanDeserializeScalars), getRow)
	assertParent(t, spans.get(t, spanExtractFeatures), getRow)
	assert.Equal(t, int64(0), spanAttribute(t, getRow, attrRowIndex).AsInt64())
	assert.Equal(t, int64(1), spanAttribute(t, spans.get(t, spanGetTable), attrNumRows).AsInt64())
}

func TestOnlineQueryBulkSpansRecordRPCErrors(t *testing.T) {
	t.Parallel()
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })

	server := startScalarsQueryServer(t, &scalarsQueryHandler{
		failWith: connect.NewError(connect.CodeUnavailable, errors.New("engine is down")),
	})
	client := newTracedGRPCClient(t, server.URL, provider)

	_, err := client.OnlineQueryBulk(
		context.Background(),
		OnlineQueryParams{}.WithInput("user.id", []int64{1}).WithOutputs("user.id"),
	)
	assert.Error(t, err)

	spans := recorded(t, recorder)
	rpc := spans.get(t, spanRPC)
	assert.Equal(t, codes.Error, rpc.Status().Code)
	assert.Equal(t, codes.Error, spans.get(t, spanOnlineQueryBulk).Status().Code)
	assert.NotEmpty(t, rpc.Events(), "the RPC error should be recorded as a span event")
	assert.Equal(t, codes.Unset, spans.get(t, spanBuildRequest).Status().Code)
}

// Tracing is opt-in: without a TracerProvider the client must not emit
// spans, and must not reach for a globally registered provider either.
func TestTracingIsOffWithoutTracerProvider(t *testing.T) {
	t.Parallel()

	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })

	server := startScalarsQueryServer(t, &scalarsQueryHandler{scalarsData: userScalarsData(t)})
	client := newTracedGRPCClient(t, server.URL, nil)

	// A span from an unrelated provider is active on the context, standing
	// in for an application that traces its own work: the query must still
	// not add spans to it.
	ctx, callerSpan := provider.Tracer("test").Start(context.Background(), "caller")
	result, err := client.OnlineQueryBulk(
		ctx,
		OnlineQueryParams{}.WithInput("user.id", []int64{1}).WithOutputs("user.id"),
	)
	assert.NoError(t, err)

	var users []user
	assert.NoError(t, result.UnmarshalInto(&users))
	assert.Len(t, users, 1)
	_, err = result.GetTable()
	assert.NoError(t, err)
	callerSpan.End()

	assert.Len(t, recorder.Ended(), 1, "expected only the caller's own span")
	assert.Equal(t, "caller", recorder.Ended()[0].Name())
}

// With tracing off there is no tracer at all, so spans are never recorded
// and never become a parent for anything downstream.
func TestNoopTracerDoesNotRecord(t *testing.T) {
	t.Parallel()
	assert.Nil(t, newTracer(nil))

	_, span := startSpan(context.Background(), newTracer(nil), spanOnlineQueryBulk)
	assert.False(t, span.IsRecording())
	assert.False(t, span.SpanContext().IsValid())
	span.End()
}

// Turning tracing off must leave the caller's context exactly as it was.
// Wrapping it in a non-recording span would make trace.SpanFromContext
// return our span instead of theirs, quietly swallowing attributes that
// other instrumentation writes to the active span.
func TestDisabledTracingLeavesContextUntouched(t *testing.T) {
	t.Parallel()
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })

	callerCtx, callerSpan := provider.Tracer("test").Start(context.Background(), "caller")
	ctx, span := startSpan(callerCtx, newTracer(nil), spanOnlineQueryBulk)
	endSpan(span, nil)

	assert.Same(t, callerSpan, trace.SpanFromContext(ctx), "the caller's span must stay active")
	assert.True(t, trace.SpanFromContext(ctx).IsRecording())
}

// Tracing that is off should also be free. This guards the call sites: a
// new span site that builds its attributes unconditionally would allocate
// here even with no provider installed.
func TestDisabledTracingDoesNotAllocate(t *testing.T) {
	tracer := newTracer(nil)
	ctx := context.Background()

	allocs := testing.AllocsPerRun(1000, func() {
		// The shape of every span site in the query path: start, then
		// attributes only if something is listening.
		var attrs []attribute.KeyValue
		if tracer != nil {
			attrs = []attribute.KeyValue{attrQueryName.String("q")}
		}
		spanCtx, span := startSpan(ctx, tracer, spanOnlineQueryBulk, attrs...)
		if span.IsRecording() {
			span.SetAttributes(attrNumRows.Int64(100))
		}
		_, child := startSpan(spanCtx, tracer, spanRPC)
		if child.IsRecording() {
			child.SetAttributes(attrScalarsBytes.Int(65536))
		}
		endSpan(child, nil)
		endSpan(span, nil)
	})
	assert.Zero(t, allocs, "tracing off must not allocate")
}
