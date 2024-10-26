package exceptions

import "errors"

var (
	ErrUnrecognizedPassword = errors.New("unrecognized password") 
	ErrUnrecognizedEmail = errors.New("unrecognized email")
)