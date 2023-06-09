package config

import (
	"internal/schemas"

	"github.com/spf13/viper"
)

func Middlewares() *schemas.Middleware {
	config := schemas.Middleware{}

	config.ContentTypes = viper.GetStringSlice("middleware.content_type")

	return &config
}
