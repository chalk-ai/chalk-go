package chalk

// DeleteFeaturesParams targets feature observations for deletion.
//
// Features and Tags are mutually exclusive. Leaving both nil targets *every* feature in
// the namespace for the given primary keys, which is the widest possible deletion. An
// explicitly empty, non-nil selector targets no features.
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
