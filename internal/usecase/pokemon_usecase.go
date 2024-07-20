package usecase

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/repository"
)

type PokemonUsecaseImpl struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonUsecase interface {
	Usecase[domain.Pokemon]
}

func NewPokemonUsecase(pokemonRepository repository.PokemonRepository) PokemonUsecase {
	return PokemonUsecaseImpl{
		PokemonRepository: pokemonRepository,
	}
}

func (p PokemonUsecaseImpl) CreateNew(pokemon domain.Pokemon) error {
	p.PokemonRepository.Create(pokemon)
	return nil
}

func (p PokemonUsecaseImpl) FindAll() []domain.Pokemon {
	return p.PokemonRepository.FindAll()
}
