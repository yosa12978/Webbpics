package helpers

import (
	"encoding/json"
	"net/http"
)

type StatusMessage struct {
	Status_code int    `json:"status_code"`
	Detail      string `json:"detail"`
}

func WriteJson(w http.ResponseWriter, status_code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status_code)
	json.NewEncoder(w).Encode(data)
}
