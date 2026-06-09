package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestColumnMapToRecordOptionalPrimitives(t *testing.T) {
	t.Parallel()
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

func TestGetValueFromFixedSizeList(t *testing.T) {
	t.Parallel()
	// Vector / embedding features arrive over the wire as Arrow FixedSizeList
	// columns (with float32 child values). Make sure we can read a row out of one.
	builder := array.NewFixedSizeListBuilder(fixtures.TestAllocator, 3, arrow.PrimitiveTypes.Float32)
	defer builder.Release()
	valueBuilder := builder.ValueBuilder().(*array.Float32Builder)

	// Row 0: [1.5, 2.5, 3.5]
	builder.Append(true)
	valueBuilder.AppendValues([]float32{1.5, 2.5, 3.5}, nil)
	// Row 1: null (AppendNull appends the child nulls itself)
	builder.AppendNull()
	// Row 2: [4.5, 5.5, 6.5]
	builder.Append(true)
	valueBuilder.AppendValues([]float32{4.5, 5.5, 6.5}, nil)

	arr := builder.NewArray()
	defer arr.Release()

	row0, err := GetValueFromArrowArray(arr, 0, false)
	assert.Nil(t, err)
	assert.Equal(t, []any{float32(1.5), float32(2.5), float32(3.5)}, row0)

	row1, err := GetValueFromArrowArray(arr, 1, false)
	assert.Nil(t, err)
	assert.Nil(t, row1)

	row2, err := GetValueFromArrowArray(arr, 2, false)
	assert.Nil(t, err)
	assert.Equal(t, []any{float32(4.5), float32(5.5), float32(6.5)}, row2)
}
