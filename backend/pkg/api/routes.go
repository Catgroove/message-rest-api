package api

import (
	"backend/pkg/handlers"
	mw "backend/pkg/middlewares"
)

// One routes file for now, but could be split into its own package in the future, and
// further divided into versions, etc.
func (a *api) routes() {
	a.router.HandleFunc("/api/v1/messages", mw.Logger(handlers.GetMessages)).Methods("GET")
	a.router.HandleFunc("/api/v1/messages", mw.Logger(handlers.CreateMessage)).Methods("POST")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(handlers.GetMessage)).Methods("GET")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(handlers.UpdateMessage)).Methods("PUT")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(handlers.DeleteMessage)).Methods("DELETE")
}
