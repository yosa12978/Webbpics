package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func New() error {
	addr := os.Getenv("MONGO_PATH")
	db_name := os.Getenv("MONGO_DATABASE")
	client_options := options.Client().ApplyURI(addr)
	client, err := mongo.Connect(context.Background(), client_options)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf(err.Error())
		return err
	}
	db = client.Database(db_name)
	return nil
}

func GetDb() *mongo.Database {
	if db == nil {
		New()
	}
	return db
}
