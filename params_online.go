package chalk

import (
	"time"

	"github.com/chalk-ai/chalk-go/expr"
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
//			context.Background(),
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

// WithQueryName returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithQueryName(queryName string) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withQueryName(queryName)
	return p
}

// WithQueryNameVersion returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithQueryNameVersion(queryNameVersion string) OnlineQueryParamsComplete {
	p.underlying = p.underlying.WithQueryNameVersion(queryNameVersion)
	return p
}

// WithInput returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithInput(feature any, value any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

// WithInputs returns a copy of Online Query parameters with the specified inputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithInputs(inputs map[any]any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withInputs(inputs)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithOutputs(features ...any) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithOutputExprs returns a copy of Online Query parameters with the specified output expressions added.
func (p OnlineQueryParamsComplete) WithOutputExprs(exprs ...expr.ExprI) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withOutputExprs(exprs...)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p OnlineQueryParamsComplete) WithStaleness(feature any, duration time.Duration) OnlineQueryParamsComplete {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

// WithTags returns a copy of Online Query parameters with the specified tags added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p OnlineQueryParamsComplete) WithTags(tags ...string) OnlineQueryParamsComplete {
	p.underlying = p.underlying.WithTags(tags...)
	return p
}

/*****************************************
 Definitions for OnlineQueryParams
******************************************/

func (p OnlineQueryParams) withInput(feature any, value any) (result OnlineQueryParams) {
	if p.rawInputs == nil {
		p.rawInputs = make(map[any]any)
	}
	p.rawInputs[feature] = value
	return p
}

func (p OnlineQueryParams) withInputs(inputs map[any]any) OnlineQueryParams {
	for key, value := range inputs {
		p = p.withInput(key, value)
	}
	return p
}

func (p OnlineQueryParams) withOutputs(features ...any) OnlineQueryParams {
	p.rawOutputs = append(p.rawOutputs, features...)
	return p
}

func (p OnlineQueryParams) withQueryName(name string) OnlineQueryParams {
	p.QueryName = name
	return p
}

func (p OnlineQueryParams) withStaleness(feature any, duration time.Duration) OnlineQueryParams {
	if p.rawStaleness == nil {
		p.rawStaleness = make(map[any]time.Duration)
	}
	p.rawStaleness[feature] = duration
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

// WithInputs returns a copy of Online Query parameters with the specified inputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithInputs(inputs map[any]any) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withInputs(inputs)
	return p
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithOutputs(features ...any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withOutputs(features...)}
}

func (p onlineQueryParamsWithInputs) WithOutputExprs(exprs ...expr.ExprI) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withOutputExprs(exprs...)}
}

// WithQueryName returns a copy of Online Query parameters with the specified query name set.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithInputs) WithQueryName(name string) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withQueryName(name)}
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithInputs) WithStaleness(feature any, duration time.Duration) onlineQueryParamsWithInputs {
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
func (p onlineQueryParamsWithOutputs) WithInput(feature any, value any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withInput(feature, value)}
}

// WithInputs returns a copy of Online Query parameters with the specified input added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithInputs(inputs map[any]any) OnlineQueryParamsComplete {
	return OnlineQueryParamsComplete{p.underlying.withInputs(inputs)}
}

// WithOutputs returns a copy of Online Query parameters with the specified outputs added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithOutputs(features ...any) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

// WithQueryName returns a copy of Online Query parameters with the specified query name set.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
func (p onlineQueryParamsWithOutputs) WithQueryName(name string) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withQueryName(name)
	return p
}

// WithStaleness returns a copy of Online Query parameters with the specified staleness added.
// For use via method chaining. See OnlineQueryParamsComplete for usage examples.
// See https://docs.chalk.ai/docs/query-caching for more information on staleness.
func (p onlineQueryParamsWithOutputs) WithStaleness(feature any, duration time.Duration) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}
