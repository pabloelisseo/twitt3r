package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&authSource=admin&appname=MongoDB%20Compass&ssl=false")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successfully connected to DB")
	return client
}

func CheckConnection() error {
	err := MongoClient.Ping(context.TODO(), nil)
	return err
}
