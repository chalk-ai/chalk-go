package config

import (
	"testing"

	"github.com/chalk-ai/chalk-go/envfs"
	"github.com/stretchr/testify/assert"
)

// TestManagerDoesNotRequireCredentials pins that config resolution succeeds
// without client credentials.
//
// It used to fail outright, which made a pre-issued JWT unusable on its own:
// resolution runs before auth.NewManager, so it rejected the caller before
// anything had looked at the token. Only auth.NewManager sees both, so only it
// can tell whether the caller has a usable credential.
func TestManagerDoesNotRequireCredentials(t *testing.T) {
	t.Parallel()

	// An empty directory, so no chalk.yml can supply credentials by accident --
	// a stray one in the ambient home directory is what masked this originally.
	emptyDir := t.TempDir()
	ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})

	manager, err := NewManager(ctx, &ManagerInputs{ConfigDir: &emptyDir})

	assert.NoError(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, "", string(manager.ClientId.Value))
	assert.Equal(t, "", string(manager.ClientSecret.Value))

	// The reason credentials are missing is retained rather than discarded, so
	// auth.NewManager can name the file it searched if nothing else supplies them.
	assert.Error(t, manager.ProjectConfigErr())
}

// TestManagerResolvesCredentialsFromEnvironment confirms the credential sources
// still work, so the check above is not passing merely because resolution broke.
func TestManagerResolvesCredentialsFromEnvironment(t *testing.T) {
	t.Parallel()

	ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{
		Env: map[string]string{
			"CHALK_CLIENT_ID":     "token-abc",
			"CHALK_CLIENT_SECRET": "ts-abc",
		},
	})
	emptyDir := t.TempDir()

	manager, err := NewManager(ctx, &ManagerInputs{ConfigDir: &emptyDir})

	assert.NoError(t, err)
	assert.Equal(t, "token-abc", string(manager.ClientId.Value))
	assert.Equal(t, "ts-abc", string(manager.ClientSecret.Value))
}
