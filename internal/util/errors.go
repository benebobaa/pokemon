package util

import "errors"

var (
	ErrUsernameNotFound        = errors.New("Username not found")
	ErrUsernameAlreadyExists   = errors.New("Username already exists")
	ErrInvalidRate             = errors.New("Invalid rate. Please provide a rate between 0 and 100")
	ErrFailedCatch             = errors.New("FAIL, it got away")
	ErrUserNotFound            = errors.New("User not found")
	ErrPokemonNotFound         = errors.New("Pokemon not found")
	ErrForbiddenReleasePokemon = errors.New("Forbidden release pokemon, its not yours")
	ErrCollectionNotFound      = errors.New("Collection not found")
)
