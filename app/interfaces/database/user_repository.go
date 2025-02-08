package database

import (
	"errors"
	"time"

	"go-docker-template/app/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) Find(db *gorm.DB) (users []domain.User, err error) {
	users = []domain.User{}
	db.Find(&users)

	return users, nil
}

var ErrUserNotFound = errors.New("user is not found")

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{
		ID:        0,
		Username:  "",
		Password:  "",
		Email:     "",
		Birthday:  time.Time{},
		ImageUrl:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	db.First(&user, id)

	if user.ID <= 0 {
		return domain.User{}, ErrUserNotFound
	}

	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, user *domain.User) (err error) {
	result := db.Create(&user)

	return result.Error
}
