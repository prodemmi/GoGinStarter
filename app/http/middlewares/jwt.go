package middlewares

import (
	"GoGinStarter/app/services/user"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func JWT(response response.Response, config *config.Config, userService user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			response.Unauthorized(ctx, "Unauthorized 401")
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.Auth.JWT.Secret), nil
		})

		if err != nil {
			response.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}

		if !token.Valid {
			response.Unauthorized(ctx, "Invalid token")
			ctx.Abort()
			return
		}

		user, _ := userService.FindByRememberToken(ctx, token.Raw)

		ctx.Set("remember_token", token)
		ctx.Set("user", user)
		ctx.Next()
	}
}
