package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string         `json:"first_name" gorm:"not null"`
	LastName      string         `json:"last_name" gorm:"not null"`
	Mobile        string         `json:"mobile;" gorm:"unique_index"`
	Email         string         `json:"email" gorm:"unique_index;not null"`
	Password      string         `json:"-" gorm:"not null"`
	RememberToken string         `json:"-" gorm:"unique"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
