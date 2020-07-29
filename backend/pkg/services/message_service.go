package services

import (
	"fmt"

	"backend/pkg/models"
)

type Messages map[int]models.Message;

// In a real application, the service wouldn't store the data, and it wouldn't be in memory
type MessageService struct {
	allMessages Messages
}

func NewMessageService() *MessageService {
	return &MessageService{
		allMessages: make(Messages),
	}
}

func (s *MessageService) GetAllMessages() Messages {
	return s.allMessages
}

func (s *MessageService) GetMessage(id int) (models.Message, error) {
	if m, found := s.allMessages[id]; found {
		return m, nil
	}

	return models.Message{}, fmt.Errorf("Message could not be found")
}

func (s *MessageService) CreateMessage(m models.Message) models.Message {
	s.allMessages[m.ID] = m;

	return m
}

func (s *MessageService) DeleteMessage(id int) error {
	if _, found := s.allMessages[id]; found {
		delete(s.allMessages, id)
		return nil
	}

	return fmt.Errorf("Message could not be found")
}

func (s *MessageService) UpdateMessage(updatedMessage models.Message) (models.Message, error) {
	if _, found := s.allMessages[updatedMessage.ID]; found {
		s.allMessages[updatedMessage.ID] = updatedMessage; 
		return updatedMessage, nil
	}

	return models.Message{}, fmt.Errorf("Message could not be found")
}
