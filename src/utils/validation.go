package utils

import (
	"github.com/go-playground/validator/v10"
)

type validationField struct {
	Field     string      `json:"field"`
	Condition string      `json:"condition"`
	Param     string      `json:"param,omitempty"`
	Actual    interface{} `json:"actual,omitempty"`
}

// ValidateModel check model by struct
func ValidateModel(model interface{}) error {
	v := validator.New()

	err := v.Struct(model)
	if err != nil {
		return &AppError{
			Err:     err,
			Code:    "INVALID_REQUEST_MODEL",
			Message: "Invalid request model",
			Details: getValidationErrors(err),
		}
	}

	return nil
}

func getValidationErrors(err error) *[]validationField {
	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	arr := make([]validationField, len(ve))

	for i, field := range ve {
		m := validationField{}

		m.Field = field.Field()
		m.Condition = field.ActualTag()

		if prm := field.Param(); prm != "" {
			m.Param = prm
		}

		if val := field.Value(); val != "" && val != nil {
			m.Actual = val
		}

		arr[i] = m
	}

	return &arr
}
