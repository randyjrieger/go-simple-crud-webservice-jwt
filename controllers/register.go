package controllers

import (
	"encoding/json"
	"go-simple-crud-webservice/models"
	"go-simple-crud-webservice/services"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds models.Creds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	td, err := services.CreateToken(creds.UserId, "customer")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate token"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token models.Token
	token.UserId = creds.UserId
	token.Role = "customer"
	token.TokenString = td.AccessToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
