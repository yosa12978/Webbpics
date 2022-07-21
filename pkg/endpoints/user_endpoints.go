package endpoints

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yosa12978/webbpics/pkg/dtos"
	"github.com/yosa12978/webbpics/pkg/helpers"
	"github.com/yosa12978/webbpics/pkg/services"
)

type IUserEndpoints interface {
	GetToken(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

type UserEndpoints struct {
}

func NewUserEndpoints() IUserEndpoints {
	return new(UserEndpoints)
}

func (ue *UserEndpoints) GetToken(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJson(w, 500, helpers.StatusMessage{Status_code: 500, Detail: err.Error()})
		return
	}
	var user_req dtos.UserReq
	err = json.Unmarshal(body, &user_req)
	if err != nil {
		helpers.WriteJson(w, 400, helpers.StatusMessage{Status_code: 400, Detail: err.Error()})
		return
	}
	userService := services.NewUserService()
	token, err := userService.LoginUser(user_req.Username, user_req.Password)
	if err != nil {
		helpers.WriteJson(w, 404, helpers.StatusMessage{Status_code: 404, Detail: err.Error()})
		return
	}
	helpers.WriteJson(w, 200, map[string]interface{}{"token": token})
}

func (ue *UserEndpoints) Signup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJson(w, 500, helpers.StatusMessage{Status_code: 500, Detail: err.Error()})
		return
	}
	var user_req dtos.UserReq
	err = json.Unmarshal(body, &user_req)
	if err != nil {
		helpers.WriteJson(w, 400, helpers.StatusMessage{Status_code: 400, Detail: err.Error()})
		return
	}
	userService := services.NewUserService()
	err = userService.Create(user_req.Username, user_req.Password)
	if err != nil {
		helpers.WriteJson(w, 400, helpers.StatusMessage{Status_code: 400, Detail: err.Error()})
		return
	}
	helpers.WriteJson(w, 201, helpers.StatusMessage{Status_code: 201, Detail: "Created"})
}
