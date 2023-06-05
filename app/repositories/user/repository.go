package user

import (
	"GoGinStarter/app/models"
	"GoGinStarter/internal/log"
	"GoGinStarter/internal/paginator"
	"GoGinStarter/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	Index(ctx *gin.Context) (*paginator.Paginator, error)
	FindById(ctx *gin.Context, id int) (*models.User, error)
}

type repository struct {
	db  *gorm.DB
	log log.Log
}

func (r *repository) Index(ctx *gin.Context) (*paginator.Paginator, error) {
	var users []models.User

	perPage := utils.IStoI(ctx.Query("per_page"))
	page := utils.IStoI(ctx.Query("page"))

	r.db = r.db.Find(&users)
	pagination := paginator.New(r.db, perPage, page)
	if err := r.db.Error; err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return pagination.WithItems(users), nil
}

func (r *repository) FindById(ctx *gin.Context, id int) (*models.User, error) {
	var user models.User
	if err := r.db.Where(id).First(&user).Error; err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func ProvideUserRepository(db *gorm.DB, log log.Log) Repository {
	return &repository{db, log}
}
