package chalk

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"github.com/chalk-ai/chalk-go/internal"
	assert "github.com/stretchr/testify/require"
)

func TestShouldUploadOfflineQueryInputAsTableDefaultsToTrue(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	resolved := &offlineQueryParamsResolved{
		inputs: map[string][]TsFeatureValue{
			"user.id": {{Value: int64(1), ObservationTime: &now}},
		},
	}

	params := OfflineQueryParams{}
	assert.True(t, shouldUploadOfflineQueryInputAsTable(&params, resolved))

	params = OfflineQueryParams{}
	complete := OfflineQueryParamsComplete{underlying: params}.WithUploadInputAsTable(false)
	assert.False(t, shouldUploadOfflineQueryInputAsTable(&complete.underlying, resolved))

	fileInput := "s3://bucket/input.parquet"
	params = OfflineQueryParams{rawFileInput: &fileInput}
	assert.False(t, shouldUploadOfflineQueryInputAsTable(&params, resolved))
}

func TestOfflineQueryInputParquetPartitions(t *testing.T) {
	t.Parallel()

	ts := time.Date(2026, 6, 30, 12, 0, 0, 0, time.UTC)
	inputs := map[string][]TsFeatureValue{
		"user.id": {
			{Value: int64(101), ObservationTime: &ts},
			{Value: int64(102), ObservationTime: &ts},
			{Value: int64(103), ObservationTime: &ts},
		},
		"user.name": {
			{Value: "a", ObservationTime: &ts},
			{Value: "b", ObservationTime: &ts},
			{Value: "c", ObservationTime: &ts},
		},
	}

	numShards := 2
	partitions, err := offlineQueryInputParquetPartitions(inputs, &numShards, memory.DefaultAllocator)
	assert.NoError(t, err)
	assert.Len(t, partitions, 2)

	first := readParquetTable(t, partitions[0])
	defer first.Release()
	assert.Equal(t, int64(2), first.NumRows())
	assert.Equal(t, []string{"user.id", "user.name", "__ts__", "__index__"}, tableFieldNames(first))

	userID := first.Column(0).Data().Chunk(0).(*array.Int64)
	assert.Equal(t, int64(101), userID.Value(0))
	assert.Equal(t, int64(102), userID.Value(1))

	index := first.Column(3).Data().Chunk(0).(*array.Int64)
	assert.Equal(t, int64(0), index.Value(0))
	assert.Equal(t, int64(1), index.Value(1))

	second := readParquetTable(t, partitions[1])
	defer second.Release()
	assert.Equal(t, int64(1), second.NumRows())

	secondIndex := second.Column(3).Data().Chunk(0).(*array.Int64)
	assert.Equal(t, int64(2), secondIndex.Value(0))

	uploadedInput := &internal.OfflineQueryInputUploadedParquetSharded{
		Filenames: []string{"first.parquet", "second.parquet"},
		Version:   offlineQueryGivensVersionSingleTSColNameWithURIPrefix,
	}
	assert.Equal(t, 3, uploadedInput.Version)
}

func TestSerializeOfflineQueryParamsWithUploadedParquetInput(t *testing.T) {
	t.Parallel()

	uploadedInput := &internal.OfflineQueryInputUploadedParquetSharded{
		Filenames: []string{"first.parquet", "second.parquet"},
		Version:   offlineQueryGivensVersionSingleTSColNameWithURIPrefix,
	}
	resolved := &offlineQueryParamsResolved{
		outputs:         []string{"user.score"},
		requiredOutputs: []string{},
	}

	serialized, err := serializeOfflineQueryParamsWithInput(&OfflineQueryParams{}, resolved, uploadedInput)
	assert.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(serialized, &result)
	assert.NoError(t, err)

	input, ok := result["input"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, []any{"first.parquet", "second.parquet"}, input["filenames"])
	assert.Equal(t, float64(3), input["version"])
	assert.NotContains(t, input, "columns")
	assert.NotContains(t, input, "values")
	assert.NotContains(t, input, "parquet_uri")
}

func readParquetTable(t *testing.T, data []byte) arrow.Table {
	t.Helper()

	table, err := pqarrow.ReadTable(
		context.Background(),
		bytes.NewReader(data),
		nil,
		pqarrow.ArrowReadProperties{},
		memory.DefaultAllocator,
	)
	assert.NoError(t, err)
	return table
}

func tableFieldNames(table arrow.Table) []string {
	names := make([]string, table.Schema().NumFields())
	for i := range names {
		names[i] = table.Schema().Field(i).Name
	}
	return names
}
