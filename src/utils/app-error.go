package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

// AppError represents http error
type AppError struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
	Err     error       `json:"cause"`
	Stack   stack       `json:"stack"`
}

// stack is a slice of StackFrames representing a stack trace
type stack []stackFrame

// stackFrame  represents a single frame of a stack trace
type stackFrame struct {
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
}

// NewAppError returns an error with error code and error messages provided in
// function params
func NewAppError(status int, code string, ErrorMsg ...string) *AppError {
	e := AppError{Status: status, Code: code}

	msgCount := len(ErrorMsg)
	if msgCount > 0 {
		e.Message = ErrorMsg[0]
	}

	if msgCount > 1 {
		e.Details = ErrorMsg[1:]
	}

	return &e
}

// Error returns a string representation of AppError. It includes at least
// error status, code and message.
func (e *AppError) Error() string {
	return fmt.Sprintf("%v-%v: %v", e.Status, e.Code, e.Message)
}

// Wrap wraps error in AppError for nested errors
func (e *AppError) Wrap(err error) {
	e.Err = err
}

// Unwrap unwraps error in AppError
func (e *AppError) Unwrap() error {
	return e.Err
}

// WriteJSON write a json representation to http.ResponseWriter
func (e *AppError) WriteJSON(res http.ResponseWriter, status int, withStack bool) error {
	var err *AppError

	if withStack {
		err = e.subset()
		stack := getStack(4)
		err.Stack = stack
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(status)
	return json.NewEncoder(res).Encode(err)
}

func (e *AppError) subset() *AppError {
	err := &AppError{
		Status:  e.Status,
		Code:    e.Code,
		Message: e.Message,
		Details: e.Details,
	}

	cause := e.Unwrap()
	if cause != nil {
		if c, ok := cause.(*AppError); ok {
			err.Wrap(c.subset())
		} else {
			err.Wrap(&AppError{Message: cause.Error()})
		}
	}

	return err
}

func getStack(skip int) []stackFrame {
	stack := make([]stackFrame, 0)

	for i := skip; ; i++ {
		pc, fn, line, ok := runtime.Caller(i)
		if !ok {
			// no more frames - we're done
			break
		}
		_, fn = filepath.Split(fn)

		f := stackFrame{File: fn, Line: line, Function: funcName(pc)}
		stack = append(stack, f)
	}

	return stack
}

// funcName gets the name of the function at pointer or "??" if one can't be found
func funcName(pc uintptr) string {
	if f := runtime.FuncForPC(pc); f != nil {
		return f.Name()
	}
	return "??"
}
