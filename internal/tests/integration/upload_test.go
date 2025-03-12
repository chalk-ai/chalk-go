package integration

import (
	"context"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/v2/internal/colls"
	"math/rand"
	"strings"
	"testing"
)

// TestUploadFeatures tests a basic features upload and
// also tests the two flavors of online query.
func TestUploadFeatures(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	// Implicitly sources config from env var
	client, err := chalk.NewClient(context.Background())
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	userIds := []int{111, 222, 333}
	socureScores := []float64{rand.Float64(), rand.Float64(), rand.Float64()}

	_, err = client.UploadFeatures(
		context.Background(),
		chalk.UploadFeaturesParams{
			Inputs: map[any]any{
				"user.id":           userIds,
				"user.socure_score": socureScores,
			},
		},
	)
	if err != nil {
		t.Fatal("Failed uploading features", err)
	}

	res, err := client.OnlineQuery(context.Background(), chalk.OnlineQueryParams{}.WithInput("user.id", userIds[0]).WithOutputs("user.socure_score"), nil)
	if err != nil {
		t.Fatal("Failed querying features", err)
	}
	ans, err := res.GetFeatureValue("user.socure_score")
	if err != nil {
		t.Fatal("Failed getting feature value for `user.socure_score`", err)
	}
	castAns, ok := ans.(float64)
	if !ok {
		t.Fatal("Failed casting feature value to float64")
	}
	if castAns != socureScores[0] {
		t.Fatalf("Queried feature 'user.socure_score' value '%v' for does not match uploaded value '%v'", castAns, socureScores[0])
	}

	bulkRes, err := client.OnlineQueryBulk(context.Background(), chalk.OnlineQueryParams{}.WithInput("user.id", userIds).WithOutputs("user.socure_score"))
	if err != nil {
		t.Fatal("Failed querying features", err)
	}
	reader := array.NewTableReader(bulkRes.ScalarsTable, 10_000)
	defer reader.Release()
	for reader.Next() {
		record := reader.Record()
		socureStrings := colls.Map(socureScores, func(val float64) string { return fmt.Sprintf("%v", val) })
		expectedString := "[" + strings.Join(socureStrings, " ") + "]"
		foundColumn := false
		for i, col := range record.Columns() {
			colName := record.ColumnName(i)
			if colName == "user.socure_score" {
				foundColumn = true
				if col.String() != expectedString {
					t.Fatalf("Queried feature 'user.socure_score' value '%v' for does not match uploaded value '%v'", col.String(), expectedString)
				}
			}
		}
		if !foundColumn {
			t.Fatal("Failed to find expected column 'user.socure_score'")
		}
	}
}
