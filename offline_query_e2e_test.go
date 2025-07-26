package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestOfflineQuery(t *testing.T) {
	// This test is a placeholder for future offline query tests.
	// It currently does not perform any assertions or checks.
	// The purpose is to ensure that the test suite runs without errors.
	t.Run("Placeholder for future offline query tests", func(t *testing.T) {
		client, err := NewClient(t.Context())
		assert.NoError(t, err)

		params := OfflineQueryParams{}.WithFileInput("gs://").WithOutputs("user.id")

		_, err = client.OfflineQuery(t.Context(), params)
	})
}
