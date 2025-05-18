package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ParseValidationErrors convert validator.ValidationErrors jadi map[field]errorMessage
func ParseValidationErrors(err error) map[string]string {
	errorsMap := make(map[string]string)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range errs {
			field := fe.Field()
			tag := fe.Tag()

			var msg string
			switch tag {
			case "required":
				msg = fmt.Sprintf("%s is required", field)
			default:
				msg = fmt.Sprintf("%s failed on the %s validation", field, tag)
			}

			errorsMap[field] = msg
		}
	}
	return errorsMap
}
