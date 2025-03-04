package handlers

import (
	"apiRUKA/internal/models"
	"apiRUKA/internal/userService" // Исправлен импорт (было `Userservice` с заглавной)
	"apiRUKA/internal/web/users"
	"context"
	"time"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers - Получение всех пользователей
func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Преобразование списка пользователей в OpenAPI-формат
	var response []users.User
	for _, u := range usersList {
		var deletedAt *time.Time
		if u.DeletedAt.Valid { // Учитываем, если поле DeletedAt может быть NULL
			deletedAt = &u.DeletedAt.Time
		}

		response = append(response, users.User{
			Id:        &u.ID,
			Email:     &u.Email,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
			DeletedAt: deletedAt,
		})
	}

	return users.GetUsers200JSONResponse(response), nil
}

// PostUsers - Создание нового пользователя
func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return users.PostUsers201JSONResponse{}, nil
	}

	newUser, err := h.service.CreateUser(models.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	})
	if err != nil {
		return nil, err
	}

	// Обрабатываем возможное NULL-значение для DeletedAt
	var deletedAt *time.Time
	if newUser.DeletedAt.Valid {
		deletedAt = &newUser.DeletedAt.Time
	}

	// Преобразование структуры в OpenAPI-модель
	response := users.User{
		Id:        &newUser.ID,
		Email:     &newUser.Email,
		CreatedAt: &newUser.CreatedAt,
		UpdatedAt: &newUser.UpdatedAt,
		DeletedAt: deletedAt, // Теперь правильно передаем указатель
	}

	return users.PostUsers201JSONResponse(response), nil
}

// PatchUsersId - Обновление данных пользователя по ID
func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	if request.Body == nil {
		return users.PatchUsersId200JSONResponse{}, nil
	}

	updatedUser, err := h.service.UpdateUserByID(uint(request.Id), models.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	})
	if err != nil {
		return nil, err
	}

	// Учитываем возможность NULL в DeletedAt
	var deletedAt *time.Time
	if updatedUser.DeletedAt.Valid {
		deletedAt = &updatedUser.DeletedAt.Time
	}

	// Преобразование структуры
	response := users.User{
		Id:        &updatedUser.ID,
		Email:     &updatedUser.Email,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
		DeletedAt: deletedAt,
	}

	return users.PatchUsersId200JSONResponse(response), nil
}

// DeleteUsersId - Удаление пользователя по ID
func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUserByID(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
