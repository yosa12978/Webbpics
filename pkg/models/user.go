package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"-" bson:"password"`
	Token    string             `json:"token" bson:"token"`
	Is_admin bool               `json:"is_admin" bson:"is_admin"`
}

func (u *User) IsValid() bool {
	return true
}
