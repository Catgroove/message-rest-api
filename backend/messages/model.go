package messages

import "time"

type Message struct {
	ID 			int `json:"id"`
	Message 	string `json:"message"`
	Created 	time.Time `json:"created"`
	Updated 	time.Time `json:"updated"`
}

type Messages []Message
