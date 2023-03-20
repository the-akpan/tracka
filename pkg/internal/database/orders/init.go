package orders

import "go.mongodb.org/mongo-driver/mongo"

const orderCollection = "orders"

func Init(db *mongo.Database) *OrderCollection {
	collection := db.Collection(orderCollection)
	order := OrderCollection{collection}

	return &order
}
