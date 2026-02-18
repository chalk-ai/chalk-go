package internal

type OnlineQueryContext struct {
	Tags                 []string `json:"tags"`
	RequiredResolverTags []string `json:"required_resolver_tags"`
}

type FeatureEncodingOptions struct {
	EncodeStructsAsObjects bool `json:"encode_structs_as_objects"`
}

type OnlineQueryRequestSerialized struct {
	Inputs           map[string]any         `json:"inputs"`
	Outputs          []string               `json:"outputs"`
	Context          OnlineQueryContext     `json:"context"`
	Staleness        map[string]string      `json:"staleness"`
	IncludeMeta      bool                   `json:"include_meta"`
	QueryName        *string                `json:"query_name"`
	CorrelationId    *string                `json:"correlation_id"`
	QueryContext     *map[string]any        `json:"query_context"`
	Meta             map[string]string      `json:"meta"`
	QueryNameVersion *string                `json:"query_name_version"`
	Now              *string                `json:"now"`
	Explain          bool                   `json:"explain"`
	StorePlanStages  bool                   `json:"store_plan_stages"`
	EncodingOptions  FeatureEncodingOptions `json:"encoding_options"`
	PlannerOptions   map[string]any         `json:"planner_options"`
	BranchId         *string                `json:"branch_id"`
}

type OfflineQueryInputSerialized struct {
	Columns []string `json:"columns"`
	Values  [][]any  `json:"values"`
}

type OfflineQueryInputUri struct {
	ParquetUri              string            `json:"parquet_uri"`
	StartRow                *int              `json:"start_row,omitempty"`
	EndRow                  *int              `json:"end_row,omitempty"`
	IsIceberg               bool              `json:"is_iceberg,omitempty"`
	IcebergSnapshotId       *int              `json:"iceberg_snapshot_id,omitempty"`
	IcebergStartPartition   *int              `json:"iceberg_start_partition,omitempty"`
	IcebergEndPartition     *int              `json:"iceberg_end_partition,omitempty"`
	IcebergFilter           *string           `json:"iceberg_filter,omitempty"`
	AwsRoleArn              *string           `json:"aws_role_arn,omitempty"`
	AwsRegion               *string           `json:"aws_region,omitempty"`
	ColumnNameToFeatureName map[string]string `json:"column_name_to_feature_name,omitempty"`
}

type ResourceRequestsSerialized struct {
	CPU                 *string `json:"cpu"`
	Memory              *string `json:"memory"`
	EphemeralVolumeSize *string `json:"ephemeral_volume_size"`
	EphemeralStorage    *string `json:"ephemeral_storage"`
	ResourceGroup       *string `json:"resource_group"`
}

type OfflineQueryRequestSerialized struct {
	// Core fields
	Input                      any                         `json:"input"`
	Output                     []string                    `json:"output"`
	OutputExpressions          []string                    `json:"output_expressions"`
	RequiredOutput             []string                    `json:"required_output"`
	RequiredOutputExpressions  []string                    `json:"required_output_expressions"`
	DestinationFormat          string                      `json:"destination_format"`
	JobId                      *string                     `json:"job_id"`
	MaxSamples                 *int                        `json:"max_samples"`
	MaxCacheAge                *int                        `json:"max_cache_age_secs"`
	ObservedAtLowerBound       *string                     `json:"observed_at_lower_bound"`
	ObservedAtUpperBound       *string                     `json:"observed_at_upper_bound"`
	DatasetName                *string                     `json:"dataset_name"`
	Branch                     *string                     `json:"branch"`
	RecomputeFeatures          any                         `json:"recompute_features"`
	SampleFeatures             *[]string                   `json:"sample_features"`
	StorePlanStages            bool                        `json:"store_plan_stages"`
	Explain                    bool                        `json:"explain"`
	Tags                       *[]string                   `json:"tags"`
	RequiredResolverTags       *[]string                   `json:"required_resolver_tags"`
	CorrelationId              *string                     `json:"correlation_id"`
	QueryContext               *map[string]any             `json:"query_context"`
	PlannerOptions             *map[string]any             `json:"planner_options"`
	UseMultipleComputers       bool                        `json:"use_multiple_computers"`
	SpineSqlQuery              *string                     `json:"spine_sql_query"`
	RecomputeRequestRevisionId *string                     `json:"recompute_request_revision_id"`
	Resources                  *ResourceRequestsSerialized `json:"resources"`
	EnvOverrides               *map[string]string          `json:"env_overrides"`
	OverrideTargetImageTag     *string                     `json:"override_target_image_tag"`
	EnableProfiling            bool                        `json:"enable_profiling"`
	StoreOnline                bool                        `json:"store_online"`
	StoreOffline               bool                        `json:"store_offline"`
	NumShards                  *int                        `json:"num_shards"`
	NumWorkers                 *int                        `json:"num_workers"`
	FeatureForLowerUpperBound  *string                     `json:"feature_for_lower_upper_bound"`
	CompletionDeadline         *string                     `json:"completion_deadline"`
	MaxRetries                 *int                        `json:"max_retries"`
	UseJobQueue                bool                        `json:"use_job_queue"`
	OverlayGraph               *string                     `json:"overlay_graph"`
}
