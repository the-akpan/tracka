package config

import (
	"encoding/json"
	"internal/schemas"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func GetControllerConfig() *schemas.Controller {
	log.Println("Loading controller config...")
	credentials := getControllerData()

	config := &schemas.Controller{}
	data, _ := json.Marshal(credentials)
	if err := json.Unmarshal(data, config); err != nil {
		log.Fatalln(err)
	}

	return config
}

func getControllerData() map[string]string {
	cMap := viper.GetStringMapString("controller")
	errs := make([]string, 0)

	keys := []string{"block_key", "secret_key", "cookie_name"}

	for _, key := range keys {
		if value, ok := cMap[key]; !ok {
			errs = append(errs, key+":: is missing")
		} else {
			cMap[key] = strings.TrimSpace(value)
			if strings.EqualFold(cMap[key], "") {
				errs = append(errs, key+":: cannot be blank")
			}
		}
	}

	if len(errs) != 0 {
		log.Fatal("Error getting admin details:\n" + strings.Join(errs, "\n"))
	}

	return cMap
}
