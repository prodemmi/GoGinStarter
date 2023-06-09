package otp

import (
	"GoGinStarter/app/repositories/otp"
	"GoGinStarter/internal/cache"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repository      otp.Repository
	log             log.Log
	cache           cache.Cache
	eventDispatcher event.Dispatcher
}

func (s Service) Create(ctx *gin.Context, mobile string) (string, error) {
	otpToken, err := s.repository.Create(ctx, mobile)
	if err != nil {
		return "", err
	}
	s.eventDispatcher.Dispatch("otp.created", map[string]any{
		"token":  otpToken.Token,
		"mobile": otpToken.Mobile,
	})

	return otpToken.Token, nil
}

func (s Service) VerifyByToken(ctx *gin.Context, mobile string, token string) error {
	return s.repository.VerifyByToken(ctx, mobile, token)
}

func ProvideOTPService(
	Repository otp.Repository,
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
