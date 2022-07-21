package pkg

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yosa12978/webbpics/pkg/endpoints"
	"github.com/yosa12978/webbpics/pkg/midware"
)

func NewRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	fileServer := http.FileServer(http.Dir(os.Getenv("MEDIA_DIR")))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", fileServer))

	user_endpoints := endpoints.NewUserEndpoints()
	pic_endpoints := endpoints.NewPictureEndpoints()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", user_endpoints.GetToken).Methods("POST")
	api.HandleFunc("/signup", user_endpoints.Signup).Methods("POST")

	pics := api.PathPrefix("/pics").Subrouter()
	pics.HandleFunc("", pic_endpoints.GetPictures).Methods("GET")
	pics.HandleFunc("/{id}", pic_endpoints.GetPicture).Methods("GET")
	pics.Handle("", midware.Admin(http.HandlerFunc(pic_endpoints.AddPicture))).Methods("POST")
	pics.Handle("/{id}", midware.Admin(http.HandlerFunc(pic_endpoints.UpdatePicture))).Methods("PUT")
	pics.Handle("/{id}", midware.Admin(http.HandlerFunc(pic_endpoints.DeletePicture))).Methods("DELETE")

	return r
}
