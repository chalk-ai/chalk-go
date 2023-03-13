package chalk

import "time"

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
func (p OnlineQueryParamsComplete) WithInput(feature string, value any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithOutputs(features ...string) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p OnlineQueryParamsComplete) WithStaleness(feature string, duration time.Duration) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

func (p OnlineQueryParams) withInput(feature string, value any) OnlineQueryParams {
	if p.inputs == nil {
		p.inputs = make(map[string]any)
	}
	p.inputs[feature] = value
	return p
}

func (p OnlineQueryParams) withOutputs(features ...string) OnlineQueryParams {
	for _, feature := range features {
		p.outputs = append(p.outputs, feature)
	}
	return p
}

func (p OnlineQueryParams) withStaleness(feature string, duration time.Duration) OnlineQueryParams {
	if p.staleness == nil {
		p.staleness = make(map[string]time.Duration)
	}
	p.staleness[feature] = duration
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
func (p onlineQueryParamsWithInputs) WithInput(feature string, value any) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithOutputs(features ...string) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withOutputs(features...)}
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithInputs) WithStaleness(feature string, duration time.Duration) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
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
func (p onlineQueryParamsWithOutputs) WithInput(feature string, value any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withInput(feature, value)}
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method git st. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithOutputs(features ...string) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithOutputs) WithStaleness(feature string, duration time.Duration) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}
