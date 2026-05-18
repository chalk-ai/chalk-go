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

func teamClient(server *testserver.MockServer) serverv1connect.TeamServiceClient {
	return serverv1connect.NewTeamServiceClient(http.DefaultClient, server.URL)
}

func TestGetTeam(t *testing.T) {
	t.Parallel()

	t.Run("Return", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		want := &serverv1.GetTeamResponse{
			Team: &serverv1.Team{
				Id: "team-id",
				Projects: []*serverv1.Project{
					{Id: "project-id", TeamId: "team-id", Name: "Features"},
				},
			},
		}
		server.OnGetTeam().Return(want)

		got, err := teamClient(server).GetTeam(
			context.Background(),
			connect.NewRequest(&serverv1.GetTeamRequest{}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Msg.Team.Id != "team-id" {
			t.Errorf("Team.Id = %q, want %q", got.Msg.Team.Id, "team-id")
		}
		if len(got.Msg.Team.Projects) != 1 || got.Msg.Team.Projects[0].Id != "project-id" {
			t.Errorf("Team.Projects = %+v, want one project with id 'project-id'", got.Msg.Team.Projects)
		}
	})

	t.Run("ReturnError", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetTeam().ReturnError(
			connect.NewError(connect.CodeFailedPrecondition, errors.New("team archived")),
		)

		_, err := teamClient(server).GetTeam(
			context.Background(),
			connect.NewRequest(&serverv1.GetTeamRequest{}),
		)
		if connect.CodeOf(err) != connect.CodeFailedPrecondition {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeFailedPrecondition)
		}
	})

	t.Run("NoMockConfigured", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		_, err := teamClient(server).GetTeam(
			context.Background(),
			connect.NewRequest(&serverv1.GetTeamRequest{}),
		)
		if connect.CodeOf(err) != connect.CodeNotFound {
			t.Errorf("error code = %v, want %v", connect.CodeOf(err), connect.CodeNotFound)
		}
	})

	t.Run("WithBehavior", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		called := false
		server.OnGetTeam().WithBehavior(func(req proto.Message) (proto.Message, error) {
			called = true
			return &serverv1.GetTeamResponse{Team: &serverv1.Team{Id: "from-behavior"}}, nil
		})

		got, err := teamClient(server).GetTeam(
			context.Background(),
			connect.NewRequest(&serverv1.GetTeamRequest{}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !called {
			t.Error("behavior callback was not invoked")
		}
		if got.Msg.Team.Id != "from-behavior" {
			t.Errorf("Team.Id = %q, want %q", got.Msg.Team.Id, "from-behavior")
		}
	})

	t.Run("CapturesRequests", func(t *testing.T) {
		t.Parallel()
		server := testserver.NewMockBuilderServer(t)
		t.Cleanup(func() { server.Close() })

		server.OnGetTeam().Return(&serverv1.GetTeamResponse{Team: &serverv1.Team{Id: "x"}})

		_, err := teamClient(server).GetTeam(
			context.Background(),
			connect.NewRequest(&serverv1.GetTeamRequest{}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		captured := server.GetCapturedRequests("GetTeam")
		if len(captured) != 1 {
			t.Fatalf("captured %d requests, want 1", len(captured))
		}
		if _, ok := captured[0].(*serverv1.GetTeamRequest); !ok {
			t.Errorf("captured request type = %T, want *serverv1.GetTeamRequest", captured[0])
		}
	})
}
