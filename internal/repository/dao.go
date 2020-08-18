package repository

import "todo-app/internal/data"

// DAO defines how to get tasks from database
type DAO interface {
	GetTask(taskID int) (task *data.Task, err error)
	GetTasks(completed string) (tasks *data.Tasks, err error)
	InsertTask(goal string) (id int64, err error)
	DeleteTask(taskID int) (err error)
	UpdateTask(taskID int, goal string, completed bool) (rowsUpdated int64, err error)
}
