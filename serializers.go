package chalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strconv"
	"time"
)

func getConvertedValues(values any) (any, error) {
	// Do preprocessing on values such as prefix each key in
	// the list of has-many features with the namespace.
	rValues := reflect.ValueOf(values)
	if rValues.Type().Kind() != reflect.Slice || rValues.Len() == 0 {
		return values, nil
	}
	elemType := rValues.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		// Not a dataclass nor a has-many feature.
		return values, nil
	}
	// This is a list of dataclasses, or a has-many list of features.
	fieldNameToPythonName := make(map[string]string)
	namespace := internal.ChalkpySnakeCase(elemType.Name())
	for i := 0; i < elemType.NumField(); i++ {
		pythonName, err := internal.ResolveFeatureName(elemType.Field(i))
		if err != nil {
			return nil, errors.New("failed to resolve field name")
		}
		if !internal.IsTypeDataclass(elemType.Field(i).Type) {
			// Has-many feature. Prepend namespace.
			pythonName = fmt.Sprintf("%s.%s", namespace, pythonName)
		}
		fieldNameToPythonName[elemType.Field(i).Name] = pythonName
	}

	newValues := make([]map[string]any, rValues.Len())

	for i := 0; i < rValues.Len(); i++ {
		newMap := make(map[string]any)
		newValues[i] = newMap
		oldValue := rValues.Index(i)
		for j := 0; j < elemType.NumField(); j++ {
			pythonName := fieldNameToPythonName[elemType.Field(j).Name]
			newMap[pythonName] = oldValue.Field(j).Interface()
		}
	}

	return newValues, nil
}

func (p OnlineQueryParams) serialize() (*internal.OnlineQueryRequestSerialized, error) {
	context := internal.OnlineQueryContext{
		Environment: internal.StringOrNil(p.EnvironmentId),
		Tags:        p.Tags,
	}

	convertedInputs := make(map[string]any)
	for fqn, values := range p.inputs {
		convertedValues, err := getConvertedValues(values)
		if err != nil {
			return nil, wrapClientError(err, "failed to preprocess input values")
		}
		convertedInputs[fqn] = convertedValues
	}

	return &internal.OnlineQueryRequestSerialized{
		Inputs:           convertedInputs,
		Outputs:          p.outputs,
		Context:          context,
		Staleness:        serializeStaleness(p.staleness),
		IncludeMeta:      p.IncludeMeta,
		IncludeMetrics:   p.IncludeMetrics,
		DeploymentId:     internal.StringOrNil(p.PreviewDeploymentId),
		QueryName:        internal.StringOrNil(p.QueryName),
		QueryNameVersion: internal.StringOrNil(p.QueryNameVersion),
		CorrelationId:    internal.StringOrNil(p.CorrelationId),
		Meta:             p.Meta,
	}, nil
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
	queryInput := internal.OfflineQueryInputSerialized{}
	globalInputTimes := make([]any, 0)

	for fqn, tsFeatureValues := range p.inputs {
		var inputValues []any
		var inputTimes []any
		for _, v := range tsFeatureValues {
			inputTimes = append(inputTimes, v.ObservationTime.Format(time.RFC3339))
			inputValues = append(inputValues, v.Value)
		}
		queryInput.Columns = append(queryInput.Columns, fqn)
		queryInput.Values = append(queryInput.Values, inputValues)
		globalInputTimes = inputTimes
	}

	queryInput.Columns = append(queryInput.Columns, "__chalk__.CHALK_TS")
	queryInput.Values = append(queryInput.Values, globalInputTimes)

	output := p.outputs
	if output == nil {
		output = make([]string, 0)
	}

	requiredOutput := p.requiredOutputs
	if requiredOutput == nil {
		requiredOutput = make([]string, 0)
	}

	serializedObj := internal.OfflineQueryRequestSerialized{
		Input:             queryInput,
		Output:            output,
		RequiredOutput:    requiredOutput,
		DatasetName:       internal.StringOrNil(p.DatasetName),
		Branch:            internal.StringOrNil(p.Branch),
		MaxSamples:        p.MaxSamples,
		DestinationFormat: "PARQUET",
		Tags:              p.Tags,
	}

	return json.Marshal(serializedObj)
}

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
