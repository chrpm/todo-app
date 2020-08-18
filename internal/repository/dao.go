package repository

import "todo-app/internal/data"

// DAO defines how to get tasks from database
type DAO interface {
	GetTask(taskID int) (task *data.Task, err error)
	InsertTask(goal string) (id int64, err error)
}
