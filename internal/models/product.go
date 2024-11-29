package models

type FullProduct struct {
	ProductId  ProductId              `json:"product_id"`
	Item       string                 `json:"item"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Product struct {
	ProductId  ProductId
	Item       string
	Attributes string
}

func (m *Product) ToFullProduct(attributes map[string]interface{}) FullProduct {
	return FullProduct{
		ProductId:  m.ProductId,
		Item:       m.Item,
		Attributes: attributes,
	}
}
