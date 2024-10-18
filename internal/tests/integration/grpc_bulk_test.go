package integration

import (
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestOnlineQueryBulkGrpc mainly tests that a
// gRPC bulk query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulkGrpc(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int64{1, 2}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.Id, testFeatures.User.SocureScore)

	res, queryErr := client.OnlineQueryBulk(req)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var users []user
	assert.Nil(t, res.UnmarshalInto(&users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.NotNil(t, users[0].Id)
	assert.Equal(t, *users[0].Id, userIds[0])
	assert.NotNil(t, users[0].SocureScore)
	assert.Equal(t, *users[0].SocureScore, socureScore)
	assert.NotNil(t, users[1].Id)
	assert.Equal(t, *users[1].Id, userIds[1])
	assert.NotNil(t, users[1].SocureScore)
	assert.Equal(t, *users[1].SocureScore, socureScore)
}
