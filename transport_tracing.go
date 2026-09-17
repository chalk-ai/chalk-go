package chalk

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"sync"
	"time"

	"connectrpc.com/connect"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const spanHTTP = "chalk.online_query_bulk.http"
const transportAttr = "chalk.transport."

type transportTraceKey struct{}

// A marker scoped to a recorded OnlineQueryBulk call. Authentication requests
// use the original HTTP client, so they cannot be mistaken for query attempts.
type transportTrace struct {
	mu       sync.Mutex
	start    time.Time
	first    time.Time
	attempts []*httpTransportTrace
}

func startTransportTrace(ctx context.Context, enabled bool, span trace.Span) (context.Context, *transportTrace) {
	if !enabled || !span.IsRecording() {
		return ctx, nil
	}
	t := &transportTrace{start: time.Now()}
	return context.WithValue(ctx, transportTraceKey{}, t), t
}

func (t *transportTrace) finish(span trace.Span) {
	if t == nil {
		return
	}
	t.mu.Lock()
	attempts, first := t.attempts, t.first
	t.mu.Unlock()
	span.SetAttributes(attribute.Int(transportAttr+"http_attempt_count", len(attempts)))
	if !first.IsZero() {
		span.SetAttributes(durationAttribute("pre_http_ms", first.Sub(t.start)))
	}
	// Normally Connect has already closed each body. Also finish on unusual
	// error paths, without closing or otherwise changing the response body.
	for _, attempt := range attempts {
		attempt.finish()
	}
}

type transportTracingClient struct {
	connect.HTTPClient
	tracer trace.Tracer
}

func (c *transportTracingClient) Do(req *http.Request) (*http.Response, error) {
	parent, _ := req.Context().Value(transportTraceKey{}).(*transportTrace)
	if parent == nil {
		return c.HTTPClient.Do(req)
	}
	t := &httpTransportTrace{start: time.Now()}
	ctx, span := startSpan(req.Context(), c.tracer, spanHTTP,
		attribute.String("server.address", req.URL.Hostname()),
	)
	t.span = span
	parent.mu.Lock()
	if parent.first.IsZero() {
		parent.first = t.start
	}
	parent.attempts = append(parent.attempts, t)
	attempt := len(parent.attempts)
	parent.mu.Unlock()
	span.SetAttributes(attribute.Int(transportAttr+"http_attempt", attempt))
	if req.ContentLength >= 0 {
		span.SetAttributes(attribute.Int64(transportAttr+"request_body_bytes", req.ContentLength))
	}
	if compression := req.Header.Get("grpc-encoding"); compression != "" {
		span.SetAttributes(attribute.String(transportAttr+"request_compression", compression))
	}
	// WithClientTrace composes with the caller's existing callbacks.
	req = req.WithContext(httptrace.WithClientTrace(ctx, t.hooks()))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		span.SetAttributes(attribute.Bool(transportAttr+"request_failed", true))
		span.SetStatus(codes.Error, "")
		// The RPC span records the error; avoid copying URLs or other error
		// strings into the additional transport information.
		t.finish()
		return res, err
	}
	span.SetAttributes(attribute.Int("http.response.status_code", res.StatusCode))
	if res.ProtoMajor > 0 {
		version := strconv.Itoa(res.ProtoMajor)
		if res.ProtoMajor == 1 {
			version += "." + strconv.Itoa(res.ProtoMinor)
		}
		span.SetAttributes(attribute.String("network.protocol.version", version))
	}
	if compression := res.Header.Get("grpc-encoding"); compression != "" {
		span.SetAttributes(attribute.String(transportAttr+"response_compression", compression))
	}
	if ms, err := strconv.ParseFloat(res.Header.Get("x-envoy-upstream-service-time"), 64); err == nil && ms >= 0 && ms < 1e12 {
		span.SetAttributes(attribute.Float64(transportAttr+"envoy_upstream_service_time_ms", ms))
	}
	if res.Body == nil {
		t.finish()
	} else {
		res.Body = &transportTracingBody{ReadCloser: res.Body, trace: t}
	}
	return res, nil
}

// httptrace callbacks may run concurrently, repeat on transport retries, or
// arrive after Do/Close returns. All callback state is guarded and late events
// are ignored. The caller's callbacks still run through httptrace composition.
type httpTransportTrace struct {
	mu                                        sync.Mutex
	span                                      trace.Span
	start                                     time.Time
	done                                      bool
	getConn, gotConn, wroteRequest, firstByte time.Time
	acquisitions, writes                      int
	connectionChanged                         bool
	conn                                      net.Conn
	reused                                    bool
	bytesRead                                 int64
	bodyEOF, bodyFailed                       bool
	phaseStarts                               map[string]time.Time
	phaseDurations                            map[string]time.Duration
	ambiguousPhases                           map[string]bool
}

func durationAttribute(name string, d time.Duration) attribute.KeyValue {
	return attribute.Float64(transportAttr+name, float64(d)/float64(time.Millisecond))
}

func (t *httpTransportTrace) eventLocked(name string, now time.Time, attrs ...attribute.KeyValue) {
	t.span.AddEvent("http."+name, trace.WithTimestamp(now), trace.WithAttributes(attrs...))
}

func (t *httpTransportTrace) phase(name, key string, begin bool, failed bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.done {
		return
	}
	now := time.Now()
	if begin {
		if t.phaseStarts == nil {
			t.phaseStarts = make(map[string]time.Time)
			t.phaseDurations = make(map[string]time.Duration)
		}
		if _, active := t.phaseStarts[key]; active {
			// DNSDone/TLSHandshakeDone have no operation identifier. When
			// same-key operations overlap, their durations cannot be paired.
			if t.ambiguousPhases == nil {
				t.ambiguousPhases = make(map[string]bool)
			}
			t.ambiguousPhases[name] = true
		}
		t.phaseStarts[key] = now
		t.eventLocked(name+"_start", now)
	} else {
		if start, ok := t.phaseStarts[key]; ok {
			t.phaseDurations[name] += now.Sub(start)
			delete(t.phaseStarts, key)
		}
		t.eventLocked(name+"_done", now, attribute.Bool("error", failed))
	}
}

func (t *httpTransportTrace) hooks() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn: func(string) {
			t.mu.Lock()
			defer t.mu.Unlock()
			if !t.done {
				t.acquisitions++
				now := time.Now()
				if t.getConn.IsZero() {
					t.getConn = now
				}
				if !t.gotConn.IsZero() {
					t.connectionChanged = true
				}
				t.eventLocked("get_conn", now)
			}
		},
		GotConn: func(info httptrace.GotConnInfo) {
			t.mu.Lock()
			defer t.mu.Unlock()
			if !t.done {
				if !t.gotConn.IsZero() {
					t.connectionChanged = true
				}
				t.gotConn = time.Now()
				t.conn, t.reused = info.Conn, info.Reused
				t.eventLocked("got_conn", t.gotConn, attribute.Bool(transportAttr+"connection_reused", info.Reused))
			}
		},
		DNSStart:     func(httptrace.DNSStartInfo) { t.phase("dns", "dns", true, false) },
		DNSDone:      func(info httptrace.DNSDoneInfo) { t.phase("dns", "dns", false, info.Err != nil) },
		ConnectStart: func(network, addr string) { t.phase("connect", network+":"+addr, true, false) },
		ConnectDone: func(network, addr string, err error) {
			t.phase("connect", network+":"+addr, false, err != nil)
		},
		TLSHandshakeStart: func() { t.phase("tls", "tls", true, false) },
		TLSHandshakeDone:  func(_ tls.ConnectionState, err error) { t.phase("tls", "tls", false, err != nil) },
		WroteHeaders: func() {
			t.mu.Lock()
			defer t.mu.Unlock()
			if !t.done {
				t.eventLocked("wrote_headers", time.Now())
			}
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			t.mu.Lock()
			defer t.mu.Unlock()
			if !t.done {
				t.writes++
				t.wroteRequest = time.Now()
				t.eventLocked("wrote_request", t.wroteRequest, attribute.Bool("error", info.Err != nil))
			}
		},
		GotFirstResponseByte: func() {
			t.mu.Lock()
			defer t.mu.Unlock()
			if !t.done {
				now := time.Now()
				if t.firstByte.IsZero() {
					t.firstByte = now
				}
				t.eventLocked("first_response_byte", now)
			}
		},
	}
}

func (t *httpTransportTrace) finish() {
	t.mu.Lock()
	if t.done {
		t.mu.Unlock()
		return
	}
	t.done = true
	now := time.Now()
	attrs := []attribute.KeyValue{
		attribute.Int(transportAttr+"connection_acquisitions", t.acquisitions),
		attribute.Int(transportAttr+"request_write_count", t.writes),
		attribute.Int64(transportAttr+"response_body_bytes", t.bytesRead),
		attribute.Bool(transportAttr+"response_body_complete", t.bodyEOF),
		attribute.Bool(transportAttr+"response_body_failed", t.bodyFailed),
	}
	addDuration := func(name string, start, end time.Time) {
		if !start.IsZero() && !end.IsZero() && !end.Before(start) {
			attrs = append(attrs, durationAttribute(name, end.Sub(start)))
		}
	}
	addDuration("response_headers_ms", t.start, t.firstByte)
	addDuration("response_read_and_decode_ms", t.firstByte, now)
	// net/http and its HTTP/2 pool can both call GetConn before a single
	// GotConn. Measure from the first GetConn so this includes the full wait
	// for a usable connection. Reacquisition after GotConn or repeated writes
	// indicate retries; keep their timeline without cross-attempt durations.
	if !t.connectionChanged && t.writes <= 1 {
		addDuration("get_connection_ms", t.getConn, t.gotConn)
		addDuration("write_request_ms", t.gotConn, t.wroteRequest)
		addDuration("wait_response_headers_ms", t.wroteRequest, t.firstByte)
	}
	for phase, duration := range t.phaseDurations {
		if !t.ambiguousPhases[phase] {
			attrs = append(attrs, durationAttribute(phase+"_ms", duration))
		}
	}
	conn := t.conn
	bodyFailed := t.bodyFailed
	if !t.gotConn.IsZero() {
		attrs = append(attrs, attribute.Bool(transportAttr+"connection_reused", t.reused))
	}
	t.mu.Unlock()
	if conn != nil {
		if addr := conn.RemoteAddr(); addr != nil {
			attrs = append(attrs, attribute.String(transportAttr+"peer_address", addr.String()))
		}
		attrs = append(attrs, tcpTraceAttributes(conn)...)
	}
	t.span.SetAttributes(attrs...)
	if bodyFailed {
		t.span.SetStatus(codes.Error, "")
	}
	t.span.End(trace.WithTimestamp(now))
}

type transportTracingBody struct {
	io.ReadCloser
	trace *httpTransportTrace
}

func (b *transportTracingBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	t := b.trace
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.done {
		t.bytesRead += int64(n)
		if err == io.EOF {
			if !t.bodyEOF {
				t.bodyEOF = true
				t.eventLocked("body_eof", time.Now())
			}
		} else if err != nil {
			t.bodyFailed = true
		}
	}
	return n, err
}

func (b *transportTracingBody) Close() error {
	err := b.ReadCloser.Close()
	if err != nil {
		b.trace.mu.Lock()
		b.trace.bodyFailed = true
		b.trace.mu.Unlock()
	}
	b.trace.finish()
	return err
}
