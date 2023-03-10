package enum

type ErrorCodeCategory struct {
	Value string
}

var (
	Request = ErrorCodeCategory{"REQUEST"}
	Field   = ErrorCodeCategory{"FIELD"}
	Network = ErrorCodeCategory{"NETWORK"}
)

var allErrorCodeCategories = map[string]ErrorCodeCategory{
	Request.Value: Request,
	Field.Value:   Field,
	Network.Value: Network,
}

var GetErrorCodeCategory = generateGetEnumFunction(allErrorCodeCategories, "error code categories")
