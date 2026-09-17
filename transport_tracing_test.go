package chalk

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"connectrpc.com/connect"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	assert "github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

type transportHTTPFunc func(*http.Request) (*http.Response, error)

func (f transportHTTPFunc) Do(req *http.Request) (*http.Response, error) { return f(req) }

func transportProvider(t *testing.T) (*sdktrace.TracerProvider, *tracetest.SpanRecorder) {
	t.Helper()
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { assert.NoError(t, provider.Shutdown(context.Background())) })
	return provider, recorder
}

func transportAttrs(span sdktrace.ReadOnlySpan) map[string]attribute.Value {
	attrs := make(map[string]attribute.Value)
	for _, attr := range span.Attributes() {
		attrs[string(attr.Key)] = attr.Value
	}
	return attrs
}

func TestTransportTracingTimingsAndBodyLifetime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		provider, recorder := transportProvider(t)
		tracer := newTracer(provider)
		ctx, rpc := startSpan(context.Background(), tracer, spanRPC)
		ctx, state := startTransportTrace(ctx, true, rpc)
		var callerHooks int
		ctx = httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
			GotConn: func(httptrace.GotConnInfo) { callerHooks++ },
		})
		time.Sleep(2 * time.Millisecond) // middleware and protobuf encoding
		client := &transportTracingClient{tracer: tracer, HTTPClient: transportHTTPFunc(func(req *http.Request) (*http.Response, error) {
			h := httptrace.ContextClientTrace(req.Context())
			h.GetConn("example.test:443")
			h.DNSStart(httptrace.DNSStartInfo{})
			time.Sleep(time.Millisecond)
			h.DNSDone(httptrace.DNSDoneInfo{})
			h.ConnectStart("tcp", "127.0.0.1:443")
			time.Sleep(2 * time.Millisecond)
			h.ConnectDone("tcp", "127.0.0.1:443", nil)
			h.TLSHandshakeStart()
			time.Sleep(3 * time.Millisecond)
			h.TLSHandshakeDone(tls.ConnectionState{}, nil)
			h.GotConn(httptrace.GotConnInfo{Reused: true})
			time.Sleep(4 * time.Millisecond)
			h.WroteHeaders()
			h.WroteRequest(httptrace.WroteRequestInfo{})
			time.Sleep(5 * time.Millisecond)
			h.GotFirstResponseByte()
			return &http.Response{
				StatusCode: 200, Proto: "HTTP/2.0", ProtoMajor: 2, Body: io.NopCloser(strings.NewReader("response")),
				Header: http.Header{"X-Envoy-Upstream-Service-Time": {"17"}, "Set-Cookie": {"private-cookie"}},
			}, nil
		})}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://example.test/query?secret=private-query", strings.NewReader("private-request"))
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer private-token")
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Empty(t, recorder.Ended(), "Do returns at headers; the HTTP span must remain open")
		time.Sleep(6 * time.Millisecond)
		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "response", string(body))
		firstEOF := time.Now()
		assert.Empty(t, recorder.Ended(), "EOF does not imply Connect has finished decoding/closing")
		time.Sleep(7 * time.Millisecond)
		// Connect may read again after EOF while finishing the response.
		n, err := res.Body.Read(make([]byte, 1))
		assert.Zero(t, n)
		assert.ErrorIs(t, err, io.EOF)
		assert.Empty(t, recorder.Ended(), "repeated EOF must not end the span")
		assert.NoError(t, res.Body.Close())
		state.finish(rpc)
		rpc.End()
		assert.Equal(t, 1, callerHooks)
		spans := recorded(t, recorder)
		httpSpan := spans.get(t, spanHTTP)
		assertParent(t, httpSpan, spans.get(t, spanRPC))
		attrs := transportAttrs(httpSpan)
		for name, expected := range map[string]float64{
			"dns_ms": 1, "connect_ms": 2, "tls_ms": 3, "get_connection_ms": 6,
			"write_request_ms": 4, "wait_response_headers_ms": 5,
			"response_headers_ms": 15, "response_read_and_decode_ms": 13,
			"envoy_upstream_service_time_ms": 17,
		} {
			assert.Equal(t, expected, attrs[transportAttr+name].AsFloat64(), name)
		}
		assert.Equal(t, float64(2), transportAttrs(spans.get(t, spanRPC))[transportAttr+"pre_http_ms"].AsFloat64())
		assert.Equal(t, int64(8), attrs[transportAttr+"response_body_bytes"].AsInt64())
		assert.True(t, attrs[transportAttr+"response_body_complete"].AsBool())
		assert.False(t, attrs[transportAttr+"response_body_failed"].AsBool())
		eofEvents := 0
		for _, event := range httpSpan.Events() {
			if event.Name == "http.body_eof" {
				eofEvents++
				assert.Equal(t, firstEOF, event.Time, "preserve the first EOF timestamp")
			}
		}
		assert.Equal(t, 1, eofEvents, "emit one EOF event per response")
		assert.True(t, attrs[transportAttr+"connection_reused"].AsBool())
		assert.Equal(t, "private-cookie", res.Header.Get("Set-Cookie"), "response metadata must remain usable")
		for _, span := range recorder.Ended() {
			for _, attr := range span.Attributes() {
				assert.NotContains(t, attr.Value.Emit(), "private-")
			}
			for _, event := range span.Events() {
				for _, attr := range event.Attributes {
					assert.NotContains(t, attr.Value.Emit(), "private-")
				}
			}
		}
	})
}

func TestTransportTracingRepeatedGetConn(t *testing.T) {
	for _, retry := range []bool{false, true} {
		name := "before-selection"
		if retry {
			name = "after-selection"
		}
		t.Run(name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				provider, recorder := transportProvider(t)
				_, span := startSpan(context.Background(), newTracer(provider), spanHTTP)
				state := &httpTransportTrace{span: span, start: time.Now()}
				hooks := state.hooks()
				hooks.GetConn("example.test:443")
				time.Sleep(3 * time.Millisecond)
				if !retry {
					// Concurrent HTTPS requests may join a shared HTTP/2
					// connection after already waiting in net/http's pool.
					hooks.GetConn("example.test:443")
				}
				hooks.GotConn(httptrace.GotConnInfo{Reused: true})
				time.Sleep(time.Millisecond)
				hooks.WroteRequest(httptrace.WroteRequestInfo{})
				if retry {
					// An unfinished reacquisition must still suppress phases,
					// even when there has only been one GotConn and one write.
					hooks.GetConn("example.test:443")
				}
				time.Sleep(5 * time.Millisecond)
				hooks.GotFirstResponseByte()
				state.finish()
				attrs := transportAttrs(recorded(t, recorder).get(t, spanHTTP))
				assert.Equal(t, int64(2), attrs[transportAttr+"connection_acquisitions"].AsInt64())
				assert.Equal(t, int64(1), attrs[transportAttr+"request_write_count"].AsInt64())
				for key, expected := range map[string]float64{
					"get_connection_ms": 3, "write_request_ms": 1, "wait_response_headers_ms": 5,
				} {
					if retry {
						assert.NotContains(t, attrs, transportAttr+key)
					} else {
						assert.Equal(t, expected, attrs[transportAttr+key].AsFloat64(), key)
					}
				}
			})
		})
	}
}

func TestTransportTracingRetriesAndLateCallbacks(t *testing.T) {
	provider, recorder := transportProvider(t)
	tracer := newTracer(provider)
	ctx, rpc := startSpan(context.Background(), tracer, spanRPC)
	ctx, state := startTransportTrace(ctx, true, rpc)
	var hooks *httptrace.ClientTrace
	client := &transportTracingClient{tracer: tracer, HTTPClient: transportHTTPFunc(func(req *http.Request) (*http.Response, error) {
		hooks = httptrace.ContextClientTrace(req.Context())
		// One HTTP Do may retry internally. Don't combine the first write
		// with the second connection acquisition into a made-up phase.
		for range 2 {
			hooks.GetConn("example.test:443")
			hooks.GotConn(httptrace.GotConnInfo{})
			hooks.WroteRequest(httptrace.WroteRequestInfo{})
		}
		hooks.GotFirstResponseByte()
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("ok"))}, nil
	})}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://example.test/query", nil)
	assert.NoError(t, err)
	res, err := client.Do(req)
	assert.NoError(t, err)
	_, err = io.Copy(io.Discard, res.Body)
	assert.NoError(t, err)
	// Exercise callbacks concurrently with completion, including late events.
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 20 {
				hooks.DNSStart(httptrace.DNSStartInfo{})
				hooks.DNSDone(httptrace.DNSDoneInfo{})
			}
		})
	}
	assert.NoError(t, res.Body.Close())
	wg.Wait()
	hooks.GotFirstResponseByte()
	state.finish(rpc)
	rpc.End()
	assert.Len(t, recorder.Ended(), 2)
	attrs := transportAttrs(recorded(t, recorder).get(t, spanHTTP))
	assert.Equal(t, int64(2), attrs[transportAttr+"connection_acquisitions"].AsInt64())
	assert.Equal(t, int64(2), attrs[transportAttr+"request_write_count"].AsInt64())
	for _, name := range []string{"get_connection_ms", "write_request_ms", "wait_response_headers_ms"} {
		assert.NotContains(t, attrs, transportAttr+name)
	}
}

type failingTransportBody struct{ err error }

func (b failingTransportBody) Read([]byte) (int, error) { return 0, b.err }
func (b failingTransportBody) Close() error             { return b.err }

func TestTransportTracingErrorsAndEarlyHeaders(t *testing.T) {
	for _, mode := range []string{"canceled", "read-error", "early-headers"} {
		t.Run(mode, func(t *testing.T) {
			provider, recorder := transportProvider(t)
			tracer := newTracer(provider)
			ctx, rpc := startSpan(context.Background(), tracer, spanRPC)
			ctx, state := startTransportTrace(ctx, true, rpc)
			client := &transportTracingClient{tracer: tracer, HTTPClient: transportHTTPFunc(func(req *http.Request) (*http.Response, error) {
				h := httptrace.ContextClientTrace(req.Context())
				if mode == "canceled" {
					return nil, context.Canceled
				}
				h.GetConn("example.test:443")
				h.GotConn(httptrace.GotConnInfo{})
				h.GotFirstResponseByte() // HTTP/2 may respond before writing finishes
				h.WroteRequest(httptrace.WroteRequestInfo{})
				var body io.ReadCloser = io.NopCloser(strings.NewReader("ok"))
				if mode == "read-error" {
					body = failingTransportBody{err: io.ErrUnexpectedEOF}
				}
				return &http.Response{StatusCode: 200, Body: body}, nil
			})}
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://example.test/query", nil)
			assert.NoError(t, err)
			res, err := client.Do(req)
			if mode == "canceled" {
				assert.ErrorIs(t, err, context.Canceled)
			} else {
				assert.NoError(t, err)
				_, readErr := io.ReadAll(res.Body)
				closeErr := res.Body.Close()
				if mode == "read-error" {
					assert.ErrorIs(t, readErr, io.ErrUnexpectedEOF)
					assert.ErrorIs(t, closeErr, io.ErrUnexpectedEOF)
				} else {
					assert.NoError(t, readErr)
					assert.NoError(t, closeErr)
				}
			}
			state.finish(rpc)
			rpc.End()
			span := recorded(t, recorder).get(t, spanHTTP)
			attrs := transportAttrs(span)
			assert.NotContains(t, attrs, transportAttr+"wait_response_headers_ms")
			if mode == "canceled" {
				assert.Equal(t, codes.Error, span.Status().Code)
			} else {
				assert.Equal(t, mode == "read-error", attrs[transportAttr+"response_body_failed"].AsBool())
			}
		})
	}
}

func TestTransportTracingDisabledDoesNotAllocateOrInstallHooks(t *testing.T) {
	provider, _ := transportProvider(t)
	ctx, span := provider.Tracer("test").Start(context.Background(), "caller")
	defer span.End()
	for _, enabled := range []bool{false, true} {
		selectedSpan := disabledSpan
		if !enabled {
			selectedSpan = span // recording provider alone isn't enough
		}
		allocs := testing.AllocsPerRun(100, func() {
			actualCtx, state := startTransportTrace(ctx, enabled, selectedSpan)
			if actualCtx != ctx || state != nil {
				panic("disabled transport tracing changed the context")
			}
			state.finish(selectedSpan)
		})
		assert.Zero(t, allocs)
	}
	client := &transportTracingClient{HTTPClient: transportHTTPFunc(func(req *http.Request) (*http.Response, error) {
		assert.Nil(t, httptrace.ContextClientTrace(req.Context()))
		return &http.Response{Body: http.NoBody}, nil
	})}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://example.test/query", nil)
	assert.NoError(t, err)
	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.NoBody, res.Body)
}

type transportQueryHandler struct {
	enginev1connect.UnimplementedQueryServiceHandler
	calls atomic.Int32
	retry bool
}

func (h *transportQueryHandler) OnlineQueryBulk(context.Context, *connect.Request[commonv1.OnlineQueryBulkRequest]) (*connect.Response[commonv1.OnlineQueryBulkResponse], error) {
	if h.calls.Add(1) == 1 && h.retry {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("refresh token"))
	}
	res := connect.NewResponse(&commonv1.OnlineQueryBulkResponse{})
	res.Header().Set("x-envoy-upstream-service-time", "17")
	res.Trailer().Set("x-test-trailer", "preserved")
	return res, nil
}

func TestTransportTracingHTTP2Integration(t *testing.T) {
	for _, mode := range []string{"h2c", "https", "insecure-https", "auth-retry", "default", "off", "no-provider", "dropped"} {
		t.Run(mode, func(t *testing.T) {
			provider, recorder := transportProvider(t)
			handler := &transportQueryHandler{retry: mode == "auth-retry"}
			mux := http.NewServeMux()
			path, queryHandler := enginev1connect.NewQueryServiceHandler(handler)
			mux.Handle(path, queryHandler)
			var server *httptest.Server
			if mode == "https" || mode == "insecure-https" {
				server = httptest.NewUnstartedServer(mux)
				server.EnableHTTP2 = true
				server.StartTLS()
				t.Cleanup(server.Close)
			} else {
				server = startUnencryptedHTTP2Server(t, mux)
			}
			cfg := &GRPCClientConfig{
				ApiServer: server.URL, QueryServer: server.URL, EnvironmentId: "test-env",
				SkipEnvironmentNameMapping: true, SkipEngineMapping: true,
				TracerProvider: provider, Tracing: &TracingOptions{Transport: mode != "off"},
				AuthProvider: rotatedAuthProvider,
				JWT:          &serverv1.GetTokenResponse{AccessToken: "mock-test-token", ExpiresIn: 3600},
			}
			if mode == "default" {
				cfg.Tracing = nil
			}
			if mode == "https" {
				cfg.HTTPClient = server.Client()
			}
			if mode == "insecure-https" {
				cfg.InsecureSkipVerify = true
			}
			if mode == "no-provider" {
				cfg.TracerProvider = nil
			}
			if mode == "dropped" {
				cfg.TracerProvider = sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.NeverSample()), sdktrace.WithSpanProcessor(recorder))
				t.Cleanup(func() {
					assert.NoError(t, cfg.TracerProvider.(*sdktrace.TracerProvider).Shutdown(context.Background()))
				})
			}
			client, err := NewGRPCClient(context.Background(), cfg)
			assert.NoError(t, err)
			returnedConfig := client.GetConfig()
			assert.Equal(t, cfg.Tracing, returnedConfig.Tracing)
			if cfg.Tracing != nil {
				// The running client owns a snapshot. Neither the input nor a
				// GetConfig result may change its instrumentation after setup.
				expectedTransport := cfg.Tracing.Transport
				cfg.Tracing.Transport = !expectedTransport
				returnedConfig.Tracing.Transport = !expectedTransport
				assert.Equal(t, expectedTransport, client.GetConfig().Tracing.Transport)
			}
			_, wrapped := client.GetConfig().HTTPClient.(*transportTracingClient)
			assert.False(t, wrapped, "GetConfig must not return the internal tracing wrapper")
			var gotConnections atomic.Int32
			ctx := httptrace.WithClientTrace(context.Background(), &httptrace.ClientTrace{
				GotConn: func(httptrace.GotConnInfo) { gotConnections.Add(1) },
			})
			for range 2 {
				res, err := client.OnlineQueryBulk(ctx, OnlineQueryParams{}.WithInput("user.id", []int64{1}).WithOutputs("user.id"))
				assert.NoError(t, err)
				assert.Equal(t, "17", res.ResponseHeaders.Get("x-envoy-upstream-service-time"))
			}
			assert.GreaterOrEqual(t, gotConnections.Load(), int32(2))
			var httpSpans []sdktrace.ReadOnlySpan
			for _, span := range recorder.Ended() {
				if span.Name() == spanHTTP {
					httpSpans = append(httpSpans, span)
				}
			}
			if mode == "default" || mode == "off" || mode == "no-provider" || mode == "dropped" {
				assert.Empty(t, httpSpans)
				return
			}
			expected := 2
			if handler.retry {
				expected++
			}
			assert.Len(t, httpSpans, expected)
			last := transportAttrs(httpSpans[len(httpSpans)-1])
			assert.True(t, last[transportAttr+"connection_reused"].AsBool())
			assert.True(t, last[transportAttr+"response_body_complete"].AsBool())
			assert.Equal(t, "2", last["network.protocol.version"].AsString())
			assert.Equal(t, float64(17), last[transportAttr+"envoy_upstream_service_time_ms"].AsFloat64())
			if handler.retry {
				assert.Equal(t, int64(1), transportAttrs(httpSpans[0])[transportAttr+"http_attempt"].AsInt64())
				assert.Equal(t, int64(2), transportAttrs(httpSpans[1])[transportAttr+"http_attempt"].AsInt64())
				assert.Equal(t, httpSpans[0].Parent().SpanID(), httpSpans[1].Parent().SpanID())
			}
		})
	}
}
