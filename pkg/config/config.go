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
}

func GetDebug() bool {
	return viper.GetBool("debug")
}
