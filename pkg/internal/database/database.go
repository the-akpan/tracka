package database

import (
	"internal/database/orders"
	"internal/database/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type table struct {
	User  *user.UserCollection
	Order *orders.OrderCollection
}

var (
	DB    *mongo.Database
	Table *table
)

func Init(db *mongo.Database) {
	DB = db

	Table = &table{
		Order: orders.Init(DB),
		User:  user.Init(DB),
	}
}
