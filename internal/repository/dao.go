package repository

// DAO defines how to get tasks from database
type DAO interface {
	InsertTask(goal string) (id int64, err error)
}
