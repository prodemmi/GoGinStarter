package routes

import (
	"GoGinStarter/app/http/handlers/api"
	"GoGinStarter/app/http/middlewares"
	"GoGinStarter/internal/container"
	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(router *gin.Engine, container *container.Container) {
	userApiHandler := api.NewUserApiHandler(container)
	apiRouter := router.Group("/api/v1", middlewares.RequestLogger(container.Log))
	{
		apiRouter.GET("/users", userApiHandler.UsersHandler)
		apiRouter.GET("/users/:id", userApiHandler.SingleHandler)
		apiRouter.POST("/users", userApiHandler.StoreHandler)
	}
}
