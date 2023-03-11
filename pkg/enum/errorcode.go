package enum

// ErrorCode indicates the type of error occurred
type ErrorCode struct {
	Value string
}

var (
	// ParseFailed indicates the query contained features that do not exist.
	ParseFailed = ErrorCode{Value: "PARSE_FAILED"}

	// ResolverNotFound indicates a resolver was required as part of running the dependency graph that could not be found.
	ResolverNotFound = ErrorCode{"RESOLVER_NOT_FOUND"}

	// InvalidQuery indicates the query is invalid. All supplied features need to be rooted in the same top-level entity.
	InvalidQuery = ErrorCode{"INVALID_QUERY"}

	// ValidationFailed indicates a feature value did not match the expected schema (e.g. `incompatible type "int"; expected "str"`)
	ValidationFailed = ErrorCode{"VALIDATION_FAILED"}

	// ResolverFailed indicates the resolver for a feature errored.
	ResolverFailed = ErrorCode{"RESOLVER_FAILED"}

	// ResolverTimedOut indicates the resolver for a feature timed out.
	ResolverTimedOut = ErrorCode{"RESOLVER_TIMED_OUT"}

	// UpstreamFailed indicates a crash in a resolver that was to produce an input for the resolver crashed, and so the resolver could not run.
	UpstreamFailed = ErrorCode{"UPSTREAM_FAILED"}

	// Unauthenticated indicates the request was submitted with an invalid authentication header.
	Unauthenticated = ErrorCode{"UNAUTHENTICATED"}

	// Unauthorized indicates the supplied credentials do not provide the right authorization to execute the request.
	Unauthorized = ErrorCode{"UNAUTHORIZED"}

	// InternalServerError indicates an unspecified error occurred.
	InternalServerError = ErrorCode{"INTERNAL_SERVER_ERROR"}

	// Cancelled indicates the operation was cancelled, typically by the caller.
	Cancelled = ErrorCode{"CANCELLED"}

	// DeadlineExceeded indicates the deadline expired before the operation could complete.
	DeadlineExceeded = ErrorCode{"DEADLINE_EXCEEDED"}
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
