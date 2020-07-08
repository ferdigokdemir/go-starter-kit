package database

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func MongoInstance() *mongo.Database {
	return db
}

// Connect is starting to mongodb connection
func MongoConnect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("database.mongodb.uri")))
	if err != nil {
		panic(err)
	}
	db = client.Database(viper.GetString("database.mongodb.database"))
	return db
}

func MongoClose() (err error) {
	return db.Client().Disconnect(context.Background())
}
