package user_http

import (
	"github.com/gofiber/fiber/v2"
	"graphql/config"
	"graphql/internal/user/user_delivery"
	reqvalidator "graphql/pkg/tools/validator"
)

type UserHandler struct {
	cfg    *config.Config
	userUC user_delivery.UserUC
}

func NewUserHandler(cfg *config.Config, userUC user_delivery.UserUC) *UserHandler {
	return &UserHandler{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (h *UserHandler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		registerUser := registerUserRequest{}
		if err := reqvalidator.ReadRequest(c, &registerUser); err != nil {
			return err
		}

		userId, err := h.userUC.RegisterUser(c.Context(), registerUser.toRegisterUser())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"user_id": userId,
		})
	}
}

func (h *UserHandler) GetUsersByFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getUser := getUserRequest{}
		if err := reqvalidator.ReadRequest(c, &getUser); err != nil {
			return err
		}

		users, err := h.userUC.GetUserByFilter(c.Context(), getUser.toGetUserByFilter())
		if err != nil {
			return err
		}

		return c.JSON(users)
	}
}
