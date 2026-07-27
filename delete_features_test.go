package chalk

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteFeaturesValidateOk(t *testing.T) {
	t.Parallel()
	require.NoError(t, DeleteFeaturesParams{
		Namespace:   "user",
		Features:    []string{"email"},
		PrimaryKeys: []string{"1"},
	}.validate())

	// Neither selector is legal: it means "every feature in the namespace".
	require.NoError(t, DeleteFeaturesParams{
		Namespace:   "user",
		PrimaryKeys: []string{"1"},
	}.validate())

	require.NoError(t, DeleteFeaturesParams{
		Namespace:   "user",
		Tags:        []string{"pii"},
		PrimaryKeys: []string{"1"},
	}.validate())
}

func TestDeleteFeaturesValidateRequiresNamespace(t *testing.T) {
	t.Parallel()
	err := DeleteFeaturesParams{PrimaryKeys: []string{"1"}}.validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "namespace is required")
}

func TestDeleteFeaturesValidateRequiresPrimaryKeys(t *testing.T) {
	t.Parallel()
	// Without primary keys the request has no target rows; the server would either
	// no-op or delete far more than intended, so fail before sending.
	err := DeleteFeaturesParams{Namespace: "user"}.validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "primary key")
}

func TestDeleteFeaturesValidateRejectsFeaturesAndTags(t *testing.T) {
	t.Parallel()
	err := DeleteFeaturesParams{
		Namespace:   "user",
		Features:    []string{"email"},
		Tags:        []string{"pii"},
		PrimaryKeys: []string{"1"},
	}.validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mutually exclusive")
}

func TestDeleteFeaturesRequestWireFormat(t *testing.T) {
	t.Parallel()
	// Must match chalkpy's FeatureObservationDeletionRequest exactly.
	body, err := json.Marshal(deleteFeaturesRequest{
		Namespace:     "user",
		Features:      []string{"email", "name"},
		Tags:          nil,
		PrimaryKeys:   []string{"1", "2"},
		RetainOffline: true,
		RetainOnline:  false,
	})
	require.NoError(t, err)

	var got map[string]any
	require.NoError(t, json.Unmarshal(body, &got))
	assert.Equal(t, "user", got["namespace"])
	assert.Equal(t, []any{"email", "name"}, got["features"])
	assert.Nil(t, got["tags"])
	assert.Equal(t, []any{"1", "2"}, got["primary_keys"])
	assert.Equal(t, true, got["retain_offline"])
	assert.Equal(t, false, got["retain_online"])

	// The booleans must be present even when false -- omitempty would drop
	// retain_online and change the server-side default.
	_, hasRetainOnline := got["retain_online"]
	assert.True(t, hasRetainOnline, "retain_online must be sent explicitly, not omitted")
}
