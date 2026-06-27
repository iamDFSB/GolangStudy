package main

import (
	"crud/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", routers.CreateUser).Methods(http.MethodPost) // Post
	router.HandleFunc("/users", routers.GetUsers).Methods(http.MethodGet) // Get
	router.HandleFunc("/users/{id}", routers.GetUserById).Methods(http.MethodGet) // Get

	log.Fatal(http.ListenAndServe(":5000", router))
}