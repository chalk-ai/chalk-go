package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go/internal"
)

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
	record, recordErr := internal.ColumnMapToRecord(fqnToValues)
	if recordErr != nil {
		return nil, fmt.Errorf("error converting a map of column values to an Arrow Record: %w", recordErr)
	}
	defer record.Release()
	return array.NewTableFromRecords(record.Schema(), []arrow.Record{record}), nil
}
