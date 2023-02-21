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

func CreateComment(w http.ResponseWriter,r *http.Request){
	decoder:=json.NewDecoder(r.Body)
	var comment model.Comment
	err:=decoder.Decode(&comment)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err=comment.ValidComment("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err=service.CreateComment(&comment,err)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&comment)
	w.WriteHeader(http.StatusOK)
}

func UpdateComment(w http.ResponseWriter,r *http.Request){
	param:=mux.Vars(r)["id"]
	id,err:=strconv.ParseUint(param,10,64)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	decoder:=json.NewDecoder(r.Body)
	var comment model.Comment
	err=decoder.Decode(&comment)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err=comment.ValidComment("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err=service.UpdateComment(comment,int(id))
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&comment)
	w.WriteHeader(http.StatusOK)
}

func DeleteComment(w http.ResponseWriter,r *http.Request){
	param:=mux.Vars(r)["id"]
	id,err:=strconv.ParseUint(param,10,64)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	service.DeleteComment(int(id))
	w.WriteHeader(http.StatusAccepted)
}