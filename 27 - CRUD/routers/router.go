package routers

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserSchema struct{
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while reading the request body"))
		return 
	}

	var user UserSchema
	if err = json.Unmarshal([]byte(requestBody), &user); err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while casting requestBody to UserSchema"))
		return 
	}
	
	db, err := database.Connection()
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while connecting on database"))
		return 
	}
	defer db.Close()


	var idInserted int
	err = db.QueryRow(
		"INSERT INTO public.usuario (nome, email) VALUES ($1, $2) RETURNING id",
		user.Nome, 
		user.Email,
	).Scan(&idInserted)

	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while while executing statement"))
		return 
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User added successfully, ID: %d", idInserted)))
}