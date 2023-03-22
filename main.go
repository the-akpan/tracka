package main

import (
	"log"
	"net/http"
	"time"

	"internal/database"
	"internal/middlewares"
	"internal/routes"

	"github.com/the-akpan/tracka/pkg/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// init app config
	config.Init()
	debug := config.GetDebug()

	// init db
	db := config.ConnectDB()
	admin := config.GetAdmin()
	database.Init(db, admin)

	// init middleware
	middleware := config.Middlewares()
	middlewares.Init(middleware)

	// init routes
	router := routes.Init(debug)
	srv := &http.Server{
		Handler:      router,
		Addr:         config.GetPort(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
