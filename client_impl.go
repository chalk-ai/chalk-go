package chalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	auth2 "github.com/chalk-ai/chalk-go/internal/auth"
	"github.com/chalk-ai/chalk-go/internal/project"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type clientImpl struct {
	ApiServer     auth2.SourcedConfig
	ClientId      auth2.SourcedConfig
	EnvironmentId auth2.SourcedConfig

	clientSecret auth2.SourcedConfig
	jwt          *auth2.JWT
	httpClient   *http.Client
	logger       *LeveledLogger
}

func (c *clientImpl) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, *ErrorResponse) {
	emptyResult := OnlineQueryResult{}

	if request.EnvironmentId == "" {
		request.EnvironmentId = c.EnvironmentId.Value
	}

	var serializedResponse onlineQueryResponseSerialized

	err := c.sendRequest(sendRequestParams{Method: "POST", URL: "v1/query/online", Body: request.serialize(), Response: &serializedResponse})
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

	return response, nil
}

func (c *clientImpl) TriggerResolverRun(request TriggerResolverRunParams) (TriggerResolverRunResult, *ErrorResponse) {
	response := TriggerResolverRunResult{}
	err := c.sendRequest(sendRequestParams{Method: "POST", URL: "v1/runs/trigger", Body: request, Response: &response})
	if err != nil {
		return TriggerResolverRunResult{}, getErrorResponse(err)
	}
	return response, nil
}

func (c *clientImpl) GetRunStatus(request GetRunStatusParams) (GetRunStatusResult, *ErrorResponse) {
	response := GetRunStatusResult{}
	err := c.sendRequest(sendRequestParams{Method: "GET", URL: fmt.Sprintf("v1/runs/%s", request.RunId), Body: request, Response: &response})
	if err != nil {
		return GetRunStatusResult{}, getErrorResponse(err)
	}
	return response, nil
}

func (c *clientImpl) getJwt() (*auth2.JWT, *ClientError) {
	body := getTokenRequest{
		ClientId:     c.ClientId.Value,
		ClientSecret: c.clientSecret.Value,
		GrantType:    "client_credentials",
	}
	response := getTokenResponse{}
	err := c.sendRequest(sendRequestParams{Method: "POST", URL: "v1/oauth/token", Body: body, Response: &response, DontRefresh: true})
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

func (c *clientImpl) sendRequest(args sendRequestParams) error {
	jsonBytes, jsonErr := json.Marshal(args.Body)
	if jsonErr != nil {
		return jsonErr
	}

	request, newRequestErr := http.NewRequest(args.Method, args.URL, bytes.NewBuffer(jsonBytes))
	if newRequestErr != nil {
		return newRequestErr
	}

	request.Header.Set("content-type", "application/json")
	request.Header.Set("accept", "application/json")
	if c.EnvironmentId.Value != "" {
		request.Header.Set("x-chalk-env-id", c.EnvironmentId.Value)
	}

	cfg, cfgErr := project.LoadProjectConfig()
	if cfgErr == nil && cfg.Project != "" {
		request.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if cfg.Project != "" {
		request.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if !args.DontRefresh {
		upsertJwtErr := c.refreshJwt(false)
		if upsertJwtErr != nil {
			(*c.logger).Debugf(fmt.Sprintf("Error pre-emptively refreshing access token: %s", upsertJwtErr.Error()))
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

	(*c.logger).Debugf("Sending request to ", request.URL)
	res, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	(*c.logger).Debugf("Response Status: ", res.Status)
	defer res.Body.Close()

	if res.StatusCode == 401 && !args.DontRefresh && request != nil {
		res, err = c.retryRequest(*request, jsonBytes, res, err)
		if err != nil {
			return err
		}
	}

	if res.StatusCode != 200 {
		clientError := getHttpError(c.logger, *res, *request)
		return &clientError
	}

	out, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(out, &args.Response)

	return err
}

func (c *clientImpl) retryRequest(
	originalRequest http.Request, originalBodyBytes []byte,
	originalResponse *http.Response, originalError error,
) (*http.Response, error) {
	upsertJwtUpon401Err := c.refreshJwt(true)
	if upsertJwtUpon401Err != nil {
		(*c.logger).Debugf("Error refreshing access token upon 401: %s", upsertJwtUpon401Err.Error())
		return originalResponse, originalError
	}

	// New request needs to be constructed otherwise we were getting the error:
	//     HTTP/1.x transport connection broken
	newRequest, err := http.NewRequest(originalRequest.Method, originalRequest.URL.String(), bytes.NewBuffer(originalBodyBytes))
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
	(*c.logger).Debugf("Response Status for retried request: ", res.Status)

	return res, nil
}

func getHttpError(logger *LeveledLogger, res http.Response, req http.Request) HTTPError {
	var errorResponse chalkHttpException
	out, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(out, &errorResponse)
	(*logger).Errorf("API error response", err, errorResponse, string(out))

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
	var cfg = (*ClientConfig)(nil)
	if len(cfgs) == 0 {
		cfg = &ClientConfig{}
	} else {
		cfg = cfgs[len(cfgs)]
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

	clientId := auth2.GetFirstNonEmptyConfig(clientIdOverride, clientIdEnvVarConfig, clientIdFileConfig)
	clientSecret := auth2.GetFirstNonEmptyConfig(clientSecretOverride, clientSecretEnvVarConfig, clientSecretFileConfig)
	if chalkYamlErr != nil && clientId.Value == "" && clientSecret.Value == "" {
		return nil, chalkYamlErr
	}

	client := &clientImpl{
		httpClient:    cfg.HTTPClient,
		logger:        cfg.Logger,
		ClientId:      clientId,
		clientSecret:  clientSecret,
		ApiServer:     auth2.GetFirstNonEmptyConfig(apiServerOverride, apiServerEnvVarConfig, apiServerFileConfig),
		EnvironmentId: auth2.GetFirstNonEmptyConfig(environmentIdOverride, environmentIdEnvVarConfig, environmentIdFileConfig),
	}

	if client.logger == nil {
		client.logger = &DefaultLeveledLogger
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
