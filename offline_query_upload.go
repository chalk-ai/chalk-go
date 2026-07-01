package chalk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/compress"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
)

const (
	offlineQueryInputTSColumnName    = "__ts__"
	offlineQueryInputIndexColumnName = "__index__"
	offlineQueryInputUploadRowLimit  = 100

	// Matches chalkpy's OfflineQueryGivensVersion.SINGLE_TS_COL_NAME_WITH_URI_PREFIX.
	offlineQueryGivensVersionSingleTSColNameWithURIPrefix = 3
)

func normalizeOfflineQueryParallelism(p *OfflineQueryParams) error {
	if p.NumShards != nil && *p.NumShards < 1 {
		return fmt.Errorf("num_shards must be greater than 0")
	}
	if p.NumWorkers != nil && *p.NumWorkers < 1 {
		return fmt.Errorf("num_workers must be greater than 0")
	}
	if p.NumShards == nil && p.NumWorkers != nil {
		numShards := *p.NumWorkers
		p.NumShards = &numShards
	}
	if p.NumWorkers == nil && p.NumShards != nil {
		numWorkers := *p.NumShards
		p.NumWorkers = &numWorkers
	}
	if p.NumShards != nil && p.NumWorkers != nil && *p.NumWorkers > *p.NumShards {
		numWorkers := *p.NumShards
		p.NumWorkers = &numWorkers
	}
	return nil
}

func shouldUploadOfflineQueryInputAsTable(p *OfflineQueryParams, resolved *offlineQueryParamsResolved) bool {
	if p.rawFileInput != nil && *p.rawFileInput != "" {
		return false
	}
	if len(resolved.inputs) == 0 {
		return false
	}
	if p.NumShards != nil || p.NumWorkers != nil {
		return true
	}
	if offlineQueryInputRowCount(resolved.inputs) > offlineQueryInputUploadRowLimit {
		return true
	}
	if p.uploadInputAsTableSet != nil {
		return *p.uploadInputAsTableSet
	}
	return true
}

func (c *clientImpl) uploadOfflineQueryInputAsTable(
	ctx context.Context,
	p *OfflineQueryParams,
	resolved *offlineQueryParamsResolved,
) (*internal.OfflineQueryInputUploadedParquetSharded, error) {
	partitions, err := offlineQueryInputParquetPartitions(resolved.inputs, p.NumShards, c.allocator)
	if err != nil {
		return nil, err
	}

	var uploadURLResponse internal.OfflineQueryParquetUploadURLResponse
	if err := c.sendRequest(ctx, &sendRequestParams{
		Method:    http.MethodGet,
		URL:       fmt.Sprintf("v1/offline_query_parquet_upload_url/%d", len(partitions)),
		Response:  &uploadURLResponse,
		Versioned: resolved.versioned,
		Branch:    &p.Branch,
	}); err != nil {
		return nil, errors.Wrap(err, "getting parquet upload URLs")
	}
	if len(uploadURLResponse.URLs) != len(partitions) {
		return nil, fmt.Errorf(
			"received %d offline query parquet upload URLs for %d input partitions",
			len(uploadURLResponse.URLs),
			len(partitions),
		)
	}

	filenames := make([]string, len(uploadURLResponse.URLs))
	for i, annotatedURL := range uploadURLResponse.URLs {
		if err := c.uploadOfflineQueryParquetPartition(ctx, annotatedURL.SignedURL, partitions[i]); err != nil {
			return nil, errors.Wrapf(err, "uploading parquet partition %d", i)
		}
		filenames[i] = annotatedURL.Filename
	}

	return &internal.OfflineQueryInputUploadedParquetSharded{
		Filenames: filenames,
		Version:   offlineQueryGivensVersionSingleTSColNameWithURIPrefix,
	}, nil
}

func (c *clientImpl) uploadOfflineQueryParquetPartition(ctx context.Context, signedURL string, parquetBytes []byte) error {
	if strings.HasPrefix(signedURL, "file://") {
		parsed, err := url.Parse(signedURL)
		if err != nil {
			return errors.Wrap(err, "parsing file upload URL")
		}
		path := parsed.Path
		if path == "" {
			return fmt.Errorf("file upload URL has empty path")
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return errors.Wrap(err, "creating file upload directory")
		}
		return os.WriteFile(path, parquetBytes, 0o644)
	}

	ctx, cancel := internal.GetContextWithTimeout(ctx, c.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, signedURL, bytes.NewReader(parquetBytes))
	if err != nil {
		return errors.Wrap(err, "creating parquet upload request")
	}
	if strings.Contains(signedURL, ".blob.core.windows.net") {
		req.Header.Set("x-ms-blob-type", "BlockBlob")
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		if len(body) > 0 {
			return fmt.Errorf("parquet upload failed with status %s: %s", res.Status, string(body))
		}
		return fmt.Errorf("parquet upload failed with status %s", res.Status)
	}
	return nil
}

func offlineQueryInputParquetPartitions(
	inputs map[string][]TsFeatureValue,
	numShards *int,
	allocator memory.Allocator,
) ([][]byte, error) {
	if allocator == nil {
		allocator = memory.DefaultAllocator
	}
	rowCount := offlineQueryInputRowCount(inputs)
	if rowCount == 0 {
		return nil, fmt.Errorf("offline query input must contain at least one row")
	}

	numPartitions := 1
	if numShards != nil {
		numPartitions = *numShards
		if numPartitions > rowCount {
			return nil, fmt.Errorf(
				"num_shards (%d) is greater than the number of rows in the input data (%d)",
				numPartitions,
				rowCount,
			)
		}
	}

	partitions := make([][]byte, 0, numPartitions)
	for partitionIndex := 0; partitionIndex < numPartitions; partitionIndex++ {
		start, end := offlineQueryPartitionBounds(rowCount, numPartitions, partitionIndex)
		parquetBytes, err := offlineQueryInputPartitionToParquet(inputs, start, end, allocator)
		if err != nil {
			return nil, errors.Wrapf(err, "serializing input partition %d", partitionIndex)
		}
		partitions = append(partitions, parquetBytes)
	}
	return partitions, nil
}

func offlineQueryInputRowCount(inputs map[string][]TsFeatureValue) int {
	for _, values := range inputs {
		return len(values)
	}
	return 0
}

func offlineQueryPartitionBounds(rowCount int, numPartitions int, partitionIndex int) (int, int) {
	rowsPerPartition := (rowCount + numPartitions - 1) / numPartitions
	start := partitionIndex * rowsPerPartition
	end := start + rowsPerPartition
	if end > rowCount {
		end = rowCount
	}
	return start, end
}

func offlineQueryInputPartitionToParquet(
	inputs map[string][]TsFeatureValue,
	start int,
	end int,
	allocator memory.Allocator,
) ([]byte, error) {
	featureNames := make([]string, 0, len(inputs))
	for fqn := range inputs {
		featureNames = append(featureNames, fqn)
	}
	slices.Sort(featureNames)

	fields := make([]arrow.Field, 0, len(featureNames)+2)
	columns := make([]arrow.Array, 0, len(featureNames)+2)
	defer func() {
		for _, column := range columns {
			column.Release()
		}
	}()

	for _, fqn := range featureNames {
		fieldType, err := inferOfflineQueryArrowType(inputs[fqn][start:end])
		if err != nil {
			return nil, errors.Wrapf(err, "inferring Arrow type for input feature %q", fqn)
		}
		column, err := offlineQueryArrowArray(inputs[fqn][start:end], fieldType, allocator)
		if err != nil {
			return nil, errors.Wrapf(err, "building Arrow array for input feature %q", fqn)
		}
		fields = append(fields, arrow.Field{Name: fqn, Type: fieldType, Nullable: true})
		columns = append(columns, column)
	}

	tsColumn, err := offlineQueryTimestampColumn(inputs[featureNames[0]][start:end], allocator)
	if err != nil {
		return nil, err
	}
	fields = append(fields, arrow.Field{
		Name:     offlineQueryInputTSColumnName,
		Type:     &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"},
		Nullable: false,
	})
	columns = append(columns, tsColumn)

	indexColumn := offlineQueryIndexColumn(start, end, allocator)
	fields = append(fields, arrow.Field{Name: offlineQueryInputIndexColumnName, Type: arrow.PrimitiveTypes.Int64, Nullable: false})
	columns = append(columns, indexColumn)

	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, columns, int64(end-start))
	defer record.Release()

	table := array.NewTableFromRecords(schema, []arrow.Record{record})
	defer table.Release()

	var buf bytes.Buffer
	writer, err := pqarrow.NewFileWriter(
		table.Schema(),
		&buf,
		parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy),
		),
		pqarrow.NewArrowWriterProperties(pqarrow.WithStoreSchema()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating parquet writer")
	}
	if err := writer.WriteTable(table, table.NumRows()); err != nil {
		_ = writer.Close()
		return nil, errors.Wrap(err, "writing parquet table")
	}
	if err := writer.Close(); err != nil {
		return nil, errors.Wrap(err, "closing parquet writer")
	}

	return buf.Bytes(), nil
}

type offlineQueryArrowKind int

const (
	offlineQueryArrowKindUnknown offlineQueryArrowKind = iota
	offlineQueryArrowKindBool
	offlineQueryArrowKindInt64
	offlineQueryArrowKindFloat64
	offlineQueryArrowKindString
	offlineQueryArrowKindTimestamp
)

func inferOfflineQueryArrowType(values []TsFeatureValue) (arrow.DataType, error) {
	kind := offlineQueryArrowKindUnknown
	for _, value := range values {
		if value.Value == nil {
			continue
		}
		nextKind, err := offlineQueryArrowKindForValue(value.Value)
		if err != nil {
			return nil, err
		}
		if kind == offlineQueryArrowKindUnknown {
			kind = nextKind
			continue
		}
		if kind == nextKind {
			continue
		}
		if (kind == offlineQueryArrowKindInt64 && nextKind == offlineQueryArrowKindFloat64) ||
			(kind == offlineQueryArrowKindFloat64 && nextKind == offlineQueryArrowKindInt64) {
			kind = offlineQueryArrowKindFloat64
			continue
		}
		return nil, fmt.Errorf("mixed input value types %s and %s are not supported", kind, nextKind)
	}

	switch kind {
	case offlineQueryArrowKindBool:
		return arrow.FixedWidthTypes.Boolean, nil
	case offlineQueryArrowKindInt64:
		return arrow.PrimitiveTypes.Int64, nil
	case offlineQueryArrowKindFloat64:
		return arrow.PrimitiveTypes.Float64, nil
	case offlineQueryArrowKindString:
		return arrow.BinaryTypes.String, nil
	case offlineQueryArrowKindTimestamp:
		return &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}, nil
	default:
		return nil, fmt.Errorf("could not infer Arrow type from all-null input values")
	}
}

func offlineQueryArrowKindForValue(value any) (offlineQueryArrowKind, error) {
	switch value.(type) {
	case bool:
		return offlineQueryArrowKindBool, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return offlineQueryArrowKindInt64, nil
	case float32, float64:
		return offlineQueryArrowKindFloat64, nil
	case string:
		return offlineQueryArrowKindString, nil
	case time.Time, *time.Time:
		return offlineQueryArrowKindTimestamp, nil
	default:
		return offlineQueryArrowKindUnknown, fmt.Errorf("unsupported input value type %T", value)
	}
}

func (k offlineQueryArrowKind) String() string {
	switch k {
	case offlineQueryArrowKindBool:
		return "bool"
	case offlineQueryArrowKindInt64:
		return "int64"
	case offlineQueryArrowKindFloat64:
		return "float64"
	case offlineQueryArrowKindString:
		return "string"
	case offlineQueryArrowKindTimestamp:
		return "timestamp"
	default:
		return "unknown"
	}
}

func offlineQueryArrowArray(values []TsFeatureValue, dataType arrow.DataType, allocator memory.Allocator) (arrow.Array, error) {
	switch dataType.ID() {
	case arrow.BOOL:
		builder := array.NewBooleanBuilder(allocator)
		defer builder.Release()
		rawValues := make([]bool, len(values))
		valid := make([]bool, len(values))
		for i, value := range values {
			if value.Value == nil {
				continue
			}
			boolValue, ok := value.Value.(bool)
			if !ok {
				return nil, fmt.Errorf("expected bool value, got %T", value.Value)
			}
			rawValues[i] = boolValue
			valid[i] = true
		}
		builder.AppendValues(rawValues, valid)
		return builder.NewArray(), nil
	case arrow.INT64:
		builder := array.NewInt64Builder(allocator)
		defer builder.Release()
		rawValues := make([]int64, len(values))
		valid := make([]bool, len(values))
		for i, value := range values {
			if value.Value == nil {
				continue
			}
			intValue, err := offlineQueryValueAsInt64(value.Value)
			if err != nil {
				return nil, err
			}
			rawValues[i] = intValue
			valid[i] = true
		}
		builder.AppendValues(rawValues, valid)
		return builder.NewArray(), nil
	case arrow.FLOAT64:
		builder := array.NewFloat64Builder(allocator)
		defer builder.Release()
		rawValues := make([]float64, len(values))
		valid := make([]bool, len(values))
		for i, value := range values {
			if value.Value == nil {
				continue
			}
			floatValue, err := offlineQueryValueAsFloat64(value.Value)
			if err != nil {
				return nil, err
			}
			rawValues[i] = floatValue
			valid[i] = true
		}
		builder.AppendValues(rawValues, valid)
		return builder.NewArray(), nil
	case arrow.STRING:
		builder := array.NewStringBuilder(allocator)
		defer builder.Release()
		rawValues := make([]string, len(values))
		valid := make([]bool, len(values))
		for i, value := range values {
			if value.Value == nil {
				continue
			}
			stringValue, ok := value.Value.(string)
			if !ok {
				return nil, fmt.Errorf("expected string value, got %T", value.Value)
			}
			rawValues[i] = stringValue
			valid[i] = true
		}
		builder.AppendValues(rawValues, valid)
		return builder.NewArray(), nil
	case arrow.TIMESTAMP:
		builder := array.NewTimestampBuilder(allocator, dataType.(*arrow.TimestampType))
		defer builder.Release()
		rawValues := make([]arrow.Timestamp, len(values))
		valid := make([]bool, len(values))
		for i, value := range values {
			if value.Value == nil {
				continue
			}
			timestamp, err := offlineQueryValueAsTimestamp(value.Value)
			if err != nil {
				return nil, err
			}
			rawValues[i] = timestamp
			valid[i] = true
		}
		builder.AppendValues(rawValues, valid)
		return builder.NewArray(), nil
	default:
		return nil, fmt.Errorf("unsupported Arrow type %s", dataType)
	}
}

func offlineQueryValueAsInt64(value any) (int64, error) {
	const maxInt64AsUint64 = uint64(1<<63 - 1)
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		if uint64(v) > maxInt64AsUint64 {
			return 0, fmt.Errorf("uint value %d overflows int64", v)
		}
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > maxInt64AsUint64 {
			return 0, fmt.Errorf("uint64 value %d overflows int64", v)
		}
		return int64(v), nil
	default:
		return 0, fmt.Errorf("expected integer value, got %T", value)
	}
}

func offlineQueryValueAsFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		intValue, err := offlineQueryValueAsInt64(v)
		if err != nil {
			return 0, err
		}
		return float64(intValue), nil
	default:
		return 0, fmt.Errorf("expected numeric value, got %T", value)
	}
}

func offlineQueryValueAsTimestamp(value any) (arrow.Timestamp, error) {
	var timestamp time.Time
	switch v := value.(type) {
	case time.Time:
		timestamp = v
	case *time.Time:
		if v == nil {
			return 0, fmt.Errorf("timestamp pointer is nil")
		}
		timestamp = *v
	default:
		return 0, fmt.Errorf("expected time.Time value, got %T", value)
	}
	return arrow.TimestampFromTime(timestamp.UTC(), arrow.Microsecond)
}

func offlineQueryTimestampColumn(values []TsFeatureValue, allocator memory.Allocator) (arrow.Array, error) {
	timestampType := &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(allocator, timestampType)
	defer builder.Release()

	rawValues := make([]arrow.Timestamp, len(values))
	for i, value := range values {
		timestamp, err := arrow.TimestampFromTime(value.ObservationTime.UTC(), arrow.Microsecond)
		if err != nil {
			return nil, errors.Wrap(err, "converting observation time")
		}
		rawValues[i] = timestamp
	}
	builder.AppendValues(rawValues, nil)
	return builder.NewArray(), nil
}

func offlineQueryIndexColumn(start int, end int, allocator memory.Allocator) arrow.Array {
	builder := array.NewInt64Builder(allocator)
	defer builder.Release()

	rawValues := make([]int64, end-start)
	for i := range rawValues {
		rawValues[i] = int64(start + i)
	}
	builder.AppendValues(rawValues, nil)
	return builder.NewArray()
}
