package chalk

import (
	"encoding/json"
	"fmt"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

func convertIfHasManyStruct(values any) (any, error) {
	// When the user passes in a list of has-many structs,
	// it gets serialized into:
	//
	// {
	//     "FullName": "John Doe",
	//     "Amount": 100,
	// }
	//
	// when we really want:
	//
	// {
	//     "user.full_name": "John Doe",
	//     "user.amount": 100,
	// }
	rValues := reflect.ValueOf(values)
	if rValues.Type().Kind() != reflect.Slice || rValues.Len() == 0 {
		return values, nil
	}
	elemType := rValues.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		// Not a dataclass nor a has-many feature.
		return values, nil
	}

	if elemType.NumField() > 0 && internal.IsTypeDataclass(elemType.Field(0).Type) {
		// Don't manually serialize dataclasses. Dataclasses need to be serialized
		// with JSON since we utilize struct tags to assign the original python
		// field name.
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
		fieldNameToPythonName[elemType.Field(i).Name] = fmt.Sprintf("%s.%s", namespace, pythonName)
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
		Environment:          internal.StringOrNil(p.EnvironmentId),
		Tags:                 p.Tags,
		RequiredResolverTags: p.RequiredResolverTags,
	}

	var now *string
	if len(p.Now) > 1 {
		p.builderErrors = append(
			p.builderErrors,
			// HACK: BuilderError should just go away and be replaced by just a list of errors.
			//       So we're just slapping together a fake BuilderError here using existing
			//       params.
			&BuilderError{
				Err: errors.Newf(
					"for non-bulk queries, there should only"+
						" be 1 `Now` value, found %d", len(p.Now),
				),
				Type:      InvalidRequest,
				Feature:   "Now",
				Value:     p.Now,
				ParamType: ParamInput,
			},
		)
	} else if len(p.Now) == 1 {
		now = lo.ToPtr(p.Now[0].Format(internal.NowTimeFormat))
	}

	convertedInputs := make(map[string]any)
	for fqn, values := range p.inputs {
		convertedValues, err := convertIfHasManyStruct(values)
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
		IncludeMeta:      p.IncludeMeta || p.Explain,
		IncludeMetrics:   p.IncludeMetrics,
		DeploymentId:     internal.StringOrNil(p.PreviewDeploymentId),
		QueryName:        internal.StringOrNil(p.QueryName),
		QueryNameVersion: internal.StringOrNil(p.QueryNameVersion),
		CorrelationId:    internal.StringOrNil(p.CorrelationId),
		Meta:             p.Meta,
		StorePlanStages:  p.StorePlanStages,
		Now:              now,
		Explain:          p.Explain,
	}, nil
}

func serializeStaleness(staleness map[string]time.Duration) map[string]string {
	res := map[string]string{}
	for k, v := range staleness {
		res[k] = internal.FormatBucketDuration(int(v.Seconds()))
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

func convertOnlineQueryParamsToProto(params *OnlineQueryParams) (*commonv1.OnlineQueryBulkRequest, error) {
	inputsFeather, err := internal.InputsToArrowBytes(params.inputs)
	if err != nil {
		return nil, errors.Wrap(err, "error serializing inputs as feather")
	}
	outputs := lo.Map(params.outputs, func(v string, _ int) *commonv1.OutputExpr {
		return &commonv1.OutputExpr{
			Expr: &commonv1.OutputExpr_FeatureFqn{
				FeatureFqn: v,
			},
		}
	})
	staleness := lo.MapValues(params.staleness, func(v time.Duration, k string) string {
		return internal.FormatBucketDuration(int(v.Seconds()))
	})

	nowProto := lo.Map(params.Now, func(v time.Time, _ int) *timestamppb.Timestamp {
		return timestamppb.New(v)
	})

	options := map[string]*structpb.Value{}
	if params.StorePlanStages {
		options["store_plan_stages"] = structpb.NewBoolValue(params.StorePlanStages)
	}
	if params.IncludeMetrics {
		options["include_metrics"] = structpb.NewBoolValue(params.IncludeMetrics)
	}

	return &commonv1.OnlineQueryBulkRequest{
		InputsFeather: inputsFeather,
		Outputs:       outputs,
		Staleness:     staleness,
		Now:           lo.Ternary(len(nowProto) == 0, nil, nowProto),
		Context: &commonv1.OnlineQueryContext{
			Environment:          params.EnvironmentId,
			Tags:                 params.Tags,
			DeploymentId:         lo.EmptyableToPtr(params.PreviewDeploymentId),
			BranchId:             params.BranchId,
			CorrelationId:        lo.EmptyableToPtr(params.CorrelationId),
			QueryName:            lo.EmptyableToPtr(params.QueryName),
			QueryNameVersion:     lo.EmptyableToPtr(params.QueryNameVersion),
			RequiredResolverTags: params.RequiredResolverTags,
			Options:              options,
		},
		ResponseOptions: &commonv1.OnlineQueryResponseOptions{
			IncludeMeta:     params.IncludeMeta || params.Explain,
			Metadata:        params.Meta,
			EncodingOptions: nil,
			Explain:         lo.Ternary(params.Explain, &commonv1.ExplainOptions{}, nil),
		},
	}, nil
}
