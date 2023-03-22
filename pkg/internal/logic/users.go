package logic

import (
	"errors"
	"internal/database"
	"internal/schemas"
)

// authenticates user details
func Login(details *schemas.Login) (*schemas.User, error) {
	user, err := database.Tables.User.Get(details.Email)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(details.Password) {
		return nil, errors.New("Unauthorized")
	}

	return user, nil
}
