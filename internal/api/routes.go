package api

import "github.com/go-chi/chi/v5"

func SetRoutes(r *chi.Mux) {
	r.Get("/tasks", GetTasks)
	r.Get("/tasks/{id}", GetTask)
	r.Post("/tasks", CreateTask)
	r.Patch("/tasks/{id}", UpdateTask)
	r.Delete("/tasks/{id}", DeleteTask)
}
