package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chalk-ai/chalk-go"
)

//lint:ignore U1000 example
func insecureClient() {
	ctx := context.Background()

	// Create a Chalk gRPC client with InsecureSkipVerify enabled.
	// WARNING: This disables SSL certificate verification and should only
	// be used in testing or development environments. Do not use in production.
	client, err := chalk.NewGRPCClient(
		ctx,
		&chalk.GRPCClientConfig{
			ClientId:           "your-client-id",
			ClientSecret:       "your-client-secret",
			ApiServer:          "https://api.chalk.ai",
			EnvironmentId:      "dev",
			InsecureSkipVerify: true, // Ignore expired or self-signed SSL certificates
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

	fmt.Println("Successfully created Chalk client with InsecureSkipVerify enabled")

	// Use the client as normal
	// For example, to get a token:
	token, err := client.GetToken(ctx)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	fmt.Printf("Token valid until: %v\n", token.ValidUntil)
}
