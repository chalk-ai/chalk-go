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
// the BuilderService and AuthService RPC interfaces. The server URL can be used
// to create real RPC clients that will interact with the mock server.
//
// The auth service returns a default mock token automatically, so tests don't need
// to configure authentication unless testing auth-specific scenarios.
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

	// Create handlers
	builderHandler := newBuilderServiceHandler(registry)
	authHandler := newAuthServiceHandler(registry)
	teamHandler := newTeamServiceHandler(registry)

	integrationsHandler := newIntegrationsServiceHandler(registry)

	// Create HTTP mux
	mux := http.NewServeMux()

	// Register BuilderService handler
	builderPath, builderRPCHandler := serverv1connect.NewBuilderServiceHandler(builderHandler)
	mux.Handle(builderPath, builderRPCHandler)

	// Register AuthService handler
	authPath, authRPCHandler := serverv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(authPath, authRPCHandler)

	// Register TeamService handler
	teamPath, teamRPCHandler := serverv1connect.NewTeamServiceHandler(teamHandler)
	mux.Handle(teamPath, teamRPCHandler)

	// Register IntegrationsService handler
	integrationsPath, integrationsRPCHandler := serverv1connect.NewIntegrationsServiceHandler(integrationsHandler)
	mux.Handle(integrationsPath, integrationsRPCHandler)

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

// OnDeleteClusterTimescaleDB configures the DeleteClusterTimescaleDB RPC method.
func (s *MockServer) OnDeleteClusterTimescaleDB() *MethodConfigBuilder[*serverv1.DeleteClusterTimescaleDBResponse] {
	return &MethodConfigBuilder[*serverv1.DeleteClusterTimescaleDBResponse]{
		methodName: "DeleteClusterTimescaleDB",
		registry:   s.registry,
	}
}

// OnGetTelemetryDeployment configures the GetTelemetryDeployment RPC method.
func (s *MockServer) OnGetTelemetryDeployment() *MethodConfigBuilder[*serverv1.GetTelemetryDeploymentResponse] {
	return &MethodConfigBuilder[*serverv1.GetTelemetryDeploymentResponse]{
		methodName: "GetTelemetryDeployment",
		registry:   s.registry,
	}
}

// OnCreateTelemetryDeployment configures the CreateTelemetryDeployment RPC method.
func (s *MockServer) OnCreateTelemetryDeployment() *MethodConfigBuilder[*serverv1.CreateTelemetryDeploymentResponse] {
	return &MethodConfigBuilder[*serverv1.CreateTelemetryDeploymentResponse]{
		methodName: "CreateTelemetryDeployment",
		registry:   s.registry,
	}
}

// OnUpdateTelemetryDeployment configures the UpdateTelemetryDeployment RPC method.
func (s *MockServer) OnUpdateTelemetryDeployment() *MethodConfigBuilder[*serverv1.UpdateTelemetryDeploymentResponse] {
	return &MethodConfigBuilder[*serverv1.UpdateTelemetryDeploymentResponse]{
		methodName: "UpdateTelemetryDeployment",
		registry:   s.registry,
	}
}

// OnDeleteTelemetryDeployment configures the DeleteTelemetryDeployment RPC method.
func (s *MockServer) OnDeleteTelemetryDeployment() *MethodConfigBuilder[*serverv1.DeleteTelemetryDeploymentResponse] {
	return &MethodConfigBuilder[*serverv1.DeleteTelemetryDeploymentResponse]{
		methodName: "DeleteTelemetryDeployment",
		registry:   s.registry,
	}
}

// OnGetEnv configures the GetEnv RPC method.
// By default, the mock server returns a valid environment with test-cluster-id.
func (s *MockServer) OnGetEnv() *MethodConfigBuilder[*serverv1.GetEnvResponse] {
	return &MethodConfigBuilder[*serverv1.GetEnvResponse]{
		methodName: "GetEnv",
		registry:   s.registry,
	}
}

// OnGetToken configures the GetToken RPC method.
// By default, the mock server returns a valid token automatically.
// Use this method only if you need to test auth failures or custom token responses.
func (s *MockServer) OnGetToken() *MethodConfigBuilder[*serverv1.GetTokenResponse] {
	return &MethodConfigBuilder[*serverv1.GetTokenResponse]{
		methodName: "GetToken",
		registry:   s.registry,
	}
}

// OnInsertIntegration configures the InsertIntegration RPC method.
func (s *MockServer) OnInsertIntegration() *MethodConfigBuilder[*serverv1.InsertIntegrationResponse] {
	return &MethodConfigBuilder[*serverv1.InsertIntegrationResponse]{
		methodName: "InsertIntegration",
		registry:   s.registry,
	}
}

// OnGetIntegration configures the GetIntegration RPC method.
func (s *MockServer) OnGetIntegration() *MethodConfigBuilder[*serverv1.GetIntegrationResponse] {
	return &MethodConfigBuilder[*serverv1.GetIntegrationResponse]{
		methodName: "GetIntegration",
		registry:   s.registry,
	}
}

// OnUpdateIntegration configures the UpdateIntegration RPC method.
func (s *MockServer) OnUpdateIntegration() *MethodConfigBuilder[*serverv1.UpdateIntegrationResponse] {
	return &MethodConfigBuilder[*serverv1.UpdateIntegrationResponse]{
		methodName: "UpdateIntegration",
		registry:   s.registry,
	}
}

// OnDeleteIntegration configures the DeleteIntegration RPC method.
func (s *MockServer) OnDeleteIntegration() *MethodConfigBuilder[*serverv1.DeleteIntegrationResponse] {
	return &MethodConfigBuilder[*serverv1.DeleteIntegrationResponse]{
		methodName: "DeleteIntegration",
		registry:   s.registry,
	}
}

// OnListIntegrations configures the ListIntegrations RPC method.
func (s *MockServer) OnListIntegrations() *MethodConfigBuilder[*serverv1.ListIntegrationsResponse] {
	return &MethodConfigBuilder[*serverv1.ListIntegrationsResponse]{
		methodName: "ListIntegrations",
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
