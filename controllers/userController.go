package controllers

import (
	"encoding/json"
	"fmt"
	"go-simple-crud-webservice/models"
	"go-simple-crud-webservice/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	err := services.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	u, err := parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	err := services.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	u, err := parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	if ID != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	err := services.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	removeErr := models.RemoveUserById(ID)
	if removeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	err := services.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	u, err := models.GetUserByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	err := services.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetUsers())
	w.WriteHeader(http.StatusOK)
}

func parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
