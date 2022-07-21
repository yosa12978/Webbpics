package services

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/yosa12978/webbpics/pkg/crypto"
	"github.com/yosa12978/webbpics/pkg/models"
	"github.com/yosa12978/webbpics/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IPictureService interface {
	AddPicture(pic io.Reader, filename string, title string) error
	GetPicture(id_hex string) (*models.Picture, error)
	GetPictures() []models.Picture
	DeletePicture(id_hex string) error
	PutPicture(id_hex string, title string) error
}

type PictureService struct {
	db *mongo.Database
}

func NewPictureService() IPictureService {
	return &PictureService{db: mongodb.GetDb()}
}

func (ps *PictureService) AddPicture(src io.Reader, filename string, title string) error {

	fext := strings.Split(filename, ".")
	path := os.Getenv("MEDIA_DIR") + "/" + crypto.NewToken(16) + "." + fext[len(fext)-1]
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	pic := models.Picture{
		Id:      primitive.NewObjectID(),
		Path:    path[1:],
		Title:   title,
		PubDate: time.Now().UnixNano(),
	}

	_, err = ps.db.Collection("pictures").InsertOne(context.Background(), pic)
	return err
}

func (ps *PictureService) GetPicture(id_hex string) (*models.Picture, error) {
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return nil, err
	}

	var pic models.Picture
	err = ps.db.Collection("pictures").FindOne(context.Background(), bson.M{"_id": id}).Decode(&pic)
	return &pic, err
}

func (ps *PictureService) GetPictures() []models.Picture {
	var pics []models.Picture
	cursor, _ := ps.db.Collection("pictures").Find(context.Background(), bson.M{}, options.Find().SetSort(bson.M{"_id": -1}))
	cursor.All(context.Background(), &pics)
	if pics == nil {
		return []models.Picture{}
	}
	return pics
}

func (ps *PictureService) DeletePicture(id_hex string) error {
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return err
	}

	pic, err := ps.GetPicture(id_hex)
	if err != nil {
		return err
	}

	os.Remove("." + pic.Path)

	_, err = ps.db.Collection("pictures").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (ps *PictureService) PutPicture(id_hex string, title string) error {
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return err
	}

	pic, err := ps.GetPicture(id_hex)
	if err != nil {
		return err
	}

	pic.Title = title
	_, err = ps.db.Collection("pictures").ReplaceOne(context.Background(), bson.M{"_id": id}, pic)
	return err
}
