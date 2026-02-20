package main

import (
	"fmt"
	"net/http"
)

var (
	tasks  = []Task{}
	lastID = 0
)

func main() {
	http.HandleFunc("POST /tasks", CreateTask)
	http.HandleFunc("GET /tasks", ReadTasks)
	http.HandleFunc("GET /tasks/{id}", GetTaskByID)
	http.HandleFunc("PUT /tasks/{id}", UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", DeleteTask)

	fmt.Println("Listening on 2525")
	http.ListenAndServe(":2525", nil)
}
