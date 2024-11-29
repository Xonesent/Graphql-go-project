//go:generate go run github.com/99designs/gqlgen generate
package order_graphql

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"graphql/config"
	"graphql/internal/models"
	"graphql/internal/order/order_delivery"
	"strconv"
)

type OrderResolver struct {
	cfg     *config.Config
	orderUC order_delivery.OrderUC
	ch      chan *models.Order
}

func NewOrderResolver(cfg *config.Config, orderUC order_delivery.OrderUC) *OrderResolver {
	return &OrderResolver{
		cfg:     cfg,
		orderUC: orderUC,
		ch:      make(chan *models.Order),
	}
}

func (r *mutationResolver) CreateOrder(ctx context.Context, request models.CreateOrderRequest) (*models.CreateOrderResponse, error) {
	userIdHeader := graphql.GetOperationContext(ctx).Headers["User_id"]
	if userIdHeader == nil || len(userIdHeader) == 0 {
		return nil, errors.New("userId header is required")
	}

	userId, err := strconv.Atoi(userIdHeader[0])
	if err != nil {
		return nil, err
	}

	orderId, err := r.orderUC.CreateOrder(ctx, toCreateProduct(request, models.UserId(userId)))
	if err != nil {
		return nil, err
	}

	r.ch <- &models.Order{
		OrderId:   orderId,
		UserId:    models.UserId(userId),
		ProductId: request.ProductID,
		Price:     int32(request.Price),
	}

	return toCreateOrderResponse(orderId), nil
}

func (r *queryResolver) GetOrdersByFilter(ctx context.Context, request models.GetOrdersRequest) (*models.GetOrdersResponse, error) {
	orders, err := r.orderUC.GetOrdersByFilter(ctx, toGetOrdersByFilter(request))
	if err != nil {
		return nil, err
	}

	return toGetOrdersResponse(orders), nil
}

func (r *subscriptionResolver) MonitorOrders(ctx context.Context) (<-chan *models.Order, error) {
	return r.ch, nil
}

func ValidateProdutId(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println(obj)

	return next(ctx)
}

func (r *OrderResolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *OrderResolver) Query() QueryResolver { return &queryResolver{r} }

func (r *OrderResolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *OrderResolver }
type queryResolver struct{ *OrderResolver }
type subscriptionResolver struct{ *OrderResolver }
