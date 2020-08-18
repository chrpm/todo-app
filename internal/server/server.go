package server

import (
	"log"
	"net/http"
	"todo-app/internal/repository"

	"github.com/gorilla/mux"
)

// Run starts web server for todo app
func Run() (err error) {

	db, err := repository.DBConnection("root", "wWeJJt3Iqc_D", "/todo_app")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	defer db.Close()

	dao := repository.NewTaskDAO(db)

	router := initalizeRoutes(dao)
	port := ":5050"
	log.Printf("Starting Http server on port %v", port)
	return http.ListenAndServe(port, router)
}

func initalizeRoutes(dao repository.DAO) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/tasks").Subrouter()

	s.HandleFunc("", getTasksHandler(dao)).Methods("GET")
	s.HandleFunc("/{id}", getTaskHandler(dao)).Methods("GET")
	s.HandleFunc("/{id}", updateTaskHandler(dao)).Methods("PUT")
	s.HandleFunc("", createTaskHandler(dao)).Methods("POST")
	s.HandleFunc("/{id}", modifyTaskHandler(dao)).Methods("PATCH")
	s.HandleFunc("/{id}", deleteTaskHandler(dao)).Methods("DELETE")

	return s
}
