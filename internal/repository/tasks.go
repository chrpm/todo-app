package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import for driver
)

// DBConnection initalizes a sql.DB instance
func DBConnection(user string, password string, database string) (db *sql.DB, err error) {
	connString := fmt.Sprintf("%s:%s@%s", user, password, database)
	db, err = sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return
}

// NewTaskDAO returns new instance of TaskDAO
func NewTaskDAO(db *sql.DB) *TaskDAO {
	return &TaskDAO{DB: db}
}

// TaskDAO implements DAO interface to write tasks to DB
type TaskDAO struct {
	DB *sql.DB
}

// InsertTask places a task object into the database
func (dao *TaskDAO) InsertTask(goal string) (id int64, err error) {
	stmt, err := dao.DB.Prepare("INSERT INTO tasks (goal) VALUES (?)")
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(goal)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
