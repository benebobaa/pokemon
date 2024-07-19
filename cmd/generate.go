package main

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/handler"
)

func generateUsers(userHandler handler.UserHandler) {

	user1 := domain.User{
		Username: "beneboba",
		Name:     "bene",
	}

	user2 := domain.User{
		Username: "phinconers",
		Name:     "phincon",
	}

	userHandler.Register(user1)
	userHandler.Register(user2)
}
