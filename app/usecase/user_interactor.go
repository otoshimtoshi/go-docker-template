package usecase

import (
	"go-docker-template/app/constants"
	"go-docker-template/app/domain"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) GetAll() (users []domain.UserResponse, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Usersの取得
	foundUsers, _ := interactor.User.Find(db)
	for _, v := range foundUsers {
		users = append(users, v.ToUserResponse())
	}

	return users, NewResultStatus(constants.StatusOK, nil)
}

func (interactor *UserInteractor) Get(id int) (user domain.UserResponse, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Userの取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UserResponse{
			ID:        0,
			Username:  "",
			Email:     "",
			Birthday:  "",
			ImageUrl:  "",
			CreatedAt: "",
			UpdatedAt: "",
		}, NewResultStatus(constants.StatusNotFound, err)
	}

	user = foundUser.ToUserResponse()

	return user, NewResultStatus(constants.StatusOK, nil)
}

//nolint:lll
func (interactor *UserInteractor) Create(user domain.User) (createdUser domain.UserResponse, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Userの登録
	err := interactor.User.Create(db, &user)
	if err != nil {
		return domain.UserResponse{
			ID:        0,
			Username:  "",
			Email:     "",
			Birthday:  "",
			ImageUrl:  "",
			CreatedAt: "",
			UpdatedAt: "",
		}, NewResultStatus(constants.StatusInternalServerError, err)
	}

	createdUser = user.ToUserResponse()

	return createdUser, NewResultStatus(constants.StatusOK, nil)
}
