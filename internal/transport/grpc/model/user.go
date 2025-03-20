package model

import (
	"errors"
	"regexp"
)

const (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	errName          = errors.New("name is required and must be less than 100 characters")
	errEmail         = errors.New("email is invalid")
	errEmptyPassword = errors.New("name is required and must be less than 100 characters")
)

type User struct {
	Name     string
	Email    string
	Password string
}

func (u *User) Validate() error {
	if u.Name == "" || len(u.Name) > 100 {
		return errName
	}

	regEx := regexp.MustCompile(emailRegex)
	if !regEx.MatchString(u.Email) || len(u.Email) < 5 || len(u.Email) > 100 {
		return errEmail
	}

	if len(u.Password) < 5 || len(u.Password) > 100 {
		return errEmptyPassword
	}

	return nil
}
