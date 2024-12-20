package internal

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var NameTag = "name"
var WindowsTag = "windows"
var ChalkTag = "chalk"

var NowTimeFormat = "2006-01-02T15:04:05.000000-07:00"

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func StringOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

// FormatBucketDuration takes an integer number of seconds and
// returns a string representation that satisfies these conditions:
//   - the largest possible unit of time (e.g. "10m" instead of "600s")
//   - a single unit of time (e.g. "601s" instead of "10m1s")
func FormatBucketDuration(duration int) string {
	units := []string{"s", "m", "h", "d"}
	divisors := []int{60, 60, 24, 7}

	for i, unit := range units {
		div := divisors[i]
		if duration%div != 0 {
			return fmt.Sprintf("%d%s", duration, unit)
		}
		duration = duration / div
	}

	return fmt.Sprintf("%dw", duration)
}

// ParseBucketDuration parses a bucket duration string
// and returns the duration in seconds.
// The input string must be of the form "Nunit" where
// N is a positive integer and unit is one of "s", "m", "h", "d", or "w".
func ParseBucketDuration(durationStr string) (int, error) {
	unitMap := map[string]int{
		"s": 1,
		"m": 60,
		"h": 3600,
		"d": 86400,
		"w": 604800,
	}

	// Parse the input string
	re := regexp.MustCompile(`^(\d+)([smhdw])$`)
	matches := re.FindStringSubmatch(durationStr)
	if matches == nil {
		return 0, fmt.Errorf("invalid bucket duration string: %s", durationStr)
	}

	// Convert the duration to seconds
	duration, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	unit := matches[2]
	seconds := duration * unitMap[unit]
	return seconds, nil
}

func ExpandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", homeDir, 1), nil
}

func getFeatureNameFromFqn(fqn string) string {
	lastPart := strings.Split(fqn, ".")
	return lastPart[len(lastPart)-1]
}

func getFqnRoot(s string) string {
	return strings.Split(s, ".")[0]
}

func Int64ToInt(value int64) (int, error) {
	// Check if the value fits in the range of an int
	if value < math.MinInt || value > math.MaxInt {
		return 0, errors.New("value out of range for int conversion")
	}
	return int(value), nil
}

// ChalkpySnakeCase aims to be in parity with
// our Python implementation of snake_case
func ChalkpySnakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isASCIIUpper(c) {
			if i > 0 && s[i-1] != '.' {
				b = append(b, '_')
			}
			c += 'a' - 'A'
		} else if isASCIIDigit(c) && i > 0 && isASCIILower(s[i-1]) {
			b = append(b, '_')
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func ResolveFeatureName(field reflect.StructField) (string, error) {
	if tag := field.Tag.Get(NameTag); tag != "" {
		return tag, nil
	}
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		parts := strings.Split(jsonTag, ",")
		if len(parts) > 0 {
			return parts[0], nil
		}
	}
	versioned := field.Tag.Get("versioned")
	fieldName := ChalkpySnakeCase(field.Name)
	if versioned == "true" {
		parts := strings.Split(fieldName, "_")
		nameErr := fmt.Errorf(
			"versioned feature must have a version suffix `VN` at the"+
				" end of the attribute name, but found '%s' instead",
			fieldName,
		)
		if len(parts) == 1 {
			return "", nameErr
		}
		lastPart := parts[len(parts)-1]
		if !strings.HasPrefix(lastPart, "v") {
			return "", nameErr
		}
		version := lastPart[1:]
		prefix := strings.Join(parts[:len(parts)-1], "_")
		if version == "1" {
			fieldName = prefix
		} else {
			fieldName = prefix + "@" + version
		}
	} else if strings.HasPrefix(versioned, "default(") && strings.HasSuffix(versioned, ")") {
		version := versioned[len("default(") : len(versioned)-len(")")]
		_, convertErr := strconv.Atoi(version)
		if convertErr != nil {
			return "", fmt.Errorf(
				"expected struct tag `versioned:\"default(N)\"` "+
					"where N is an integer, but found %s instead",
				versioned,
			)
		}
		if version != "1" {
			fieldName = fieldName + "@" + version
		}
	} else if versioned != "" {
		return "", fmt.Errorf(
			"expected struct tag `versioned:\"true\"` or `versioned:\"default(N)\"` "+
				"where N is an integer, but found '%s' instead",
			versioned,
		)
	}
	return fieldName, nil
}

func GetWindowBucketsFromStructTag(field reflect.StructField) ([]string, error) {
	tag := field.Tag.Get(WindowsTag)
	tags := strings.Split(tag, ",")
	if tag == "" || len(tags) == 0 {
		return nil, errors.Newf("Window bucket struct tag missing or empty, e.g. `%s:\"1m,5m,...\"`", WindowsTag)
	}
	return tags, nil
}

func HasDontOmitTag(field reflect.StructField) bool {
	chalkTags := strings.Split(field.Tag.Get(ChalkTag), ",")
	return colls.Contains(chalkTags, "dontomit")
}

func GetWindowBucketsSecondsFromStructTag(field reflect.StructField) ([]int, error) {
	buckets, err := GetWindowBucketsFromStructTag(field)
	if err != nil {
		return nil, err
	}

	seconds := make([]int, len(buckets))
	for i, bucket := range buckets {
		val, err := ParseBucketDuration(bucket)
		if err != nil {
			return nil, err
		}
		seconds[i] = val
	}
	return seconds, nil
}

func allValid(l int) []bool {
	valid := make([]bool, l)
	for i := range valid {
		valid[i] = true
	}
	return valid
}

func GetBucketFromFqn(fqn string) (string, error) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]
	lastSectionSplit := strings.Split(lastSection, "__")
	formatErr := fmt.Errorf(
		"error unmarshalling value for windowed bucket feature %s: "+
			"expected windowed bucket feature to have fqn of the format "+
			"`{fqn}__{bucket seconds}__` ",
		fqn,
	)
	if len(lastSectionSplit) < 2 {
		return "", formatErr
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return "", formatErr
	}
	return FormatBucketDuration(seconds), nil
}

func SingleInputsToBulkInputs(singleInputs map[string]any) (map[string]any, error) {
	bulkInputs := make(map[string]any)
	for k, singleValue := range singleInputs {
		slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(singleValue)), 1, 1)
		slice.Index(0).Set(reflect.ValueOf(singleValue))
		bulkInputs[k] = slice.Interface()
	}
	return bulkInputs, nil
}

func getFieldToPythonName(structType reflect.Type) (map[string]string, error) {
	isDataclass := IsTypeDataclass(structType)
	res := make(map[string]string)
	namespace := ChalkpySnakeCase(structType.Name())
	for i := 0; i < structType.NumField(); i++ {
		pythonName, err := ResolveFeatureName(structType.Field(i))
		if err != nil {
			return nil, errors.New("failed to resolve field name")
		}
		if !isDataclass {
			// Don't prepend namespace if it is a dataclass
			pythonName = fmt.Sprintf("%s.%s", namespace, pythonName)
		}
		res[structType.Field(i).Name] = pythonName
	}
	return res, nil
}

func preprocessStructSingle(structValue reflect.Value, fieldToPythonName map[string]string, expectedFieldType *reflect.Type) (any, error) {
	var fields []reflect.StructField
	var values []reflect.Value
	structType := structValue.Type()
	for i := 0; i < structType.NumField(); i++ {
		pythonName := fieldToPythonName[structType.Field(i).Name]
		converted, err := PreprocessIfStruct(structValue.Field(i).Interface())
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"failed to convert inner feature struct for field '%s'",
				structType.Field(i).Name,
			)
		}

		rConverted := reflect.ValueOf(converted)
		if (rConverted.IsValid() && !rConverted.IsNil()) ||
			IsTypeDataclass(structType) ||
			HasDontOmitTag(structType.Field(i)) {
			// We omit nil fields unless `chalk:"dontomit"`
			// is specified or if the struct is a dataclass
			newTag := fmt.Sprintf(`json:"%s"`, pythonName)
			if tag := structType.Field(i).Tag; tag != "" {
				newTag = newTag + " " + string(tag)
			}
			fields = append(fields, reflect.StructField{
				Name: structType.Field(i).Name,
				Type: reflect.TypeOf(converted),
				Tag:  reflect.StructTag(newTag),
			})
			values = append(values, rConverted)
		}
	}

	newStructType := reflect.StructOf(fields)
	if len(fields) == 0 {
		return reflect.Zero(reflect.PointerTo(newStructType)), nil
	}

	newStruct := reflect.New(newStructType)
	if expectedFieldType != nil {
		// CHA-5430 - inefficient with large number of has-many structs, but unblocks a customer.
		if newStructType.NumField() != (*expectedFieldType).NumField() {
			return nil, errors.Newf(
				"expected struct to have %d fields but got %d",
				(*expectedFieldType).NumField(),
				newStructType.NumField(),
			)
		}
		for i := 0; i < newStructType.NumField(); i++ {
			expectedField := (*expectedFieldType).Field(i)
			newField := newStructType.Field(i)
			if expectedField.Name != newField.Name {
				return nil, errors.Newf(
					"expected field '%s' but got '%s'",
					expectedField.Name,
					newField.Name,
				)
			}
			if expectedField.Type != newField.Type {
				return nil, errors.Newf(
					"expected field '%s' to have type '%s' but got '%s'",
					expectedField.Name,
					expectedField.Type,
					newField.Type,
				)
			}
			if expectedField.Tag != newField.Tag {
				return nil, errors.Newf(
					"expected field '%s' to have tag '%s' but got '%s'",
					expectedField.Name,
					expectedField.Tag,
					newField.Tag,
				)
			}
		}
	}
	for i, value := range values {
		newStruct.Elem().Field(i).Set(value)
	}
	return newStruct.Interface(), nil
}

func PreprocessIfStruct(values any) (any, error) {
	// Does (unfortunately) two things:
	// (1) Removes nil fields from feature structs, unless the `chalk:"dontomit"` tag is present.
	// (2) Converts feature structs and dataclasses to a format that the server expects.
	//     i.e. When the user passes in a has-one feature struct or a list of has-many structs,
	//          it gets serialized into:
	//
	// {
	//     "FullName": "John Doe",
	//     "Amount": 100,
	// }
	//
	// when we really want:
	//
	// {
	//     "user.full_name": "John Doe",
	//     "user.amount": 100,
	// }
	//
	// Meanwhile, dataclasses are serialized by default as:
	//
	// {
	//     "Lat": 37.7749,
	//     "Lng": 122.4194,
	// }
	//
	// when we want
	// {
	//     "lat": 37.7749,
	//     "lng": 122.4194,
	// }
	//
	rValues := reflect.ValueOf(values)
	if rValues.Kind() == reflect.Ptr {
		rValues = rValues.Elem()
	}
	if !rValues.IsValid() {
		return values, nil
	}

	if IsStruct(rValues.Type()) {
		// This is a has-one feature
		fieldNameToPythonName, err := getFieldToPythonName(rValues.Type())
		if err != nil {
			return nil, errors.Wrap(err, "failed to get feature struct field to python name mapping")
		}
		return preprocessStructSingle(rValues, fieldNameToPythonName, nil)
	}

	if rValues.Type().Kind() != reflect.Slice || rValues.Len() == 0 {
		return values, nil
	}
	elemType := rValues.Type().Elem()
	if !IsStruct(elemType) {
		return values, nil
	}

	// This is a list of dataclasses, or a has-many list of features.
	fieldNameToPythonName, err := getFieldToPythonName(elemType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get feature struct field to python name mapping")
	}

	var newSlice *reflect.Value
	var expectedFieldType *reflect.Type
	for i := 0; i < rValues.Len(); i++ {
		newStruct, err := preprocessStructSingle(rValues.Index(i), fieldNameToPythonName, expectedFieldType)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert feature struct: %v", rValues.Index(i).Interface())
		}
		if newSlice == nil {
			// CHA-5430 - we can't assume that the entire slice is of the same type as the first element
			newExpectedFieldType := reflect.TypeOf(newStruct).Elem()
			expectedFieldType = &newExpectedFieldType
			newSliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(newStruct)), rValues.Len(), rValues.Len())
			newSlice = &newSliceValue
		}
		newSlice.Index(i).Set(reflect.ValueOf(newStruct))
	}
	return newSlice.Interface(), nil
}
