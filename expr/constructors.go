package expr

import (
	"fmt"
	"time"

	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
)

// Literal constructors

// Most popular types (save space in serialized message)

// Null creates a null literal expression
func Null() Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_NullValue{
				NullValue: &arrowv1.ArrowType{},
			},
		},
	}
}

// Bool creates a boolean literal expression
func Bool(value bool) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_BoolValue{
				BoolValue: value,
			},
		},
	}
}

// Float64 creates a float64 literal expression
func Float64(value float64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Float64Value{
				Float64Value: value,
			},
		},
	}
}

// Float is an alias for Float64 for convenience
func Float(value float64) Expr {
	return Float64(value)
}

// Int64 creates an int64 literal expression
func Int64(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Int64Value{
				Int64Value: value,
			},
		},
	}
}

// Int is an alias for Int64 for convenience
func Int(value int64) Expr {
	return Int64(value)
}

// LargeUtf8 creates a large UTF-8 string literal expression
func LargeUtf8(value string) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_LargeUtf8Value{
				LargeUtf8Value: value,
			},
		},
	}
}

// Date64 creates a date64 literal expression (days since epoch)
func Date64(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Date_64Value{
				Date_64Value: value,
			},
		},
	}
}

// Duration constructors

// DurationSecond creates a duration in seconds
func DurationSecond(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_DurationSecondValue{
				DurationSecondValue: value,
			},
		},
	}
}

// DurationMillisecond creates a duration in milliseconds
func DurationMillisecond(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_DurationMillisecondValue{
				DurationMillisecondValue: value,
			},
		},
	}
}

// DurationMicrosecond creates a duration in microseconds
func DurationMicrosecond(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_DurationMicrosecondValue{
				DurationMicrosecondValue: value,
			},
		},
	}
}

// DurationNanosecond creates a duration in nanoseconds
func DurationNanosecond(value int64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_DurationNanosecondValue{
				DurationNanosecondValue: value,
			},
		},
	}
}

// String and integer type constructors

// Utf8 creates a UTF-8 string literal expression
func Utf8(value string) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Utf8Value{
				Utf8Value: value,
			},
		},
	}
}

// String is an alias for Utf8 for convenience
func String(value string) Expr {
	return Utf8(value)
}

// Int8 creates an int8 literal expression
func Int8(value int8) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Int8Value{
				Int8Value: int32(value),
			},
		},
	}
}

// Int16 creates an int16 literal expression
func Int16(value int16) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Int16Value{
				Int16Value: int32(value),
			},
		},
	}
}

// Int32 creates an int32 literal expression
func Int32(value int32) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Int32Value{
				Int32Value: value,
			},
		},
	}
}

// Uint8 creates a uint8 literal expression
func Uint8(value uint8) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Uint8Value{
				Uint8Value: uint32(value),
			},
		},
	}
}

// Uint16 creates a uint16 literal expression
func Uint16(value uint16) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Uint16Value{
				Uint16Value: uint32(value),
			},
		},
	}
}

// Uint32 creates a uint32 literal expression
func Uint32(value uint32) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Uint32Value{
				Uint32Value: value,
			},
		},
	}
}

// Uint64 creates a uint64 literal expression
func Uint64(value uint64) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Uint64Value{
				Uint64Value: value,
			},
		},
	}
}

// Float16 creates a float16 literal expression
func Float16(value float32) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Float16Value{
				Float16Value: value,
			},
		},
	}
}

// Float32 creates a float32 literal expression
func Float32(value float32) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Float32Value{
				Float32Value: value,
			},
		},
	}
}

// Date32 creates a date32 literal expression (days since epoch as int32)
func Date32(value int32) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_Date_32Value{
				Date_32Value: value,
			},
		},
	}
}

// Binary creates a binary literal expression
func Binary(value []byte) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_BinaryValue{
				BinaryValue: value,
			},
		},
	}
}

// LargeBinary creates a large binary literal expression
func LargeBinary(value []byte) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_LargeBinaryValue{
				LargeBinaryValue: value,
			},
		},
	}
}

// TimestampSecond creates a timestamp literal expression with second precision
func TimestampSecond(value time.Time) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_TimestampValue{
				TimestampValue: &arrowv1.ScalarTimestampValue{
					Value: &arrowv1.ScalarTimestampValue_TimeSecondValue{
						TimeSecondValue: value.Unix(),
					},
					Timezone: value.Location().String(),
				},
			},
		},
	}
}

// TimestampMillisecond creates a timestamp literal expression with millisecond precision
func TimestampMillisecond(value time.Time) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_TimestampValue{
				TimestampValue: &arrowv1.ScalarTimestampValue{
					Value: &arrowv1.ScalarTimestampValue_TimeMillisecondValue{
						TimeMillisecondValue: value.UnixMilli(),
					},
					Timezone: value.Location().String(),
				},
			},
		},
	}
}

// TimestampMicrosecond creates a timestamp literal expression with microsecond precision
func TimestampMicrosecond(value time.Time) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_TimestampValue{
				TimestampValue: &arrowv1.ScalarTimestampValue{
					Value: &arrowv1.ScalarTimestampValue_TimeMicrosecondValue{
						TimeMicrosecondValue: value.UnixMicro(),
					},
					Timezone: value.Location().String(),
				},
			},
		},
	}
}

// TimestampNanosecond creates a timestamp literal expression with nanosecond precision
func TimestampNanosecond(value time.Time) Expr {
	return &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_TimestampValue{
				TimestampValue: &arrowv1.ScalarTimestampValue{
					Value: &arrowv1.ScalarTimestampValue_TimeNanosecondValue{
						TimeNanosecondValue: value.UnixNano(),
					},
					Timezone: value.Location().String(),
				},
			},
		},
	}
}

// Timestamp creates a timestamp literal expression with nanosecond precision (alias for convenience)
func Timestamp(value time.Time) Expr {
	return TimestampNanosecond(value)
}

// Date creates a date literal expression using Date64 (days since epoch)
func Date(value time.Time) Expr {
	// Calculate days since Unix epoch (1970-01-01)
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	days := int64(value.Sub(epoch).Hours() / 24)

	return Date64(days)
}

// Column and identifier constructors

// Col creates a column reference expression
func Col(name string) Expr {
	return &ColumnExpr{
		Name: name,
	}
}

// maybe replace with a relation.name call to Col?
func ColIn(name string, relation string) Expr {
	return &ColumnExpr{
		Name:     name,
		Relation: relation,
	}
}

// Identifier creates an identifier expression
func Identifier(name string) Expr {
	return &IdentifierExpr{
		Name: name,
	}
}

// ColumnExpr represents a column reference
type ColumnExpr struct {
	Expr
	Name     string
	Relation string
}

func (e *ColumnExpr) exprType() string {
	return "column"
}

func (e *ColumnExpr) String() string {
	if e.Relation != "" {
		return fmt.Sprintf("%s.%s", e.Relation, e.Name)
	}
	return e.Name
}

// WithRelation adds a relation qualifier to the column
func (e *ColumnExpr) WithRelation(relation string) *ColumnExpr {
	return &ColumnExpr{
		Name:     e.Name,
		Relation: relation,
	}
}

// Implement Expr interface for ColumnExpr
func (e *ColumnExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *ColumnExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *ColumnExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *ColumnExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *ColumnExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *ColumnExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *ColumnExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *ColumnExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *ColumnExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *ColumnExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *ColumnExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *ColumnExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *ColumnExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *ColumnExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *ColumnExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *ColumnExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *ColumnExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *ColumnExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// Function call constructors

// FunctionCall creates a function call expression
func FunctionCall(name string, args ...Expr) Expr {
	return &CallExpr{
		Function: Identifier(name),
		Args:     args,
		Kwargs:   make(map[string]Expr),
	}
}

// FunctionCallWithKwargs creates a function call expression with keyword arguments
func FunctionCallWithKwargs(name string, args []Expr, kwargs map[string]Expr) Expr {
	if kwargs == nil {
		kwargs = make(map[string]Expr)
	}
	return &CallExpr{
		Function: Identifier(name),
		Args:     args,
		Kwargs:   kwargs,
	}
}

func ChalkNow() Expr {
	return Col("_").Attr("chalk_now")
}

// DataFrame creates a dataframe reference for aggregations
func DataFrame(name string) DataFrameExpr {
	return &DataFrameExprImpl{
		Name: name,
	}
}

// DataFrameExprImpl represents a dataframe for aggregation operations
type DataFrameExprImpl struct {
	Expr
	Name       string
	Conditions []Expr // Store accumulated filter conditions
	Selection  Expr   // Store expression to be aggregated
}

func (e *DataFrameExprImpl) exprType() string {
	return "dataframe"
}

func (e *DataFrameExprImpl) String() string {
	if len(e.Conditions) == 0 {
		return e.Name
	}

	// Build filter chain string
	result := e.Name
	for _, condition := range e.Conditions {
		result = fmt.Sprintf("%s.filter(%s)", result, condition)
	}
	return result
}

// Implement Expr interface for DataFrameExprImpl
func (e *DataFrameExprImpl) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *DataFrameExprImpl) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *DataFrameExprImpl) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *DataFrameExprImpl) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *DataFrameExprImpl) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *DataFrameExprImpl) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *DataFrameExprImpl) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *DataFrameExprImpl) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *DataFrameExprImpl) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *DataFrameExprImpl) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *DataFrameExprImpl) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *DataFrameExprImpl) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *DataFrameExprImpl) Not() Expr           { return unaryOp("NOT", e) }
func (e *DataFrameExprImpl) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *DataFrameExprImpl) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *DataFrameExprImpl) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *DataFrameExprImpl) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *DataFrameExprImpl) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// Implement DataFrameExpr interface
func (e *DataFrameExprImpl) Filter(condition ExprI) DataFrameExpr {
	// Chain conditions with AND logic
	conditions := make([]Expr, len(e.Conditions)+1)
	copy(conditions, e.Conditions)
	conditions[len(e.Conditions)] = condition.(Expr)

	return &DataFrameExprImpl{
		Name:       e.Name,
		Conditions: conditions,
		Selection:  e.Selection,
	}
}

func (e *DataFrameExprImpl) Select(selection Expr) DataFrameExpr {
	return &DataFrameExprImpl{
		Name:       e.Name,
		Conditions: e.Conditions,
		Selection:  selection,
	}
}

func (e *DataFrameExprImpl) Agg(aggFunc string) Expr {
	return &AggregateExprImpl{
		Function:   aggFunc,
		DataFrame:  e,
		Conditions: e.Conditions,
		Selection:  e.Selection,
	}
}

// AggregateExprImpl represents an aggregation expression that returns a scalar
type AggregateExprImpl struct {
	Expr
	Function   string
	DataFrame  DataFrameExpr
	Conditions []Expr // Accumulated filter conditions
	Distinct   bool
	Selection  Expr
}

func (e *AggregateExprImpl) exprType() string {
	return "aggregate"
}

func (e *AggregateExprImpl) String() string {
	// Use the DataFrame's string representation which includes filters
	return fmt.Sprintf("%s.agg(%s)", e.DataFrame, e.Function)
}

// Implement Expr interface for AggregateExprImpl
func (e *AggregateExprImpl) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *AggregateExprImpl) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *AggregateExprImpl) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *AggregateExprImpl) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *AggregateExprImpl) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *AggregateExprImpl) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *AggregateExprImpl) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *AggregateExprImpl) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *AggregateExprImpl) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *AggregateExprImpl) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *AggregateExprImpl) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *AggregateExprImpl) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *AggregateExprImpl) Not() Expr           { return unaryOp("NOT", e) }
func (e *AggregateExprImpl) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *AggregateExprImpl) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *AggregateExprImpl) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *AggregateExprImpl) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *AggregateExprImpl) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// WithDistinct adds DISTINCT to the aggregation
func (e *AggregateExprImpl) WithDistinct() *AggregateExprImpl {
	return &AggregateExprImpl{
		Function:  e.Function,
		DataFrame: e.DataFrame,
		Selection: e.Selection,
		Distinct:  true,
	}
}
