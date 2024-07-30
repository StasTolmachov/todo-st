package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func ConnectionDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbNane := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbNane, dbPass)

	//connect to db
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	//check connection to db
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}
