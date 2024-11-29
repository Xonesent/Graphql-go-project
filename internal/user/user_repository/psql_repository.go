package user_repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"graphql/config"
	"graphql/internal/db_store"
	"graphql/internal/models"
)

type UserPsqlRepository struct {
	cfg    *config.Config
	psqlDB *pgxpool.Pool
}

func NewUserPsqlRepository(cfg *config.Config, psqlDB *pgxpool.Pool) *UserPsqlRepository {
	return &UserPsqlRepository{
		cfg:    cfg,
		psqlDB: psqlDB,
	}
}

func (r *UserPsqlRepository) CreateUser(ctx context.Context, userParams *AddUser) (models.UserId, error) {
	query, args, err := sq.Insert(db_store.UsersTableName).
		Columns(db_store.UserNameColumnName).
		Values(userParams.Name).
		Suffix("RETURNING " + db_store.UserIdColumnName).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return -1, err
	}

	var userId models.UserId

	if err := r.psqlDB.QueryRow(ctx, query, args...).Scan(&userId); err != nil {
		return -1, err
	}

	return userId, nil
}

func (r *UserPsqlRepository) GetUsersByFilter(ctx context.Context, userFilter *GetUserByFilter) ([]models.User, error) {
	query, args, err := sq.Select(db_store.UserColumns...).
		From(db_store.UsersTableName).
		Where(getUserFilter(userFilter)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var users []models.User

	rows, err := r.psqlDB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}

func getUserFilter(userFilter *GetUserByFilter) sq.And {
	filter := sq.And{}

	if len(userFilter.UserIds) != 0 {
		filter = append(filter, sq.Eq{
			db_store.UserIdColumnName: userFilter.UserIds,
		})
	}

	if len(userFilter.Names) != 0 {
		filter = append(filter, sq.Eq{
			db_store.UserNameColumnName: userFilter.Names,
		})
	}

	return filter
}
