package messages

import (
	"time"
	"math/rand"
)

type messageService struct {
	allMessages Messages
}

func (s *messageService) GetAllMessages () Messages {
	return s.allMessages
}

func (s *messageService) GetMessage (id int) Message {
	for _, m := range s.allMessages {
		if m.ID == id {
			return m
		}
	}

	return Message{}
}

func (s *messageService) CreateMessage(m Message) Message {
	newMessage := Message{
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

func (s *messageService) UpdateMessage(id int, updatedMessage Message) Message {
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

	return Message{}
}

var MessageService = messageService{}
