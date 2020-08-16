package main

import (
	"fmt"
	"todo-app/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		fmt.Printf("Error running server: %v\n", err)
	}
	return
}
