package domain

type Product struct {
	id int32 `json:"id"`
	name string `json:"name"`
	price float32 `json:"price"`
}

func NewProduct (name string, price float32) *Product {
	//return &product{id: 1, name: name, price: price}
	
	object := Product{id: 1, name: name, price: price}
	return &object
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) GetId() int32 {
	return p.id
}