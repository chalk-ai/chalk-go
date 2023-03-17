package chalk

import (
	"net/http"
)

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

/*
   def offline_query(
       self,
       input: Optional[Union[Mapping[Union[str, Feature, Any], Any], pd.DataFrame, pl.DataFrame, DataFrame]] = None,
       input_times: Union[Sequence[datetime], datetime, None] = None,
       output: Sequence[Union[str, Feature, Any]] = (),
       required_output: Sequence[Union[str, Feature, Any]] = (),
       environment: Optional[EnvironmentId] = None,
       dataset_name: Optional[str] = None,
       branch: Optional[BranchId] = None,
       max_samples: Optional[int] = None,
   ) -> DatasetImpl:


        req = CreateOfflineQueryJobRequest(
            output=optional_output,
            required_output=required_output,
            destination_format="PARQUET",
            input=query_input,
            max_samples=max_samples,
            dataset_name=dataset_name,
            branch=branch,
        )
*/

type offlineQueryInputSerialized struct {
	Columns []string          `json:"columns"`
	Values  [][]FeatureResult `json:"values"`
}
type offlineQueryRequestSerialized struct {
	Input             offlineQueryInputSerialized `json:"input"`
	Output            []string                    `json:"output"`
	RequiredOutput    []string                    `json:"required_output"`
	DatasetName       *string                     `json:"dataset_name"`
	Branch            *string                     `json:"branch"`
	MaxSamples        *int                        `json:"max_samples"`
	DestinationFormat string                      `json:"destination_format"`
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

	EnvironmentOverride string
	PreviewDeploymentId string
}

type getTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type getTokenResponse struct {
	AccessToken        string `json:"access_token"`
	TokenType          string `json:"token_type"`
	ExpiresIn          int    `json:"expires_in"`
	ApiServer          string `json:"api_server"`
	PrimaryEnvironment string `json:"primary_environment"`
}

type chalkErrorSerialized struct {
	Code      string             `json:"code"`
	Category  string             `json:"category"`
	Message   string             `json:"message"`
	Exception *ResolverException `json:"exception"`
	Feature   string             `json:"feature"`
	Resolver  string             `json:"resolver"`
}
