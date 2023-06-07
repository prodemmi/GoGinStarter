package notifications

import (
	"GoGinStarter/internal/notification"
	"time"
)

func SendExampleSMSNotification(delay time.Duration) (*notification.SMS, error) {
	notify := notification.SMS{SendFrom: "09130243519", SendTo: []string{"09130243519"}, Delay: delay}
	return notify.Send()
}
