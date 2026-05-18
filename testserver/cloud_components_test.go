package testserver_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/testserver"
	"google.golang.org/protobuf/proto"
)

// cloudComponentsClient returns a real CloudComponentsService client wired
// against the mock server's HTTP endpoint.
func cloudComponentsClient(server *testserver.MockServer) serverv1connect.CloudComponentsServiceClient {
	return serverv1connect.NewCloudComponentsServiceClient(http.DefaultClient, server.URL)
}

func TestGetCloudComponentCluster(t *testing.T) {
	t.Parallel()

	t.Run("Return", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		want := &serverv1.GetCloudComponentClusterResponse{
			Cluster: &serverv1.CloudComponentClusterResponse{Id: "cluster-id", Kind: "EKS_STANDARD"},
		}
		server.OnGetCloudComponentCluster().Return(want)

		got, err := cloudComponentsClient(server).GetCloudComponentCluster(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentClusterRequest{Id: "cluster-id"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Cluster.Id != "cluster-id" {
			t.Errorf("Cluster.Id = %q, want %q", got.Msg.Cluster.Id, "cluster-id")
		}
		if got.Msg.Cluster.Kind != "EKS_STANDARD" {
			t.Errorf("Cluster.Kind = %q, want %q", got.Msg.Cluster.Kind, "EKS_STANDARD")
		}
	})

	t.Run("ReturnError", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentCluster().ReturnError(
			connect.NewError(connect.CodePermissionDenied, errors.New("nope")),
		)

		_, err := cloudComponentsClient(server).GetCloudComponentCluster(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentClusterRequest{Id: "cluster-id"}),
		)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if connect.CodeOf(err) != connect.CodePermissionDenied {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodePermissionDenied)
		}
	})

	t.Run("NoMockConfigured", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		_, err := cloudComponentsClient(server).GetCloudComponentCluster(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentClusterRequest{Id: "cluster-id"}),
		)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeNotFound)
		}
	})

	t.Run("WithBehavior", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentCluster().WithBehavior(func(req proto.Message) (proto.Message, error) {
			r := req.(*serverv1.GetCloudComponentClusterRequest)
			return &serverv1.GetCloudComponentClusterResponse{
				Cluster: &serverv1.CloudComponentClusterResponse{Id: r.Id + "-echo"},
			}, nil
		})

		got, err := cloudComponentsClient(server).GetCloudComponentCluster(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentClusterRequest{Id: "abc"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Cluster.Id != "abc-echo" {
			t.Errorf("Cluster.Id = %q, want %q", got.Msg.Cluster.Id, "abc-echo")
		}
	})

	t.Run("CapturesRequests", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentCluster().Return(&serverv1.GetCloudComponentClusterResponse{
			Cluster: &serverv1.CloudComponentClusterResponse{Id: "x"},
		})

		_, err := cloudComponentsClient(server).GetCloudComponentCluster(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentClusterRequest{Id: "captured-id"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		captured := server.GetCapturedRequests("GetCloudComponentCluster")
		if len(captured) != 1 {
			t.Fatalf("captured %d requests, want 1", len(captured))
		}
		req := captured[0].(*serverv1.GetCloudComponentClusterRequest)
		if req.Id != "captured-id" {
			t.Errorf("captured Id = %q, want %q", req.Id, "captured-id")
		}
	})
}

func TestGetCloudComponentVpc(t *testing.T) {
	t.Parallel()

	t.Run("Return", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		want := &serverv1.GetCloudComponentVpcResponse{
			Vpc: &serverv1.CloudComponentVpcResponse{Id: "vpc-id", Kind: "aws"},
		}
		server.OnGetCloudComponentVpc().Return(want)

		got, err := cloudComponentsClient(server).GetCloudComponentVpc(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentVpcRequest{Id: "vpc-id"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Vpc.Id != "vpc-id" {
			t.Errorf("Vpc.Id = %q, want %q", got.Msg.Vpc.Id, "vpc-id")
		}
		if got.Msg.Vpc.Kind != "aws" {
			t.Errorf("Vpc.Kind = %q, want %q", got.Msg.Vpc.Kind, "aws")
		}
	})

	t.Run("ReturnError", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentVpc().ReturnError(
			connect.NewError(connect.CodeInternal, errors.New("boom")),
		)

		_, err := cloudComponentsClient(server).GetCloudComponentVpc(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentVpcRequest{Id: "vpc-id"}),
		)
		if connect.CodeOf(err) != connect.CodeInternal {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeInternal)
		}
	})

	t.Run("NoMockConfigured", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		_, err := cloudComponentsClient(server).GetCloudComponentVpc(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentVpcRequest{Id: "vpc-id"}),
		)
		if connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeNotFound)
		}
	})

	t.Run("WithBehavior", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentVpc().WithBehavior(func(req proto.Message) (proto.Message, error) {
			r := req.(*serverv1.GetCloudComponentVpcRequest)
			return &serverv1.GetCloudComponentVpcResponse{
				Vpc: &serverv1.CloudComponentVpcResponse{Id: r.Id + "-echo"},
			}, nil
		})

		got, err := cloudComponentsClient(server).GetCloudComponentVpc(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentVpcRequest{Id: "abc"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Vpc.Id != "abc-echo" {
			t.Errorf("Vpc.Id = %q, want %q", got.Msg.Vpc.Id, "abc-echo")
		}
	})

	t.Run("CapturesRequests", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudComponentVpc().Return(&serverv1.GetCloudComponentVpcResponse{
			Vpc: &serverv1.CloudComponentVpcResponse{Id: "x"},
		})

		_, err := cloudComponentsClient(server).GetCloudComponentVpc(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudComponentVpcRequest{Id: "captured-vpc"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		captured := server.GetCapturedRequests("GetCloudComponentVpc")
		if len(captured) != 1 {
			t.Fatalf("captured %d requests, want 1", len(captured))
		}
		req := captured[0].(*serverv1.GetCloudComponentVpcRequest)
		if req.Id != "captured-vpc" {
			t.Errorf("captured Id = %q, want %q", req.Id, "captured-vpc")
		}
	})
}
