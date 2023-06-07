package notifications

import (
	"GoGinStarter/internal/notification"
	"time"
)

func SendExampleEmailNotification(delay time.Duration) (*notification.Email, error) {
	notify := notification.Email{
		SendFrom: "09130243519",
		SendTo:   []string{"09130243519"},
		Subject:  "Subject",
		Message:  "Message",
		Delay:    delay,
	}
	return notify.Send(0)
}
