package endpoints

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosa12978/webbpics/pkg/dtos"
	"github.com/yosa12978/webbpics/pkg/helpers"
	"github.com/yosa12978/webbpics/pkg/services"
)

type IPictureEndpoints interface {
	GetPicture(w http.ResponseWriter, r *http.Request)
	GetPictures(w http.ResponseWriter, r *http.Request)
	AddPicture(w http.ResponseWriter, r *http.Request)
	DeletePicture(w http.ResponseWriter, r *http.Request)
	UpdatePicture(w http.ResponseWriter, r *http.Request)
}

type PictureEndpoints struct {
}

func NewPictureEndpoints() IPictureEndpoints {
	return new(PictureEndpoints)
}

func (pe *PictureEndpoints) GetPicture(w http.ResponseWriter, r *http.Request) {
	pictureService := services.NewPictureService()
	vars := mux.Vars(r)
	pic, err := pictureService.GetPicture(vars["id"])
	if err != nil {
		helpers.WriteJson(w, 404, helpers.StatusMessage{Status_code: 404, Detail: err.Error()})
		return
	}
	helpers.WriteJson(w, 200, *pic)
}

func (pe *PictureEndpoints) GetPictures(w http.ResponseWriter, r *http.Request) {
	pictureService := services.NewPictureService()
	pics := pictureService.GetPictures()
	helpers.WriteJson(w, 200, pics)
}

func (pe *PictureEndpoints) AddPicture(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		helpers.WriteJson(w, 400, helpers.StatusMessage{Status_code: 400, Detail: err.Error()})
		return
	}
	title := r.FormValue("title")
	pictureService := services.NewPictureService()
	err = pictureService.AddPicture(file, handler.Filename, title)
	if err != nil {
		helpers.WriteJson(w, 500, helpers.StatusMessage{Status_code: 500, Detail: err.Error()})
		return
	}
	helpers.WriteJson(w, 201, helpers.StatusMessage{Status_code: 201, Detail: "Created"})
}

func (pe *PictureEndpoints) DeletePicture(w http.ResponseWriter, r *http.Request) {
	pictureService := services.NewPictureService()
	vars := mux.Vars(r)
	err := pictureService.DeletePicture(vars["id"])
	if err != nil {
		helpers.WriteJson(w, 404, helpers.StatusMessage{Status_code: 404, Detail: "Not found"})
		return
	}
	helpers.WriteJson(w, 200, helpers.StatusMessage{Status_code: 200, Detail: "Success"})
}

func (pe *PictureEndpoints) UpdatePicture(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJson(w, 500, helpers.StatusMessage{Status_code: 500, Detail: err.Error()})
		return
	}
	var dto dtos.PicUpdate
	json.Unmarshal(body, &dto)

	pictureService := services.NewPictureService()
	err = pictureService.PutPicture(id, dto.Title)
	if err != nil {
		helpers.WriteJson(w, 404, helpers.StatusMessage{Status_code: 404, Detail: err.Error()})
		return
	}

	helpers.WriteJson(w, 200, helpers.StatusMessage{Status_code: 200, Detail: "Success"})
}
