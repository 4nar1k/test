package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	result := db.Find(&messages)
	if result.Error != nil {
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var messageReq Message
	err := json.NewDecoder(r.Body).Decode(&messageReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result := db.Create(&messageReq)
	if result.Error != nil {
		http.Error(w, "Error creating message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messageReq)
	fmt.Fprint(w)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var message Message
	result := db.First(&message, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "Error fetching message", http.StatusInternalServerError)
		return
	}
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	validUpdates := make(map[string]interface{})
	for key, val := range updates {
		switch key {
		case "task":
			if task, ok := val.(string); ok {
				validUpdates["task"] = task
			} else {
				http.Error(w, "Invalid type for task", http.StatusBadRequest)
				return
			}
		case "is_done":
			if isDone, ok := val.(bool); ok {
				validUpdates["is_done"] = isDone
			} else {
				http.Error(w, "Invalid type for is_done", http.StatusBadRequest)
				return
			}
		default:
			continue
		}
	}
	if len(validUpdates) == 0 {
		http.Error(w, "No valid fields to update", http.StatusBadRequest)
		return
	}
	result = db.Model(&message).Updates(validUpdates)
	if result.Error != nil {
		http.Error(w, "Error updating message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var message Message
	result := db.First(&message, "id = ?", id) // Проверяем, есть ли запись в базе
	if result.Error != nil {
		http.Error(w, `{"error": "Message not found"}`, http.StatusNotFound)
		return
	}

	result = db.Delete(&message) // Удаляем найденное сообщение
	if result.Error != nil {
		http.Error(w, `{"error": "Failed to delete message"}`, http.StatusInternalServerError)
		return
	}

	// Формируем JSON-ответ
	response := map[string]string{
		"message": "Message deleted successfully",
		"id":      id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK вместо 204 No Content
	json.NewEncoder(w).Encode(response)
}

func main() {
	InitDB()
	db.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
