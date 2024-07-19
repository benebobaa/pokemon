package handler

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/usecase"
	"pokemon_solid/internal/util"
)

type PokemonHandlerImpl struct {
	PokemonUsecase usecase.PokemonUsecase
}

type PokemonHandler interface {
	FindAll()
	Create(pokemon domain.Pokemon)
}

func NewPokemonHandler(pokemonUsecase usecase.PokemonUsecase) PokemonHandler {
	return PokemonHandlerImpl{
		PokemonUsecase: pokemonUsecase,
	}
}

func (p PokemonHandlerImpl) Create(pokemon domain.Pokemon) {
	p.PokemonUsecase.CreateNew(pokemon)
}

func (p PokemonHandlerImpl) FindAll() {

	pokemons := p.PokemonUsecase.FindAll()

	util.TablePokemon(pokemons)
}
