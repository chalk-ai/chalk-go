package chalk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type integrationsRPCHandler struct {
	serverv1connect.UnimplementedIntegrationsServiceHandler
	mu sync.Mutex

	// existing is returned by GetIntegrationByName; nil means not found.
	existing *serverv1.Integration

	insertRequest *serverv1.InsertIntegrationRequest
	updateRequest *serverv1.UpdateIntegrationRequest
	headers       http.Header

	deletedIDs []string
	// deleteErr is returned by DeleteIntegration for this ID.
	deleteErrForID string
}

func (h *integrationsRPCHandler) GetIntegrationByName(
	_ context.Context,
	_ *connect.Request[serverv1.GetIntegrationByNameRequest],
) (*connect.Response[serverv1.GetIntegrationByNameResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	return connect.NewResponse(&serverv1.GetIntegrationByNameResponse{Integration: h.existing}), nil
}

func (h *integrationsRPCHandler) InsertIntegration(
	_ context.Context,
	req *connect.Request[serverv1.InsertIntegrationRequest],
) (*connect.Response[serverv1.InsertIntegrationResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.insertRequest = req.Msg
	h.headers = req.Header().Clone()
	return connect.NewResponse(&serverv1.InsertIntegrationResponse{
		Integration: &serverv1.Integration{
			Id:            "integration-1",
			Name:          ptr.OrNil(req.Msg.GetName()),
			Kind:          req.Msg.GetIntegrationKind(),
			EnvironmentId: "test-env",
			CreatedAt:     timestamppb.New(time.Unix(1700000000, 0)),
			UpdatedAt:     timestamppb.New(time.Unix(1700000000, 0)),
		},
	}), nil
}

func (h *integrationsRPCHandler) UpdateIntegration(
	_ context.Context,
	req *connect.Request[serverv1.UpdateIntegrationRequest],
) (*connect.Response[serverv1.UpdateIntegrationResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.updateRequest = req.Msg
	return connect.NewResponse(&serverv1.UpdateIntegrationResponse{
		Integration: &serverv1.Integration{
			Id:            req.Msg.GetIntegrationId(),
			Name:          ptr.OrNil(req.Msg.GetName()),
			Kind:          serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL,
			EnvironmentId: "test-env",
		},
	}), nil
}

func (h *integrationsRPCHandler) DeleteIntegration(
	_ context.Context,
	req *connect.Request[serverv1.DeleteIntegrationRequest],
) (*connect.Response[serverv1.DeleteIntegrationResponse], error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.deleteErrForID != "" && req.Msg.GetId() == h.deleteErrForID {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("not allowed"))
	}
	h.deletedIDs = append(h.deletedIDs, req.Msg.GetId())
	return connect.NewResponse(&serverv1.DeleteIntegrationResponse{}), nil
}

func newIntegrationsTestClient(t *testing.T, handler *integrationsRPCHandler) Client {
	t.Helper()

	mux := http.NewServeMux()
	authPath, authHandler := serverv1connect.NewAuthServiceHandler(&computeAuthHandler{})
	mux.Handle(authPath, authHandler)
	integrationsPath, integrationsHandler := serverv1connect.NewIntegrationsServiceHandler(handler)
	mux.Handle(integrationsPath, integrationsHandler)

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient(context.Background(), &ClientConfig{
		ClientId:      "client-id",
		ClientSecret:  "client-secret",
		ApiServer:     server.URL,
		EnvironmentId: "test-env",
	})
	require.NoError(t, err)
	return client
}

func TestClientCreateIntegrationSendsConfig(t *testing.T) {
	handler := &integrationsRPCHandler{}
	client := newIntegrationsTestClient(t, handler)

	integration, err := client.CreateIntegration(context.Background(), IntegrationParams{
		Name: "my_postgres",
		Kind: "postgresql",
		Variables: map[string]string{
			"PGHOST":     "localhost",
			"PGPASSWORD": "hunter2",
		},
	})
	require.NoError(t, err)
	require.Equal(t, "integration-1", integration.ID)
	require.Equal(t, "my_postgres", integration.Name)
	require.Equal(t, "postgresql", integration.Kind)
	require.Equal(t, "test-env", integration.EnvironmentID)
	require.Equal(t, time.Unix(1700000000, 0).UTC(), integration.CreatedAt)

	require.NotNil(t, handler.insertRequest)
	require.Equal(t, "my_postgres", handler.insertRequest.GetName())
	require.Equal(t, serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL, handler.insertRequest.GetIntegrationKind())
	require.Equal(t, "localhost", handler.insertRequest.GetConfig()["PGHOST"].GetLiteral())
	require.Equal(t, "hunter2", handler.insertRequest.GetConfig()["PGPASSWORD"].GetLiteral())

	require.Equal(t, "Bearer test-token", handler.headers.Get("Authorization"))
	require.Equal(t, "go-api", handler.headers.Get("X-Chalk-Server"))
	require.Equal(t, "test-env", handler.headers.Get("X-Chalk-Env-Id"))
}

func TestClientApplyIntegrationCreatesWhenMissing(t *testing.T) {
	handler := &integrationsRPCHandler{}
	client := newIntegrationsTestClient(t, handler)

	res, err := client.ApplyIntegration(context.Background(), IntegrationParams{
		Name:      "my_postgres",
		Kind:      "postgresql",
		Variables: map[string]string{"PGHOST": "localhost"},
	})
	require.NoError(t, err)
	require.True(t, res.Created)
	require.Equal(t, "integration-1", res.Integration.ID)
	require.NotNil(t, handler.insertRequest)
	require.Nil(t, handler.updateRequest)
}

func TestClientApplyIntegrationUpdatesWhenPresent(t *testing.T) {
	handler := &integrationsRPCHandler{
		existing: &serverv1.Integration{
			Id:   "integration-existing",
			Name: ptr.OrNil("my_postgres"),
			Kind: serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL,
		},
	}
	client := newIntegrationsTestClient(t, handler)

	res, err := client.ApplyIntegration(context.Background(), IntegrationParams{
		Name:      "my_postgres",
		Kind:      "postgresql",
		Variables: map[string]string{"PGHOST": "db.example.com"},
	})
	require.NoError(t, err)
	require.False(t, res.Created)
	require.Equal(t, "integration-existing", res.Integration.ID)
	require.Nil(t, handler.insertRequest)
	require.NotNil(t, handler.updateRequest)
	require.Equal(t, "integration-existing", handler.updateRequest.GetIntegrationId())
	require.Equal(t, "db.example.com", handler.updateRequest.GetConfig()["PGHOST"].GetLiteral())
}

func TestClientApplyIntegrationRejectsKindChange(t *testing.T) {
	handler := &integrationsRPCHandler{
		existing: &serverv1.Integration{
			Id:   "integration-existing",
			Name: ptr.OrNil("my_source"),
			Kind: serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL,
		},
	}
	client := newIntegrationsTestClient(t, handler)

	_, err := client.ApplyIntegration(context.Background(), IntegrationParams{
		Name: "my_source",
		Kind: "snowflake",
	})
	require.ErrorContains(t, err, "already exists with kind 'postgresql'")
	require.Nil(t, handler.insertRequest)
	require.Nil(t, handler.updateRequest)
}

func TestClientDeleteIntegrationByNameResolvesID(t *testing.T) {
	handler := &integrationsRPCHandler{
		existing: &serverv1.Integration{
			Id:   "integration-existing",
			Name: ptr.OrNil("my_postgres"),
			Kind: serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL,
		},
	}
	client := newIntegrationsTestClient(t, handler)

	require.NoError(t, client.DeleteIntegrationByName(context.Background(), "my_postgres"))
	require.Equal(t, []string{"integration-existing"}, handler.deletedIDs)
}

func TestClientDeleteIntegrationByNameErrorsWhenMissing(t *testing.T) {
	handler := &integrationsRPCHandler{}
	client := newIntegrationsTestClient(t, handler)

	err := client.DeleteIntegrationByName(context.Background(), "my_postgres")
	require.ErrorContains(t, err, "integration 'my_postgres' not found")
	require.Empty(t, handler.deletedIDs)
}

func TestClientDeleteIntegrationByNameRejectsEmptyName(t *testing.T) {
	handler := &integrationsRPCHandler{}
	client := newIntegrationsTestClient(t, handler)

	require.ErrorContains(t, client.DeleteIntegrationByName(context.Background(), ""), "integration name is required")
	require.Empty(t, handler.deletedIDs)
}

func TestClientDeleteIntegrationByNameWrapsRPCError(t *testing.T) {
	handler := &integrationsRPCHandler{
		existing: &serverv1.Integration{
			Id:   "integration-existing",
			Name: ptr.OrNil("my_postgres"),
			Kind: serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL,
		},
		deleteErrForID: "integration-existing",
	}
	client := newIntegrationsTestClient(t, handler)

	err := client.DeleteIntegrationByName(context.Background(), "my_postgres")
	require.ErrorContains(t, err, "deleting integration 'my_postgres'")
	require.Empty(t, handler.deletedIDs)
}

func TestIntegrationParamsResolveKind(t *testing.T) {
	kind, err := IntegrationParams{Name: "my_source", Kind: "postgresql"}.resolveKind()
	require.NoError(t, err)
	require.Equal(t, serverv1.IntegrationKind_INTEGRATION_KIND_POSTGRESQL, kind)

	kind, err = IntegrationParams{Name: "my_source", Kind: "INTEGRATION_KIND_SNOWFLAKE"}.resolveKind()
	require.NoError(t, err)
	require.Equal(t, serverv1.IntegrationKind_INTEGRATION_KIND_SNOWFLAKE, kind)

	_, err = IntegrationParams{Name: "my_source", Kind: "mongodb"}.resolveKind()
	require.ErrorContains(t, err, "unknown integration kind 'mongodb'")
	require.ErrorContains(t, err, "postgresql")
}

func TestIntegrationParamsValidation(t *testing.T) {
	client := newIntegrationsTestClient(t, &integrationsRPCHandler{})

	_, err := client.CreateIntegration(context.Background(), IntegrationParams{Kind: "postgresql"})
	require.ErrorContains(t, err, "name is required")

	_, err = client.CreateIntegration(context.Background(), IntegrationParams{Name: "my_postgres"})
	require.ErrorContains(t, err, "kind is required")
}
