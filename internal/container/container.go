package container

import (
	ur "GoGinStarter/app/repositories/user"
	us "GoGinStarter/app/services/user"
	"GoGinStarter/internal/cache"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/response"
	"gorm.io/gorm"
)

type Container struct {
	UserService     us.Service
	UserRepository  ur.Repository
	Cache           cache.Cache
	DB              *gorm.DB
	Config          *config.Config
	Response        response.Response
	Log             log.Log
	EventDispatcher event.Dispatcher
}

func ProvideContainer(
	UserService us.Service,
	UserRepository ur.Repository,
	Cache cache.Cache,
	Config *config.Config,
	DB *gorm.DB,
	Response response.Response,
	Log log.Log,
	EventDispatcher event.Dispatcher,
) *Container {
	return &Container{
		UserService:     UserService,
		UserRepository:  UserRepository,
		Cache:           Cache,
		Config:          Config,
		DB:              DB,
		Response:        Response,
		Log:             Log,
		EventDispatcher: EventDispatcher,
	}
}
