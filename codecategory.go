package chalk

type ErrorCodeCategory struct {
	Value string
}

var (
	// Request errors are raised before execution of your
	// resolver code. They may occur due to invalid feature
	// names in the input or a request that cannot be satisfied
	// by the resolvers you have defined.
	Request = ErrorCodeCategory{"REQUEST"}

	// Field errors are raised while running a feature resolver
	// for a particular field. For this type of error, you'll
	// find a feature and resolver attribute in the error type.
	// When a feature resolver crashes, you will receive null
	// value in the response. To differentiate from a resolver
	// returning a null value and a failure in the resolver,
	// you need to check the error schema.
	Field = ErrorCodeCategory{"FIELD"}

	// Network errors are thrown outside your resolvers.
	// For example, your request was unauthenticated,
	// connection failed, or an error occurred within Chalk.
	Network = ErrorCodeCategory{"NETWORK"}
)
