package repository

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/util"
)

type UserRepositoryImpl struct {
	Users  []domain.User
	lastID int
}

type UserRepository interface {
	Repository[domain.User]
	FindByUsername(username string) (*domain.User, error)
	// FindAll() []domain.User
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		Users: []domain.User{},
	}
}

func (r *UserRepositoryImpl) Create(value domain.User) {
	r.lastID++

	value.ID = r.lastID

	r.Users = append(r.Users, value)

}

func (r UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {

	for _, v := range r.Users {
		if v.Username == username {
			return &v, nil
		}
	}

	return nil, util.ErrUsernameNotFound
}

func (r UserRepositoryImpl) FindAll() []domain.User {
	return r.Users
}
