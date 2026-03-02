package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/talesmasoero/go-task-tracker/internal/domain"
)

type TaskService interface {
	CreateTask(task domain.Task) (domain.Task, error)
	ReadTasks() ([]domain.Task, error)
	GetTaskByID(id int) (domain.Task, error)
	UpdateTask(newTask domain.Task) error
	DeleteTask(id int) error
}

type TaskHandler struct {
	svc TaskService
}

func NewTaskHandler(svc TaskService) *TaskHandler {
	return &TaskHandler{
		svc: svc,
	}
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	task, err := th.svc.CreateTask(task)
	if err != nil {
		http.Error(w, "could not save task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "could not encode task", http.StatusInternalServerError)
	}
}

func (th *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.svc.ReadTasks()
	if err != nil {
		http.Error(w, "could not read tasks", http.StatusInternalServerError)
		return
	}

	json, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		http.Error(w, "error reading tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (th *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := th.getUrlID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := th.svc.GetTaskByID(id)
	if err != nil {
		http.Error(w, "task does not exist", http.StatusInternalServerError)
		return
	}

	json, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		http.Error(w, "error parsing task into json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := th.getUrlID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask := domain.Task{
		ID: id,
	}

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if err := th.svc.UpdateTask(newTask); err != nil {
		http.Error(w, "could not update task", http.StatusInternalServerError)
	}
}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := th.getUrlID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.svc.DeleteTask(id); err != nil {
		http.Error(w, "could not delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (th *TaskHandler) getUrlID(r *http.Request) (int, error) {
	strID := r.PathValue("id")
	if strID == "" {
		return 0, errors.New("error getting id from url")
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		return 0, errors.New("task id must be a number")
	}
	return id, nil
}
