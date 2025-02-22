package taskService

import (
	"apiRUKA/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task models.Task) (models.Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]models.Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task models.Task) (models.Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
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
func (r *taskRepository) UpdateTaskByID(id uint, task models.Task) (models.Task, error) {
	var existing models.Task
	//  Проверяем, существует ли задача в базе
	if err := r.db.First(&existing, id).Error; err != nil {
		return models.Task{}, err // Если нет, возвращаем ошибку
	}

	//   Обновляем только переданные поля
	result := r.db.Model(&existing).Updates(task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return existing, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	return result.Error
}
