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

func getParams() chalk.OnlineQueryParamsComplete {
	return chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
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
}

func testUserValues(t *testing.T, testUser *user) {
	t.Helper()
	assert.Equal(t, *testUser.Id, int64(1))
	assert.Equal(t, *testUser.Gender, "f")
	assert.NotNil(t, testUser.Today)
	assert.Equal(t, *testUser.NiceNewFeature, int64(9))
	assert.Equal(t, *testUser.SocureScore, 123.0)
	assert.Equal(t, *testUser.FavoriteNumbers, []int64{1, 2, 3})
	assert.Equal(t, *testUser.FavoriteColors, []string{"red", "green", "blue"})
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
			certPool, err := x509.SystemCertPool()
			if err != nil {
				t.Fatal("Failed creating a system cert pool", err)
			}
			httpClient := http.Client{
				Transport: &http2.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: certPool,
					},
				},
			}

			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc, HTTPClient: &httpClient})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			res, queryErr := client.OnlineQuery(getParams(), &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}

			var explicitUser user
			err = res.UnmarshalInto(&explicitUser)
			// TODO: We need to fix this nil check
			//       to just be `!= nil`
			if err != (*chalk.ClientError)(nil) {
				t.Fatal("Failed unmarshaling result", err)
			}

			testUserValues(t, &implicitUser)
			testUserValues(t, &explicitUser)
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
			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
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

			_, queryErr := client.OnlineQuery(params, &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}
			assert.Equal(t, 123.0, *implicitUser.SocureScore)
		})
	}
}

// TestGRPCOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
//
// This test is also notably different from the E2E test
// where a gRPC client is also tested but is built on top
// of the existing REST `Client` interface.
func TestGRPCOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewGRPCClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}

	res, queryErr := client.OnlineQuery(context.Background(), getParams())
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var testUser user
	assert.NoError(t, chalk.UnmarshalOnlineQueryResponse(res, &testUser))
	testUserValues(t, &testUser)
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

			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
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

			_, err = client.OnlineQueryBulk(req)
			assert.NoError(t, err)
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
			if fixture.useGrpc {
				t.Skip("CHA-4780")
			}
			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

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
				WithInput(testFeatures.User.Id, 1).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			_, err = client.OnlineQuery(req, nil)
			assert.NoError(t, err)
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

			client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: fixture.useGrpc, HTTPClient: &httpClient})
			if fixture.shouldFail {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			var userObj user
			_, queryErr := client.OnlineQuery(getParams(), &userObj)
			assert.NoError(t, queryErr)
		})
	}
}
