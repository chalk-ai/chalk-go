package chalk

import (
	"encoding/json"
	"fmt"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (p OnlineQueryParams) serialize() (*internal.OnlineQueryRequestSerialized, error) {
	outputs := p.outputs
	if outputs == nil {
		// If we are passing query name, we don't need to pass outputs,
		// so outputs is empty, but when JSON serialized should never
		// be `null`.
		outputs = []string{}
	}

	context := internal.OnlineQueryContext{
		Environment:          internal.StringOrNil(p.EnvironmentId),
		Tags:                 p.Tags,
		RequiredResolverTags: p.RequiredResolverTags,
	}

	var now *string
	if len(p.Now) > 1 {
		return nil, fmt.Errorf(
			"for non-bulk queries, there should only"+
				" be 1 `Now` value, found %d", len(p.Now),
		)
	} else if len(p.Now) == 1 {
		n := p.Now[0].Format(internal.NowTimeFormat)
		now = &n
	}

	convertedInputs := make(map[string]any)
	for fqn, values := range p.inputs {
		convertedValues, err := internal.PreprocessIfStruct(values)
		if err != nil {
			return nil, wrapClientError(err, "failed to convert structs in input feature values")
		}
		convertedInputs[fqn] = convertedValues
	}

	var encodingOptions internal.FeatureEncodingOptions
	if p.EncodingOptions == nil {
		encodingOptions = internal.FeatureEncodingOptions{
			EncodeStructsAsObjects: false,
		}
	} else {
		encodingOptions = internal.FeatureEncodingOptions{
			EncodeStructsAsObjects: p.EncodingOptions.EncodeStructsAsObjects,
		}
	}

	return &internal.OnlineQueryRequestSerialized{
		Inputs:           convertedInputs,
		Outputs:          outputs,
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
		EncodingOptions:  encodingOptions,
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
	var timeObj time.Time
	if feature.Timestamp != "" {
		parsed, err := time.Parse(time.RFC3339, feature.Timestamp)
		if err != nil {
			return FeatureResult{}, err
		}
		timeObj = parsed
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
	outputs := colls.Map(params.outputs, func(v string) *commonv1.OutputExpr {
		return &commonv1.OutputExpr{
			Expr: &commonv1.OutputExpr_FeatureFqn{
				FeatureFqn: v,
			},
		}
	})

	staleness := make(map[string]string)
	for k, v := range params.staleness {
		staleness[k] = internal.FormatBucketDuration(int(v.Seconds()))
	}

	nowProto := colls.Map(params.Now, func(v time.Time) *timestamppb.Timestamp {
		return timestamppb.New(v)
	})

	options := map[string]*structpb.Value{}
	if params.StorePlanStages {
		options["store_plan_stages"] = structpb.NewBoolValue(params.StorePlanStages)
	}
	if params.IncludeMetrics {
		options["include_metrics"] = structpb.NewBoolValue(params.IncludeMetrics)
	}

	now := nowProto
	if len(nowProto) == 0 {
		now = nil
	}

	var explainOptions *commonv1.ExplainOptions
	if params.Explain {
		explainOptions = &commonv1.ExplainOptions{}
	}

	return &commonv1.OnlineQueryBulkRequest{
		InputsFeather: inputsFeather,
		Outputs:       outputs,
		Staleness:     staleness,
		Now:           now,
		Context: &commonv1.OnlineQueryContext{
			Environment:          params.EnvironmentId,
			Tags:                 params.Tags,
			DeploymentId:         ptr.PtrOrNil(params.PreviewDeploymentId),
			BranchId:             params.BranchId,
			CorrelationId:        ptr.PtrOrNil(params.CorrelationId),
			QueryName:            ptr.PtrOrNil(params.QueryName),
			QueryNameVersion:     ptr.PtrOrNil(params.QueryNameVersion),
			RequiredResolverTags: params.RequiredResolverTags,
			Options:              options,
		},
		ResponseOptions: &commonv1.OnlineQueryResponseOptions{
			IncludeMeta:     params.IncludeMeta || params.Explain,
			Metadata:        params.Meta,
			EncodingOptions: nil,
			Explain:         explainOptions,
		},
	}, nil
}
