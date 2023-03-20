package orders

import (
	"context"
	"internal/schemas"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderCollection struct {
	*mongo.Collection
}

func (coll *OrderCollection) Get(page int64, limit int64) ([]*schemas.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var orders []*schemas.Order = make([]*schemas.Order, 0)
	skip := (page - 1) * limit
	opt := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "createdAt", Value: -1}}).
		SetProjection(bson.D{{Key: "updates", Value: 0}})

	cursor, err := coll.Find(ctx, bson.D{{}}, opt)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (coll *OrderCollection) GetOne(tracking_num string) (*schemas.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var order schemas.Order

	if err := coll.
		FindOne(ctx, bson.D{{Key: "tracking_num", Value: tracking_num}}).
		Decode(&order); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatalf("order get: %+v", err)
	}

	return &order, nil
}

func (coll *OrderCollection) Create(data *schemas.Order) (*schemas.Order, error) {
	return nil, nil
}
