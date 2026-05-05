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

	log.Fatal(http.ListenAndServe(":5000", router))
}