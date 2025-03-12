package internal

import (
	"github.com/chalk-ai/chalk-go/v2/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestColumnMapToRecordOptionalPrimitives(t *testing.T) {
	// Tests that we can convert a slice of pointers to primitive
	// types to an Arrow column with nulls correctly set.
	num := int64(1)
	numFloat := 1.0
	in := map[string]any{
		"int":   []*int64{&num, nil, &num},
		"float": []*float64{&numFloat, &numFloat, nil},
	}
	record, convertErr := ColumnMapToRecord(in, fixtures.TestAllocator)
	assert.Nil(t, convertErr)
	for i := 0; i < int(record.NumRows()); i++ {
		for j, col := range record.Columns() {
			name := record.ColumnName(j)
			slice := in[name]
			if col.IsNull(i) {
				assert.True(t, reflect.ValueOf(slice).Index(i).IsNil())
			} else {
				val, valErr := GetValueFromArrowArray(col, i, false)
				assert.Nil(t, valErr)
				assert.Equal(t, reflect.ValueOf(slice).Index(i).Elem().Interface(), val)
			}
		}
	}
}
