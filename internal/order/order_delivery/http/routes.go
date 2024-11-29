package order_http

import "github.com/gofiber/fiber/v2"

type OrderHDL interface {
	CreateOrder() fiber.Handler
	GetOrdersByFilter() fiber.Handler
}

func MapOrderRoutes(group fiber.Router, h OrderHDL) {
	group.Post(":product_id", h.CreateOrder())
	group.Get("", h.GetOrdersByFilter())
}
