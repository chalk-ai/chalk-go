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
func (p OnlineQueryParamsComplete) WithStaleness(feature string, duration time.Duration) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

/********************************************
 Definitions for OnlineQueryParams
*********************************************/

// WithInput returns a copy of Online Query parameters with the specified inputs added.
// For use via method chaining. See [OnlineQueryParamsComplete] for usage examples.
func (p OnlineQueryParams) WithInput(feature string, value any) onlineQueryParamsWithInputs {
	return onlineQueryParamsWithInputs{underlying: p.withInput(feature, value)}
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParams) WithOutputs(features ...string) onlineQueryParamsWithOutputs {
	return onlineQueryParamsWithOutputs{underlying: p.withOutputs(features...)}
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParams) WithStaleness(feature string, duration time.Duration) OnlineQueryParams {
	return p.withStaleness(feature, duration)
}

func (p OnlineQueryParams) withInput(feature string, value any) OnlineQueryParams {
	if p.Inputs == nil {
		p.Inputs = make(map[string]any)
	}
	p.Inputs[feature] = value
	return p
}

func (p OnlineQueryParams) withOutputs(features ...string) OnlineQueryParams {
	for _, feature := range features {
		p.Outputs = append(p.Outputs, feature)
	}
	return p
}

func (p OnlineQueryParams) withStaleness(feature string, duration time.Duration) OnlineQueryParams {
	p.Staleness[feature] = duration
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
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithOutputs(features ...string) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithStaleness(feature string, duration time.Duration) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}
