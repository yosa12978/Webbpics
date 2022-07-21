package services

import (
	"context"
	"errors"

	"github.com/yosa12978/webbpics/pkg/crypto"
	"github.com/yosa12978/webbpics/pkg/models"
	"github.com/yosa12978/webbpics/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	Create(username string, password string) error
	GetUserByToken(token string) (*models.User, error)
	LoginUser(username string, password string) (string, error)
	IsUsernameTaken(username string) bool
}

type UserService struct {
	db *mongo.Database
}

func NewUserService() IUserService {
	return &UserService{db: mongodb.GetDb()}
}

func (us *UserService) Create(username string, password string) error {
	if us.IsUsernameTaken(username) {
		return errors.New("this username is already in use")
	}

	user := models.User{
		Id:       primitive.NewObjectID(),
		Username: username,
		Password: crypto.NewMD5(password),
		Token:    crypto.NewToken(32),
		Is_admin: false,
	}

	_, err := us.db.Collection("users").InsertOne(context.Background(), user)
	return err
}

func (us *UserService) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	err := us.db.Collection("users").FindOne(context.Background(), bson.M{"token": token}).Decode(&user)
	return &user, err
}

func (us *UserService) LoginUser(username string, password string) (string, error) {
	var user models.User

	filter := bson.M{
		"$and": []bson.M{
			{"username": username},
			{"password": crypto.NewMD5(password)},
		},
	}

	err := us.db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	return user.Token, err
}

func (us *UserService) IsUsernameTaken(username string) bool {
	var user *models.User
	err := us.db.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	return err == nil
}
