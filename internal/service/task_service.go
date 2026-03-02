package service

import (
	"errors"
	"fmt"

	"github.com/talesmasoero/go-task-tracker/internal/domain"
)

type TaskRepository interface {
	Save(task domain.Task) (domain.Task, error)
	ReadAll() ([]domain.Task, error)
	GetByID(id int) (domain.Task, error)
	Update(task domain.Task) error
	Delete(id int) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (ts TaskService) CreateTask(task domain.Task) (domain.Task, error) {
	if task.Description == "" {
		return domain.Task{}, errors.New("task description cannot be empty")
	}

	task, err := ts.repo.Save(task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("could not save task: %w", err)
	}
	return task, nil
}
