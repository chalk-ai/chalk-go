package internal

import (
	"time"
)

type OnlineQueryContext struct {
	Environment          *string  `json:"environment"`
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
	IncludeMetrics   bool                   `json:"include_metrics"`
	DeploymentId     *string                `json:"deployment_id"`
	QueryName        *string                `json:"query_name"`
	CorrelationId    *string                `json:"correlation_id"`
	Meta             map[string]string      `json:"meta"`
	QueryNameVersion *string                `json:"query_name_version"`
	Now              *string                `json:"now"`
	Explain          bool                   `json:"explain"`
	StorePlanStages  bool                   `json:"store_plan_stages"`
	EncodingOptions  FeatureEncodingOptions `json:"encoding_options"`
}

type OfflineQueryInputSerialized struct {
	Columns []string `json:"columns"`
	Values  [][]any  `json:"values"`
}

type OfflineQueryRequestSerialized struct {
	Input                OfflineQueryInputSerialized `json:"input"`
	Output               []string                    `json:"output"`
	RequiredOutput       []string                    `json:"required_output"`
	DatasetName          *string                     `json:"dataset_name"`
	Branch               *string                     `json:"branch"`
	MaxSamples           *int                        `json:"max_samples"`
	DestinationFormat    string                      `json:"destination_format"`
	JobId                *string                     `json:"job_id"`
	MaxCacheAge          *int                        `json:"max_cache_age_secs"`
	ObservedAtLowerBound *time.Time                  `json:"observed_at_lower_bound"`
	ObservedAtUpperBound *time.Time                  `json:"observed_at_upper_bound"`
	Tags                 []string                    `json:"tags"`
}
