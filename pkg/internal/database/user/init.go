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

func Init(db *mongo.Database, admin *schemas.User) *UserCollection {
	collection := db.Collection(userCollection)
	userColl := &UserCollection{collection}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if num, err := userColl.CountDocuments(ctx, bson.D{}); num == 0 {
		log.Println("Creating admin...")
		_, err = userColl.Create(admin.Email, string(admin.Password))
		if err != nil {
			log.Fatalln(err)
		}
	} else if err != nil {
		log.Fatalln(err)
	}

	return userColl
}
