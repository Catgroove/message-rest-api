package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"backend/pkg/models"
	"backend/pkg/services"
	"backend/pkg/utils"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := services.MessageService.GetAllMessages()

	utils.ResponseToJSON(w, r, messages, http.StatusOK)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	err := utils.RequestFromJSON(w, r, &m)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	message, err := services.MessageService.CreateMessage(m)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	utils.ResponseToJSON(w, r, message, http.StatusCreated)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	message, err := services.MessageService.GetMessage(messageId)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusNotFound)
		return
	}

	utils.ResponseToJSON(w, r, message, http.StatusOK)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
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

	message, err := services.MessageService.UpdateMessage(messageId, m)
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

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusBadRequest)
		return
	}

	err = services.MessageService.DeleteMessage(messageId)
	if err != nil {
		utils.ErrorToJSON(w, r, err, http.StatusNotFound)
		return
	}

	utils.ResponseToJSON(w, r, map[string]bool{"deleted": true}, http.StatusOK)
}
