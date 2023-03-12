package chalk

import "time"

func (p OnlineQueryParams) withInput(feature string, value any) OnlineQueryParams {
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

type onlineQueryParamsWithInputs struct {
	underlying OnlineQueryParams
}

type onlineQueryParamsWithOutputs struct {
	underlying OnlineQueryParams
}

type onlineQueryParamsComplete struct {
	underlying OnlineQueryParams
}

func (p OnlineQueryParams) WithInput(feature string, value any) onlineQueryParamsWithInputs {
	return onlineQueryParamsWithInputs{underlying: p.withInput(feature, value)}
}
func (p OnlineQueryParams) WithOutputs(features ...string) onlineQueryParamsWithOutputs {
	return onlineQueryParamsWithOutputs{underlying: p.withOutputs(features...)}
}
func (p OnlineQueryParams) WithStaleness(feature string, duration time.Duration) OnlineQueryParams {
	return p.withStaleness(feature, duration)
}

func (p onlineQueryParamsWithInputs) WithInput(feature string, value any) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withInput(feature, value)
	return p
}

func (p onlineQueryParamsWithInputs) WithOutputs(features ...string) onlineQueryParamsComplete {
	return onlineQueryParamsComplete{p.underlying.withOutputs(features...)}
}

func (p onlineQueryParamsWithInputs) WithStaleness(feature string, duration time.Duration) onlineQueryParamsWithInputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

func (p onlineQueryParamsWithOutputs) WithInput(feature string, value any) onlineQueryParamsComplete {
	return onlineQueryParamsComplete{p.underlying.withInput(feature, value)}
}

func (p onlineQueryParamsWithOutputs) WithOutputs(features ...string) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withOutputs(features...)
	return p
}

func (p onlineQueryParamsWithOutputs) WithStaleness(feature string, duration time.Duration) onlineQueryParamsWithOutputs {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}

func (p onlineQueryParamsComplete) WithStaleness(feature string, duration time.Duration) onlineQueryParamsComplete {
	p.underlying = p.underlying.withStaleness(feature, duration)
	return p
}
