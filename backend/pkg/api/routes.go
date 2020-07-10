package api

import (
	h "backend/pkg/handlers"
	mw "backend/pkg/middlewares"
)

// One routes file for now, but could be split into its own package in the future, and
// further divided into versions, etc.
func (a *api) routes() {
	a.router.HandleFunc("/api/v1/messages", mw.Logger(h.GetMessages)).Methods("GET")
	a.router.HandleFunc("/api/v1/messages", mw.Logger(h.CreateMessage)).Methods("POST")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(h.GetMessage)).Methods("GET")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(h.UpdateMessage)).Methods("PUT")
	a.router.HandleFunc("/api/v1/messages/{id:[0-9]+}", mw.Logger(h.DeleteMessage)).Methods("DELETE")
}
