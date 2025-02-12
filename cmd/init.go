package main

import (
	"log"
	"pokemon_solid/internal/handler"
	"pokemon_solid/internal/repository"
	"pokemon_solid/internal/usecase"
)

func initHandler() (handler.UserHandler, handler.PokemonHandler, handler.CollectionsHandler) {
	log.Println("init user")
	ur := repository.NewUserRepository()
	uc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uc)
	generateUsers(uh)

	log.Println("init pokemon")
	pr := repository.NewPokemonRepository()
	pc := usecase.NewPokemonUsecase(pr)
	ph := handler.NewPokemonHandler(pc)
	generatePokemons(ph)

	log.Println("init collections")
	cr := repository.NewCollectionsRepository()
	cc := usecase.NewCollectionsUsecase(cr, ur, pr)
	hc := handler.NewCollectionsHandler(cc)

	return uh, ph, hc
}
