package domain

type IProduct interface {
	Save(product Product) error
	GetAll() ([]Product, error)
	GetByID(id int32) (*Product, error)
}
