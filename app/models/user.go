package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string         `json:"first_name" gorm:"not null"`
	LastName  string         `json:"last_name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique_index;not null"`
	Password  string         `json:"-" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
