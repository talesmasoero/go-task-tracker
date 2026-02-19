package main

import (
	"encoding/json"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err == nil {
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
