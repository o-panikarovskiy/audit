package utils

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// AppError represents http error
type AppError struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Details interface{}  `json:"details"`
	Err     error        `json:"cause"`
	Stack   []StackFrame `json:"stack"`
}

// StackFrame represents stack frame
type StackFrame struct {
	File     string `json:"file"`
	FuncName string `json:"func"`
	Line     int    `json:"line"`
}

// NewAppError returns an error with error code and error messages provided in
// function params
func NewAppError(code string, msg ...string) *AppError {
	e := AppError{Code: code}

	msgCount := len(msg)
	if msgCount > 0 {
		e.Message = msg[0]
	}

	if msgCount > 1 {
		e.Details = msg[1:]
	}

	return &e
}

// Error returns a string representation of AppError. It includes at least
// error status, code and message.
func (e *AppError) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Message)
}

// GetErrorStack wraps error and return stack frames
func GetErrorStack(err error, skip int) []StackFrame {
	err = errors.WithStack(err)
	stackError, ok := err.(interface{ StackTrace() errors.StackTrace })
	if !ok {
		return nil
	}

	stack := stackError.StackTrace()
	result := make([]StackFrame, 0, len(stack)-skip)

	for i := skip; i < len(stack); i++ {
		frame := stack[i]
		sframe := StackFrame{}

		sframe.File = fmt.Sprintf("%s", frame)
		sframe.FuncName = fmt.Sprintf("%n", frame)
		line, le := strconv.Atoi(fmt.Sprintf("%d", frame))
		if le == nil {
			sframe.Line = line
		}

		result = append(result, sframe)
	}

	return result
}
