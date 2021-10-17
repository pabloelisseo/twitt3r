package models

import "time"

type RecordTweet struct {
	UserId    string    `bson:"userID" json:"userID,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
}
