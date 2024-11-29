package product_usecase

import (
	"context"
	"graphql/internal/models"
	product_mng_repo "graphql/internal/product/product_repository/mongo"
	product_psql_repo "graphql/internal/product/product_repository/postgres"
)

type ProductMngRepo interface {
	AddProductAttributes(ctx context.Context, attributes map[string]interface{}) (string, error)
	GetProductAttributesByFilter(ctx context.Context, attributesFilter *product_mng_repo.AttributesFilter) ([]product_mng_repo.ProductAttributes, error)
}

type ProductPsqlRepo interface {
	CreateProduct(ctx context.Context, productParams *product_psql_repo.CreateProduct) (models.ProductId, error)
	GetProductsByFilter(ctx context.Context, productFilter *product_psql_repo.GetProductsByFilter) ([]models.Product, error)
}
