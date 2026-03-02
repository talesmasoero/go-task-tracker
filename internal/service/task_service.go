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

func (ts TaskService) ReadTasks() ([]domain.Task, error) {
	tasks, err := ts.repo.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read tasks: %w", err)
	}
	return tasks, nil
}

func (ts TaskService) GetTaskByID(id int) (domain.Task, error) {
	task, err := ts.repo.GetByID(id)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (ts TaskService) UpdateTask(newTask domain.Task) error {
	if newTask.Description == "" {
		return errors.New("task description cannot be empty")
	}

	if err := ts.repo.Update(newTask); err != nil {
		return fmt.Errorf("could not update task: %w", err)
	}
	return nil
}

func (ts TaskService) DeleteTask(id int) error {
	if err := ts.repo.Delete(id); err != nil {
		return fmt.Errorf("could not delete task: %w", err)
	}
	return nil
}
