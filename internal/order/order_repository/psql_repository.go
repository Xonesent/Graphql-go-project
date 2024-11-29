package order_repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"graphql/config"
	"graphql/internal/db_store"
	"graphql/internal/models"
)

type OrderPsqlRepository struct {
	cfg    *config.Config
	psqlDB *pgxpool.Pool
}

func NewOrderPsqlRepository(cfg *config.Config, psqlDB *pgxpool.Pool) *OrderPsqlRepository {
	return &OrderPsqlRepository{
		cfg:    cfg,
		psqlDB: psqlDB,
	}
}

func (r *OrderPsqlRepository) CreateOrder(ctx context.Context, orderParams *CreateOrder) (models.OrderId, error) {
	query, args, err := sq.Insert(db_store.OrdersTableName).
		Columns(db_store.InsertOrdersColumns...).
		Values(
			orderParams.UserId,
			orderParams.ProductId,
			orderParams.Price,
		).
		Suffix("RETURNING " + db_store.OrderIdColumnName).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return -1, err
	}

	var orderId models.OrderId

	if err := r.psqlDB.QueryRow(ctx, query, args...).Scan(&orderId); err != nil {
		return -1, err
	}

	return orderId, nil
}

func (r *OrderPsqlRepository) GetOrdersByFilter(ctx context.Context, orderFilter *GetOrdersByFilter) ([]models.Order, error) {
	query, args, err := sq.Select(db_store.OrdersColumns...).
		From(db_store.OrdersTableName).
		Where(getOrderFilter(orderFilter)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var orders []models.Order

	rows, err := r.psqlDB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderId, &order.UserId, &order.ProductId, &order.Price); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return orders, nil
}

func getOrderFilter(orderFilter *GetOrdersByFilter) sq.And {
	filter := sq.And{}

	if len(orderFilter.OrderIds) != 0 {
		filter = append(filter, sq.Eq{db_store.OrderIdColumnName: orderFilter.OrderIds})
	}

	if len(orderFilter.UserIds) != 0 {
		filter = append(filter, sq.Eq{db_store.UserIdColumnName: orderFilter.UserIds})
	}

	if len(orderFilter.ProductIds) != 0 {
		filter = append(filter, sq.Eq{db_store.ProductItemColumnName: orderFilter.ProductIds})
	}

	if len(orderFilter.Prices) != 0 {
		filter = append(filter, sq.Eq{db_store.OrderPriceColumnName: orderFilter.Prices})
	}

	return filter
}
