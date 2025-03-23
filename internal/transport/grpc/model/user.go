package model

import (
	"errors"
	"regexp"
)

const (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	minChars   = 5
	maxChars   = 100
)

var (
	errName          = errors.New("name is required and must be less than 100 characters")
	errEmail         = errors.New("email is invalid")
	errEmptyPassword = errors.New("password is required and must be less than 100 characters")
)

type User struct {
	Name     string
	Email    string
	Password string
}

func (u *User) Validate() error {
	if u.Name == "" || len(u.Name) > maxChars {
		return errName
	}

	regEx := regexp.MustCompile(emailRegex)
	if !regEx.MatchString(u.Email) || len(u.Email) < minChars || len(u.Email) > maxChars {
		return errEmail
	}

	if len(u.Password) < minChars || len(u.Password) > maxChars {
		return errEmptyPassword
	}

	return nil
}
