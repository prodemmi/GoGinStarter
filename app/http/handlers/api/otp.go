package api

import (
	"GoGinStarter/app/services/otp"
	"GoGinStarter/app/services/user"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/event"
	"GoGinStarter/internal/jwt"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/response"
	"github.com/gin-gonic/gin"
)

type OTPApiHandler struct {
	userService     user.Service
	otpService      otp.Service
	response        response.Response
	log             log.Log
	config          *config.Config
	eventDispatcher event.Dispatcher
}

func NewTOPApiHandler(container *container.Container) *OTPApiHandler {
	return &OTPApiHandler{
		container.UserService,
		container.OTPService,
		container.Response,
		container.Log,
		container.Config,
		container.EventDispatcher,
	}
}

type SentRequest struct {
	Mobile string `json:"mobile"`
}

func (o *OTPApiHandler) SentHandler(ctx *gin.Context) {
	var data SentRequest
	if err := ctx.BindJSON(&data); err != nil {
		o.response.BadRequest(ctx, err.Error(), nil)
		return
	}

	_, err := o.otpService.Create(ctx, data.Mobile)
	if err != nil {
		o.response.Error(ctx, 500, err.Error())
		return
	}

	o.response.Success(ctx, nil, "OTP successfully sent")
}

type VerifyRequest struct {
	Mobile string `json:"mobile"`
	Token  string `json:"token"`
}

func (o *OTPApiHandler) VerifyHandler(ctx *gin.Context) {
	var data VerifyRequest
	if err := ctx.BindJSON(&data); err != nil {
		o.response.BadRequest(ctx, err.Error(), nil)
		return
	}

	if err := o.otpService.VerifyByToken(ctx, data.Mobile, data.Token); err != nil {
		o.response.Error(ctx, 500, err.Error())
		return
	}

	user, err := o.userService.FirstOrCreate(ctx, data.Mobile)
	if err != nil {
		o.response.Error(ctx, 500, err.Error())
		return
	}

	tokenString, err := jwt.CreateToken(user.ID, o.config.Auth.JWT.Secret, o.config.Auth.JWT.ExpirationTime)
	if err != nil {
		o.response.Error(ctx, 500, err.Error())
		return
	}

	user, err = o.userService.AddRememberToken(ctx, user.ID, tokenString)
	if err != nil {
		o.response.Error(ctx, 500, err.Error())
		return
	}

	ctx.Header("Authorization", "Bearer "+tokenString)

	o.response.Success(ctx, nil, "Authentication successfully")
}
