package internal

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestSerialization(t *testing.T) {
	// Create the input values
	recordBuilder := array.NewRecordBuilder(
		memory.NewGoAllocator(),
		arrow.NewSchema(
			[]arrow.Field{
				{Name: "feature.abc", Type: arrow.PrimitiveTypes.Int32},
				{Name: "feature.def", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		),
	)
	defer recordBuilder.Release()
	recordBuilder.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	recordBuilder.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	record := recordBuilder.NewRecord()
	defer record.Release()

	_, err := CreateRequestBody(record)
	assert.Nil(t, err)
}
