package container

import (
	otpr "GoGinStarter/app/repositories/otp"
	ur "GoGinStarter/app/repositories/user"
	otps "GoGinStarter/app/services/otp"
	us "GoGinStarter/app/services/user"
	"GoGinStarter/internal/cache"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/response"
	"GoGinStarter/internal/scheduler"
	"GoGinStarter/internal/seeder"
	"GoGinStarter/internal/session"
	"gorm.io/gorm"
)

type Container struct {
	OTPService      otps.Service
	OTPRepository   otpr.Repository
	UserService     us.Service
	UserRepository  ur.Repository
	Cache           cache.Cache
	DB              *gorm.DB
	Config          *config.Config
	Response        response.Response
	Log             log.Log
	EventDispatcher event.Dispatcher
	Seeder          seeder.Seeder
	Session         session.Session
	Scheduler       scheduler.Schedule
}

func ProvideContainer(
	OTPService otps.Service,
	OTPRepository otpr.Repository,
	UserService us.Service,
	UserRepository ur.Repository,
	Cache cache.Cache,
	Config *config.Config,
	DB *gorm.DB,
	Response response.Response,
	Log log.Log,
	EventDispatcher event.Dispatcher,
	Seeder seeder.Seeder,
	Session session.Session,
	Scheduler scheduler.Schedule,
) *Container {
	return &Container{
		OTPService:      OTPService,
		OTPRepository:   OTPRepository,
		UserService:     UserService,
		UserRepository:  UserRepository,
		Cache:           Cache,
		Config:          Config,
		DB:              DB,
		Response:        Response,
		Log:             Log,
		EventDispatcher: EventDispatcher,
		Seeder:          Seeder,
		Session:         Session,
		Scheduler:       Scheduler,
	}
}
