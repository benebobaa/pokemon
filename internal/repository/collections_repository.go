package repository

import (
	"pokemon_solid/internal/domain"
	"time"
)

type CollectionsRepositoryImpl struct {
	Collections []domain.Collections
	lastId      int
}

type CollectionsRepository interface {
	Repository[domain.Collections]
}

func NewCollectionsRepository() CollectionsRepository {
	return &CollectionsRepositoryImpl{
		Collections: []domain.Collections{},
	}
}

func (c *CollectionsRepositoryImpl) Create(value domain.Collections) {
	c.lastId++
	value.ID = c.lastId
	value.CreatedAt = time.Now().Format(time.RFC1123)

	c.Collections = append(c.Collections, value)
}

func (c *CollectionsRepositoryImpl) FindAll() []domain.Collections {
	return c.Collections
}

// FindById implements CollectionsRepository.
func (c *CollectionsRepositoryImpl) FindById(id int) (*domain.Collections, error) {
	panic("unimplemented")
}
