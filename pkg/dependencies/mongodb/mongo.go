package mongoConn

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigMongo struct {
	Host     string `envconfig:"MONGO_HOST" validate:"required"`
	Port     string `envconfig:"MONGO_PORT" validate:"required"`
	User     string `envconfig:"MONGO_USER" validate:"required"`
	Password string `envconfig:"MONGO_PASSWORD" validate:"required"`
	Database string `envconfig:"MONGO_DATABASE" validate:"required"`
}

func NewMongoDB(cfg *ConfigMongo) (*mongo.Client, error) {
	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}
	connUrl := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(connUrl).SetAuth(credential)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
