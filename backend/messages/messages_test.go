package messages

import (
	"testing"
	"reflect"
)

func TestGetAllMessages(t *testing.T) {
	MessageService.CreateMessage(Message{
		Message: "This is a message",
	})

	got := MessageService.GetAllMessages()

	if (reflect.TypeOf(got) != reflect.TypeOf(Messages{})) {
		t.Errorf("Expected slice to be of type %s, but got %s", reflect.TypeOf(got), reflect.TypeOf(Messages{}))
	}
}

func TestGetMessage(t *testing.T) {
	m := MessageService.CreateMessage(Message{
		Message: "This is a message",
	})

	got := MessageService.GetMessage(m.ID)

	if !reflect.DeepEqual(got.ID, m.ID) {
		t.Errorf("Expected created message to be equal to retrieval message")
	}
}

func TestCreateMessage(t *testing.T) {
	got := MessageService.CreateMessage(Message{
		Message: "This is a message",
	})

	if got.Message != "This is a message" {
		t.Errorf("Expected %s, got %s", "This is a message", got.Message)
	}

	if reflect.TypeOf(got) != reflect.TypeOf(Message{}) {
		t.Errorf("Expected message to be of type %s, got %s", reflect.TypeOf(Message{}), reflect.TypeOf(got))
	}
}

func TestDeleteMessage(t *testing.T) {
	createdMessage := MessageService.CreateMessage(Message{
		Message: "This is a message",
	})

	MessageService.DeleteMessage(createdMessage.ID)
	messages := MessageService.GetAllMessages()

	for _, m := range messages {
		if m.ID == createdMessage.ID {
			t.Errorf("Found deleted message in slice of messages")
		}
	}
}

func TestUpdateMessage(t *testing.T) {
	createdMessage := MessageService.CreateMessage(Message{
		Message: "This is a message",
	})

	updateMessage := Message{Message: "New message"}
	newMessage := MessageService.UpdateMessage(createdMessage.ID, updateMessage)

	if (newMessage.Message != updateMessage.Message) {
		t.Errorf("Expected newMessage to have message of \"%s\", but got \"%s\"", updateMessage.Message, newMessage.Message)
	}
}
