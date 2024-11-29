package product_delivery

import (
	"context"
	"graphql/internal/models"
	"graphql/internal/product/product_usecase"
)

type ProductUC interface {
	CreateProduct(ctx context.Context, createProductParams *product_usecase.CreateProduct) (models.ProductId, error)
	GetProductsByFilter(ctx context.Context, productParams *product_usecase.ProductsFilter) ([]models.FullProduct, error)
}
