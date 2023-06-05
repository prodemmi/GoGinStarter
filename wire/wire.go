//go:build wireinject
// +build wireinject

package wire

import (
	"GoGinStarter/app/internal/cache"
	"GoGinStarter/app/internal/config"
	"GoGinStarter/app/internal/container"
	"GoGinStarter/app/internal/db"
	"GoGinStarter/app/internal/event"
	"GoGinStarter/app/internal/log"
	"GoGinStarter/app/internal/response"
	ur "GoGinStarter/app/repositories/user"
	us "GoGinStarter/app/services/user"
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
		ur.ProvideUserRepository,
		us.ProvideUserService,
		container.ProvideContainer,
	)

	return &container.Container{}
}
