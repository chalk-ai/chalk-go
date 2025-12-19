package expr

import (
	"testing"
)

func TestBasicLiterals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{"int64", Int(42), "42"},
		{"int32", Int32(123), "123"},
		{"int16", Int16(456), "456"},
		{"int8", Int8(78), "78"},
		{"uint64", Uint64(42), "42"},
		{"uint32", Uint32(123), "123"},
		{"uint16", Uint16(456), "456"},
		{"uint8", Uint8(78), "78"},
		{"float", Float(3.14), "3.14"},
		{"float64", Float64(3.14), "3.14"},
		{"float32", Float32(2.71), "2.71"},
		{"string_utf8", String("hello"), "hello"},
		{"string_large_utf8", LargeUtf8("world"), "world"},
		{"bool_true", Bool(true), "true"},
		{"bool_false", Bool(false), "false"},
		{"null", Null(), "null"},
		{"date32", Date32(18000), "date32(18000)"},
		{"date64", Date64(18000), "date64(18000)"},
		{"duration_second", DurationSecond(3600), "duration_s(3600)"},
		{"duration_millisecond", DurationMillisecond(1000), "duration_ms(1000)"},
		{"duration_microsecond", DurationMicrosecond(1000000), "duration_us(1000000)"},
		{"duration_nanosecond", DurationNanosecond(1000000000), "duration_ns(1000000000)"},
		{"binary", Binary([]byte{1, 2, 3}), "binary(3 bytes)"},
		{"large_binary", LargeBinary([]byte{4, 5, 6, 7}), "large_binary(4 bytes)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.expr.String()
			if actual != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, actual)
			}
		})
	}
}

// Test that demonstrates fluent API usage similar to polars
func TestFluentAPI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		// Complex condition chains
		{
			"complex_condition_chain",
			Col("user.transactions").
				Attr("amount").
				Gt(Int(100)).
				And(Col("user.age").Ge(Int(18))).
				Or(Col("user.premium").Eq(Bool(true))),
			"OR(AND(>(user.transactions.amount, 100), >=(user.age, 18)), ==(user.premium, true))",
		},

		// Column references
		{
			"column_reference_simple",
			Col("user_id"),
			"user_id",
		},
		{
			"column_reference_dotted",
			Col("user.name"),
			"user.name",
		},

		// Arithmetic operations
		{
			"arithmetic_basic",
			Int(4).Add(Int(5)),
			"+(4, 5)",
		},
		{
			"arithmetic_chained",
			Int(4).Add(Int(5)).Mul(Int(2)),
			"*(+(4, 5), 2)",
		},
		{
			"math_expression_chain",
			Int(10).Add(Int(5)).Mul(Int(2)).Sub(Int(3)),
			"-(*(+(10, 5), 2), 3)",
		},

		// Comparisons
		{
			"comparison_greater_than",
			Col("age").Gt(Int(18)),
			">(age, 18)",
		},

		// Logical operations
		{
			"logical_and_comparison",
			Col("age").Gt(Int(18)).And(Col("active").Eq(Bool(true))),
			"AND(>(age, 18), ==(active, true))",
		},

		// Function calls
		{
			"function_call_single_arg",
			FunctionCall("abs", Col("value")),
			"abs(value)",
		},
		{
			"function_call_multiple_args",
			FunctionCall("concat", String("Hello"), String(" "), String("World")),
			"concat(Hello,  , World)",
		},

		// Attribute access
		{
			"attribute_access",
			Col("array").Attr("length"),
			"array.length",
		},

		// Aggregations
		{
			"aggregation_with_filter",
			DataFrame("transactions").Filter(Col("amount").Gt(Int(100))).Agg("sum"),
			"transactions.filter(>(amount, 100)).agg(sum)",
		},

		// Alias
		{
			"alias_expression",
			Col("name").As("full_name"),
			"name AS full_name",
		},

		// Null checks
		{
			"null_check_is_not_null",
			Col("email").IsNotNull(),
			"IS_NOT_NULL(email)",
		},
		{
			"null_check_is_null",
			Col("phone").IsNull(),
			"IS_NULL(phone)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.expr.String()
			if actual != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, actual)
			}
		})
	}
}
