package service

import "errors"

var (
	ErrGeneratePassword  = errors.New("failed to generate password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
)
