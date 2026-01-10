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
//	Query ID: 56a928b4-975f-4940-a590-141f6469d8dc
//	Environment ID: <your-environment-id>
//	Query Timestamp: 2026-01-07 18:38:43.060682528 +0000 UTC

//lint:ignore U1000 example
func multiNamespaceQuery() {
	// User represents the user feature namespace
	// In practice, use `chalk codegen` to generate these structs
	type User struct {
		Id          *int64
		Email       *string
		Name        *string
		CreditScore *int64
	}

	// Author represents the author feature namespace
	// In practice, use `chalk codegen` to generate these structs
	type Author struct {
		Id           *string
		Name         *string
		ArticleCount *int64
		Rating       *float64
	}

	// MultiNamespaceResult combines features from multiple namespaces
	type MultiNamespaceResult struct {
		User   User
		Author Author
	}

	// Initialize the Chalk gRPC client
	// By default, this will use environment variables or ~/.chalk.yml configuration
	ctx := context.Background()
	chalkClient, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

	// Execute an online query with multiple input namespaces
	// This example queries features from both user and author namespaces
	// Note: Use OnlineQueryBulk with single-element slices for a single query
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
				// Uncomment additional features if they exist in your environment:
				// "user.email",
				// "user.name",
				// "user.credit_score",
				// Author namespace features
				"author.id",
				// Uncomment additional features if they exist in your environment:
				// "author.name",
				// "author.article_count",
				// "author.rating",
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
	if result.User.Email != nil {
		fmt.Printf("  user.email: %s\n", *result.User.Email)
	}
	if result.User.Name != nil {
		fmt.Printf("  user.name: %s\n", *result.User.Name)
	}
	if result.User.CreditScore != nil {
		fmt.Printf("  user.credit_score: %d\n", *result.User.CreditScore)
	}

	fmt.Println("\nAuthor Features:")
	if result.Author.Id != nil {
		fmt.Printf("  author.id: %s\n", *result.Author.Id)
	}
	if result.Author.Name != nil {
		fmt.Printf("  author.name: %s\n", *result.Author.Name)
	}
	if result.Author.ArticleCount != nil {
		fmt.Printf("  author.article_count: %d\n", *result.Author.ArticleCount)
	}
	if result.Author.Rating != nil {
		fmt.Printf("  author.rating: %.2f\n", *result.Author.Rating)
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
