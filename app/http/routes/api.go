package routes

import (
	"GoGinStarter/app/http/handlers/api"
	"GoGinStarter/app/http/middlewares"
	"GoGinStarter/internal/container"
	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(router *gin.Engine, container *container.Container) {
	OTPApiHandler := api.NewTOPApiHandler(container)
	userApiHandler := api.NewUserApiHandler(container)

	jwtMiddleware := middlewares.JWT(container.Response, container.Config, container.UserService)

	apiRouter := router.Group("/api/v1", middlewares.RequestLogger(container.Log))

	authRouter := apiRouter.Group("/auth")
	authRouter.POST("/sent-otp", OTPApiHandler.SentHandler)
	authRouter.POST("/verify-otp", OTPApiHandler.VerifyHandler)

	userRouter := apiRouter.Group("/users", jwtMiddleware)
	userRouter.GET("/index", userApiHandler.UsersHandler)
	userRouter.POST("/store", userApiHandler.StoreHandler)
	userRouter.GET("/:id/single", userApiHandler.SingleHandler)
}
