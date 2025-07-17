package integration

import (
	"github.com/chalk-ai/chalk-go"
	"testing"
)

// Supplied an incorrect environment used to cause a panic
func TestWrongEnvironment(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	_, err := chalk.NewGRPCClient(
		t.Context(),
		&chalk.GRPCClientConfig{
			EnvironmentId: "wrong",
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk GRPC Client", err)
	}
	_, err = chalk.NewClient(
		t.Context(),
		&chalk.ClientConfig{
			EnvironmentId: "wrong",
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
}
