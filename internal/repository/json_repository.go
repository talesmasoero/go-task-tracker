package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/talesmasoero/go-task-tracker/internal/domain"
)

const filemode = 0644 // Read/Writer for owner, Read for others

type JSONRepository struct {
	filepath string
}

func NewJSONRepository(filepath string) *JSONRepository {
	return &JSONRepository{
		filepath: filepath,
	}
}

func (repo *JSONRepository) Save(task domain.Task) (int, error) {
	tasks, err := repo.loadTasks()
	if err != nil {
		return 0, err
	}

	task.ID = 1
	if len(tasks) > 0 {
		task.ID = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, task)

	newJson, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return 0, fmt.Errorf("could not marshal json: %w", err)
	}

	if err := os.WriteFile(repo.filepath, newJson, filemode); err != nil {
		return 0, fmt.Errorf("could not save file: %w", err)
	}
	return task.ID, nil
}

func (repo *JSONRepository) ReadAll() ([]domain.Task, error) {
	return repo.loadTasks()
}

func (repo *JSONRepository) loadTasks() ([]domain.Task, error) {
	var tasks []domain.Task

	jsonData, err := os.ReadFile(repo.filepath)
	if errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	}
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	if len(jsonData) == 0 {
		return tasks, nil
	}

	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		return nil, fmt.Errorf("could not unmarshal json: %w", err)
	}
	return tasks, nil
}
