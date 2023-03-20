package schemas

import (
	"fmt"
	"strings"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) Validate() map[string]string {
	errors := map[string]string{}

	login.Email = strings.TrimSpace(login.Email)
	if len(login.Email) != 0 {
		if email, err := validateEmail(login.Email); err != nil {
			errors["email"] = fmt.Sprintf(INVALID, "email")
		} else {
			login.Email = email
		}
	} else {
		errors["email"] = REQUIRED
	}

	if len(login.Password) == 0 {
		errors["password"] = REQUIRED
	}

	return errors
}
