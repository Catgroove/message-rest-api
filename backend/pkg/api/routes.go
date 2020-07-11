package api

import (
	"github.com/gorilla/mux"

	h "backend/pkg/handlers"
)

// One routes file for now, but could be split into its own package in the future, and
// further divided into versions, etc.
func (a *api) routes(r *mux.Router) {
	messageRouter := r.PathPrefix("/messages").Subrouter()
	messageRouter.HandleFunc("", h.GetMessages).Methods("GET")
	messageRouter.HandleFunc("", h.CreateMessage).Methods("POST")
	messageRouter.HandleFunc("/{id:[0-9]+}", h.GetMessage).Methods("GET")
	messageRouter.HandleFunc("/{id:[0-9]+}", h.UpdateMessage).Methods("PUT")
	messageRouter.HandleFunc("/{id:[0-9]+}", h.DeleteMessage).Methods("DELETE")
}
