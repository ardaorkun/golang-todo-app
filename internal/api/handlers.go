package api

import (
	"encoding/json"
	"github.com/ardaorkun/go-todo-app/internal/db"
	"github.com/ardaorkun/go-todo-app/internal/models"
	"github.com/ardaorkun/go-todo-app/scripts"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := db.GetTasks()
	scripts.HandleError(w, err, 500)

	responseData, err := json.Marshal(tasks)
	scripts.HandleError(w, err, 500)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseData)
	if err != nil {
		return
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !requireID(w, id) {
		return
	}

	task, err := db.GetTask(id)
	scripts.HandleError(w, err, 500)

	if task == nil {
		http.Error(w, scripts.ErrTaskNotFound, http.StatusNotFound)
		return
	}

	responseData, err := json.Marshal(task)
	scripts.HandleError(w, err, 500)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseData)
	if err != nil {
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	scripts.HandleError(w, err, 400)

	if task.Title == "" || task.Description == "" {
		http.Error(w, scripts.ErrMsgTitleAndDescriptionReq, http.StatusBadRequest)
		return
	}

	err = db.CreateTask(task.Title, task.Description)
	scripts.HandleError(w, err, 500)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Task created successfully"))
	if err != nil {
		return
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !requireID(w, id) {
		return
	}

	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	scripts.HandleError(w, err, 400)

	var title *string
	var description *string
	var completed *bool

	if val, ok := updateData["title"].(string); ok {
		title = &val
	}

	if val, ok := updateData["description"].(string); ok {
		description = &val
	}

	if val, ok := updateData["completed"].(bool); ok {
		completed = &val
	}

	if title == nil && description == nil && completed == nil {
		http.Error(w, scripts.ErrMsgNoValidUpdateData, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(id, title, description, completed)
	scripts.HandleError(w, err, 500)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Task updated successfully"))
	if err != nil {
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !requireID(w, id) {
		return
	}

	err := db.DeleteTask(id)
	scripts.HandleError(w, err, 500)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Task deleted successfully"))
	if err != nil {
		return
	}
}

func requireID(w http.ResponseWriter, id string) bool {
	if id == "" {
		http.Error(w, scripts.ErrMsgTaskIDRequired, http.StatusBadRequest)
		return false
	}
	return true
}
