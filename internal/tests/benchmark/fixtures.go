package benchmark

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"github.com/cockroachdb/errors"
	"net/http/httptest"
)

type TestFixture struct {
	Client chalk.Client
	Server *httptest.Server
}

func NewTestFixture(serverConfig *fixtures.MockServerConfig) (*TestFixture, error) {
	server, err := fixtures.NewMockServer(serverConfig)
	if err != nil {
		return nil, errors.Wrap(err, "creating mock server")
	}

	client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{
		ApiServer:    server.URL,
		QueryServer:  server.URL,
		HTTPClient:   server.Client(),
		ClientId:     "bogus",
		ClientSecret: "bogus",
		UseGrpc:      true,
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
