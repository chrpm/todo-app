package data

import "errors"

var (
	// ErrRecordNotFound is used when a record is missing from the database
	ErrRecordNotFound = errors.New("record not found in db")
)
