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
