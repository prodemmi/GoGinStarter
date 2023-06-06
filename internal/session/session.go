package session

import (
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"gorm.io/gorm"
)

type Session struct {
	db    *gorm.DB
	Store gormsessions.Store
}

func ProvideSession(db *gorm.DB) Session {
	store := gormsessions.NewStore(db, true, []byte("secret"))
	return Session{db, store}
}
