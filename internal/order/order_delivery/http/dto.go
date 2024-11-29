package order_http

import (
	"graphql/internal/models"
	"graphql/internal/order/order_usecase"
)

type createOrderRequest struct {
	Price int `json:"price"`
}

func (r *createOrderRequest) toCreateOrder(userId, productId int) *order_usecase.CreateOrder {
	return &order_usecase.CreateOrder{
		UserId:    models.UserId(userId),
		ProductId: models.ProductId(productId),
		Price:     r.Price,
	}
}

type getOrdersRequest struct {
	OrderIds   []models.OrderId   `json:"order_ids"`
	UserIds    []models.UserId    `json:"user_ids"`
	ProductIds []models.ProductId `json:"product_ids"`
	Prices     []int              `json:"prices"`
}

func (r *getOrdersRequest) toGetOrdersByFilter() *order_usecase.GetOrdersByFilter {
	return &order_usecase.GetOrdersByFilter{
		OrderIds:   r.OrderIds,
		UserIds:    r.UserIds,
		ProductIds: r.ProductIds,
		Prices:     r.Prices,
	}
}
