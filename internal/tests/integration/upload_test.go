package integration

import (
	"github.com/chalk-ai/chalk-go"
	"os"
	"testing"
)

func SkipIfNotIntegrationTester(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTER") != "true" {
		t.Skip()
	}
}

// TestUploadFeatures tests a basic features upload
// with all primitive data types
func TestUploadFeatures(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewClient() // Implicitly sources config from env var
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	_, err = client.UploadFeatures(chalk.UploadFeaturesParams{
		Inputs: map[any]any{
			"user.id":           []string{"1", "2", "3"},
			"user.socure_score": []float64{625.0, 636.0, 5525.0},
		},
	})
	if err != nil {
		t.Fatal("Failed uploading features", err)
	}
}
