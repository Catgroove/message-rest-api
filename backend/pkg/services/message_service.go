package services

import (
	"fmt"

	"backend/pkg/models"
)

// In a real application, the service wouldn't store the data
type messageService struct {
	allMessages models.Messages
}

func (s *messageService) GetAllMessages() models.Messages {
	return s.allMessages
}

func (s *messageService) GetMessage(id int) (models.Message, error) {
	for _, m := range s.allMessages {
		if m.ID == id {
			return m, nil
		}
	}

	return models.Message{}, fmt.Errorf("Message could not be found")
}

func (s *messageService) CreateMessage(m models.Message) models.Message {
	s.allMessages = append(s.allMessages, m)
	return m
}

func (s *messageService) DeleteMessage(id int) error {
	var found bool
	for index, m := range s.allMessages {
		if m.ID == id {
			found = true
			s.allMessages = append(s.allMessages[:index], s.allMessages[index+1:]...)
			break
		}
	}

	if !found {
		return fmt.Errorf("Message could not be found")
	}

	return nil
}

func (s *messageService) UpdateMessage(updatedMessage models.Message) (models.Message, error) {
	for index, m := range s.allMessages {
		if m.ID == updatedMessage.ID {
			s.allMessages = append(s.allMessages[:index], s.allMessages[index+1:]...)
			s.allMessages = append(s.allMessages, updatedMessage)
			return updatedMessage, nil
		}
	}

	return models.Message{}, fmt.Errorf("Message could not be found")
}

var MessageService = messageService{
	allMessages: models.Messages{},
}
