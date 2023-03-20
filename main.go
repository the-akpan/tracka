package main

import (
	"log"
	"net/http"
	"time"

	"internal/routes"

	"github.com/the-akpan/tracka/pkg/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize app config
	conf := config.Init()

	router := routes.Init()
	srv := &http.Server{
		Handler:      router,
		Addr:         conf.Port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
