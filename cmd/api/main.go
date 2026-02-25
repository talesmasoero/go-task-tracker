package main

import (
	"fmt"
	"net/http"

	"github.com/talesmasoero/go-task-tracker/internal/handler"
	"github.com/talesmasoero/go-task-tracker/internal/repository"
)

func main() {
	repo := repository.NewJSONRepository("test.json")
	h := handler.NewTaskHandler(repo)

	http.HandleFunc("POST /tasks", h.CreateTask)
	http.HandleFunc("GET /tasks", h.ReadTasks)
	http.HandleFunc("GET /tasks/{id}", h.GetTaskByID)
	http.HandleFunc("PUT /tasks/{id}", h.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", h.DeleteTask)

	fmt.Println("Listening on 2525")
	http.ListenAndServe(":2525", nil)
}
