package chalk

import (
	"fmt"

	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
)

// UnloadAllResolversFqn is the sentinel FQN that tells the server to auto-detect every
// resolver eligible for unloading, rather than naming them individually. It cannot carry
// partition expressions -- there is no single join key that applies to every resolver.
const UnloadAllResolversFqn = "*"

// UnloadResolver names a resolver whose output should be pre-computed and unloaded to
// cloud storage ahead of the shard jobs that read it.
//
// Leaving PartitionBy empty unloads the resolver's output flat. Supplying partition
// expressions hash-buckets that output on the join key, so a sharded query has each shard
// read only the buckets it owns instead of scanning everything.
type UnloadResolver struct {
	// Fqn is the resolver's fully-qualified name, or [UnloadAllResolversFqn] ("*") to let
	// the server pick every eligible resolver.
	Fqn string

	// PartitionBy holds the expressions to bucket the unloaded output on. Empty means an
	// unpartitioned (flat) unload. More than one expression is accepted by the wire format
	// but the engine currently rejects it at plan time.
	PartitionBy []UnloadPartition
}

// UnloadPartition is a single partition expression for an unloaded resolver.
//
// Build one with [PartitionByEquality] (the cross-namespace form that actually buckets) or
// [PartitionBySelf]. Both Left and Right accept the same feature references as
// [OfflineQueryParamsComplete.WithOutputs] -- a codegen'd feature field or a raw FQN string.
type UnloadPartition struct {
	// Left is the left-hand feature of the equality.
	Left any

	// Right is the right-hand feature. A nil Right means self-equality on Left.
	Right any
}

// PartitionByEquality buckets an unloaded resolver's output on an equality between two
// features, e.g. PartitionByEquality(Txn.UserId, User.Id) for a resolver resolving Txn
// joined against a User spine.
//
// This is the form that produces useful bucketing: the scan key has to be joined to another
// namespace, which self-equality cannot express.
func PartitionByEquality(left any, right any) UnloadPartition {
	return UnloadPartition{Left: left, Right: right}
}

// PartitionBySelf buckets on a single feature, shorthand for `feature == feature`.
func PartitionBySelf(feature any) UnloadPartition {
	return UnloadPartition{Left: feature}
}

// UnloadAllResolvers returns the spec asking the server to auto-detect and unload every
// eligible resolver.
func UnloadAllResolvers() UnloadResolver {
	return UnloadResolver{Fqn: UnloadAllResolversFqn}
}

// UnloadResolversByName returns unpartitioned unload specs for the named resolvers.
func UnloadResolversByName(fqns ...string) []UnloadResolver {
	specs := make([]UnloadResolver, 0, len(fqns))
	for _, fqn := range fqns {
		specs = append(specs, UnloadResolver{Fqn: fqn})
	}
	return specs
}

// serialize renders a partition expression into the string the engine decodes.
//
// The engine (engine/chalkengine/metaplanner/manifest_utils.py::_decode_partition_expr)
// accepts a raw FQN, a plain-text "lhs == rhs" equality, or a base64-encoded LogicalExprNode.
// We emit the plain-text forms: they express everything chalk-go needs, including the
// cross-namespace join, without building and encoding a proto. The decoder attempts base64
// first with validate=True, and "lhs == rhs" is never valid base64 (`=` is only accepted as
// trailing padding), so the two forms cannot be confused.
func (u UnloadPartition) serialize() (string, error) {
	leftFqn, _, err := getFqn(u.Left)
	if err != nil {
		return "", errors.Wrap(err, "validating left side of unload partition expression")
	}
	if u.Right == nil {
		return leftFqn, nil
	}
	rightFqn, _, err := getFqn(u.Right)
	if err != nil {
		return "", errors.Wrap(err, "validating right side of unload partition expression")
	}
	return fmt.Sprintf("%s == %s", leftFqn, rightFqn), nil
}

// serializeUnloadResolvers renders the specs into their wire representation, returning nil
// when there is nothing to send so the field is omitted entirely.
func serializeUnloadResolvers(specs []UnloadResolver) ([]internal.UnloadResolverSpecSerialized, error) {
	if len(specs) == 0 {
		return nil, nil
	}

	serialized := make([]internal.UnloadResolverSpecSerialized, 0, len(specs))
	for _, spec := range specs {
		if spec.Fqn == "" {
			return nil, errors.New(
				"unload resolver spec is missing a resolver FQN - set Fqn, or use " +
					"chalk.UnloadAllResolvers() to auto-detect every eligible resolver",
			)
		}
		// Mirrors chalkpy's encode_unload_resolvers, which raises on the same combination.
		// "*" stands for every eligible resolver, so no single join key can apply to it.
		if spec.Fqn == UnloadAllResolversFqn && len(spec.PartitionBy) > 0 {
			return nil, errors.Newf(
				"unload resolver selector %q does not accept partition expressions - "+
					"name the resolver explicitly to partition its output", UnloadAllResolversFqn,
			)
		}

		partitions := make([]string, 0, len(spec.PartitionBy))
		for _, partition := range spec.PartitionBy {
			rendered, err := partition.serialize()
			if err != nil {
				return nil, errors.Wrapf(err, "serializing partition for unload resolver %q", spec.Fqn)
			}
			partitions = append(partitions, rendered)
		}

		serialized = append(serialized, internal.UnloadResolverSpecSerialized{
			Fqn:         spec.Fqn,
			PartitionBy: partitions,
		})
	}
	return serialized, nil
}
