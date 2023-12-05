package chalk

type UploadFeaturesParams struct {
	Inputs              map[any]any
	BranchOverride      string
	EnvironmentOverride string
	PreviewDeploymentId string
}
