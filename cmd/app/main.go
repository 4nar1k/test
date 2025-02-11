package main

import (
	"apiRUKA/internal/database"
	"apiRUKA/internal/handlers"
	"apiRUKA/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	if database.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	database.DB.AutoMigrate(&taskService.Task{})

	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)

	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
