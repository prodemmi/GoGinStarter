package notification

import (
	"fmt"
	"time"
)

type SMS struct {
	SendFrom string
	SendTo   []string
	Delay    time.Duration
}

func (s *SMS) Send() (*SMS, error) {
	// Implement sending sms
	fmt.Printf("Sending SMS notification to %s:\n", s.SendFrom)
	return s, nil
}
