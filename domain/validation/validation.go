package validation

import (
	"fmt"
	"unicode"
)

type ValidationError struct {
	Message string
}

var (
	ErrPasswordTooLong  = &ValidationError{"password cannot be longer than 18 characters"}
	ErrPasswordNoNumber = &ValidationError{"password must contain at least one number"}
	ErrPasswordTooShort = &ValidationError{"password must be at least 8 characters long"}
)

func (e *ValidationError) Error() string {
	return e.Message
}

func containsNumber(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			fmt.Printf("Found digit: %c\n", r)
			return true
		}
	}
	return false
}

func ValidatePassword(password string) error {
	if !containsNumber(password) {
		return ErrPasswordNoNumber
	}

	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	if len(password) > 18 {
		return ErrPasswordTooLong
	}
	return nil
}
