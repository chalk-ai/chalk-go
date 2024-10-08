package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestOnlineQueryBulk mainly tests that a
// real query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulk(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient()
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
		WithOutputs(testFeatures.User.SocureScore)

	res, queryErr := client.OnlineQueryBulk(req)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var users []user
	assert.Nil(t, res.UnmarshalInto(&users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.Equal(t, *users[0].Id, userIds[0])
	assert.Equal(t, *users[0].SocureScore, socureScore)
	assert.Equal(t, *users[1].Id, userIds[1])
	assert.Equal(t, *users[1].SocureScore, socureScore)
}

// TestOnlineQueryBulkGrpcNative mainly tests that a
// real query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulkGrpcNative(t *testing.T) {
	t.Skip("CHA-4780")
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewGRPCClient()
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
		WithOutputs(testFeatures.User.SocureScore)

	res, queryErr := client.OnlineQueryBulk(context.Background(), req)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var users []user
	assert.Nil(t, chalk.UnmarshalOnlineQueryBulkResponse(res, &users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.Equal(t, *users[0].Id, userIds[0])
	assert.Equal(t, *users[0].SocureScore, socureScore)
	assert.Equal(t, *users[1].Id, userIds[1])
	assert.Equal(t, *users[1].SocureScore, socureScore)
}
