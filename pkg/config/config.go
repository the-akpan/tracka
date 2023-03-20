package config

import (
	"internal/database"
	"internal/middlewares"
	"internal/routes"
	"internal/schemas"
	"log"

	"github.com/spf13/viper"
)

var config schemas.Config

func Get() *schemas.Config {
	return &config
}

func Init() *schemas.Config {
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

	getAdmin()

	db := connectDB()
	database.Init(db)

	config.Port = getPort()
	debug := viper.GetBool("debug")
	log.Println(debug)
	routes.SetDebug(debug)
	config.Debug = debug

	middleware := configMiddlewares()
	middlewares.Init(middleware)

	return &config
}

func setDefaults() {
	viper.SetDefault("port", "8000")
	viper.SetDefault("debug", true)
	viper.SetDefault("database", map[string]string{
		"uri":  "mongodb://admin:password@localhost:27017",
		"name": "tracka",
	})
}
