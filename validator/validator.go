package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorListValidation struct {
	Validation []ErrorListItem `json:"validation"`
	Entity     string          `json:"entity"`
}

type ErrorListItem struct {
	Field string `json:"field"`
	Key   string `json:"key"`
}

func (l *ErrorListValidation) Error() string {
	builder := strings.Builder{}
	builder.WriteString("In Struct ")
	builder.WriteString(l.Entity)
	builder.WriteString("\n")
	for _, i := range l.Validation {
		builder.WriteString("Field: ")
		builder.WriteString(i.Field)
		builder.WriteString(" is ")
		builder.WriteString(i.Key)
		builder.WriteString("\n")
	}
	return builder.String()
}

func IsValid[T any](s T) error {
	var errorList ErrorListValidation
	validate := validator.New()

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	for _, error := range validationErrors {
		errorList.Validation = append(errorList.Validation, ErrorListItem{
			Field: error.StructField(),
			Key:   error.Tag(),
		})
	}

	if len(errorList.Validation) > 0 {
		errorList.Entity = fmt.Sprintf("%T", s)
		return &errorList
	}

	return nil
}
