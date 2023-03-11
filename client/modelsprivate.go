package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"net/http"
)

type chalkClientImpl struct {
	ApiServer     config
	ClientId      config
	EnvironmentId config

	clientSecret config
	jwt          *auth.JWT
	httpClient   *http.Client
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
	Data   []featureResultSerialized `json:"data"`
	Errors []chalkErrorSerialized    `json:"errors"`
	Meta   *QueryMeta                `json:"meta"`
}

type onlineQueryContext struct {
	Environment *string  `json:"environment"`
	Tags        []string `json:"tags"`
}

type featureResultSerialized struct {
	Field     string                 `json:"field"`
	Value     any                    `json:"value"`
	Pkey      any                    `json:"pkey"`
	Timestamp string                 `json:"ts"`
	Meta      *FeatureResolutionMeta `json:"meta"`
	Error     *chalkErrorSerialized  `json:"error"`
}

type chalkHttpException struct {
	Detail *string `json:"detail"`
	Trace  *string `json:"trace"`
}

type sendRequestParams struct {
	Request *http.Request

	Body   any
	Method string
	URL    string

	Response    any
	DontRefresh bool
}

type config struct {
	Value  string
	Source string
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

type chalkErrorSerialized struct {
	Code      string          `json:"code"`
	Category  string          `json:"category"`
	Message   string          `json:"message"`
	Exception *ChalkException `json:"exception"`
	Feature   string          `json:"feature"`
	Resolver  string          `json:"resolver"`
}
