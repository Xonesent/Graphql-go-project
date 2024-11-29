//go:generate go run github.com/99designs/gqlgen generate
package product_graphql

import (
	"context"
	"graphql/config"
	"graphql/internal/models"
	"graphql/internal/product/product_delivery"
)

type ProductResolver struct {
	cfg       *config.Config
	ProductUC product_delivery.ProductUC
}

func NewProductResolver(cfg *config.Config, ProductUC product_delivery.ProductUC) *ProductResolver {
	return &ProductResolver{
		cfg:       cfg,
		ProductUC: ProductUC,
	}
}

func (r *mutationResolver) CreateProduct(ctx context.Context, request models.CreateProductRequest) (*models.CreateProductResponse, error) {
	productId, err := r.ProductUC.CreateProduct(ctx, toCreateProduct(&request))
	if err != nil {
		return nil, err
	}

	return toCreateProductResponse(productId), nil
}

func (r *queryResolver) GetProductsByFilter(ctx context.Context, request models.GetProductsRequest) (*models.GetProductsResponse, error) {
	products, err := r.ProductUC.GetProductsByFilter(ctx, toGetProducts(&request))
	if err != nil {
		return nil, err
	}

	return toGetProductsResponse(products), nil
}

func (r *ProductResolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *ProductResolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *ProductResolver }
type queryResolver struct{ *ProductResolver }
