package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"backend/messages"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/messages", getMessages).Methods("GET")
	r.HandleFunc("/api/v1/messages", createMessage).Methods("POST")
	r.HandleFunc("/api/v1/messages/{id:[0-9]+}", getMessage).Methods("GET")
	r.HandleFunc("/api/v1/messages/{id:[0-9]+}", updateMessage).Methods("PUT")
	r.HandleFunc("/api/v1/messages/{id:[0-9]+}", deleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	messages := messages.MessageService.GetAllMessages()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var m messages.Message
	_ = json.NewDecoder(r.Body).Decode(&m)
	message := messages.MessageService.CreateMessage(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])
	message := messages.MessageService.GetMessage(messageId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])

	var m messages.Message
	_ = json.NewDecoder(r.Body).Decode(&m)

	message := messages.MessageService.UpdateMessage(messageId, m)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])
	messages.MessageService.DeleteMessage(messageId);
	messages := messages.MessageService.GetAllMessages();

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
