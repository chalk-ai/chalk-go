package chalk

import (
	"fmt"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"time"
)

func queryMetaFromProto(metaRaw *commonv1.OnlineQueryMetadata) *QueryMeta {
	if metaRaw == nil {
		return nil
	}
	var executionDuration float64
	if metaRaw.GetExecutionDuration() != nil {
		executionDuration = metaRaw.GetExecutionDuration().AsDuration().Seconds()
	}

	var queryTimestamp *time.Time
	if metaRaw.GetQueryTimestamp() != nil {
		queryTimestamp = ptr.Ptr(metaRaw.GetQueryTimestamp().AsTime())
	}

	return &QueryMeta{
		ExecutionDurationS: executionDuration,
		DeploymentId:       metaRaw.DeploymentId,
		EnvironmentId:      metaRaw.EnvironmentId,
		EnvironmentName:    metaRaw.EnvironmentName,
		QueryId:            metaRaw.QueryId,
		QueryTimestamp:     queryTimestamp,
		QueryHash:          metaRaw.QueryHash,
	}
}

var errorCodeMap = map[commonv1.ErrorCode]ErrorCode{
	commonv1.ErrorCode_ERROR_CODE_INTERNAL_SERVER_ERROR_UNSPECIFIED: InternalServerError,
	commonv1.ErrorCode_ERROR_CODE_PARSE_FAILED:                      ParseFailed,
	commonv1.ErrorCode_ERROR_CODE_RESOLVER_NOT_FOUND:                ResolverNotFound,
	commonv1.ErrorCode_ERROR_CODE_INVALID_QUERY:                     InvalidQuery,
	commonv1.ErrorCode_ERROR_CODE_VALIDATION_FAILED:                 ValidationFailed,
	commonv1.ErrorCode_ERROR_CODE_RESOLVER_FAILED:                   ResolverFailed,
	commonv1.ErrorCode_ERROR_CODE_RESOLVER_TIMED_OUT:                ResolverTimedOut,
	commonv1.ErrorCode_ERROR_CODE_UPSTREAM_FAILED:                   UpstreamFailed,
	commonv1.ErrorCode_ERROR_CODE_UNAUTHENTICATED:                   Unauthenticated,
	commonv1.ErrorCode_ERROR_CODE_UNAUTHORIZED:                      Unauthorized,
	commonv1.ErrorCode_ERROR_CODE_CANCELLED:                         Cancelled,
	commonv1.ErrorCode_ERROR_CODE_DEADLINE_EXCEEDED:                 DeadlineExceeded,
}

var errorCodeCategoryMap = map[commonv1.ErrorCodeCategory]ErrorCodeCategory{
	commonv1.ErrorCodeCategory_ERROR_CODE_CATEGORY_NETWORK_UNSPECIFIED: Network,
	commonv1.ErrorCodeCategory_ERROR_CODE_CATEGORY_REQUEST:             Request,
	commonv1.ErrorCodeCategory_ERROR_CODE_CATEGORY_FIELD:               Field,
}

func exceptionFromProto(e *commonv1.ChalkException) *ResolverException {
	if e == nil {
		return nil
	}
	return &ResolverException{
		Stacktrace: e.GetStacktrace(),
		Message:    e.GetMessage(),
		Kind:       e.GetKind(),
	}
}

func serverErrorFromProto(e *commonv1.ChalkError) (*ServerError, error) {
	if e == nil {
		return nil, nil
	}
	code, ok := errorCodeMap[e.GetCode()]
	if !ok {
		return nil, fmt.Errorf("unknown error code: %v", e.GetCode())
	}
	category, ok := errorCodeCategoryMap[e.GetCategory()]
	if !ok {
		return nil, fmt.Errorf("unknown error code category: %v", e.GetCategory())
	}

	return &ServerError{
		Code:      code,
		Message:   e.GetMessage(),
		Category:  category,
		Feature:   e.GetFeature(),
		Resolver:  e.GetResolver(),
		Exception: exceptionFromProto(e.GetException()),
	}, nil
}

func serverErrorsFromProto(e []*commonv1.ChalkError) ([]ServerError, error) {
	var errs []ServerError
	for _, err := range e {
		s, err := serverErrorFromProto(err)
		if err != nil {
			return nil, err
		}
		errs = append(errs, *s)
	}
	return errs, nil
}
