package validationerrorformat

import (
	"gopkg.in/go-playground/validator.v9"
)

type validationErrorItem struct {
	Tag string `json:"tag"`
	Field string `json:"field"`
}

type validationErrorObj struct {
	Message string `json:"message"`
    Errors []validationErrorItem `json:"errors"`
}

type ValidationErrorFormat struct {
	Error validationErrorObj `json:"error"`
}

func New(errs validator.ValidationErrors) *ValidationErrorFormat {
	errorItems := make([]validationErrorItem, len(errs))
	for i, e := range errs {
		errorItems[i] = validationErrorItem {
			Tag: e.Tag(),
			Field:  e.Field(),
		}
	}

	return &ValidationErrorFormat{
		Error: validationErrorObj {
			Message: "ValidationErrors",
			Errors: errorItems,
		},
	}
}