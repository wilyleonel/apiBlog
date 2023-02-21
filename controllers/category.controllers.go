package controllers

import (
	"apiBlog/model"
	"apiBlog/responses"
	"apiBlog/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetCategory()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&result)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var category model.Category
	err := decoder.Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err=category.ValidCategory("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = service.CreateCategory(&category, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&category)
	w.WriteHeader(http.StatusOK)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	service.DeleteCategory(int(id))
	w.WriteHeader(http.StatusAccepted)
}
