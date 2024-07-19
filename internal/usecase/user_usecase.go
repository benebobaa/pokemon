package usecase

import "pokemon_solid/internal/repository"

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

type UserUsecase interface {
	CreateNew()
	FindById()
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return UserUsecaseImpl{
		UserRepository: userRepository,
	}
}

func (u UserUsecaseImpl) CreateNew() {

}

func (u UserUsecaseImpl) FindById() {

}
