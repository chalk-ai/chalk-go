package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chalk-ai/chalk-go"
)

// User represents a feature namespace for user features
// In practice, this would be generated using `chalk codegen`
type User struct {
	Id           *string
	Email        *string
	Name         *string
	CreditScore  *int64
	AccountValue *float64
}

func onlineQueryBulk() {
	// Initialize the Chalk gRPC client
	// By default, this will use environment variables or ~/.chalk.yml configuration:
	// - CHALK_CLIENT_ID or config file
	// - CHALK_CLIENT_SECRET or config file
	// - CHALK_API_SERVER or config file
	// - CHALK_ACTIVE_ENVIRONMENT or config file
	ctx := context.Background()
	chalkClient, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

	// Alternatively, you can explicitly configure the client:
	// chalkClient, err := chalk.NewGRPCClient(
	// 	ctx,
	// 	&chalk.GRPCClientConfig{
	// 		ClientId:      "id-89140a6614886982a6782106759e30",
	// 		ClientSecret:  "sec-b1ba98e658d7ada4ff4c7464fb0fcee65fe2cbd86b3dd34141e16f6314267b7b",
	// 		ApiServer:     "https://api.chalk.ai",
	// 		EnvironmentId: "prod",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to create Chalk client: %v", err)
	// }

	// Execute a bulk online query for multiple users
	// In this example, we query features for 3 users at once
	res, err := chalkClient.OnlineQueryBulk(
		ctx,
		chalk.OnlineQueryParams{
			// Optional: Add query metadata
			QueryName:     "user_credit_check",
			CorrelationId: "bulk-query-example-123",
			IncludeMeta:   true,
		}.
			// Specify input features with slices for bulk queries
			WithInput("user.id", []string{"user_1", "user_2", "user_3"}).
			// Specify which features to compute and return
			WithOutputs(
				"user.id",
				"user.email",
				"user.name",
				"user.credit_score",
				"user.account_value",
			),
	)

	// Note: If the query returns errors, both `res` and `err` may be non-nil
	// The result may contain partial data even when errors occur
	if err != nil {
		log.Printf("Query completed with errors: %v", err)
		// You can still access partial results if available
	}

	// Method 1: Unmarshal results into a slice of structs
	var users []User
	if err := res.UnmarshalInto(&users); err != nil {
		log.Fatalf("Failed to unmarshal results: %v", err)
	}

	fmt.Printf("Successfully queried %d users:\n", len(users))
	for i, user := range users {
		fmt.Printf("\nUser %d:\n", i+1)
		if user.Id != nil {
			fmt.Printf("  ID: %s\n", *user.Id)
		}
		if user.Email != nil {
			fmt.Printf("  Email: %s\n", *user.Email)
		}
		if user.Name != nil {
			fmt.Printf("  Name: %s\n", *user.Name)
		}
		if user.CreditScore != nil {
			fmt.Printf("  Credit Score: %d\n", *user.CreditScore)
		}
		if user.AccountValue != nil {
			fmt.Printf("  Account Value: $%.2f\n", *user.AccountValue)
		}
	}

	// Method 2: Access individual row results
	fmt.Println("\n--- Accessing individual rows ---")
	row, err := res.GetRow(0)
	if err != nil {
		log.Fatalf("Failed to get row: %v", err)
	}

	// Access feature values from a row
	emailFeature, err := row.GetFeature("user.email")
	if err != nil {
		log.Fatalf("Failed to get email feature: %v", err)
	}
	fmt.Printf("First user's email: %v\n", emailFeature.Value)

	// Method 3: Access query metadata
	if res.GetQueryMeta() != nil {
		meta := res.GetQueryMeta()
		fmt.Printf("\n--- Query Metadata ---\n")
		fmt.Printf("Query ID: %s\n", meta.QueryId)
		fmt.Printf("Deployment ID: %s\n", meta.DeploymentId)
		fmt.Printf("Environment ID: %s\n", meta.EnvironmentId)
		if meta.QueryTimestamp != nil {
			fmt.Printf("Query Timestamp: %v\n", *meta.QueryTimestamp)
		}
	}

	// Method 4: Check for query errors
	if errors, err := res.GetErrors(); err == nil && len(errors) > 0 {
		fmt.Printf("\n--- Query Errors ---\n")
		for i, queryErr := range errors {
			fmt.Printf("Error %d: %s\n", i+1, queryErr.Message)
			fmt.Printf("  Code: %s\n", queryErr.Code)
			if queryErr.Feature != "" {
				fmt.Printf("  Feature: %s\n", queryErr.Feature)
			}
		}
	}

	// Method 5: Access the Arrow table directly (for advanced use cases)
	table, err := res.GetTable()
	if err != nil {
		log.Fatalf("Failed to get table: %v", err)
	}
	defer table.Release()
	fmt.Printf("\n--- Arrow Table ---\n")
	fmt.Printf("Number of rows: %d\n", table.NumRows())
	fmt.Printf("Number of columns: %d\n", table.NumCols())
}
