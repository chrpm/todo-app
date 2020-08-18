package data

// Task represents the task object in the database
type Task struct {
	ID        int    `json:"id"`
	Goal      string `json:"goal"`
	Completed bool   `json:"completed"`
}

// CreateTaskRequest is used to get data from creation request
type CreateTaskRequest struct {
	Goal string `json:"goal"`
}
