package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() (err error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbSettings := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)
	DB, err = sql.Open("postgres", dbSettings)
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
