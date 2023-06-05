package api

import (
	us "GoGinStarter/app/services/user"
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/response"
	"GoGinStarter/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserApiHandler struct {
	userService     us.Service
	response        response.Response
	log             log.Log
	eventDispatcher event.Dispatcher
}

func NewUserApiHandler(container *container.Container) *UserApiHandler {
	return &UserApiHandler{
		container.UserService,
		container.Response,
		container.Log,
		container.EventDispatcher,
	}
}

func (h *UserApiHandler) UsersHandler(ctx *gin.Context) {
	paginatedUsers, err := h.userService.Index(ctx)
	if err != nil {
		h.response.Error(ctx, 500, err.Error())
		return
	}
	h.response.WithPaginate(ctx, paginatedUsers, "")
}

func (h *UserApiHandler) SingleHandler(c *gin.Context) {
	user, err := h.userService.FindById(c, utils.IStoI(c.Param("id")))
	if err != nil {
		h.response.Error(c, 500, err.Error())
		return
	}

	h.response.Success(c, user, "")
}

func (h *UserApiHandler) StoreHandler(c *gin.Context) {
}
