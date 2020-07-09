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

	return nil
}