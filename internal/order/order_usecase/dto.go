package order_usecase

import (
	"graphql/internal/models"
	"graphql/internal/order/order_repository"
)

type CreateOrder struct {
	UserId    models.UserId
	ProductId models.ProductId
	Price     int
}

func (r *CreateOrder) toCreateOrder() *order_repository.CreateOrder {
	return &order_repository.CreateOrder{
		UserId:    r.UserId,
		ProductId: r.ProductId,
		Price:     r.Price,
	}
}

type GetOrdersByFilter struct {
	OrderIds   []models.OrderId
	UserIds    []models.UserId
	ProductIds []models.ProductId
	Prices     []int
}

func (r *GetOrdersByFilter) toGetOrdersByFilter() *order_repository.GetOrdersByFilter {
	return &order_repository.GetOrdersByFilter{
		OrderIds:   r.OrderIds,
		UserIds:    r.UserIds,
		ProductIds: r.ProductIds,
		Prices:     r.Prices,
	}
}
