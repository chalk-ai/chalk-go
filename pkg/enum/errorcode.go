package enum

type ErrorCode struct {
	Value string
}

var (
	ParseFailed         = ErrorCode{"PARSE_FAILED"}
	ResolverNotFound    = ErrorCode{"RESOLVER_NOT_FOUND"}
	InvalidQuery        = ErrorCode{"INVALID_QUERY"}
	ValidationFailed    = ErrorCode{"VALIDATION_FAILED"}
	ResolverFailed      = ErrorCode{"RESOLVER_FAILED"}
	ResolverTimedOut    = ErrorCode{"RESOLVER_TIMED_OUT"}
	UpstreamFailed      = ErrorCode{"UPSTREAM_FAILED"}
	Unauthenticated     = ErrorCode{"UNAUTHENTICATED"}
	Unauthorized        = ErrorCode{"UNAUTHORIZED"}
	InternalServerError = ErrorCode{"INTERNAL_SERVER_ERROR"}
	Cancelled           = ErrorCode{"CANCELLED"}
	DeadlineExceeded    = ErrorCode{"DEADLINE_EXCEEDED"}
)

var allErrorCodes = map[string]ErrorCode{
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
}

var GetErrorCode = generateGetEnumFunction(allErrorCodes, "error codes")
