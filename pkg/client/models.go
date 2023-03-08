package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

type Client struct {
	BaseUrl       string
	jwt           *auth.JWT
	httpClient    *http.Client
	EnvironmentId string
	ClientId      string
	ClientSecret  string
}

type OnlineQueryRequest struct {
	Inputs         map[string]any      `json:"inputs,string"`
	Outputs        []string            `json:"outputs"`
	Context        *OnlineQueryContext `json:"context"`
	Staleness      map[string]string   `json:"staleness"`
	IncludeMeta    bool                `json:"include_meta"`
	IncludeMetrics bool                `json:"include_metrics"`
	DeploymentId   string              `json:"deployment_id"`
	QueryName      string              `json:"query_name"`
	CorrelationId  string              `json:"correlation_id"`
	Meta           map[string]string   `json:"meta"`
}

type onlineQueryHttpRequest struct {
	Inputs         map[string]any      `json:"inputs,string"`
	Outputs        []string            `json:"outputs"`
	Context        *OnlineQueryContext `json:"context"`
	Staleness      map[string]string   `json:"staleness"`
	IncludeMeta    bool                `json:"include_meta"`
	IncludeMetrics bool                `json:"include_metrics"`
	DeploymentId   *string             `json:"deployment_id"`
	QueryName      *string             `json:"query_name"`
	CorrelationId  *string             `json:"correlation_id"`
	Meta           map[string]string   `json:"meta"`
}

type OnlineQueryContext struct {
	Environment string   `json:"environment"`
	Tags        []string `json:"tags"`
}

type ChalkException struct {
	Kind       string `json:"kind"`
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
}

type QueryMeta struct {
	ExecutionDurationS float64 `json:"execution_duration_s"`
	DeploymentId       string  `json:"deployment_id"`
	QueryId            string  `json:"query_id"`
}

type OnlineQueryResponse struct {
	Data []FeatureResult `json:"data"`
	Meta QueryMeta       `json:"meta"`
}

type onlineQueryHttpResponse struct {
	Data   []FeatureResult `json:"data"`
	Errors []ChalkError    `json:"errors"`
	Meta   QueryMeta       `json:"meta"`
}

type FeatureResult struct {
	Field     string         `json:"field"`
	Value     any            `json:"value"`
	Timestamp string         `json:"ts"`
	Meta      map[string]any `json:"meta"`
	Error     *ChalkError    `json:"error"`
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
