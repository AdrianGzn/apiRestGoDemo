package application
import "demob/src/domain"

type ViewAllProductUseCase struct {
	db domain.IProduct
}

func NewViewAllUseCase(db domain.IProduct) *ViewAllProductUseCase {
	return &ViewAllProductUseCase{db: db}
}

func (uc *ViewAllProductUseCase) Run() ([]domain.Product, error) {
	return uc.db.GetAll()
}
