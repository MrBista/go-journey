package utils

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

type NotFoundError struct {
	Message string
}

func NewValidationError(message, field string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal untuk field %s dengan message %s",
		e.Field, e.Message)
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return NewValidationError("Format email tidak sesuai", "email")
	}
	return nil
}

func CustomErrorLearn() {
	errValidate := validateEmail("ask")
	if errValidate != nil {
		fmt.Println("isinya apa ya %w", errValidate)
	}
}
