package expr

import (
	"testing"
)

func TestChalkFunctions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		// String manipulation functions
		{
			name:     "Upper function",
			expr:     Upper(String("hello")),
			expected: "upper(hello)",
		},
		{
			name:     "Lower function",
			expr:     Lower(String("HELLO")),
			expected: "lower(HELLO)",
		},
		{
			name:     "Trim function without chars",
			expr:     Trim(String("  hello  ")),
			expected: "trim(  hello  )",
		},
		{
			name:     "Trim function with chars",
			expr:     Trim(String("xxhelloxx"), String("x")),
			expected: "trim(xxhelloxx, x)",
		},
		{
			name:     "Length function",
			expr:     Length(String("hello")),
			expected: "length(hello)",
		},
		{
			name:     "Split function",
			expr:     Split(String("a,b,c"), String(",")),
			expected: "split(a,b,c, ,)",
		},
		{
			name:     "Split function with max split",
			expr:     Split(String("a,b,c"), String(","), Int(1)),
			expected: "split(a,b,c, ,, 1)",
		},
		{
			name:     "StartsWith function",
			expr:     StartsWith(String("hello"), String("he")),
			expected: "starts_with(hello, he)",
		},
		{
			name:     "EndsWith function",
			expr:     EndsWith(String("hello"), String("lo")),
			expected: "ends_with(hello, lo)",
		},
		{
			name:     "Replace function",
			expr:     Replace(String("hello world"), String("world"), String("universe")),
			expected: "replace(hello world, world, universe)",
		},
		{
			name:     "Concat function",
			expr:     Concat(String("hello"), String(" "), String("world")),
			expected: "concat(hello,  , world)",
		},

		// Mathematical functions
		{
			name:     "Sin function",
			expr:     Sin(Float64(3.14)),
			expected: "sin(3.14)",
		},
		{
			name:     "Cos function",
			expr:     Cos(Float64(0)),
			expected: "cos(0)",
		},
		{
			name:     "Sqrt function",
			expr:     Sqrt(Float64(16)),
			expected: "sqrt(16)",
		},
		{
			name:     "Abs function",
			expr:     Abs(Int(-5)),
			expected: "abs(-5)",
		},
		{
			name:     "Round function without decimals",
			expr:     Round(Float64(3.14159)),
			expected: "round(3.14159)",
		},
		{
			name:     "Round function with decimals",
			expr:     Round(Float64(3.14159), Int(2)),
			expected: "round(3.14159, 2)",
		},
		{
			name:     "Floor function",
			expr:     Floor(Float64(3.9)),
			expected: "floor(3.9)",
		},
		{
			name:     "Ceiling function",
			expr:     Ceiling(Float64(3.1)),
			expected: "ceiling(3.1)",
		},

		// Date/time functions
		{
			name:     "DayOfMonth function",
			expr:     DayOfMonth(Col("timestamp_col")),
			expected: "day_of_month(timestamp_col)",
		},
		{
			name:     "Month function",
			expr:     Month(Col("timestamp_col")),
			expected: "month(timestamp_col)",
		},
		{
			name:     "Year function",
			expr:     Year(Col("timestamp_col")),
			expected: "year(timestamp_col)",
		},
		{
			name:     "Hour function",
			expr:     Hour(Col("timestamp_col")),
			expected: "hour(timestamp_col)",
		},
		{
			name:     "CurrentDate function",
			expr:     CurrentDate(),
			expected: "current_date()",
		},

		// Array functions
		{
			name:     "ArrayJoin function",
			expr:     ArrayJoin(Col("array_col"), String(",")),
			expected: "array_join(array_col, ,)",
		},
		{
			name:     "ArrayConstructor function",
			expr:     ArrayConstructor(Int(1), Int(2), Int(3)),
			expected: "array_constructor(1, 2, 3)",
		},
		{
			name:     "Cardinality function",
			expr:     Cardinality(Col("array_col")),
			expected: "cardinality(array_col)",
		},
		{
			name:     "ElementAt function",
			expr:     ElementAt(Col("array_col"), Int(0)),
			expected: "element_at(array_col, 0)",
		},
		{
			name:     "ArraySort function",
			expr:     ArraySort(Col("array_col")),
			expected: "array_sort(array_col)",
		},
		{
			name:     "ArrayMax function",
			expr:     ArrayMax(Col("numeric_array")),
			expected: "array_max(numeric_array)",
		},

		// Hash functions
		{
			name:     "Md5 function",
			expr:     Md5(Col("data_col")),
			expected: "md5(data_col)",
		},
		{
			name:     "Sha256 function",
			expr:     Sha256(Col("data_col")),
			expected: "sha256(data_col)",
		},

		// Logical functions
		{
			name:     "And function",
			expr:     And(Bool(true), Bool(false)),
			expected: "and(true, false)",
		},
		{
			name:     "Eq function",
			expr:     Eq(Int(1), Int(2)),
			expected: "eq(1, 2)",
		},
		{
			name:     "Lt function",
			expr:     Lt(Int(1), Int(2)),
			expected: "lt(1, 2)",
		},

		// Control flow functions
		{
			name:     "IfElse function",
			expr:     IfElse(Col("condition"), String("yes"), String("no")),
			expected: "if_else(condition, yes, no)",
		},
		{
			name:     "Coalesce function",
			expr:     Coalesce(Col("nullable1"), Col("nullable2"), String("default")),
			expected: "coalesce(nullable1, nullable2, default)",
		},

		// Conversion functions
		{
			name:     "FromBase function",
			expr:     FromBase(String("1010"), Int(2)),
			expected: "from_base(1010, 2)",
		},

		// String similarity functions
		{
			name:     "SequenceMatcherRatio function",
			expr:     SequenceMatcherRatio(String("hello"), String("hallo")),
			expected: "sequence_matcher_ratio(hello, hallo)",
		},
		{
			name:     "JaccardSimilarity function",
			expr:     JaccardSimilarity(String("abc"), String("bcd")),
			expected: "jaccard_similarity(abc, bcd)",
		},

		// Binary functions
		{
			name:     "FromBigEndian64 function",
			expr:     FromBigEndian64(Col("binary_data")),
			expected: "from_big_endian_64(binary_data)",
		},
		{
			name:     "ToHex function",
			expr:     ToHex(Col("binary_data")),
			expected: "to_hex(binary_data)",
		},
		{
			name:     "FromUtf8 function",
			expr:     FromUtf8(Col("binary_data")),
			expected: "from_utf8(binary_data)",
		},
		{
			name:     "ToUtf8 function",
			expr:     ToUtf8(String("hello")),
			expected: "to_utf8(hello)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.expr.String()
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func TestComplexChalkFunctionCombinations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name:     "Nested string functions",
			expr:     Upper(Trim(Lower(Col("name")))),
			expected: "upper(trim(lower(name)))",
		},
		{
			name:     "Array operations with filtering",
			expr:     ArrayMax(ArraySort(Col("scores"))),
			expected: "array_max(array_sort(scores))",
		},
		{
			name:     "Mathematical expression with functions",
			expr:     Round(Sqrt(Col("area").Div(Float64(3.14159))), Int(2)),
			expected: "round(sqrt(/(area, 3.14159)), 2)",
		},
		{
			name:     "Date extraction chain",
			expr:     Year(FromIso8601Timestamp(Col("date_string"))),
			expected: "year(from_iso8601_timestamp(date_string))",
		},
		{
			name: "Conditional with string functions",
			expr: IfElse(
				StartsWith(Col("email"), String("admin")),
				Upper(String("admin")),
				Lower(Col("role")),
			),
			expected: "if_else(starts_with(email, admin), upper(admin), lower(role))",
		},
		{
			name:     "Array manipulation chain",
			expr:     ArrayJoin(ArrayDistinct(ArraySort(Col("tags"))), String(", ")),
			expected: "array_join(array_distinct(array_sort(tags)), , )",
		},
		{
			name:     "Hash chain with encoding",
			expr:     ToHex(Sha256(ToUtf8(Col("password")))),
			expected: "to_hex(sha256(to_utf8(password)))",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.expr.String()
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func TestFunctionFluentAPI(t *testing.T) {
	t.Parallel()

	// Test that function results can be used in fluent API
	result := Upper(Col("name")).Eq(String("JOHN")).And(Length(Col("email")).Gt(Int(5)))
	expected := "AND(=(upper(name), JOHN), >(length(email), 5))"

	actual := result.String()
	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestOptionalParameters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name:     "Trim without optional parameter",
			expr:     Trim(String("  hello  ")),
			expected: "trim(  hello  )",
		},
		{
			name:     "Trim with optional parameter",
			expr:     Trim(String("xxhelloxx"), String("x")),
			expected: "trim(xxhelloxx, x)",
		},
		{
			name:     "Round without decimal places",
			expr:     Round(Float64(3.14159)),
			expected: "round(3.14159)",
		},
		{
			name:     "Round with decimal places",
			expr:     Round(Float64(3.14159), Int(2)),
			expected: "round(3.14159, 2)",
		},
		{
			name:     "Split without max split",
			expr:     Split(String("a,b,c"), String(",")),
			expected: "split(a,b,c, ,)",
		},
		{
			name:     "Split with max split",
			expr:     Split(String("a,b,c"), String(","), Int(1)),
			expected: "split(a,b,c, ,, 1)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.expr.String()
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}
		})
	}
}
