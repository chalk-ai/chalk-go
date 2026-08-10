package chalk

import (
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/envfs"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Every client must be constructible from a pre-issued JWT alone, with no client
// credentials and no readable chalk.yml. That is the whole point of accepting a
// JWT: the caller has already authenticated and has deliberately not been given
// a secret.
//
// This used to fail for all three clients, because config.NewManager rejected
// empty credentials before auth.NewManager ever looked at the token. It only
// appeared to work when a chalk.yml happened to exist, so these tests point
// ConfigDir at an empty directory and inject an empty environment rather than
// trusting whatever the ambient machine has.

func preIssuedTestJWT() *serverv1.GetTokenResponse {
	return &serverv1.GetTokenResponse{
		AccessToken: "header.payload.signature",
		ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
	}
}

func TestClientsAcceptPreIssuedJWTWithoutCredentials(t *testing.T) {
	t.Parallel()

	t.Run("json client", func(t *testing.T) {
		t.Parallel()
		emptyDir := t.TempDir()
		ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

		client, err := NewClient(ctx, &ClientConfig{
			ApiServer:                  "https://api.chalk.ai",
			EnvironmentId:              "env-abc",
			ConfigDir:                  &emptyDir,
			JWT:                        preIssuedTestJWT(),
			SkipEnvironmentNameMapping: true,
			SkipEngineMapping:          true,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("grpc client", func(t *testing.T) {
		t.Parallel()
		emptyDir := t.TempDir()
		ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

		client, err := NewGRPCClient(ctx, &GRPCClientConfig{
			ApiServer:                  "https://api.chalk.ai",
			QueryServer:                "https://engine.chalk.ai",
			EnvironmentId:              "env-abc",
			ConfigDir:                  &emptyDir,
			JWT:                        preIssuedTestJWT(),
			SkipEnvironmentNameMapping: true,
			SkipEngineMapping:          true,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("volume client", func(t *testing.T) {
		t.Parallel()
		emptyDir := t.TempDir()
		ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

		client, err := NewVolumeClient(ctx, &VolumeClientConfig{
			ApiServer:                  "https://api.chalk.ai",
			EnvironmentId:              "env-abc",
			ConfigDir:                  &emptyDir,
			JWT:                        preIssuedTestJWT(),
			SkipEnvironmentNameMapping: true,
			SkipEngineMapping:          true,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})
}

// TestClientRejectsNeitherJWTNorCredentials keeps the invariant that removing the
// config-level check did not remove it altogether -- auth.NewManager still
// refuses when the caller has nothing usable, and says so clearly.
func TestClientRejectsNeitherJWTNorCredentials(t *testing.T) {
	t.Parallel()

	emptyDir := t.TempDir()
	ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

	_, err := NewClient(ctx, &ClientConfig{
		ApiServer:     "https://api.chalk.ai",
		EnvironmentId: "env-abc",
		ConfigDir:     &emptyDir,
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no JWT and no ClientId/ClientSecret provided")
	// The unreadable chalk.yml is surfaced as the reason they were missing. That
	// detail used to come from config, which raised the error itself; moving the
	// check to auth means it now has to cross the package boundary.
	assert.Contains(t, err.Error(), "auth config file")
}

// TestClientRejectsPartialCredentials pins that an id without a secret still
// fails locally. config.NewManager used to catch this with an || check; moving
// the invariant into auth.NewManager kept the || rather than loosening it to &&,
// so this does not become a confusing server-side rejection.
func TestClientRejectsPartialCredentials(t *testing.T) {
	t.Parallel()

	emptyDir := t.TempDir()
	ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

	_, err := NewClient(ctx, &ClientConfig{
		ApiServer:     "https://api.chalk.ai",
		EnvironmentId: "env-abc",
		ConfigDir:     &emptyDir,
		ClientId:      "token-abc",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no JWT and no ClientId/ClientSecret provided")
}
