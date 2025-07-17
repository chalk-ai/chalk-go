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
//				context.Background(),
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

// WithRunAsynchronously returns a copy of Offline Query parameters with RunAsynchronously set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithRunAsynchronously(runAsynchronously bool) OfflineQueryParamsComplete {
	p.underlying.RunAsynchronously = runAsynchronously
	return p
}

// WithNumShards returns a copy of Offline Query parameters with NumShards set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithNumShards(numShards int) OfflineQueryParamsComplete {
	p.underlying.NumShards = &numShards
	return p
}

// WithNumWorkers returns a copy of Offline Query parameters with NumWorkers set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithNumWorkers(numWorkers int) OfflineQueryParamsComplete {
	p.underlying.NumWorkers = &numWorkers
	return p
}

// WithResources returns a copy of Offline Query parameters with Resources set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithResources(resources *ResourceRequests) OfflineQueryParamsComplete {
	p.underlying.Resources = resources
	return p
}

// WithCompletionDeadline returns a copy of Offline Query parameters with CompletionDeadline set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithCompletionDeadline(deadline time.Duration) OfflineQueryParamsComplete {
	p.underlying.CompletionDeadline = &deadline
	return p
}

// WithMaxRetries returns a copy of Offline Query parameters with MaxRetries set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithMaxRetries(maxRetries int) OfflineQueryParamsComplete {
	p.underlying.MaxRetries = &maxRetries
	return p
}

// WithStoreOnline returns a copy of Offline Query parameters with StoreOnline set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithStoreOnline(storeOnline bool) OfflineQueryParamsComplete {
	p.underlying.StoreOnline = storeOnline
	return p
}

// WithStoreOffline returns a copy of Offline Query parameters with StoreOffline set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithStoreOffline(storeOffline bool) OfflineQueryParamsComplete {
	p.underlying.StoreOffline = storeOffline
	return p
}

// WithUseMultipleComputers returns a copy of Offline Query parameters with UseMultipleComputers set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithUseMultipleComputers(useMultipleComputers bool) OfflineQueryParamsComplete {
	p.underlying.UseMultipleComputers = useMultipleComputers
	return p
}

// WithUploadInputAsTable returns a copy of Offline Query parameters with UploadInputAsTable set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithUploadInputAsTable(uploadInputAsTable bool) OfflineQueryParamsComplete {
	p.underlying.UploadInputAsTable = uploadInputAsTable
	return p
}

// WithEnvOverrides returns a copy of Offline Query parameters with EnvOverrides set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithEnvOverrides(envOverrides map[string]string) OfflineQueryParamsComplete {
	p.underlying.EnvOverrides = envOverrides
	return p
}

// WithEnableProfiling returns a copy of Offline Query parameters with EnableProfiling set.
// For use via method chaining. See OfflineQueryParamsComplete for usage examples.
func (p OfflineQueryParamsComplete) WithEnableProfiling(enableProfiling bool) OfflineQueryParamsComplete {
	p.underlying.EnableProfiling = enableProfiling
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
	timestampedValues := p.getTimestampedFeatures(values)
	if p.rawInputs == nil {
		p.rawInputs = make(map[any][]TsFeatureValue)
	}
	p.rawInputs[feature] = append(p.rawInputs[feature], timestampedValues...)
	return p
}

func (p OfflineQueryParams) withInputs(inputs map[any][]any) OfflineQueryParams {
	for key, values := range inputs {
		p = p.withInput(key, values)
	}
	return p
}

func (p OfflineQueryParams) withOutputs(features ...any) OfflineQueryParams {
	p.rawOutputs = append(p.rawOutputs, features...)
	return p
}

func (p OfflineQueryParams) withRequiredOutputs(features ...any) OfflineQueryParams {
	p.rawRequiredOutputs = append(p.rawRequiredOutputs, features...)
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
