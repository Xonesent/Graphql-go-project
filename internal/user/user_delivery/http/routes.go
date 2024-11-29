package user_http

import "github.com/gofiber/fiber/v2"

type UserHDL interface {
	RegisterUser() fiber.Handler
	GetUsersByFilter() fiber.Handler
}

func MapUserRoutes(group fiber.Router, h UserHDL) {
	group.Post("/register", h.RegisterUser())
	group.Get("", h.GetUsersByFilter())
}
