package main

import (
	"go-simple-crud-webservice/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("ACCESS_SECRET", "e2922901-f374-4283-a1ad-0e3c6d06011f")

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:3000", r)
}
