package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/talesmasoero/go-task-tracker/internal/domain"
)

type TaskRepository interface {
	Save(task domain.Task) (domain.Task, error)
	ReadAll() ([]domain.Task, error)
	GetByID(id int) (domain.Task, error)
	Update(task domain.Task) error
	Delete(id int) error
}

type TaskHandler struct {
	repo TaskRepository
}

func NewTaskHandler(repo TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if task.Description == "" {
		http.Error(w, "task description cannot be empty", http.StatusBadRequest)
		return
	}

	task, err := th.repo.Save(task)
	if err != nil {
		http.Error(w, "could not save task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.repo.ReadAll()
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

	tasks, err := th.repo.ReadAll()
	if err != nil {
		http.Error(w, "could not read tasks", http.StatusInternalServerError)
		return
	}

	for _, task := range tasks {
		if id == task.ID {
			json, err := json.MarshalIndent(task, "", "  ")
			if err != nil {
				http.Error(w, "error parsing task into json", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
			return
		}
	}
	http.Error(w, "task id doesn't exists", http.StatusNotFound)
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

	if newTask.Description == "" {
		http.Error(w, "task description cannot be empty", http.StatusBadRequest)
		return
	}

	// Talvez retornar a task, igual na Save para receber a task atualizada no front
	if err := th.repo.Update(newTask); err != nil {
		http.Error(w, "error updating task", http.StatusInternalServerError)
		return
	}
}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := th.getUrlID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.repo.Delete(id); err != nil {
		http.Error(w, "error deleting task", http.StatusInternalServerError)
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
