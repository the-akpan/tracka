package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults...")
		} else {
			log.Fatal(err)
		}
	}
}

func setDefaults() {
	viper.SetDefault("port", "8000")
	viper.SetDefault("debug", true)
	viper.SetDefault("controller", map[string]string{
		"secret_key":  "BC236FF1AB9F84D39BF500B9758F43F6",
		"block_key":   "CF33022D399FC30106D60AEE76BF363B",
		"cookie_name": "tracka",
	})
}

func GetDebug() bool {
	return viper.GetBool("debug")
}
