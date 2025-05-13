package repository

import (
	"context"
	"time"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageRepository struct {
	client *mongo.Client
}

func (repository *PageRepository) connectToCollection() (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collection := repository.client.Database("content").Collection("pages")

	return collection, ctx, cancel
}

func (repository *PageRepository) Insert(page model.Page) error {
	collection, ctx, cancel := repository.connectToCollection()
	defer cancel()

	_, err := collection.InsertOne(ctx, page)
	return err
}
