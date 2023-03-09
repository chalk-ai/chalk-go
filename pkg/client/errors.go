package client

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/client/clientenums"
	"strings"
)

type ChalkErrorResponse struct {
	ServerErrors []ChalkServerError
	ClientError  *ChalkClientError
}

func (e *ChalkErrorResponse) Error() string {
	stringifiedErrors := make([]string, 0)
	for _, err := range e.ServerErrors {
		stringifiedErrors = append(stringifiedErrors, err.Error())
	}

	return strings.Join(stringifiedErrors, "\n")
}

type ChalkClientError struct {
	Message string
}

func (e *ChalkClientError) Error() string {
	return e.Message
}

type ChalkServerError struct {
	Code      clientenums.ErrorCode
	Category  clientenums.ErrorCodeCategory
	Message   string
	Exception *ChalkException
	Feature   string
	Resolver  string
}

func (e *ChalkServerError) Error() string {
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

type chalkErrorSerialized struct {
	Code      string          `json:"code"`
	Category  string          `json:"category"`
	Message   string          `json:"message"`
	Exception *ChalkException `json:"exception"`
	Feature   string          `json:"feature"`
	Resolver  string          `json:"resolver"`
}

type ChalkHttpError struct {
	Path          string
	Message       string
	StatusCode    int
	ContentLength int64
	Trace         *string
}

func (e *ChalkHttpError) Error() string {
	if e.Trace != nil {
		return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d, trace=%q",
			e.Path, e.Message, e.StatusCode, e.ContentLength, *e.Trace)
	}
	return fmt.Sprintf("HTTP Error: path=%q, message=%q, status=%d, content-length=%d",
		e.Path, e.Message, e.StatusCode, e.ContentLength)
}
