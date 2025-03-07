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

	client, err := chalk.NewGRPCClient(context.Background())
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int64{1, 2}

	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, userIds).
		WithOutputs(testFeatures.User.Id, testFeatures.User.SocureScore)

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
}

// TestOnlineQueryGrpcIncludeMeta mainly tests that the response
// includes the correct metadata when requested.
func TestOnlineQueryGrpcIncludeMeta(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	userId := int64(432)
	expectedSocureScore := 123.0

	restClient, err := chalk.NewClient(context.Background())
	assert.NoError(t, err)
	_, err = restClient.UploadFeatures(
		context.Background(),
		chalk.UploadFeaturesParams{
			Inputs: map[any]any{
				testFeatures.User.Id:          []int64{userId},
				testFeatures.User.SocureScore: []float64{expectedSocureScore},
			},
		},
	)
	assert.NoError(t, err)

	grpcClient, err := chalk.NewGRPCClient(context.Background())
	assert.NoError(t, err)
	req := chalk.OnlineQueryParams{IncludeMeta: true}.
		WithInput(testFeatures.User.Id, []int64{userId}).
		WithOutputs(testFeatures.User.Id, testFeatures.User.SocureScore, testFeatures.User.Today)
	res, err := grpcClient.OnlineQueryBulk(context.Background(), req)
	assert.NoError(t, err)

	row, err := res.GetRow(0)
	assert.NoError(t, err)

	socureScore, err := row.GetFeature("user.socure_score")
	assert.Nil(t, err)
	assert.NotNil(t, socureScore)
	assert.NotNil(t, socureScore.Meta)
	assert.Equal(t, expectedSocureScore, socureScore.Value)
	assert.Equal(t, true, socureScore.Meta.SourceType == chalk.SourceTypeOnlineStore)
	assert.Equal(t, userId, socureScore.Meta.Pkey)

	today, err := row.GetFeature("user.today")
	assert.Nil(t, err)
	assert.Equal(t, "neobank.resolvers.get_today", today.Meta.ResolverFqn)
	assert.Equal(t, userId, socureScore.Meta.Pkey)
}

// TestOnlineQueryGrpcErringScalar tests requests with an erring scalar feature as the sole output
func TestOnlineQueryGrpcErringScalar(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient(context.Background())
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, []int{1}).
		WithOutputs(testFeatures.User.CrashingFeature)
	resp, err := client.OnlineQueryBulk(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, resp.RawResponse.Errors)

	row, err := resp.GetRow(0)
	assert.NoError(t, err)

	crashingFeature, err := row.GetFeature("user.crashing_feature")
	assert.NoError(t, err)

	assert.NotNil(t, crashingFeature)
	assert.Nil(t, crashingFeature.Value)
}

// TestOnlineQueryGrpcErringHasMany tests requests with an erring has-many feature as the sole output
func TestOnlineQueryGrpcErringHasMany(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient(context.Background())
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Id, []int{1}).
		WithOutputs(testFeatures.Series.CrashingInvestors)
	resp, err := client.OnlineQueryBulk(context.Background(), params)
	assert.NoError(t, err)
	assert.Greater(t, len(resp.RawResponse.GetErrors()), 0)

	row, err := resp.GetRow(0)
	assert.NoError(t, err)

	crashingInvestors, err := row.GetFeature("series.crashing_investors")
	assert.NoError(t, err)
	assert.NotNil(t, crashingInvestors)
	castVal, ok := crashingInvestors.Value.([]any)
	assert.True(t, ok)
	assert.Equal(t, 0, len(castVal))
}

// TestOnlineQueryGrpcSoleHasManyOutput tests requests with a has-many feature as the sole output
func TestOnlineQueryGrpcSoleHasManyOutput(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	if initFeaturesErr != nil {
		t.Fatal("Failed initializing features", initFeaturesErr)
	}

	client, err := chalk.NewGRPCClient(context.Background())
	assert.NoError(t, err)
	params := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.Series.Id, []string{"seed"}).
		WithOutputs(testFeatures.Series.Investors)
	resp, err := client.OnlineQueryBulk(context.Background(), params)
	assert.NoError(t, err)
	assert.Nil(t, resp.RawResponse.GetErrors())

	row, err := resp.GetRow(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(row.Features))
	mySeries := []series{}
	if err := resp.UnmarshalInto(&mySeries); err != nil {
		assert.FailNow(t, "Failed to unmarshal response", err)
	}
	assert.NotNil(t, mySeries)
	assert.Equal(t, 1, len(mySeries))
	assert.NotNil(t, mySeries[0].Investors)
	assert.Equal(t, 50002, len(*mySeries[0].Investors))
	assert.NotNil(t, (*mySeries[0].Investors)[0].SeriesId)
	assert.Equal(t, "seed", *(*mySeries[0].Investors)[0].SeriesId)
}
