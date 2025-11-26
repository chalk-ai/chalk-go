package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
)

func redeployDeployment() {
	// Initialize the Chalk client
	ctx := context.Background()
	apiServer := "https://api.chalk.ai"
	httpClient := http.DefaultClient
	client := serverv1connect.NewBuilderServiceClient(
		httpClient,
		apiServer,
	)

	// Create a redeploy request
	req := &serverv1.RedeployDeploymentRequest{
		ExistingDeploymentId: "your-deployment-id", // Replace with actual deployment ID
		EnableProfiling:      false,
		DeploymentTags:       []string{"production", "v2"},
		// Optional: Set base image override
		// BaseImageOverride: &someBaseImage,
		// Optional: Override the graph
		// OverrideGraph: &graphv1.Graph{...},
	}

	// Execute the redeploy
	result, err := client.RedeployDeployment(ctx, connect.NewRequest(req))
	if err != nil {
		log.Fatalf("Failed to redeploy deployment: %v", err)
	}

	// Access the response data
	fmt.Printf("Redeploy initiated successfully!\n")
	fmt.Printf("New Deployment ID: %s\n", result.Msg.DeploymentId)

	// The build_id field is deprecated but may still be returned
	if result.Msg.BuildId != "" {
		fmt.Printf("Build ID (deprecated): %s\n", result.Msg.BuildId)
	}
}
