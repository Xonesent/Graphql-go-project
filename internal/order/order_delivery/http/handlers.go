package order_http

import (
	"github.com/gofiber/fiber/v2"
	"graphql/config"
	"graphql/internal/order/order_delivery"
	reqvalidator "graphql/pkg/tools/validator"
	"strconv"
)

type OrderHandler struct {
	cfg     *config.Config
	orderUC order_delivery.OrderUC
}

func NewOrderHandler(cfg *config.Config, orderUC order_delivery.OrderUC) *OrderHandler {
	return &OrderHandler{
		cfg:     cfg,
		orderUC: orderUC,
	}
}

func (h *OrderHandler) CreateOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		createOrder := createOrderRequest{}
		if err := reqvalidator.ReadRequest(c, &createOrder); err != nil {
			return err
		}

		productId, err := strconv.Atoi(c.Params("product_id"))
		if err != nil {
			return err
		}

		userId, err := strconv.Atoi(c.Get("user_id"))
		if err != nil {
			return err
		}

		orderId, err := h.orderUC.CreateOrder(c.Context(), createOrder.toCreateOrder(userId, productId))
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"order_id": orderId,
		})
	}
}

func (h *OrderHandler) GetOrdersByFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		getOrder := getOrdersRequest{}
		if err := reqvalidator.ReadRequest(c, &getOrder); err != nil {
			return err
		}

		orders, err := h.orderUC.GetOrdersByFilter(c.Context(), getOrder.toGetOrdersByFilter())
		if err != nil {
			return err
		}

		return c.JSON(orders)
	}
}
