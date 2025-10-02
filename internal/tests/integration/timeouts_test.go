package integration

import (
	"context"
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

var timeouts = []struct {
	name       string
	timeout    time.Duration
	shouldFail bool
}{
	{name: "1 nanosecond", timeout: 1 * time.Nanosecond, shouldFail: true},
	{name: "5 seconds", timeout: 5 * time.Second},
	{name: "unspecified (zero value)", timeout: 0},
}

func TestTimeoutClientLevel(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				var err error
				if useGrpc {
					_, err = chalk.NewGRPCClient(t.Context(), &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
				} else {
					_, err = chalk.NewClient(t.Context(), &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
				}
				if timeoutFixture.shouldFail {
					assert.Error(t, err)
					return
				} else {
					assert.NoError(t, err)
				}

			})
		}
	}
}

func TestTimeoutClientOverrides(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				ctx, cancelFunc := context.WithTimeout(t.Context(), time.Minute*1)
				defer cancelFunc()
				var err error
				if useGrpc {
					_, err = chalk.NewGRPCClient(ctx, &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
				} else {
					_, err = chalk.NewClient(ctx, &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
				}
				// Since we've passed in a context with lenient timeout override,
				// all client instantiation should succeed.
				assert.NoError(t, err)

			})
		}
	}
}

func TestTimeoutRequestOverrides(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	testFeatures, initErr := GetTestFeatures()
	assert.NoError(t, initErr)
	lenientTimeout := time.Minute * 1
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.AllTypes.Id, []int{1}).
		WithOutputs(testFeatures.AllTypes.StrFeat)
	for _, useGrpc := range []bool{false, true} {
		for _, timeoutFixture := range timeouts {
			t.Run(fmt.Sprintf("grpc=%v, timeoutFixture=%v", useGrpc, timeoutFixture.name), func(t *testing.T) {
				t.Parallel()
				ctx, cancelFunc := context.WithTimeout(t.Context(), lenientTimeout)
				defer cancelFunc()
				if useGrpc {
					timeoutClient, err := chalk.NewGRPCClient(ctx, &chalk.GRPCClientConfig{Timeout: timeoutFixture.timeout})
					assert.NoError(t, err)

					// lenient override
					requestCtx, requestCancelFunc := context.WithTimeout(ctx, lenientTimeout)
					defer requestCancelFunc()
					res, err := timeoutClient.OnlineQueryBulk(requestCtx, params)
					assert.NoError(t, err)
					assert.Equal(t, 0, len(res.RawResponse.GetErrors()))

					// no override
					_, err = timeoutClient.OnlineQueryBulk(t.Context(), params)
					if timeoutFixture.shouldFail {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				} else {
					timeoutClient, err := chalk.NewClient(ctx, &chalk.ClientConfig{Timeout: timeoutFixture.timeout})
					assert.NoError(t, err)

					// lenient override
					requestCtx, requestCancelFunc := context.WithTimeout(ctx, lenientTimeout)
					defer requestCancelFunc()
					_, err = timeoutClient.OnlineQueryBulk(requestCtx, params)
					assert.NoError(t, err)

					// no override
					_, err = timeoutClient.OnlineQueryBulk(t.Context(), params)
					if timeoutFixture.shouldFail {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				}
			})
		}
	}
}
