package infraestructure

import (
	"database/sql"
	"demob/src/domain"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	DB *sql.DB
}


func NewMysql() *Mysql {
	dsn := "goUser:goPass123@tcp(54.91.181.99:3306)/demo_db"


	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Error al conectar a MySQL: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("MySQL no responde: " + err.Error())
	}

	fmt.Println("Conectado a MySQL correctamente ✔")

	return &Mysql{DB: db}
}


func (mysql *Mysql) Save(product domain.Product) error {

	query := "INSERT INTO products (name, price, stock, created_at) VALUES (?, ?, ?, ?)"

	_, err := mysql.DB.Exec(query,
		product.GetName(),
		product.GetPrice(),
		product.GetStock(),
		product.GetCreatedAt(),
	)

	if err != nil {
		return err
	}

	return nil
}


func (mysql *Mysql) GetAll() ([]domain.Product, error) {

	rows, err := mysql.DB.Query("SELECT id, name, price, stock, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var p domain.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}


func (mysql *Mysql) GetByID(id int32) (*domain.Product, error) {

	query := "SELECT id, name, price, stock, created_at FROM products WHERE id = ?"

	var p domain.Product

	err := mysql.DB.QueryRow(query, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("producto no encontrado")
	} else if err != nil {
		return nil, err
	}

	return &p, nil
}
