package db

import (
	"context"
	"time"

	"github.com/pabloelisseo/twitt3r/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifySignUp(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twitt3r")
	col := db.Collection("users")

	signUp := make(map[string]interface{})
	if len(u.Name) > 0 {
		signUp["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		signUp["lastName"] = u.LastName
	}
	if !u.Birthdate.IsZero() {
		signUp["birthdate"] = u.Birthdate
	}
	if len(u.Email) > 0 {
		signUp["email"] = u.Email
	}
	if len(u.Avatar) > 0 {
		signUp["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		signUp["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		signUp["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		signUp["location"] = u.Location
	}
	if len(u.Website) > 0 {
		signUp["website"] = u.Website
	}

	updateString := bson.M{
		"$set": signUp,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	_, err := col.UpdateByID(ctx, objID, updateString)
	if err != nil {
		return false, err
	}
	return true, err
}
