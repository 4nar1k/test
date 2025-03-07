package taskService

import (
	"apiRUKA/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task models.Task) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTasksByUserID(userID uint) ([]models.Task, error) //
	UpdateTaskByID(id uint, task models.Task) (models.Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task models.Task) (models.Task, error) {
	var existingTask models.Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return models.Task{}, err
	}

	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone

	result := r.db.Save(&existingTask)
	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	return result.Error
}
