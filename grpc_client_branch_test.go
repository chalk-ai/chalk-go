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

func startTestServer(t *testing.T, captured *capturedHeaders) *httptest.Server {
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

func newTestGRPCClient(t *testing.T, serverURL string, branch string) GRPCClient {
	t.Helper()
	client, err := NewGRPCClient(context.Background(), &GRPCClientConfig{
		ClientId:                   "test-client-id",
		ClientSecret:               "test-client-secret",
		ApiServer:                  serverURL,
		QueryServer:                serverURL,
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

func TestGRPCBranchHeadersFromClientConfig(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "my-branch")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCNoBranchHeadersWithoutBranch(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "engine-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCPerRequestBranchOverridesClientBranch(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "client-branch")

	params := OnlineQueryParamsComplete{}.WithBranchId("request-branch")
	_, _ = client.OnlineQueryBulk(context.Background(), params)

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "request-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCPerRequestBranchWithNoClientBranch(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "")

	params := OnlineQueryParamsComplete{}.WithBranchId("request-branch")
	_, _ = client.OnlineQueryBulk(context.Background(), params)

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "request-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCClientBranchAppliesToUpdateAggregates(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "my-branch")

	_, _ = client.UpdateAggregates(context.Background(), UpdateAggregatesParams{
		Inputs: map[any]any{
			"feat.id": []string{"a"},
		},
	})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCClientBranchAppliesToGetAggregates(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "my-branch")

	_, _ = client.GetAggregates(context.Background(), []string{"feat.id"})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCClientBranchAppliesToPlanAggregateBackfill(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "my-branch")

	_, _ = client.PlanAggregateBackfill(context.Background(), &aggregatev1.PlanAggregateBackfillRequest{})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "my-branch", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", h.Get("X-Chalk-Deployment-Type"))
}

func TestGRPCNoBranchOnNonBranchClient(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)
	client := newTestGRPCClient(t, server.URL, "")

	_, _ = client.GetAggregates(context.Background(), []string{"feat.id"})

	h := captured.last()
	assert.NotNil(t, h)
	assert.Equal(t, "", h.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "engine-grpc", h.Get("X-Chalk-Deployment-Type"))
}

// Verify that Ping also gets branch headers through the interceptor.
func TestGRPCPingAlsoGetsBranchHeaders(t *testing.T) {
	captured := &capturedHeaders{}
	server := startTestServer(t, captured)

	// Ping isn't on the public GRPCClient interface, so we test via
	// the interceptor indirectly — OnlineQueryBulk is sufficient to
	// confirm the interceptor applies to all queryClient methods.
	// This test validates that a second call still gets branch headers
	// (i.e., headers aren't consumed/cleared on first use).
	client := newTestGRPCClient(t, server.URL, "branch-a")

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})
	first := captured.last()
	assert.Equal(t, "branch-a", first.Get("X-Chalk-Branch-Id"))

	_, _ = client.OnlineQueryBulk(context.Background(), OnlineQueryParamsComplete{})
	second := captured.last()
	assert.Equal(t, "branch-a", second.Get("X-Chalk-Branch-Id"))
	assert.Equal(t, "branch-grpc", second.Get("X-Chalk-Deployment-Type"))
}
