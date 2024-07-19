package repository

import (
	"pokemon_solid/internal/domain"
)

type PokemonRepositoryImpl struct {
	Pokemons map[int]domain.Pokemon
	lastID   int
}

type PokemonRepository interface {
	Repository[domain.Pokemon]
}

func NewPokemonRepository() PokemonRepository {
	return &PokemonRepositoryImpl{
		Pokemons: make(map[int]domain.Pokemon),
	}
}

func (p *PokemonRepositoryImpl) Create(value domain.Pokemon) {
	p.lastID++
	value.ID = p.lastID

	p.Pokemons[p.lastID] = value
}

func (p *PokemonRepositoryImpl) FindAll() []domain.Pokemon {
	var pokemons []domain.Pokemon

	for _, v := range p.Pokemons {
		pokemons = append(pokemons, v)
	}

	return pokemons
}
