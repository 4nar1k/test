package userService

import (
	"apiRUKA/internal/models"
	"apiRUKA/internal/taskService"
)

type UserService struct {
	repo        UserRepository
	taskService *taskService.TaskService
}

func NewUserService(repo UserRepository, taskService *taskService.TaskService) *UserService {
	return &UserService{repo: repo, taskService: taskService}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUserByID(id uint, user models.User) (models.User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	db := s.repo.GetDB()
	return db.Unscoped().Delete(&models.User{}, id).Error
}

func (s *UserService) GetTasksForUser(userID uint) ([]models.Task, error) {
	return s.taskService.GetTasksByUserID(userID)
}
