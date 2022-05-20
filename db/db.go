package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	// conf := config.GetConfig()
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func CreateCon() *sql.DB {
	return db
}
