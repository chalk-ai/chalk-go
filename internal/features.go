package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func SerializeStaleness(staleness map[string]time.Duration) map[string]string {
	res := map[string]string{}
	for k, v := range staleness {
		res[k] = strconv.Itoa(int(v.Seconds())) + "s"
	}
	return res
}

type Numbers interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Feature struct {
	Fqn string
}

func unwrapFeature(t any) (*Feature, error) {
	reflectValue := reflect.ValueOf(t)
	if featureClass := getFeatureClassFromMember(reflectValue); featureClass != nil {
		// If the user is querying a feature class, e.g.
		//   Features.User.CreditReport instead of Features.User.CreditReport.CreditScore
		return featureClass, nil
	} else if reflectValue.Kind() == reflect.Ptr {
		// Everything but windowed features
		return (*Feature)(reflectValue.UnsafePointer()), nil
	} else if reflectValue.Kind() == reflect.Map {
		// Base windowed feature is typed as a Map.
		// But it is natural for a user to try querying
		// a base windowed feature when they want to query
		// every pseudofeature in the windowed feature.
		// So here we return the base windowed feature
		// with a valid FQN.
		windowedFeatureMap := reflectValue
		keys := windowedFeatureMap.MapKeys()
		if len(keys) == 0 {
			return nil, fmt.Errorf("exception occurred obtaining buckets for windowed feature - no buckets found")
		}
		key := keys[0]
		ptrToPseudofeature := windowedFeatureMap.MapIndex(key)
		castedPtrToPseudofeature := (*Feature)(ptrToPseudofeature.UnsafePointer())

		baseWindowedFeatureFqn := strings.Split(castedPtrToPseudofeature.Fqn, "__")[0]
		baseWindowedFeature := Feature{Fqn: baseWindowedFeatureFqn}
		return &baseWindowedFeature, nil
	}
	typePointedTo := reflect.ValueOf(t).Elem().Type()
	return nil, fmt.Errorf("cannot unwrap object of unsupported type: %s", typePointedTo)
}

func UnwrapFeature(t any) (result *Feature, returnErr error) {
	defer func() {
		panicContents := recover()
		if panicContents == nil {
			return
		}
		detail := "details irretrievable"
		switch typedContents := panicContents.(type) {
		case *reflect.ValueError:
			detail = typedContents.Error()
		case string:
			detail = typedContents
		}
		result = nil
		returnErr = fmt.Errorf("unexpected unwrap feature internal error: %s", detail)
	}()
	return unwrapFeature(t)
}

func convertNumber[T Numbers](anyNumber any) (T, error) {
	// TODO: Possibly unmarshal numbers as the correct type (instead of float64)
	// into FeatureResult, instead of converting them here.
	switch typedNumber := anyNumber.(type) {
	case float64:
		return T(typedNumber), nil
	default:
		castedNumber, ok := anyNumber.(T)
		if !ok {
			var t T
			return t, fmt.Errorf("cannot cast the number '%s' of type '%s' to the specified type '%s'", anyNumber, reflect.TypeOf(typedNumber), reflect.TypeOf(t))
		}
		return castedNumber, nil
	}
}

func convertSliceyNumbers[T Numbers](anySlice []any) ([]T, error) {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		convRes, convErr := convertNumber[T](v)
		if convErr != nil {
			return nil, fmt.Errorf("error converting number-slice element: %w", convErr)
		}
		typedSlice[i] = convRes
	}
	return typedSlice, nil
}

func convertSliceyNonNumbers[T any](anySlice []any) ([]T, error) {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		castRes, ok := v.(T)
		if !ok {
			var t T
			return []T{}, fmt.Errorf("cannot cast the slice element '%s' of type '%s' to the specified type '%s'", v, reflect.TypeOf(v), reflect.TypeOf(t))
		}
		typedSlice[i] = castRes
	}
	return typedSlice, nil
}

func convertIfNumber(value any, kind reflect.Kind) (any, error) {
	// TODO: Figure out if we could possibly
	// do the equivalent by creating a new
	// reflect.Value with reflect.New
	// and reflect.Value.Set.
	var err error
	switch kind {
	case reflect.Int8:
		value, err = convertNumber[int8](value)
	case reflect.Int16:
		value, err = convertNumber[int16](value)
	case reflect.Int32:
		value, err = convertNumber[int32](value)
	case reflect.Int64:
		value, err = convertNumber[int64](value)
	case reflect.Uint8:
		value, err = convertNumber[uint8](value)
	case reflect.Uint16:
		value, err = convertNumber[uint16](value)
	case reflect.Uint32:
		value, err = convertNumber[uint32](value)
	case reflect.Uint64:
		value, err = convertNumber[uint64](value)
	case reflect.Float32:
		value, err = convertNumber[float32](value)
	case reflect.Float64:
		value, err = convertNumber[float64](value)
	}

	return value, err
}

func convertNumberSlice(sliceElemKind reflect.Kind, value any) (any, error) {
	anySlice := value.([]any)
	switch sliceElemKind {
	case reflect.Int8:
		return convertSliceyNumbers[int8](anySlice)
	case reflect.Int16:
		return convertSliceyNumbers[int16](anySlice)
	case reflect.Int32:
		return convertSliceyNumbers[int32](anySlice)
	case reflect.Int64:
		return convertSliceyNumbers[int64](anySlice)
	case reflect.Uint8:
		return convertSliceyNumbers[uint8](anySlice)
	case reflect.Uint16:
		return convertSliceyNumbers[uint16](anySlice)
	case reflect.Uint32:
		return convertSliceyNumbers[uint32](anySlice)
	case reflect.Uint64:
		return convertSliceyNumbers[uint64](anySlice)
	case reflect.Float32:
		return convertSliceyNumbers[float32](anySlice)
	case reflect.Float64:
		return convertSliceyNumbers[float64](anySlice)
	case reflect.String:
		return convertSliceyNonNumbers[string](anySlice)
	case reflect.Bool:
		return convertSliceyNonNumbers[bool](anySlice)
	default:
		return nil, fmt.Errorf("unsupported slice type '%s' when converting number slice", sliceElemKind)
	}
}

func getPointerToCopied(elemType reflect.Type, value any) reflect.Value {
	copied := reflect.New(reflect.TypeOf(value))
	copied.Elem().Set(reflect.ValueOf(value))
	castedPointer := reflect.NewAt(elemType, copied.UnsafePointer())
	return castedPointer
}
func getReflectValue(value any, elemType reflect.Type) (reflect.Value, error) {
	value, convErr := convertIfNumber(value, elemType.Kind())
	if convErr != nil {
		return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
	}
	if elemType.String() == "time.Time" {
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			return reflect.ValueOf(&timeValue), nil
		}

		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return reflect.Value{}, timeErr
		}
		return reflect.ValueOf(&dateValue), nil
	} else if elemType.Kind() == reflect.Slice || elemType.Kind() == reflect.Array {
		value, convErr = convertNumberSlice(elemType.Elem().Kind(), value)
		if convErr != nil {
			return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
		}
		return getPointerToCopied(elemType, value), nil
	} else {
		return getPointerToCopied(elemType, value), nil
	}
}
func getFeatureClassFromMember(field reflect.Value) *Feature {
	if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct && field.Type().Elem().String() != "time.Time" {
		structValue := field.Elem()
		for i := 0; i < structValue.NumField(); i++ {
			memberField := structValue.Field(i)
			if memberField.Kind() != reflect.Ptr {
				continue
			}
			memberFeature, unwrapErr := UnwrapFeature(memberField.Interface())
			if unwrapErr != nil {
				continue
			}
			memberFqn := memberFeature.Fqn

			var featureClassFqn string
			if IsDataclass(structValue) {
				// Dataclass feature classes have members
				// that share the same FQN as the feature class.
				featureClassFqn = memberFqn
			} else {
				featureClassFqn = DesuffixFqn(memberFqn)
			}
			return &Feature{
				Fqn: featureClassFqn,
			}
		}
	}
	return nil
}

var FieldNotFoundError = errors.New("field not found")

func DesuffixFqn(fqn string) string {
	sections := strings.Split(fqn, ".")
	return strings.Join(sections[:len(sections)-1], ".")
}

func getWindowedPseudofeatureMeta(fqn string, fieldMap FqnToField) (*int, *reflect.Value) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]

	lastSectionSplit := strings.Split(lastSection, "__")
	if len(lastSectionSplit) < 2 {
		return nil, nil
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return nil, nil
	}

	featureClassFqn := DesuffixFqn(fqn)
	baseFeatureFqn := featureClassFqn + "." + lastSectionSplit[0]
	baseFeatureField, ok := fieldMap[baseFeatureFqn]
	if !ok {
		return nil, nil
	}

	return &seconds, &baseFeatureField
}

type FqnToField map[string]reflect.Value

func IsDataclass(field reflect.Value) bool {
	if field.Kind() == reflect.Struct {
		if field.NumField() == 0 {
			return false
		}
		for i := 0; i < field.NumField(); i++ {
			fieldMeta := field.Type().Field(i)
			if fieldMeta.Tag.Get("dataclass_field") != "true" {
				return false
			}
		}
		return true
	}

	return false
}
func isDataclassPointer(field reflect.Value) bool {
	if field.Kind() == reflect.Ptr && IsDataclass(field.Elem()) {
		return true
	}
	return false
}
func (t FqnToField) SetFeature(fqn string, value any) error {
	if field, fieldFound := t[fqn]; fieldFound && isDataclassPointer(field) {
		structValue := field.Elem()
		dataclassValues, ok := value.([]any)
		if !ok {
			return fmt.Errorf("error unmarshalling value for dataclass feature %s: value is not a slice", fqn)
		}
		if len(dataclassValues) != structValue.NumField() {
			return fmt.Errorf("error unmarshalling value for dataclass feature %s: expected %d fields, got %s", fqn, structValue.NumField(), dataclassValues)
		}
		for idx, memberValue := range dataclassValues {
			memberFieldMeta := structValue.Type().Field(idx)
			memberField := structValue.Field(idx)
			pythonName := SnakeCase(memberFieldMeta.Name)
			if memberField == (reflect.Value{}) {
				return fmt.Errorf("error unmarshalling value for dataclass feature %s: field %s not found in struct %s", fqn, pythonName, structValue.Type().Name())
			}
			memberFqn := fqn + "." + pythonName
			if err := t.SetFeature(memberFqn, memberValue); err != nil {
				return fmt.Errorf("error unmarshalling value '%s' for dataclass feature '%s': %w", pythonName, fqn, err)
			}
		}
	} else if bucketDuration, baseFeatureField := getWindowedPseudofeatureMeta(fqn, t); bucketDuration != nil && baseFeatureField != nil {
		tagValue := reflect.ValueOf(FormatBucketDuration(*bucketDuration))

		if baseFeatureField.Kind() != reflect.Map {
			return fmt.Errorf(fmt.Sprintf("exception setting windowed feature '%s'", fqn))
		}

		reflectValue, err := getReflectValue(value, baseFeatureField.Type().Elem().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for windowed feature %s: %w", fqn, err)
		}

		baseFeatureField.SetMapIndex(tagValue, reflectValue)
	} else {
		field, fieldFound = t[fqn]
		if !fieldFound {
			return FieldNotFoundError
		}
		reflectValue, err := getReflectValue(value, field.Type().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for feature %s: %w", fqn, err)
		}
		field.Set(reflectValue)
	}

	return nil
}

// InitFeatureInternal is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as
// a pointer to a Feature struct with the appropriate FQN.
func InitFeatureInternal(structValue reflect.Value, fqn string, visited map[string]bool, fieldMap FqnToField) error {
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("feature initialization function argument must be a reflect.Value of the kind reflect.Struct, found %s instead", structValue.Kind().String())
	}

	namespace := structValue.Type().Name()
	if fqn == "" && namespace != "" {
		fqn = SnakeCase(namespace) + "."
	}

	if isVisited, ok := visited[namespace]; ok && isVisited {
		// This is not memoization. Simply a cycle checker while DFSing.
		return nil
	} else {
		visited[namespace] = true
	}

	for i := 0; i < structValue.NumField(); i++ {
		f := structValue.Field(i)
		fieldMeta := structValue.Type().Field(i)

		isTimeField := f.Type().Elem().Kind() == reflect.Struct && f.Type().Elem().String() == "time.Time"

		attributeName := fieldMeta.Name
		updatedFqn := fqn + SnakeCase(attributeName)

		if !f.CanSet() {
			continue
		}

		if f.Type().Elem().Kind() == reflect.Struct && !isTimeField {
			// RECURSIVE CASE.
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}
			featureSet := reflect.New(f.Type().Elem())
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			initErr := InitFeatureInternal(featureSetInDisguise, updatedFqn+".", visited, fieldMap)
			if initErr != nil {
				return initErr
			}
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			}
		} else if f.Kind() == reflect.Map {
			// Creates a map of tag values to pointers to Features.
			// For example, if we have the tag "windows=6h,12h,1d",
			// then the map will be:
			//
			// 		map[string]*int64{
			// 			"6h": &Feature{Fqn: "user.clicks__21600__"},
			// 			"12h": &Feature{Fqn: "user.clicks__43200__"},
			// 			"1d": &Feature{Fqn: "user.clicks__86400__"},
			// 		}
			//
			// Notice that while the values is typed as *int64, it is
			// actually a pointer to a Feature struct. See BASE CASE
			// section below.
			mapValueType := f.Type().Elem()
			if mapValueType.Kind() != reflect.Pointer {
				return fmt.Errorf("the map type for Windowed features should a pointer as its value type, but found %s instead", mapValueType.Kind())
			}
			newMap := reflect.MakeMap(f.Type())
			windows := fieldMeta.Tag.Get("windows")
			for _, tag := range strings.Split(windows, ",") {
				seconds, parseErr := ParseBucketDuration(tag)
				if parseErr != nil {
					return fmt.Errorf("error parsing bucket duration: %s", parseErr)
				}
				windowFqn := updatedFqn + fmt.Sprintf("__%d__", seconds)
				if fieldMap == nil {
					feature := Feature{Fqn: windowFqn}
					ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
					newMap.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
				} else {
					nilPointer := reflect.New(f.Type().Elem()).Elem()
					nilPointer.Set(reflect.Zero(nilPointer.Type()))
					newMap.SetMapIndex(reflect.ValueOf(tag), nilPointer)
				}
			}
			f.Set(newMap)
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			}
		} else {
			// BASE CASE.
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			} else {
				// Create new Feature instance and point to it.
				// The equivalent way of doing it without 'reflect':
				//
				//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
				//

				// Dataclass fields are not actually real features,
				// so when we are initializing the root Features struct (fieldMap == nil),
				// we want to return the parent (real) feature FQN
				// instead of the fake FQN of the dataclass field.
				if IsDataclass(structValue) {
					updatedFqn = DesuffixFqn(updatedFqn)
				}

				feature := Feature{Fqn: updatedFqn}
				ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
				f.Set(ptrInDisguiseToFeature)
			}
		}
	}

	visited[namespace] = false
	return nil
}

func pointerCheck(field reflect.Value) error {
	if field.Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer type but found %s -- make sure the generated feature structs are unchanged, and that every field is of a pointer type except for Windowed feature types", field.Kind())
	}
	return nil
}
