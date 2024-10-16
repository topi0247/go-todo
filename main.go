package main

import (
	"fmt"
	"udemy-todo-app/app/controllers"
	"udemy-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
