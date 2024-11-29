package order_delivery

import (
	"context"
	"graphql/internal/models"
	"graphql/internal/order/order_usecase"
)

type OrderUC interface {
	CreateOrder(ctx context.Context, createOrderParams *order_usecase.CreateOrder) (models.OrderId, error)
	GetOrdersByFilter(ctx context.Context, orderFilter *order_usecase.GetOrdersByFilter) ([]models.Order, error)
}
