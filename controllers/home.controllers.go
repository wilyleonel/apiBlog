package controllers

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HomeResponse{
		Message: "Hello World!d----------",
	})
}