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

//lint:ignore U1000 example
func redeployDeployment() {
	// Initialize the Chalk client
	ctx := context.Background()
	apiServer := "https://api.chalk.ai"
	httpClient := http.DefaultClient
	client := serverv1connect.NewBuilderServiceClient(
		httpClient,
		apiServer,
	)

	// Optional: pin the engine platform version for this redeploy.
	platformVersion := "v3.27.30"

	// Create a redeploy request
	req := &serverv1.RedeployDeploymentRequest{
		ExistingDeploymentId: "your-deployment-id", // Replace with actual deployment ID
		EnableProfiling:      false,
		DeploymentTags:       []string{"production", "v2"},
		// Pin the engine platform version (a tag/digest selector applied to the default
		// engine base image). Mutually exclusive with BaseImageOverride — setting both
		// returns an InvalidArgument error. When unset, the existing deployment's pinned
		// platform version is inherited.
		PlatformVersion: &platformVersion,
		// Optional: Set base image override (mutually exclusive with PlatformVersion)
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
	// To read the resolved platform version, call GetDeployment with result.Msg.DeploymentId;
	// the returned Deployment carries PinnedPlatformVersion.
}
