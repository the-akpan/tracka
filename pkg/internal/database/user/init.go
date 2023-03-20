package user

import (
	"context"
	"internal/schemas"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

func Init(db *mongo.Database) *UserCollection {
	collection := db.Collection(userCollection)
	user := &UserCollection{collection}

	var admin schemas.User
	credentials := map[string]string{}
	query := bson.D{{Key: "email", Value: credentials}}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := user.FindOne(ctx, query).Decode(&admin); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Creating admin...")
			user.Create("", "")
		} else {
			log.Fatal(err)
		}
	}

	return user
}
