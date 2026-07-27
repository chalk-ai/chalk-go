package chalk

import (
	"github.com/cockroachdb/errors"
)

// DeleteFeaturesParams targets feature observations for deletion.
//
// Features and Tags are mutually exclusive. Leaving both empty targets *every* feature in
// the namespace for the given primary keys, which is the widest possible deletion -- set
// one of them unless that is genuinely what you want.
type DeleteFeaturesParams struct {
	// Namespace in which the features targeted for deletion reside. Required.
	Namespace string

	// Features names the features to delete for the targeted primary keys. Names are
	// namespace-relative (e.g. "email", not "user.email"). Mutually exclusive with Tags.
	Features []string

	// Tags targets every feature carrying one of these tags. Mutually exclusive with Features.
	Tags []string

	// PrimaryKeys of the observations to delete. Required.
	PrimaryKeys []string

	// RetainOffline keeps the targeted observations in the offline store.
	RetainOffline bool

	// RetainOnline keeps the targeted observations in the online store.
	RetainOnline bool
}

// DeleteFeaturesResult holds any errors encountered while deleting. Deletion can
// partially succeed, so a non-empty Errors does not mean nothing was deleted.
type DeleteFeaturesResult struct {
	Errors ServerErrors `json:"errors"`
}

// deleteFeaturesRequest is the wire form, matching chalkpy's
// FeatureObservationDeletionRequest.
type deleteFeaturesRequest struct {
	Namespace     string   `json:"namespace"`
	Features      []string `json:"features"`
	Tags          []string `json:"tags"`
	PrimaryKeys   []string `json:"primary_keys"`
	RetainOffline bool     `json:"retain_offline"`
	RetainOnline  bool     `json:"retain_online"`
}

func (p DeleteFeaturesParams) validate() error {
	if p.Namespace == "" {
		return errors.New("namespace is required for feature deletion")
	}
	if len(p.PrimaryKeys) == 0 {
		return errors.New("at least one primary key is required for feature deletion")
	}
	// Mirrors chalkpy: the server picks targets from one selector or the other, so
	// supplying both is ambiguous rather than additive.
	if len(p.Features) > 0 && len(p.Tags) > 0 {
		return errors.New(
			"features and tags are mutually exclusive for feature deletion - " +
				"specify one, or neither to target every feature in the namespace",
		)
	}
	return nil
}
