package controllers

import (
	"net/http"
	"udemy-todo-app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
