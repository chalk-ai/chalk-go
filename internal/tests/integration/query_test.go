package integration

import (
	"github.com/chalk-ai/chalk-go"
	"testing"
)

// TestOnlineQueryAllTypesUnmarshalling only tests
// unmarshalling real data from the staging server
// does not crash. Correctness is tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryAllTypesUnmarshalling(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}
	req := chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(
			testFeatures.User.Id,
			testFeatures.User.Gender,
			testFeatures.User.Today,
			testFeatures.User.SocureScore,
			testFeatures.User.FavoriteNumbers,
			testFeatures.User.FavoriteColors,
			testFeatures.User.FranchiseSet,
		)

	res, queryErr := client.OnlineQuery(req, nil)
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}
	var testUser user
	err = res.UnmarshalInto(&testUser)
	// TODO: We need to fix this nil check
	//       to just be `!= nil`
	if err != (*chalk.ClientError)(nil) {
		t.Fatal("Failed unmarshaling result", err)
	}
}
