package mongo

import (
	"context"
	"time"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageRepository struct {
	Client *mongo.Client
}

func (repository *PageRepository) connect() (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collection := repository.Client.Database("content").Collection("pages")

	return collection, ctx, cancel
}

func (repository *PageRepository) Insert(p *page.Page) error {
	collection, ctx, cancel := repository.connect()
	defer cancel()

	_, err := collection.InsertOne(ctx, p)
	return err
}

func (repository *PageRepository) Update(p *page.Page) error {
	collection, ctx, cancel := repository.connect()
	defer cancel()

	_, err := collection.UpdateByID(ctx, p.ID, p)
	return err
}
