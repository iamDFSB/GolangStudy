package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connection() (*sql.DB, error) {
	const Host = "localhost"
	const Password = 1234
	const User = "postgres"
	const DbName = "postgres"

	DataSourceName := fmt.Sprintf("postgres://%s:%d@%s:5432/%s", User, Password, Host, DbName)
	db, err := sql.Open("pgx", DataSourceName)
	
	if err != nil{
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}
	
	return db, nil
}