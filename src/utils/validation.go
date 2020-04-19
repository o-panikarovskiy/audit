package utils

import (
	"fmt"
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
