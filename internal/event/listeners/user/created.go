package user

import (
	"fmt"
)

type CreatedListener struct{}

func (l CreatedListener) HandleEvent(data interface{}) {
	fmt.Println("User created:", data)
}
