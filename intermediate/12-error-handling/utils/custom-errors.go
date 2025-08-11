package utils

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

type NotFoundError struct {
	Message string
}

type User struct {
	Email string
	Name  string
}

func NewValidationError(message, field string) *ValidationError {
	return &ValidationError{
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
		return ValidationError{
			Field:   "email",
			Message: "Format email tidak valid",
			Value:   email,
		}
	}
	return nil
}

func SaveData(id string, data any) error {
	if id == "" {
		return ValidationError{
			Field:   "id",
			Message: "id kosong tidak dapat menyimpan data",
			Value:   id,
		}
	}

	if data, ok := data.(User); ok {
		if err := validateEmail(data.Email); err != nil {
			return ValidationError{
				Field:   "email",
				Message: "Format email tidak sesuai",
				Value:   data.Email,
			}
		}
	}

	return nil
}

func CustomErrorLearn() {
	// errValidate := validateEmail("ask")
	// if errValidate != nil {
	// 	fmt.Println("isinya apa ya %w", errValidate)
	// }
	datUser := User{
		Email: "manboy",
		Name:  "bismen",
	}
	err := SaveData("32", datUser)

	if err != nil {
		// fmt.Println("Error when save data ", err)
		fmt.Println(err)
		return
	}
}
