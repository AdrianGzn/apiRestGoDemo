package infraestructure
import "fmt"
import "demob/src/domain"
import "errors"

type Mysql struct {
	products []domain.Product
}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (mysql *Mysql) Save(product domain.Product) error {
	mysql.products = append(mysql.products, product)
	fmt.Println("Product saved:", product)
	return nil
} 

func (mysql *Mysql) GetAll() ([]domain.Product, error) {
	return mysql.products, nil
}

func (mysql *Mysql) GetByID(id int32) (*domain.Product, error) {
	for _, myProduct := range mysql.products {
		if myProduct.GetId() == id {
			return &myProduct, nil
		}
	}
	
	return nil, errors.New("Producto no encontrado")
}

func (mysql *Mysql) Update(id int32, updatedProduct domain.Product) error {
	for i, myProduct := range mysql.products {
		if myProduct.GetId() == id {
			mysql.products[i] = updatedProduct
			fmt.Println("Product: ", updatedProduct)
			return nil
		}
	}

	return errors.New("Producto no encontrado")
}

func (mysql *Mysql) Delete(id int32) error {
	for i, myProduct := range mysql.products {
		if myProduct.GetId() == id {
			mysql.products = append(mysql.products[:i], mysql.products[i+1:]...)
			fmt.Println("Producto eliminado: ", id)
			return nil
		}
	}
	
	return errors.New("Producto no encontrado")
}