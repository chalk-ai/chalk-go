package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
)

// CreateFeatureTable creates an Arrow table from a map of feature to values.
// A feature can be either a string or a codegen'd feature reference.
// The values should be a slice of the appropriate type for the feature.
func CreateFeatureTable(featureToValues map[any]any) (arrow.Table, error) {
	return buildTableFromFeatureToValuesMap(featureToValues)
}

// buildTableFromFeatureToValuesMap builds an Arrow record from a map of features to values.
// The features should be codegen-ed `Feature` objects.
func buildTableFromFeatureToValuesMap(featureToValues map[any]any) (arrow.Table, error) {
	fqnToValues := make(map[string]any)
	for featureRaw, values := range featureToValues {
		feature, unwrapErr := UnwrapFeature(featureRaw)
		if unwrapErr != nil {
			return nil, unwrapErr
		}
		fqnToValues[feature.Fqn] = values
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
