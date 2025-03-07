package benchmark

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"github.com/cockroachdb/errors"
	"net/http/httptest"
)

type TestFixture struct {
	Client chalk.GRPCClient
	Server *httptest.Server
}

func NewTestFixture(serverConfig *fixtures.MockServerConfig) (*TestFixture, error) {
	server, err := fixtures.NewMockServer(serverConfig)
	if err != nil {
		return nil, errors.Wrap(err, "creating mock server")
	}

	client, err := chalk.NewGRPCClient(context.Background(), &chalk.GRPCClientConfig{
		ApiServer:    server.URL,
		QueryServer:  server.URL,
		HTTPClient:   server.Client(),
		ClientId:     "bogus",
		ClientSecret: "bogus",
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating mock client")
	}

	return &TestFixture{
		Client: client,
		Server: server,
	}, nil
}

func (m *TestFixture) Close() {
	m.Server.Close()
}
