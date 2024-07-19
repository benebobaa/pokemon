package usecase

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/repository"
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

	userExists, err := u.UserRepository.FindByUsername(user.Username)

	if userExists != nil {
		return err
	}

	err = u.UserRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecaseImpl) FindByUsername(username string) (*domain.User, error) {
	return u.UserRepository.FindByUsername(username)
}

func (u UserUsecaseImpl) FindAll() []domain.User {
	return u.UserRepository.FindAll()
}
