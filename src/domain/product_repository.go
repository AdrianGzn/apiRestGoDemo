package domain

type IProduct interface {
	Save(product Product) error
	GetAll() ([]Product, error)
	GetByID(id int32) (*Product, error)
	Update(id int32, product Product) error
	Delete(id int32) error
}
