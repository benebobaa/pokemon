package domain

import "time"

type Pokemon struct {
	ID             int
	Name           string
	Type           string
	CatchRate      int
	IsRare         bool
	RegisteredDate time.Time
}
