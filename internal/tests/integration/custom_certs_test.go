package integration

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"net/http"
	"testing"
)

func TestCustomCerts(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	systemCertPool, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal("Failed creating a system cert pool", err)
	}
	emptyCertPool := x509.NewCertPool()

	for _, fixture := range []struct {
		useGrpc    bool
		certPool   *x509.CertPool
		shouldFail bool
	}{
		{useGrpc: false, certPool: systemCertPool, shouldFail: false},
		{useGrpc: false, certPool: emptyCertPool, shouldFail: true},
		{useGrpc: true, certPool: systemCertPool, shouldFail: false},
		{useGrpc: true, certPool: emptyCertPool, shouldFail: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v, shouldFail=%v", fixture.useGrpc, fixture.shouldFail), func(t *testing.T) {
			t.Parallel()
			httpClient := http.Client{
				Transport: &http2.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: fixture.certPool,
					},
				},
			}

			if fixture.useGrpc {
				client, err := chalk.NewGRPCClient(context.Background(), &chalk.GRPCClientConfig{
					HTTPClient: &httpClient,
				})
				if fixture.shouldFail {
					assert.Error(t, err)
					return
				} else {
					assert.NoError(t, err)
				}
				_, err = client.OnlineQueryBulk(
					context.Background(),
					chalk.OnlineQueryParams{}.
						WithInput(testFeatures.AllTypes.Id, []int{1}).
						WithOutputs(testFeatures.AllTypes.StrFeat),
				)
				assert.NoError(t, err)
			} else {
				client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{HTTPClient: &httpClient})
				if fixture.shouldFail {
					assert.Error(t, err)
					return
				} else {
					assert.NoError(t, err)
				}
				params := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.AllTypes.Id, 1).
					WithOutputs(testFeatures.AllTypes.StrFeat)
				_, err = client.OnlineQuery(context.Background(), params, nil)
				assert.NoError(t, err)
			}
		})
	}
}
