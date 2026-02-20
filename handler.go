package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
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
	strID := r.PathValue("id")
	if strID == "" {
		http.Error(w, "error getting id from url", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "task id must be a number", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if ID == task.ID {
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
