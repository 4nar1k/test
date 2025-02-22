package main

import (
	"apiRUKA/internal/database"
	"apiRUKA/internal/handlers"
	"apiRUKA/internal/middleware"
	"apiRUKA/internal/taskService"
	"apiRUKA/internal/web/tasks"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создание репозитория и сервиса
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	// Создание обработчика
	handler := handlers.NewHandler(service)

	// Инициализация Echo
	e := echo.New()

	// Подключение стандартных middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	// Middleware для связывания context.Context и echo.Context
	e.Use(middleware.AttachEchoContextMiddleware)

	// Регистрация обработчиков
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
