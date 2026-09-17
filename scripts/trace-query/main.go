// Run from the repository root: go run ./scripts/trace-query
// Requires CHALK_CLIENT_ID and CHALK_CLIENT_SECRET. No tracing agent is needed.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"connectrpc.com/connect"
	chalk "github.com/chalk-ai/chalk-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const queryCount = 10

type settings struct {
	APIServer   string `json:"api_server"`
	Environment string `json:"environment"`
	QueryServer string `json:"query_server"`
	Concurrency int    `json:"concurrency"`
	Timeout     string `json:"timeout"`
	Transport   bool   `json:"transport_tracing"`
	Platform    string `json:"platform"`
}

type queryResult struct {
	Query             int                 `json:"query"`
	TraceID           string              `json:"trace_id"`
	CorrelationID     string              `json:"correlation_id"`
	SentTraceparent   string              `json:"sent_traceparent,omitempty"`
	ChalkQueryID      string              `json:"chalk_query_id,omitempty"`
	Headers           map[string][]string `json:"response_trace_headers,omitempty"`
	Trailers          map[string][]string `json:"response_trace_trailers,omitempty"`
	MetadataIDs       map[string]string   `json:"response_metadata_ids,omitempty"`
	ServerExecutionMS *float64            `json:"server_execution_ms,omitempty"`
	QueryMS           float64             `json:"query_ms"`
	DecodeMS          float64             `json:"decode_ms"`
	UserID            *int64              `json:"user_id,omitempty"`
	Error             string              `json:"error,omitempty"`
}

type resultKey struct{}

// The SDK emits spans but leaves propagation to the application's interceptor.
// Only retain tracing metadata, not arbitrary response headers or auth tokens.
func traceInterceptor(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		propagation.TraceContext{}.Inject(ctx, propagation.HeaderCarrier(req.Header()))
		result, _ := ctx.Value(resultKey{}).(*queryResult)
		if result != nil {
			result.SentTraceparent = req.Header().Get("traceparent")
		}
		res, err := next(ctx, req)
		if result != nil {
			if res != nil {
				result.Headers = traceHeaders(res.Header())
				result.Trailers = traceHeaders(res.Trailer())
			} else {
				var connectErr *connect.Error
				if errors.As(err, &connectErr) {
					result.Headers = traceHeaders(connectErr.Meta())
				}
			}
		}
		return res, err
	}
}

func isTraceKey(key string) bool {
	key = strings.NewReplacer("_", "", "-", "").Replace(strings.ToLower(key))
	switch key {
	case "traceparent", "traceid", "queryid", "correlationid", "requestid",
		"xtraceid", "xqueryid", "xcorrelationid", "xrequestid",
		"xchalktraceid", "xchalkqueryid", "xchalkcorrelationid", "xdatadogtraceid":
		return true
	}
	return false
}

func traceHeaders(headers http.Header) map[string][]string {
	ids := make(map[string][]string)
	for key, values := range headers {
		if isTraceKey(key) {
			ids[strings.ToLower(key)] = append([]string(nil), values...)
		}
	}
	return ids
}

type User struct {
	Id *int64
}

func query(ctx context.Context, client chalk.GRPCClient, tracer trace.Tracer, timeout time.Duration, number int) queryResult {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	ctx, root := tracer.Start(ctx, "trace-query", trace.WithAttributes(attribute.Int("query", number)))
	defer root.End()
	traceID := root.SpanContext().TraceID().String()
	result := queryResult{Query: number, TraceID: traceID, CorrelationID: traceID}
	ctx = context.WithValue(ctx, resultKey{}, &result)
	start := time.Now()
	res, err := client.OnlineQueryBulk(ctx, chalk.OnlineQueryParams{
		IncludeMeta: true, CorrelationId: traceID,
	}.WithInput("user.id", []int64{1}).WithOutputs("user.id"))
	result.QueryMS = milliseconds(time.Since(start))
	if res != nil {
		meta := res.RawResponse.GetResponseMeta()
		result.ChalkQueryID = meta.GetQueryId()
		if duration := meta.GetExecutionDuration(); duration != nil {
			ms := milliseconds(duration.AsDuration())
			result.ServerExecutionMS = &ms
		}
		result.MetadataIDs = make(map[string]string)
		for key, value := range meta.GetMetadata() {
			if isTraceKey(key) {
				result.MetadataIDs["metadata."+key] = value
			}
		}
		for key, value := range meta.GetAdditionalMetadata() {
			if isTraceKey(key) {
				result.MetadataIDs["additional_metadata."+key] = value.GetStringValue()
			}
		}
		if err == nil {
			var users []User
			start = time.Now()
			err = res.UnmarshalInto(&users)
			result.DecodeMS = milliseconds(time.Since(start))
			if err == nil {
				if len(users) != 1 || users[0].Id == nil || *users[0].Id != 1 {
					err = errors.New("expected one row with user.id = 1")
				} else {
					result.UserID = users[0].Id
				}
			}
		}
	}
	if err != nil {
		result.Error = err.Error()
		root.RecordError(err)
		root.SetStatus(codes.Error, err.Error())
	}
	return result
}

// Keep recording lightweight; convert spans and write JSON after the burst.
type recorder struct {
	mu    sync.Mutex
	spans []sdktrace.ReadOnlySpan
}

func (*recorder) OnStart(context.Context, sdktrace.ReadWriteSpan) {}
func (*recorder) Shutdown(context.Context) error                  { return nil }
func (*recorder) ForceFlush(context.Context) error                { return nil }
func (r *recorder) OnEnd(span sdktrace.ReadOnlySpan) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.spans = append(r.spans, span)
}

type spanEvent struct {
	Name       string         `json:"name"`
	Time       time.Time      `json:"time"`
	Attributes map[string]any `json:"attributes"`
}

type recordedSpan struct {
	Name        string         `json:"name"`
	Scope       string         `json:"scope"`
	TraceID     string         `json:"trace_id"`
	SpanID      string         `json:"span_id"`
	ParentID    string         `json:"parent_span_id"`
	Start       time.Time      `json:"start"`
	DurationMS  float64        `json:"duration_ms"`
	Status      string         `json:"status"`
	Description string         `json:"status_description,omitempty"`
	Attributes  map[string]any `json:"attributes"`
	Events      []spanEvent    `json:"events,omitempty"`
}

func attributes(attrs []attribute.KeyValue) map[string]any {
	values := make(map[string]any, len(attrs))
	for _, attr := range attrs {
		values[string(attr.Key)] = attr.Value.AsInterface()
	}
	return values
}

func (r *recorder) snapshot() []recordedSpan {
	r.mu.Lock()
	defer r.mu.Unlock()
	spans := make([]recordedSpan, 0, len(r.spans))
	for _, s := range r.spans {
		span := recordedSpan{
			Name: s.Name(), Scope: s.InstrumentationScope().Name,
			TraceID: s.SpanContext().TraceID().String(), SpanID: s.SpanContext().SpanID().String(),
			ParentID: s.Parent().SpanID().String(), Start: s.StartTime().UTC(),
			DurationMS: milliseconds(s.EndTime().Sub(s.StartTime())),
			Status:     s.Status().Code.String(), Description: s.Status().Description,
			Attributes: attributes(s.Attributes()),
		}
		for _, event := range s.Events() {
			span.Events = append(span.Events, spanEvent{event.Name, event.Time.UTC(), attributes(event.Attributes)})
		}
		spans = append(spans, span)
	}
	sort.Slice(spans, func(i, j int) bool { return spans[i].Start.Before(spans[j].Start) })
	return spans
}

func milliseconds(d time.Duration) float64 { return float64(d) / float64(time.Millisecond) }

func metric(attrs map[string]any, key string) string {
	value, ok := attrs["chalk.transport."+key]
	if !ok {
		return "-"
	}
	if number, ok := value.(float64); ok {
		return fmt.Sprintf("%.3f", number)
	}
	return fmt.Sprint(value)
}

func printResults(results []queryResult, spans []recordedSpan) {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "Q/HTTP\tquery_ms\tdecode_ms\trpc_ms\tpre_http_ms\tconn_ms\twrite_ms\twait_ms\tread_ms\trtt_ms\tenvoy_ms\tserver_ms\treused\tresult")
	for _, result := range results {
		var rpc recordedSpan
		var attempts []recordedSpan
		for _, span := range spans {
			if span.TraceID != result.TraceID {
				continue
			}
			switch span.Name {
			case "chalk.online_query_bulk.rpc":
				rpc = span
			case "chalk.online_query_bulk.http":
				attempts = append(attempts, span)
			}
		}
		if len(attempts) == 0 {
			attempts = append(attempts, recordedSpan{})
		}
		status := "OK"
		if result.Error != "" {
			status = "ERROR"
		}
		rpcMS := "-"
		if rpc.Name != "" {
			rpcMS = fmt.Sprintf("%.3f", rpc.DurationMS)
		}
		serverMS := "-"
		if result.ServerExecutionMS != nil {
			serverMS = fmt.Sprintf("%.3f", *result.ServerExecutionMS)
		}
		for _, attempt := range attempts {
			a := attempt.Attributes
			fmt.Fprintf(w, "%d/%s\t%.3f\t%.3f\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				result.Query, metric(a, "http_attempt"), result.QueryMS, result.DecodeMS, rpcMS,
				metric(rpc.Attributes, "pre_http_ms"), metric(a, "get_connection_ms"), metric(a, "write_request_ms"),
				metric(a, "wait_response_headers_ms"), metric(a, "response_read_and_decode_ms"),
				metric(a, "tcp_rtt_ms"), metric(a, "envoy_upstream_service_time_ms"), serverMS, metric(a, "connection_reused"), status)
		}
	}
	w.Flush()
	fmt.Println("\nIDs (trace ID is also the sent Chalk correlation ID):")
	for _, result := range results {
		queryID := result.ChalkQueryID
		if queryID == "" {
			queryID = "(not returned)"
		}
		fmt.Printf("%2d trace=%s chalk_query=%s\n", result.Query, result.TraceID, queryID)
		var returnedTraceID string
		if values := result.Trailers["x-chalk-trace-id"]; len(values) > 0 {
			returnedTraceID = values[0]
		} else if values := result.Headers["x-chalk-trace-id"]; len(values) > 0 {
			returnedTraceID = values[0]
		}
		if returnedTraceID != "" {
			fmt.Printf("   returned_trace=%s matches_sent=%t\n", returnedTraceID, returnedTraceID == result.TraceID)
		}
		if result.Error != "" {
			fmt.Printf("   error: %s\n", result.Error)
		}
	}
}

func run() error {
	var cfg settings
	flag.StringVar(&cfg.APIServer, "api-server", "https://api.meta.ci.chalk.dev", "Chalk API server")
	flag.StringVar(&cfg.Environment, "environment", "test45948e3f", "Chalk environment ID")
	flag.IntVar(&cfg.Concurrency, "concurrency", queryCount, "Concurrent queries (1 to 10); always sends 10 total")
	flag.BoolVar(&cfg.Transport, "transport", true, "Enable extra transport tracing")
	timeout := flag.Duration("timeout", 15*time.Second, "Timeout per query; initialization uses 30s")
	out := flag.String("out", filepath.Join(os.TempDir(), "chalk-trace-"+time.Now().Format("20060102-150405.000000000")+".json"), "JSON report path (must not already exist)")
	flag.Parse()
	if flag.NArg() != 0 || cfg.Concurrency < 1 || cfg.Concurrency > queryCount || *timeout <= 0 || cfg.APIServer == "" || cfg.Environment == "" {
		return errors.New("invalid arguments; run with -h for usage")
	}
	clientID, clientSecret := os.Getenv("CHALK_CLIENT_ID"), os.Getenv("CHALK_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return errors.New("set CHALK_CLIENT_ID and CHALK_CLIENT_SECRET before running")
	}
	// Check the output location before sending any queries and keep the report private.
	file, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return fmt.Errorf("create report: %w", err)
	}
	defer file.Close()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	recorder := &recorder{}
	provider := sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithSpanProcessor(recorder))
	defer provider.Shutdown(context.Background())
	initCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	client, err := chalk.NewGRPCClient(initCtx, &chalk.GRPCClientConfig{
		ClientId: clientID, ClientSecret: clientSecret,
		ApiServer: cfg.APIServer, EnvironmentId: cfg.Environment,
		TracerProvider: provider, Tracing: &chalk.TracingOptions{Transport: cfg.Transport},
		Interceptors: []connect.Interceptor{connect.UnaryInterceptorFunc(traceInterceptor)},
	})
	cancel()
	if err != nil {
		return fmt.Errorf("initialize Chalk client: %w", err)
	}
	cfg.QueryServer = client.GetConfig().QueryServer
	cfg.Timeout, cfg.Platform = timeout.String(), runtime.GOOS+"/"+runtime.GOARCH
	fmt.Printf("Query server: %s\n10 queries, concurrency=%d, no warm-up queries; authentication completed.\n\n", cfg.QueryServer, cfg.Concurrency)
	results := make([]queryResult, queryCount)
	start := time.Now().UTC()
	gate := make(chan struct{})
	var wg sync.WaitGroup
	for worker := range cfg.Concurrency {
		wg.Go(func() {
			<-gate
			for i := worker; i < queryCount; i += cfg.Concurrency {
				results[i] = query(ctx, client, provider.Tracer("chalk-go/trace-query"), *timeout, i+1)
			}
		})
	}
	close(gate)
	wg.Wait()
	spans := recorder.snapshot()
	report := struct {
		Started  time.Time      `json:"started"`
		Settings settings       `json:"settings"`
		Results  []queryResult  `json:"results"`
		Spans    []recordedSpan `json:"spans"`
	}{start, cfg, results, spans}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close report: %w", err)
	}
	printResults(results, spans)
	fmt.Printf("\nReport: %s\n", *out)
	fmt.Println("Missing values are '-'. wait_ms includes server work; read_ms includes protobuf decoding.")
	fmt.Println("envoy_ms is x-envoy-upstream-service-time, when returned, not the full gateway duration.")
	fmt.Println("server_ms is Chalk's reported execution duration; it excludes time outside that server interval.")
	if runtime.GOOS != "linux" {
		fmt.Println("TCP RTT is available only on Linux; it will be absent on this computer.")
	}
	fmt.Println("Match individual trace/query IDs to gateway timings; 10 queries are a smoke test, not a p95 estimate.")
	failed := 0
	for _, result := range results {
		if result.Error != "" {
			failed++
		}
	}
	if failed > 0 {
		return fmt.Errorf("%d of 10 queries failed (report saved)", failed)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
