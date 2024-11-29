package order_usecase

import (
	"context"
	"graphql/config"
	"graphql/internal/models"
)

type OrderUseCase struct {
	cfg           *config.Config
	orderPsqlRepo OrderPsqlRepo
}

func NewOrderUseCase(
	cfg *config.Config,
	orderPsqlRepo OrderPsqlRepo,
) *OrderUseCase {
	return &OrderUseCase{
		cfg:           cfg,
		orderPsqlRepo: orderPsqlRepo,
	}
}

func (u *OrderUseCase) CreateOrder(ctx context.Context, createOrderParams *CreateOrder) (models.OrderId, error) {
	orderId, err := u.orderPsqlRepo.CreateOrder(ctx, createOrderParams.toCreateOrder())
	if err != nil {
		return -1, err
	}

	return orderId, nil
}

func (u *OrderUseCase) GetOrdersByFilter(ctx context.Context, orderFilter *GetOrdersByFilter) ([]models.Order, error) {
	orders, err := u.orderPsqlRepo.GetOrdersByFilter(ctx, orderFilter.toGetOrdersByFilter())
	if err != nil {
		return nil, err
	}

	return orders, nil
}
