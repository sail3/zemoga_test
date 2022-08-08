package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoClient *Client
	mongoOnce   sync.Once
)

type Client struct {
	*mongo.Client
}

func NewMongoClient(mongoURL string) *Client {
	mongoOnce.Do(func() {
		opt := options.Client()
		opt.ApplyURI(mongoURL)
		opt.SetMaxConnIdleTime(2 * time.Minute)
		opt.SetMaxConnecting(25)
		opt.SetMaxPoolSize(25)
		client, err := mongo.Connect(context.Background(), opt)
		if err != nil {
			panic(err)
		}

		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			panic(err)
		}
		mongoClient = &Client{client}
	})

	return mongoClient
}
