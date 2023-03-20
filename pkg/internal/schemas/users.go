package schemas

import (
	"time"
)

type User struct {
	Email     string    `json:"email" bson:"email"`
	Password  []byte    `json:"-" bson:"password"`
	CreatedAt time.Time `json:"-" bson:"createdAt"`
}

func (user *User) SetPassword(password string) error {
	hash, err := hashValue([]byte(password))
	if err == nil {
		user.Password = hash
	}

	return err
}

func (user User) CheckPassword(password string) bool {
	return compareHashValue([]byte(password), []byte(user.Password))
}
