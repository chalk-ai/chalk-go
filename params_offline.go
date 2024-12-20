package chalk

import (
	"time"
)

/*****************************************
 Definitions for OfflineQueryParamsComplete
******************************************/

// OfflineQueryParamsComplete is the only type of object
// accepted as an argument to Client.OfflineQuery.
// OfflineQueryParamsComplete is obtained by calling a chain
// of methods starting with any method of [OfflineQueryParams].
//
// Example:
//
//	     defaultObservedAt := time.Now().Add(-time.Hour)
//			observedAt, _ := time.Parse(time.RFC822, "02 Jan 22 15:04 PST")
//			client.OfflineQuery(
//				OfflineQueryParams{
//					EnvironmentId: "pipkjlfc3gtmn",
//				}.
//		 		WithInput(Features.User.Id, []any{1, chalk.TsFeatureValue{Value: 2, ObservationTime: &observedAt}}).
//		 		WithRequiredOutputs(Features.User.Email, Features.User.Card.Id),
//			)
//
// It is mandatory to call [OfflineQueryParams.WithOutput]
// or [OfflineQueryParams.WithRequiredOutputs] at least once
// for OfflineQueryParamsComplete to be returned.
// Otherwise, an incomplete type will be returned, and it cannot
// be passed into Client.OfflineQuery.
type OfflineQueryParamsComplete struct {
	underlying OfflineQueryParams
}

// WithInput returns a copy of Offline Query parameters with the specified input added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithInput(feature any, values []any) OfflineQueryParamsComplete {
	p.underlying = p.underlying.withInput(feature, values)
	return p
}

// WithOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithOutputs(features ...any) OfflineQueryParamsComplete {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithRequiredOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithRequiredOutputs(features ...any) OfflineQueryParamsComplete {
	p.underlying = p.underlying.withRequiredOutputs(features...)
	return p
}

/********************************************
 Definitions for offlineQueryParamsWithInputs
*********************************************/

type offlineQueryParamsWithInputs struct {
	underlying OfflineQueryParams
}

// WithInput returns a copy of Offline Query parameters with the specified input added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p offlineQueryParamsWithInputs) WithInput(feature any, values []any) offlineQueryParamsWithInputs {
	p.underlying = p.underlying.withInput(feature, values)
	return p
}

// WithOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p offlineQueryParamsWithInputs) WithOutputs(features ...any) OfflineQueryParamsComplete {
	return OfflineQueryParamsComplete{p.underlying.withOutputs(features...)}
}

// WithRequiredOutputs returns a copy of Offline Query parameters with the specified outputs added.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p offlineQueryParamsWithInputs) WithRequiredOutputs(features ...any) OfflineQueryParamsComplete {
	return OfflineQueryParamsComplete{p.underlying.withRequiredOutputs(features...)}
}

/***********************************
 Definitions for OfflineQueryParams
***********************************/

func (p OfflineQueryParams) withInput(feature any, values []any) OfflineQueryParams {
	validateErr := validateFeature(feature, ParamInput)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}

	timestampedValues := p.getTimestampedFeatures(values)

	if p.inputs == nil {
		p.inputs = make(map[string][]TsFeatureValue)
	}

	var key string
	if fqn, ok := feature.(string); ok {
		key = fqn
	} else {
		castedFeature, castErr := UnwrapFeature(feature)
		if castErr != nil {
			builderError := BuilderError{
				Err:       castErr,
				Type:      UnwrapFeatureError,
				Feature:   feature,
				Value:     values,
				ParamType: ParamInput,
			}
			p.builderErrors = append(p.builderErrors, &builderError)
			return p
		}
		p.versioned = true
		key = castedFeature.Fqn
	}
	p.inputs[key] = append(p.inputs[key], timestampedValues...)
	return p
}

func (p OfflineQueryParams) withInputs(inputs map[any][]any) OfflineQueryParams {
	for key, values := range inputs {
		p = p.withInput(key, values)
	}
	return p
}

func (p OfflineQueryParams) withOutputs(features ...any) OfflineQueryParams {
	validateErr := validateFeatures(features, ParamOutput)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}

	for _, feature := range features {
		var key string
		if fqn, ok := feature.(string); ok {
			key = fqn
		} else {
			castedFeature, castErr := UnwrapFeature(feature)
			if castErr != nil {
				builderError := BuilderError{
					Err:       castErr,
					Type:      UnwrapFeatureError,
					Feature:   feature,
					ParamType: ParamOutput,
				}
				p.builderErrors = append(p.builderErrors, &builderError)
				return p
			}
			p.versioned = true
			key = castedFeature.Fqn
		}
		p.outputs = append(p.outputs, key)
	}
	return p
}

func (p OfflineQueryParams) withRequiredOutputs(features ...any) OfflineQueryParams {
	validateErr := validateFeatures(features, ParamRequiredOutput)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}

	for _, feature := range features {
		var key string
		if fqn, ok := feature.(string); ok {
			key = fqn
		} else {
			castedFeature, castErr := UnwrapFeature(feature)
			if castErr != nil {
				builderError := BuilderError{
					Err:       castErr,
					Type:      UnwrapFeatureError,
					Feature:   feature,
					ParamType: ParamRequiredOutput,
				}
				p.builderErrors = append(p.builderErrors, &builderError)
				return p
			}
			p.versioned = true
			key = castedFeature.Fqn
		}
		p.requiredOutputs = append(p.requiredOutputs, key)
	}
	return p
}

func (p OfflineQueryParams) getTimestampedFeatures(values []any) []TsFeatureValue {
	castedValues := make([]TsFeatureValue, 0)
	localTz, err := time.LoadLocation("Local")
	if err != nil {
		localTz, _ = time.LoadLocation("UTC")
	}
	for _, value := range values {
		var castedVal TsFeatureValue
		if tsFeature, ok := value.(TsFeatureValue); ok && tsFeature.ObservationTime != nil {
			castedVal = tsFeature
			if castedVal.ObservationTime.Location() == nil {
				localTime := castedVal.ObservationTime.In(localTz)
				castedVal.ObservationTime = &localTime
			}
		} else {
			observedAt := time.Now()
			if p.DefaultTime != nil {
				observedAt = *p.DefaultTime
			}
			castedVal = TsFeatureValue{Value: value, ObservationTime: &observedAt}
		}
		utcTime := castedVal.ObservationTime.UTC()
		castedVal.ObservationTime = &utcTime
		castedValues = append(castedValues, castedVal)
	}
	return castedValues
}

func (p OfflineQueryParams) withQueryContext(queryContext *QueryContext) OfflineQueryParams {
	p.QueryContext = queryContext
	return p
}
