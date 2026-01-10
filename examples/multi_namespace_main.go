package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chalk-ai/chalk-go"
)

// This example demonstrates querying Chalk with multiple input namespaces.
// It shows how to provide inputs from different namespaces (user.id and author.id)
// and retrieve features from both namespaces in a single query.
//
// Expected output:
//
//	=== Multi-Namespace Query Results ===
//
//	User Features:
//	  user.id: 1
//
//	Author Features:
//	  author.id: 1
//
//	--- Accessing Arrow Table ---
//	Number of rows: 1
//	Number of columns: 5
//	Column names: [user.id author.id __id__ __chalk__.__result_metadata__.user.id __chalk__.__result_metadata__.author.id]
//
//	--- Query Metadata ---
//	Query ID: 60c54890-4ef6-4121-ad1b-3ee065553ac9
//	Environment ID: <your-environment-id>
//	Query Timestamp: 2026-01-07 18:38:21.505303883 +0000 UTC

func main() {
	// User represents the user feature namespace
	type User struct {
		Id *int64
	}

	// Author represents the author feature namespace
	type Author struct {
		Id *string
	}

	// MultiNamespaceResult combines features from multiple namespaces
	type MultiNamespaceResult struct {
		User   User
		Author Author
	}

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
	// 		ClientId:      "your-client-id",
	// 		ClientSecret:  "your-client-secret",
	// 		ApiServer:     "https://api.chalk.ai",
	// 		EnvironmentId: "your-environment-id",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to create Chalk client: %v", err)
	// }

	// Execute an online query with multiple input namespaces
	// This example queries features from both user and author namespaces
	// Note: OnlineQueryBulk is used with single-element slices for a single query
	res, err := chalkClient.OnlineQueryBulk(
		ctx,
		chalk.OnlineQueryParams{
			QueryName:     "multi_namespace_example",
			CorrelationId: "multi-ns-123",
			IncludeMeta:   true,
		}.
			// Provide inputs from multiple namespaces
			// Use slices with one element for a single query
			WithInput("user.id", []int64{1}).
			WithInput("author.id", []string{"1"}).
			// Request outputs from both namespaces
			// Including the input features themselves
			WithOutputs(
				// User namespace features
				"user.id",
				// Author namespace features
				"author.id",
			),
	)

	if err != nil {
		log.Printf("Query completed with errors: %v", err)
	}

	// Method 1: Unmarshal into a slice (bulk results)
	var results []MultiNamespaceResult
	if err := res.UnmarshalInto(&results); err != nil {
		log.Fatalf("Failed to unmarshal results: %v", err)
	}

	// Get the first result (since we queried with one set of inputs)
	if len(results) == 0 {
		log.Fatal("No results returned")
	}
	result := results[0]

	fmt.Println("=== Multi-Namespace Query Results ===")
	fmt.Println("\nUser Features:")
	if result.User.Id != nil {
		fmt.Printf("  user.id: %d\n", *result.User.Id)
	}

	fmt.Println("\nAuthor Features:")
	if result.Author.Id != nil {
		fmt.Printf("  author.id: %s\n", *result.Author.Id)
	}

	// Method 2: Access the Arrow table directly
	fmt.Println("\n--- Accessing Arrow Table ---")
	table, err := res.GetTable()
	if err != nil {
		log.Fatalf("Failed to get table: %v", err)
	}
	defer table.Release()
	fmt.Printf("Number of rows: %d\n", table.NumRows())
	fmt.Printf("Number of columns: %d\n", table.NumCols())

	// Print column names
	schema := table.Schema()
	columnNames := make([]string, len(schema.Fields()))
	for i, field := range schema.Fields() {
		columnNames[i] = field.Name
	}
	fmt.Printf("Column names: %v\n", columnNames)

	// Method 3: Access query metadata
	if meta := res.GetQueryMeta(); meta != nil {
		fmt.Printf("\n--- Query Metadata ---\n")
		fmt.Printf("Query ID: %s\n", meta.QueryId)
		fmt.Printf("Environment ID: %s\n", meta.EnvironmentId)
		if meta.QueryTimestamp != nil {
			fmt.Printf("Query Timestamp: %v\n", *meta.QueryTimestamp)
		}
	}

	// Method 4: Check for any errors
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
}
