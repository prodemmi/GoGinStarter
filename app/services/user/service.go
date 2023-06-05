package user

import (
	"GoGinStarter/app/models"
	"GoGinStarter/app/repositories/user"
	"GoGinStarter/internal/cache"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/paginator"
	"GoGinStarter/internal/utils"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repository      user.Repository
	log             log.Log
	cache           cache.Cache
	eventDispatcher event.Dispatcher
}

func (s Service) Index(ctx *gin.Context) (*paginator.Paginator, error) {
	paginatedUsers, err := s.repository.Index(ctx)
	if err != nil {
		return nil, err
	}
	s.eventDispatcher.Dispatch("user.created", nil)

	return paginatedUsers, nil
}

func (s Service) FindById(ctx *gin.Context, id int) (*models.User, error) {
	return s.repository.FindById(ctx, utils.IStoI(ctx.Param("id")))
}

func ProvideUserService(
	Repository user.Repository,
	Log log.Log,
	Cache cache.Cache,
	EventDispatcher event.Dispatcher,
) Service {
	return Service{
		repository:      Repository,
		log:             Log,
		cache:           Cache,
		eventDispatcher: EventDispatcher,
	}
}
