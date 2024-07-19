package repository

import (
	"pokemon_solid/internal/domain"
	intf "pokemon_solid/internal/interface"
)

type UserRepositoryImpl struct {
	DB map[int]domain.User
}

type UserRepository interface {
	intf.Crud
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{
		DB: make(map[int]domain.User),
	}
}

func (r UserRepositoryImpl) Create() {
}

func (r UserRepositoryImpl) Read() {
}

func (r UserRepositoryImpl) Update() {

}
func (r UserRepositoryImpl) Delete() {

}
