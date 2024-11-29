package product_psql_repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"graphql/config"
	"graphql/internal/db_store"
	"graphql/internal/models"
)

type ProductPsqlRepository struct {
	cfg    *config.Config
	psqlDB *pgxpool.Pool
}

func NewProductPsqlRepository(cfg *config.Config, psqlDB *pgxpool.Pool) *ProductPsqlRepository {
	return &ProductPsqlRepository{
		cfg:    cfg,
		psqlDB: psqlDB,
	}
}

func (r *ProductPsqlRepository) CreateProduct(ctx context.Context, productParams *CreateProduct) (models.ProductId, error) {
	query, args, err := sq.Insert(db_store.ProductsTableName).
		Columns(db_store.ProductItemColumnName, db_store.ProductAttributesIdColumnName).
		Values(productParams.Item, productParams.AttributeId).
		Suffix("RETURNING " + db_store.ProductIdColumnName).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return -1, err
	}

	var productId models.ProductId

	if err := r.psqlDB.QueryRow(ctx, query, args...).Scan(&productId); err != nil {
		return -1, err
	}

	return productId, nil
}

func (r *ProductPsqlRepository) GetProductsByFilter(ctx context.Context, productFilter *GetProductsByFilter) ([]models.Product, error) {
	query, args, err := sq.Select(db_store.ProductColumns...).
		From(db_store.ProductsTableName).
		Where(getProductFilter(productFilter)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var products []models.Product

	rows, err := r.psqlDB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ProductId, &product.Item, &product.Attributes); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return products, nil
}

func getProductFilter(productFilter *GetProductsByFilter) sq.And {
	filter := sq.And{}

	if len(productFilter.ProductIds) != 0 {
		filter = append(filter, sq.Eq{
			db_store.ProductIdColumnName: productFilter.ProductIds,
		})
	}

	if len(productFilter.Items) != 0 {
		filter = append(filter, sq.Eq{
			db_store.ProductItemColumnName: productFilter.Items,
		})
	}

	return filter
}
