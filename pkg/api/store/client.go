package store

import (
	"context"
	"time"

	"github.com/sid-sun/notes-api/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// ClientInterface defines the client type
type ClientInterface interface {
	GetCollection() (*mongo.Client, *mongo.Collection, error)
}

type client struct {
	config config.DBConfig
	logger *zap.Logger
	timeout time.Duration
}

// NewClient creates a new client instance
func NewClient(config config.DBConfig, lgr *zap.Logger) ClientInterface {
	return client{config: config, logger: lgr, timeout: time.Second*time.Duration(config.TimeoutInSec())}
}

// GetCollection connects to db, fetches collection from db and client
func (c client) GetCollection() (*mongo.Client, *mongo.Collection, error) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.config.Address()))
	if err != nil {
		c.logger.Sugar().Errorf("%s : %v", "[Client] [GetCollection] [Connect]", err)
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		c.logger.Sugar().Errorf("%s : %v", "[Client] [GetCollection] [Ping]", err)
		return nil, nil, err
	}

	return client, client.Database(c.config.Database()).Collection(c.config.Collection()), nil
}
