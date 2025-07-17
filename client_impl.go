package chalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow/memory"
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
	config    *configManager
	allocator memory.Allocator

	Branch        string
	QueryServer   string
	DeploymentTag string
	resourceGroup *string

	timeout    *time.Duration
	httpClient HTTPClient

	logger LeveledLogger
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
}

func (c *clientImpl) OfflineQuery(ctx context.Context, params OfflineQueryParamsComplete) (Dataset, error) {
	request := params.underlying

	resolved, err := request.resolve()
	if err != nil {
		return Dataset{}, errors.Wrap(err, "resolving params")
	}

	body, err := serializeOfflineQueryParams(&request, resolved)
	if err != nil {
		return Dataset{}, errors.Wrap(err, "serializing offline query params")
	}

	response := Dataset{}
	if err = c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:              "POST",
			URL:                 "v4/offline_query",
			Body:                body,
			Response:            &response,
			EnvironmentOverride: request.EnvironmentId,
			Versioned:           resolved.versioned,
			Branch:              &request.Branch,
		},
	); err != nil {
		return Dataset{}, errors.Wrap(err, "sending request")
	}

	if len(response.Errors) > 0 {
		return Dataset{}, response.Errors
	}

	// Set the environment ID
	if response.EnvironmentID == "" {
		response.EnvironmentID = request.EnvironmentId
		if response.EnvironmentID == "" {
			response.EnvironmentID = c.config.environmentId.Value
		}
	}

	// Set the client reference for Wait() method
	response.client = c

	for idx := range response.Revisions {
		response.Revisions[idx].client = c
		// Set environment ID for each revision if not already set
		if response.Revisions[idx].EnvironmentID == "" {
			response.Revisions[idx].EnvironmentID = response.EnvironmentID
		}
	}

	return response, nil
}

func (c *clientImpl) OnlineQueryBulk(ctx context.Context, params OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	request := params.underlying

	resolved, err := request.resolveBulk()
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Newf("resolving bulk query params")
	}

	for _, input := range resolved.inputs {
		kind := reflect.ValueOf(input).Kind()
		if !(kind == reflect.Slice || kind == reflect.Array) {
			return OnlineQueryBulkResult{}, errors.Newf(
				"inputs to bulk online query must be a slice or array, found: %s", kind.String(),
			)
		}
	}
	data, err := params.ToBytes(
		&SerializationOptions{
			ClientConfigBranchId: c.Branch,
			Allocator:            c.allocator,
			resolved:             resolved,
		},
	)
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "serializing online query params")
	}

	var resourceGroupOverride *string
	if params.underlying.ResourceGroup != "" {
		resourceGroupOverride = &params.underlying.ResourceGroup
	}

	response := OnlineQueryBulkResponse{allocator: c.allocator}
	err = c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:                "POST",
			URL:                   "v1/query/feather",
			Body:                  data,
			Response:              &response,
			EnvironmentOverride:   params.underlying.EnvironmentId,
			PreviewDeploymentId:   params.underlying.PreviewDeploymentId,
			ResourceGroupOverride: resourceGroupOverride,
			Versioned:             resolved.versioned,
			Branch:                params.underlying.BranchId,
			IsEngineRequest:       true,
		},
	)

	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "sending request")
	}

	singleBulkResult, ok := response.QueryResults["0"]
	if !ok {
		return OnlineQueryBulkResult{}, errors.New("unexpected bulk online query response from server")
	}

	if len(singleBulkResult.Errors) > 0 {
		return OnlineQueryBulkResult{}, singleBulkResult.Errors
	}

	return OnlineQueryBulkResult{
		ScalarsTable: singleBulkResult.ScalarData,
		GroupsTables: singleBulkResult.GroupsData,
		Meta:         singleBulkResult.Meta,
		allocator:    c.allocator,
	}, nil
}

func (c *clientImpl) UploadFeatures(ctx context.Context, params UploadFeaturesParams) (UploadFeaturesResult, error) {
	convertedInputs, err := getConvertedInputsMap(params.Inputs)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrap(err, "converting input map keys")
	}
	recordBytes, err := internal.InputsToArrowBytes(convertedInputs, c.allocator)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrap(err, "converting inputs to Arrow Record bytes")
	}
	attrs := map[string]any{
		"features":          colls.Keys(convertedInputs),
		"table_compression": "uncompressed",
		"table_bytes":       recordBytes,
	}

	body, err := internal.ChalkMarshal(attrs)
	if err != nil {
		return UploadFeaturesResult{}, errors.Wrap(err, "marshaling upload features request")
	}

	response := UploadFeaturesResult{}
	err = c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:              "POST",
			URL:                 "v1/upload_features/multi",
			Body:                body,
			Response:            &response,
			EnvironmentOverride: params.EnvironmentOverride,
			PreviewDeploymentId: params.PreviewDeploymentId,
			IsEngineRequest:     true,
		},
	)
	return response, errors.Wrap(err, "sending request")
}

func (c *clientImpl) OnlineQuery(ctx context.Context, params OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	request := params.underlying

	resolved, err := request.resolveSingle()
	if err != nil {
		return OnlineQueryResult{}, errors.Wrap(err, "resolving single query params")
	}

	serializedRequest, err := serializeOnlineQueryParams(&request, resolved)
	if err != nil {
		return OnlineQueryResult{}, errors.Wrap(err, "serializing online query params")
	}

	var resourceGroupOverride *string
	if params.underlying.ResourceGroup != "" {
		resourceGroupOverride = &params.underlying.ResourceGroup
	}

	var response onlineQueryResponseSerialized
	if err = c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:                "POST",
			URL:                   "v1/query/online",
			Body:                  *serializedRequest,
			Response:              &response,
			EnvironmentOverride:   request.EnvironmentId,
			PreviewDeploymentId:   request.PreviewDeploymentId,
			Versioned:             resolved.versioned,
			Branch:                params.underlying.BranchId,
			ResourceGroupOverride: resourceGroupOverride,
			IsEngineRequest:       true,
		},
	); err != nil {
		return OnlineQueryResult{}, errors.Wrap(err, "sending request")
	}
	if len(response.Errors) > 0 {
		serverErrors, err := deserializeChalkErrors(response.Errors)
		if err != nil {
			return OnlineQueryResult{}, errors.Wrap(err, "deserializing Chalk errors")
		}

		return OnlineQueryResult{}, serverErrors
	}

	features := make(map[string]FeatureResult)

	deserializedData, err := deserializeFeatureResults(response.Data)
	if err != nil {
		return OnlineQueryResult{}, errors.Wrap(err, "deserializing feature results")
	}

	for _, result := range deserializedData {
		features[result.Field] = result
	}

	result := OnlineQueryResult{
		Data:      deserializedData,
		Meta:      response.Meta,
		features:  features,
		allocator: c.allocator,
	}

	if resultHolder != nil {
		if err := result.UnmarshalInto(resultHolder); err != nil {
			return result, errors.Wrap(err, "unmarshaling result")
		}
	}

	return result, nil
}

func (c *clientImpl) TriggerResolverRun(ctx context.Context, request TriggerResolverRunParams) (TriggerResolverRunResult, error) {
	response := TriggerResolverRunResult{}
	err := c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:              "POST",
			URL:                 "v1/runs/trigger",
			Body:                request,
			Response:            &response,
			EnvironmentOverride: request.EnvironmentId,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	return response, errors.Wrap(err, "triggering resolver run")
}

func (c *clientImpl) GetRunStatus(ctx context.Context, request GetRunStatusParams) (GetRunStatusResult, error) {
	response := GetRunStatusResult{}
	err := c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:              "GET",
			URL:                 fmt.Sprintf("v1/runs/%s", request.RunId),
			Body:                request,
			Response:            &response,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	return response, errors.Wrap(err, "getting run status")
}

func (c *clientImpl) getDatasetUrls(ctx context.Context, RevisionId string, EnvironmentId string) ([]string, error) {
	response := GetOfflineQueryJobResponse{}
	for !response.IsFinished {
		err := c.sendRequest(
			ctx,
			&sendRequestParams{
				Method:              "GET",
				URL:                 fmt.Sprintf("v2/offline_query/%s", RevisionId),
				EnvironmentOverride: EnvironmentId,
				Response:            &response,
			},
		)
		if err != nil {
			return []string{}, errors.Wrap(err, "getting dataset urls")
		}
		time.Sleep(500 * time.Millisecond)
	}

	return response.Urls, nil
}

func (c *clientImpl) saveUrlToDirectory(URL string, directory string) (err error) {
	resp, err := c.httpClient.Get(URL)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		}
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
	return errors.Wrapf(os.WriteFile(destinationPath, data, os.ModePerm), "saving file to %s", destinationPath)
}

func (c *clientImpl) GetToken(ctx context.Context) (*TokenResult, error) {
	res, err := c.getToken(ctx, c.config.clientId.Value, c.config.clientSecret.Value)
	if err != nil {
		return nil, err // Intentionally not wrapping
	}
	return &TokenResult{
		AccessToken:        res.AccessToken,
		PrimaryEnvironment: res.PrimaryEnvironment,
		ValidUntil:         res.ValidUntil,
		Engines:            res.Engines,
	}, nil
}

func (c *clientImpl) getToken(ctx context.Context, clientId string, clientSecret string) (*getTokenResult, error) {
	body := getTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}
	response := getTokenResponse{}
	err := c.sendRequest(
		ctx,
		&sendRequestParams{
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

func (c *clientImpl) sendRequest(ctx context.Context, args *sendRequestParams) error {
	body, getBufferErr := getBodyBuffer(args.Body)
	if getBufferErr != nil {
		return getBufferErr
	}

	ctx, cancel := internal.GetContextWithTimeout(ctx, c.timeout)
	defer cancel()
	request, newRequestErr := http.NewRequestWithContext(ctx, args.Method, args.URL, body)
	if newRequestErr != nil {
		(c.logger).Debugf("error sending request: %s", newRequestErr.Error())
		return newRequestErr
	}

	headers := c.getHeaders(args.EnvironmentOverride, args.PreviewDeploymentId, args.Branch, args.ResourceGroupOverride)
	request.Header = headers

	if !args.DontRefresh {
		if err := c.config.refresh(ctx, false); err != nil {
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
		res, err = c.retryRequest(ctx, *request, args.Body, res, err)
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
		return castResponse.Unmarshal(out)
	}
	return json.Unmarshal(out, args.Response)
}

func (c *clientImpl) retryRequest(
	ctx context.Context,
	originalRequest http.Request, originalBody any,
	originalResponse *http.Response, originalError error,
) (*http.Response, error) {
	if err := c.config.refresh(ctx, true); err != nil {
		(c.logger).Debugf("Error refreshing access token upon 401: %s", err.Error())
		return originalResponse, originalError
	}

	originalBodyBuffer, getBufferErr := getBodyBuffer(originalBody)
	if getBufferErr != nil {
		return nil, getBufferErr
	}

	// New request needs to be constructed otherwise we were getting the error:
	//     HTTP/1.x transport connection broken
	ctx, cancel := internal.GetContextWithTimeout(ctx, c.timeout)
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
		if err != nil {
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

func (c *clientImpl) GetOfflineQueryStatus(ctx context.Context, request GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
	response := GetOfflineQueryStatusResult{}
	err := c.sendRequest(
		ctx,
		&sendRequestParams{
			Method:              "GET",
			URL:                 fmt.Sprintf("v4/offline_query/%s/status", request.JobId),
			Response:            &response,
			PreviewDeploymentId: request.PreviewDeploymentId,
		},
	)
	return response, errors.Wrap(err, "getting offline query status")
}

func newClientImpl(
	ctx context.Context,
	cfg *ClientConfig,
) (*clientImpl, error) {
	config, err := newConfigManager(
		cfg.ApiServer,
		cfg.ClientId,
		cfg.ClientSecret,
		cfg.EnvironmentId,
		cfg.Logger,
	)
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

	allocator := memory.DefaultAllocator
	if cfg.Allocator != nil {
		allocator = cfg.Allocator
	}

	client := &clientImpl{
		Branch:        cfg.Branch,
		DeploymentTag: cfg.DeploymentTag,
		QueryServer:   cfg.QueryServer,
		resourceGroup: resourceGroup,
		timeout:       timeout,

		logger:     logger,
		httpClient: httpClient,
		allocator:  allocator,

		config: config,
	}
	client.config.getToken = client.getToken
	return client, errors.Wrap(client.config.refresh(ctx, false), "refreshing config on client creation")
}
