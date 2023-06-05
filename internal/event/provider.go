package event

import (
	"GoGinStarter/app/listeners/user"
)

var Listeners = map[string][]Listener{
	"user.created": {
		user.CreatedListener{},
	},
}
