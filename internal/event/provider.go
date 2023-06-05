package event

import (
	"GoGinStarter/internal/event/listeners/user"
)

var Listeners = map[string][]Listener{
	"user.created": {
		user.CreatedListener{},
	},
}
