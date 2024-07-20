package usecase

type Usecase[T any] interface {
	FindAll() []T
	CreateNew(any T) error
}
