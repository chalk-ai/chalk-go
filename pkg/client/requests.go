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

func (c *Client) sendRequest(params requestParams) error {
	params.Request.Header.Set("content-type", "application/json")
	params.Request.Header.Set("accept", "application/json")
	if c.EnvironmentId.Value != "" {
		params.Request.Header.Set("x-chalk-env-id", c.EnvironmentId.Value)
	}

	cfg, cfgErr := project.LoadProjectConfig()
	if cfgErr == nil && cfg.Project != "" {
		params.Request.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if cfg.Project != "" {
		params.Request.Header.Set("x-chalk-project-name", cfg.Project)
	}

	if !params.DontRefresh {
		upsertJwtErr := c.upsertJwt()
		logrus.Debug(fmt.Sprintf("Error pre-emptively refreshing access token: %s", upsertJwtErr.Error()))

	}
	if c.jwt != nil && c.jwt.Token != "" {
		params.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt.Token))
	}

	if !strings.HasPrefix(params.Request.URL.String(), "http:") && !strings.HasPrefix(params.Request.URL.String(), "https:") {
		var err error
		params.Request.URL, err = url.Parse(fmt.Sprintf("%s/%s", c.ApiServer.Value, params.Request.URL.String()))
		if err != nil {
			return err
		}
	}

	logrus.Debug("Sending params.Requestuest to ", params.Request.URL)

	res, err := c.httpClient.Do(params.Request)
	if err != nil {
		return err
	}

	logrus.Debug("Response Status: ", res.Status)
	defer res.Body.Close()

	if res.StatusCode == 401 && !params.DontRefresh {
		upsertJwtUpon401Err := c.upsertJwt()
		if upsertJwtUpon401Err != nil {
			logrus.Debug(fmt.Sprintf("Error refreshing access token upon 401: %s", upsertJwtUpon401Err.Error()))
		} else {
			res, err = c.httpClient.Do(params.Request)
			if err != nil {
				return err
			}
			logrus.Debug("Response Status for retry params.Requestuest: ", res.Status)
		}
	}

	if res.StatusCode != 200 {
		clientError := getHttpError(*res, *params.Request)
		return &clientError
	}

	out, _ := io.ReadAll(res.Body)
	logrus.Trace("Body: ", string(out))
	err = json.Unmarshal(out, &params.Response)

	return err
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
