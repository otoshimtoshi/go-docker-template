package usecase

import (
	"go-docker-template/app/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Find(db *gorm.DB) (users []domain.User, err error)
	FindByID(db *gorm.DB, id int) (user domain.User, err error)
	Create(db *gorm.DB, user *domain.User) (err error)
}
