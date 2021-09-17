package repository

import (
	"context"

	"github.com/spf13/viper"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/model/audience"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SubscriptionsMongo defines subscriptions repository for MongoDB database.
type SubscriptionsMongo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewSubscriptionsMongo constructs new SubscriptionsMongo repository instance.
func NewSubscriptionsMongo(client *mongo.Client) *SubscriptionsMongo {
	return &SubscriptionsMongo{
		client:     client,
		collection: client.Database(viper.GetString("mongo_database")).Collection("subscriptions"),
	}
}

// GetByToken retrieves audience.SubscriptionTicket from the collection by given `token` of intention.EventConcern.
func (r *SubscriptionsMongo) GetByToken(token string) ([]audience.SubscriptionTicket, error) {
	var (
		results     []audience.SubscriptionTicket
		ctx, cancel = context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))
	)


	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{
		"concern_token": token,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, err
}

// Insert stores audience.SubscriptionTicket in the database.
func (r *SubscriptionsMongo) Insert(tickets ...audience.SubscriptionTicket) error {
	var (
		docs []interface{}
		ctx, cancel = context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))
	)

	defer cancel()

	for i := range tickets {
		docs = append(docs, tickets[i])
	}

	_, err := r.collection.InsertMany(ctx, docs)

	return err
}

// DeleteByToken removes audience.SubscriptionTicket from the database by given intention.EventConcern `token`.
func (r *SubscriptionsMongo) DeleteByToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))

	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.M{
		"concern_token": token,
	})

	return err
}

// DeleteByUserID removes audience.SubscriptionTicket from the database by given `userID`.
func (r *SubscriptionsMongo) DeleteByUserID(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))

	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.M{
		"user_id": userID,
	})

	return err
}
