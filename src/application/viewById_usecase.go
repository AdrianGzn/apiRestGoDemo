package application
import "demob/src/domain"

type ViewByIdProductUseCase struct {
	db domain.IProduct
}

func NewViewByIdProductUseCase(db domain.IProduct) *ViewByIdProductUseCase {
	return &ViewByIdProductUseCase{db: db}
}

func (uc *ViewByIdProductUseCase) Run (id int32) (*domain.Product, error) {
	return uc.db.GetByID(id)
}
