package chalk

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"strings"
)

func (e *ErrorResponse) Error() string {
	if len(e.ServerErrors) > 0 {
		stringifiedServerErrors := make([]string, 0)
		for _, err := range e.ServerErrors {
			stringifiedServerErrors = append(stringifiedServerErrors, err.Error())
		}
		return fmt.Sprintf("Server Errors: %s", strings.Join(stringifiedServerErrors, "\n"))
	} else if e.HttpError != nil {
		return e.HttpError.Error()
	} else if e.ClientError != nil {
		return e.ClientError.Error()
	} else {
		return "Unexpected chalk.Client error. Please contact Chalk if this persists."
	}
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("Client Error: %s", e.Message)
}

func (e *ServerError) Error() string {
	detailArr := make([]string, 0)
	if e.Message != "" {
		detailArr = append(detailArr, "Message: "+e.Message)
	}
	if e.Exception != nil && e.Exception.Message != "" {
		detailArr = append(detailArr, "Exception: "+e.Exception.Message)
	}
	if e.Exception != nil && e.Exception.Stacktrace != "" {
		detailArr = append(detailArr, "Stacktrace: "+e.Exception.Stacktrace)
	}
	if e.Resolver != "" {
		detailArr = append(detailArr, "Resolver: "+e.Resolver)
	}
	if e.Feature != "" {
		detailArr = append(detailArr, "Feature: "+e.Feature)
	}

	details := ""
	if len(detailArr) > 0 {
		contents := strings.Join(detailArr, " | ")
		details = ": [ " + contents + " ]"
	}
	return fmt.Sprintf("Chalk Error occurred%s", details)
}

func (e *HTTPError) Error() string {
	if e.Trace != nil {
		return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d, trace=%q",
			e.Path, e.Message, e.StatusCode, e.ContentLength, *e.Trace)
	}
	return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d",
		e.Path, e.Message, e.StatusCode, e.ContentLength)
}

func deferFunctionWithError(function func() error, originalError error) error {
	err := originalError
	if bodyCloseErr := function(); bodyCloseErr != nil && err == nil {
		err = bodyCloseErr
	}
	return err
}

func newServerError(errs []ServerError) *ErrorResponse {
	return &ErrorResponse{ServerErrors: errs}
}

func newClientError(err error) *ErrorResponse {
	return &ErrorResponse{ClientError: &ClientError{Message: err.Error()}}
}

func wrapClientError(err error, message string) *ErrorResponse {
	return newClientError(errors.Wrap(err, message))
}
