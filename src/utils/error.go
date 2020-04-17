package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// APIError is main error struct
type APIError struct {
	Status  int       `json:"status"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Details StringMap `json:"details"`
	Stack   []string  `json:"stack"`
}

func (error *APIError) Error() string {
	return fmt.Sprintf("%v-%v: %v", error.Status, error.Code, error.Message)
}

// NewAPIError returns APIError
func NewAPIError(status int, code string, msg string, details ...StringMap) error {
	var dt StringMap

	for _, item := range details {
		for key, value := range item {
			dt[key] = value
		}
	}

	return &APIError{
		Status:  status,
		Code:    code,
		Message: msg,
		Details: dt,
		Stack:   *stackTrace(1),
	}
}

func stackTrace(skip int) *[]string {
	const maxSize = 50

	stack := make([]uintptr, maxSize)
	length := runtime.Callers(2+skip, stack[:])
	stack = stack[:length]

	if len(stack) == 0 {
		return nil
	}

	curdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	res := make([]string, len(stack)-1)
	frames := runtime.CallersFrames(stack)
	frame, more := frames.Next()
	for i := 0; more && i < len(res); i++ {
		res[i] = fmt.Sprintf(
			"%s %s:%d",
			strings.Replace(frame.Function, curdir, ``, 1),
			strings.Replace(frame.File, curdir, ``, 1),
			frame.Line,
		)

		frame, more = frames.Next()
	}

	return &res
}
