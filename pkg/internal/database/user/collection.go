package user

import (
	"context"
	"internal/schemas"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	*mongo.Collection
}

func (coll *UserCollection) Get(email string) (*schemas.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var user schemas.User

	if err := coll.
		FindOne(ctx, bson.D{{Key: "email", Value: email}}).
		Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (coll *UserCollection) Create(email string, password string) (*schemas.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var user schemas.User

	if _, err := coll.Get(email); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	user = schemas.User{Email: email, CreatedAt: time.Now()}
	user.SetPassword(password)

	if _, err := coll.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (coll *UserCollection) Update(user schemas.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	filter := bson.D{{Key: "email", Value: user.Email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: user.Password}}}}
	_, err := coll.UpdateOne(ctx, filter, update)

	return err
}
