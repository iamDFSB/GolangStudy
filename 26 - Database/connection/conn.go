package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type UserSchema struct{
	ID uint `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	const Host = "localhost"
	const Password = 1234
	const User = "postgres"
	const DbName = "postgres"

	// postgres://username:password@localhost:5432/database_name
	DataSourceName := fmt.Sprintf("postgres://%s:%d@%s:5432/%s", User, Password, Host, DbName)
	db, err := sql.Open("pgx", DataSourceName)
	
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()
	
	rows, err := db.Query("select * from public.usuario")

	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var result UserSchema
		err := rows.Scan(
			&result.ID,
			&result.Nome,
			&result.Email,
		)

		if err != nil{
			log.Fatal(err)
		}

		fmt.Println(result)
	}

}