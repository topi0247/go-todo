package controllers

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/app/helpers"
	"todo-app/app/models"
	"todo-app/infrastructure/db"

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
	userUUID := helpers.GetSession(r)
	log.Println("userUUID:", userUUID)
	if userUUID != "" {
		generateHTML(w, helpers.GetFlashes(w, r), "layout", "public_navbar", "flash", "login")
	} else {
		_, err := models.Users(
			qm.Where("uuid = ?", userUUID),
		).One(r.Context(), db.DB)
		if err != nil {
			log.Println(err)
			flash := helpers.GetFlashes(w, r)
			log.Println(flash)
			generateHTML(w, flash, "layout", "public_navbar", "flash", "login")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		helpers.AppendFlash(w, r, helpers.FlashError, "入力値が不正です")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	user, err := models.Users(
		qm.Where("email = ?", r.PostFormValue("email")),
	).One(r.Context(), db.DB)
	if err != nil {
		log.Println(err)
		helpers.AppendFlash(w, r, helpers.FlashError, "認証に失敗しました。再度ログインしてください")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(r.PostFormValue("password"))); err != nil {
		log.Println(err)
		helpers.AppendFlash(w, r, helpers.FlashError, "認証に失敗しました。再度ログインしてください")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	helpers.CreateSession(w, r, user.UUID)
	helpers.AppendFlash(w, r, helpers.FlashSuccess, "ログインしました")
	http.Redirect(w, r, "/todos", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	helpers.ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
