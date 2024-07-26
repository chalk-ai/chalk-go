package chalk

import (
	"time"
)

/*****************************************
 Definitions for OnlineQueryParamsComplete
******************************************/

// OnlineQueryParamsComplete is the only type of object
// accepted as an argument to Client.OnlineQuery.
// OnlineQueryParamsComplete is obtained by calling a chain
// of methods starting with any method of [OnlineQueryParams].
//
// Example:
//
//		client.OnlineQuery(
//			OnlineQueryParams{
//				IncludeMeta: true,
//				EnvironmentId: "pipkjlfc3gtmn",
//			}.
//	 		WithInput(Features.User.Card.Id, 4).
//	 		WithOutputs(Features.User.Email, Features.User.Card.Id),
//		)
//
// [OnlineQueryParams.WithInput] and [OnlineQueryParams.WithOutputs]
// are mandatory methods. This means they must each be called at
// least once for OnlineQueryParamsComplete to be returned.
// Otherwise, an incomplete type will be returned, and it cannot
// be passed into Client.OnlineQuery.
type OnlineQueryParamsComplete struct {
	underlying OnlineQueryParams
}

// WithInput returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithInput(feature any, value any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithOutputs(features ...any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p OnlineQueryParamsComplete) WithStaleness(feature any, duration time.Duration) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

func (p OnlineQueryParamsComplete) WithBranchId(branchId string) OnlineQueryParamsComplete {
	p.underlying = p.underlying.WithBranchId(branchId)
	return p
}

/*****************************************
 Definitions for OnlineQueryParams
******************************************/

func (p OnlineQueryParams) withInput(feature any, value any) (result OnlineQueryParams) {
	validateErr := validateFeature(feature, ParamInput)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}

	if p.inputs == nil {
		p.inputs = make(map[string]any)
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
				ParamType: ParamInput,
			}
			p.builderErrors = append(p.builderErrors, &builderError)
			return p
		}
		p.versioned = true
		key = castedFeature.Fqn
	}

	p.inputs[key] = value
	return p
}

func (p OnlineQueryParams) withOutputs(features ...any) OnlineQueryParams {
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

func (p OnlineQueryParams) withStaleness(feature any, duration time.Duration) OnlineQueryParams {
	validateErr := validateFeature(feature, ParamStaleness)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}

	if p.staleness == nil {
		p.staleness = make(map[string]time.Duration)
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
				Value:     duration,
				ParamType: ParamStaleness,
			}
			p.builderErrors = append(p.builderErrors, &builderError)
			return p
		}
		p.versioned = true
		key = castedFeature.Fqn
	}
	p.staleness[key] = duration
	return p
}

func (p OnlineQueryParams) withTags(tags []string) OnlineQueryParams {
	validateErr := validateTags(tags, ParamTags)
	if validateErr != nil {
		p.builderErrors = append(p.builderErrors, validateErr)
		return p
	}
	if p.Tags == nil {
		p.Tags = tags
	} else {
		p.Tags = append(p.Tags, tags...)
	}
	return p
}

/********************************************
 Definitions for onlineQueryParamsWithInputs
*********************************************/

type onlineQueryParamsWithInputs struct {
	underlying OnlineQueryParams
}

// WithInput returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithInput(feature any, value any) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithOutputs(features ...any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withOutputs(features...)}
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithInputs) WithStaleness(feature any, duration time.Duration) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

// WithTags returns a copy of Online Query parameters with the specified tags added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/resolver-tags for more information on tags.
func (p onlineQueryParamsWithInputs) WithTags(tags []string) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withTags(tags)
	return p
}

/********************************************
 Definitions for onlineQueryParamsWithOutputs
*********************************************/

type onlineQueryParamsWithOutputs struct {
	underlying OnlineQueryParams
}

// WithInput returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithInput(feature any, value any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withInput(feature, value)}
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method git st. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithOutputs(features ...any) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithOutputs) WithStaleness(feature any, duration time.Duration) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}
