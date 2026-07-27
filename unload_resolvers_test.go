package chalk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializeUnloadResolversEmpty(t *testing.T) {
	t.Parallel()
	// nil rather than an empty slice, so `omitempty` drops the key entirely and the
	// request stays byte-identical to one that never mentioned unload resolvers.
	serialized, err := serializeUnloadResolvers(nil)
	require.NoError(t, err)
	assert.Nil(t, serialized)

	serialized, err = serializeUnloadResolvers([]UnloadResolver{})
	require.NoError(t, err)
	assert.Nil(t, serialized)
}

func TestSerializeUnloadResolversAll(t *testing.T) {
	t.Parallel()
	serialized, err := serializeUnloadResolvers([]UnloadResolver{UnloadAllResolvers()})
	require.NoError(t, err)
	require.Len(t, serialized, 1)
	assert.Equal(t, "*", serialized[0].Fqn)
	assert.Empty(t, serialized[0].PartitionBy)
}

func TestSerializeUnloadResolversByName(t *testing.T) {
	t.Parallel()
	serialized, err := serializeUnloadResolvers(UnloadResolversByName("resolver_a", "resolver_b"))
	require.NoError(t, err)
	require.Len(t, serialized, 2)
	assert.Equal(t, "resolver_a", serialized[0].Fqn)
	assert.Equal(t, "resolver_b", serialized[1].Fqn)
	assert.Empty(t, serialized[0].PartitionBy)
}

func TestSerializeUnloadResolversPartitionByEquality(t *testing.T) {
	t.Parallel()
	// The cross-namespace form. This is what actually buckets the unload, and the engine
	// parses it as plain text rather than a base64 proto.
	serialized, err := serializeUnloadResolvers([]UnloadResolver{{
		Fqn:         "unload_txns",
		PartitionBy: []UnloadPartition{PartitionByEquality("txn.user_id", "user.id")},
	}})
	require.NoError(t, err)
	require.Len(t, serialized, 1)
	assert.Equal(t, "unload_txns", serialized[0].Fqn)
	assert.Equal(t, []string{"txn.user_id == user.id"}, serialized[0].PartitionBy)
}

func TestSerializeUnloadResolversPartitionBySelf(t *testing.T) {
	t.Parallel()
	// Self-equality is sent as the bare FQN, the shorthand the engine expands.
	serialized, err := serializeUnloadResolvers([]UnloadResolver{{
		Fqn:         "unload_txns",
		PartitionBy: []UnloadPartition{PartitionBySelf("txn.user_id")},
	}})
	require.NoError(t, err)
	require.Len(t, serialized, 1)
	assert.Equal(t, []string{"txn.user_id"}, serialized[0].PartitionBy)
}

func TestSerializeUnloadResolversRejectsPartitionedWildcard(t *testing.T) {
	t.Parallel()
	// Mirrors chalkpy, which raises on the same combination: "*" stands for every eligible
	// resolver, so no single join key can apply to it.
	_, err := serializeUnloadResolvers([]UnloadResolver{{
		Fqn:         UnloadAllResolversFqn,
		PartitionBy: []UnloadPartition{PartitionByEquality("txn.user_id", "user.id")},
	}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not accept partition expressions")
}

func TestSerializeUnloadResolversRejectsMissingFqn(t *testing.T) {
	t.Parallel()
	_, err := serializeUnloadResolvers([]UnloadResolver{{Fqn: ""}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing a resolver FQN")
}

func TestSerializeUnloadResolversRejectsBadFeatureReference(t *testing.T) {
	t.Parallel()
	// A non-feature, non-string reference must fail loudly rather than serializing to
	// something the engine would silently misparse.
	_, err := serializeUnloadResolvers([]UnloadResolver{{
		Fqn:         "unload_txns",
		PartitionBy: []UnloadPartition{PartitionByEquality(12345, "user.id")},
	}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unload_txns")
}

func TestWithUnloadResolversAppends(t *testing.T) {
	t.Parallel()
	params := OfflineQueryParams{}.
		WithInput("user.id", []any{1}).
		WithOutputs("user.name").
		WithUnloadResolvers(UnloadResolver{Fqn: "resolver_a"}).
		WithUnloadResolvers(UnloadResolver{Fqn: "resolver_b"})

	require.Len(t, params.underlying.UnloadResolvers, 2)
	assert.Equal(t, "resolver_a", params.underlying.UnloadResolvers[0].Fqn)
	assert.Equal(t, "resolver_b", params.underlying.UnloadResolvers[1].Fqn)
}

func TestWithUnloadAllResolvers(t *testing.T) {
	t.Parallel()
	params := OfflineQueryParams{}.
		WithInput("user.id", []any{1}).
		WithOutputs("user.name").
		WithUnloadAllResolvers()

	require.Len(t, params.underlying.UnloadResolvers, 1)
	assert.Equal(t, UnloadAllResolversFqn, params.underlying.UnloadResolvers[0].Fqn)
}
