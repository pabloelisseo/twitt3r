package db

import (
	"context"
	"time"

	"github.com/pabloelisseo/twitt3r/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twitt3r")
	col := db.Collection("users")

	var result models.User
	query := bson.M{"email": email}

	err := col.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
