package usecase

import (
	"fmt"
	"math/rand"
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/domain/request"
	"pokemon_solid/internal/repository"
	"pokemon_solid/internal/util"
	"time"
)

type CollectionsUsecaseImpl struct {
	CollectionsRepository repository.CollectionsRepository
	UserRepository        repository.UserRepository
	PokemonRepository     repository.PokemonRepository
}

type CollectionsUsecase interface {
	FindAll() []domain.Collections
	TryCatch(request request.CatchRequest) (string, error)
	ReleaseCollection(request request.ReleaseRequest) error
	FindAllByUserID(id int) []domain.Collections
	catchProbability(rate int) (string, error)
}

func NewCollectionsUsecase(
	collectionsRepository repository.CollectionsRepository,
	userRepository repository.UserRepository,
	pokemonRepository repository.PokemonRepository,
) CollectionsUsecase {
	return CollectionsUsecaseImpl{
		CollectionsRepository: collectionsRepository,
		UserRepository:        userRepository,
		PokemonRepository:     pokemonRepository,
	}
}

// CreateNew implements CollectionsUsecase.
func (c CollectionsUsecaseImpl) TryCatch(request request.CatchRequest) (string, error) {
	user, err := c.UserRepository.FindById(request.UserID)

	if err != nil {
		return "", util.ErrUserNotFound
	}

	pokemon, err := c.PokemonRepository.FindById(request.PokemonID)

	if err != nil {
		return "", util.ErrPokemonNotFound
	}

	collection := domain.Collections{
		User:    *user,
		Pokemon: *pokemon,
	}

	message, err := c.catchProbability(pokemon.CatchRate)

	if err != nil {
		return "", err
	}

	c.CollectionsRepository.Create(collection)

	return message, nil
}

// FindAll implements CollectionsUsecase.
func (c CollectionsUsecaseImpl) FindAll() []domain.Collections {
	return c.CollectionsRepository.FindAll()
}

// Release implements CollectionsUsecase.
func (c CollectionsUsecaseImpl) ReleaseCollection(request request.ReleaseRequest) error {

	collection, err := c.CollectionsRepository.FindById(request.CollectionID)

	if err != nil {
		return err
	}

	if collection.User.ID != request.UserID {
		return util.ErrForbiddenReleasePokemon
	}

	err = c.CollectionsRepository.Delete(collection.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindAllByUserID implements CollectionsUsecase.
func (c CollectionsUsecaseImpl) FindAllByUserID(id int) []domain.Collections {
	return c.CollectionsRepository.FindByUserID(id)
}

func (_ CollectionsUsecaseImpl) catchProbability(rate int) (string, error) {
	// Fake Loading
	fmt.Print("\nLoading catch")

	for i := 0; i < 5; i++ {
		fmt.Print(" . ")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("\n\n")

	if rate < 0 || rate > 100 {
		return "", util.ErrInvalidRate
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS, you caught it", nil
	}

	return "", util.ErrFailedCatch
}
