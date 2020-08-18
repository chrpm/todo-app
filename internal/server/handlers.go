package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/internal/config"
	"todo-app/internal/data"
	"todo-app/internal/repository"

	"github.com/gorilla/mux"
)

func getTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idParam, 10, 32)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		task, err := dao.GetTask(int(id))
		if err != nil {
			if errors.Is(err, data.ErrRecordNotFound) {
				http.Error(w, "", http.StatusNotFound)
			} else {
				http.Error(w, "", http.StatusInternalServerError)
			}
			return
		}

		taskJSON, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(taskJSON))
	}
}
func getTasksHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Return empty object list when none found
		w.Write([]byte("Get Many\n")) // Return Objects
	}
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

func updateTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update One\n"))
	}
}

func modifyTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Modify One\n"))
	}
}

func deleteTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		// return 404 if not found
		w.Write([]byte("Delete One\n"))
	}
}
