package models

type Order struct {
	OrderId   OrderId   `json:"order_id"`
	UserId    UserId    `json:"user_id"`
	ProductId ProductId `json:"product_id"`
	Price     int32     `bson:"price"`
}
