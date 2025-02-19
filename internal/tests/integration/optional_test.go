package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestQueryOptionalFeatures(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	userIds := []int{1}

	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	res := chalk.OnlineQueryParams{}.WithInput(testFeatures.User.Id, userIds[0]).WithOutputs(testFeatures.User.FullNameOptional, testFeatures.User.SocureScore)
	result := user{}
	_, err = client.OnlineQuery(context.Background(), res, &result)
	if err != nil {
		t.Fatal("Failed querying features", err)
	}
	assert.NotNil(t, result.SocureScore)
	assert.Equal(t, 123.0, *result.SocureScore)
	assert.Nil(t, result.FullNameOptional)
}
