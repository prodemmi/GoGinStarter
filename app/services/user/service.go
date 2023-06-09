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
	//s.eventDispatcher.Dispatch("user.created", nil)
	//_, _ = notifications.SendExampleSMSNotification(0)
	//_, _ = notifications.SendExampleEmailNotification(0)

	return paginatedUsers, nil
}

func (s Service) FindById(ctx *gin.Context, id int) (*models.User, error) {
	return s.repository.FindById(ctx, utils.IStoI(ctx.Param("id")))
}

func (s Service) FirstOrCreate(ctx *gin.Context, mobile string) (*models.User, error) {
	user, err := s.repository.FirstOrCreate(ctx, mobile)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) AddRememberToken(ctx *gin.Context, id uint, token string) (*models.User, error) {
	return s.repository.AddRememberToken(ctx, id, token)
}

func (s Service) FindByRememberToken(ctx *gin.Context, token string) (*models.User, error) {
	return s.repository.FindByRememberToken(ctx, token)
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
