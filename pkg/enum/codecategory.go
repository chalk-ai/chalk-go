package enum

type ErrorCodeCategory struct {
	Value string
}

var (
	REQUEST = ErrorCodeCategory{"REQUEST"}
	FIELD   = ErrorCodeCategory{"FIELD"}
	NETWORK = ErrorCodeCategory{"NETWORK"}
)

var allErrorCodeCategories = map[string]ErrorCodeCategory{
	REQUEST.Value: REQUEST,
	FIELD.Value:   FIELD,
	NETWORK.Value: NETWORK,
}

var GetErrorCodeCategory = generateGetEnumFunction(allErrorCodeCategories, "error code categories")
