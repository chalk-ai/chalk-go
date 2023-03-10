package client

import (
	"github.com/chalk-ai/chalk-go/pkg/enum"
	"strconv"
	"time"
)

func (request *OnlineQueryParams) serialize() onlineQueryRequestSerialized {
	context := onlineQueryContext{
		Environment: stringPointerOrNil(request.EnvironmentId),
		Tags:        request.Tags,
	}

	body := onlineQueryRequestSerialized{
		Inputs:         request.Inputs,
		Outputs:        request.Outputs,
		Context:        context,
		Staleness:      deserializeStaleness(request.Staleness),
		IncludeMeta:    request.IncludeMeta,
		IncludeMetrics: request.IncludeMetrics,
		DeploymentId:   stringPointerOrNil(request.DeploymentId),
		QueryName:      stringPointerOrNil(request.QueryName),
		CorrelationId:  stringPointerOrNil(request.CorrelationId),
		Meta:           request.Meta,
	}

	return body
}

func deserializeStaleness(staleness map[string]time.Duration) map[string]string {
	res := map[string]string{}
	for k, v := range staleness {
		res[k] = strconv.Itoa(int(v.Seconds())) + "s"
	}
	return res
}

func (feature featureResultSerialized) deserialize() (FeatureResult, error) {
	timeObj, err := time.Parse(time.RFC3339, feature.Timestamp)
	if err != nil {
		return FeatureResult{}, err
	}

	var dError *ChalkServerError = nil
	if feature.Error != nil {
		dErrorObj, err := feature.Error.deserialize()
		if err != nil {
			return FeatureResult{}, err
		}
		dError = &dErrorObj
	}

	return FeatureResult{
		Field:     feature.Field,
		Value:     feature.Value,
		Timestamp: timeObj,
		Meta:      feature.Meta,
		Error:     dError,
	}, nil
}

func (feature FeatureResult) serialize() (featureResultSerialized, error) {
	sError, err := feature.Error.serialize()
	if err != nil {
		return featureResultSerialized{}, err
	}

	return featureResultSerialized{
		Field:     feature.Field,
		Value:     feature.Value,
		Timestamp: feature.Timestamp.String(),
		Meta:      feature.Meta,
		Error:     &sError,
	}, nil
}

func (response *onlineQueryResponseSerialized) deserialize() (OnlineQueryResult, error) {
	features := make(map[string]FeatureResult)

	deserializedData, err := deserializeFeatureResults(response.Data)
	if err != nil {
		return OnlineQueryResult{}, err
	}

	for _, result := range deserializedData {
		features[result.Field] = result
	}

	return OnlineQueryResult{
		Data:     deserializedData,
		Meta:     response.Meta,
		features: features,
	}, nil
}

func deserializeFeatureResults(results []featureResultSerialized) ([]FeatureResult, *ChalkClientError) {
	deserializedResults := make([]FeatureResult, 0)
	for _, sResult := range results {
		dResult, dErr := sResult.deserialize()
		if dErr != nil {
			return []FeatureResult{}, &ChalkClientError{
				Message: dErr.Error(),
			}
		}
		deserializedResults = append(deserializedResults, dResult)

	}
	return deserializedResults, nil
}

func (e *ChalkServerError) serialize() (chalkErrorSerialized, error) {
	return chalkErrorSerialized{
		Code:      e.Code.Value,
		Category:  e.Category.Value,
		Message:   e.Message,
		Exception: e.Exception,
		Feature:   e.Feature,
		Resolver:  e.Resolver,
	}, nil
}

func (e *chalkErrorSerialized) deserialize() (ChalkServerError, error) {
	errorCode, getErrorCodeErr := enum.GetErrorCode(e.Code)
	if getErrorCodeErr != nil {
		return ChalkServerError{}, getErrorCodeErr
	}

	errorCodeCategory, getCategoryErr := enum.GetErrorCodeCategory(e.Category)
	if getCategoryErr != nil {
		return ChalkServerError{}, getCategoryErr
	}

	return ChalkServerError{
		Code:      *errorCode,
		Category:  *errorCodeCategory,
		Message:   e.Message,
		Exception: e.Exception,
		Feature:   e.Feature,
		Resolver:  e.Resolver,
	}, nil
}

func deserializeChalkErrors(errors []chalkErrorSerialized) ([]ChalkServerError, *ChalkClientError) {
	deserializedErrors := make([]ChalkServerError, 0)
	for _, serializedErr := range errors {
		deserializedError, deserializationFailure := serializedErr.deserialize()
		if deserializationFailure != nil {
			return []ChalkServerError{}, &ChalkClientError{
				Message: deserializationFailure.Error(),
			}
		}
		deserializedErrors = append(deserializedErrors, deserializedError)

	}
	return deserializedErrors, nil
}
