package taskService

import "apiRUKA/internal/models"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task models.Task) (models.Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
