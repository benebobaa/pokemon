package handler

import (
	"fmt"
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/usecase"
	"pokemon_solid/internal/util"
)

type UserHandlerImpl struct {
	UserUsecase usecase.UserUsecase
}

type UserHandler interface {
	Handler
	Register(user domain.User)
	Access(user *domain.User) bool
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandlerImpl{
		UserUsecase: userUsecase,
	}
}

func (h UserHandlerImpl) FindAll() {

	users := h.UserUsecase.FindAll()

	util.TableUser(users)
}

func (h UserHandlerImpl) Register(user domain.User) {

	err := h.UserUsecase.CreateNew(user)

	if err != nil {
		fmt.Println("Register error: ", err.Error())
	}
}

func (h UserHandlerImpl) Access(user *domain.User) bool {

	userExists, err := h.UserUsecase.FindByUsername(user.Username)

	if err != nil {
		fmt.Println("Access error: ", err.Error())
		return false
	}

	*user = *userExists

	fmt.Println("Success access: ", user.Username)

	return true
}
