package chalk

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/ipc"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
)

func (p OnlineQueryParamsComplete) toBytes() (*[]byte, error) {
	inputsAsArrow, err := p.inputsToArrow(p.underlying.inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inputs to Arrow: %w", err)
	}
	inputBytes, err := internal.CreateRequestBody(inputsAsArrow, p.underlying.outputs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inputs to Arrow: %w", err)
	}
	return inputBytes, nil
}

func (p OnlineQueryParamsComplete) inputsToArrow(inputs map[string]any) (arrow.Record, error) {
	golangToArrowType := map[reflect.Kind]arrow.DataType{
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
		reflect.String:  arrow.BinaryTypes.String,
		reflect.Bool:    arrow.FixedWidthTypes.Boolean,
	}

	for k, v := range inputs {
		if reflect.TypeOf(v).Kind() != reflect.Array && reflect.TypeOf(v).Kind() != reflect.Slice {
			return nil, fmt.Errorf("conversion of inputs to Arrow requires all input values to be an array, instead found type '%s' for feature '%s': ", reflect.TypeOf(v).Kind(), k)
		}
		length := 0
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			length = reflect.ValueOf(v).Len()
		} else {
			length = reflect.ValueOf(v).Type().Len()
		}
		if length == 0 {
			return nil, fmt.Errorf("conversion of inputs to Arrow requires all input values to be non-empty, instead found empty array for feature '%s': ", k)
		}
	}

	// Create the input values
	var schema []arrow.Field
	for k, v := range inputs {
		arrowType, ok := golangToArrowType[reflect.ValueOf(v).Type().Elem().Kind()]
		if !ok {
			return nil, fmt.Errorf("unsupported input type found for feature '%s' when converting to arrow: %s", k, reflect.TypeOf(v).Kind())
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
		reflectKind := reflect.ValueOf(values).Type().Elem().Kind()
		switch reflectKind {
		case reflect.Int:
			arrayValues := values.([]int)
			convertedValues := []int64{}
			for _, value := range arrayValues {
				convertedValues = append(convertedValues, int64(value))
			}
			recordBuilder.Field(idx).(*array.Int64Builder).AppendValues(convertedValues, nil)
		case reflect.Int8:
			recordBuilder.Field(idx).(*array.Int8Builder).AppendValues(values.([]int8), nil)
		case reflect.Int16:
			recordBuilder.Field(idx).(*array.Int16Builder).AppendValues(values.([]int16), nil)
		case reflect.Int32:
			recordBuilder.Field(idx).(*array.Int32Builder).AppendValues(values.([]int32), nil)
		case reflect.Int64:
			recordBuilder.Field(idx).(*array.Int64Builder).AppendValues(values.([]int64), nil)
		case reflect.Uint:
			arrayValues := values.([]uint)
			convertedValues := []uint64{}
			for _, value := range arrayValues {
				convertedValues = append(convertedValues, uint64(value))
			}
			recordBuilder.Field(idx).(*array.Uint64Builder).AppendValues(convertedValues, nil)
		case reflect.Uint8:
			recordBuilder.Field(idx).(*array.Uint8Builder).AppendValues(values.([]uint8), nil)
		case reflect.Uint16:
			recordBuilder.Field(idx).(*array.Uint16Builder).AppendValues(values.([]uint16), nil)
		case reflect.Uint32:
			recordBuilder.Field(idx).(*array.Uint32Builder).AppendValues(values.([]uint32), nil)
		case reflect.Uint64:
			recordBuilder.Field(idx).(*array.Uint64Builder).AppendValues(values.([]uint64), nil)
		case reflect.Float32:
			recordBuilder.Field(idx).(*array.Float32Builder).AppendValues(values.([]float32), nil)
		case reflect.Float64:
			recordBuilder.Field(idx).(*array.Float64Builder).AppendValues(values.([]float64), nil)
		case reflect.String:
			recordBuilder.Field(idx).(*array.StringBuilder).AppendValues(values.([]string), nil)
		case reflect.Bool:
			recordBuilder.Field(idx).(*array.BooleanBuilder).AppendValues(values.([]bool), nil)
		default:
			if reflectKind.String() == "time.Time" {
				// TODO: Support this
				//recordBuilder.Field(idx).(*array.TimestampBuilder).AppendValues(values.([]time.Time), nil)
			}
			return nil, fmt.Errorf("unsupported input type found for feature '%s' when converting to arrow: %s", field.Name, reflectKind.String())
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

func consumeMagicStr(startIdx int, bytes []byte) (int, error) {
	magicBytes := []byte("CHALK_BYTE_TRANSMISSION")
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
		fmt.Errorf("failed to process byte items length map")
	}
	fmt.Println(">>> BYTE ITEMS MAP CAST")
	fmt.Println(byteItemsMapCast)
	fmt.Println(">>> REAL SOHAI")
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

func ChalkUnmarshal(body []byte) (map[string]any, error) {
	idx, err := consumeMagicStr(0, body)
	if err != nil {
		return nil, err
	}
	idx, err, jsonBody := consumeJsonAttrs(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume json attrs: %w", err)
	}
	fmt.Println(">>> THE JSON ATTRIBUTES")
	fmt.Println(jsonBody)

	idx, err, pydanticJsonBody := consumePydanticAttrs(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume pydantic attrs: %w", err)
	}
	fmt.Println(">>> THE PYDANTIC ATTRIBUTES")
	fmt.Println(pydanticJsonBody)

	idx, err, byteItems := consumeByteItems(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume raw byte items: %w", err)
	}
	fmt.Println(">>> THE BYTE ITEMS")
	fmt.Println(byteItems)

	idx, err, deserializableByteItems := consumeByteItems(idx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to consume deserializable byte items: %w", err)
	}
	fmt.Println(">>> THE DESERIALIZABLE BYTE ITEMS")
	fmt.Println(deserializableByteItems)

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

func convertBytesToRecord(byteArr []byte) (*arrow.Record, error) {
	result := bytes.NewBuffer(byteArr)
	//buffer := bufio.NewReader(byteArr)
	reader, err := ipc.NewReader(result)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader unmarshaling result table: %w", err)
	}
	record, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal result table: %w", err)
	}
	return &record, nil
}

func (r *OnlineQueryBulkResponse) Unmarshal(body []byte) error {
	res, err := ChalkUnmarshal(body)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bytes: %w", err)
	}
	resultMap := map[QueryName]onlineQueryResultFeather{}

	queryNameToBytesInBytes, ok := res["query_results_bytes"]
	if !ok {
		return fmt.Errorf("malformed bulk online query response - missing 'query_results_bytes' attribute")
	}

	queryNameToBytesInBytesCast, ok := queryNameToBytesInBytes.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert 'query_results_bytes' value to a byte array")
	}
	queryNameToBytesInMap, err := ChalkUnmarshal(queryNameToBytesInBytesCast)

	for queryName, queryResultBytes := range queryNameToBytesInMap {
		queryResultBytesCast, ok := queryResultBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast bytes to byte array for query name: %s", queryName)
		}
		resultFeather := onlineQueryResultFeather{}
		err := resultFeather.Unmarshal(queryResultBytesCast)
		if err != nil {
			return fmt.Errorf("failed to unmarshal result bytes for query name '%s': %w", queryName, err)
		}
		resultMap[queryName] = resultFeather
	}

	r.QueryResults = resultMap
	fmt.Println(">>> HELLO")
	fmt.Println(res)
	return nil
}

func (r *onlineQueryResultFeather) UnmarshalInto(resultHolder any) error {
	// Scalar data
	// Loop through each row
	// 	Call UnmarshalInto for each value using its column name
	//

	// Groups data
}

func (r *onlineQueryResultFeather) Unmarshal(body []byte) error {
	res, err := ChalkUnmarshal(body)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bytes: %w", err)
	}

	hasData, ok := res["has_data"]
	if !ok {
		return fmt.Errorf("missing attribute 'has_data'")
	}
	hasDataBool, ok := hasData.(bool)
	if !ok {
		return fmt.Errorf("cannot cast attribute 'has_data' to bool")
	}

	var record *arrow.Record
	var groupsRecords map[string]*arrow.Record
	if hasDataBool {
		scalarDataBytes, ok := res["scalar_data"]
		if !ok {
			return fmt.Errorf("missing attribute 'scalar_data'")
		}
		scalarDataBytesCast, ok := scalarDataBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast scalar data bytes to bytes array")
		}
		record, err = convertBytesToRecord(scalarDataBytesCast)
		if err != nil {
			return fmt.Errorf("failed to convert scalar data bytes to an Arrow IPC Record: %w", err)
		}

		groupsDataBytes, ok := res["groups_data"]
		if !ok {
			return fmt.Errorf("missing attribute 'groups_data'")
		}
		groupsDataBytesCast, ok := groupsDataBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast groups data bytes to bytes array")
		}
		groupsDataMap, err := ChalkUnmarshal(groupsDataBytesCast)
		if err != nil {
			return fmt.Errorf("failed to unmarshal 'groups_data' value: %w", err)
		}

		groupsRecords = map[string]*arrow.Record{}
		for k, v := range groupsDataMap {
			vBytes, ok := v.([]byte)
			if !ok {
				return fmt.Errorf("failed to cast data for has-many feature '%s': %w", k, err)
			}
			vRecord, err := convertBytesToRecord(vBytes)
			if err != nil {
				return fmt.Errorf("failed to convert bytes for has-many feature '%s' to Arrow record batch: %w", k, err)
			}
			groupsRecords[k] = vRecord
		}
	}

	r.HasData = hasDataBool
	r.ScalarData = record
	r.GroupsData = groupsRecords

	metaString, ok := res["meta"]
	if !ok {
		return fmt.Errorf("missing attribute 'meta'")
	}
	if metaString == nil {

	} else {
		metaStringCast, ok := metaString.(string)
		if !ok {
			return fmt.Errorf("'meta' attribute cannot be casted to string")
		}
		meta := QueryMeta{}
		json.Unmarshal([]byte(metaStringCast), &meta)
		r.Meta = &meta
	}

	return nil
}
