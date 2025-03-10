package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strings"
	"time"
)

type Feature struct {
	Fqn string
}

func DesuffixFqn(fqn string) string {
	sections := strings.Split(fqn, ".")
	return strings.Join(sections[:len(sections)-1], ".")
}

func getFeatureClassFromMember(field reflect.Value) *Feature {
	if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct && field.Type().Elem() != reflect.TypeOf(time.Time{}) {
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
			if internal.IsDataclass(structValue) {
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
