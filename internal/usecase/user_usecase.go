package usecase

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/repository"
	"pokemon_solid/internal/util"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

type UserUsecase interface {
	CreateNew(user domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindAll() []domain.User
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return UserUsecaseImpl{
		UserRepository: userRepository,
	}
}

func (u UserUsecaseImpl) CreateNew(user domain.User) error {

	userExists, _ := u.UserRepository.FindByUsername(user.Username)

	if userExists != nil {
		return util.ErrUsernameAlreadyExists
	}

	u.UserRepository.Create(user)

	return nil
}

func (u UserUsecaseImpl) FindByUsername(username string) (*domain.User, error) {
	return u.UserRepository.FindByUsername(username)
}

func (u UserUsecaseImpl) FindAll() []domain.User {
	return u.UserRepository.FindAll()
}
