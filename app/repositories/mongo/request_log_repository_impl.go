package mongo

import (
	"context"
	"log/slog"

	"github.com/anurag925/aero/app/models"
	"github.com/anurag925/aero/core/logger"
	"github.com/anurag925/mongoboiler"
	"go.mongodb.org/mongo-driver/mongo"
)

type requestLogRepositoryImpl struct {
	collection *mongoboiler.Collection
}

func NewRequestLogRepositoryImpl(client *mongo.Client, dbName string) *requestLogRepositoryImpl {
	return &requestLogRepositoryImpl{
		collection: mongoboiler.New(client, dbName).NewCollection("request_logs"),
	}
}

func (r *requestLogRepositoryImpl) Log(ctx context.Context, log *models.RequestLog) error {
	d, err := r.collection.InsertOne(ctx, log)
	logger.Info(ctx, "request log inserted", slog.Any("log", d))
	if err != nil {
		logger.Error(ctx, "unable to insert request log", err)
		return err
	}
	return nil
}
