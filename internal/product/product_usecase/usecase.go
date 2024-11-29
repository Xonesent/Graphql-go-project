package product_usecase

import (
	"context"
	"graphql/config"
	"graphql/internal/models"
)

type ProductUseCase struct {
	cfg             *config.Config
	productPsqlRepo ProductPsqlRepo
	productMngRepo  ProductMngRepo
}

func NewProductUseCase(
	cfg *config.Config,
	productPsqlRepo ProductPsqlRepo,
	productMngRepo ProductMngRepo,
) *ProductUseCase {
	return &ProductUseCase{
		cfg:             cfg,
		productPsqlRepo: productPsqlRepo,
		productMngRepo:  productMngRepo,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, createProductParams *CreateProduct) (models.ProductId, error) {
	attributesId, err := u.productMngRepo.AddProductAttributes(ctx, createProductParams.Attributes)
	if err != nil {
		return -1, err
	}

	productId, err := u.productPsqlRepo.CreateProduct(ctx, createProductParams.toCreateProduct(attributesId))
	if err != nil {
		return -1, err
	}

	return productId, nil
}

func (u *ProductUseCase) GetProductsByFilter(ctx context.Context, productParams *ProductsFilter) ([]models.FullProduct, error) {
	products, err := u.productPsqlRepo.GetProductsByFilter(ctx, productParams.toGetProductsByFilter())
	if err != nil {
		return nil, err
	}

	attributesIds, attributesMap := GetProductAttributes(products)

	filteredAttributes, err := u.productMngRepo.GetProductAttributesByFilter(ctx, productParams.PrepareMngFilter(attributesIds))
	if err != nil {
		return nil, err
	}

	filteredProducts := make([]models.FullProduct, 0)

	for _, attributes := range filteredAttributes {
		filteredProducts = append(filteredProducts, attributesMap[attributes.ProductId].ToFullProduct(attributes.Attributes))
	}

	return filteredProducts, nil
}
