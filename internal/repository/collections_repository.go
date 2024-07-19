package repository

import (
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/util"
	"time"
)

type CollectionsRepositoryImpl struct {
	Collections []domain.Collections
	lastId      int
}

type CollectionsRepository interface {
	Repository[domain.Collections]
	Delete(id int) error
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

	for _, v := range c.Collections {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, util.ErrCollectionNotFound
}

// Delete implements CollectionsRepository.
func (c *CollectionsRepositoryImpl) Delete(id int) error {

	for i, v := range c.Collections {
		if id == v.ID {
			c.Collections = append(c.Collections[:i], c.Collections[i+1:]...)
			return nil
		}
	}

	return util.ErrCollectionNotFound
}
