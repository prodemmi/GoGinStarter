//go:build wireinject
// +build wireinject

package wire

import (
	ur "GoGinStarter/app/repositories/user"
	us "GoGinStarter/app/services/user"
	"GoGinStarter/internal/cache"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/db"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/response"
	"GoGinStarter/internal/seeder"
	"GoGinStarter/internal/session"
	"github.com/google/wire"
)

func InitializeContainer() *container.Container {
	wire.Build(
		event.ProvideDispatcher,
		response.ProvideResponse,
		config.ProvideConfig,
		log.ProvideLog,
		cache.ProvideCache,
		db.ProvideDB,
		session.ProvideSession,
		seeder.ProvideSeeder,
		ur.ProvideUserRepository,
		us.ProvideUserService,
		container.ProvideContainer,
	)

	return &container.Container{}
}
