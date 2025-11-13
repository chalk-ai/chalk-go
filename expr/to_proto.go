package expr

import (
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/ipc"
	"github.com/apache/arrow/go/v16/arrow/memory"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
	"github.com/chalk-ai/chalk-go/internal"
)

func ToIdentifierLiteral(name string) *expressionv1.LogicalExprNode {
	return &expressionv1.LogicalExprNode{
		ExprForm: &expressionv1.LogicalExprNode_Identifier{
			Identifier: &expressionv1.Identifier{
				Name: name,
			},
		},
	}
}

func toProtos(exprs ...Expr) ([]*expressionv1.LogicalExprNode, error) {
	protos := make([]*expressionv1.LogicalExprNode, len(exprs))
	for i, e := range exprs {
		proto, err := ToProto(e)
		if err != nil {
			return nil, err
		}
		protos[i] = proto
	}
	return protos, nil
}

// ToProto converts an ExprI to a LogicalExprNode proto message
func ToProto(expr ExprI) (*expressionv1.LogicalExprNode, error) {
	if expr == nil {
		return nil, nil
	}

	switch e := expr.(type) {
	case *IdentifierExpr:
		return ToIdentifierLiteral(e.Name), nil

	case *LiteralExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
				LiteralValue: &expressionv1.ExprLiteral{
					Value:               e.ScalarValue,
					IsArrowScalarObject: e.IsArrowScalarObject,
				},
			},
		}, nil

	case *ColumnExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Identifier{
				Identifier: &expressionv1.Identifier{
					Name: formatColumnName(e),
				},
			},
		}, nil

	case *GetAttributeExpr:
		parent, err := ToProto(e.Parent)
		if err != nil {
			return nil, err
		}
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
				GetAttribute: &expressionv1.ExprGetAttribute{
					Parent: parent,
					Attribute: &expressionv1.Identifier{
						Name: e.Attribute,
					},
				},
			},
		}, nil

	case *CallExpr:
		args, err := toProtos(e.Args...)
		if err != nil {
			return nil, err
		}

		kwargs := make(map[string]*expressionv1.LogicalExprNode)
		for key, value := range e.Kwargs {
			proto, err := ToProto(value)
			if err != nil {
				return nil, err
			}
			kwargs[key] = proto
		}

		// Special handling for list_literal function
		if id, ok := e.Function.(*IdentifierExpr); ok && id.Name == "list_literal" {
			if len(kwargs) != 0 {
				return nil, fmt.Errorf("list_literal cannot have kwargs")
			}

			if len(args) == 0 {
				return nil, fmt.Errorf("list_literal must have at least one argument")
			}

			scalarValues := make([]*arrowv1.ScalarValue, len(args))
			for i, arg := range args {
				litForm, ok := arg.GetExprForm().(*expressionv1.LogicalExprNode_LiteralValue)
				if !ok || litForm.LiteralValue == nil {
					return nil, fmt.Errorf("arguments to list_literal must be literals")
				}
				scalarValues[i] = litForm.LiteralValue.Value
			}

			listValue, err := scalarValuesToScalarListValue(scalarValues, memory.DefaultAllocator)
			if err != nil {
				return nil, fmt.Errorf("failed to create list value: %w", err)
			}

			return &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_ListValue{
								ListValue: listValue,
							},
						},
						IsArrowScalarObject: false,
					},
				},
			}, nil
		}

		fn, err := ToProto(e.Function)
		if err != nil {
			return nil, err
		}
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Call{
				Call: &expressionv1.ExprCall{
					Func:   fn,
					Args:   args,
					Kwargs: kwargs,
				},
			},
		}, nil

	case *AliasExpr:
		// Aliases are typically handled at a higher level in the query plan,
		// but we can represent them as the underlying expression for now
		return ToProto(e.Expression)

	case *DataFrameExprImpl:
		// DataFrame reference as identifier
		// return ToProto(Identifier("_").Attr(e.Name))
		return ToIdentifierLiteral(e.Name), nil

	case *AggregateExprImpl:
		// Build the base DataFrame reference
		dataframeNode, err := ToProto(e.DataFrame)
		if err != nil {
			return nil, err
		}

		// If there are filter conditions or selections, wrap the DataFrame in a GetSubscript
		if len(e.Conditions) > 0 || e.Selection != nil {
			subscriptNodes := make([]*expressionv1.LogicalExprNode, 0, len(e.Conditions)+1)

			// Convert selection to proto node if it exists
			if e.Selection != nil {
				proto, err := ToProto(e.Selection)
				if err != nil {
					return nil, err
				}
				subscriptNodes = append(subscriptNodes, proto)
			}

			// Convert all conditions to proto nodes
			for _, condition := range e.Conditions {
				proto, err := ToProto(condition)
				if err != nil {
					return nil, err
				}
				subscriptNodes = append(subscriptNodes, proto)
			}

			// Wrap the DataFrame in a GetSubscript with all filter conditions
			dataframeNode = &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_GetSubscript{
					GetSubscript: &expressionv1.ExprGetSubscript{
						Parent:    dataframeNode,
						Subscript: subscriptNodes,
					},
				},
			}
		}

		args, err := toProtos(e.Arguments...)
		if err != nil {
			return nil, err
		}
		kwargs := make(map[string]*expressionv1.LogicalExprNode, 1)

		expectedNumArgs := 0
		endsWithK := false
		switch e.Function {
		case "approx_top_k":
			expectedNumArgs = 1
			endsWithK = true
		case "min_by":
			expectedNumArgs = 1
		case "max_by":
			expectedNumArgs = 1
		case "min_by_n":
			expectedNumArgs = 2
		case "max_by_n":
			expectedNumArgs = 2
		}
		if len(e.Arguments) != expectedNumArgs {
			return nil, fmt.Errorf("expecting exactly %d arguments to %s, got %d", expectedNumArgs, e.Function, len(e.Arguments))
		}

		if endsWithK {
			kwargs["k"] = args[len(args)-1]
		}

		// Apply the aggregation function as a GetAttribute on the (possibly filtered) DataFrame
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Call{
				Call: &expressionv1.ExprCall{
					Func: &expressionv1.LogicalExprNode{
						ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
							GetAttribute: &expressionv1.ExprGetAttribute{
								Attribute: &expressionv1.Identifier{
									Name: e.Function,
								},
								Parent: dataframeNode,
							},
						},
					},
					Args:   args,
					Kwargs: kwargs,
				},
			},
		}, nil

	default:
		return nil, fmt.Errorf("unknown (%T)", e)
	}
}

// formatColumnName formats a column expression as a string
func formatColumnName(col *ColumnExpr) string {
	if col.Relation != "" {
		return col.Relation + "." + col.Name
	}
	return col.Name
}

// CreateEmptyList creates an empty list with the specified arrow type
func CreateEmptyList(dataType arrow.DataType, allocator memory.Allocator) (*arrowv1.ScalarListValue, error) {
	// Create an empty builder of the appropriate type
	var builder array.Builder

	switch dataType.(type) {
	case *arrow.Int64Type:
		builder = array.NewInt64Builder(allocator)
	case *arrow.Float64Type:
		builder = array.NewFloat64Builder(allocator)
	case *arrow.StringType:
		builder = array.NewStringBuilder(allocator)
	case *arrow.LargeStringType:
		builder = array.NewLargeStringBuilder(allocator)
	case *arrow.BooleanType:
		builder = array.NewBooleanBuilder(allocator)
	default:
		return nil, fmt.Errorf("unsupported arrow type for empty list: %s", dataType.Name())
	}
	defer builder.Release()

	// Build empty array
	arr := builder.NewArray()
	defer arr.Release()

	// Create a schema with a single field
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "item", Type: dataType, Nullable: true},
	}, nil)

	// Create a record with the empty array
	record := array.NewRecord(schema, []arrow.Array{arr}, 0)
	defer record.Release()

	// Serialize to Arrow IPC format
	bws := &internal.BufferWriteSeeker{}
	writer, err := ipc.NewFileWriter(bws, ipc.WithSchema(schema), ipc.WithAllocator(allocator))
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow IPC writer: %w", err)
	}

	if err := writer.Write(record); err != nil {
		return nil, fmt.Errorf("failed to write Arrow record: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close Arrow writer: %w", err)
	}

	// Convert Arrow schema to protobuf Schema
	protoSchema, err := arrowSchemaToProto(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to convert schema to proto: %w", err)
	}

	return &arrowv1.ScalarListValue{
		ArrowData: bws.Bytes(),
		Schema:    protoSchema,
	}, nil
}

// scalarValuesToScalarListValue converts a slice of ScalarValue to a ScalarListValue
func scalarValuesToScalarListValue(values []*arrowv1.ScalarValue, allocator memory.Allocator) (*arrowv1.ScalarListValue, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("cannot infer type for empty list")
	}

	// Determine the data type from the first value
	var dataType arrow.DataType
	var builder array.Builder

	// Check all values are the same type as the first non-null value
	firstValue := values[0]

	switch firstValue.Value.(type) {
	case *arrowv1.ScalarValue_Int64Value:
		dataType = arrow.PrimitiveTypes.Int64
		builder = array.NewInt64Builder(allocator)
		defer builder.Release()

		int64Builder := builder.(*array.Int64Builder)
		for _, val := range values {
			if v, ok := val.Value.(*arrowv1.ScalarValue_Int64Value); ok {
				int64Builder.Append(v.Int64Value)
			} else if _, ok := val.Value.(*arrowv1.ScalarValue_NullValue); ok {
				int64Builder.AppendNull()
			} else {
				return nil, fmt.Errorf("mixed types in list not supported")
			}
		}

	case *arrowv1.ScalarValue_Float64Value:
		dataType = arrow.PrimitiveTypes.Float64
		builder = array.NewFloat64Builder(allocator)
		defer builder.Release()

		float64Builder := builder.(*array.Float64Builder)
		for _, val := range values {
			if v, ok := val.Value.(*arrowv1.ScalarValue_Float64Value); ok {
				float64Builder.Append(v.Float64Value)
			} else if _, ok := val.Value.(*arrowv1.ScalarValue_NullValue); ok {
				float64Builder.AppendNull()
			} else {
				return nil, fmt.Errorf("mixed types in list not supported")
			}
		}

	case *arrowv1.ScalarValue_Utf8Value:
		dataType = arrow.BinaryTypes.String
		builder = array.NewStringBuilder(allocator)
		defer builder.Release()

		stringBuilder := builder.(*array.StringBuilder)
		for _, val := range values {
			if v, ok := val.Value.(*arrowv1.ScalarValue_Utf8Value); ok {
				stringBuilder.Append(v.Utf8Value)
			} else if _, ok := val.Value.(*arrowv1.ScalarValue_NullValue); ok {
				stringBuilder.AppendNull()
			} else {
				return nil, fmt.Errorf("mixed types in list not supported")
			}
		}

	case *arrowv1.ScalarValue_LargeUtf8Value:
		dataType = arrow.BinaryTypes.LargeString
		builder = array.NewLargeStringBuilder(allocator)
		defer builder.Release()

		stringBuilder := builder.(*array.LargeStringBuilder)
		for _, val := range values {
			if v, ok := val.Value.(*arrowv1.ScalarValue_LargeUtf8Value); ok {
				stringBuilder.Append(v.LargeUtf8Value)
			} else if _, ok := val.Value.(*arrowv1.ScalarValue_NullValue); ok {
				stringBuilder.AppendNull()
			} else {
				return nil, fmt.Errorf("mixed types in list not supported")
			}
		}

	case *arrowv1.ScalarValue_BoolValue:
		dataType = arrow.FixedWidthTypes.Boolean
		builder = array.NewBooleanBuilder(allocator)
		defer builder.Release()

		boolBuilder := builder.(*array.BooleanBuilder)
		for _, val := range values {
			if v, ok := val.Value.(*arrowv1.ScalarValue_BoolValue); ok {
				boolBuilder.Append(v.BoolValue)
			} else if _, ok := val.Value.(*arrowv1.ScalarValue_NullValue); ok {
				boolBuilder.AppendNull()
			} else {
				return nil, fmt.Errorf("mixed types in list not supported")
			}
		}

	default:
		return nil, fmt.Errorf("unsupported scalar value type for list: %T", firstValue.Value)
	}

	// Build the array
	arr := builder.NewArray()
	defer arr.Release()

	// Create a schema with a single field
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "item", Type: dataType, Nullable: true},
	}, nil)

	// Create a record with the array
	record := array.NewRecord(schema, []arrow.Array{arr}, int64(len(values)))
	defer record.Release()

	// Serialize to Arrow IPC format
	bws := &internal.BufferWriteSeeker{}
	writer, err := ipc.NewFileWriter(bws, ipc.WithSchema(schema), ipc.WithAllocator(allocator))
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow IPC writer: %w", err)
	}

	if err := writer.Write(record); err != nil {
		return nil, fmt.Errorf("failed to write Arrow record: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close Arrow writer: %w", err)
	}

	// Convert Arrow schema to protobuf Schema
	protoSchema, err := arrowSchemaToProto(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to convert schema to proto: %w", err)
	}

	return &arrowv1.ScalarListValue{
		ArrowData: bws.Bytes(),
		Schema:    protoSchema,
	}, nil
}

// arrowSchemaToProto converts an Arrow schema to a protobuf Schema
func arrowSchemaToProto(schema *arrow.Schema) (*arrowv1.Schema, error) {
	columns := make([]*arrowv1.Field, len(schema.Fields()))

	for i, field := range schema.Fields() {
		arrowType, err := arrowDataTypeToProto(field.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field %s: %w", field.Name, err)
		}

		columns[i] = &arrowv1.Field{
			Name:      field.Name,
			Nullable:  field.Nullable,
			ArrowType: arrowType,
		}
	}

	return &arrowv1.Schema{
		Columns: columns,
	}, nil
}

// arrowDataTypeToProto converts an Arrow DataType to a protobuf ArrowType
func arrowDataTypeToProto(dataType arrow.DataType) (*arrowv1.ArrowType, error) {
	switch dataType.(type) {
	case *arrow.Int64Type:
		return &arrowv1.ArrowType{
			ArrowTypeEnum: &arrowv1.ArrowType_Int64{
				Int64: &arrowv1.EmptyMessage{},
			},
		}, nil

	case *arrow.Float64Type:
		return &arrowv1.ArrowType{
			ArrowTypeEnum: &arrowv1.ArrowType_Float64{
				Float64: &arrowv1.EmptyMessage{},
			},
		}, nil

	case *arrow.StringType:
		return &arrowv1.ArrowType{
			ArrowTypeEnum: &arrowv1.ArrowType_Utf8{
				Utf8: &arrowv1.EmptyMessage{},
			},
		}, nil

	case *arrow.LargeStringType:
		return &arrowv1.ArrowType{
			ArrowTypeEnum: &arrowv1.ArrowType_LargeUtf8{
				LargeUtf8: &arrowv1.EmptyMessage{},
			},
		}, nil

	case *arrow.BooleanType:
		return &arrowv1.ArrowType{
			ArrowTypeEnum: &arrowv1.ArrowType_Bool{
				Bool: &arrowv1.EmptyMessage{},
			},
		}, nil

	default:
		return nil, fmt.Errorf("unsupported Arrow data type: %s", dataType.Name())
	}
}

// ProtoToArrowDataType converts a protobuf ArrowType to an Arrow DataType
func ProtoToArrowDataType(protoType *arrowv1.ArrowType) (arrow.DataType, error) {
	if protoType == nil {
		return nil, fmt.Errorf("proto arrow type is nil")
	}

	switch protoType.ArrowTypeEnum.(type) {
	case *arrowv1.ArrowType_Int64:
		return arrow.PrimitiveTypes.Int64, nil
	case *arrowv1.ArrowType_Float64:
		return arrow.PrimitiveTypes.Float64, nil
	case *arrowv1.ArrowType_Utf8:
		return arrow.BinaryTypes.String, nil
	case *arrowv1.ArrowType_LargeUtf8:
		return arrow.BinaryTypes.LargeString, nil
	case *arrowv1.ArrowType_Bool:
		return arrow.FixedWidthTypes.Boolean, nil
	case *arrowv1.ArrowType_LargeList:
		largeList := protoType.ArrowTypeEnum.(*arrowv1.ArrowType_LargeList)
		if largeList.LargeList.FieldType != nil {
			return ProtoToArrowDataType(largeList.LargeList.FieldType.ArrowType)
		}
		return nil, fmt.Errorf("large list has no field type")
	default:
		return nil, fmt.Errorf("unsupported protobuf Arrow type: %T", protoType.ArrowTypeEnum)
	}
}
