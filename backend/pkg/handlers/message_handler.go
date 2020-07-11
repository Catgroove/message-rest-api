package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"backend/pkg/models"
	"backend/pkg/services"
	"backend/pkg/utils"
)

type MessageHandler struct {
	ms *services.MessageService
}

func NewMessageHandler(ms *services.MessageService, r *mux.Router) {
	mh := MessageHandler{ms}
	messageRouter := r.PathPrefix("/messages").Subrouter()
	messageRouter.HandleFunc("", mh.GetMessages).Methods("GET")
	messageRouter.HandleFunc("", mh.CreateMessage).Methods("POST")
	messageRouter.HandleFunc("/{id:[0-9]+}", mh.GetMessage).Methods("GET")
	messageRouter.HandleFunc("/{id:[0-9]+}", mh.UpdateMessage).Methods("PUT")
	messageRouter.HandleFunc("/{id:[0-9]+}", mh.DeleteMessage).Methods("DELETE")
}

func (mh MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := mh.ms.GetAllMessages()

	utils.ResponseToJSON(w, r, messages, http.StatusOK)
}

func (mh MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	err := utils.RequestFromJSON(w, r, &m)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	err = m.Validate()
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	m.BeforeCreate()
	message := mh.ms.CreateMessage(m)

	utils.ResponseToJSON(w, r, message, http.StatusCreated)
}

func (mh MessageHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	message, err := mh.ms.GetMessage(messageId)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusNotFound)
		return
	}

	utils.ResponseToJSON(w, r, message, http.StatusOK)
}

func (mh MessageHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	var m models.Message
	err = utils.RequestFromJSON(w, r, &m)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	err = m.Validate()
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	m.BeforeUpdate()
	m.ID = messageId
	message, err := mh.ms.UpdateMessage(m)
	if err != nil {
		if err.Error() == "Message could not be found" {
			utils.ErrorToJSON(w, r, err, http.StatusNotFound)
		} else {
			utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		}
		return
	}

	utils.ResponseToJSON(w, r, message, http.StatusOK)
}

func (mh MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	err = mh.ms.DeleteMessage(messageId)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusNotFound)
		return
	}

	utils.ResponseToJSON(w, r, map[string]bool{"deleted": true}, http.StatusOK)
}
