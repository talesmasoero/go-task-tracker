package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"

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

func (repo *JSONRepository) Save(task domain.Task) (domain.Task, error) {
	tasks, err := repo.loadTasks()
	if err != nil {
		return domain.Task{}, err
	}

	task.ID = 1
	if len(tasks) > 0 {
		task.ID = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, task)

	if err := repo.saveTasks(tasks); err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (repo *JSONRepository) ReadAll() ([]domain.Task, error) {
	return repo.loadTasks()
}

func (repo *JSONRepository) GetByID(id int) (domain.Task, error) {
	tasks, err := repo.loadTasks()
	if err != nil {
		return domain.Task{}, err
	}

	idx, err := repo.hasTask(tasks, id)
	if err != nil {
		return domain.Task{}, err
	}
	return tasks[idx], nil
}

func (repo *JSONRepository) Update(newTask domain.Task) error {
	tasks, err := repo.loadTasks()
	if err != nil {
		return err
	}

	idx, err := repo.hasTask(tasks, newTask.ID)
	if err != nil {
		return err
	}

	tasks[idx].Description = newTask.Description

	if err := repo.saveTasks(tasks); err != nil {
		return err
	}
	return nil
}

func (repo *JSONRepository) Delete(id int) error {
	tasks, err := repo.loadTasks()
	if err != nil {
		return err
	}

	idx, err := repo.hasTask(tasks, id)
	if err != nil {
		return err
	}

	tasks = slices.Delete(tasks, idx, idx+1)

	if err := repo.saveTasks(tasks); err != nil {
		return err
	}
	return nil
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

func (repo *JSONRepository) saveTasks(tasks []domain.Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	if err := os.WriteFile(repo.filepath, jsonData, filemode); err != nil {
		return fmt.Errorf("could not save file: %w", err)
	}
	return nil
}

// Returns task index if exists
func (repo *JSONRepository) hasTask(tasks []domain.Task, id int) (int, error) {
	for i, task := range tasks {
		if id == task.ID {
			return i, nil
		}
	}
	return 0, errors.New("could not find task")
}
