package config

import (
	"log"
	"net/mail"
	"strings"

	"internal/schemas"

	"github.com/spf13/viper"
)

func GetAdmin() *schemas.User {
	log.Println("Loading root account...")
	credentials := getAdminDetails()

	user := &schemas.User{
		Email:    credentials["email"],
		Password: credentials["password"],
	}

	return user
}

func getAdminDetails() map[string]string {
	userMap := viper.GetStringMapString("admin")
	errs := make([]string, 0)

	keys := []string{"email", "password"}

	for _, key := range keys {
		if value, ok := userMap[key]; !ok {
			errs = append(errs, key+":: is missing")
		} else {
			userMap[key] = strings.TrimSpace(value)
			if strings.EqualFold(userMap[key], "") {
				errs = append(errs, key+":: cannot be blank")
			} else if key == keys[0] {
				if _, err := mail.ParseAddress(userMap[key]); err != nil {
					errs = append(errs, key+":: "+err.Error())
				}
			}
		}
	}

	if len(errs) != 0 {
		log.Fatal("Error getting admin details:\n" + strings.Join(errs, "\n"))
	}

	return userMap
}
