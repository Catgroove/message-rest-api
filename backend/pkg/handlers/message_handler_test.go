package handlers_test

import (
	"github.com/gorilla/mux"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/models"
	"backend/pkg/handlers"
	"backend/pkg/services"
)

func TestGetMessages(t *testing.T) {
	ms, ts := preTest()
	defer ts.Close()

	m := ms.CreateMessage(models.Message{
		Message: "Test Message",
	})

	res, err := http.Get(ts.URL + "/messages")
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusOK)

	messages := models.Messages{}
	err = json.NewDecoder(res.Body).Decode(&messages)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if messages[0].ID != m.ID {
		t.Errorf("Expected first message to have id %v, got %v", m.ID, messages[0].ID)
	}
}

func TestCreateMessage(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	message := []byte(`{"message":"test"}`)
	res, err := http.Post(ts.URL + "/messages", "application/json", bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusCreated)

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

func TestCreateMessageMissingBody(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	res, err := http.Post(ts.URL + "/messages", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusBadRequest)
}

func TestCreateMessageMissingMessage(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	message := []byte(`{"bad_field":"test"}`)
	res, err := http.Post(ts.URL + "/messages", "application/json", bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusBadRequest)

	var m map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if m["error"] != "Message is required" {
		t.Errorf("Expected error \"%v\", got %v", "Message is required", m["error"])
	}
}

func TestCreateMessageTooLongMesage(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	message := []byte(`{"message":"testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest"}`)
	res, err := http.Post(ts.URL + "/messages", "application/json", bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusBadRequest)

	var m map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if m["error"] != "Message length greater than 200 characters" {
		t.Errorf("Expected error \"%v\", got %v", "Message length greater than 200 characters", m["error"])
	}
}

func TestGetMessage(t *testing.T) {
	ms, ts := preTest()
	defer ts.Close()

	m := ms.CreateMessage(models.Message{
		Message: "Test Message",
	})

	res, err := http.Get(ts.URL + fmt.Sprintf("/messages/%v", m.ID))
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusOK)

	message := models.Message{}
	err = json.NewDecoder(res.Body).Decode(&message)
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

func TestGetMessageNotFound(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	res, err := http.Get(ts.URL + "/messages/123")
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, res.StatusCode, http.StatusNotFound)
}

func TestUpdateMessage(t *testing.T) {
	ms, ts := preTest()
	defer ts.Close()

	m := ms.CreateMessage(models.Message{
		Message: "Test Message",
	})

	message := []byte(`{"message":"updated message"}`)
	req, err := http.NewRequest(http.MethodPut, ts.URL + fmt.Sprintf("/messages/%v", m.ID), bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusOK)

	updatedMessage := models.Message{}
	err = json.NewDecoder(res.Body).Decode(&updatedMessage)
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

func TestUpdateMessageMissingBody(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	req, err := http.NewRequest(http.MethodPut, ts.URL + "/messages/123", nil)
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusBadRequest)
}

func TestUpdateMessageNotFound(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	message := []byte(`{"message":"updated message"}`)
	req, err := http.NewRequest(http.MethodPut, ts.URL + "/messages/123", bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusNotFound)
}

func TestUpdateMessageMissingMessage(t *testing.T) {
	ms, ts := preTest()
	defer ts.Close()

	m := ms.CreateMessage(models.Message{
		Message: "Test Message",
	})

	message := []byte(`{"unknown_field":"updated message"}`)
	req, err := http.NewRequest(http.MethodPut, ts.URL + fmt.Sprintf("/messages/%v", m.ID), bytes.NewBuffer(message))
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusBadRequest)
}

func TestDeleteMessage(t *testing.T) {
	ms, ts := preTest()
	defer ts.Close()

	m := ms.CreateMessage(models.Message{
		Message: "Test Message",
	})

	req, err := http.NewRequest(http.MethodDelete, ts.URL + fmt.Sprintf("/messages/%v", m.ID), nil)
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusOK)

	var status map[string]bool
	err = json.NewDecoder(res.Body).Decode(&status)
	if err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if !status["deleted"] {
		t.Errorf("Expected deleted to be true, but got %v", status["deleted"])
	}
}

func TestDeleteMessageNotFound(t *testing.T) {
	_, ts := preTest()
	defer ts.Close()

	req, err := http.NewRequest(http.MethodDelete, ts.URL + "/messages/123", nil)
	if err != nil {
		t.Fatal(nil)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(nil)
	}

	checkResponseCode(t, res.StatusCode, http.StatusNotFound)
}

func preTest() (*services.MessageService, *httptest.Server) {
	r := mux.NewRouter().StrictSlash(true)

	ms := services.NewMessageService()
	handlers.NewMessageHandler(ms, r)

	ts := httptest.NewServer(r)

	return ms, ts
}

func checkResponseCode(t *testing.T, actual, expected int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d", expected, actual)
	}
}
