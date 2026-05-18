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

func cloudCredentialsClient(server *testserver.MockServer) serverv1connect.CloudAccountCredentialsServiceClient {
	return serverv1connect.NewCloudAccountCredentialsServiceClient(http.DefaultClient, server.URL)
}

// TestGetCloudCredentials covers the entire CloudAccountCredentialsService —
// this service is registered on the mock mux in server.go alongside the others.
// The "Return" subtest exercising a real client call against server.URL
// implicitly verifies that mux registration is correct.
func TestGetCloudCredentials(t *testing.T) {
	t.Parallel()

	t.Run("Return", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		want := &serverv1.GetCloudCredentialsResponse{
			Credentials: &serverv1.CloudCredentialsResponse{Id: "cred-id", Name: "test", Kind: "aws"},
		}
		server.OnGetCloudCredentials().Return(want)

		got, err := cloudCredentialsClient(server).GetCloudCredentials(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudCredentialsRequest{Id: "cred-id"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Credentials.Id != "cred-id" {
			t.Errorf("Credentials.Id = %q, want %q", got.Msg.Credentials.Id, "cred-id")
		}
		if got.Msg.Credentials.Kind != "aws" {
			t.Errorf("Credentials.Kind = %q, want %q", got.Msg.Credentials.Kind, "aws")
		}
	})

	t.Run("ReturnError", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudCredentials().ReturnError(
			connect.NewError(connect.CodeUnauthenticated, errors.New("token expired")),
		)

		_, err := cloudCredentialsClient(server).GetCloudCredentials(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudCredentialsRequest{Id: "cred-id"}),
		)
		if connect.CodeOf(err) != connect.CodeUnauthenticated {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeUnauthenticated)
		}
	})

	t.Run("NoMockConfigured", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		_, err := cloudCredentialsClient(server).GetCloudCredentials(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudCredentialsRequest{Id: "cred-id"}),
		)
		if connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeNotFound)
		}
	})

	t.Run("WithBehavior", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudCredentials().WithBehavior(func(req proto.Message) (proto.Message, error) {
			r := req.(*serverv1.GetCloudCredentialsRequest)
			return &serverv1.GetCloudCredentialsResponse{
				Credentials: &serverv1.CloudCredentialsResponse{Id: r.Id + "-echo"},
			}, nil
		})

		got, err := cloudCredentialsClient(server).GetCloudCredentials(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudCredentialsRequest{Id: "abc"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Credentials.Id != "abc-echo" {
			t.Errorf("Credentials.Id = %q, want %q", got.Msg.Credentials.Id, "abc-echo")
		}
	})

	t.Run("CapturesRequests", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetCloudCredentials().Return(&serverv1.GetCloudCredentialsResponse{
			Credentials: &serverv1.CloudCredentialsResponse{Id: "x"},
		})

		_, err := cloudCredentialsClient(server).GetCloudCredentials(
			context.Background(),
			connect.NewRequest(&serverv1.GetCloudCredentialsRequest{Id: "captured-cred"}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		captured := server.GetCapturedRequests("GetCloudCredentials")
		if len(captured) != 1 {
			t.Fatalf("captured %d requests, want 1", len(captured))
		}
		req := captured[0].(*serverv1.GetCloudCredentialsRequest)
		if req.Id != "captured-cred" {
			t.Errorf("captured Id = %q, want %q", req.Id, "captured-cred")
		}
	})
}
