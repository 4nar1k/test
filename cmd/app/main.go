package main

import (
	"apiRUKA/internal/database"
	"apiRUKA/internal/handlers"
	"apiRUKA/internal/middleware"
	"apiRUKA/internal/taskService"
	"apiRUKA/internal/userService"
	"apiRUKA/internal/web/tasks"
	"apiRUKA/internal/web/users"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создание репозиториев
	tasksRepo := taskService.NewTaskRepository(database.DB)
	userRepo := userService.NewUserRepository(database.DB)

	// Создание сервисов
	tasksService := taskService.NewTaskService(tasksRepo)
	userService := userService.NewUserService(userRepo)

	// Создание обработчиков
	tasksHandler := handlers.NewTaskHandler(tasksService)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализация Echo
	e := echo.New()

	// Подключение стандартных middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	// Middleware для связывания context.Context и echo.Context
	e.Use(middleware.AttachEchoContextMiddleware)

	// Регистрация обработчиков
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)

	tasks.RegisterHandlers(e, tasksStrictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
