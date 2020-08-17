package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run starts web server for todo app
func Run() (err error) {
	r := mux.NewRouter()
	s := r.PathPrefix("/tasks").Subrouter()

	// Single object endpoints
	s.HandleFunc("/{id}", getTask).Methods("GET")
	s.HandleFunc("/{id}", updateTask).Methods("PUT")
	s.HandleFunc("/{id}", createTask).Methods("POST")
	s.HandleFunc("/{id}", modifyTask).Methods("PATCH")
	s.HandleFunc("/{id}", deleteTask).Methods("DELETE")

	// Multiple object endpoints
	s.HandleFunc("", getTasks).Methods("GET")
	s.HandleFunc("", updateTasks).Methods("PUT")
	s.HandleFunc("", createTasks).Methods("POST")
	s.HandleFunc("", modifyTasks).Methods("PATCH")
	s.HandleFunc("", deleteTasks).Methods("DELETE")

	port := ":8080"
	log.Printf("Starting Http server on port %v", port)
	return http.ListenAndServe(port, s)
}
