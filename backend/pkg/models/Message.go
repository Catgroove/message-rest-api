package models

import (
	"errors"
	"math/rand"
	"time"
)

type Message struct {
	ID      int       `json:"id"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (m *Message) BeforeCreate() {
	m.ID = rand.Intn(100000000)
	m.Created = time.Now()
	m.Updated = time.Now()
}

func (m *Message) BeforeUpdate() {
	m.Updated = time.Now()
}

func (m *Message) Validate() error {
	if m.Message == "" {
		return errors.New("Message is required")
	}

	if len(m.Message) > 200 {
		return errors.New("Message length greater than 200 characters")
	}

	return nil
}
