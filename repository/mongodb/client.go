package mongodb

import (
	"context"
	"time"

	"github.com/marcelo-r/shortener/shortener"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI        string        `env:"URI"`
	Database   string        `env:"DATABASE"`
	Collection string        `env:"COLLECTION"`
	Timeout    time.Duration `env:"TIMEOUT"`
}

type client struct {
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(ctx context.Context, config Config) (shortener.Repository, error) {
	withTimeout, cancel := context.WithTimeout(ctx, config.Timeout)
	defer cancel()
	mongoClient, err := mongo.Connect(withTimeout, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}
	collection := mongoClient.Database(config.Database).Collection(config.Collection)
	return &client{
		ctx:        ctx,
		client:     mongoClient,
		collection: collection,
		timeout:    config.Timeout,
	}, nil
}

func (c *client) FindByID(ctx context.Context, id string) (shortener.Redirect, error) {
	redirect := shortener.Redirect{}
	return redirect, nil
}

func (c *client) Store(ctx context.Context, redirect shortener.Redirect) error {
	return nil
}
