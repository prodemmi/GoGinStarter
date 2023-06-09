package models

import (
	"gorm.io/gorm"
	"time"
)

type OtpTokens struct {
	gorm.Model
	Mobile    string    `json:"mobile" gorm:"index:not null"`
	Token     string    `json:"token" gorm:"size:6;not null"`
	SentAt    time.Time `json:"sent_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (o *OtpTokens) TableName() string {
	return "otp_tokens"
}
