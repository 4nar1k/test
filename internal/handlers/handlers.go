package handlers

import (
	"apiRUKA/internal/models"
	"apiRUKA/internal/taskService"
	"apiRUKA/internal/web/tasks"
	"context"
	"errors"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// GetTasks возвращает список всех задач
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	return h.getTasks(ctx, request)
}

func (h *Handler) getTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		userId := int(tsk.UserID) // Преобразование uint в int
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &userId,
		}
		response = append(response, task)
	}
	return response, nil
}

// GetUsersIdTasks возвращает список задач конкретного пользователя
func (h *Handler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	tasksByUser, err := h.Service.GetTasksByUserID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	response := tasks.GetUsersIdTasks200JSONResponse{}
	for _, tsk := range tasksByUser {
		userId := int(tsk.UserID)
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &userId,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body.UserId == nil {
		return nil, errors.New("user_id is required")
	}

	taskToCreate := models.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
		UserID: uint(*request.Body.UserId), // Преобразование int в uint
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	userId := int(createdTask.UserID)
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &userId,
	}

	return response, nil
}

// PatchTasksId обновляет задачу по ID
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}
	taskToUpdate := models.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	userId := int(updatedTask.UserID)
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &userId,
	}
	return response, nil
}

// DeleteTasksId удаляет задачу по ID
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
