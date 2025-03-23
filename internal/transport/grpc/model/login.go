package model

import "regexp"

type Login struct {
	Email    string
	Password string
}

func (l *Login) Validate() error {
	regEx := regexp.MustCompile(emailRegex)
	if !regEx.MatchString(l.Email) || len(l.Email) < minChars || len(l.Email) > maxChars {
		return errEmail
	}

	if len(l.Password) < minChars || len(l.Password) > maxChars {
		return errEmptyPassword
	}

	return nil
}
