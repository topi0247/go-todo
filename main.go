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

    // ユーザー作成
    /*
    u := &models.User{}
    u.Name = "test"
    u.Email = "test@example.com"
    u.Password = "testtest"
    fmt.Println(u)

    u.CreateUser()
    */

    // ユーザー情報取得
    /*
    u, _ := models.GetUser(1)
    fmt.Println(u)
    */

    // ユーザー情報更新
    /*
    u, _ := models.GetUser(1)
    fmt.Println(u)
    u.Name = "test2"
    u.Email = "test2@example.com"
    u.UpdateUser()
    u, _ = models.GetUser(1)
    fmt.Println(u)
    */

    // ユーザー削除
    u, _ := models.GetUser(1)
    fmt.Println(u)
    u.DeleteUser()
    u, _ = models.GetUser(1)
    fmt.Println(u)
}

// func outputConfig(cfg config.ConfigList) {
//     fmt.Printf("%+v\n", cfg)
// }
