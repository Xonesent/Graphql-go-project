package product_http

import (
	"github.com/gofiber/fiber/v2"
	"graphql/config"
	"graphql/internal/product/product_delivery"
	reqvalidator "graphql/pkg/tools/validator"
)

type ProductHandler struct {
	cfg       *config.Config
	ProductUC product_delivery.ProductUC
}

func NewProductHandler(cfg *config.Config, ProductUC product_delivery.ProductUC) *ProductHandler {
	return &ProductHandler{
		cfg:       cfg,
		ProductUC: ProductUC,
	}
}

func (h *ProductHandler) CreateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		createProduct := createProductRequest{}
		if err := reqvalidator.ReadRequest(c, &createProduct); err != nil {
			return err
		}

		productId, err := h.ProductUC.CreateProduct(c.Context(), createProduct.toCreateProduct())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"product_id": productId,
		})
	}
}

func (h *ProductHandler) GetProductsByFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		productFilter := productFilterRequest{}
		if err := reqvalidator.ReadRequest(c, &productFilter); err != nil {
			return err
		}

		products, err := h.ProductUC.GetProductsByFilter(c.Context(), productFilter.toCreateProduct())
		if err != nil {
			return err
		}

		return c.JSON(products)
	}
}
