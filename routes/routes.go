package routes

import (
	"go-simple-crud-webservice/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/signup/", controllers.SignUp).Methods("GET")
	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.RemoveUser).Methods("DELETE")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
}
