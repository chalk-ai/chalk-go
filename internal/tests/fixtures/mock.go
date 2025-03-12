package fixtures

import (
	"connectrpc.com/connect"
	"context"
	aggregatev1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/common/v1"
	enginev1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/engine/v1"
	"github.com/chalk-ai/chalk-go/v2/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/v2/gen/chalk/server/v1/serverv1connect"
	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"net/http/httptest"
	"time"
)

type mockQueryServer struct {
	queryBulkResponse *connect.Response[commonv1.OnlineQueryBulkResponse]
}

func (m *mockQueryServer) OnlineQueryBulk(
	ctx context.Context,
	req *connect.Request[commonv1.OnlineQueryBulkRequest],
) (*connect.Response[commonv1.OnlineQueryBulkResponse], error) {
	return m.queryBulkResponse, nil
}

func (m *mockQueryServer) Ping(
	ctx context.Context,
	req *connect.Request[enginev1.PingRequest],
) (*connect.Response[enginev1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) OnlineQuery(
	ctx context.Context,
	req *connect.Request[commonv1.OnlineQueryRequest],
) (*connect.Response[commonv1.OnlineQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) OnlineQueryMulti(
	ctx context.Context,
	req *connect.Request[commonv1.OnlineQueryMultiRequest],
) (*connect.Response[commonv1.OnlineQueryMultiResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) UploadFeaturesBulk(
	ctx context.Context,
	req *connect.Request[commonv1.UploadFeaturesBulkRequest],
) (*connect.Response[commonv1.UploadFeaturesBulkResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) UploadFeatures(
	ctx context.Context,
	req *connect.Request[commonv1.UploadFeaturesRequest],
) (*connect.Response[commonv1.UploadFeaturesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) PlanAggregateBackfill(
	ctx context.Context,
	req *connect.Request[aggregatev1.PlanAggregateBackfillRequest],
) (*connect.Response[aggregatev1.PlanAggregateBackfillResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockQueryServer) GetAggregates(
	ctx context.Context,
	req *connect.Request[aggregatev1.GetAggregatesRequest],
) (*connect.Response[aggregatev1.GetAggregatesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

type mockAuthServer struct{}

func (m *mockAuthServer) GetToken(
	ctx context.Context,
	req *connect.Request[serverv1.GetTokenRequest],
) (*connect.Response[serverv1.GetTokenResponse], error) {
	return connect.NewResponse(&serverv1.GetTokenResponse{
		AccessToken: "abc",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
		ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour * 24 * 30)),
		ApiServer:   "http://localhost:8080",
	}), nil
}

func (m *mockAuthServer) CreateLinkSession(
	ctx context.Context,
	req *connect.Request[serverv1.CreateLinkSessionRequest],
) (*connect.Response[serverv1.CreateLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockAuthServer) GetLinkSession(
	ctx context.Context,
	req *connect.Request[serverv1.GetLinkSessionRequest],
) (*connect.Response[serverv1.GetLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

func (m *mockAuthServer) UpdateLinkSession(
	ctx context.Context,
	req *connect.Request[serverv1.UpdateLinkSessionRequest],
) (*connect.Response[serverv1.UpdateLinkSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not implemented"))
}

type MockServerConfig struct {
	QueryBulkResponse *commonv1.OnlineQueryBulkResponse
}

func NewMockServer(config *MockServerConfig) (*httptest.Server, error) {
	queryServer := &mockQueryServer{
		queryBulkResponse: connect.NewResponse(config.QueryBulkResponse),
	}
	authServer := &mockAuthServer{}

	mux := http.NewServeMux()
	mux.Handle(enginev1connect.NewQueryServiceHandler(queryServer))
	mux.Handle(serverv1connect.NewAuthServiceHandler(authServer))
	server := httptest.NewUnstartedServer(mux)
	http2Server := &http2.Server{
		// Otherwise we get `executing online query: unavailable: http2: frame too large`
		MaxReadFrameSize: 10 * 1024 * 1024,
	}
	if err := http2.ConfigureServer(server.Config, http2Server); err != nil {
		return nil, err
	}
	server.EnableHTTP2 = true
	server.StartTLS()
	return server, nil
}
