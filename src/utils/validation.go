package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateModel check model by struct
func ValidateModel(model interface{}) error {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(model)
	if err != nil {
		return BadRequestModel(err)
	}

	return nil
}

// BadRequestModel returns APPError
func BadRequestModel(err interface{}) *APPError {
	var details *StringMap
	message := "Invalid request model"

	switch e := err.(type) {
	case *APPError:
		return e
	case string:
		message = e
	case validator.ValidationErrors:
		details = parseValidationErrors(e)
	case error:
		message = e.Error()
	}

	return &APPError{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_REQUEST_MODEL",
		Message: message,
		Details: details,
		Stack:   *stackTrace(1),
	}
}

func parseValidationErrors(err validator.ValidationErrors) *StringMap {
	var sb strings.Builder
	sm := make(StringMap)
	for _, field := range err {
		sb.Reset()
		sb.WriteString("condition: " + field.ActualTag())
		if prm := field.Param(); prm != "" {
			sb.WriteString("(" + prm + ")")
		}
		if val := field.Value(); val != "" && val != nil {
			sb.WriteString(fmt.Sprintf("; actual: %v", val))
		}
		sm[field.Field()] = sb.String()
	}
	return &sm
}
