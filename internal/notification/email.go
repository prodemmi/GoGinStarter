package notification

import (
	"fmt"
	"time"
)

type Email struct {
	SendFrom string
	SendTo   []string
	Subject  string
	Message  string
	Delay    time.Duration
}

func (s *Email) Send(delay time.Duration) (*Email, error) {
	// Implement sending Email
	fmt.Printf("Sending Email to %s:\n", s.SendFrom)
	return s, nil
}
