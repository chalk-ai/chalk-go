package integration

import (
	"github.com/chalk-ai/chalk-go"
	"testing"
)

// TestOnlineQueryBulkGrpc mainly tests that a
// gRPC bulk query works e2e. Correctness is
// tested elsewhere.
func TestWrongEnvironment(t *testing.T) {
	_, err := chalk.NewClient(
		&chalk.ClientConfig{
			UseGrpc:       true,
			EnvironmentId: "wrong",
		})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
}
