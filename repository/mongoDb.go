package repository

import (
	"context"
	"go-ms-session/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Ctx               context.Context
	ProfileCollection *mongo.Collection
	MongoClient       *mongo.Client
}


func NewMongoRepository(c *mongo.Client, cfg *config.ServerConfig) *MongoRepository {
	return &MongoRepository{
		Ctx:               context.Background(),
		ProfileCollection: c.Database(cfg.MongoDb.Database).Collection("access-session"),
		MongoClient:       c,
	}
}


