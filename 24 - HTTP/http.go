package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Route"))
	})	
	

	log.Fatal(http.ListenAndServe(":5000", nil))
}