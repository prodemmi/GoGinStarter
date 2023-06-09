package event

import (
	"GoGinStarter/app/listeners/otp"
	"GoGinStarter/app/listeners/user"
)

var Listeners = map[string][]Listener{
	"user.created": {
		user.CreatedListener{},
	},
	"otp.created": {
		otp.CreatedListener{},
	},
}
