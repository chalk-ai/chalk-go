package integration

import (
	"context"
	"testing"

	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
)

var initFeaturesErr error

func init() {
	initFeaturesErr = chalk.InitFeatures(&testFeatures)
}

// TestOnlineQueryBulkGrpc mainly tests that a
// gRPC bulk query works e2e. Correctness is
// tested elsewhere.
func TestOnlineQueryBulkGrpc(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int64{1, 2}

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
	assert.Equal(t, userIds[0], *users[0].Id)
	assert.NotNil(t, users[0].SocureScore)
	assert.Equal(t, socureScore, *users[0].SocureScore)
	assert.NotNil(t, users[1].Id)
	assert.Equal(t, userIds[1], *users[1].Id)
	assert.NotNil(t, users[1].SocureScore)
	assert.Equal(t, socureScore, *users[1].SocureScore)
}

// TestOnlineQueryGrpcErringScalar tests requests with an erring scalar feature as the sole output
func TestOnlineQueryGrpcErringScalar(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient()
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(testFeatures.User.CrashingFeature)
	resp, err := client.OnlineQuery(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Errors)
	assert.NotNil(t, resp.GetData().GetResults()[0].GetValue().GetNullValue())
}

// TestOnlineQueryGrpcErringHasMany tests requests with an erring has-many feature as the sole output
func TestOnlineQueryGrpcErringHasMany(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient()
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Id, 1).
		WithOutputs(testFeatures.Series.CrashingInvestors)
	resp, err := client.OnlineQuery(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Errors)
	assert.Equal(t, 0, len(resp.GetData().GetResults()[0].GetValue().GetListValue().GetValues()))
}

// TestOnlineQueryGrpcSoleHasManyOutput tests requests with a has-many feature as the sole output
func TestOnlineQueryGrpcSoleHasManyOutput(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient()
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Id, "seed").
		WithOutputs(testFeatures.Series.Investors)
	resp, err := client.OnlineQuery(context.Background(), params)
	assert.NoError(t, err)
	assert.Nil(t, resp.Errors)
	mySeries := series{}
	if err := chalk.UnmarshalOnlineQueryResponse(resp, &mySeries); err != nil {
		assert.FailNow(t, "Failed to unmarshal response", err)
	}
	assert.NotNil(t, mySeries.Investors)
	assert.Equal(t, 50002, len(*mySeries.Investors))
}
