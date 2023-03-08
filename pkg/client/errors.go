package client

import (
	"fmt"
	"strings"
)

type ChalkErrorResponse struct {
	ChalkErrors []ChalkError
}

func (e *ChalkErrorResponse) Error() string {
	stringifiedErrors := make([]string, 0)
	for _, err := range e.ChalkErrors {
		stringifiedErrors = append(stringifiedErrors, err.Error())
	}

	return strings.Join(stringifiedErrors, "\n")
}

type ChalkError struct {
	Code      string          `json:"code"`
	Category  string          `json:"category"`
	Message   string          `json:"message"`
	Exception *ChalkException `json:"exception"`
	Feature   string          `json:"feature"`
	Resolver  string          `json:"resolver"`
}

func (e *ChalkError) Error() string {
	detailArr := make([]string, 0)
	if e.Message != "" {
		detailArr = append(detailArr, "Message: "+e.Message)
	}
	if e.Exception.Message != "" {
		detailArr = append(detailArr, "Exception: "+e.Exception.Message)
	}
	if e.Exception.Stacktrace != "" {
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

type ClientError struct {
	Path          string
	Message       string
	StatusCode    int
	ContentLength int64
	Trace         *string
}

func (e *ClientError) Error() string {
	if e.Trace != nil {
		return fmt.Sprintf("httpClient Error: path=%q, message=%q, status=%d, content-length=%d, trace=%q",
			e.Path, e.Message, e.StatusCode, e.ContentLength, *e.Trace)
	}
	return fmt.Sprintf("httpClient Error: path=%q, message=%q, status=%d, content-length=%d",
		e.Path, e.Message, e.StatusCode, e.ContentLength)
}
