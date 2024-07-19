package repository

// type Crud interface {
// 	Create(value any) error
// }

type Repository[T any] interface {
	Create(value T)
	FindAll() []T
	FindById(id int) (*T, error)
}
