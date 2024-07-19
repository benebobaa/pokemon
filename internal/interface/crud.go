package intf

type Crud interface {
	Create(value any) error
}
