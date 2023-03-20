package schemas

import (
	"net/mail"
	"strings"
)

func validateEmail(email string) (string, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return "", err
	}

	return strings.ToLower(email), nil
}
