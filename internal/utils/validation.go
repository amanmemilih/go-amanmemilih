package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError maps validator.ValidationErrors to Laravel-like format
func FormatValidationError(err error) map[string][]string {
	errors := map[string][]string{}

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := toSnakeCase(fe.Field())
			message := getErrorMessage(field, fe.Tag(), fe.Param())
			errors[field] = append(errors[field], message)
		}
	}

	return errors
}

// toSnakeCase converts CamelCase to snake_case
func toSnakeCase(str string) string {
	snake := regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

// getErrorMessage returns readable messages for each validation tag
func getErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s field is required", field)
	case "min":
		return fmt.Sprintf("%s field min is %s", field, param)
	case "max":
		return fmt.Sprintf("%s field max is %s", field, param)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	// tambahkan case lain sesuai kebutuhan
	default:
		return fmt.Sprintf("%s is not valid", field)
	}
}
