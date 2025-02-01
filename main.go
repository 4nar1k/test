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
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Message created successfully")
}

func main() {
	InitDB()
	db.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
