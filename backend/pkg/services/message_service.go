package services

import (
	"time"
	"math/rand"

	"backend/pkg/models"
)

type messageService struct {
	allMessages models.Messages
}

func (s *messageService) GetAllMessages () models.Messages {
	return s.allMessages
}

func (s *messageService) GetMessage (id int) models.Message {
	for _, m := range s.allMessages {
		if m.ID == id {
			return m
		}
	}

	return models.Message{}
}

func (s *messageService) CreateMessage(m models.Message) models.Message {
	newMessage := models.Message{
		ID: rand.Intn(100000000),
		Message: m.Message,
		Created: time.Now(),
		Updated: time.Now(),
	}
	s.allMessages = append(s.allMessages, newMessage);
	return newMessage
}

func (s *messageService) DeleteMessage(id int) {
	for index, m := range s.allMessages {
		if m.ID == id {
			s.allMessages = append(s.allMessages[:index], s.allMessages[index+1:]...)
			break
		}
	}
}

func (s *messageService) UpdateMessage(id int, updatedMessage models.Message) models.Message {
	for index, m := range s.allMessages {
		if m.ID == id {
			message := s.allMessages[index]
			message.Message = updatedMessage.Message
			message.Updated = time.Now()

			s.allMessages = append(s.allMessages[:index], s.allMessages[index+1:]...)			
			s.allMessages = append(s.allMessages, message)
			return message
		}
	}

	return models.Message{}
}

var MessageService = messageService{}
