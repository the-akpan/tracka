package config

import (
	"log"
	"strings"

	"internal/schemas"

	"github.com/spf13/viper"
)

func getAdmin() *schemas.User {
	log.Println("Loading root account...")
	getAdminDetails()

	return &schemas.User{}
}

func getAdminDetails() map[string]string {
	userMap := viper.GetStringMapString("admin")
	errs := make([]string, 0)

	for key, value := range userMap {

		userMap[key] = strings.TrimSpace(value)
		if strings.EqualFold(userMap[key], "") {
			errs = append(errs, key+" cannot be blank")
		}
	}

	if len(errs) != 0 {
		log.Fatal(strings.Join(errs, "\n"))
	}

	return userMap
}
