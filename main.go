package main

import (
	"github.com/ardaorkun/go-todo-app/config"
	"github.com/ardaorkun/go-todo-app/internal/api"
	"github.com/ardaorkun/go-todo-app/internal/db"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	db.InitializeDatabase()

	r := chi.NewRouter()
	api.SetRoutes(r)

	http.ListenAndServe(config.GetConfig().Port, r)
}
