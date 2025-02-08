package controllers

import (
	"log"
	"strconv"
	"time"

	"go-docker-template/app/constants"
	"go-docker-template/app/domain"
	"go-docker-template/app/interfaces/database"
	"go-docker-template/app/usecase"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

type StoreJsonRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	ImageUrl string `json:"image_url"`
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *UsersController) Index(c Context) {
	users, res := controller.Interactor.GetAll()
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))

		return
	}

	c.JSON(res.StatusCode, NewH("success", users))
}

func (controller *UsersController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))

		return
	}

	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UsersController) Store(c Context) {
	var json StoreJsonRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(constants.StatusBadRequest, NewH(err.Error(), nil))

		return
	}

	birthday, err := time.Parse("2006-01-02", json.Birthday)
	if err != nil {
		log.Println("Invalid date format:", err)
		birthday = time.Time{} // デフォルト値を設定
	}

	user := domain.User{
		ID:        0,
		Username:  json.Username,
		Password:  json.Password,
		Email:     json.Email,
		Birthday:  birthday,
		ImageUrl:  json.ImageUrl,
		CreatedAt: time.Now(), // ここで適切な値を設定
		UpdatedAt: time.Now(), // ここで適切な値を設定
	}

	createdUser, res := controller.Interactor.Create(user)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))

		return
	}

	c.JSON(res.StatusCode, NewH("success", createdUser))
}
