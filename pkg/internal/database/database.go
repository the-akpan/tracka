package database

import (
	"internal/database/orders"
	"internal/database/user"
	"internal/schemas"

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

func Init(db *mongo.Database, admin *schemas.User) {
	DB = db

	Table = &table{
		Order: orders.Init(DB),
		User:  user.Init(DB, admin),
	}
}
