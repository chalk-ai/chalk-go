package chalk

import (
	"fmt"
	"strings"
)

type serverErrorsT []ServerError

func (e serverErrorsT) Error() string {
	stringifiedServerErrors := make([]string, 0)
	for _, err := range e {
		stringifiedServerErrors = append(stringifiedServerErrors, err.Error())
	}
	return strings.Join(stringifiedServerErrors, "\n")
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
	if e != nil && e.Trace != nil {
		return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d, trace=%q",
			e.Path, e.Message, e.StatusCode, e.ContentLength, *e.Trace)
	}
	return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d",
		e.Path, e.Message, e.StatusCode, e.ContentLength)
}
