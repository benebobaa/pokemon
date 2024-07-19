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
	FindAll()
	Register(user domain.User)
	Access(username string) bool
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

func (h UserHandlerImpl) Access(username string) bool {

	_, err := h.UserUsecase.FindByUsername(username)

	if err != nil {
		fmt.Println("Access error: ", err.Error())
		return false
	}

	fmt.Println("Success access: ", username)

	return true
}
