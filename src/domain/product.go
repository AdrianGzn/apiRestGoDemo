package domain

type Product struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Stock     int32   `json:"stock"`
	CreatedAt string  `json:"created_at"`
}

func (p *Product) GetName() string        { return p.Name }
func (p *Product) GetPrice() float32      { return p.Price }
func (p *Product) GetStock() int32        { return p.Stock }
func (p *Product) GetCreatedAt() string   { return p.CreatedAt }
func (p *Product) GetId() int32           { return p.ID }
