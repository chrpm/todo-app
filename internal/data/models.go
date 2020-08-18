package data

// Task represents the task object in the database
type Task struct {
	ID        int    `json:"id"`
	Goal      string `json:"goal"`
	Completed bool   `json:"completed"`
}

// Tasks represents a group of tasks for api response
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// CreateTaskRequest is used to get data from creation request
type CreateTaskRequest struct {
	Goal string `json:"goal"`
}

// UpdateTaskRequest is used to update a task
type UpdateTaskRequest struct {
	Goal      string `json:"goal"`
	Completed bool   `json:"completed"`
}
