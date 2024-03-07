package chalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type clientImpl struct {
	ApiServer     auth2.SourcedConfig
	ClientId      auth2.SourcedConfig
	EnvironmentId auth2.SourcedConfig
	Branch        string

	clientSecret       auth2.SourcedConfig
	jwt                *auth2.JWT
	httpClient         HTTPClient
	logger             LeveledLogger
	initialEnvironment auth2.SourcedConfig
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
}

func (c *clientImpl) OfflineQuery(params OfflineQueryParamsComplete) (Dataset, error) {
	request := params.underlying

	if len(request.builderErrors) > 0 {
		builderErrString := request.builderErrors.Error()
		clientErrString := "error building offline query params:\n" + builderErrString
		return Dataset{}, &ErrorResponse{ClientError: &ClientError{clientErrString}}
	}

	emptyResult := Dataset{}
	response := Dataset{}

	err := c.sendRequest(
		sendRequestParams{
			Method:              "POST",
			URL:                 "v3/offline_query",
			Body:                request,
			Response:            &response,
			EnvironmentOverride: request.EnvironmentId,
			Branch:              request.Branch,
		},
	)
	if err != nil {
		return emptyResult, getErrorResponse(err)
	}

	if len(response.Errors) > 0 {
		return emptyResult, &ErrorResponse{ServerErrors: response.Errors}
	}

	for idx, _ := range response.Revisions {
		response.Revisions[idx].client = c
	}

	return response, nil
}

func (c *clientImpl) OnlineQueryBulk(params OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	emptyResult := OnlineQueryBulkResult{}
	request := params.underlying

	if len(request.builderErrors) > 0 {
		builderErrString := request.builderErrors.Error()
		clientErrString := "error building bulk online query params:\n" + builderErrString
		return emptyResult, &ErrorResponse{ClientError: &ClientError{clientErrString}}
	}

	validationErrors := params.validatePostBuild()
	if len(validationErrors) > 0 {
		return emptyResult, &ErrorResponse{ClientError: &ClientError{validationErrors.Error()}}
	}

	for _, input := range request.inputs {
		if !(reflect.ValueOf(input).Kind() == reflect.Slice || reflect.ValueOf(input).Kind() == reflect.Array) {
			return emptyResult, &ErrorResponse{
				ClientError: &ClientError{
					"Inputs to bulk online query must be a slice or array",
				},
			}
		}
	}
	data, err := params.ToBytes()
	if err != nil {
		return emptyResult, &ErrorResponse{ClientError: &ClientError{fmt.Errorf("error serializing online query params: %w", err).Error()}}
	}
	var response OnlineQueryBulkResponse
	err = c.sendRequest(
		sendRequestParams{
			Method:              "POST",
			URL:                 "v1/query/feather",
			Body:                data,
			Response:            &response,
			EnvironmentOverride: params.underlying.EnvironmentId,
			PreviewDeploymentId: params.underlying.PreviewDeploymentId,
		},
	)

	if err != nil {
		return emptyResult, getErrorResponse(err)
	}

	singleBulkResult, ok := response.QueryResults["0"]
	if !ok {
		return emptyResult, &ErrorResponse{ClientError: &ClientError{"unexpected bulk online query response from server"}}
	}

	if len(singleBulkResult.Errors) > 0 {
		return emptyResult, &ErrorResponse{ServerErrors: singleBulkResult.Errors}
	}

	return OnlineQueryBulkResult{
		ScalarsTable: singleBulkResult.ScalarData,
		GroupsTables: singleBulkResult.GroupsData,
		Meta:         singleBulkResult.Meta,
	}, nil
}

func (c *clientImpl) UploadFeatures(params UploadFeaturesParams) (UploadFeaturesResult, error) {
	castMap := make(map[string]any)

	allLength := -1
	for k, v := range params.Inputs {
		var fqn string
		if _, ok := k.(string); ok {
			fqn = k.(string)
		} else {
			feature, err := UnwrapFeature(k)
			if err != nil {
				msg := fmt.Sprintf("Invalid inputs key '%v' with type '%T'. Expected `string` or `Feature`", k, k)
				return UploadFeaturesResult{}, &ErrorResponse{ClientError: &ClientError{Message: msg}}
			}
			fqn = feature.Fqn
		}
		castMap[fqn] = v

		currLength := -1
		if reflect.TypeOf(v).Kind() == reflect.Slice || reflect.TypeOf(v).Kind() == reflect.Array {
			currLength = reflect.ValueOf(v).Len()
		} else {
			return UploadFeaturesResult{}, &ErrorResponse{
				ClientError: &ClientError{
					Message: fmt.Sprintf("Values for feature '%s' must be a slice or array", fqn),
				},
			}
		}

		if allLength == -1 {
			allLength = currLength
		}
		if allLength != currLength {
			err := &ClientError{
				Message: fmt.Sprintf("All input slices or arrays must be the same length - found length %d for feature '%s' but expected length %d", currLength, fqn, allLength),
			}
			return UploadFeaturesResult{}, &ErrorResponse{ClientError: err}
		}
		if currLength == 0 {
			err := &ClientError{
				Message: fmt.Sprintf("All input slices or arrays must be non-empty - found length %d for feature '%s'", currLength, fqn),
			}
			return UploadFeaturesResult{}, &ErrorResponse{ClientError: err}
		}
	}

	body, err := internal.CreateUploadFeaturesBody(castMap)
	if err != nil {
		return UploadFeaturesResult{}, &ErrorResponse{ClientError: &ClientError{Message: err.Error()}}
	}

	response := UploadFeaturesResult{}
	err = c.sendRequest(
		sendRequestParams{
			Method:              "POST",
			URL:                 "v1/upload_features/multi",
			Body:                body,
			Response:            &response,
			EnvironmentOverride: params.EnvironmentOverride,
			PreviewDeploymentId: params.PreviewDeploymentId,
		},
	)
	if err != nil {
		return UploadFeaturesResult{}, getErrorResponse(err)
	}
	return response, nil
}

func (c *clientImpl) OnlineQuery(params OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	request := params.underlying

	if len(request.builderErrors) > 0 {
		builderErrString := request.builderErrors.Error()
		clientErrString := "error building online query params:\n" + builderErrString
		return OnlineQueryResult{}, &ErrorResponse{ClientError: &ClientError{clientErrString}}
	}

	validationErrors := params.validatePostBuild()
	if len(validationErrors) > 0 {
		return OnlineQueryResult{}, &ErrorResponse{ClientError: &ClientError{validationErrors.Error()}}
	}

	for _, input := range request.inputs {
		if reflect.ValueOf(input).Kind() == reflect.Slice || reflect.ValueOf(input).Kind() == reflect.Array {
			return OnlineQueryResult{}, &ErrorResponse{
				ClientError: &ClientError{
					"inputs to online query must be a scalar value, found slice or array - did you mean to use OnlineQueryBulk?",
				},
			}
		}
	}

	emptyResult := OnlineQueryResult{}

	var serializedResponse onlineQueryResponseSerialized

	err := c.sendRequest(
		sendRequestParams{
			Method:              "POST",
			URL:                 "v1/query/online",
			Body:                request.serialize(),
			Response:            &serializedResponse,
			EnvironmentOverride: request.EnvironmentId,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	if err != nil {
		return emptyResult, getErrorResponse(err)
	}
	if len(serializedResponse.Errors) > 0 {
		serverErrors, deserializationErr := deserializeChalkErrors(serializedResponse.Errors)
		if deserializationErr != nil {
			return emptyResult, &ErrorResponse{
				ClientError: &ClientError{deserializationErr.Error()},
			}
		}

		return emptyResult, &ErrorResponse{ServerErrors: serverErrors}
	}

	response, err := serializedResponse.deserialize()
	if err != nil {
		return emptyResult, &ErrorResponse{
			ClientError: &ClientError{err.Error()},
		}
	}

	response.expectedOutputs = params.underlying.outputs
	if resultHolder != nil {
		unmarshalErr := response.UnmarshalInto(resultHolder)
		if unmarshalErr != nil {
			return response, &ErrorResponse{
				ClientError: unmarshalErr,
			}
		}
	}

	return response, nil
}

func (c *clientImpl) TriggerResolverRun(request TriggerResolverRunParams) (TriggerResolverRunResult, error) {
	response := TriggerResolverRunResult{}
	err := c.sendRequest(
		sendRequestParams{
			Method:              "POST",
			URL:                 "v1/runs/trigger",
			Body:                request,
			Response:            &response,
			EnvironmentOverride: request.EnvironmentId,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	if err != nil {
		return TriggerResolverRunResult{}, getErrorResponse(err)
	}
	return response, nil
}

func (c *clientImpl) GetRunStatus(request GetRunStatusParams) (GetRunStatusResult, error) {
	response := GetRunStatusResult{}
	err := c.sendRequest(
		sendRequestParams{
			Method:              "GET",
			URL:                 fmt.Sprintf("v1/runs/%s", request.RunId),
			Body:                request,
			Response:            &response,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	if err != nil {
		return GetRunStatusResult{}, getErrorResponse(err)
	}
	return response, nil
}

func (c *clientImpl) getDatasetUrls(RevisionId string, EnvironmentId string) ([]string, error) {
	response := GetOfflineQueryJobResponse{}

	for !response.IsFinished {
		err := c.sendRequest(
			sendRequestParams{
				Method:              "GET",
				URL:                 fmt.Sprintf("v2/offline_query/%s", RevisionId),
				EnvironmentOverride: EnvironmentId,
				Response:            &response,
			},
		)
		if err != nil {
			return []string{}, getErrorResponse(err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	return response.Urls, nil
}

func (c *clientImpl) saveUrlToDirectory(URL string, directory string) error {
	resp, err := c.httpClient.Get(URL)
	if err != nil {
		return err
	}
	defer func() {
		err = deferFunctionWithError(resp.Body.Close, err)
	}()

	parsedUrl, urlParseErr := url.Parse(URL)
	if urlParseErr != nil {
		return urlParseErr
	}
	destinationFilepath := filepath.Join(parsedUrl.Path[4:])

	directory, expandErr := internal.ExpandTilde(directory)
	if expandErr != nil {
		return fmt.Errorf("error getting home directory for path '%s' - please try an absolute path", directory)
	}
	destinationDirectory := filepath.Join(directory, filepath.Dir(destinationFilepath))

	if err = os.MkdirAll(destinationDirectory, os.ModePerm); err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	destinationPath := filepath.Join(directory, destinationFilepath)
	err = os.WriteFile(destinationPath, data, os.ModePerm)

	return err
}

func (c *clientImpl) getJwt() (*auth2.JWT, *ClientError) {
	body := getTokenRequest{
		ClientId:     c.ClientId.Value,
		ClientSecret: c.clientSecret.Value,
		GrantType:    "client_credentials",
	}
	response := getTokenResponse{}
	err := c.sendRequest(
		sendRequestParams{
			Method:      "POST",
			URL:         "v1/oauth/token",
			Body:        body,
			Response:    &response,
			DontRefresh: true,
		},
	)
	if err != nil {
		return nil, &ClientError{Message: fmt.Sprintf(
			"Error obtaining access token: %s.\n"+
				"  Auth config:\n"+
				"    api_server=%q (source: %s),\n"+
				"    client_id=%q (source: %s),\n"+
				"    client_secret=*** (source: %s),\n"+
				"    environment_id=%q (source: %s)\n",
			err.Error(),
			c.ApiServer.Value,
			c.ApiServer.Source,
			c.ClientId.Value,
			c.ClientId.Source,
			c.clientSecret.Source,
			c.EnvironmentId.Value,
			c.EnvironmentId.Source,
		)}
	}

	if c.initialEnvironment.Value == "" {
		c.EnvironmentId = auth2.SourcedConfig{
			Value:  response.PrimaryEnvironment,
			Source: "Primary Environment from credentials exchange response",
		}
	} else {
		c.EnvironmentId = c.initialEnvironment
	}

	expiry := time.Now().UTC().Add(time.Duration(response.ExpiresIn) * time.Second)
	jwt := &auth2.JWT{
		Token:      response.AccessToken,
		ValidUntil: expiry,
	}
	return jwt, nil
}

func (c *clientImpl) refreshJwt(forceRefresh bool) *ClientError {
	if !forceRefresh && c.jwt != nil && !time.Time.IsZero(c.jwt.ValidUntil) &&
		c.jwt.ValidUntil.After(time.Now().UTC().Add(-10*time.Second)) {
		return nil
	}

	jwt, getJwtErr := c.getJwt()
	if getJwtErr != nil {
		return getJwtErr
	}
	c.jwt = jwt
	return nil
}

func getBodyBuffer(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	switch v := body.(type) {
	case []byte:
		return bytes.NewBuffer(v), nil
	default:
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(jsonBytes), nil

	}
}

func (c *clientImpl) sendRequest(args sendRequestParams) error {
	body, getBufferErr := getBodyBuffer(args.Body)
	if getBufferErr != nil {
		return getBufferErr
	}

	request, newRequestErr := http.NewRequest(args.Method, args.URL, body)
	if newRequestErr != nil {
		(c.logger).Debugf("error sending request: %s", newRequestErr.Error())
		return newRequestErr
	}

	headers := c.getHeaders(args.EnvironmentOverride, args.PreviewDeploymentId, args.Branch)
	request.Header = headers

	if !args.DontRefresh {
		upsertJwtErr := c.refreshJwt(false)
		if upsertJwtErr != nil {
			(c.logger).Debugf("Error pre-emptively refreshing access token: %s", upsertJwtErr)
		}
	}
	if c.jwt != nil && c.jwt.Token != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.Token))
	}

	if !strings.HasPrefix(request.URL.String(), "http:") && !strings.HasPrefix(request.URL.String(), "https:") {
		var err error
		request.URL, err = url.Parse(fmt.Sprintf("%s/%s", c.ApiServer.Value, request.URL.String()))
		if err != nil {
			return err
		}
	}

	(c.logger).Debugf("Sending request to ", request.URL)
	res, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	(c.logger).Debugf("Response Status: ", res.Status)
	defer res.Body.Close()

	if res.StatusCode == 401 && !args.DontRefresh && request != nil {
		res, err = c.retryRequest(*request, args.Body, res, err)
		if err != nil {
			return err
		}
	}

	if res.StatusCode != 200 {
		clientError := getHttpError(c.logger, *res, *request)
		return &clientError
	}

	out, _ := io.ReadAll(res.Body)
	castResponse, isBulkResponse := args.Response.(*OnlineQueryBulkResponse)
	if isBulkResponse {
		err = castResponse.Unmarshal(out)
	} else {
		err = json.Unmarshal(out, args.Response)
	}

	return err
}

func (c *clientImpl) retryRequest(
	originalRequest http.Request, originalBody any,
	originalResponse *http.Response, originalError error,
) (*http.Response, error) {
	upsertJwtUpon401Err := c.refreshJwt(true)
	if upsertJwtUpon401Err != nil {
		(c.logger).Debugf("Error refreshing access token upon 401: %s", upsertJwtUpon401Err.Error())
		return originalResponse, originalError
	}

	originalBodyBuffer, getBufferErr := getBodyBuffer(originalBody)
	if getBufferErr != nil {
		return nil, getBufferErr
	}

	// New request needs to be constructed otherwise we were getting the error:
	//     HTTP/1.x transport connection broken
	newRequest, err := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), originalBodyBuffer)
	if err != nil {
		return nil, err
	}
	newRequest.Header = originalRequest.Header
	if c.jwt != nil && c.jwt.Token != "" {
		newRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.Token))
	}

	res, err := c.httpClient.Do(newRequest)
	if err != nil {
		return nil, err
	}
	(c.logger).Debugf("Response Status for retried request: ", res.Status)

	return res, nil
}

func (c *clientImpl) getHeaders(environmentOverride string, previewDeploymentId string, branchOverride string) http.Header {
	headers := http.Header{}

	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", "chalk-go-0.0")
	headers.Set("X-Chalk-Client-Id", c.ClientId.Value)

	if branchOverride != "" {
		headers.Set("X-Chalk-Branch-Id", branchOverride)
	} else if c.Branch != "" {
		headers.Set("X-Chalk-Branch-Id", c.Branch)
	}

	if environmentOverride == "" {
		headers.Set("X-Chalk-Env-Id", c.EnvironmentId.Value)
	} else {
		headers.Set("X-Chalk-Env-Id", environmentOverride)
	}
	if previewDeploymentId != "" {
		headers.Set("X-Chalk-Preview-Deployment", previewDeploymentId)
	}

	return headers
}

func getHttpError(logger LeveledLogger, res http.Response, req http.Request) HTTPError {
	var errorResponse chalkHttpException
	out, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(out, &errorResponse)
	logger.Errorf("API error response", err, errorResponse, string(out))

	clientError := HTTPError{
		Message:       "Unknown Chalk Server Error",
		Path:          req.URL.String(),
		StatusCode:    res.StatusCode,
		ContentLength: res.ContentLength,
		Trace:         errorResponse.Trace,
	}

	if errorResponse.Detail != nil {
		clientError.Message = *errorResponse.Detail
	}

	return clientError
}

func getErrorResponse(err error) *ErrorResponse {
	httpError, ok := err.(*HTTPError)
	if ok {
		return &ErrorResponse{HttpError: httpError}
	}
	return &ErrorResponse{ClientError: &ClientError{Message: err.Error()}}
}

func newClientImpl(
	cfgs ...*ClientConfig,
) (*clientImpl, error) {
	var cfg *ClientConfig
	if len(cfgs) == 0 {
		cfg = &ClientConfig{}
	} else {
		cfg = cfgs[len(cfgs)-1]
	}

	chalkYamlConfig, chalkYamlErr := auth2.GetProjectAuthConfig()

	apiServerOverride := auth2.GetChalkClientArgConfig(cfg.ApiServer)
	clientIdOverride := auth2.GetChalkClientArgConfig(cfg.ClientId)
	clientSecretOverride := auth2.GetChalkClientArgConfig(cfg.ClientSecret)
	environmentIdOverride := auth2.GetChalkClientArgConfig(cfg.EnvironmentId)

	apiServerEnvVarConfig := auth2.GetEnvVarConfig(internal.ApiServerEnvVarKey)
	clientIdEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientIdEnvVarKey)
	clientSecretEnvVarConfig := auth2.GetEnvVarConfig(internal.ClientSecretEnvVarKey)
	environmentIdEnvVarConfig := auth2.GetEnvVarConfig(internal.EnvironmentEnvVarKey)

	apiServerFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ApiServer)
	clientIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientId)
	clientSecretFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ClientSecret)
	environmentIdFileConfig := auth2.GetChalkYamlConfig(chalkYamlConfig.ActiveEnvironment)

	apiServer := auth2.GetFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig)
	clientId := auth2.GetFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig)
	clientSecret := auth2.GetFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig)
	environmentId := auth2.GetFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig)

	if chalkYamlErr != nil && clientId.Value == "" && clientSecret.Value == "" {
		return nil, chalkYamlErr
	}

	client := &clientImpl{
		ClientId:      clientId,
		ApiServer:     apiServer,
		EnvironmentId: environmentId,
		Branch:        cfg.Branch,

		logger:             cfg.Logger,
		httpClient:         cfg.HTTPClient,
		clientSecret:       clientSecret,
		initialEnvironment: environmentId,
	}

	if client.logger == nil {
		client.logger = DefaultLeveledLogger
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}

	err := client.refreshJwt(false)
	if cfg.Logger != nil {
		client.logger = cfg.Logger
	}
	if err != nil {
		return nil, err
	}
	return client, nil
}
