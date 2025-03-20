package model

import "regexp"

type Login struct {
	Email    string
	Password string
}

func (l *Login) Validate() error {
	regEx := regexp.MustCompile(emailRegex)
	if !regEx.MatchString(l.Email) || len(l.Email) < 5 || len(l.Email) > 100 {
		return errEmail
	}

	if len(l.Password) < 5 || len(l.Password) > 100 {
		return errEmptyPassword
	}

	return nil
}
