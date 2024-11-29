package product_usecase

import (
	"graphql/internal/models"
	product_mng_repo "graphql/internal/product/product_repository/mongo"
	"graphql/internal/product/product_repository/postgres"
)

type CreateProduct struct {
	Item       string
	Attributes map[string]interface{}
}

func (r *CreateProduct) toCreateProduct(attributeId string) *product_psql_repo.CreateProduct {
	return &product_psql_repo.CreateProduct{
		Item:        r.Item,
		AttributeId: attributeId,
	}
}

type ProductsFilter struct {
	ProductIds []models.ProductId
	Items      []string
	Attributes map[string]interface{}
}

func (r *ProductsFilter) toGetProductsByFilter() *product_psql_repo.GetProductsByFilter {
	return &product_psql_repo.GetProductsByFilter{
		ProductIds: r.ProductIds,
		Items:      r.Items,
	}
}

func GetProductAttributes(products []models.Product) ([]string, map[string]*models.Product) {
	attributesIds := make([]string, len(products))
	attributesMap := make(map[string]*models.Product, len(products))

	for i, product := range products {
		attributesIds[i] = product.Attributes
		attributesMap[product.Attributes] = &product
	}

	return attributesIds, attributesMap
}

func (r *ProductsFilter) PrepareMngFilter(attributesIds []string) *product_mng_repo.AttributesFilter {
	return &product_mng_repo.AttributesFilter{
		AttributesIds: attributesIds,
		Attributes:    r.Attributes,
	}
}
