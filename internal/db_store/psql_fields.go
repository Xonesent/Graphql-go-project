package db_store

const (
	UsersTableName                = "shop.users"
	ProductsTableName             = "shop.products"
	OrdersTableName               = "shop.orders"
	UserIdColumnName              = "user_id"
	UserNameColumnName            = "name"
	ProductIdColumnName           = "product_id"
	ProductItemColumnName         = "item"
	ProductAttributesIdColumnName = "attributes_id"
	OrderIdColumnName             = "order_id"
	OrderPriceColumnName          = "price"
)

var (
	UserColumns = []string{
		UserIdColumnName,
		UserNameColumnName,
	}
	ProductColumns = []string{
		ProductIdColumnName,
		ProductItemColumnName,
		ProductAttributesIdColumnName,
	}
	OrdersColumns = []string{
		OrderIdColumnName,
		UserIdColumnName,
		ProductIdColumnName,
		OrderPriceColumnName,
	}
	InsertOrdersColumns = []string{
		UserIdColumnName,
		ProductIdColumnName,
		OrderPriceColumnName,
	}
)
