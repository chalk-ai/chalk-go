package chalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
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
	Client
	config *configManager

	Branch        string
	QueryServer   string
	DeploymentTag string
	resourceGroup *string
	timeout       *time.Duration

	httpClient HTTPClient
	logger     LeveledLogger
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
}

func (c *clientImpl) GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error) {
	return nil, errors.New("not implemented")
}

func (c *clientImpl) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	return nil, errors.New("not implemented")
}

func (c *clientImpl) OfflineQuery(params OfflineQueryParamsComplete) (Dataset, error) {
	request := params.underlying

	if len(request.builderErrors) > 0 {
		return Dataset{}, errors.Wrapf(request.builderErrors, "building offline query params")
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
			Versioned:           params.underlying.versioned,
			Branch:              &request.Branch,
		},
	)
	if err != nil {
		return emptyResult, errors.Wrap(err, "send request")
	}

	if len(response.Errors) > 0 {
		return emptyResult, &ErrorResponse{ServerErrors: response.Errors}
	}

	for idx := range response.Revisions {
		response.Revisions[idx].client = c
	}

	return response, nil
}

func (c *clientImpl) OnlineQueryBulk(params OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	emptyResult := OnlineQueryBulkResult{}
	request := params.underlying

	if len(request.builderErrors) > 0 {
		return emptyResult, errors.Wrapf(request.builderErrors, "building params")
	}

	validationErrors := params.validatePostBuild()
	if len(validationErrors) > 0 {
		return emptyResult, validationErrors
	}

	for _, input := range request.inputs {
		kind := reflect.ValueOf(input).Kind()
		if !(kind == reflect.Slice || kind == reflect.Array) {
			return emptyResult, errors.Newf("Inputs to bulk online query must be a slice or array, found: ", kind.String())
		}
	}
	data, err := params.ToBytes(&SerializationOptions{ClientConfigBranchId: c.Branch})
	if err != nil {
		return emptyResult, errors.Wrapf(err, "serializing online query params")
	}

	var resourceGroupOverride *string
	if params.underlying.ResourceGroup != "" {
		resourceGroupOverride = &params.underlying.ResourceGroup
	}

	var response OnlineQueryBulkResponse
	err = c.sendRequest(
		sendRequestParams{
			Method:                "POST",
			URL:                   "v1/query/feather",
			Body:                  data,
			Response:              &response,
			EnvironmentOverride:   params.underlying.EnvironmentId,
			PreviewDeploymentId:   params.underlying.PreviewDeploymentId,
			ResourceGroupOverride: resourceGroupOverride,
			Versioned:             params.underlying.versioned,
			Branch:                params.underlying.BranchId,
			IsEngineRequest:       true,
		},
	)

	if err != nil {
		return emptyResult, errors.Wrap(err, "send request")
	}

	singleBulkResult, ok := response.QueryResults["0"]
	if !ok {
		return emptyResult, errors.Newf("unexpected bulk online query response from server")
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
	convertedInputs, err := getConvertedInputsMap(params.Inputs)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrapf(err, "convert input map keys")
	}
	recordBytes, err := internal.InputsToArrowBytes(convertedInputs)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrapf(err, "convert inputs to Arrow Record bytes")
	}
	attrs := map[string]any{
		"features":          colls.Keys(convertedInputs),
		"table_compression": "uncompressed",
		"table_bytes":       recordBytes,
	}

	body, err := internal.ChalkMarshal(attrs)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrapf(err, "marshal upload features request")
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
			IsEngineRequest:     true,
		},
	)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrap(err, "send request")
	}
	return response, nil
}

func (c *clientImpl) OnlineQuery(params OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	request := params.underlying

	if len(request.builderErrors) > 0 {
		return OnlineQueryResult{}, errors.Wrap(request.builderErrors, "building online query params")
	}

	validationErrors := params.validatePostBuild()
	if len(validationErrors) > 0 {
		return OnlineQueryResult{}, validationErrors
	}

	emptyResult := OnlineQueryResult{}

	var serializedResponse onlineQueryResponseSerialized

	serializedRequest, err := request.serialize()
	if err != nil {
		return emptyResult, errors.Wrap(err, "serializing online query params")
	}

	var resourceGroupOverride *string
	if params.underlying.ResourceGroup != "" {
		resourceGroupOverride = &params.underlying.ResourceGroup
	}

	if err = c.sendRequest(
		sendRequestParams{
			Method:                "POST",
			URL:                   "v1/query/online",
			Body:                  *serializedRequest,
			Response:              &serializedResponse,
			EnvironmentOverride:   request.EnvironmentId,
			PreviewDeploymentId:   request.PreviewDeploymentId,
			Versioned:             params.underlying.versioned,
			Branch:                params.underlying.BranchId,
			ResourceGroupOverride: resourceGroupOverride,
			IsEngineRequest:       true,
		},
	); err != nil {
		return emptyResult, errors.Wrap(err, "send request")
	}
	if len(serializedResponse.Errors) > 0 {
		serverErrors, deserializationErr := deserializeChalkErrors(serializedResponse.Errors)
		if deserializationErr != nil {
			return emptyResult, errors.Wrap(deserializationErr, "deserializing Chalk errors")
		}

		return emptyResult, &ErrorResponse{ServerErrors: serverErrors}
	}

	response, err := serializedResponse.deserialize()
	if err != nil {
		return emptyResult, errors.Wrap(err, "deserializing online query response")
	}

	response.expectedOutputs = params.underlying.outputs
	if resultHolder != nil {
		unmarshalErr := response.UnmarshalInto(resultHolder)
		if unmarshalErr != nil {
			return response, errors.Wrap(unmarshalErr, "unmarshaling result")
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
		return TriggerResolverRunResult{}, errors.Wrap(err, "send request")
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
		return GetRunStatusResult{}, errors.Wrap(err, "send request")
	}
	return response, nil
}

func (c *clientImpl) UpdateAggregates(params UpdateAggregatesParams) (UpdateAggregatesResult, error) {
	return UpdateAggregatesResult{}, errors.New("not implemented")
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
			return []string{}, errors.Wrap(err, "send request")
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

func (c *clientImpl) GetToken() (*TokenResult, error) {
	getTokenResult, err := c.getToken(c.config.clientId.Value, c.config.clientSecret.Value)
	if err != nil {
		return nil, err // Intentionally not wrapping
	}
	return &TokenResult{
		AccessToken:        getTokenResult.AccessToken,
		PrimaryEnvironment: getTokenResult.PrimaryEnvironment,
		ValidUntil:         getTokenResult.ValidUntil,
		Engines:            getTokenResult.Engines,
	}, nil
}

func (c *clientImpl) getToken(clientId string, clientSecret string) (*getTokenResult, error) {
	body := getTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
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
		return nil, errors.Newf(
			"Error obtaining access token: %s.\n"+
				"  Auth config:\n"+
				"    api_server=%q (source: %s),\n"+
				"    client_id=%q (source: %s),\n"+
				"    client_secret=*** (source: %s),\n"+
				"    environment_id=%q (source: %s)\n",
			err.Error(),
			c.config.apiServer.Value,
			c.config.apiServer.Source,
			c.config.clientId.Value,
			c.config.clientId.Source,
			c.config.clientSecret.Source,
			c.config.environmentId.Value,
			c.config.environmentId.Source,
		)
	}
	expiry := time.Now().UTC().Add(time.Duration(response.ExpiresIn) * time.Second)
	return &getTokenResult{
		ValidUntil:         expiry,
		AccessToken:        response.AccessToken,
		PrimaryEnvironment: response.PrimaryEnvironment,
		Engines:            response.Engines,
	}, nil
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

	ctx, cancel := internal.GetContextWithTimeout(context.Background(), c.timeout)
	defer cancel()
	request, newRequestErr := http.NewRequestWithContext(ctx, args.Method, args.URL, body)
	if newRequestErr != nil {
		(c.logger).Debugf("error sending request: %s", newRequestErr.Error())
		return newRequestErr
	}

	headers := c.getHeaders(args.EnvironmentOverride, args.PreviewDeploymentId, args.Branch, args.ResourceGroupOverride)
	request.Header = headers

	if !args.DontRefresh {
		if err := c.config.refresh(false); err != nil {
			(c.logger).Debugf("Error pre-emptively refreshing access token: %s", err)
		}
	}
	if c.config.jwt != nil && c.config.jwt.Token != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.jwt.Token))
	}
	if args.Versioned {
		request.Header.Set("X-Chalk-Features-Versioned", "true")
	}

	if !strings.HasPrefix(request.URL.String(), "http:") && !strings.HasPrefix(request.URL.String(), "https:") {
		var err error
		request.URL, err = url.Parse(fmt.Sprintf(
			"%s/%s",
			c.GetResolvedServer(args.EnvironmentOverride, args.IsEngineRequest),
			request.URL.String(),
		))
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
		clientError, err := getHttpError(c.logger, *res, *request)
		if err != nil {
			return errors.Wrap(err, "deserializing http error")
		}
		return clientError
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
	if err := c.config.refresh(true); err != nil {
		(c.logger).Debugf("Error refreshing access token upon 401: %s", err.Error())
		return originalResponse, originalError
	}

	originalBodyBuffer, getBufferErr := getBodyBuffer(originalBody)
	if getBufferErr != nil {
		return nil, getBufferErr
	}

	// New request needs to be constructed otherwise we were getting the error:
	//     HTTP/1.x transport connection broken
	ctx, cancel := internal.GetContextWithTimeout(context.Background(), c.timeout)
	defer cancel()
	newRequest, err := http.NewRequestWithContext(
		ctx,
		originalRequest.Method,
		originalRequest.URL.String(),
		originalBodyBuffer,
	)
	if err != nil {
		return nil, err
	}
	newRequest.Header = originalRequest.Header
	if c.config.jwt != nil && c.config.jwt.Token != "" {
		newRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.jwt.Token))
	}

	res, err := c.httpClient.Do(newRequest)
	if err != nil {
		return nil, err
	}
	(c.logger).Debugf("Response Status for retried request: ", res.Status)

	return res, nil
}

func (c *clientImpl) getResolvedEnvironment(envOverride string) string {
	if envOverride == "" {
		return c.config.environmentId.Value
	}
	return envOverride
}

func (c *clientImpl) GetResolvedServer(envOverride string, useQueryServer bool) string {
	if !useQueryServer {
		return c.config.apiServer.Value
	}

	if c.QueryServer != "" {
		return c.QueryServer
	}

	env := c.getResolvedEnvironment(envOverride)
	if engine, foundEngine := c.config.engines[env]; foundEngine && env != "" {
		return engine
	}

	return c.config.apiServer.Value
}

func (c *clientImpl) getHeaders(environmentOverride string, previewDeploymentId string, branchOverride *string, resourceGroupOverride *string) http.Header {
	headers := http.Header{}

	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", "chalk-go-0.0")
	headers.Set("X-Chalk-Client-Id", c.config.clientId.Value)

	var branchResolved string
	if branchOverride != nil && *branchOverride != "" {
		branchResolved = *branchOverride
	} else if c.Branch != "" {
		branchResolved = c.Branch
	}

	if branchResolved == "" {
		headers.Set("X-Chalk-Deployment-Type", "engine")
	} else {
		headers.Set("X-Chalk-Deployment-Type", "branch")
		headers.Set("X-Chalk-Branch-Id", branchResolved)
	}
	if c.DeploymentTag != "" {
		headers.Set("X-Chalk-Deployment-Tag", c.DeploymentTag)
	}

	headers.Set("X-Chalk-Env-Id", c.getResolvedEnvironment(environmentOverride))
	if previewDeploymentId != "" {
		headers.Set("X-Chalk-Preview-Deployment", previewDeploymentId)
	}

	if resourceGroupOverride != nil {
		headers.Set(HeaderKeyResourceGroup, *resourceGroupOverride)
	} else if c.resourceGroup != nil {
		headers.Set(HeaderKeyResourceGroup, *c.resourceGroup)
	}

	return headers
}

func getHttpError(logger LeveledLogger, res http.Response, req http.Request) (*HTTPError, error) {
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
		errDetailJson, err := json.Marshal(errorResponse.Detail)
		if err == nil {
			return nil, errors.Wrapf(
				err,
				"error marshalling error detail into string - original error detail: %v",
				errorResponse.Detail,
			)
		}
		clientError.Message = string(errDetailJson)
	}

	return &clientError, nil
}

func newClientImpl(
	cfg ClientConfig,
) (*clientImpl, error) {
	config, err := newConfigManager(cfg.ApiServer, cfg.ClientId, cfg.ClientSecret, cfg.EnvironmentId, cfg.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}

	logger := cfg.Logger
	if logger == nil {
		logger = DefaultLeveledLogger
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	var resourceGroup *string
	if cfg.ResourceGroup != "" {
		resourceGroup = &cfg.ResourceGroup
	}

	var timeout *time.Duration
	if cfg.Timeout != 0 { // If unspecified (zero value)
		timeout = &cfg.Timeout
	}

	client := &clientImpl{
		Branch:        cfg.Branch,
		DeploymentTag: cfg.DeploymentTag,
		QueryServer:   cfg.QueryServer,
		resourceGroup: resourceGroup,
		timeout:       timeout,

		logger:     logger,
		httpClient: httpClient,

		config: config,
	}
	client.config.getToken = client.getToken
	if err := client.config.refresh(false); err != nil {
		return nil, errors.Wrap(err, "error fetching initial config")
	}
	return client, nil
}
