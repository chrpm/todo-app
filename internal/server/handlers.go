package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/data"
	"todo-app/internal/repository"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Return 404 not found
	w.Write([]byte("Get One\n")) // Return object
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Return empty object list when none found
	w.Write([]byte("Get Many\n")) // Return Objects
}

func createTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		dec := json.NewDecoder(r.Body)
		var t data.CreateTaskRequest
		err := dec.Decode(&t)
		if err != nil {
			http.Error(w, fmt.Sprintf("Couldnt parse json body: %s", err.Error()), http.StatusBadRequest)
			return
		}

		goal := t.Goal

		// Request Validation
		if len(goal) == 0 {
			http.Error(w, "Goal cannot be empty", http.StatusBadRequest)
			return
		}
		if len(goal) > config.MaxGoalLength {
			http.Error(w, fmt.Sprintf("Goal max length is %d characters", config.MaxGoalLength), http.StatusBadRequest)
			return
		}

		id, err := dao.InsertTask(goal)
		if err != nil {
			http.Error(w, "DB Write Failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/tasks/%d", id))
		w.WriteHeader(http.StatusCreated)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update One\n"))
}

func updateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update Many\n"))
}

func modifyTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Modify One\n"))
}

func modifyTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Modify Many"))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	// return 404 if not found
	w.Write([]byte("Delete One\n"))
}

func deleteTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Delete Many\n"))
}
