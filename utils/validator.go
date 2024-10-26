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
	Recognize(string) error
}

type RegexPasswordValidator struct {
}

func (v *RegexPasswordValidator) Recognize(password string) error {
	// Ensure the password length is between 6 and 24 characters
    if len(password) < 6 || len(password) > 24 {
        return exceptions.ErrUnrecognizedPassword
    }

    // Check if it contains at least one uppercase letter
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    // Check if it contains at least one lowercase letter
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    // Check if it contains at least one digit
    hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

    // Ensure all conditions are met
    if !(hasUpper && hasLower && hasDigit) {
        return exceptions.ErrUnrecognizedPassword
    }

    return nil
}

func ValidatePassword(validator PasswordValidator, password string) error {
	return validator.Recognize(password)
}

type EmailValidator interface {
	Recognize(string) error
}

type RegexEmailValidator struct {
}

func (ev *RegexEmailValidator) Recognize(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return exceptions.ErrUnrecognizedEmail
	}
	return nil

}

func ValidateEmail(validator EmailValidator, email string) error {
	return validator.Recognize(email)
}