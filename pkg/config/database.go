package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectDB extracts db credentials from config
// and attempts to create a db connection
func connectDB() *mongo.Database {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Println("Connecting to DB...")
	db := viper.GetStringMapString("database")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db["uri"]))
	if err != nil {
		log.Fatalf("connectDB: %+v", err)
	}

	return client.Database(db["name"])
}
