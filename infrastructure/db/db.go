package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbSettings := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)
	log.Println("Connecting to database with settings:", dbSettings)
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
