package userService

import (
	"apiRUKA/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateUser - Создает нового пользователя и возвращает его
	CreateUser(user models.User) (models.User, error)
	// GetAllUsers - Возвращает массив всех пользователей из БД
	GetAllUsers() ([]models.User, error)
	// GetUserByID - Возвращает пользователя по ID
	GetUserByID(id uint) (models.User, error)
	// UpdateUserByID - Обновляет данные пользователя по ID
	UpdateUserByID(id uint, user models.User) (models.User, error)
	// DeleteUserByID - Удаляет пользователя по ID
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(id uint, user models.User) (models.User, error) {
	var existing models.User
	if err := r.db.First(&existing, id).Error; err != nil {
		return models.User{}, err
	}

	result := r.db.Model(&existing).Updates(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return existing, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
