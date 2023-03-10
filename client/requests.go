package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/project"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) sendRequest(args sendRequestParams) error {
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
			logrus.Debug(fmt.Sprintf("Error pre-emptively refreshing access token: %s", upsertJwtErr.Error()))
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

	logrus.Debug("Sending request to ", request.URL)
	res, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	logrus.Debug("Response Status: ", res.Status)
	defer res.Body.Close()

	if res.StatusCode == 401 && !args.DontRefresh && request != nil {
		res, err = c.retryRequest(*request, jsonBytes, res, err)
		if err != nil {
			return err
		}
	}

	if res.StatusCode != 200 {
		clientError := getHttpError(*res, *request)
		return &clientError
	}

	out, _ := io.ReadAll(res.Body)
	logrus.Trace("Body: ", string(out))
	err = json.Unmarshal(out, &args.Response)

	return err
}

func (c *Client) retryRequest(originalRequest http.Request, originalBodyBytes []byte, originalResponse *http.Response, originalError error) (*http.Response, error) {
	upsertJwtUpon401Err := c.refreshJwt(true)
	if upsertJwtUpon401Err != nil {
		logrus.Debug(fmt.Sprintf("Error refreshing access token upon 401: %s", upsertJwtUpon401Err.Error()))
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
	logrus.Debug("Response Status for retried request: ", res.Status)

	return res, nil
}

func getHttpError(res http.Response, req http.Request) ChalkHttpError {
	var errorResponse chalkHttpException
	out, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(out, &errorResponse)
	logrus.Debug("API error response", err, errorResponse, string(out))

	clientError := ChalkHttpError{
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
