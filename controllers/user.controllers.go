package controllers

import (
	"apiBlog/database"
	"apiBlog/model"
	"apiBlog/pass"
	"apiBlog/responses"
	"apiBlog/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	result, err := service.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&result)

}
func GetIdUser(w http.ResponseWriter, r *http.Request) {
	Param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(Param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := service.GetIdUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = user.ValidUser("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = service.Create(&user, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err = decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = user.ValidUser("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = service.Update(user, int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	service.Delete(int(id))
	w.WriteHeader(http.StatusAccepted)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var authDetails model.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err pass.Error
		err = pass.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser model.User
	database.DB.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
		var err pass.Error
		err = pass.SetError(err, "email or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := pass.CheckPassword(authDetails.Password, authUser.Password)

	if !check {
		var err pass.Error
		err = pass.SetError(err, "email or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := pass.Token(authUser.Email)
	if err != nil {
		var err pass.Error
		err = pass.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token pass.Token1
	token.Email = authUser.Email
	token.TokenString = validToken
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
