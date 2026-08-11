package auth

import (
	"context"
	"testing"
	"time"

	"github.com/chalk-ai/chalk-go/config"
	"github.com/chalk-ai/chalk-go/envfs"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// These exercise the AuthProvider path without a network. A provider replaces
// the client-credentials exchange entirely, so every assertion here would
// otherwise require a live api-server to disprove.

func tokenValidFor(d time.Duration) *serverv1.GetTokenResponse {
	return &serverv1.GetTokenResponse{
		AccessToken: "header.payload.signature",
		ExpiresAt:   timestamppb.New(time.Now().Add(d)),
	}
}

func authValidFor(environmentID string, d time.Duration) *AuthSnapshot {
	return &AuthSnapshot{Token: tokenValidFor(d), EnvironmentID: environmentID}
}

// managerWithoutCredentials builds a config.Manager with no credentials from any
// source, so anything that works is working via the token or the provider.
func managerWithoutCredentials(t *testing.T) *config.Manager {
	t.Helper()
	emptyDir := t.TempDir()
	ctx := envfs.ContextWithEnvironmentGetter(t.Context(), envfs.MapEnvironmentGetter{})
	manager, err := config.NewManager(ctx, &config.ManagerInputs{
		APIServer:     "https://api.chalk.ai",
		EnvironmentId: "env-abc",
		ConfigDir:     &emptyDir,
	})
	require.NoError(t, err)
	require.Equal(t, "", string(manager.ClientId.Value))
	return manager
}

func newManagerForTest(t *testing.T, in *Inputs) (*Manager, error) {
	t.Helper()
	in.Config = managerWithoutCredentials(t)
	in.SkipEnvironmentNameMapping = true
	in.SkipEngineMapping = true
	return NewManager(t.Context(), in)
}

func TestAuthProviderSuppliesFirstSnapshot(t *testing.T) {
	t.Parallel()

	calls := 0
	m, err := newManagerForTest(t, &Inputs{
		AuthProvider: func(context.Context) (*AuthSnapshot, error) {
			calls++
			return authValidFor("env-provider", time.Hour), nil
		},
	})

	// A provider alone is sufficient: no Token and no credentials.
	require.NoError(t, err)
	assert.Equal(t, 1, calls)
	authSnapshot, err := m.GetAuth(t.Context(), time.Now())
	assert.NoError(t, err)
	assert.Equal(t, "header.payload.signature", authSnapshot.Token.AccessToken)
	assert.Equal(t, "env-provider", authSnapshot.EnvironmentID)
}

func TestAuthProviderRefreshesTokenAndEnvironmentTogether(t *testing.T) {
	t.Parallel()

	calls := 0
	// Valid at construction, but not by enough to satisfy the headroom callers
	// ask for, so the next GetJWT must go back to the provider.
	m, err := newManagerForTest(t, &Inputs{
		Token: tokenValidFor(30 * time.Second),
		AuthProvider: func(context.Context) (*AuthSnapshot, error) {
			calls++
			return authValidFor("env-rotated", time.Hour), nil
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 0, calls, "the supplied token should be used as-is at construction")

	authSnapshot, err := m.GetAuth(t.Context(), time.Now().Add(time.Minute))

	// Without the provider this would attempt a credentials exchange with two
	// empty strings and fail against the api-server.
	assert.NoError(t, err)
	assert.Equal(t, 1, calls)
	assert.Equal(t, "env-rotated", authSnapshot.EnvironmentID)
}

func TestAuthProviderNotCalledWhileSnapshotIsFresh(t *testing.T) {
	t.Parallel()

	calls := 0
	m, err := newManagerForTest(t, &Inputs{
		Token: tokenValidFor(time.Hour),
		AuthProvider: func(context.Context) (*AuthSnapshot, error) {
			calls++
			return authValidFor("env-provider", time.Hour), nil
		},
	})
	require.NoError(t, err)

	for range 5 {
		_, err = m.GetJWT(t.Context(), time.Now().Add(time.Minute))
		assert.NoError(t, err)
	}

	assert.Equal(t, 0, calls, "a fresh token must be served from cache")
}

func TestAuthProviderErrorIsPropagated(t *testing.T) {
	t.Parallel()

	_, err := newManagerForTest(t, &Inputs{
		AuthProvider: func(context.Context) (*AuthSnapshot, error) {
			return nil, assert.AnError
		},
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "refreshing pre-issued token")
}

// TestAuthProviderResultIsValidated covers a callback misbehaving. These are
// caller-supplied functions, and each of these returns would otherwise be cached
// and fail much later -- as a nil dereference in an interceptor, or a bare
// "Bearer " header, or an endless retry loop.
func TestAuthProviderResultIsValidated(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		returned *AuthSnapshot
		expected string
	}{
		{
			name:     "nil token",
			returned: nil,
			expected: "auth snapshot is nil",
		},
		{
			name:     "empty access token",
			returned: &AuthSnapshot{Token: &serverv1.GetTokenResponse{ExpiresAt: timestamppb.New(time.Now().Add(time.Hour))}, EnvironmentID: "env-provider"},
			expected: "empty access token",
		},
		{
			// A missing ExpiresAt reads as the Unix epoch, so it lands in the
			// staleness check rather than needing its own rule.
			name:     "no expiry",
			returned: &AuthSnapshot{Token: &serverv1.GetTokenResponse{AccessToken: "a.b.c"}, EnvironmentID: "env-provider"},
			expected: "may no longer be refreshed",
		},
		{
			name:     "already expired",
			returned: authValidFor("env-provider", -time.Hour),
			expected: "may no longer be refreshed",
		},
		{
			name:     "empty environment",
			returned: authValidFor("", time.Hour),
			expected: "empty environment ID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := newManagerForTest(t, &Inputs{
				AuthProvider: func(context.Context) (*AuthSnapshot, error) {
					return tc.returned, nil
				},
			})

			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.expected)
		})
	}
}

// TestSuppliedTokenIsValidated pins the checks on a caller-supplied Token.
//
// Deliberately narrower than the provider path: a missing ExpiresAt is not
// rejected here, because a Token can be combined with client credentials that
// pick up the slack, and this repo's own grpc_client_branch_test.go does exactly
// that. Rejecting it would be a breaking change for a shipped pattern.
func TestSuppliedTokenIsValidated(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		token    *serverv1.GetTokenResponse
		expected string
	}{
		{
			name:     "empty access token",
			token:    &serverv1.GetTokenResponse{ExpiresAt: timestamppb.New(time.Now().Add(time.Hour))},
			expected: "empty access token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := newManagerForTest(t, &Inputs{Token: tc.token})

			require.Error(t, err)
			assert.Contains(t, err.Error(), "invalid pre-issued JWT")
			assert.Contains(t, err.Error(), tc.expected)
		})
	}
}
