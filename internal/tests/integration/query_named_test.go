package integration

import (
	"context"
	"fmt"
	chalk "github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestNamedQueries tests that querying with a query name works.
func TestNamedQueries(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	testCases := []struct {
		name         string
		queryName    string
		version      string
		expectedName string
	}{
		{
			name:         "simple_named_query",
			queryName:    "nqtest_get_name",
			version:      "",
			expectedName: "Portrait Of the Artist as a Young Man",
		},
		{
			name:         "versioned_tagged_query",
			queryName:    "nqtest_get_tags",
			version:      "2.0.0",
			expectedName: "hardcoded name tag2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for _, useGrpc := range []bool{false, true} {
				t.Run(fmt.Sprintf("grpc=%v", useGrpc), func(t *testing.T) {
					t.Parallel()

					if useGrpc {
						params := chalk.OnlineQueryParams{}.
							WithInput(testFeatures.NQFeatures.Id, []int64{1}).
							WithQueryName(tc.queryName).
							WithQueryNameVersion(tc.version)

						res, err := grpcClient.OnlineQueryBulk(context.Background(), params)
						assert.NoError(t, err)

						var results []nqFeatures
						assert.NoError(t, res.UnmarshalInto(&results))
						assert.NotEmpty(t, results)
						assert.NotNil(t, results[0].Name)

						assert.Equal(t, tc.expectedName, *results[0].Name)
					} else {
						params := chalk.OnlineQueryParams{}.
							WithInput(testFeatures.NQFeatures.Id, int64(1)).
							WithQueryName(tc.queryName).
							WithQueryNameVersion(tc.version)

						var result nqFeatures
						_, err := restClient.OnlineQuery(context.Background(), params, &result)
						assert.NoError(t, err)
						assert.NotNil(t, result.Name)

						assert.Equal(t, tc.expectedName, *result.Name)
					}
				})
			}
		})
	}
}
