package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/chalk-ai/chalk-go"
)

func main() {
	// Initialize the Chalk gRPC client
	ctx := context.Background()
	chalkClient, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

	// Execute a query that might return errors
	res, err := chalkClient.OnlineQueryBulk(
		ctx,
		chalk.OnlineQueryParams{
			QueryName: "error_handling_example",
		}.
			WithInput("user.id", []string{"user_1", "user_2"}).
			WithOutputs("user.id", "user.email", "user.credit_score"),
	)

	// Method 1: Check for ServerErrors using errors.As
	// This allows you to specifically handle Chalk server errors
	// (e.g., resolver failures, feature computation errors)
	var serverErrs chalk.ServerErrors
	if errors.As(err, &serverErrs) {
		fmt.Printf("Encountered %d server error(s):\n", len(serverErrs))
		for i, serverErr := range serverErrs {
			fmt.Printf("\nError %d:\n", i+1)
			fmt.Printf("  Code: %s\n", serverErr.Code)
			fmt.Printf("  Category: %s\n", serverErr.Category)
			fmt.Printf("  Message: %s\n", serverErr.Message)

			if serverErr.Feature != "" {
				fmt.Printf("  Feature: %s\n", serverErr.Feature)
			}
			if serverErr.Resolver != "" {
				fmt.Printf("  Resolver: %s\n", serverErr.Resolver)
			}
			if serverErr.Exception != nil {
				fmt.Printf("  Exception: %s\n", serverErr.Exception.Message)
				if serverErr.Exception.Stacktrace != "" {
					fmt.Printf("  Stacktrace:\n%s\n", serverErr.Exception.Stacktrace)
				}
			}
		}
		// You can still access partial results even when server errors occur
		if res.GetQueryMeta() != nil {
			fmt.Printf("\nPartial results available for query %s\n", res.GetQueryMeta().QueryId)
		}
	}

	// Method 2: Check for HTTP errors
	// These are errors related to API communication (auth, network, etc.)
	var httpErr *chalk.HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("HTTP Error occurred:\n")
		fmt.Printf("  Status Code: %d\n", httpErr.StatusCode)
		fmt.Printf("  Path: %s\n", httpErr.Path)
		fmt.Printf("  Message: %s\n", httpErr.Message)
		if httpErr.Trace != nil {
			fmt.Printf("  Trace ID: %s\n", *httpErr.Trace)
		}
		return // HTTP errors typically mean no partial results are available
	}

	// Method 3: Generic error handling
	if err != nil {
		log.Printf("Query encountered an error: %v", err)
		// Note: Even with errors, res might contain partial data
	}

	// Method 4: Access errors from the result object
	// This is useful when you have a successful response but want to check
	// if any features failed to compute
	if errors, err := res.GetErrors(); err == nil && len(errors) > 0 {
		fmt.Printf("\nFound %d feature-level errors in results:\n", len(errors))
		for _, featureErr := range errors {
			fmt.Printf("  - %s: %s\n", featureErr.Feature, featureErr.Message)
		}
	}

	// Method 5: Type assertion (alternative to errors.As)
	if err != nil {
		switch e := err.(type) {
		case chalk.ServerErrors:
			fmt.Printf("Directly type-asserted ServerErrors: %d error(s)\n", len(e))
		case *chalk.HTTPError:
			fmt.Printf("Directly type-asserted HTTPError: %d\n", e.StatusCode)
		default:
			fmt.Printf("Unknown error type: %v\n", err)
		}
	}

	// If no errors occurred, process the results normally
	if err == nil {
		fmt.Println("Query completed successfully!")
		// Process results...
	}
}
