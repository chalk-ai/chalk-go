package chalk

import (
	"encoding/json"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestUnmarshalDatasetResponse_WithDeltaTimestamp(t *testing.T) {
	t.Parallel()
	jsonData := []byte(`{
		"is_finished": false,
		"version": 6,
		"dataset_id": "test-dataset-id",
		"dataset_name": "test-dataset-name",
		"environment_id": "test-env",
		"revisions": [
			{
				"revision_id": "rev-1",
				"creator_id": "creator-1",
				"environment_id": "test-env",
				"outputs": ["feature1"],
				"status": 7,
				"filters": {
					"sample_filters": {
						"lower_bound": -2592000,
						"upper_bound": -86400,
						"max_samples": 1000
					},
					"max_cache_age_secs": null
				},
				"num_partitions": 1
			}
		],
		"errors": []
	}`)

	var dataset Dataset
	err := json.Unmarshal(jsonData, &dataset)
	assert.NoError(t, err)
	assert.NotNil(t, dataset.Revisions)
	assert.Len(t, dataset.Revisions, 1)

	sf := dataset.Revisions[0].Filters.SampleFilters
	assert.NotNil(t, sf.LowerBound)
	assert.NotNil(t, sf.LowerBound.Delta)
	assert.InDelta(t, -2592000, *sf.LowerBound.Delta, 0.01)
	assert.Nil(t, sf.LowerBound.Time)

	assert.NotNil(t, sf.UpperBound)
	assert.NotNil(t, sf.UpperBound.Delta)
	assert.InDelta(t, -86400, *sf.UpperBound.Delta, 0.01)
	assert.Nil(t, sf.UpperBound.Time)

	assert.NotNil(t, sf.MaxSamples)
	assert.Equal(t, 1000, *sf.MaxSamples)
}

func TestUnmarshalDatasetResponse_WithNonRFC3339Timestamp(t *testing.T) {
	jsonData := []byte(`{
		"is_finished": true,
		"version": 6,
		"dataset_id": "test-dataset-id",
		"dataset_name": "test-dataset-name",
		"environment_id": "test-env",
		"revisions": [
			{
				"revision_id": "rev-1",
				"creator_id": "creator-1",
				"environment_id": "test-env",
				"outputs": ["feature1"],
				"status": 7,
				"filters": {
					"sample_filters": {
						"lower_bound": "2025-09-27T10:00:00",
						"upper_bound": "2025-09-28T10:27:50.948512",
						"max_samples": 200
					},
					"max_cache_age_secs": null
				},
				"num_partitions": 1
			}
		],
		"errors": []
	}`)

	var dataset Dataset
	err := json.Unmarshal(jsonData, &dataset)
	assert.NoError(t, err)
	assert.NotNil(t, dataset.Revisions)
	assert.Len(t, dataset.Revisions, 1)

	sf := dataset.Revisions[0].Filters.SampleFilters
	assert.NotNil(t, sf.LowerBound)
	assert.NotNil(t, sf.LowerBound.Time)
	assert.Nil(t, sf.LowerBound.Delta)
	assert.Equal(t, 2025, sf.LowerBound.Time.Year())
	assert.Equal(t, 9, int(sf.LowerBound.Time.Month()))
	assert.Equal(t, 27, sf.LowerBound.Time.Day())
	assert.Equal(t, 10, sf.LowerBound.Time.Hour())
	assert.Equal(t, 0, sf.LowerBound.Time.Minute())

	assert.NotNil(t, sf.UpperBound)
	assert.NotNil(t, sf.UpperBound.Time)
	assert.Nil(t, sf.UpperBound.Delta)
	assert.Equal(t, 2025, sf.UpperBound.Time.Year())
	assert.Equal(t, 9, int(sf.UpperBound.Time.Month()))
	assert.Equal(t, 28, sf.UpperBound.Time.Day())
	assert.Equal(t, 10, sf.UpperBound.Time.Hour())
	assert.Equal(t, 27, sf.UpperBound.Time.Minute())
	assert.Equal(t, 50, sf.UpperBound.Time.Second())

	assert.NotNil(t, sf.MaxSamples)
	assert.Equal(t, 200, *sf.MaxSamples)
}

func TestMarshalDatasetResponse_WithDeltaTimestamp(t *testing.T) {
	lowerDelta := -2592000.0
	upperDelta := -86400.0
	maxSamples := 1000
	datasetId := "test-dataset-id"
	dataset := Dataset{
		Version:       6,
		DatasetId:     &datasetId,
		EnvironmentID: "test-env",
		Revisions: []DatasetRevision{
			{
				RevisionId:    "rev-1",
				EnvironmentID: "test-env",
				Status:        QueryStatusSuccessful,
				Filters: DatasetFilter{
					SampleFilters: DatasetSampleFilter{
						LowerBound: &TimeBound{Delta: &lowerDelta},
						UpperBound: &TimeBound{Delta: &upperDelta},
						MaxSamples: &maxSamples,
					},
				},
			},
		},
	}

	data, err := json.Marshal(dataset)
	assert.NoError(t, err)

	var roundTripped Dataset
	err = json.Unmarshal(data, &roundTripped)
	assert.NoError(t, err)

	sf := roundTripped.Revisions[0].Filters.SampleFilters
	assert.NotNil(t, sf.LowerBound)
	assert.NotNil(t, sf.LowerBound.Delta)
	assert.InDelta(t, -2592000, *sf.LowerBound.Delta, 0.01)
	assert.Nil(t, sf.LowerBound.Time)

	assert.NotNil(t, sf.UpperBound)
	assert.NotNil(t, sf.UpperBound.Delta)
	assert.InDelta(t, -86400, *sf.UpperBound.Delta, 0.01)
	assert.Nil(t, sf.UpperBound.Time)
}

func TestMarshalDatasetResponse_WithStringTimestamp(t *testing.T) {
	lower := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	upper := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
	maxSamples := 500
	datasetId := "test-dataset-id"
	dataset := Dataset{
		Version:       6,
		DatasetId:     &datasetId,
		EnvironmentID: "test-env",
		Revisions: []DatasetRevision{
			{
				RevisionId:    "rev-1",
				EnvironmentID: "test-env",
				Status:        QueryStatusSuccessful,
				Filters: DatasetFilter{
					SampleFilters: DatasetSampleFilter{
						LowerBound: &TimeBound{Time: &lower},
						UpperBound: &TimeBound{Time: &upper},
						MaxSamples: &maxSamples,
					},
				},
			},
		},
	}

	data, err := json.Marshal(dataset)
	assert.NoError(t, err)

	var roundTripped Dataset
	err = json.Unmarshal(data, &roundTripped)
	assert.NoError(t, err)

	sf := roundTripped.Revisions[0].Filters.SampleFilters
	assert.NotNil(t, sf.LowerBound)
	assert.NotNil(t, sf.LowerBound.Time)
	assert.Nil(t, sf.LowerBound.Delta)
	assert.True(t, lower.Equal(*sf.LowerBound.Time))

	assert.NotNil(t, sf.UpperBound)
	assert.NotNil(t, sf.UpperBound.Time)
	assert.Nil(t, sf.UpperBound.Delta)
	assert.True(t, upper.Equal(*sf.UpperBound.Time))
}
