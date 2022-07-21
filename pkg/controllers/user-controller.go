package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leonardodfelix/mysql-crud/pkg/models"
	"github.com/leonardodfelix/mysql-crud/pkg/utils"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]
	userId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	userDetails, _ := models.GetUserById(userId)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]
	userId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	userDetails := models.DeleteUser(userId)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var UpdateUser = &models.User{}
	utils.ParseBody(r, UpdateUser)
	params := mux.Vars(r)
	id := params["userId"]
	userId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	userDetails, db := models.GetUserById(userId)
	if UpdateUser.Name != "" {
		userDetails.Name = UpdateUser.Name
	}
	if UpdateUser.Email != "" {
		userDetails.Email = UpdateUser.Email
	}
	if UpdateUser.Password != "" {
		userDetails.Password = UpdateUser.Password
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
