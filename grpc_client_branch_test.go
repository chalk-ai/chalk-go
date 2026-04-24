package chalk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"connectrpc.com/connect"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	assert "github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type capturedHeaders struct {
	mu      sync.Mutex
	headers []http.Header
}

func (c *capturedHeaders) capture(h http.Header) {
	c.mu.Lock()
	defer c.mu.Unlock()
	clone := h.Clone()
	c.headers = append(c.headers, clone)
}

func (c *capturedHeaders) last() http.Header {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.headers) == 0 {
		return nil
	}
	return c.headers[len(c.headers)-1]
}

func (c *capturedHeaders) count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.headers)
}

type headerCapturingQueryHandler struct {
	enginev1connect.UnimplementedQueryServiceHandler
	captured *capturedHeaders
}

func (h *headerCapturingQueryHandler) OnlineQueryBulk(
	_ context.Context,
	req *connect.Request[commonv1.OnlineQueryBulkRequest],
) (*connect.Response[commonv1.OnlineQueryBulkResponse], error) {
	h.captured.capture(req.Header())
	return connect.NewResponse(&commonv1.OnlineQueryBulkResponse{}), nil
}

func (h *headerCapturingQueryHandler) UploadFeaturesBulk(
	_ context.Context,
	req *connect.Request[commonv1.UploadFeaturesBulkRequest],
) (*connect.Response[commonv1.UploadFeaturesBulkResponse], error) {
	h.captured.capture(req.Header())
	return connect.NewResponse(&commonv1.UploadFeaturesBulkResponse{}), nil
}

func (h *headerCapturingQueryHandler) GetAggregates(
	_ context.Context,
	req *connect.Request[aggregatev1.GetAggregatesRequest],
) (*connect.Response[aggregatev1.GetAggregatesResponse], error) {
	h.captured.capture(req.Header())
	return connect.NewResponse(&aggregatev1.GetAggregatesResponse{}), nil
}

func (h *headerCapturingQueryHandler) PlanAggregateBackfill(
	_ context.Context,
	req *connect.Request[aggregatev1.PlanAggregateBackfillRequest],
) (*connect.Response[aggregatev1.PlanAggregateBackfillResponse], error) {
	h.captured.capture(req.Header())
	return connect.NewResponse(&aggregatev1.PlanAggregateBackfillResponse{}), nil
}

type minimalAuthHandler struct {
	serverv1connect.UnimplementedAuthServiceHandler
}

func (h *minimalAuthHandler) GetToken(
	_ context.Context,
	_ *connect.Request[serverv1.GetTokenRequest],
) (*connect.Response[serverv1.GetTokenResponse], error) {
	return connect.NewResponse(&serverv1.GetTokenResponse{
		AccessToken: "mock-test-token",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}), nil
}

type minimalGraphHandler struct {
	serverv1connect.UnimplementedGraphServiceHandler
}

func startQueryServer(t *testing.T, captured *capturedHeaders) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	queryPath, queryHandler := enginev1connect.NewQueryServiceHandler(&headerCapturingQueryHandler{captured: captured})
	mux.Handle(queryPath, queryHandler)
	h2cHandler := h2c.NewHandler(mux, &http2.Server{})
	server := httptest.NewServer(h2cHandler)
	t.Cleanup(server.Close)
	return server
}

func startAPIServer(t *testing.T, captured *capturedHeaders) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	queryPath, queryHandler := enginev1connect.NewQueryServiceHandler(&headerCapturingQueryHandler{captured: captured})
	mux.Handle(queryPath, queryHandler)
	authPath, authHandler := serverv1connect.NewAuthServiceHandler(&minimalAuthHandler{})
	mux.Handle(authPath, authHandler)
	graphPath, graphHandler := serverv1connect.NewGraphServiceHandler(&minimalGraphHandler{})
	mux.Handle(graphPath, graphHandler)
	h2cHandler := h2c.NewHandler(mux, &http2.Server{})
	server := httptest.NewServer(h2cHandler)
	t.Cleanup(server.Close)
	return server
}

func newTestGRPCClient(t *testing.T, apiServerURL string, queryServerURL string, branch string) GRPCClient {
	t.Helper()
	client, err := NewGRPCClient(context.Background(), &GRPCClientConfig{
		ClientId:                   "test-client-id",
		ClientSecret:               "test-client-secret",
		ApiServer:                  apiServerURL,
		QueryServer:                queryServerURL,
		EnvironmentId:              "test-env",
		Branch:                     branch,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken: "mock-test-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		},
	})
	assert.NoError(t, err)
	return client
}

func TestGRPCBranchQueryRoutesToAPIServer(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "my-branch")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
	// Branch queries route through api.chalk.ai's Envoy, which dispatches
	// based on x-chalk-server. It must be "go-api" — sending "engine" makes
	// Envoy try to forward to the engine cluster and the upstream resets,
	// surfacing as "reset reason: protocol error".
	assert.Equal(t, "go-api", h.Get("X-Chalk-Server"))
}

func TestGRPCNoBranchQueryRoutesToEngine(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})

	assert.Equal(t, 1, engineCaptured.count())
	assert.Equal(t, 0, apiCaptured.count())
	h := engineCaptured.last()
	assert.Equal(t, "", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "engine-grpc", h.Get("X-Chalk-Deployment-Type"))
	// Non-branch queries go directly to the engine.
	assert.Equal(t, "engine", h.Get("X-Chalk-Server"))
}

func TestGRPCPerRequestBranchRoutesToAPIServer(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "")

	params := OnlineQueryParamsComplete{}.WithBranchId("request-branch")
	_, _ = client.OnlineQueryBulk(context.Background(), params)

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "request-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCPerRequestBranchOverridesClientBranch(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "client-branch")

	params := OnlineQueryParamsComplete{}.WithBranchId("request-branch")
	_, _ = client.OnlineQueryBulk(context.Background(), params)

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "request-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCBranchUpdateAggregatesRoutesToAPIServer(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "my-branch")

	_, _ = client.UpdateAggregates(context.Background(), UpdateAggregatesParams{
		Inputs: map[any]any{
			"feat.id": []string{"a"},
		},
	})

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCBranchGetAggregatesRoutesToAPIServer(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "my-branch")

	_, _ = client.GetAggregates(context.Background(), []string{"feat.id"})

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCBranchPlanAggregateBackfillRoutesToAPIServer(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "my-branch")

	_, _ = client.PlanAggregateBackfill(context.Background(), &aggregatev1.PlanAggregateBackfillRequest{})

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 1, apiCaptured.count())
	h := apiCaptured.last()
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCNoBranchAggregatesRouteToEngine(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "")

	_, _ = client.GetAggregates(context.Background(), []string{"feat.id"})

	assert.Equal(t, 1, engineCaptured.count())
	assert.Equal(t, 0, apiCaptured.count())
	h := engineCaptured.last()
	assert.Equal(t, "", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "engine-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCBranchHeadersPersistAcrossCalls(t *testing.T) {
	engineCaptured := &capturedHeaders{}
	apiCaptured := &capturedHeaders{}
	engineServer := startQueryServer(t, engineCaptured)
	apiServer := startAPIServer(t, apiCaptured)
	client := newTestGRPCClient(t, apiServer.URL, engineServer.URL, "branch-a")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})
	first := apiCaptured.last()
	assert.Equal(t, "branch-a", first.Get("X-Chalk-Branch-Id"))

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})
	second := apiCaptured.last()
	assert.Equal(t, "branch-a", second.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", second.Get("X-Chalk-Deployment-Type"))

	assert.Equal(t, 0, engineCaptured.count())
	assert.Equal(t, 2, apiCaptured.count())
}
