package main

import (
	"log"
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/handler"
	"time"
)

func generateUsers(userHandler handler.UserHandler) {
	log.Println("generate users")

	user1 := domain.User{
		Username: "beneboba",
		Name:     "bene",
	}

	user3 := domain.User{
		Username: "beneboba",
		Name:     "siapa saya",
	}

	user2 := domain.User{
		Username: "phinconers",
		Name:     "phincon",
	}

	userHandler.Register(user1)
	userHandler.Register(user2)
	userHandler.Register(user3)
}

func generatePokemons(ph handler.PokemonHandler) {
	log.Println("generate pokemons")
	now := time.Now()

	pokemons := []domain.Pokemon{
		{Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now.Format(time.RFC1123)},
		{Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0).Format(time.RFC1123)},
		{Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0).Format(time.RFC1123)},
		{Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0).Format(time.RFC1123)},
		{Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0).Format(time.RFC1123)},
	}

	for _, v := range pokemons {
		ph.Create(v)
	}
}
