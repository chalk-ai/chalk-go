package internal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/ipc"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

var golangToArrowPrimitiveType = map[reflect.Kind]arrow.DataType{
	reflect.Int:     arrow.PrimitiveTypes.Int64,
	reflect.Int8:    arrow.PrimitiveTypes.Int8,
	reflect.Int16:   arrow.PrimitiveTypes.Int16,
	reflect.Int32:   arrow.PrimitiveTypes.Int32,
	reflect.Int64:   arrow.PrimitiveTypes.Int64,
	reflect.Uint:    arrow.PrimitiveTypes.Uint64,
	reflect.Uint8:   arrow.PrimitiveTypes.Uint8,
	reflect.Uint16:  arrow.PrimitiveTypes.Uint16,
	reflect.Uint32:  arrow.PrimitiveTypes.Uint32,
	reflect.Uint64:  arrow.PrimitiveTypes.Uint64,
	reflect.Float32: arrow.PrimitiveTypes.Float32,
	reflect.Float64: arrow.PrimitiveTypes.Float64,
	reflect.String:  arrow.BinaryTypes.LargeString,
	reflect.Bool:    arrow.FixedWidthTypes.Boolean,
}

// inputsToArrowBytes converts map of FQNs to slice of values to an Arrow Record, serialized.
func inputsToArrowBytes(inputs map[string]any) ([]byte, error) {
	record, recordErr := ColumnMapToRecord(inputs)
	if recordErr != nil {
		return nil, recordErr
	}
	defer record.Release()
	return recordToBytes(record)
}

func convertReflectToArrowType(value reflect.Type) (arrow.DataType, error) {
	kind := value.Kind()
	if arrowType, isPrimitive := golangToArrowPrimitiveType[kind]; isPrimitive {
		return arrowType, nil
	} else if kind == reflect.Slice || kind == reflect.Array {
		elemKind := value.Elem()
		if elemType, err := convertReflectToArrowType(elemKind); err == nil {
			return arrow.LargeListOf(elemType), nil
		} else {
			return nil, fmt.Errorf("arrow conversion failed - a slice of anything "+
				"but primitives is currently unsupported, found type: %s",
				elemKind,
			)
		}
	} else if kind == reflect.Struct {
		if value == reflect.TypeOf(time.Time{}) {
			return &arrow.TimestampType{
				Unit:     arrow.Nanosecond,
				TimeZone: "UTC",
			}, nil
		}
		var arrowFields []arrow.Field
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			isPointer := field.Type.Kind() == reflect.Ptr
			rType := field.Type
			if isPointer {
				rType = field.Type.Elem()
			}
			dtype, dtypeErr := convertReflectToArrowType(rType)
			if dtypeErr != nil {
				return nil, errors.Wrapf(
					dtypeErr,
					"arrow conversion failed - failed to convert struct field type '%v' to arrow type",
					rType,
				)
			}
			arrowFields = append(arrowFields, arrow.Field{
				Name:     resolveFeatureName(field),
				Type:     dtype,
				Nullable: isPointer,
			})
		}
		return arrow.StructOf(arrowFields...), nil
	} else {
		return nil, fmt.Errorf("arrow conversion failed - unsupported type: %s", kind)
	}
}

func setBuilderValues(builder array.Builder, slice reflect.Value, nullMask []bool) error {
	sliceType := slice.Type()
	if sliceType.Kind() != reflect.Slice {
		return errors.Errorf(
			"conversion of inputs to Arrow requires all input values "+
				"to be a slice, instead found type '%s'",
			slice.Kind(),
		)
	}
	if slice.Len() == 0 {
		return nil // No values to append
	}

	elemType := sliceType.Elem()
	elemKind := elemType.Kind()
	values := slice.Interface()
	switch elemKind {
	case reflect.Ptr:
		// elemType is `*T`, so we need to get the type `T`
		nonPtrSliceType := reflect.SliceOf(elemType.Elem())
		nonPtrSlice := reflect.MakeSlice(nonPtrSliceType, 0, slice.Len())
		nilMask := make([]bool, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			v := slice.Index(i)
			if v.IsNil() {
				nilMask[i] = true
				nonPtrSlice = reflect.Append(nonPtrSlice, reflect.Zero(elemType.Elem()))
			} else {
				nonPtrSlice = reflect.Append(nonPtrSlice, v.Elem())
			}
		}
		return setBuilderValues(builder, nonPtrSlice, nilMask)
	case reflect.Int:
		arrayValues := values.([]int)
		convertedValues := []int64{}
		for _, value := range arrayValues {
			convertedValues = append(convertedValues, int64(value))
		}
		builder.(*array.Int64Builder).AppendValues(convertedValues, nullMask)
	case reflect.Int8:
		builder.(*array.Int8Builder).AppendValues(values.([]int8), nullMask)
	case reflect.Int16:
		builder.(*array.Int16Builder).AppendValues(values.([]int16), nullMask)
	case reflect.Int32:
		builder.(*array.Int32Builder).AppendValues(values.([]int32), nullMask)
	case reflect.Int64:
		builder.(*array.Int64Builder).AppendValues(values.([]int64), nullMask)
	case reflect.Uint:
		arrayValues := values.([]uint)
		convertedValues := []uint64{}
		for _, value := range arrayValues {
			convertedValues = append(convertedValues, uint64(value))
		}
		builder.(*array.Uint64Builder).AppendValues(convertedValues, nullMask)
	case reflect.Uint8:
		builder.(*array.Uint8Builder).AppendValues(values.([]uint8), nullMask)
	case reflect.Uint16:
		builder.(*array.Uint16Builder).AppendValues(values.([]uint16), nullMask)
	case reflect.Uint32:
		builder.(*array.Uint32Builder).AppendValues(values.([]uint32), nullMask)
	case reflect.Uint64:
		builder.(*array.Uint64Builder).AppendValues(values.([]uint64), nullMask)
	case reflect.Float32:
		builder.(*array.Float32Builder).AppendValues(values.([]float32), nullMask)
	case reflect.Float64:
		builder.(*array.Float64Builder).AppendValues(values.([]float64), nullMask)
	case reflect.String:
		builder.(*array.LargeStringBuilder).AppendValues(values.([]string), nullMask)
	case reflect.Bool:
		builder.(*array.BooleanBuilder).AppendValues(values.([]bool), nullMask)
	case reflect.Struct:
		if elemType == reflect.TypeOf(time.Time{}) {
			timeSlice := values.([]time.Time)
			timestampSlice := make([]arrow.Timestamp, 0, len(timeSlice))
			for _, t := range timeSlice {
				timestampSlice = append(timestampSlice, arrow.Timestamp(t.UnixNano()))
			}
			builder.(*array.TimestampBuilder).AppendValues(timestampSlice, nullMask)
		} else {
			sBuilder, builderOk := builder.(*array.StructBuilder)
			if !builderOk {
				return errors.Errorf("internal error: expected struct builder, found %T", builder)
			}

			numFieldsReflect := elemType.NumField()
			numFieldsArrow := sBuilder.NumField()
			if numFieldsReflect != numFieldsArrow {
				return errors.Errorf(
					"expected number of fields in struct to match number of fields in Arrow struct schema, "+
						"found %d fields in struct and %d fields in Arrow struct schema",
					numFieldsReflect,
					numFieldsArrow,
				)
			}

			var namesReflect []string
			var namesArrow []string
			arrowStructType, typeOk := sBuilder.Type().(*arrow.StructType)
			if !typeOk {
				return errors.Errorf(
					"internal error: expected struct type as StructBuilder type, found %T",
					sBuilder.Type(),
				)
			}
			for i := 0; i < numFieldsReflect; i++ {
				namesReflect = append(namesReflect, resolveFeatureName(elemType.Field(i)))
				namesArrow = append(namesArrow, arrowStructType.Field(i).Name)
			}
			if !reflect.DeepEqual(namesReflect, namesArrow) {
				return errors.Errorf(
					"expected field names and their ordering in struct to match field names in Arrow struct schema, "+
						"found %v in struct and %v in Arrow struct schema",
					namesReflect,
					namesArrow,
				)
			}

			var columns []reflect.Value
			for j := 0; j < numFieldsReflect; j++ {
				fieldSliceType := reflect.SliceOf(elemType.Field(j).Type)
				fieldSlice := reflect.MakeSlice(fieldSliceType, 0, slice.Len())
				for i := 0; i < slice.Len(); i++ {
					fieldSlice = reflect.Append(fieldSlice, slice.Index(i).Field(j))
				}
				columns = append(columns, fieldSlice)
			}

			for i := 0; i < sBuilder.NumField(); i++ {
				if err := setBuilderValues(sBuilder.FieldBuilder(i), columns[i], nullMask); err != nil {
					return errors.Wrapf(err, "failed to set values for struct field '%d'", i)
				}
			}
		}
	default:
		return errors.Errorf(
			"unsupported input type found when converting to arrow: %s",
			elemKind.String(),
		)
	}
	return nil
}

// ColumnMapToRecord converts a map of column names to slices of values to an Arrow Record.
func ColumnMapToRecord(inputs map[string]any) (arrow.Record, error) {
	// Create the input values
	var schema []arrow.Field
	for k, v := range inputs {
		columnVal := reflect.ValueOf(v)
		columnElemType := columnVal.Type().Elem()
		arrowType, convErr := convertReflectToArrowType(columnElemType)
		if convErr != nil {
			return nil, fmt.Errorf("failed to convert values for column '%s': %w", k, convErr)
		}
		schema = append(schema, arrow.Field{Name: k, Type: arrowType})
	}

	recordBuilder := array.NewRecordBuilder(memory.NewGoAllocator(), arrow.NewSchema(schema, nil))
	defer recordBuilder.Release()

	for idx, field := range schema {
		values, ok := inputs[field.Name]
		if !ok {
			return nil, fmt.Errorf("failed to find input values for feature '%s'", field.Name)
		}

		if err := setBuilderValues(recordBuilder.Field(idx), reflect.ValueOf(values), nil); err != nil {
			return nil, errors.Wrapf(err, "failed to set values for feature '%s'", field.Name)
		}
	}
	return recordBuilder.NewRecord(), nil
}

func consume8ByteLen(startIdx int, bytes []byte) (int, error, uint64) {
	numBytesThatRepresentsLength := 8
	err := checkLen(startIdx, bytes, numBytesThatRepresentsLength)
	if err != nil {
		return startIdx, err, 0
	}
	length := binary.BigEndian.Uint64(bytes[startIdx : startIdx+numBytesThatRepresentsLength])
	return startIdx + numBytesThatRepresentsLength, nil, length
}

func consumeMagicStr(magicStr string, startIdx int, bytes []byte) (int, error) {
	magicBytes := []byte(magicStr)
	err := checkLen(startIdx, bytes, len(magicBytes))
	if err != nil {
		return startIdx, fmt.Errorf("failed to find enough bytes to consume magic string")
	}
	ptr := startIdx
	for _, b := range magicBytes {
		if bytes[ptr] != b {
			return 0, fmt.Errorf("magic string bytes do not match")
		}
		ptr++
	}
	return ptr, nil
}

// consumeJsonAttrs converts bytes to a JSON struct
func consumeJsonAttrs(startIdx int, bytes []byte) (int, error, any) {
	midIdx, err, jsonBodyLen := consume8ByteLen(startIdx, bytes)
	if err != nil {
		return startIdx, fmt.Errorf("failed to consume length: %w", err), nil
	}
	var jsonBody any
	endIdx := midIdx + int(jsonBodyLen)
	err = json.Unmarshal(bytes[midIdx:endIdx], &jsonBody)
	if err != nil {
		return startIdx, fmt.Errorf("failed to unmarshal: %w", err), nil
	}
	return endIdx, nil, jsonBody
}

// consumePydanticAttrs converts bytes to JSON, but not to pydantic because
// pydantic is not supported in Go.
func consumePydanticAttrs(startIdx int, bytes []byte) (int, error, any) {
	midIdx, err, jsonBodyLen := consume8ByteLen(startIdx, bytes)
	if err != nil {
		return startIdx, fmt.Errorf("failed to consume length: %w", err), nil
	}
	var jsonBody any
	endIdx := midIdx + int(jsonBodyLen)
	err = json.Unmarshal(bytes[midIdx:endIdx], &jsonBody)
	if err != nil {
		return startIdx, fmt.Errorf("failed to unmarshal: %w", err), nil
	}
	// Loop through the struct and convert the values from a JSON string to
	// the appropriate type
	return endIdx, nil, jsonBody
}

func consumeByteItemsData(startIdx int, bytes []byte, byteItemsMap map[string]int) (int, error, map[string][]byte) {
	byteItems := map[string][]byte{}
	idx := startIdx
	for key, length := range byteItemsMap {
		err := checkLen(idx, bytes, length)
		if err != nil {
			return startIdx, fmt.Errorf("failed to find enough bytes to consume byte item with key '%s': %w", key, err), nil
		}
		byteItems[key] = bytes[idx : idx+int(length)]
		idx += int(length)
	}
	return idx, nil, byteItems
}

func consumeByteItems(startIdx int, bytes []byte) (int, error, map[string][]byte) {
	idx, err, byteItemsMap := consumeJsonAttrs(startIdx, bytes)
	if err != nil {
		return startIdx, fmt.Errorf("failed to consume byte items length map: %w", err), nil
	}
	byteItemsMapCast, ok := byteItemsMap.(map[string]any)
	if !ok {
		return idx, fmt.Errorf("failed to process byte items length map"), nil
	}
	byteItemsMapInt := map[string]int{}
	for key, val := range byteItemsMapCast {
		intVal, err := convertNumber[int](val)
		if err != nil {
			return idx, fmt.Errorf("failed to cast value in byte items length map to `int` - possibly a wider integer type needed"), nil
		}
		byteItemsMapInt[key] = intVal
	}

	idx, err, byteItemsData := consumeByteItemsData(idx, bytes, byteItemsMapInt)
	if err != nil {
		return startIdx, fmt.Errorf("failed to consume byte items data: %w", err), nil
	}
	return idx, nil, byteItemsData
}

// checkLen checks if there are enough bytes to consume
// so that we don't panic upon indexing out of bounds
func checkLen(startIdx int, bytes []byte, length int) error {
	if len(bytes) < startIdx+length {
		return fmt.Errorf("failed to find enough bytes to consume")
	}
	return nil
}

func produceLen(length int, ioWriter *bufio.Writer) error {
	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, uint64(length))
	_, err := ioWriter.Write(lengthBytes)
	if err != nil {
		return err
	}
	return nil
}

func produceJsonAttrs(jsonAttrs map[string]any, ioWriter *bufio.Writer) error {
	jsonBytes, err := json.Marshal(jsonAttrs)
	if err != nil {
		return fmt.Errorf("failed to serialize JSON attributes to JSON: %w", err)
	}
	err = produceLen(len(jsonBytes), ioWriter)
	_, err = ioWriter.Write(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

func produceByteAttrs(byteAttrs map[string][]byte, ioWriter *bufio.Writer) error {
	byteAttrsMap := map[string]any{}
	for k, v := range byteAttrs {
		byteAttrsMap[k] = len(v)
	}
	err := produceJsonAttrs(byteAttrsMap, ioWriter)
	if err != nil {
		return err
	}
	for _, v := range byteAttrs {
		_, err = ioWriter.Write(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func recordToBytes(record arrow.Record) ([]byte, error) {
	bws := &BufferWriteSeeker{}
	fileWriter, err := ipc.NewFileWriter(bws, ipc.WithSchema(record.Schema()), ipc.WithAllocator(memory.NewGoAllocator()))
	err = fileWriter.Write(record)
	if err != nil {
		return nil, fmt.Errorf("failed to write Arrow Table to request: %w", err)
	}
	err = fileWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close Arrow Table writer: %w", err)
	}

	return bws.Bytes(), nil
}

// ChalkMarshal converts a map to a byte array.
// Follows the byte-packing format as described in the Chalk
// Python repo's `byte_transmit.serialize()` function.
func ChalkMarshal(attrs map[string]any) ([]byte, error) {
	// Magic str
	// attrs json len
	// attrs json
	// pydantic len
	// pydantic json
	// attr and byte offset json len
	// attr and byte offset json
	// concatenated byte objects
	// attr and byte offset json len
	// attr and byte offsets for serializables
	// concatenated byte objects

	jsonAttrs := map[string]any{}
	byteAttrs := map[string][]byte{}

	for k, v := range attrs {
		if byteList, ok := v.([]byte); ok {
			byteAttrs[k] = byteList
		} else {
			jsonAttrs[k] = v
		}
	}

	// Serialize the message
	var result bytes.Buffer
	ioWriter := bufio.NewWriter(&result)

	// Magic string header
	_, err := ioWriter.WriteString("CHALK_BYTE_TRANSMISSION")
	if err != nil {
		return nil, err
	}

	// JSON attrs
	err = produceJsonAttrs(jsonAttrs, ioWriter)
	if err != nil {
		return nil, err
	}

	// Pydantic attrs
	err = produceJsonAttrs(map[string]any{}, ioWriter)
	if err != nil {
		return nil, err
	}

	// Byte attrs
	err = produceByteAttrs(byteAttrs, ioWriter)

	// ByteSerializables
	err = produceJsonAttrs(map[string]any{}, ioWriter)

	err = ioWriter.Flush()
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func CreateUploadFeaturesBody(inputs map[string]any) ([]byte, error) {
	recordBytes, err := inputsToArrowBytes(inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inputs to Arrow Record bytes: %w", err)
	}

	attrs := map[string]any{
		"features":          colls.Keys(inputs),
		"table_compression": "uncompressed",
		"table_bytes":       recordBytes,
	}

	return ChalkMarshal(attrs)
}

type FeatherRequestHeader struct {
	Outputs  []string `json:"outputs"`
	BranchId *string  `json:"branch_id"`
	Explain  bool     `json:"explain"`
}

func CreateOnlineQueryBulkBody(inputs map[string]any, header FeatherRequestHeader) ([]byte, error) {
	arrowBytes, err := inputsToArrowBytes(inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inputs to Arrow: %w", err)
	}

	// Serialize the message
	var result bytes.Buffer
	ioWriter := bufio.NewWriter(&result)

	// Magic string header
	magicStringLen, err := ioWriter.WriteString("chal1")
	if err != nil {
		return nil, err
	}

	// Placeholder for the size of the header
	placeholder := make([]byte, 8)
	_, err = ioWriter.Write(placeholder)
	if err != nil {
		return nil, err
	}

	// Header: TODO: get the other parameters for this.
	jsonBytes, err := json.Marshal(header)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize header to JSON: %w", err)
	}
	headerLength, err := ioWriter.Write(jsonBytes)
	if err != nil {
		return nil, err
	}

	// Placeholder for the size of the body
	_, err = ioWriter.Write(placeholder)
	if err != nil {
		return nil, err
	}

	// Body
	_, err = ioWriter.Write(arrowBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to write Arrow Table bytes to request: %w", err)
	}

	err = ioWriter.Flush()
	if err != nil {
		return nil, err
	}

	// Fill in the sizes
	resultBytes := result.Bytes()
	binary.BigEndian.PutUint64(resultBytes[magicStringLen:], uint64(headerLength))
	nonBodyLength := magicStringLen + 8 + headerLength + 8
	bodyLength := len(resultBytes) - nonBodyLength
	binary.BigEndian.PutUint64(resultBytes[magicStringLen+8+headerLength:], uint64(bodyLength))
	return resultBytes, nil
}

func GetHeaderFromSerializedOnlineQueryBulkBody(body []byte) (map[string]any, error) {
	idx, err := consumeMagicStr("chal1", 0, body)
	if err != nil {
		return nil, err
	}

	idx, err, header := consumeJsonAttrs(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume header: %w", err)
	}
	headerMap, ok := header.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("failed to cast header to map")
	}
	return headerMap, nil
}

func ChalkUnmarshal(body []byte) (map[string]any, error) {
	idx, err := consumeMagicStr("CHALK_BYTE_TRANSMISSION", 0, body)
	if err != nil {
		return nil, err
	}
	idx, err, jsonBody := consumeJsonAttrs(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume json attrs: %w", err)
	}

	idx, err, pydanticJsonBody := consumePydanticAttrs(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume pydantic attrs: %w", err)
	}

	idx, err, byteItems := consumeByteItems(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume raw byte items: %w", err)
	}

	idx, err, deserializableByteItems := consumeByteItems(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume deserializable byte items: %w", err)
	}

	res := map[string]any{}
	if jsonBody != nil {
		for k, v := range jsonBody.(map[string]any) {
			res[k] = v
		}
	}
	if pydanticJsonBody != nil {
		for k, v := range pydanticJsonBody.(map[string]any) {
			res[k] = v
		}
	}
	if byteItems != nil {
		for k, v := range byteItems {
			res[k] = v
		}
	}
	if deserializableByteItems != nil {
		for k, v := range deserializableByteItems {
			res[k] = v
		}
	}
	return res, nil
}

func ConvertBytesToTable(byteArr []byte) (result arrow.Table, err error) {
	bytesReader := bytes.NewReader(byteArr)
	alloc := memory.NewCheckedAllocator(memory.DefaultAllocator)
	fileReader, err := ipc.NewFileReader(bytesReader, ipc.WithAllocator(alloc))
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow file reader: %w", err)
	}
	defer func() {
		err = fileReader.Close()
	}()

	records := make([]arrow.Record, fileReader.NumRecords())
	for i := 0; i < fileReader.NumRecords(); i++ {
		rec, err := fileReader.Record(i)
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}
		rec.Retain()
		records[i] = rec
	}
	result = array.NewTableFromRecords(fileReader.Schema(), records)
	for _, rec := range records {
		rec.Release()
	}
	return result, err
}
