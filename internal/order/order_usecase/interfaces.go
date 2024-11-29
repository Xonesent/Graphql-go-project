package order_usecase

import (
	"context"
	"graphql/internal/models"
	"graphql/internal/order/order_repository"
)

type OrderPsqlRepo interface {
	CreateOrder(ctx context.Context, orderParams *order_repository.CreateOrder) (models.OrderId, error)
	GetOrdersByFilter(ctx context.Context, orderFilter *order_repository.GetOrdersByFilter) ([]models.Order, error)
}
