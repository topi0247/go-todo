package main

import (
    "fmt"
    // "log"
    // "udemy-todo-app/config"
    "udemy-todo-app/app/models"
)

func main() {
    // cfg := config.Config
    // outputConfig(cfg)

    // log.Println("test")

    fmt.Println(models.Db)

    u := &models.User{}
    u.Name = "test"
    u.Email = "test@example.com"
    u.Password = "testtest"
    fmt.Println(u)

    u.CreateUser()
}

// func outputConfig(cfg config.ConfigList) {
//     fmt.Printf("%+v\n", cfg)
// }
