// Package testserver provides an in-process mock RPC server for testing
// applications that use Connect RPC.
//
// The mock server is designed to be reusable across multiple projects and
// provides a realistic testing environment by implementing actual HTTP/RPC
// handlers rather than using interface mocks.
//
// # Basic Usage
//
// Create a mock server and configure responses:
//
//	func TestMyFeature(t *testing.T) {
//	    server := testserver.NewMockBuilderServer(t)
//	    defer server.Close()
//
//	    // Configure mock response
//	    server.OnGetClusterTimescaleDB().Return(&serverv1.GetClusterTimescaleDBResponse{
//	        Id: "test-db-id",
//	        Name: "test-db",
//	    })
//
//	    // Use server.URL to create a real client
//	    client := serverv1connect.NewBuilderServiceClient(http.DefaultClient, server.URL)
//	    resp, err := client.GetClusterTimescaleDB(ctx, connect.NewRequest(&serverv1.GetClusterTimescaleDBRequest{
//	        Id: "test-db-id",
//	    }))
//
//	    require.NoError(t, err)
//	    assert.Equal(t, "test-db-id", resp.Msg.Id)
//	}
//
// # Error Testing
//
// Configure methods to return errors:
//
//	server.OnUpdateClusterTimescaleDB().ReturnError(errors.New("update failed"))
//
//	client := serverv1connect.NewBuilderServiceClient(http.DefaultClient, server.URL)
//	_, err := client.UpdateClusterTimescaleDB(ctx, connect.NewRequest(...))
//
//	require.Error(t, err)
//	assert.Contains(t, err.Error(), "update failed")
//
// # Request Capture
//
// Capture and assert on requests sent to the server:
//
//	server.OnGetClusterTimescaleDB().Return(&serverv1.GetClusterTimescaleDBResponse{...})
//
//	client := serverv1connect.NewBuilderServiceClient(http.DefaultClient, server.URL)
//	client.GetClusterTimescaleDB(ctx, connect.NewRequest(&serverv1.GetClusterTimescaleDBRequest{
//	    Id: "test-id",
//	}))
//
//	requests := server.GetCapturedRequests("GetClusterTimescaleDB")
//	require.Len(t, requests, 1)
//	capturedReq := requests[0].(*serverv1.GetClusterTimescaleDBRequest)
//	assert.Equal(t, "test-id", capturedReq.Id)
//
// # Custom Behaviors
//
// For complex scenarios, use custom behavior functions:
//
//	server.OnGetClusterTimescaleDB().WithBehavior(func(req proto.Message) (proto.Message, error) {
//	    getReq := req.(*serverv1.GetClusterTimescaleDBRequest)
//	    if getReq.Id == "invalid" {
//	        return nil, errors.New("not found")
//	    }
//	    return &serverv1.GetClusterTimescaleDBResponse{
//	        Id: getReq.Id,
//	        Name: "dynamic-name",
//	    }, nil
//	})
//
// # Thread Safety
//
// The mock server is thread-safe and can be used with t.Parallel() tests.
// Each test should create its own server instance to ensure isolation.
package testserver
