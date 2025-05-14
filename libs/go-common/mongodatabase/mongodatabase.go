package mongodatabase

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	host = os.Getenv("MONGODB_HOST")
	port = os.Getenv("MONGODB_PORT")
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))
	if err != nil {
		return nil, err
	}
	return client, nil
}
