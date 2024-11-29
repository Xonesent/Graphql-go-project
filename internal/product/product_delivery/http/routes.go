package product_http

import "github.com/gofiber/fiber/v2"

type ProductHDL interface {
	CreateProduct() fiber.Handler
	GetProductsByFilter() fiber.Handler
}

func MapProductRoutes(group fiber.Router, h ProductHDL) {
	group.Post("", h.CreateProduct())
	group.Get("", h.GetProductsByFilter())
}
