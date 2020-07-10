package api

import "backend/pkg/handlers"

// One routes file for now, but could be split into its own package in the future, and
// further divided into versions, etc.
func (a *api) routes() {
	a.router.HandleFunc("/api/v1/messages", handlers.GetMessages).Methods("GET")
	a.router.HandleFunc("/api/v1/messages", handlers.CreateMessage).Methods("POST")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", handlers.GetMessage).Methods("GET")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", handlers.UpdateMessage).Methods("PUT")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", handlers.DeleteMessage).Methods("DELETE")
}