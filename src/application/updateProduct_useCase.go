package application

import "demob/src/domain"

type UpdateProductUseCase struct {
	db domain.IProduct
}

func NewUpdateProductUseCase (db domain.IProduct) *UpdateProductUseCase {
	return &UpdateProductUseCase{db: db}
}

func (uc *UpdateProductUseCase) Run(id int32, product domain.Product) error {
	return uc.db.Update(id, product)
}

