package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/api"
	"backend/pkg/models"
	"backend/pkg/services"
)

func TestGetMessages(t *testing.T) {
	m := services.MessageService.CreateMessage(models.Message{
		Message: "Test Message",
	})

	req := httptest.NewRequest("GET", "/api/v1/messages", nil)
	res := request(req)
	checkResponseCode(t, res.Code, http.StatusOK)

	messages := models.Messages{}
	err := json.NewDecoder(res.Body).Decode(&messages)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if messages[0].ID != m.ID {
		t.Errorf("Expected first message to have id %v, got %v", m.ID, messages[0].ID)
	}
}

func TestCreateMessage(t *testing.T) {
	// Missing body
	req := httptest.NewRequest("POST", "/api/v1/messages", nil)
	res := request(req)
	checkResponseCode(t, res.Code, http.StatusBadRequest)

	// Missing message
	message := []byte(`{"bad_field":"test"}`)
	req = httptest.NewRequest("POST", "/api/v1/messages", bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusBadRequest)

	var m map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if m["error"] != "Message is required" {
		t.Errorf("Expected error \"%v\", got %v", "Message is required", m["error"])
	}

	// messagae > 256
	message = []byte(`{"message":"testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest"}`)
	req = httptest.NewRequest("POST", "/api/v1/messages", bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusBadRequest)

	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if m["error"] != "Message length greater than 200 characters" {
		t.Errorf("Expected error \"%v\", got %v", "Message length greater than 200 characters", m["error"])
	}

	// Valid message
	message = []byte(`{"message":"test"}`)
	req = httptest.NewRequest("POST", "/api/v1/messages", bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusCreated)

	createdMessage := models.Message{}
	err = json.NewDecoder(res.Body).Decode(&createdMessage)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if createdMessage.ID == 0 {
		t.Errorf("Expected message to have an id > 0")
	}

	if createdMessage.Message != "test" {
		t.Errorf("Expected message to be \"%v\", but got \"%v\"", "test", createdMessage.Message)
	}
}

func TestGetMessage(t *testing.T) {
	// Id not found
	req := httptest.NewRequest("GET", "/api/v1/messages/123", nil)
	res := request(req)
	checkResponseCode(t, res.Code, http.StatusNotFound)

	// Id found
	m := services.MessageService.CreateMessage(models.Message{
		Message: "Test Message",
	})

	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/messages/%v", m.ID), nil)
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusOK)

	message := models.Message{}
	err := json.NewDecoder(res.Body).Decode(&message)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if message.ID != m.ID {
		t.Errorf("Expected message to have id %v, got %v", m.ID, message.ID)
	}

	if message.Message != m.Message {
		t.Errorf("Expected message to be \"%v\", but got \"%v\"", m.Message, message.Message)
	}
}

func TestUpdateMessage(t *testing.T) {
	// Missing body
	req := httptest.NewRequest("PUT", "/api/v1/messages/123", nil)
	res := request(req)
	checkResponseCode(t, res.Code, http.StatusBadRequest)

	// Id not found
	message := []byte(`{"message":"updated message"}`)
	req = httptest.NewRequest("PUT", "/api/v1/messages/123", bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusNotFound)

	// Id found but missing message
	m := services.MessageService.CreateMessage(models.Message{
		Message: "Test Message",
	})
	message = []byte(`{"unknown_field":"updated message"}`)
	req = httptest.NewRequest("PUT", fmt.Sprintf("/api/v1/messages/%v", m.ID), bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusBadRequest)

	// Id found with correct message
	message = []byte(`{"message":"updated message"}`)
	req = httptest.NewRequest("PUT", fmt.Sprintf("/api/v1/messages/%v", m.ID), bytes.NewBuffer(message))
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusOK)

	updatedMessage := models.Message{}
	err := json.NewDecoder(res.Body).Decode(&updatedMessage)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if updatedMessage.ID != m.ID {
		t.Errorf("Expected to get back id of %v, but got %v", m.ID, updatedMessage.ID)
	}

	if updatedMessage.Message != "updated message" {
		t.Errorf("Expected message to be \"%v\"", "updated message")
	}
}

func TestDeleteMessage(t *testing.T) {
	// Id not found
	req := httptest.NewRequest("DELETE", "/api/v1/messages/123", nil)
	res := request(req)
	checkResponseCode(t, res.Code, http.StatusNotFound)

	// Id found
	m := services.MessageService.CreateMessage(models.Message{
		Message: "Test Message",
	})

	req = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/messages/%v", m.ID), nil)
	res = request(req)
	checkResponseCode(t, res.Code, http.StatusOK)

	var status map[string]bool
	err := json.NewDecoder(res.Body).Decode(&status)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if !status["deleted"] {
		t.Errorf("Expected deleted to be true, but got %v", status["deleted"])
	}
}

func request(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	API := api.CreateAPI()
	API.ServeHTTP(rec, req)

	return rec
}

func checkResponseCode(t *testing.T, actual, expected int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d", expected, actual)
	}
}
