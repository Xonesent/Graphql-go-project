package order_repository

import "graphql/internal/models"

type CreateOrder struct {
	UserId    models.UserId
	ProductId models.ProductId
	Price     int
}

type GetOrdersByFilter struct {
	OrderIds   []models.OrderId
	UserIds    []models.UserId
	ProductIds []models.ProductId
	Prices     []int
}
