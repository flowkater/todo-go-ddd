package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		return "validation failed"
	}
	return e.Errors[0].Message
}

// Validate validates any struct using tags
func Validate(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errors []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", err.Field())
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
			case "max":
				message = fmt.Sprintf("%s must not exceed %s characters", err.Field(), err.Param())
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", err.Field())
			case "url":
				message = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "oneof":
				message = fmt.Sprintf("%s must be one of [%s]", err.Field(), err.Param())
			default:
				message = fmt.Sprintf("%s is invalid", err.Field())
			}
			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Message: message,
			})
		}
		return &ValidationErrors{Errors: errors}
	}
	return nil
}
