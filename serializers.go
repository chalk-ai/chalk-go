package chalk

import (
	"encoding/json"
	"github.com/chalk-ai/chalk-go/internal"
	"strconv"
	"time"
)

func (p OnlineQueryParams) serialize() onlineQueryRequestSerialized {
	context := onlineQueryContext{
		Environment: internal.StringOrNil(p.EnvironmentId),
		Tags:        p.Tags,
	}

	body := onlineQueryRequestSerialized{
		Inputs:         p.inputs,
		Outputs:        p.outputs,
		Context:        context,
		Staleness:      serializeStaleness(p.staleness),
		IncludeMeta:    p.IncludeMeta,
		IncludeMetrics: p.IncludeMetrics,
		DeploymentId:   internal.StringOrNil(p.PreviewDeploymentId),
		QueryName:      internal.StringOrNil(p.QueryName),
		CorrelationId:  internal.StringOrNil(p.CorrelationId),
		Meta:           p.Meta,
	}

	return body
}

func serializeStaleness(staleness map[string]time.Duration) map[string]string {
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

	var dError *ServerError = nil
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
		Pkey:      feature.Pkey,
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
		Pkey:      feature.Pkey,
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

func deserializeFeatureResults(results []featureResultSerialized) ([]FeatureResult, error) {
	deserializedResults := make([]FeatureResult, 0)
	for _, sResult := range results {
		dResult, dErr := sResult.deserialize()
		if dErr != nil {
			return []FeatureResult{}, &ClientError{
				Message: dErr.Error(),
			}
		}
		deserializedResults = append(deserializedResults, dResult)

	}
	return deserializedResults, nil
}

func (e *ServerError) serialize() (chalkErrorSerialized, error) {
	return chalkErrorSerialized{
		Code:      e.Code.Value,
		Category:  e.Category.Value,
		Message:   e.Message,
		Exception: e.Exception,
		Feature:   e.Feature,
		Resolver:  e.Resolver,
	}, nil
}

func (e *ErrorCode) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	errorCode, getErrorCodeErr := getErrorCode(str)
	if getErrorCodeErr != nil {
		return getErrorCodeErr
	}

	*e = *errorCode
	return nil
}

func (c *ErrorCodeCategory) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	errorCodeCategory, getCategoryErr := getErrorCodeCategory(str)
	if getCategoryErr != nil {
		return getCategoryErr
	}

	*c = *errorCodeCategory
	return nil
}

func (p OfflineQueryParams) MarshalJSON() ([]byte, error) {
	queryInput := offlineQueryInputSerialized{}
	for fqn, values := range p.Inputs {
		queryInput.Columns = append(queryInput.Columns, fqn)
		queryInput.Values = append(queryInput.Values, values)
	}

	requiredOutput := p.Output
	if requiredOutput == nil {
		requiredOutput = p.Output
	}

	return json.Marshal(offlineQueryRequestSerialized{
		Input:             queryInput,
		Output:            p.Output,
		RequiredOutput:    requiredOutput,
		DatasetName:       internal.StringOrNil(p.DatasetName),
		Branch:            internal.StringOrNil(p.Branch),
		MaxSamples:        p.MaxSamples,
		DestinationFormat: "PARQUET",
	})
}

//func (e ServerError) UnmarshalJSON(data []byte) error {
//	c := chalkErrorSerialized{}
//	unmarshalErr := json.Unmarshal(data, &c)
//	if unmarshalErr != nil {
//		return unmarshalErr
//	}
//
//	errorCode, getErrorCodeErr := getErrorCode(c.Code)
//	if getErrorCodeErr != nil {
//		return getErrorCodeErr
//	}
//
//	errorCodeCategory, getCategoryErr := getErrorCodeCategory(c.Category)
//	if getCategoryErr != nil {
//		return getCategoryErr
//	}
//
//	e.Code = *errorCode
//	e.Category = *errorCodeCategory
//	e.Message = c.Message
//	e.Exception = c.Exception
//	e.Feature = c.Feature
//	e.Resolver = c.Resolver
//
//	return nil
//}

func (e *chalkErrorSerialized) deserialize() (ServerError, error) {
	errorCode, getErrorCodeErr := getErrorCode(e.Code)
	if getErrorCodeErr != nil {
		return ServerError{}, getErrorCodeErr
	}

	errorCodeCategory, getCategoryErr := getErrorCodeCategory(e.Category)
	if getCategoryErr != nil {
		return ServerError{}, getCategoryErr
	}

	return ServerError{
		Code:      *errorCode,
		Category:  *errorCodeCategory,
		Message:   e.Message,
		Exception: e.Exception,
		Feature:   e.Feature,
		Resolver:  e.Resolver,
	}, nil
}

func deserializeChalkErrors(errors []chalkErrorSerialized) ([]ServerError, error) {
	deserializedErrors := make([]ServerError, 0)
	for _, serializedErr := range errors {
		deserializedError, deserializationFailure := serializedErr.deserialize()
		if deserializationFailure != nil {
			return []ServerError{}, &ClientError{
				Message: deserializationFailure.Error(),
			}
		}
		deserializedErrors = append(deserializedErrors, deserializedError)

	}
	return deserializedErrors, nil
}

var getErrorCode = internal.GenerateGetEnumFunction(
	map[string]ErrorCode{
		ParseFailed.Value:         ParseFailed,
		ResolverTimedOut.Value:    ResolverTimedOut,
		ResolverNotFound.Value:    ResolverNotFound,
		InvalidQuery.Value:        InvalidQuery,
		ValidationFailed.Value:    ValidationFailed,
		ResolverFailed.Value:      ResolverFailed,
		UpstreamFailed.Value:      UpstreamFailed,
		Unauthenticated.Value:     Unauthenticated,
		Unauthorized.Value:        Unauthorized,
		InternalServerError.Value: InternalServerError,
		Cancelled.Value:           Cancelled,
		DeadlineExceeded.Value:    DeadlineExceeded,
	},
	"error codes",
)

var getErrorCodeCategory = internal.GenerateGetEnumFunction(
	map[string]ErrorCodeCategory{
		Request.Value: Request,
		Field.Value:   Field,
		Network.Value: Network,
	},
	"error code categories",
)
