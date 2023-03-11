package chalk

import "github.com/chalk-ai/chalk-go/internal"

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

var GetErrorCodeCategory = internal.GenerateGetEnumFunction(allErrorCodeCategories, "error code categories")
