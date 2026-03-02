package main

import (
	"fmt"
	"net/http"

	"github.com/talesmasoero/go-task-tracker/internal/handler"
	"github.com/talesmasoero/go-task-tracker/internal/repository"
	"github.com/talesmasoero/go-task-tracker/internal/service"
)

func main() {
	repo := repository.NewJSONRepository("test.json")
	svc := service.NewTaskService(repo)
	hdlr := handler.NewTaskHandler(svc)

	http.HandleFunc("POST /tasks", hdlr.CreateTask)
	http.HandleFunc("GET /tasks", hdlr.ReadTasks)
	http.HandleFunc("GET /tasks/{id}", hdlr.GetTaskByID)
	http.HandleFunc("PUT /tasks/{id}", hdlr.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", hdlr.DeleteTask)

	fmt.Println("Listening on 2525")
	http.ListenAndServe(":2525", nil)
}
