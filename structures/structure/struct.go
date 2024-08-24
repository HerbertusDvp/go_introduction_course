package structure

type Product struct {
	Id    uint     `json:"id"`
	Name  string   `jason:"name"`
	Type  Type     `jason:"type"`
	Count uint16   `jason:"count"`
	Price float64  `jason:"price"`
	Tag   []string `jason:"tag"`
}

type Type struct {
	ID          uint   `jason:"id"`
	Code        string `jason:"code"`
	Description string `jason:"description"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}
