package handlers

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/pkg/services"
	"backend/pkg/models"
	"backend/pkg/utils"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := services.MessageService.GetAllMessages()

	utils.ResponseToJSON(w, r, messages, http.StatusOK);
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	_ = json.NewDecoder(r.Body).Decode(&m)
	message := services.MessageService.CreateMessage(m)

	utils.ResponseToJSON(w, r, message, http.StatusCreated)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])
	message := services.MessageService.GetMessage(messageId)

	utils.ResponseToJSON(w, r, message, http.StatusOK);
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])

	var m models.Message
	_ = json.NewDecoder(r.Body).Decode(&m)

	message := services.MessageService.UpdateMessage(messageId, m)

	utils.ResponseToJSON(w, r, message, http.StatusOK);
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageId, _ := strconv.Atoi(params["id"])
	services.MessageService.DeleteMessage(messageId);
	messages := services.MessageService.GetAllMessages();

	utils.ResponseToJSON(w, r, messages, http.StatusOK);
}
