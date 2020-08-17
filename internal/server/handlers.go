package server

import "net/http"

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

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "objectID")
	w.WriteHeader(http.StatusCreated)
	// return 400 bad request when cant create
	w.Write([]byte("Create One\n")) // Return body of object
}

func createTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// return 400 bad request when cant create
	w.Write([]byte("Create Many\n")) // Return body of objects
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
