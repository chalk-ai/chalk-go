package chalk

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/expr"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func serializeOnlineQueryParams(p *OnlineQueryParams, resolved *onlineQueryParamsResolved) (*internal.OnlineQueryRequestSerialized, error) {
	outputs := resolved.outputs
	if outputs == nil {
		// If we are passing query name, we don't need to pass outputs,
		// so outputs is empty, but when JSON serialized should never
		// be `null`.
		outputs = []string{}
	}

	var now *string
	if len(p.Now) > 1 {
		return nil, fmt.Errorf(
			"for non-bulk queries, there should only be 1 `Now` value, found %d",
			len(p.Now),
		)
	} else if len(p.Now) == 1 {
		n := p.Now[0].Format(internal.NowTimeFormat)
		now = &n
	}

	convertedInputs := make(map[string]any)
	for fqn, values := range resolved.inputs {
		convertedValues, err := internal.PreprocessIfStruct(values)
		if err != nil {
			return nil, errors.Wrap(err, "convert structs in input feature values")
		}
		convertedInputs[fqn] = convertedValues
	}

	result := &internal.OnlineQueryRequestSerialized{
		Inputs:  convertedInputs,
		Outputs: outputs,
		Context: internal.OnlineQueryContext{
			Tags:                 p.Tags,
			RequiredResolverTags: p.RequiredResolverTags,
		},
		Staleness:        serializeStaleness(resolved.staleness),
		IncludeMeta:      p.IncludeMeta || p.Explain,
		QueryName:        internal.StringOrNil(p.QueryName),
		QueryNameVersion: internal.StringOrNil(p.QueryNameVersion),
		CorrelationId:    internal.StringOrNil(p.CorrelationId),
		QueryContext:     p.QueryContext.ToMap(),
		Meta:             p.Meta,
		StorePlanStages:  p.StorePlanStages,
		Now:              now,
		Explain:          p.Explain,
		EncodingOptions: internal.FeatureEncodingOptions{
			// To ensure backcompat with codegen'd structs.
			// See https://github.com/chalk-ai/chalk-go/pull/159
			// And also makes unmarshalling easier. This is
			// unspecifiable by the user.
			EncodeStructsAsObjects: true,
		},
		PlannerOptions: p.PlannerOptions,
		BranchId:       p.BranchId,
	}

	return result, nil
}

func serializeStaleness(staleness map[string]time.Duration) map[string]string {
	res := map[string]string{}
	for k, v := range staleness {
		res[k] = internal.FormatBucketDuration(int(v.Seconds()))
	}
	return res
}

func (feature featureResultSerialized) deserialize() (FeatureResult, error) {
	var timeObj *time.Time
	if feature.Timestamp != "" {
		parsed, err := time.Parse(time.RFC3339, feature.Timestamp)
		if err != nil {
			return FeatureResult{}, err
		}
		timeObj = &parsed
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

func deserializeFeatureResults(results []featureResultSerialized) ([]FeatureResult, error) {
	deserializedResults := make([]FeatureResult, 0)
	for _, sResult := range results {
		dResult, err := sResult.deserialize()
		if err != nil {
			return []FeatureResult{}, err
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

// getRecomputeFeaturesValue returns the appropriate value for recompute_features field
// based on whether RecomputeFeatures or RecomputeFeaturesList is set
func getRecomputeFeaturesValue(p *OfflineQueryParams) interface{} {
	// If RecomputeFeaturesList is set, use that
	if len(p.RecomputeFeaturesList) > 0 {
		return p.RecomputeFeaturesList
	}
	// Otherwise use the boolean value
	return p.RecomputeFeatures
}

func serializeOfflineQueryParams(p *OfflineQueryParams, resolved *offlineQueryParamsResolved) ([]byte, error) {
	var queryInput interface{}

	// Check if rawFileInput is provided
	if p.rawFileInput != nil && *p.rawFileInput != "" {
		// Use OfflineQueryInputUri format
		queryInput = &internal.OfflineQueryInputUri{
			ParquetUri: *p.rawFileInput,
		}
	} else if len(resolved.inputs) > 0 {
		// Use OfflineQueryInputSerialized format
		queryInputSerialized := &internal.OfflineQueryInputSerialized{}
		globalInputTimes := make([]any, 0)
		for fqn, tsFeatureValues := range resolved.inputs {
			var inputValues []any
			var inputTimes []any
			for _, v := range tsFeatureValues {
				inputTimes = append(inputTimes, v.ObservationTime.Format(time.RFC3339))
				inputValues = append(inputValues, v.Value)
			}
			queryInputSerialized.Columns = append(queryInputSerialized.Columns, fqn)
			queryInputSerialized.Values = append(queryInputSerialized.Values, inputValues)
			globalInputTimes = inputTimes
		}
		queryInputSerialized.Columns = append(queryInputSerialized.Columns, "__chalk__.CHALK_TS")
		queryInputSerialized.Values = append(queryInputSerialized.Values, globalInputTimes)
		queryInput = queryInputSerialized
	}

	output := resolved.outputs
	if output == nil {
		output = make([]string, 0)
	}

	requiredOutput := resolved.requiredOutputs
	if requiredOutput == nil {
		requiredOutput = make([]string, 0)
	}

	// Convert ResourceRequests to ResourceRequestsSerialized
	var resourcesSerialized *internal.ResourceRequestsSerialized
	if p.Resources != nil {
		resourcesSerialized = &internal.ResourceRequestsSerialized{
			CPU:                 p.Resources.CPU,
			Memory:              p.Resources.Memory,
			EphemeralVolumeSize: p.Resources.EphemeralVolumeSize,
			EphemeralStorage:    p.Resources.EphemeralStorage,
			ResourceGroup:       p.Resources.ResourceGroup,
		}
	}

	// Convert CompletionDeadline to string format
	var completionDeadlineStr *string
	if p.CompletionDeadline != nil {
		duration := p.CompletionDeadline.String()
		completionDeadlineStr = &duration
	}

	// Convert Tags to pointer if not empty
	var tagsPtr *[]string
	if len(p.Tags) > 0 {
		tagsPtr = &p.Tags
	}

	// Convert EnvOverrides to pointer if not empty
	var envOverridesPtr *map[string]string
	if len(p.EnvOverrides) > 0 {
		envOverridesPtr = &p.EnvOverrides
	}

	// Convert RequiredResolverTags to pointer if not empty
	var requiredResolverTagsPtr *[]string
	if len(p.RequiredResolverTags) > 0 {
		requiredResolverTagsPtr = &p.RequiredResolverTags
	}

	// Convert CorrelationId to pointer if not empty
	var correlationIdPtr *string
	if p.CorrelationId != "" {
		correlationIdPtr = &p.CorrelationId
	}

	// Convert PlannerOptions to pointer if not empty
	var plannerOptionsPtr *map[string]any
	if len(p.PlannerOptions) > 0 {
		plannerOptionsPtr = &p.PlannerOptions
	}

	// Convert SampleFeatures to pointer if not empty
	var sampleFeaturesPtr *[]string
	if len(p.SampleFeatures) > 0 {
		sampleFeaturesPtr = &p.SampleFeatures
	}

	// Convert remaining string fields to pointers if not empty
	var spineSqlQueryPtr *string
	if p.SpineSqlQuery != "" {
		spineSqlQueryPtr = &p.SpineSqlQuery
	}

	var recomputeRequestRevisionIdPtr *string
	if p.RecomputeRequestRevisionId != "" {
		recomputeRequestRevisionIdPtr = &p.RecomputeRequestRevisionId
	}

	var overrideTargetImageTagPtr *string
	if p.OverrideTargetImageTag != "" {
		overrideTargetImageTagPtr = &p.OverrideTargetImageTag
	}

	var featureForLowerUpperBoundPtr *string
	if p.FeatureForLowerUpperBound != "" {
		featureForLowerUpperBoundPtr = &p.FeatureForLowerUpperBound
	}

	// Convert time bounds to string format (using foreign branch's approach)
	var lowerBoundStr *string
	if p.ObservedAtLowerBound != nil {
		formatted := strings.Replace(p.ObservedAtLowerBound.Format(time.RFC3339Nano), "Z", "+00:00", 1)
		lowerBoundStr = &formatted
	}

	var upperBoundStr *string
	if p.ObservedAtUpperBound != nil {
		formatted := strings.Replace(p.ObservedAtUpperBound.Format(time.RFC3339Nano), "Z", "+00:00", 1)
		upperBoundStr = &formatted
	}

	// Build the serialized object to match Python structure exactly
	serializedObj := internal.OfflineQueryRequestSerialized{
		// Core fields
		Input:                      queryInput,
		Output:                     output,
		OutputExpressions:          []string{},
		RequiredOutput:             requiredOutput,
		RequiredOutputExpressions:  []string{},
		DestinationFormat:          "PARQUET",
		JobId:                      nil, // Always nil - server auto-generates
		MaxSamples:                 p.MaxSamples,
		MaxCacheAge:                nil,           // Deprecated in Python - always nil
		ObservedAtLowerBound:       lowerBoundStr, // Using foreign branch's inline approach
		ObservedAtUpperBound:       upperBoundStr, // Using foreign branch's inline approach
		DatasetName:                internal.StringOrNil(p.DatasetName),
		Branch:                     internal.StringOrNil(p.Branch),
		RecomputeFeatures:          getRecomputeFeaturesValue(p),
		SampleFeatures:             sampleFeaturesPtr,
		StorePlanStages:            p.StorePlanStages,
		Explain:                    p.Explain,
		Tags:                       tagsPtr,
		RequiredResolverTags:       requiredResolverTagsPtr,
		CorrelationId:              correlationIdPtr,
		QueryContext:               p.QueryContext.ToMap(),
		PlannerOptions:             plannerOptionsPtr,
		UseMultipleComputers:       p.UseMultipleComputers || p.RunAsynchronously,
		SpineSqlQuery:              spineSqlQueryPtr,
		RecomputeRequestRevisionId: recomputeRequestRevisionIdPtr,
		Resources:                  resourcesSerialized,
		EnvOverrides:               envOverridesPtr,
		OverrideTargetImageTag:     overrideTargetImageTagPtr,
		EnableProfiling:            p.EnableProfiling,
		StoreOnline:                p.StoreOnline,
		StoreOffline:               p.StoreOffline,
		NumShards:                  p.NumShards,
		NumWorkers:                 p.NumWorkers,
		FeatureForLowerUpperBound:  featureForLowerUpperBoundPtr,
		CompletionDeadline:         completionDeadlineStr,
		MaxRetries:                 p.MaxRetries,
		UseJobQueue:                p.UseJobQueue,
		OverlayGraph:               nil,
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

func deserializeChalkErrors(errors []chalkErrorSerialized) (ServerErrors, error) {
	deserializedErrors := make([]ServerError, 0)
	for _, serializedErr := range errors {
		deserializedError, err := serializedErr.deserialize()
		if err != nil {
			return []ServerError{}, err
		}
		deserializedErrors = append(deserializedErrors, deserializedError)

	}
	return deserializedErrors, nil
}

func generateGetEnumFunction[K comparable](valueToEnum map[string]K, enumName string) func(string) (*K, error) {
	return func(value string) (*K, error) {
		enum, found := valueToEnum[value]
		if !found {
			return nil, fmt.Errorf("cannot find enum value '%s' among all %s", value, enumName)
		}

		return &enum, nil
	}
}

var getErrorCode = generateGetEnumFunction(
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

var getErrorCodeCategory = generateGetEnumFunction(
	map[string]ErrorCodeCategory{
		Request.Value: Request,
		Field.Value:   Field,
		Network.Value: Network,
	},
	"error code categories",
)

func convertOnlineQueryParamsToProto(params *OnlineQueryParams, allocator memory.Allocator) (*commonv1.OnlineQueryBulkRequest, error) {
	resolved, err := params.resolveBulk()
	if err != nil {
		return nil, errors.Wrap(err, "resolving params")
	}
	inputsFeather, err := internal.InputsToArrowBytes(resolved.inputs, allocator)
	if err != nil {
		return nil, errors.Wrap(err, "error serializing inputs as feather")
	}

	var outputs []*commonv1.OutputExpr
	for _, o := range resolved.outputs {
		outputs = append(outputs, &commonv1.OutputExpr{
			Expr: &commonv1.OutputExpr_FeatureFqn{
				FeatureFqn: o,
			},
		})
	}
	for _, o := range params.OutputExprs {
		outputColumnName := ""
		if casted, ok := o.(*expr.AliasExpr); ok {
			outputColumnName = casted.Alias
		}
		exproto, err := expr.ToProto(o)
		if err != nil {
			return nil, err
		}
		outputs = append(
			outputs,
			&commonv1.OutputExpr{
				Expr: &commonv1.OutputExpr_FeatureExpression{
					FeatureExpression: &commonv1.FeatureExpression{
						OutputColumnName: outputColumnName,
						Namespace:        "",
						Expr:             exproto,
					},
				},
			},
		)
	}

	staleness := make(map[string]string)
	for k, v := range resolved.staleness {
		staleness[k] = internal.FormatBucketDuration(int(v.Seconds()))
	}

	nowProto := make([]*timestamppb.Timestamp, len(params.Now))
	for i, v := range params.Now {
		nowProto[i] = timestamppb.New(v)
	}

	options := map[string]*structpb.Value{}
	if params.StorePlanStages {
		options["store_plan_stages"] = structpb.NewBoolValue(params.StorePlanStages)
	}
	for k, v := range params.PlannerOptions {
		protoVal, err := structpb.NewValue(v)
		if err != nil {
			return nil, errors.Wrapf(err, "converting planner option value for '%s' to proto: %v", k, v)
		}
		options[k] = protoVal
	}

	now := nowProto
	if len(nowProto) == 0 {
		now = nil
	}

	var explainOptions *commonv1.ExplainOptions
	if params.Explain {
		explainOptions = &commonv1.ExplainOptions{}
	}
	queryContextProto, err := params.QueryContext.toProtoMap()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert query context to protobuf")
	}

	return &commonv1.OnlineQueryBulkRequest{
		Inputs: &commonv1.OnlineQueryBulkRequest_InputsFeather{
			InputsFeather: inputsFeather,
		},
		Outputs:       outputs,
		Staleness:     staleness,
		Now:           now,
		Context: &commonv1.OnlineQueryContext{
			Tags:                 params.Tags,
			BranchId:             params.BranchId,
			CorrelationId:        ptr.OrNil(params.CorrelationId),
			QueryName:            ptr.OrNil(params.QueryName),
			QueryNameVersion:     ptr.OrNil(params.QueryNameVersion),
			QueryContext:         queryContextProto,
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
