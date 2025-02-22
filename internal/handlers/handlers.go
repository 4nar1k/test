package handlers

import (
	"apiRUKA/internal/models"
	"apiRUKA/internal/taskService"
	"apiRUKA/internal/web/tasks"
	"context"
)

// Handler структура, которая реализует StrictServerInterface
type Handler struct {
	Service *taskService.TaskService
}

// NewHandler создает новый экземпляр Handler
func NewHandler(service *taskService.TaskService) *Handler {
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
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

// PostTasks создает новую задачу
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	return h.postTasks(ctx, request)
}

func (h *Handler) postTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := models.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

// PatchTasksId обновляет задачу по ID
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	return h.patchTasksId(ctx, request)
}

func (h *Handler) patchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := models.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(taskID), taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

// DeleteTasksId удаляет задачу по ID
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	return h.deleteTasksId(ctx, request)
}

func (h *Handler) deleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id
	err := h.Service.DeleteTaskByID(uint(taskID))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
