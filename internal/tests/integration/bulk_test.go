package integration

import (
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

// TestOnlineQueryBulkResultUnmarshal tests that we can unmarshal
// the result of an online query bulk request into structs.
func TestOnlineQueryBulkResultUnmarshal(t *testing.T) {
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

	users := make([]user, 0)
	assert.Nil(t, res.UnmarshalInto(&users))
	assert.Equal(t, 2, len(users))

	socureScore := 123.0
	assert.True(t, reflect.DeepEqual(users[0], user{Id: &userIds[0], SocureScore: &socureScore}))
	assert.True(t, reflect.DeepEqual(users[1], user{Id: &userIds[1], SocureScore: &socureScore}))
}
