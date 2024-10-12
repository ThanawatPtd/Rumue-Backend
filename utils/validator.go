package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	valid "github.com/go-playground/validator/v10"
)

var validate = valid.New()

type ValidateError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ValidateStruct[T any](payload T) *ValidateError {
	err := validate.Struct(payload)
	errMsg := ""
	if err != nil {
		for _, err := range err.(valid.ValidationErrors) {
			tmp := strings.Split(err.StructNamespace(), ".")
			msg := fmt.Sprintf("%s is %s", tmp[len(tmp)-1], err.Tag())
			msg = strings.ToLower(string(msg[0])) + msg[1:]
			errMsg = errMsg + msg + ", "
		}

		return &ValidateError{
			Error:   "Invalid request",
			Message: errMsg[:len(errMsg)-2],
		}
	}

	return nil
}

type PasswordValidator interface {
	Match(string) error
}

type RegexPasswordValidator struct {
}

func (v *RegexPasswordValidator) Match(password string) error {
	// Ensure the password length is between 6 and 24 characters
    if len(password) < 6 || len(password) > 24 {
        return exceptions.ErrUnmatchPassword
    }

    // Check if it contains at least one uppercase letter
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    // Check if it contains at least one lowercase letter
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    // Check if it contains at least one digit
    hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

    // Ensure all conditions are met
    if !(hasUpper && hasLower && hasDigit) {
        return exceptions.ErrUnmatchPassword
    }

    return nil
}

func ValidatePassword(validator PasswordValidator, password string) error {
	return validator.Match(password)
}