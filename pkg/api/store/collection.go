package store

import (
	"context"
	"github.com/sid-sun/notes-api/cmd/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// CollectionInterface defines a db instance interface
type CollectionInterface interface {
	Find([]byte) (*mongo.SingleResult, error)
	Delete([]byte) (*mongo.DeleteResult, error)
	Insert([]byte) (*mongo.InsertOneResult, error)
}

// Collection implements CollectionInterface with map
type Collection struct {
	*mongo.Collection
	logger  *zap.Logger
	timeout time.Duration
}

// NewCollection creates a new instance for db
func NewCollection(cl *mongo.Collection, logger *zap.Logger, cfg config.DBConfig) CollectionInterface {
	return Collection{cl, logger, time.Second * time.Duration(cfg.TimeoutInSec())}
}

// Delete deletes the db Data instance corresponding to id
func (c Collection) Delete(condition []byte) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	res, err := c.Collection.DeleteOne(ctx, condition)
	if err != nil {
		c.logger.Sugar().Errorf("%s : %v", "[Collection] [Delete] [DeleteOne]", err)
		return nil, err
	}

	return res, err
}

// Get returns a db Data instance corresponding to id
func (c Collection) Find(condition []byte) (*mongo.SingleResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	res := c.Collection.FindOne(ctx, condition)
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		c.logger.Sugar().Errorf("%s : %v", "[Collection] [Find] [FindOne]", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

// Insert unconditionally sets db record of id to provided data
func (c Collection) Insert(document []byte) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	res, err := c.Collection.InsertOne(ctx, document)
	if err != nil {
		c.logger.Sugar().Errorf("%s : %v", "[Collection] [Insert] [InsertOne]", err)
		return nil, err
	}

	return res, nil
}
