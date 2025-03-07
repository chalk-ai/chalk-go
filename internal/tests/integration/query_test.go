package integration

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
	"net/http"
	"testing"
	"time"
)

func testUserValues(t *testing.T, testUser *user) {
	t.Helper()
	assert.NotNil(t, testUser)
	assert.NotNil(t, testUser.Id)
	assert.Equal(t, int64(1), *testUser.Id)
	assert.NotNil(t, testUser.Gender)
	assert.Equal(t, "f", *testUser.Gender)
	assert.NotNil(t, testUser.Today)
	assert.NotNil(t, testUser.NiceNewFeature)
	assert.Equal(t, int64(9), *testUser.NiceNewFeature)
	assert.NotNil(t, testUser.SocureScore)
	assert.Equal(t, 123.0, *testUser.SocureScore)
	assert.NotNil(t, testUser.FavoriteNumbers)
	assert.Equal(t, []int64{1, 2, 3}, *testUser.FavoriteNumbers)
	assert.NotNil(t, testUser.FavoriteColors)
	assert.Equal(t, []string{"red", "green", "blue"}, *testUser.FavoriteColors)
	assert.NotNil(t, testUser.FranchiseSet)
}

// TestOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			params := chalk.OnlineQueryParams{}.
				WithOutputs(
					testFeatures.User.Id,
					testFeatures.User.Gender,
					testFeatures.User.Today,
					testFeatures.User.NiceNewFeature,
					testFeatures.User.SocureScore,
					testFeatures.User.FavoriteNumbers,
					testFeatures.User.FavoriteColors,
					testFeatures.User.FranchiseSet,
				)

			if fixture.useGrpc {
				client, err := chalk.NewGRPCClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a GRPC Chalk Client", err)
				}

				var users []user
				res, err := client.OnlineQueryBulk(
					context.Background(),
					params.WithInput(testFeatures.User.Id, []int{1}),
				)
				if err != nil {
					t.Fatal("Failed querying features", err)
				}
				assert.NoError(t, res.UnmarshalInto(&users))
				testUserValues(t, &users[0])
			} else {
				client, err := chalk.NewClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a Chalk Client", err)
				}

				var implicitUser user
				res, queryErr := client.OnlineQuery(
					context.Background(),
					params.WithInput(testFeatures.User.Id, 1),
					&implicitUser,
				)
				if queryErr != nil {
					t.Fatal("Failed querying features", queryErr)
				}

				var explicitUser user
				assert.NoError(t, res.UnmarshalInto(&explicitUser))
				testUserValues(t, &implicitUser)
				testUserValues(t, &explicitUser)
			}
		})
	}
}

// TestNamedQueriesE2E tests that querying with a query name works.
func TestNamedQueriesE2E(t *testing.T) {
	t.Skip("CHA-5086")
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			client, err := chalk.NewClient(context.Background())
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			params := chalk.OnlineQueryParams{}.
				WithInput("user.id", 1).
				WithQueryName("user_socure_score")

			_, queryErr := client.OnlineQuery(context.Background(), params, &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}
			assert.Equal(t, 123.0, *implicitUser.SocureScore)
		})
	}
}

// TestOnlineQueryBulkParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInFeatherHeader. Correctness
// of the results is *not* tested here.
func TestOnlineQueryBulkParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			err := chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			userIds := []int{1, 2}

			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now(), time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.User.Id, userIds).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			if fixture.useGrpc {
				client, err := chalk.NewGRPCClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a GRPC Chalk Client", err)
				}
				_, err = client.OnlineQueryBulk(context.Background(), req)
				assert.NoError(t, err)
			} else {
				client, err := chalk.NewClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a Chalk Client", err)
				}
				_, err = client.OnlineQueryBulk(context.Background(), req)
				assert.NoError(t, err)
			}
		})
	}
}

// TestOnlineQueryParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInOnlineQuery. Correctness
// of the results is *not* tested here.
func TestOnlineQueryParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			if fixture.useGrpc {
				client, err := chalk.NewGRPCClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a GRPC Chalk Client", err)
				}
				_, err = client.OnlineQueryBulk(context.Background(), req.WithInput("user.id", []int{1}))
				assert.NoError(t, err)
			} else {
				client, err := chalk.NewClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a Chalk Client", err)
				}
				_, err = client.OnlineQuery(
					context.Background(),
					req.WithInput("user.id", 1),
					nil,
				)
				assert.NoError(t, err)
			}
		})
	}
}

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
						WithInput(testFeatures.User.Id, []int{1}).
						WithOutputs(testFeatures.User.SocureScore),
				)
				assert.NoError(t, err)
			} else {
				client, err := chalk.NewClient(context.Background())
				if fixture.shouldFail {
					assert.Error(t, err)
					return
				} else {
					assert.NoError(t, err)
				}
				params := chalk.OnlineQueryParams{}.
					WithInput(testFeatures.User.Id, 1).
					WithOutputs(testFeatures.User.SocureScore)
				_, queryErr := client.OnlineQuery(context.Background(), params, nil)
				assert.NoError(t, queryErr)
			}
		})
	}
}
