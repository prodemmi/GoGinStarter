package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(userID uint, secret string, expirationAt int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Second * time.Duration(expirationAt)).Unix(),
	})
	return token.SignedString([]byte(secret))
}
