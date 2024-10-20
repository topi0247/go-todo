package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() (err error) {
	DB, err = sql.Open("postgres", "host=db user=postgres dbname=webapp sslmode=disable password=password")
	if err != nil {
		log.Println("Failed to open database:", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Println("Database ping failed:", err)
		return err
	}

	log.Println("Database connection established")
	return err
}
