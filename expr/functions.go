package expr

import arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"

// Chalk function constructors for all supported functions
// These create function call expressions that can be used in the fluent API

// String and byte manipulation functions

// FromBigEndian64 converts binary data to int64 using big-endian byte order
func FromBigEndian64(data Expr) Expr {
	return FunctionCall("from_big_endian_64", data)
}

// FromBigEndian32 converts binary data to int32 using big-endian byte order
func FromBigEndian32(data Expr) Expr {
	return FunctionCall("from_big_endian_32", data)
}

// ToHex converts binary data to hexadecimal string representation
func ToHex(data Expr) Expr {
	return FunctionCall("to_hex", data)
}

// FromUtf8 converts to binary data from UTF-8 string
func FromUtf8(data Expr) Expr {
	return FunctionCall("from_utf8", data)
}

// ToUtf8 converts string to UTF-8 binary data
func ToUtf8(str Expr) Expr {
	return FunctionCall("to_utf8", str)
}

// BytesToUtf8 converts from binary data to UTF-8 string (same as chalkpy bytes_to_string(..., encoding='utf-8'))
func BytesToUtf8(bytes Expr) Expr {
	return FunctionCall("bytes_to_string_utf8", bytes)
}

// String manipulation functions

// Strpos finds the position of substring in string (0-based)
func Strpos(str, substr Expr) Expr {
	return FunctionCall("strpos", str, substr)
}

// Strrpos finds the last position of substring in string (0-based)
func Strrpos(str, substr Expr) Expr {
	return FunctionCall("strrpos", str, substr)
}

// Trim removes whitespace from both ends of string
func Trim(str Expr, chars ...Expr) Expr {
	if len(chars) == 0 {
		return FunctionCall("trim", str)
	}
	return FunctionCall("trim", str, chars[0])
}

// Rtrim removes whitespace from right end of string
func Rtrim(str Expr, chars ...Expr) Expr {
	if len(chars) == 0 {
		return FunctionCall("rtrim", str)
	}
	return FunctionCall("rtrim", str, chars[0])
}

// Ltrim removes whitespace from left end of string
func Ltrim(str Expr, chars ...Expr) Expr {
	if len(chars) == 0 {
		return FunctionCall("ltrim", str)
	}
	return FunctionCall("ltrim", str, chars[0])
}

// Upper converts string to uppercase
func Upper(str Expr) Expr {
	return FunctionCall("upper", str)
}

// Lower converts string to lowercase
func Lower(str Expr) Expr {
	return FunctionCall("lower", str)
}

// Length returns the length of a string
func Length(str Expr) Expr {
	return FunctionCall("length", str)
}

// Split splits string by delimiter
func Split(str, delimiter Expr, maxSplit ...Expr) Expr {
	if len(maxSplit) == 0 {
		return FunctionCall("split", str, delimiter)
	}
	return FunctionCall("split", str, delimiter, maxSplit[0])
}

// StartsWith checks if string starts with prefix
func StartsWith(str, prefix Expr) Expr {
	return FunctionCall("starts_with", str, prefix)
}

// EndsWith checks if string ends with suffix
func EndsWith(str, suffix Expr) Expr {
	return FunctionCall("ends_with", str, suffix)
}

// Like performs SQL LIKE pattern matching
func Like(str, pattern Expr) Expr {
	return FunctionCall("like", str, pattern)
}

// RegexpLike performs regular expression matching
func RegexpLike(str, pattern Expr) Expr {
	return FunctionCall("regexp_like", str, pattern)
}

// BoostRegexpFindall finds all matches of regular expression
func BoostRegexpFindall(str, pattern Expr) Expr {
	return FunctionCall("boost_regexp_findall", str, pattern)
}

// RegexpReplace replaces text matching regular expression
func RegexpReplace(str, pattern Expr, replacement ...Expr) Expr {
	if len(replacement) == 0 {
		return FunctionCall("regexp_replace", str, pattern)
	}
	return FunctionCall("regexp_replace", str, pattern, replacement[0])
}

// Replace replaces all occurrences of old substring with new
func Replace(str, old, new Expr) Expr {
	return FunctionCall("replace", str, old, new)
}

// Reverse reverses a string
func Reverse(str Expr) Expr {
	return FunctionCall("reverse", str)
}

// Substr extracts substring
func Substr(str, start Expr, length ...Expr) Expr {
	if len(length) == 0 {
		return FunctionCall("substr", str, start)
	}
	return FunctionCall("substr", str, start, length[0])
}

// ZiSplitPart splits string and returns part at index (0-based)
func ZiSplitPart(str, delimiter, index Expr) Expr {
	return FunctionCall("zi_split_part", str, delimiter, index)
}

// Concat concatenates strings or arrays
func Concat(first Expr, rest ...Expr) Expr {
	args := append([]Expr{first}, rest...)
	return FunctionCall("concat", args...)
}

// String similarity functions

// SequenceMatcherRatio calculates similarity ratio between two strings
func SequenceMatcherRatio(left, right Expr) Expr {
	return FunctionCall("sequence_matcher_ratio", left, right)
}

// JaccardSimilarity calculates Jaccard similarity between two strings
func JaccardSimilarity(left, right Expr) Expr {
	return FunctionCall("jaccard_similarity", left, right)
}

// JaroWinklerDistance calculates Jaro-Winkler distance
func JaroWinklerDistance(left, right, threshold Expr) Expr {
	return FunctionCall("jaro_winkler_distance", left, right, threshold)
}

// PartialRatio calculates partial ratio similarity
func PartialRatio(left, right Expr) Expr {
	return FunctionCall("partial_ratio", left, right)
}

// TokenSetRatio calculates token set ratio similarity
func TokenSetRatio(left, right Expr) Expr {
	return FunctionCall("token_set_ratio", left, right)
}

// TokenSortRatio calculates token sort ratio similarity
func TokenSortRatio(left, right Expr) Expr {
	return FunctionCall("token_sort_ratio", left, right)
}

// Date and time functions

// ParseDatetime parses datetime string with format
func ParseDatetime(dateStr, format Expr) Expr {
	return FunctionCall("parse_datetime", dateStr, format)
}

// ToIso8601 converts timestamp to ISO8601 string
func ToIso8601(timestamp Expr) Expr {
	return FunctionCall("to_iso8601", timestamp)
}

// FromIso8601Timestamp parses ISO8601 timestamp string
func FromIso8601Timestamp(str Expr) Expr {
	return FunctionCall("from_iso8601_timestamp", str)
}

// LastDayOfMonth returns the last day of the month for given timestamp
func LastDayOfMonth(timestamp Expr) Expr {
	return FunctionCall("last_day_of_month", timestamp)
}

// DayOfMonth extracts day of month from timestamp or date
func DayOfMonth(dateTime Expr) Expr {
	return FunctionCall("day_of_month", dateTime)
}

// Month extracts month from timestamp or date
func Month(dateTime Expr) Expr {
	return FunctionCall("month", dateTime)
}

// Year extracts year from timestamp or date
func Year(dateTime Expr) Expr {
	return FunctionCall("year", dateTime)
}

// Hour extracts hour from timestamp
func Hour(timestamp Expr) Expr {
	return FunctionCall("hour", timestamp)
}

// Minute extracts minute from timestamp
func Minute(timestamp Expr) Expr {
	return FunctionCall("minute", timestamp)
}

// Second extracts second from timestamp
func Second(timestamp Expr) Expr {
	return FunctionCall("second", timestamp)
}

// ToUnixtime converts timestamp to Unix timestamp (seconds since epoch)
func ToUnixtime(timestamp Expr) Expr {
	return FunctionCall("to_unixtime", timestamp)
}

// FromUnixSeconds converts Unix timestamp to timestamp
func FromUnixSeconds(seconds Expr) Expr {
	return FunctionCall("from_unixtime", seconds)
}

// DateTrunc truncates timestamp to specified unit
func DateTrunc(unit, timestamp Expr) Expr {
	return FunctionCall("date_trunc", unit, timestamp)
}

// FormatDatetime formats timestamp as string
func FormatDatetime(timestamp, format Expr) Expr {
	return FunctionCall("format_datetime", timestamp, format)
}

// CurrentDate returns current date
func CurrentDate() Expr {
	return FunctionCall("current_date")
}

// Mathematical functions

// Sin calculates sine
func Sin(x Expr) Expr {
	return FunctionCall("sin", x)
}

// Cos calculates cosine
func Cos(x Expr) Expr {
	return FunctionCall("cos", x)
}

// Tan calculates tangent
func Tan(x Expr) Expr {
	return FunctionCall("tan", x)
}

// Sqrt calculates square root
func Sqrt(x Expr) Expr {
	return FunctionCall("sqrt", x)
}

// Log10 calculates base-10 logarithm
func Log10(x Expr) Expr {
	return FunctionCall("log10", x)
}

// Abs calculates absolute value
func Abs(x Expr) Expr {
	return FunctionCall("abs", x)
}

// Round rounds to nearest integer or specified decimal places
func Round(x Expr, decimals ...Expr) Expr {
	if len(decimals) == 0 {
		return FunctionCall("round", x)
	}
	return FunctionCall("round", x, decimals[0])
}

// BankersRound rounds using banker's rounding (round half to even)
func BankersRound(x Expr, decimals ...Expr) Expr {
	if len(decimals) == 0 {
		return FunctionCall("bankers_round", x)
	}
	return FunctionCall("bankers_round", x, decimals[0])
}

// Truncate truncates to integer part
func Truncate(x Expr) Expr {
	return FunctionCall("truncate", x)
}

// Floor rounds down to nearest integer
func Floor(x Expr) Expr {
	return FunctionCall("floor", x)
}

// Ceiling rounds up to nearest integer
func Ceiling(x Expr) Expr {
	return FunctionCall("ceiling", x)
}

// Negate negates a number
func Negate(x Expr) Expr {
	return FunctionCall("negate", x)
}

// Mod calculates modulo
func Mod(a, b Expr) Expr {
	return FunctionCall("%", a, b)
}

// ScalarMin returns minimum of two values
func ScalarMin(a, b Expr) Expr {
	return FunctionCall("scalar_min", a, b)
}

// ScalarMax returns maximum of two values
func ScalarMax(a, b Expr) Expr {
	return FunctionCall("scalar_max", a, b)
}

// IsNan checks if value is NaN
func IsNan(x Expr) Expr {
	return FunctionCall("is_nan", x)
}

// Nan returns NaN value
func Nan() Expr {
	return FunctionCall("nan")
}

// Logical functions

// And performs logical AND
func And(a, b Expr) Expr {
	return FunctionCall("and", a, b)
}

// Lt performs less than comparison
func Lt(a, b Expr) Expr {
	return FunctionCall("lt", a, b)
}

// Gt performs greater than comparison
func Gt(a, b Expr) Expr {
	return FunctionCall("gt", a, b)
}

// Eq performs equality comparison
func Eq(a, b Expr) Expr {
	return FunctionCall("eq", a, b)
}

// Neq performs inequality comparison
func Neq(a, b Expr) Expr {
	return FunctionCall("neq", a, b)
}

// Lte performs less than or equal comparison
func Lte(a, b Expr) Expr {
	return FunctionCall("lte", a, b)
}

// Gte performs greater than or equal comparison
func Gte(a, b Expr) Expr {
	return FunctionCall("gte", a, b)
}

// Array manipulation functions

// ArrayJoin joins array elements with separator
func ArrayJoin(array, separator Expr) Expr {
	return FunctionCall("array_join", array, separator)
}

// ArrayConstructor creates array from elements
func ArrayConstructor(elements ...Expr) Expr {
	return FunctionCall("array_constructor", elements...)
}

// ArrayFrequency returns frequency map of array elements
func ArrayFrequency(array Expr) Expr {
	return FunctionCall("array_frequency", array)
}

// Cardinality returns length of array
func Cardinality(array Expr) Expr {
	return FunctionCall("cardinality", array)
}

// ElementAt returns element at index (1-indexed)
func ElementAt(array, index Expr) Expr {
	return FunctionCall("element_at", array, index)
}

// PythonElementAt returns element at index (0-based for Python compatibility)
func PythonElementAt(array, index Expr) Expr {
	return FunctionCall("python_element_at", array, index)
}

// Slice extracts slice from array
func Slice(array, start, length Expr) Expr {
	return FunctionCall("slice", array, start, length)
}

// Contains checks if array contains element
func Contains(array, element Expr) Expr {
	return FunctionCall("contains", array, element)
}

// ArrayAverage calculates average of numeric array
func ArrayAverage(array Expr) Expr {
	return FunctionCall("array_average", array)
}

// ArraySort sorts array in ascending order
func ArraySort(array Expr) Expr {
	return FunctionCall("array_sort", array)
}

// ArraySortDesc sorts array in descending order
func ArraySortDesc(array Expr) Expr {
	return FunctionCall("array_sort_desc", array)
}

// ArrayDistinct returns unique elements from array
func ArrayDistinct(array Expr) Expr {
	return FunctionCall("array_distinct", array)
}

// ArrayMax returns maximum element from array
func ArrayMax(array Expr) Expr {
	return FunctionCall("array_max", array)
}

// ArrayMin returns minimum element from array
func ArrayMin(array Expr) Expr {
	return FunctionCall("array_min", array)
}

// ArraySum calculates sum of numeric array
func ArraySum(array Expr) Expr {
	return FunctionCall("array_sum", array)
}

// ArrayStddev calculates standard deviation of array
func ArrayStddev(array, isSample Expr) Expr {
	return FunctionCall("array_stddev", array, isSample)
}

// ArrayMedian calculates median of numeric array
func ArrayMedian(array Expr) Expr {
	return FunctionCall("array_median", array)
}

// ArrayMode calculates mode of array
func ArrayMode(array Expr, mode ...Expr) Expr {
	if len(mode) == 0 {
		return FunctionCall("array_mode", array)
	}
	return FunctionCall("array_mode", array, mode[0])
}

// Map manipulation functions

// MapKeysbyTopNValues returns keys with top N values from map
func MapKeysByTopNValues(mapExpr, n Expr) Expr {
	return FunctionCall("map_keys_by_top_n_values", mapExpr, n)
}

// MapGet gets value from map by key
func MapGet(mapExpr, key Expr) Expr {
	return FunctionCall("map_get", mapExpr, key)
}

// MapContains checks if map contains key matching predicate
func MapContains(mapExpr, predicate Expr) Expr {
	return FunctionCall("map_contains", mapExpr, predicate)
}

// Hash functions

// SpookyHashV264 calculates 64-bit SpookyHash
func SpookyHashV264(data Expr) Expr {
	return FunctionCall("spooky_hash_v2_64", data)
}

// SpookyHashV232 calculates 32-bit SpookyHash
func SpookyHashV232(data Expr) Expr {
	return FunctionCall("spooky_hash_v2_32", data)
}

// Md5 calculates MD5 hash
func Md5(data Expr) Expr {
	return FunctionCall("md5", data)
}

// Sha1 calculates SHA-1 hash
func Sha1(data Expr) Expr {
	return FunctionCall("sha1", data)
}

// Sha256 calculates SHA-256 hash
func Sha256(data Expr) Expr {
	return FunctionCall("sha256", data)
}

// Sha512 calculates SHA-512 hash
func Sha512(data Expr) Expr {
	return FunctionCall("sha512", data)
}

// Conversion functions

// FromBase converts string from specified base to integer
func FromBase(str, base Expr) Expr {
	return FunctionCall("from_base", str, base)
}

// Cast converts a value to a target type
func Cast(expr Expr, to *arrowv1.ArrowType) Expr {
	return FunctionCall("cast", expr, &LiteralExpr{
		ScalarValue: &arrowv1.ScalarValue{
			Value: &arrowv1.ScalarValue_NullValue{
				NullValue: to,
			},
		},
		IsArrowScalarObject: true,
	})
}

// Control flow functions

// IfElse conditional expression
func IfElse(condition, trueValue, falseValue Expr) Expr {
	return FunctionCall("if_else", condition, trueValue, falseValue)
}

// Coalesce returns first non-null value
func Coalesce(values ...Expr) Expr {
	return FunctionCall("coalesce", values...)
}

// Fail throws an error with message
func Fail(message Expr) Expr {
	return FunctionCall("fail", message)
}

// GetJsonValue extracts value from JSON
func GetJsonValue(json Expr, path string) Expr {
	return FunctionCall("get_json_value", json, String(path))
}

// Jsonify turns a value into a JSON string
func Jsonify(json Expr) Expr {
	return FunctionCall("jsonify", json)
}

// Now returns current timestamp (convenience function)
func Now() Expr {
	// This would typically be implemented as a special function
	// For now, we'll use a placeholder
	return FunctionCall("now")
}

// TimestampFromString parses timestamp from ISO8601 string
func TimestampFromString(str Expr) Expr {
	return FromIso8601Timestamp(str)
}

// StringFromTimestamp converts timestamp to ISO8601 string
func StringFromTimestamp(timestamp Expr) Expr {
	return ToIso8601(timestamp)
}
