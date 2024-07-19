package handler

import "pokemon_solid/internal/usecase"

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) {

}
