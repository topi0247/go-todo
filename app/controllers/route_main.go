package controllers

import (
	"log"
	"net/http"
	"udemy-todo-app/app/helpers"
	"udemy-todo-app/app/models"
	"udemy-todo-app/infrastructure/db"
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
	userUUID := helpers.GetSession(r)
	if userUUID == "" {
		helpers.AppendFlash(w, r, helpers.FlashError, "ログインしてください")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	generateHTML(w, helpers.GetFlashes(w, r), "layout", "private_navbar", "flash", "index")
}

// func todoNew(w http.ResponseWriter, r *http.Request) {
// 	_, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
// 	}
// }

// func todoSave(w http.ResponseWriter, r *http.Request) {
// 	sess, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		err = r.ParseForm()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		user, err := sess.GetUserBySession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		content := r.PostFormValue("content")
// 		if err := user.CreateTodo(content); err != nil {
// 			log.Println(err)
// 		}

// 		http.Redirect(w, r, "/todos", http.StatusFound)
// 	}
// }

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
