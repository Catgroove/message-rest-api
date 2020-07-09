package models

import (
	"time"
	"errors"
)

type Message struct {
	ID 			int `json:"id"`
	Message 	string `json:"message"`
	Created 	time.Time `json:"created"`
	Updated 	time.Time `json:"updated"`
}

type Messages []Message

func (m *Message) Validate() error {
	if m.Message == "" {
		return errors.New("Message is required")
	}

	if len(m.Message) > 200 {
		return errors.New("Message length greater than 200 characters")
	}

	return nil
}