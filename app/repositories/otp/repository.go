package otp

import (
	"GoGinStarter/app/models"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/otp"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	time "time"
)

type Repository interface {
	Create(ctx *gin.Context, mobile string) (*models.OtpTokens, error)
	VerifyByToken(ctx *gin.Context, mobile string, token string) error
}

type repository struct {
	db     *gorm.DB
	log    log.Log
	config *config.Config
}

func (r *repository) Create(ctx *gin.Context, mobile string) (*models.OtpTokens, error) {
	otpToken := models.OtpTokens{
		Mobile:    mobile,
		Token:     otp.GenerateOTP(r.config.Auth.OTP.TokenLength),
		SentAt:    time.Now(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(r.config.Auth.OTP.ExpirationTime)),
	}
	r.db.Create(&otpToken)

	return &otpToken, nil
}

func (r *repository) VerifyByToken(ctx *gin.Context, mobile string, token string) error {
	otpToken := models.OtpTokens{}
	db := r.db.Where("mobile = ? AND token = ?", mobile, token).First(&otpToken)
	if err := db.Error; err != nil {
		return err
	}

	if time.Now().After(otpToken.ExpiresAt) {
		return fmt.Errorf("token has exipred")
	}

	if err := db.Delete(&otpToken).Error; err != nil {
		return err
	}

	return nil
}

func ProvideOtpRepository(db *gorm.DB, log log.Log, config *config.Config) Repository {
	return &repository{db, log, config}
}
