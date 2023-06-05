package user

import (
	"fmt"
)

type UpdatedListener struct{}

func (l UpdatedListener) HandleEvent(data interface{}) {
	fmt.Printf("User updated")
}
