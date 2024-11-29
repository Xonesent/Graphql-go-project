package product_http

import (
	"graphql/internal/models"
	"graphql/internal/product/product_usecase"
)

type createProductRequest struct {
	Item       string                 `json:"item" validate:"required"`
	Attributes map[string]interface{} `json:"attributes"`
}

func (r *createProductRequest) toCreateProduct() *product_usecase.CreateProduct {
	return &product_usecase.CreateProduct{
		Item:       r.Item,
		Attributes: r.Attributes,
	}
}

type productFilterRequest struct {
	ProductIds []models.ProductId     `json:"product_ids"`
	Items      []string               `json:"items"`
	Attributes map[string]interface{} `json:"attributes"`
}

func (r *productFilterRequest) toCreateProduct() *product_usecase.ProductsFilter {
	return &product_usecase.ProductsFilter{
		ProductIds: r.ProductIds,
		Items:      r.Items,
		Attributes: r.Attributes,
	}
}
