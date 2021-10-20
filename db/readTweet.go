package db

import (
	"context"
	"log"
	"time"

	"github.com/pabloelisseo/twitt3r/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweet(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twitt3r")
	col := db.Collection("tweets")

	var result []*models.ReturnTweets

	condition := bson.M{
		"userId": ID,
	}

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var document models.ReturnTweets
		err := cursor.Decode(&document)
		if err != nil {
			log.Fatal(err.Error())
			return result, false
		}
		result = append(result, &document)
	}
	return result, true
}
