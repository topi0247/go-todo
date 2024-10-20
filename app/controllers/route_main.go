package controllers

import (
	"log"
	"net/http"
	"udemy-todo-app/app/helpers"
	"udemy-todo-app/app/models"
	"udemy-todo-app/infrastructure/db"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func top(w http.ResponseWriter, r *http.Request) {
	userUUID := helpers.GetSession(r)
	log.Printf("userUUID: %s", userUUID)
	if userUUID != "" {
		_, err := models.Users(
			models.UserWhere.UUID.EQ(userUUID),
		).One(r.Context(), db.DB)
		if err == nil {
			http.Redirect(w, r, "/todos", http.StatusFound)
			return
		}
	}
	generateHTML(w, nil, "layout", "public_navbar", "top")
}

func index(w http.ResponseWriter, r *http.Request) {
	if !helpers.Authenticate(w, r) {
		return
	}

	user := helpers.CurrentUser(r)

	todos, err := models.Todos(
		models.TodoWhere.UserID.EQ(user.ID),
	).All(r.Context(), db.DB)
	if err != nil {
		helpers.AppendFlash(w, r, helpers.FlashError, "予期せぬエラーが発生しました。再ログインしてください")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	flash := helpers.GetFlashes(w, r)

	data := map[string]interface{}{
		"User":         user,
		"Todos":        todos,
		"FlashSuccess": flash.FlashSuccess,
		"FlashError":   flash.FlashError,
		"FlashNotice":  flash.FlashNotice,
	}
	generateHTML(w, data, "layout", "private_navbar", "flash", "index")
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	if !helpers.Authenticate(w, r) {
		return
	}

	generateHTML(w, helpers.GetFlashes(w, r), "layout", "private_navbar", "todo_new")
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	if !helpers.Authenticate(w, r) {
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		helpers.AppendFlash(w, r, helpers.FlashError, "入力値が不正です")
		http.Redirect(w, r, "/todos/new", http.StatusFound)
		return
	}

	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	user := helpers.CurrentUser(r)
	todo := &models.Todo{
		Title:       title,
		Description: description,
		UserID:      user.ID,
	}

	if err := todo.Insert(r.Context(), db.DB, boil.Infer()); err != nil {
		log.Println(err)
		helpers.AppendFlash(w, r, helpers.FlashError, "タスクの作成に失敗しました")
		http.Redirect(w, r, "/todos/new", http.StatusFound)
		return
	}

	helpers.AppendFlash(w, r, helpers.FlashSuccess, "タスクを作成しました")
	http.Redirect(w, r, "/todos", http.StatusFound)
}

// func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
// 	sess, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		_, err := sess.GetUserBySession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		t, err := models.GetTodo(id)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
// 	}
// }

// func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
// 	sess, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		user, err := sess.GetUserBySession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		content := r.PostFormValue("content")
// 		t := &models.Todo{
// 			ID:      id,
// 			Content: content,
// 			UserID:  user.ID,
// 		}
// 		if err := t.UpdateTodo(); err != nil {
// 			log.Println(err)
// 		}
// 		http.Redirect(w, r, "/todos", http.StatusFound)
// 	}
// }

// func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
// 	sess, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		_, err := sess.GetUserBySession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		t, err := models.GetTodo(id)
// 		if err != nil {
// 			log.Fatalln()
// 		}
// 		if err := t.DeleteTodo(); err != nil {
// 			log.Println(err)
// 		}
// 		http.Redirect(w, r, "/todos", http.StatusFound)
// 	}
// }
