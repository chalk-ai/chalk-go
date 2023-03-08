package client

import (
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/project"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) sendRequest(req *http.Request, response any) error {
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	if c.EnvironmentId != "" {
		req.Header.Set("x-chalk-env-id", c.EnvironmentId)
	}

	cfg, cfgErr := project.LoadProjectConfig()
	if cfgErr == nil && cfg.Project != "" {
		req.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if cfg.Project != "" {
		req.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if c.jwt != nil && c.jwt.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.Token))
	}

	if !strings.HasPrefix(req.URL.String(), "http:") && !strings.HasPrefix(req.URL.String(), "https:") {
		var err error
		req.URL, err = url.Parse(fmt.Sprintf("%s/%s", c.BaseUrl, req.URL.String()))
		if err != nil {
			return err
		}
	}

	logrus.Debug("Sending request to ", req.URL)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	logrus.Debug("Response Status: ", res.Status)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		clientError := getClientError(*res, *req)
		return &clientError
	}

	out, _ := io.ReadAll(res.Body)
	logrus.Trace("Body: ", string(out))
	err = json.Unmarshal(out, &response)

	return err
}

func getClientError(res http.Response, req http.Request) ClientError {
	var errorResponse chalkHttpException
	out, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(out, &errorResponse)
	logrus.Debug("API error response", err, errorResponse, string(out))

	clientError := ClientError{
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
