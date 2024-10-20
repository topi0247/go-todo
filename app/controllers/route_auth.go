package controllers

import (
	"fmt"
	"log"
	"net/http"
	"udemy-todo-app/app/helpers"
	"udemy-todo-app/app/models"
	"udemy-todo-app/infrastructure/db"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func encrypt(password string) (passwordDigest string) {
	passwordDigestBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	passwordDigest = string(passwordDigestBytes)
	return passwordDigest
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, helpers.GetFlashes(w, r), "layout", "public_navbar", "flash", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			helpers.AppendFlash(w, r, helpers.FlashError, "入力値が不正です")
		}

		user := &models.User{
			UUID:           uuid.New().String(),
			Name:           r.PostFormValue("name"),
			Email:          r.PostFormValue("email"),
			PasswordDigest: encrypt(r.PostFormValue("password")),
		}

		if err := user.Insert(r.Context(), db.DB, boil.Infer()); err != nil {
			log.Println("Error occurred while inserting user:", err)
			fmt.Printf("Failed to create user: %s\n", err)
			helpers.AppendFlash(w, r, helpers.FlashError, "ユーザーの作成に失敗しました")
			http.Redirect(w, r, "/signup", http.StatusFound)
		} else {
			helpers.AppendFlash(w, r, helpers.FlashSuccess, "新規登録しました")
			helpers.CreateSession(w, r, user.UUID)
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// if err != nil {
	// 	generateHTML(w, nil, "layout", "public_navbar", "login")
	// } else {
	// 	http.Redirect(w, r, "/todos", http.StatusFound)
	// }
	userUUID := helpers.GetSession(r)
	if userUUID != "" {
		generateHTML(w, helpers.GetFlashes(w, r), "layout", "public_navbar", "flash", "login")
	} else {
		user, err := models.Users(
			qm.Where("uuid = ?", userUUID),
		).One(r.Context(), db.DB)
		if err != nil || user == nil {
			log.Println(err)
			generateHTML(w, helpers.GetFlashes(w, r), "layout", "public_navbar", "flash", "login")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	}
}

// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	user, err := models.GetUserByEmail(r.PostFormValue("email"))
// 	if err != nil {
// 		log.Println(err)
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	}

// 	if user.Password == models.Encrypt(r.PostFormValue("password")) {
// 		session, err := user.CreateSession()
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		cookie := http.Cookie{
// 			Name:     "_cookie",
// 			Value:    session.UUID,
// 			HttpOnly: true,
// 		}
// 		http.SetCookie(w, &cookie)

// 		http.Redirect(w, r, "/", http.StatusFound)
// 	} else {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	}
// }

// func logout(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("_cookie")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err != http.ErrNoCookie {
// 		session := models.Session{UUID: cookie.Value}
// 		session.DeleteSessionByUUID()
// 	}
// 	http.Redirect(w, r, "/login", http.StatusFound)
// }
