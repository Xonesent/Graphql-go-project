package product_graphql

import (
	"graphql/internal/models"
	"graphql/internal/product/product_usecase"
)

func toCreateProduct(request *models.CreateProductRequest) *product_usecase.CreateProduct {
	return &product_usecase.CreateProduct{
		Item:       request.Item,
		Attributes: request.Attributes,
	}
}

func toCreateProductResponse(productId models.ProductId) *models.CreateProductResponse {
	return &models.CreateProductResponse{
		ProductID: productId,
	}
}

func toGetProducts(request *models.GetProductsRequest) *product_usecase.ProductsFilter {
	return &product_usecase.ProductsFilter{
		ProductIds: request.ProductIds,
		Items:      request.Items,
		Attributes: request.Attributes,
	}
}

func toGetProductsResponse(products []models.FullProduct) *models.GetProductsResponse {
	return &models.GetProductsResponse{
		Products: products,
	}
}
