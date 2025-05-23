package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	//def strings de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//abrindo a conexão com o postgresql
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	//verificação da conexão
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
