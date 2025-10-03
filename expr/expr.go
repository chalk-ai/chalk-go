package expr

import (
	"fmt"

	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
)

// Expr represents a logical expression node that can be built fluently

type ExprI interface {
	// String representation
	String() string
	// Internal method to get the expression type for building
	exprType() string
}

type Expr interface {
	ExprI

	// Arithmetic operations
	Add(other Expr) Expr
	Sub(other Expr) Expr
	Mul(other Expr) Expr
	Div(other Expr) Expr

	// Comparison operations
	Eq(other Expr) Expr
	Ne(other Expr) Expr
	Lt(other Expr) Expr
	Le(other Expr) Expr
	Gt(other Expr) Expr
	Ge(other Expr) Expr

	// Logical operations
	And(other Expr) Expr
	Or(other Expr) Expr
	Not() Expr

	// Null checking
	IsNull() Expr
	IsNotNull() Expr

	// Attribute access
	Attr(attribute string) Expr

	// Alias
	As(alias string) Expr

	// Bracket access

	// Function application
	Apply(args ...Expr) Expr
}

// DataFrameExpr represents expressions that operate on DataFrames
type DataFrameExpr interface {
	ExprI

	Filter(condition ExprI) DataFrameExpr
	Select(selection Expr) DataFrameExpr
	Agg(aggFunc string, args ...Expr) Expr
}

// Binary operation helper
func binaryOp(left Expr, op string, right Expr) Expr {
	return &CallExpr{
		Function: Identifier(op),
		Args:     []Expr{left, right},
	}
}

// Unary operation helper
func unaryOp(op string, operand Expr) Expr {
	return &CallExpr{
		Function: Identifier(op),
		Args:     []Expr{operand},
	}
}

// IdentifierExpr represents an identifier (variable name, column name, etc.)
type IdentifierExpr struct {
	Expr
	Name string
}

func (e *IdentifierExpr) exprType() string {
	return "identifier"
}

func (e *IdentifierExpr) String() string {
	return e.Name
}

// Implement Expr interface for IdentifierExpr
func (e *IdentifierExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *IdentifierExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *IdentifierExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *IdentifierExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *IdentifierExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *IdentifierExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *IdentifierExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *IdentifierExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *IdentifierExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *IdentifierExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *IdentifierExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *IdentifierExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *IdentifierExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *IdentifierExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *IdentifierExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *IdentifierExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *IdentifierExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *IdentifierExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// LiteralExpr represents a literal value using Arrow scalar values
type LiteralExpr struct {
	Expr
	ScalarValue *arrowv1.ScalarValue
}

func (e *LiteralExpr) exprType() string {
	return "literal"
}

func (e *LiteralExpr) String() string {
	if e.ScalarValue == nil {
		return "null"
	}

	switch v := e.ScalarValue.Value.(type) {
	case *arrowv1.ScalarValue_NullValue:
		return "null"
	case *arrowv1.ScalarValue_BoolValue:
		return fmt.Sprintf("%t", v.BoolValue)
	case *arrowv1.ScalarValue_Float64Value:
		return fmt.Sprintf("%g", v.Float64Value)
	case *arrowv1.ScalarValue_Int64Value:
		return fmt.Sprintf("%d", v.Int64Value)
	case *arrowv1.ScalarValue_LargeUtf8Value:
		return v.LargeUtf8Value
	case *arrowv1.ScalarValue_Date_64Value:
		return fmt.Sprintf("date64(%d)", v.Date_64Value)
	case *arrowv1.ScalarValue_DurationSecondValue:
		return fmt.Sprintf("duration_s(%d)", v.DurationSecondValue)
	case *arrowv1.ScalarValue_DurationMillisecondValue:
		return fmt.Sprintf("duration_ms(%d)", v.DurationMillisecondValue)
	case *arrowv1.ScalarValue_DurationMicrosecondValue:
		return fmt.Sprintf("duration_us(%d)", v.DurationMicrosecondValue)
	case *arrowv1.ScalarValue_DurationNanosecondValue:
		return fmt.Sprintf("duration_ns(%d)", v.DurationNanosecondValue)
	case *arrowv1.ScalarValue_Utf8Value:
		return v.Utf8Value
	case *arrowv1.ScalarValue_Int8Value:
		return fmt.Sprintf("%d", v.Int8Value)
	case *arrowv1.ScalarValue_Int16Value:
		return fmt.Sprintf("%d", v.Int16Value)
	case *arrowv1.ScalarValue_Int32Value:
		return fmt.Sprintf("%d", v.Int32Value)
	case *arrowv1.ScalarValue_Uint8Value:
		return fmt.Sprintf("%d", v.Uint8Value)
	case *arrowv1.ScalarValue_Uint16Value:
		return fmt.Sprintf("%d", v.Uint16Value)
	case *arrowv1.ScalarValue_Uint32Value:
		return fmt.Sprintf("%d", v.Uint32Value)
	case *arrowv1.ScalarValue_Uint64Value:
		return fmt.Sprintf("%d", v.Uint64Value)
	case *arrowv1.ScalarValue_Float16Value:
		return fmt.Sprintf("%g", v.Float16Value)
	case *arrowv1.ScalarValue_Float32Value:
		return fmt.Sprintf("%g", v.Float32Value)
	case *arrowv1.ScalarValue_Date_32Value:
		return fmt.Sprintf("date32(%d)", v.Date_32Value)
	case *arrowv1.ScalarValue_BinaryValue:
		return fmt.Sprintf("binary(%d bytes)", len(v.BinaryValue))
	case *arrowv1.ScalarValue_LargeBinaryValue:
		return fmt.Sprintf("large_binary(%d bytes)", len(v.LargeBinaryValue))
	case *arrowv1.ScalarValue_TimestampValue:
		if v.TimestampValue != nil {
			var timeValue int64
			var unit string
			switch ts := v.TimestampValue.Value.(type) {
			case *arrowv1.ScalarTimestampValue_TimeSecondValue:
				timeValue = ts.TimeSecondValue
				unit = "s"
			case *arrowv1.ScalarTimestampValue_TimeMillisecondValue:
				timeValue = ts.TimeMillisecondValue
				unit = "ms"
			case *arrowv1.ScalarTimestampValue_TimeMicrosecondValue:
				timeValue = ts.TimeMicrosecondValue
				unit = "μs"
			case *arrowv1.ScalarTimestampValue_TimeNanosecondValue:
				timeValue = ts.TimeNanosecondValue
				unit = "ns"
			default:
				return "timestamp(unknown)"
			}
			timezone := v.TimestampValue.Timezone
			if timezone == "" {
				timezone = "UTC"
			}
			return fmt.Sprintf("timestamp(%d%s, %s)", timeValue, unit, timezone)
		}
		return "timestamp(null)"
	default:
		return fmt.Sprintf("unknown (%T)", e.ScalarValue.Value)
	}
}

// Implement Expr interface for LiteralExpr
func (e *LiteralExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *LiteralExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *LiteralExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *LiteralExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *LiteralExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *LiteralExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *LiteralExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *LiteralExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *LiteralExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *LiteralExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *LiteralExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *LiteralExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *LiteralExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *LiteralExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *LiteralExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *LiteralExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *LiteralExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *LiteralExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// GetAttributeExpr represents field access like arr.length
type GetAttributeExpr struct {
	Parent    Expr
	Attribute string
}

func (e *GetAttributeExpr) exprType() string {
	return "get_attribute"
}

func (e *GetAttributeExpr) String() string {
	return fmt.Sprintf("%s.%s", e.Parent.String(), e.Attribute)
}

// Implement Expr interface for GetAttributeExpr
func (e *GetAttributeExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *GetAttributeExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *GetAttributeExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *GetAttributeExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *GetAttributeExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *GetAttributeExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *GetAttributeExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *GetAttributeExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *GetAttributeExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *GetAttributeExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *GetAttributeExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *GetAttributeExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *GetAttributeExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *GetAttributeExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *GetAttributeExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *GetAttributeExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *GetAttributeExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *GetAttributeExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// CallExpr represents function calls and method calls
type CallExpr struct {
	Expr
	Function Expr
	Args     []Expr
	Kwargs   map[string]Expr
}

func (e *CallExpr) exprType() string {
	return "call"
}

func (e *CallExpr) String() string {
	args := ""
	for i, arg := range e.Args {
		if i > 0 {
			args += ", "
		}
		args += arg.String()
	}

	for key, value := range e.Kwargs {
		if len(args) > 0 {
			args += ", "
		}
		args += fmt.Sprintf("%s=%s", key, value.String())
	}

	return fmt.Sprintf("%s(%s)", e.Function, args)
}

// Implement Expr interface for CallExpr
func (e *CallExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *CallExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *CallExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *CallExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *CallExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *CallExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *CallExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *CallExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *CallExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *CallExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *CallExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *CallExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *CallExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *CallExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *CallExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *CallExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *CallExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *CallExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}

// AliasExpr represents an aliased expression
type AliasExpr struct {
	Expression Expr
	Alias      string
}

func (e *AliasExpr) exprType() string {
	return "alias"
}

func (e *AliasExpr) String() string {
	return fmt.Sprintf("%s AS %s", e.Expression, e.Alias)
}

// Implement Expr interface for AliasExpr
func (e *AliasExpr) Add(other Expr) Expr { return binaryOp(e, "+", other) }
func (e *AliasExpr) Sub(other Expr) Expr { return binaryOp(e, "-", other) }
func (e *AliasExpr) Mul(other Expr) Expr { return binaryOp(e, "*", other) }
func (e *AliasExpr) Div(other Expr) Expr { return binaryOp(e, "/", other) }
func (e *AliasExpr) Eq(other Expr) Expr  { return binaryOp(e, "=", other) }
func (e *AliasExpr) Ne(other Expr) Expr  { return binaryOp(e, "!=", other) }
func (e *AliasExpr) Lt(other Expr) Expr  { return binaryOp(e, "<", other) }
func (e *AliasExpr) Le(other Expr) Expr  { return binaryOp(e, "<=", other) }
func (e *AliasExpr) Gt(other Expr) Expr  { return binaryOp(e, ">", other) }
func (e *AliasExpr) Ge(other Expr) Expr  { return binaryOp(e, ">=", other) }
func (e *AliasExpr) And(other Expr) Expr { return binaryOp(e, "AND", other) }
func (e *AliasExpr) Or(other Expr) Expr  { return binaryOp(e, "OR", other) }
func (e *AliasExpr) Not() Expr           { return unaryOp("NOT", e) }
func (e *AliasExpr) IsNull() Expr        { return unaryOp("IS_NULL", e) }
func (e *AliasExpr) IsNotNull() Expr     { return unaryOp("IS_NOT_NULL", e) }
func (e *AliasExpr) Attr(attribute string) Expr {
	return &GetAttributeExpr{Parent: e, Attribute: attribute}
}
func (e *AliasExpr) As(alias string) Expr {
	return &AliasExpr{Expression: e, Alias: alias}
}
func (e *AliasExpr) Apply(args ...Expr) Expr {
	return &CallExpr{Function: e, Args: args}
}
