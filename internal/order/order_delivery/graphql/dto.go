package order_graphql

import (
	"graphql/internal/models"
	"graphql/internal/order/order_usecase"
)

func toCreateProduct(request models.CreateOrderRequest, userId models.UserId) *order_usecase.CreateOrder {
	return &order_usecase.CreateOrder{
		UserId:    userId,
		ProductId: request.ProductID,
		Price:     request.Price,
	}
}

func toCreateOrderResponse(orderId models.OrderId) *models.CreateOrderResponse {
	return &models.CreateOrderResponse{
		OrderID: orderId,
	}
}

func toGetOrdersByFilter(request models.GetOrdersRequest) *order_usecase.GetOrdersByFilter {
	return &order_usecase.GetOrdersByFilter{
		OrderIds:   request.OrderIds,
		UserIds:    request.UserIds,
		ProductIds: request.ProductIds,
		Prices:     request.Prices,
	}
}

func toGetOrdersResponse(orders []models.Order) *models.GetOrdersResponse {
	return &models.GetOrdersResponse{
		Orders: orders,
	}
}
