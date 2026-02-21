package main

import (
	"fmt"
	"net/http"

	"github.com/talesmasoero/go-task-tracker/internal/handler"
)

func main() {
	http.HandleFunc("POST /tasks", handler.CreateTask)
	http.HandleFunc("GET /tasks", handler.ReadTasks)
	http.HandleFunc("GET /tasks/{id}", handler.GetTaskByID)
	http.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	fmt.Println("Listening on 2525")
	http.ListenAndServe(":2525", nil)
}
