package seeds

import (
	"time"

	"go-docker-template/app/domain"

	"github.com/jinzhu/gorm"
)

//nolint:lll
func CreateUser(db *gorm.DB, username string, password string, email string, birthday time.Time, imageURL string) error {
	return db.Create(&domain.User{
		ID:        0,
		Username:  username,
		Password:  password,
		Email:     email,
		Birthday:  birthday,
		ImageUrl:  imageURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}
