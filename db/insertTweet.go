package db

import (
	"context"
	"time"

	"github.com/pabloelisseo/twitt3r/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twitt3r")
	col := db.Collection("tweets")

	tweet := bson.M{
		"userId":    t.UserId,
		"message":   t.Message,
		"createdAt": t.CreatedAt,
	}

	result, err := col.InsertOne(ctx, tweet)
	if err != nil {
		return string(""), false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil
}
