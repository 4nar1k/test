package handlers

import (
	"apiRUKA/internal/models"
	"apiRUKA/internal/userService"
	"apiRUKA/internal/web/users"
	"context"
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

	var response []users.User
	for _, u := range usersList {
		response = append(response, users.User{
			Id:    &u.ID,
			Email: &u.Email,
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

	response := users.User{
		Id:    &newUser.ID,
		Email: &newUser.Email,
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
	// Преобразование структуры
	response := users.User{
		Id:    &updatedUser.ID,
		Email: &updatedUser.Email,
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
