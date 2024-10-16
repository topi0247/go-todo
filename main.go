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
	/*
	   u, _ := models.GetUser(1)
	   fmt.Println(u)
	   u.DeleteUser()
	   u, _ = models.GetUser(1)
	   fmt.Println(u)
	*/

	// Todo作成
	/*
	   user, _ := models.GetUser(2)
	   user.CreateTodo("First Todo")
	*/

	// Todo取得
	/*
	   t, _ := models.GetTodo(1)
	   fmt.Println(t)
	*/

	// todo全取得
	/*
	   user, _ := models.GetUser(2)
	   user.CreateTodo("Second Todo")

	   todos, _ := models.GetTodos()
	   for _, v := range todos {
	       fmt.Println(v)
	   }
	*/

	// 1ユーザーのTodo全取得
	/*
	   u := &models.User{}
	   u.Name = "test3"
	   u.Email = "test3@example.com"
	   u.Password = "testtest"
	   fmt.Println(u)
	   u.CreateUser()

	   user, _ := models.GetUser(3)
	   user.CreateTodo("Third Todo")

	   user2, _ := models.GetUser(2)
	   todos, _ := user2.GetTodosByUser()
	   for _, v := range todos {
	       fmt.Println(v)
	   }
	*/

	// todo更新
	/*
		t, _ := models.GetTodo(1)
		t.Content = "First Todo Update"
		t.UpdateTodo()
	*/

	// todo削除
	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}

// func outputConfig(cfg config.ConfigList) {
//     fmt.Printf("%+v\n", cfg)
// }
