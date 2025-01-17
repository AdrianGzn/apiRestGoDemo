package application

import "demob/src/domain"

type DeleteProductUseCase struct {
	db domain.IProduct
}

func NewDeleteProductUseCase(db domain.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{db: db}
}

func (uc *DeleteProductUseCase) Run(id int32) error {
	return uc.db.Delete(id)
}