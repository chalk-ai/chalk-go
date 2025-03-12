package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go/v2"
	assert "github.com/stretchr/testify/require"
	"testing"
)

// TestOnlineQueryBulk mainly tests that a
// real query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulk(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient(context.Background())
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
		WithOutputs(testFeatures.User.SocureScore, testFeatures.User.Today)

	res, queryErr := client.OnlineQueryBulk(context.Background(), req)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var users []user
	assert.Nil(t, res.UnmarshalInto(&users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.Equal(t, userIds[0], *users[0].Id)
	assert.Equal(t, socureScore, *users[0].SocureScore)
	assert.Equal(t, userIds[1], *users[1].Id)
	assert.Equal(t, socureScore, *users[1].SocureScore)
	assert.NotNil(t, users[0].Today)
}

// TestOnlineQueryBulkGrpcNative mainly tests that a
// real query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulkGrpcNative(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewGRPCClient(context.Background())
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
		WithOutputs(testFeatures.User.Id, testFeatures.User.SocureScore, testFeatures.User.Today)

	res, queryErr := client.OnlineQueryBulk(context.Background(), req)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var users []user
	assert.Nil(t, res.UnmarshalInto(&users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.NotNil(t, users[0].Id)
	assert.Equal(t, userIds[0], *users[0].Id)
	assert.NotNil(t, users[0].SocureScore)
	assert.Equal(t, socureScore, *users[0].SocureScore)
	assert.NotNil(t, users[1].Id)
	assert.Equal(t, userIds[1], *users[1].Id)
	assert.NotNil(t, users[1].SocureScore)
	assert.Equal(t, socureScore, *users[1].SocureScore)
	assert.NotNil(t, users[0].Today)
}
