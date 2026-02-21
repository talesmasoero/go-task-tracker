package handler

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/talesmasoero/go-task-tracker/internal/domain"
)

var (
	tasks  []domain.Task
	lastID int
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if task.Description == "" {
		http.Error(w, "task description cannot be empty", http.StatusBadRequest)
		return
	}
	task.ID = lastID + 1

	tasks = append(tasks, task)
	lastID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func ReadTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		http.Error(w, "error reading tasks", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskID(r)
	if err != "" {
		http.Error(w, err, http.StatusBadRequest)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskID(r)
	if err != "" {
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	var newTask domain.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if newTask.Description == "" {
		http.Error(w, "task description cannot be empty", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if id == task.ID {
			tasks[i].Description = newTask.Description

			json, err := json.MarshalIndent(tasks[i], "", "  ")
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskID(r)
	if err != "" {
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if id == task.ID {
			tasks = slices.Delete(tasks, i, i+1)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "task id doesn't exists", http.StatusNotFound)
}

func getTaskID(r *http.Request) (int, string) {
	strID := r.PathValue("id")
	if strID == "" {
		return 0, "error getting id from url"
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		return 0, "task id must be a number"
	}
	return id, ""
}
