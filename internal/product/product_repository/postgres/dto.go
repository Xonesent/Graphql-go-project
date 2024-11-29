package product_psql_repo

import "graphql/internal/models"

type CreateProduct struct {
	Item        string
	AttributeId string
}

type GetProductsByFilter struct {
	ProductIds []models.ProductId
	Items      []string
}
