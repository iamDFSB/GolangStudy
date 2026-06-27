package routers

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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


func GetUsers(w http.ResponseWriter, r *http.Request){
	db, err := database.Connection()
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while connecting on database"))
		return 
	}
	defer db.Close()

	rows, err := db.Query("SELECT nome, email FROM public.usuario")
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while executing query"))
		return 
	}
	defer rows.Close()
	
	var users []UserSchema = []UserSchema{}
	for rows.Next(){
		var user UserSchema
		if err := rows.Scan(&user.Nome, &user.Email); err != nil{
			 w.Write([]byte("An error occured while compiling the rows"))
			 return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil{
		w.Write([]byte("An error occured while encoding users list"))
		return 
	}

}


func GetUserById(w http.ResponseWriter, r *http.Request){
	
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil{
		w.Write([]byte("An error occured while parsing id param"))
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

	rows, err := db.Query(`
		SELECT nome, email FROM public.usuario
		WHERE id = $1
	`, id)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An error occurred while executing query"))
		return 
	}
	defer rows.Close()
	
	rows.Next()

	var user UserSchema
	if err := rows.Scan(&user.Nome, &user.Email); err != nil{
			w.Write([]byte("An error occured while compiling the rows"))
			return
	}
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil{
		w.Write([]byte("An error occured while encoding users list"))
		return 
	}
}