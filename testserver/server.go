package testserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"google.golang.org/protobuf/proto"
)

// MockServer wraps httptest.Server with a configuration API for setting
// up mock RPC responses.
type MockServer struct {
	*httptest.Server
	registry *ResponseRegistry
}

// NewMockBuilderServer creates a new in-process HTTP server that implements
// the BuilderService RPC interface. The server URL can be used to create
// real RPC clients that will interact with the mock server.
//
// Example:
//
//	server := testserver.NewMockBuilderServer(t)
//	defer server.Close()
//
//	server.OnGetClusterTimescaleDB().Return(&serverv1.GetClusterTimescaleDBResponse{
//	    Id: "test-id",
//	})
//
//	client := serverv1connect.NewBuilderServiceClient(http.DefaultClient, server.URL)
//	resp, err := client.GetClusterTimescaleDB(ctx, connect.NewRequest(...))
func NewMockBuilderServer(t testing.TB) *MockServer {
	registry := NewResponseRegistry()
	handler := newBuilderServiceHandler(registry)

	// Create Connect RPC handler
	path, rpcHandler := serverv1connect.NewBuilderServiceHandler(handler)

	// Create HTTP mux and register handler
	mux := http.NewServeMux()
	mux.Handle(path, rpcHandler)

	// Create httptest server
	httpServer := httptest.NewServer(mux)

	return &MockServer{
		Server:   httpServer,
		registry: registry,
	}
}

// MethodConfigBuilder provides a fluent API for configuring mock responses
// for a specific RPC method.
type MethodConfigBuilder[T proto.Message] struct {
	methodName string
	registry   *ResponseRegistry
}

// Return configures the method to return the given response.
func (b *MethodConfigBuilder[T]) Return(response T) {
	b.registry.SetResponse(b.methodName, response)
}

// ReturnError configures the method to return the given error.
func (b *MethodConfigBuilder[T]) ReturnError(err error) {
	b.registry.SetError(b.methodName, err)
}

// WithBehavior configures the method to use a custom behavior function
// that can implement complex logic.
func (b *MethodConfigBuilder[T]) WithBehavior(fn BehaviorFunc) {
	b.registry.SetBehavior(b.methodName, fn)
}

// OnGetClusterTimescaleDB configures the GetClusterTimescaleDB RPC method.
func (s *MockServer) OnGetClusterTimescaleDB() *MethodConfigBuilder[*serverv1.GetClusterTimescaleDBResponse] {
	return &MethodConfigBuilder[*serverv1.GetClusterTimescaleDBResponse]{
		methodName: "GetClusterTimescaleDB",
		registry:   s.registry,
	}
}

// OnUpdateClusterTimescaleDB configures the UpdateClusterTimescaleDB RPC method.
func (s *MockServer) OnUpdateClusterTimescaleDB() *MethodConfigBuilder[*serverv1.UpdateClusterTimescaleDBResponse] {
	return &MethodConfigBuilder[*serverv1.UpdateClusterTimescaleDBResponse]{
		methodName: "UpdateClusterTimescaleDB",
		registry:   s.registry,
	}
}

// OnCreateClusterTimescaleDB configures the CreateClusterTimescaleDB RPC method.
func (s *MockServer) OnCreateClusterTimescaleDB() *MethodConfigBuilder[*serverv1.CreateClusterTimescaleDBResponse] {
	return &MethodConfigBuilder[*serverv1.CreateClusterTimescaleDBResponse]{
		methodName: "CreateClusterTimescaleDB",
		registry:   s.registry,
	}
}

// GetCapturedRequests returns all requests captured for the given method name.
// This is useful for test assertions.
func (s *MockServer) GetCapturedRequests(methodName string) []proto.Message {
	return s.registry.GetCapturedRequests(methodName)
}

// Reset clears all configured responses, errors, behaviors, and captured requests.
func (s *MockServer) Reset() {
	s.registry.Reset()
}
