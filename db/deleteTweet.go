package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twitt3r")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)
	userObjID, _ := primitive.ObjectIDFromHex(UserID)

	condition := bson.M{
		"_id":    objID,
		"userId": userObjID,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err
}
