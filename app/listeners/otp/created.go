package otp

import (
	"fmt"
)

type CreatedListener struct{}

func (l CreatedListener) HandleEvent(data interface{}) {
	if d, ok := data.(map[string]interface{}); ok {
		token := d["token"].(string)
		mobile := d["mobile"].(string)

		// Implement OTP send

		fmt.Println("OTP created is " + token + " for mobile " + mobile)
	}
}
