package main

import (
	"fmt"
	"log"
	"todo-app/app/controllers"
	"todo-app/infrastructure/db"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Failed to connect to the database: %s\n", err)
	}

	boil.SetDB(db.DB)
	controllers.StartMainServer()
}
