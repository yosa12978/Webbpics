package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Picture struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Title   string             `json:"title" bson:"title"`
	Path    string             `json:"path" bson:"path"`
	PubDate int64              `json:"pub_date" bson:"pub_date"`
}
