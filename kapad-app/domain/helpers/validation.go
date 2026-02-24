package helpers

import (
	"errors"
	"unicode"
)

func ValidateUsername(username string) error {
	username = trimSpaces(username)
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	return nil
}

func ValidatePassword(password string) error {
	password = trimSpaces(password)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasLetter := false
	hasNumber := false
	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		}
		if unicode.IsDigit(ch) {
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return errors.New("password must contain at least one letter and one number")
	}

	return nil
}

func trimSpaces(s string) string {
	return string([]rune(s))
}
