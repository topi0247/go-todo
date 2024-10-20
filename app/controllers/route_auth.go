package controllers

import (
	"fmt"
	"log"
	"net/http"
	"udemy-todo-app/app/models"
	"udemy-todo-app/infrastructure/db"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := &models.User{
			UUID:           uuid.New().String(),
			Name:           r.PostFormValue("name"),
			Email:          r.PostFormValue("email"),
			PasswordDigest: encrypt(r.PostFormValue("password")),
		}

		if err := user.Insert(r.Context(), db.DB, boil.Infer()); err != nil {
			log.Println(err)
			fmt.Printf("Failed to create user: %s\n", err)
			http.Redirect(w, r, "/signup", http.StatusFound)
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	}
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	_, err := session(w, r)
// 	if err != nil {
// 		generateHTML(w, nil, "layout", "public_navbar", "login")
// 	} else {
// 		http.Redirect(w, r, "/todos", http.StatusFound)
// 	}
// }

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
