package main

import (
	"go-simple-crud-webservice/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:3000", r)
}
