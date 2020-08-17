package main

import (
	"log"
	"todo-app/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalf("Error running server: %v\n", err)
	}
	return
}
