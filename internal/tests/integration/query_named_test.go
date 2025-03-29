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
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			var resultUser user
			if fixture.useGrpc {
				client, err := chalk.NewGRPCClient(context.Background())
				assert.NoError(t, err)
				res, err := client.OnlineQueryBulk(
					context.Background(),
					chalk.OnlineQueryParams{}.
						WithInput("user.id", []int{1}).
						WithQueryName("user_socure_score").
						WithQueryNameVersion("1.0.0"),
				)
				assert.NoError(t, err)
				results := []user{}
				assert.NoError(t, res.UnmarshalInto(&results))
				resultUser = results[0]
			} else {
				client, err := chalk.NewClient(context.Background())
				if err != nil {
					t.Fatal("Failed creating a Chalk Client", err)
				}
				params := chalk.OnlineQueryParams{}.
					WithInput("user.id", 1).
					WithQueryName("user_socure_score").
					WithQueryNameVersion("1.0.0")

				_, queryErr := client.OnlineQuery(context.Background(), params, &resultUser)
				if queryErr != nil {
					t.Fatal("Failed querying features", queryErr)
				}
			}
			assert.Equal(t, 123.0, *resultUser.SocureScore)

		})
	}
}
