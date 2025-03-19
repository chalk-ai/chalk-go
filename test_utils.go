package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"github.com/cockroachdb/errors"
)

type FeatureTable struct {
	Table arrow.Table

	allocator memory.Allocator
}

func (f *FeatureTable) ToBytes() ([]byte, error) {
	allocator := f.allocator
	if allocator == nil {
		allocator = memory.DefaultAllocator
	}
	return internal.TableToBytes(f.Table, allocator)
}

func (f *FeatureTable) Release() {
	f.Table.Release()
}

type MakeFeatureTableOptions struct {
	// Defaults to memory.DefaultAllocator
	Allocator memory.Allocator
}

// MakeFeatureTable creates an Arrow table from a map of feature to values.
// A feature can be either a string or a codegen'd feature reference.
// The values should be a slice of the appropriate type for the feature.
// The first return value is the Arrow table, and the second return value
// is the serialized Arrow table, which you can pass to raw response objects
// such as GRPCOnlineQueryBulkResult.
func MakeFeatureTable(featureToValues map[any]any, options ...MakeFeatureTableOptions) (*FeatureTable, error) {
	allocator := memory.DefaultAllocator
	if len(options) > 1 {
		return nil, errors.New("too many options provided")
	} else if len(options) == 1 {
		allocator = options[0].Allocator

	}
	fqnToValues, err := convertFeatureToValues(featureToValues)
	if err != nil {
		return nil, errors.Wrap(err, "mapping features to FQN")
	}

	record, recordErr := internal.ColumnMapToRecord(fqnToValues, allocator)
	if recordErr != nil {
		return nil, fmt.Errorf("error converting a map of column values to an Arrow Record: %w", recordErr)
	}
	defer record.Release()
	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})
	return &FeatureTable{Table: table}, nil
}

func convertFeatureToValues(featureToValues map[any]any) (map[string]any, error) {
	fqnToValues := make(map[string]any)
	for featureRaw, values := range featureToValues {
		if fqn, ok := featureRaw.(string); ok {
			fqnToValues[fqn] = values
		} else if feature, unwrapErr := UnwrapFeature(featureRaw); unwrapErr == nil {
			fqnToValues[feature.Fqn] = values
		} else {
			return nil, errors.Newf(
				"please pass either a string or a codegen'd feature reference as the key, got %T",
				featureRaw,
			)
		}
	}
	return fqnToValues, nil
}

// buildTableFromFeatureToValuesMap builds an Arrow record from a map of features to values.
// The features should be codegen-ed `Feature` objects.
func buildTableFromFeatureToValuesMap(featureToValues map[any]any) (arrow.Table, error) {
	fqnToValues, err := convertFeatureToValues(featureToValues)
	if err != nil {
		return nil, errors.Wrap(err, "mapping features to FQN")
	}
	return tableFromFqnToValues(fqnToValues)
}

func tableFromFqnToValues(fqnToValues map[string]any) (arrow.Table, error) {
	record, recordErr := internal.ColumnMapToRecord(fqnToValues, fixtures.TestAllocator)
	if recordErr != nil {
		return nil, fmt.Errorf("error converting a map of column values to an Arrow Record: %w", recordErr)
	}
	defer record.Release()
	return array.NewTableFromRecords(record.Schema(), []arrow.Record{record}), nil
}
