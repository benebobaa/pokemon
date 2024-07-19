package usecase

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/repository"
)

type PokemonUsecaseImpl struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonUsecase interface {
	CreateNew(pokemon domain.Pokemon)
	FindAll() []domain.Pokemon
}

func NewPokemonUsecase(pokemonRepository repository.PokemonRepository) PokemonUsecase {
	return PokemonUsecaseImpl{
		PokemonRepository: pokemonRepository,
	}
}

func (p PokemonUsecaseImpl) CreateNew(pokemon domain.Pokemon) {
	p.PokemonRepository.Create(pokemon)
}

func (p PokemonUsecaseImpl) FindAll() []domain.Pokemon {
	return p.PokemonRepository.FindAll()
}
