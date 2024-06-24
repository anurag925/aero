package initializers

import (
	"context"

	"github.com/anurag925/aero/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// default mongo db provider
func InitMongoDB() (*mongo.Client, error) {
	ctx := context.TODO()
	url := config.Secrets().MongoDBUrl
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
