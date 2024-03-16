package mongo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type mongoConfig struct{
	port     string
	host     string
	username string
	password string
}


func NewMongoClient() (*mongo.Client, error) {
	config := mongoConfig{
		port:     viper.GetString("mongo.port"),
		host:     viper.GetString("mongo.host"),
		username: viper.GetString("mongo.username"),
		password: viper.GetString("mongo.password"),
	}

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.username, config.password, config.host, config.port)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
