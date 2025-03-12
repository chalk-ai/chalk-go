package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	"testing"
)

// Supplied an incorrect environment used to cause a panic
func TestWrongEnvironment(t *testing.T) {
	SkipIfNotIntegrationTester(t)
	_, err := chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			UseGrpc:       true,
			EnvironmentId: "wrong",
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk GRPC Client", err)
	}
	_, err = chalk.NewClient(
		context.Background(),
		&chalk.ClientConfig{
			EnvironmentId: "wrong",
		},
	)
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
}
