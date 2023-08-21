package internal

import (
	"bufio"
	"bytes"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/ipc"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"testing"
)

func TestSerialization(t *testing.T) {
	// Create the schema for the inputs table
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "feature.abc", Type: arrow.PrimitiveTypes.Int32},
			{Name: "feature.def", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Create the input values
	pool := memory.NewGoAllocator()
	recordBuilder := array.NewRecordBuilder(pool, schema)
	defer recordBuilder.Release()
	recordBuilder.Field(0).(*array.Int32Builder).AppendValues([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	recordBuilder.Field(1).(*array.Float64Builder).AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	record := recordBuilder.NewRecord()
	defer record.Release()

	// Serialize the message
	var result bytes.Buffer
	ioWriter := bufio.NewWriter(&result)

	ipcWriter := ipc.NewWriter(ioWriter, ipc.WithSchema(schema))

	// Magic string header
	magicStringLen, _ := ioWriter.WriteString("chal1")

	// Placeholder for the size of the header
	ioWriter.WriteByte(0)

	// Header
	headerLength, _ := ioWriter.WriteString(`{"output": [...]}`)

	// Placeholder for the size of the body
	ioWriter.WriteByte(0)

	// Body
	ipcWriter.Write(record)

	// Fill in the sizes
	ipcWriter.Close()
	ioWriter.Flush()
	resultBytes := result.Bytes()
	resultBytes[magicStringLen] = uint8(headerLength)
	resultBytes[magicStringLen+1+headerLength] = uint8(
		len(resultBytes) - magicStringLen - 1 - headerLength,
	)

}
