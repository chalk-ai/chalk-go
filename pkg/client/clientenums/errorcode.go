package clientenums

import "github.com/chalk-ai/chalk-go/pkg/utils"

type ErrorCode struct {
	Value string
}

var (
	PARSE_FAILED          = ErrorCode{"PARSE_FAILED"}
	RESOLVER_NOT_FOUND    = ErrorCode{"RESOLVER_NOT_FOUND"}
	INVALID_QUERY         = ErrorCode{"INVALID_QUERY"}
	VALIDATION_FAILED     = ErrorCode{"VALIDATION_FAILED"}
	RESOLVER_FAILED       = ErrorCode{"RESOLVER_FAILED"}
	RESOLVER_TIMED_OUT    = ErrorCode{"RESOLVER_TIMED_OUT"}
	UPSTREAM_FAILED       = ErrorCode{"UPSTREAM_FAILED"}
	UNAUTHENTICATED       = ErrorCode{"UNAUTHENTICATED"}
	UNAUTHORIZED          = ErrorCode{"UNAUTHORIZED"}
	INTERNAL_SERVER_ERROR = ErrorCode{"INTERNAL_SERVER_ERROR"}
	CANCELLED             = ErrorCode{"CANCELLED"}
	DEADLINE_EXCEEDED     = ErrorCode{"DEADLINE_EXCEEDED"}
)

var AllErrorCodes = map[string]ErrorCode{
	PARSE_FAILED.Value:          PARSE_FAILED,
	RESOLVER_NOT_FOUND.Value:    RESOLVER_NOT_FOUND,
	INVALID_QUERY.Value:         INVALID_QUERY,
	VALIDATION_FAILED.Value:     VALIDATION_FAILED,
	RESOLVER_FAILED.Value:       RESOLVER_FAILED,
	UPSTREAM_FAILED.Value:       UPSTREAM_FAILED,
	UNAUTHENTICATED.Value:       UNAUTHENTICATED,
	UNAUTHORIZED.Value:          UNAUTHORIZED,
	INTERNAL_SERVER_ERROR.Value: INTERNAL_SERVER_ERROR,
	CANCELLED.Value:             CANCELLED,
	DEADLINE_EXCEEDED.Value:     DEADLINE_EXCEEDED,
}

var GetErrorCode = utils.GenerateGetEnumFunction(AllErrorCodes, "error codes")
