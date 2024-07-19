package util

import "errors"

var (
	ErrUsernameNotFound      = errors.New("Username not found")
	ErrUsernameAlreadyExists = errors.New("Username already exists")
)
