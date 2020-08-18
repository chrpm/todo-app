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

		completed := r.URL.Query().Get("completed")
		if completed != "true" && completed != "false" {
			completed = ""
		}

		tasks, err := dao.GetTasks(completed)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		tasksJSON, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tasksJSON))
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

		msg := validateGoal(t.Goal)
		if msg != "" {
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		id, err := dao.InsertTask(t.Goal)
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
		idParam := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idParam, 10, 32)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		dec := json.NewDecoder(r.Body)
		var t data.UpdateTaskRequest
		err = dec.Decode(&t)
		if err != nil {
			http.Error(w, fmt.Sprintf("Couldnt parse json body: %s", err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Printf("%v", t)

		msg := validateGoal(t.Goal)
		if msg != "" {
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		rowsUpdated, err := dao.UpdateTask(int(id), t.Goal, t.Completed)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if rowsUpdated < 1 {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func deleteTaskHandler(dao repository.DAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idParam, 10, 32)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = dao.DeleteTask(int(id))
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func validateGoal(goal string) string {
	// Request Validation
	if len(goal) == 0 {
		return "Goal cannot be empty"
	}
	if len(goal) > config.MaxGoalLength {
		return fmt.Sprintf("Goal max length is %d characters", config.MaxGoalLength)
	}
	return ""
}
