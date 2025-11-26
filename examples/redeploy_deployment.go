package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chalk-ai/chalk-go"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
)

func main() {
	// Initialize the Chalk client
	ctx := context.Background()
	client, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

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
	result, err := client.RedeployDeployment(ctx, req)
	if err != nil {
		log.Fatalf("Failed to redeploy deployment: %v", err)
	}

	// Access the response data
	fmt.Printf("Redeploy initiated successfully!\n")
	fmt.Printf("New Deployment ID: %s\n", result.RawResponse.DeploymentId)

	// The build_id field is deprecated but may still be returned
	if result.RawResponse.BuildId != "" {
		fmt.Printf("Build ID (deprecated): %s\n", result.RawResponse.BuildId)
	}
}