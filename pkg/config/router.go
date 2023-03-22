package config

import (
	"log"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Extract app port from config file.
// Returns port as string
func GetPort() string {
	log.Println("Getting app port...")
	port := viper.GetString("port")
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("getPort: %+v\n", err)
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
