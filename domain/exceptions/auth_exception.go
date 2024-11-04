package exceptions

import "errors"

var (
	ErrLoginFailed = errors.New("login failed")
)