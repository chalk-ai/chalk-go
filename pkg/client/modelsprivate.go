package client

type config struct {
	Value  string
	Source string
}

type onlineQueryRequestSerialized struct {
	Inputs         map[string]any     `json:"inputs,string"`
	Outputs        []string           `json:"outputs"`
	Context        onlineQueryContext `json:"context"`
	Staleness      map[string]string  `json:"staleness"`
	IncludeMeta    bool               `json:"include_meta"`
	IncludeMetrics bool               `json:"include_metrics"`
	DeploymentId   *string            `json:"deployment_id"`
	QueryName      *string            `json:"query_name"`
	CorrelationId  *string            `json:"correlation_id"`
	Meta           map[string]string  `json:"meta"`
}

type onlineQueryResponseSerialized struct {
	Data   []FeatureResult        `json:"data"`
	Errors []chalkErrorSerialized `json:"errors"`
	// Make query meta pointer.
	Meta *queryMeta `json:"meta"`
}

type onlineQueryContext struct {
	Environment *string  `json:"environment"`
	Tags        []string `json:"tags"`
}

type chalkHttpException struct {
	Detail *string `json:"detail"`
	Trace  *string `json:"trace"`
}

type getTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type getTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ApiServer   string `json:"api_server"`
}

// TODO create public ChalkException?
type chalkException struct {
	Kind       string `json:"kind"`
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
}

// TODO create public QueryMeta
type queryMeta struct {
	ExecutionDurationS float64 `json:"execution_duration_s"`
	DeploymentId       string  `json:"deployment_id"`
	QueryId            string  `json:"query_id"`
}

type chalkErrorSerialized struct {
	Code      string          `json:"code"`
	Category  string          `json:"category"`
	Message   string          `json:"message"`
	Exception *chalkException `json:"exception"`
	Feature   string          `json:"feature"`
	Resolver  string          `json:"resolver"`
}
